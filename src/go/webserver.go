// internal/utils/webserver.go
package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// WebhookPayload representa a estrutura básica de um payload de webhook recebido
type WebhookPayload struct {
	Timestamp string                 `json:"timestamp"`
	Method    string                 `json:"method"`
	Path      string                 `json:"path"`
	Headers   map[string][]string    `json:"headers"`
	Body      json.RawMessage        `json:"body"`
	Query     map[string][]string    `json:"query"`
}

// WebhookEvent representa um evento de webhook com a URL que o disparou
type WebhookEvent struct {
	URL     string
	Payload WebhookPayload
}

// WebhookServer gerencia um servidor HTTP para receber webhooks
type WebhookServer struct {
	server      *http.Server
	eventCh     chan WebhookEvent
	listenAddr  string
	routes      map[string]struct{} // Caminhos que estamos monitorando
	mu          sync.Mutex          // Protege o mapa de rotas
}

// NewWebhookServer cria uma nova instância de WebhookServer
func NewWebhookServer(listenAddr string) *WebhookServer {
	return &WebhookServer{
		eventCh:    make(chan WebhookEvent, 100), // Buffer de eventos
		listenAddr: listenAddr,
		routes:     make(map[string]struct{}),
	}
}

// AddRoute adiciona uma rota (caminho) para o servidor monitorar
func (ws *WebhookServer) AddRoute(path string) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.routes[path] = struct{}{}
	fmt.Printf("Webserver: Monitorando rota de webhook: %s%s\n", ws.listenAddr, path)
}

// Start inicia o servidor HTTP em uma goroutine
func (ws *WebhookServer) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ws.mu.Lock()
		_, isMonitoredRoute := ws.routes[r.URL.Path]
		ws.mu.Unlock()

		if !isMonitoredRoute {
			http.NotFound(w, r)
			return
		}

		fmt.Printf("Webserver: Recebido webhook em %s %s\n", r.Method, r.URL.Path)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Erro ao ler corpo da requisição", http.StatusInternalServerError)
			return
		}

		payload := WebhookPayload{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
			Method:    r.Method,
			Path:      r.URL.Path,
			Headers:   r.Header,
			Body:      body,
			Query:     r.URL.Query(),
		}

		// Envia o evento para o canal
		select {
		case ws.eventCh <- WebhookEvent{URL: r.URL.String(), Payload: payload}:
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "Webhook recebido com sucesso!")
		case <-time.After(1 * time.Second): // Timeout para não bloquear o handler
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, "Servidor ocupado, tente novamente.")
		}
	})

	ws.server = &http.Server{
		Addr:    ws.listenAddr,
		Handler: mux,
	}

	go func() {
		fmt.Printf("Webserver: Iniciando servidor de webhook em %s\n", ws.listenAddr)
		if err := ws.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "Webserver: Erro ao iniciar servidor HTTP: %v\n", err)
		}
	}()
}

// Stop encerra o servidor HTTP
func (ws *WebhookServer) Stop(ctx context.Context) error {
	fmt.Println("Webserver: Encerrando servidor de webhook...")
	close(ws.eventCh) // Fecha o canal para sinalizar o fim dos eventos
	return ws.server.Shutdown(ctx)
}

// Events retorna o canal de eventos de webhook
func (ws *WebhookServer) Events() <-chan WebhookEvent {
	return ws.eventCh
}
