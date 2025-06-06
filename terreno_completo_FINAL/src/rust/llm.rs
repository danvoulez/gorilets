// Handlers de LLM (summarize, plan)
use actix_web::{post, web, HttpResponse, Responder};

#[post("/api/llm/summary")]
pub async fn summary() -> impl Responder {
    HttpResponse::Ok().body("LLM summary handler")
}
