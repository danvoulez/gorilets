use actix_files::Files;
use actix_web::{
    get, post,
    web::{self, Data},
    App, HttpResponse, HttpServer, Responder,
};
use std::sync::Arc;
use tokio::sync::Mutex;
use serde::Deserialize;
use sqlx::{Pool, Sqlite};
use std::time::SystemTime;

mod handlers;
mod logline_runner;
mod utils;

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    // Inicializa logger
    env_logger::init();

    // Cria pool SQLite (para protótipo; em produção use Postgres)
    let pool = Pool::<Sqlite>::connect("sqlite://minicontratos.db").await.unwrap();

    // Se quiser rodar migrações (usando sqlx-cli)
    // sqlx::migrate!("./migrations").run(&pool).await.unwrap();

    // Estado compartilhado
    let app_state = Data::new(AppState {
        db: Arc::new(Mutex::new(pool)),
    });

    HttpServer::new(move || {
        App::new()
            .app_data(app_state.clone())
            // Serve arquivos estáticos do frontend (WASM, CSS, HTML)
            .service(Files::new("/static", "./frontend/static").show_files_listing())
            // Endpoints da UI (via WASM)
            .service(Files::new("/contracts/ui", "./contracts/ui"))
            // Endpoints:
            .service(handlers::login_page)
            .service(handlers::login_api)
            .service(handlers::contracts_list)
            .service(handlers::llm_summary)
    })
    .bind(("127.0.0.1", 8080))?
    .run()
    .await
}

pub struct AppState {
    // Pool de DB protegido por Mutex para concorrência simples
    db: Arc<Mutex<Pool<Sqlite>>>,
}
