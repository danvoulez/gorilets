// Handlers de WhatsApp (listar conversas, enviar mensagem)
use actix_web::{post, web, HttpResponse, Responder};

#[post("/api/whatsapp/send")]
pub async fn send() -> impl Responder {
    HttpResponse::Ok().body("WhatsApp send handler")
}
