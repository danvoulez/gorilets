// internal/snapshotter/snapshotter.go
package snapshotter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	errorsInternal "loglineos/internal/errors"
	"loglineos/internal/executor" // Importa o executor para EntityRecord
	"loglineos/internal/spans"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/google/uuid"
)

// SnapshotMetadata contém informações sobre o snapshot
type SnapshotMetadata struct {
	SnapshotID  string                 `json:"snapshot_id"`
	Timestamp   string                 `json:"timestamp"`
	Goal        string                 `json:"goal"`
	Message     string                 `json:"message,omitempty"`
	Inputs      []string               `json:"inputs_committed,omitempty"`
	SignedBy    []string               `json:"signed_by,omitempty"`
	SpansCount  int                    `json:"spans_count"`
	EntitiesCount int                  `json:"entities_count"`
	GitHash     string                 `json:"git_hash,omitempty"`
	GitBranch   string                 `json:"git_branch,omitempty"`
	Signed      bool                   `json:"signed,omitempty"`
	Signature   string                 `json:"signature_file,omitempty"` // Caminho para o arquivo .asc ou base64 da assinatura
	Context     map[string]interface{} `json:"context,omitempty"`
}

// CreateSnapshot gera um novo snapshot do estado atual
func CreateSnapshot(
	goal string,
	inputs []string,
	signedBy []string,
	message string,
	globalSpansFilePath string, // Caminho para o arquivo spans.jsonl principal
	entityStore map[string]executor.EntityRecord, // EntityStore do executor
	baseOutputDir string, // Diretório base onde os snapshots serão criados (e.g., ./data)
	signGPG bool, // Flag para assinar com GPG
) (string, error) {
	
	// 1. Gerar ID e caminho para o snapshot
	snapshotID := uuid.New().String()
	snapshotDir := filepath.Join(baseOutputDir, "snapshots", snapshotID)

	if err := os.MkdirAll(snapshotDir, 0755); err != nil {
		return "", fmt.Errorf("snapshotter: falha ao criar diretório de snapshot '%s': %w", snapshotDir, err)
	}
	fmt.Printf("   Criado diretório de snapshot: %s\n", snapshotDir)

	// 2. Coletar e filtrar spans relevantes
	allSpans, err := spans.ReadAllSpans(globalSpansFilePath)
	if err != nil {
		return "", fmt.Errorf("snapshotter: falha ao ler spans globais: %w", err)
	}

	filteredSpans := []spans.Span{}
	
	inputMap := make(map[string]bool)
	for _, id := range inputs {
		inputMap[id] = true
	}

	for _, s := range allSpans {
		if inputMap[s.TaskID] || inputMap[s.SpanID] {
			filteredSpans = append(filteredSpans, s)
		}
	}
	
	spansSnapshotPath := filepath.Join(snapshotDir, "spans_snapshot.jsonl")
	writer, err := spans.NewJSONLWriter(spansSnapshotPath)
	if err != nil {
		return "", fmt.Errorf("snapshotter: falha ao criar escritor de spans para snapshot: %w", err)
	}
	defer writer.Close()
	for _, s := range filteredSpans {
		if err := writer.WriteSpan(s); err != nil {
			return "", fmt.Errorf("snapshotter: falha ao escrever span no snapshot: %w", err)
		}
	}
	fmt.Printf("   Salvos %d spans relevantes em: %s\n", len(filteredSpans), spansSnapshotPath)

	// Salva as entidades registradas do EntityStore
	entitiesSlice := make([]executor.EntityRecord, 0, len(entityStore))
	for _, entity := range entityStore {
		entitiesSlice = append(entitiesSlice, entity)
	}

	entitiesSnapshotPath := filepath.Join(snapshotDir, "entities_snapshot.json")
	entitiesData, err := json.MarshalIndent(entitiesSlice, "", "  ")
	if err != nil {
		return "", fmt.Errorf("snapshotter: falha ao serializar entidades: %w", err)
	}
	if err := ioutil.WriteFile(entitiesSnapshotPath, entitiesData, 0644); err != nil {
		return "", fmt.Errorf("snapshotter: falha ao escrever entidades no snapshot: %w", err)
	}
	fmt.Printf("   Salvas %d entidades registradas em: %s\n", len(entitiesSlice), entitiesSnapshotPath)


	// 3. Gerar metadados do snapshot
	metadata := SnapshotMetadata{
		SnapshotID:  snapshotID,
		Timestamp:   time.Now().UTC().Format(time.RFC3339),
		Goal:        goal,
		Message:     message,
		Inputs:      inputs,
		SignedBy:    signedBy,
		SpansCount:  len(filteredSpans),
		EntitiesCount: len(entitiesSlice),
		Context: map[string]interface{}{
			"logline_version": "v0.4", // Atualiza a versão
		},
	}
	metadataPath := filepath.Join(snapshotDir, "metadata.json")
	metadataBytes, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return "", fmt.Errorf("snapshotter: falha ao serializar metadados do snapshot: %w", err)
	}
	if err := ioutil.WriteFile(metadataPath, metadataBytes, 0644); err != nil {
		return "", fmt.Errorf("snapshotter: falha ao escrever metadados do snapshot: %w", err)
	}
	fmt.Printf("   Metadados do snapshot salvos em: %s\n", metadataPath)


	// 4. Integração com Git Real (`go-git`)
	repoPath := filepath.Clean(filepath.Join(baseOutputDir, ".."))

	r, err := git.PlainOpen(repoPath)
	if err != nil {
		fmt.Printf("   [INFO] Repositório Git não encontrado ou erro ao abrir (%v). Pulando integração Git.\n", err)
	} else {
		fmt.Println("   [INFO] Integrando com Git REAL...")
		w, err := r.Worktree()
		if err != nil {
			return "", fmt.Errorf("snapshotter: falha ao obter Worktree Git: %w", err)
		}

		relSnapshotPath, err := filepath.Rel(repoPath, snapshotDir)
		if err != nil {
			return "", fmt.Errorf("snapshotter: falha ao obter caminho relativo do snapshot: %w", err)
		}
		_, err = w.Add(relSnapshotPath)
		if err != nil {
			return "", fmt.Errorf("snapshotter: falha ao adicionar snapshot ao Git staging: %w", err)
		}
		fmt.Printf("   Git: Adicionado '%s' ao staging.\n", relSnapshotPath)

		commitMsg := fmt.Sprintf("LogLine Commit: %s (Snapshot ID: %s)", goal, snapshotID)
		if message != "" {
			commitMsg = fmt.Sprintf("%s\n\n%s", commitMsg, message)
		}

		authorName := "LogLineOS Automated Commit"
		authorEmail := "loglineos@example.com"
		if len(signedBy) > 0 {
			authorName = strings.Join(signedBy, ", ") + " (via LogLineOS)"
		}

		commitHash, err := w.Commit(commitMsg, &git.CommitOptions{
			Author: &object.Signature{
				Name:  authorName,
				Email: authorEmail,
				When:  time.Now(),
			},
		})
		if err != nil {
			return "", fmt.Errorf("snapshotter: falha ao criar commit Git: %w", err)
		}
		fmt.Printf("   Git: Commit criado com sucesso: %s\n", commitHash.String())

		metadata.GitHash = commitHash.String()
		head, err := r.Head()
		if err == nil {
			metadata.GitBranch = head.Name().Short()
		}
		metadataBytesUpdated, err := json.MarshalIndent(metadata, "", "  ")
		if err == nil {
			ioutil.WriteFile(metadataPath, metadataBytesUpdated, 0644)
		}
	}

	// 5. Assinatura Digital GPG Real (via exec.Command)
	if signGPG {
		fmt.Println("   [INFO] Tentando assinatura GPG REAL...")
		cmd := exec.Command("gpg", "--batch", "--yes", "--armor", "--detach-sign", "--output", metadataPath+".asc", metadataPath)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		
		err := cmd.Run()
		if err != nil {
			log.Printf("⚠️ GPG_SIGN_FAIL: %v. Verifique se o GPG está instalado, sua chave está configurada e a senha pode ser fornecida (ou use --batch).\n", err)
			metadata.Signed = false
		} else {
			metadata.Signed = true
			metadata.Signature = filepath.Base(metadataPath) + ".asc"
			fmt.Println("   GPG: Metadados do snapshot GPG-assinados com sucesso.")
			metadataBytesUpdated, err := json.MarshalIndent(metadata, "", "  ")
			if err == nil {
				ioutil.WriteFile(metadataPath, metadataBytesUpdated, 0644)
			}
		}
	}

	return snapshotID, nil
}
