use actix_web::{get, post, web, HttpResponse, Responder};
use sqlx::SqlitePool;
use crate::AppState;
use serde::Deserialize;
use std::time::SystemTime;
use crate::logline_runner;
use crate::utils::record_span;

// Página de login: redireciona para o frontend WASM que irá buscar splash.logline
#[get("/login")]
pub async fn login_page() -> impl Responder {
    // Simplesmente devolve o index.html do WASM; o WASM vai carregar /contracts/ui/login.logline
    HttpResponse::Ok()
        .content_type("text/html")
        .body(include_str!("../frontend/static/index.html"))
}

// API de login: recebe JSON (ou form) e responde com fragmento HTML
#[derive(Deserialize)]
pub struct LoginData {
    email: String,
    senha: String,
}

#[post("/api/auth/login")]
pub async fn login_api(
    data: web::Json<LoginData>,
    state: web::Data<AppState>,
) -> impl Responder {
    // Gera span de auditoria
    record_span("auth.login", &format!("user:{}", data.email)).await;

    // Simulação de autenticação (substituir por consulta real a DB)
    if data.email.contains("@") && data.senha.len() >= 6 {
        let html = r#"
        <div class="p-6 bg-green-100 rounded">
          <h2 class="text-xl font-semibold mb-4">Login bem-sucedido!</h2>
          <a href="/dashboard" class="text-blue-600 hover:underline">Ir para o dashboard</a>
        </div>
        "#;
        HttpResponse::Ok().content_type("text/html").body(html)
    } else {
        let html = r#"
        <div class="p-6 bg-red-100 rounded">
          <h2 class="text-xl font-semibold mb-4 text-red-700">Credenciais inválidas</h2>
          <a href="/login" class="text-blue-600 hover:underline">Tentar novamente</a>
        </div>
        "#;
        HttpResponse::Ok().content_type("text/html").body(html)
    }
}

// Listar contratos ativos (JSON ou HTML fragment para HTMX)
#[get("/api/contracts/list")]
pub async fn contracts_list(state: web::Data<AppState>) -> impl Responder {
    let pool = state.db.lock().await;
    let rows = sqlx::query!("SELECT id, title, status FROM contracts")
        .fetch_all(&*pool)
        .await
        .unwrap_or_else(|_| vec![]);

    // Gera fragmento HTML da tabela
    let mut html = String::new();
    html.push_str(r#"<table class="min-w-full bg-gray-50 shadow overflow-hidden rounded-lg">"#);
    html.push_str(r#"<thead><tr><th class="px-4 py-2 border-b">ID</th><th class="px-4 py-2 border-b">Título</th><th class="px-4 py-2 border-b">Status</th></tr></thead><tbody>"#);
    for row in rows {
        html.push_str(&format!(
            r#"<tr><td class="px-4 py-2 border-b">{}</td><td class="px-4 py-2 border-b">{}</td><td class="px-4 py-2 border-b">{}</td></tr>"#,
            row.id, row.title, row.status
        ));
    }
    html.push_str("</tbody></table>");
    HttpResponse::Ok().content_type("text/html").body(html)
}

// Handler LLM: recebe prompt via JSON e devolve JSON com “conteúdo”
#[derive(Deserialize)]
pub struct LlmRequest {
    prompt: String,
}

#[post("/api/llm/summary")]
pub async fn llm_summary(req: web::Json<LlmRequest>) -> impl Responder {
    // Executa Span llm.gateway via função auxiliar
    let summary = logline_runner::run_llm_gateway(&req.prompt).await.unwrap_or_else(|e| e);
    HttpResponse::Ok().json(serde_json::json!({ "summary": summary }))
}
