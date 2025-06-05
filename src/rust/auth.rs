// Handlers de autenticação (login, refresh, revoke)
// Exemplo mínimo
use actix_web::{post, web, HttpResponse, Responder};

#[post("/api/auth/login")]
pub async fn login_api() -> impl Responder {
    HttpResponse::Ok().body("Login handler")
}
