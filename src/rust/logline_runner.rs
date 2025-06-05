use std::time::SystemTime;
use tokio::process::Command;
use serde_json::Value;

// Função que simula “llm.gateway”: tenta local Mistral, senão cloud
pub async fn run_llm_gateway(prompt: &str) -> Result<String, String> {
    // Primeiro, tenta LLM local (Mistral via CLI)
    let output = Command::new("ollama")
        .arg("run")
        .arg("mistral")
        .arg(prompt)
        .output()
        .await
        .map_err(|e| format!("Erro ao executar LLM local: {}", e))?;

    if output.status.success() {
        let resp = String::from_utf8_lossy(&output.stdout).into_owned();
        // Grava span llm.response
        record_span("llm.route", "mistral").await;
        record_span("llm.response", &resp).await;
        return Ok(resp);
    }
    // Se falhou, chama API OpenAI (curl)
    record_span("llm.fallback", "true").await;
    let api_key = std::env::var("OPENAI_API_KEY").unwrap_or_default();
    let body = serde_json::json!({
      "model": "gpt-4",
      "messages": [{ "role": "user", "content": prompt }],
      "temperature": 0.3
    });
    let resp = reqwest::Client::new()
        .post("https://api.openai.com/v1/chat/completions")
        .bearer_auth(api_key)
        .json(&body)
        .send()
        .await
        .map_err(|e| format!("Erro OpenAI: {}", e))?;

    let data: Value = resp.json().await.map_err(|e| format!("Parse JSON falhou: {}", e))?;
    let content = data["choices"][0]["message"]["content"].as_str().unwrap_or("").to_string();
    record_span("llm.route", "gpt-4").await;
    record_span("llm.response", &content).await;
    Ok(content)
}

// Função auxiliar para gravar um span genérico em JSONL (por simplicidade, grava no console)
pub async fn record_span(span_type: &str, message: &str) {
    let now = SystemTime::now();
    let timestamp = chrono::DateTime::<chrono::Utc>::from(now).to_rfc3339();
    let span = format!(r#"{{"timestamp":"{}","type":"{}","message":"{}"}}"#, timestamp, span_type, message);
    println!("{}", span); // Em produção, grave em arquivo .jsonl ou serviço de logs
}
