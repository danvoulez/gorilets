// Handlers de infraestrutura (executar spans de infra)
use actix_web::{post, web, HttpResponse, Responder};

#[post("/api/infra/deploy")]
pub async fn deploy() -> impl Responder {
    HttpResponse::Ok().body("Infra deploy handler")
}
