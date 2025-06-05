// Handlers de contratos (CRUD, execução, versionamento)
use actix_web::{get, post, web, HttpResponse, Responder};

#[get("/api/contracts/list")]
pub async fn list() -> impl Responder {
    HttpResponse::Ok().body("List contracts handler")
}
