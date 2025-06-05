# 📦 LogLineOS: Instruções Finais de Instalação e Execução

Este documento traz o passo a passo final, detalhado, para instalar dependências, compilar, executar e testar o LogLineOS em sua versão premium.

---

## 🛠️ Instalação e Build

1. **Entre no diretório do projeto:**
   ```
   cd loglineos
   ```

2. **Instale TODAS as dependências Go:**
   ```
   go mod tidy
   ```
   - Se houver erros, revise os arquivos `.go` e repita.

3. **Configure o `config/logline.yml`:**
   - Se for usar OpenAI, coloque sua chave real no campo `api_key`.
   - Se NÃO for usar OpenAI, altere `provider: "openai"` para `provider: "mock"` ou `provider: "ollama"`.
   - Para Ollama, garanta que o Ollama esteja rodando localmente e um modelo carregado.

4. **Inicialize o repositório Git (opcional, mas recomendado):**
   ```
   git init .
   git config user.name "LogLineOS User"
   git config user.email "user@loglineos.com"
   echo "data/*" > .gitignore
   echo "!data/snapshots/" >> .gitignore
   git add .gitignore
   git commit -m "Initial LogLineOS repository setup"
   ```

5. **Configure GPG (opcional, para assinatura):**
   - Instale o GnuPG (`brew install gnupg` ou `sudo apt install gnupg`).
   - Gere uma chave se necessário: `gpg --full-generate-key`.

6. **Compile o LogLineOS:**
   ```
   go build -o logline
   ```
   - Se houver erros de build, revise os códigos e repita `go mod tidy`.

7. **Limpe o diretório de dados (opcional):**
   ```
   rm -rf ./data
   mkdir -p ./data
   ```

---

## 🌟 Como Testar (Experiência Premium)

### 1. **Execução de Workflow com LLM**
   ```
   rm -f ./data/spans.jsonl && ./logline run examples/workflows/llm_hook_workflow.logline
   ```
   - Observe hooks, chamadas LLM, disparos condicionais e o arquivo `data/spans.jsonl`.

### 2. **Loops e Contexto Dinâmico**
   ```
   rm -f ./data/spans.jsonl && ./logline run examples/workflows/iterative_process.logline
   ```
   - Verifique spans dinâmicos e variáveis de loop em `data/spans.jsonl`.

### 3. **Monitoramento Contínuo e Webhook**
   ```
   rm -f ./data/spans.jsonl && ./logline watch examples/workflows/file_watcher.logline
   ```
   - Em outro terminal:
     ```
     curl -X POST -H 'Content-Type: application/json' -d '{"status": "DEPLOY_COMPLETE", "version": "v1.0.0"}' http://localhost:8080/webhooks/deploy-status
     ```
   - Veja os eventos e spans gerados.

### 4. **PromptPad - IDE Terminal**
   ```
   rm -f ./data/spans.jsonl && ./logline promptpad
   ```
   - Use comandos como `open`, `simulate`, `run`, `audit`, `commit`, `revert`, `replay`, `? <pergunta>` no PromptPad.
   - Valide, edite e visualize grafos dos workflows, auditando tudo em tempo real.

---

## 💡 Dicas Finais

- **Qualquer erro de dependência ou build:** revise os códigos, salve corretamente e repita `go mod tidy` e `go build`.
- **Para logs e auditoria:** consulte sempre `data/spans.jsonl` e use o comando `logline audit` para relatórios.
- **Explore todos os exemplos em `examples/workflows/`.**

---

**Parabéns! Você está com o LogLineOS pronto para operar em sua melhor versão — auditável, inteligente, visual, integrável e com uma IDE de terminal premium.**

---
