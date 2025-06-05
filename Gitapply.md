 (cd "$(git rev-parse --show-toplevel)" && git apply --3way <<'EOF' 
diff --git a/flipapp_runtime_full/README.md b/flipapp_runtime_full/README.md
new file mode 100644
index 0000000000000000000000000000000000000000..1f15729e256dddba2961df8803f26e17206e5a3e
--- /dev/null
+++ b/flipapp_runtime_full/README.md
@@ -0,0 +1,70 @@
+```markdown
+# FlipApp LogLineOS + WhatsApp Integration
+
+![LogLineOS Logo](https://raw.githubusercontent.com/dally/logline-os/main/logo.svg)
+
+## 🚀 Uma Nova Era de Software: LogLineOS com Módulo WhatsApp
+
+Bem-vindo ao **FlipApp LogLineOS**, o projeto flagship que materializa a visão da linguagem LogLine com **integração completa do WhatsApp**. Este não é apenas um aplicativo; é um **manifesto operacional** onde toda a lógica, interface de usuário, estado, comportamento e até as regras de execução do sistema são descritas de forma **100% declarativa** como fluxos de spans em arquivos `.logline`.
+
+**Este projeto não depende mais de Docker, Nginx, Vite, Jest ou Node.js para construir sua interface.** A única dependência de execução é o binário Go (`flip`) que interpreta spans LogLine, mais eventuais proxies LLM em Python ou WASM.
+
+### ✨ Recursos do FlipApp LogLineOS + WhatsApp
+
+- **Integração WhatsApp Web**: UI pixel-perfect, simulação de áudio, tradução automática, ghost‐sale detection, sistema de julgamento e treinamento contínuo.
+- **Micro‐Kernel Declarativo**: Executor Go/WASM que lê `execution_rule` em LogLine.
+- **Auditabilidade Total & Reversível**: Cada operação é um span persistido.
+- **Offline‐First**: Design robusto sem necessidade de container ou servidor externo para UI.
+
+---
+
+## 🛠️ Pré-requisitos
+
+- **Go 1.20+** para compilar o binário CLI `flip`
+- **Python 3.8+** (opcional) para proxies LLM locais
+- Nenhuma dependência de Node.js, Docker ou bundlers
+
+## 🚀 Como Rodar Localmente
+
+1. Clone o repositório  
+   ```
+   git clone https://github.com/danvoulez/dollars.git
+   cd dollars
+   ```
+
+2. Compile o Executor  
+   ```
+   make build
+   ```
+
+3. Carregue variáveis de ambiente (se tiver `load_env.sh`)  
+   ```
+   source load_env.sh
+   ```
+
+4. Execute o FlipApp  
+   ```
+   ./bin/flip run --watch
+   ```
+
+5. (Opcional) Proxy LLM  
+   ```
+   python3 llm_proxy_hard.py --port 8001
+   ```
+
+---
+
+## 🧪 Testes & Validações
+
+- `make validate` – sintaxe de todos os spans `.logline` e `.jsonl`  
+- `make test` – testes LogLine  
+- Não há dependência de Jest, Vite ou Playwright
+
+---
+
+## 📚 Mais
+
+- CHANGELOG.md  
+- CONTRIBUTING.md  
+- CHEATSHEET.md  
+- LICENSE  
\ No newline at end of file
diff --git a/flipapp_runtime_full/app.js b/flipapp_runtime_full/app.js
new file mode 100644
index 0000000000000000000000000000000000000000..e040ebc4751856440506d1733c452e3d2e2c7606
--- /dev/null
+++ b/flipapp_runtime_full/app.js
@@ -0,0 +1,5 @@
+fetch("ui/flipapp_ui.logline")
+  .then(res => res.text())
+  .then(text => {
+    document.getElementById("app").innerText = "FlipApp carregado com spans:\n" + text;
+  });
\ No newline at end of file
diff --git a/flipapp_runtime_full/cli/flip_commands.logline b/flipapp_runtime_full/cli/flip_commands.logline
new file mode 100644
index 0000000000000000000000000000000000000000..547596c4acddecd53bbd456c2e0bfce7270ae7f8
--- /dev/null
+++ b/flipapp_runtime_full/cli/flip_commands.logline
@@ -0,0 +1,187 @@
+[
+  {
+    "id": "flip_cli_definition",
+    "timestamp": "2025-06-04T02:22:10Z",
+    "type": "cli_definition",
+    "name": "flip",
+    "description": "LogLineOS Command Line Interface",
+    "version": "1.0.0-loglineos"
+  },
+
+  {
+    "id": "command_run_with_simulation",
+    "timestamp": "2025-06-04T02:22:15Z",
+    "type": "cli_command",
+    "name": "run",
+    "description": "Start the LogLineOS application",
+    "flags": [
+      {
+        "name": "watch",
+        "short": "w",
+        "description": "Watch for file changes and reload",
+        "type": "boolean",
+        "default": false
+      },
+      {
+        "name": "config",
+        "short": "c",
+        "description": "Path to config file",
+        "type": "string",
+        "default": "runtime/config.logline"
+      },
+      {
+        "name": "simulate",
+        "short": "s",
+        "description": "Run in simulation mode (dry-run)",
+        "type": "boolean",
+        "default": false
+      },
+      {
+        "name": "port",
+        "short": "p",
+        "description": "HTTP server port",
+        "type": "integer",
+        "default": 8080
+      }
+    ],
+    "execution_contract": "cli_run_application"
+  },
+
+  {
+    "id": "contract_cli_run_application",
+    "timestamp": "2025-06-04T02:22:20Z",
+    "type": "contract_definition",
+    "name": "cli_run_application",
+    "version": "1.0.0-loglineos",
+    "description": "Executes the run command with all options",
+    "created_by": "LogLineOS CLI",
+    "args_schema": {
+      "watch": { "type": "boolean", "default": false },
+      "config": { "type": "string", "default": "runtime/config.logline" },
+      "simulate": { "type": "boolean", "default": false },
+      "port": { "type": "integer", "default": 8080 }
+    },
+    "effects": [
+      {
+        "type": "when",
+        "when": "args.simulate === true",
+        "effects": [
+          {
+            "type": "contract_call",
+            "contract_name": "enable_simulation_mode",
+            "args": { "simulation_level": "full" }
+          }
+        ]
+      },
+      {
+        "type": "audit_log",
+        "level": "info",
+        "message": "🚀 Starting FlipApp LogLineOS",
+        "context": {
+          "mode": "{{args.simulate ? 'SIMULATION' : 'NORMAL'}}",
+          "watch": "{{args.watch}}",
+          "port": "{{args.port}}"
+        }
+      },
+      {
+        "type": "kernel_action",
+        "action_type": "start_http_server",
+        "port_template": "{{args.port}}",
+        "static_dir": "public",
+        "watch_mode": "{{args.watch}}"
+      },
+      {
+        "type": "when",
+        "when": "args.watch === true",
+        "effects": [
+          {
+            "type": "kernel_action",
+            "action_type": "start_file_watcher",
+            "watch_patterns": ["*.logline", "**/*.logline"],
+            "on_change_contract": "reload_application"
+          }
+        ]
+      }
+    ]
+  },
+
+  {
+    "id": "command_simulate",
+    "timestamp": "2025-06-04T02:22:25Z",
+    "type": "cli_command",
+    "name": "simulate",
+    "description": "Run contracts in simulation mode",
+    "arguments": [
+      {
+        "name": "contract_name",
+        "description": "Name of the contract to simulate",
+        "required": true
+      }
+    ],
+    "flags": [
+      {
+        "name": "args",
+        "short": "a",
+        "description": "JSON arguments for the contract",
+        "type": "string",
+        "default": "{}"
+      },
+      {
+        "name": "level",
+        "short": "l",
+        "description": "Simulation detail level",
+        "type": "string",
+        "enum": ["basic", "full", "verbose"],
+        "default": "full"
+      }
+    ],
+    "execution_contract": "cli_simulate_contract"
+  },
+
+  {
+    "id": "contract_cli_simulate_contract",
+    "timestamp": "2025-06-04T02:22:30Z",
+    "type": "contract_definition",
+    "name": "cli_simulate_contract",
+    "version": "1.0.0-loglineos",
+    "description": "Simulates a specific contract execution",
+    "created_by": "LogLineOS CLI",
+    "args_schema": {
+      "contract_name": { "type": "string", "required": true },
+      "args": { "type": "string", "default": "{}" },
+      "level": { "type": "string", "enum": ["basic", "full", "verbose"], "default": "full" }
+    },
+    "effects": [
+      {
+        "type": "contract_call",
+        "contract_name": "enable_simulation_mode",
+        "args": { "simulation_level": "{{args.level}}" }
+      },
+      {
+        "type": "audit_log",
+        "level": "info",
+        "message": "🔍 Simulating contract: {{args.contract_name}}",
+        "context": { "simulation_started": true }
+      },
+      {
+        "type": "kernel_action",
+        "action_type": "parse_json_string",
+        "json_string_template": "{{args.args}}",
+        "response_path": "temp.parsed_args"
+      },
+      {
+        "type": "contract_call",
+        "contract_name": "{{args.contract_name}}",
+        "args": "{{temp.parsed_args}}"
+      },
+      {
+        "type": "contract_call",
+        "contract_name": "simulation_summary"
+      },
+      {
+        "type": "contract_call",
+        "contract_name": "disable_simulation_mode"
+      }
+    ]
+  }
+]
\ No newline at end of file
diff --git a/flipapp_runtime_full/contracts/audit.logline b/flipapp_runtime_full/contracts/audit.logline
new file mode 100644
index 0000000000000000000000000000000000000000..c4bc725b71ac857be1283a2b07431f57d90ceef4
--- /dev/null
+++ b/flipapp_runtime_full/contracts/audit.logline
@@ -0,0 +1,29 @@
+# contracts/audit.logline
+# Definição de spans de auditoria e contratos de logging.
+
+- type: span_definition
+  name: "audit_log_info"
+  properties:
+    message: "string"
+    context: "object"
+
+- type: contract
+  name: log_debug
+  args:
+    message: "string"
+    context: "object"
+  effects:
+    - type: audit_log_debug
+      message: "{{args.message}}"
+      context: "{{args.context}}"
+
+- type: contract
+  name: log_error
+  args:
+    message: "string"
+    error_obj: "object"
+  effects:
+    - type: audit_log_error
+      message: "{{args.message}}"
+      error: "{{args.error_obj.message}}"
+      stack: "{{args.error_obj.stack}}"
diff --git a/flipapp_runtime_full/contracts/chat.logline b/flipapp_runtime_full/contracts/chat.logline
new file mode 100644
index 0000000000000000000000000000000000000000..4270b99320cf91b76ef601858f2f4b52817efd05
--- /dev/null
+++ b/flipapp_runtime_full/contracts/chat.logline
@@ -0,0 +1,229 @@
+[
+  {
+    "id": "contract_send_message",
+    "timestamp": "2025-06-04T02:12:18Z",
+    "type": "contract_definition",
+    "name": "send_message",
+    "version": "1.0.0-loglineos",
+    "description": "Sends a user message in the chat and triggers an LLM response.",
+    "created_by": "LogLineOS Kernel",
+    "args_schema": {
+      "message": { "type": "string", "required": true, "min_length": 1, "max_length": 1000 },
+      "conversation_id": { "type": "string", "required": true }
+    },
+    "middleware": {
+      "before": [
+        { "type": "audit_log", "level": "info", "message": "Attempting to send message to {{args.conversation_id}}" }
+      ],
+      "after": [
+        { "type": "audit_log", "level": "info", "message": "Message sent successfully." }
+      ],
+      "onError": [
+        { "type": "audit_log", "level": "error", "message": "Failed to send message: {{error.message}}" },
+        { "type": "ui_notification", "level": "error", "content_template": "Erro ao enviar mensagem: {{error.message}}" }
+      ]
+    },
+    "rate_limit": { "max_calls": 10, "window_ms": 60000 },
+    "effects": [
+      {
+        "id": "effect_clear_input",
+        "type": "state_mutation",
+        "path": "chat.message_input",
+        "value": ""
+      },
+      {
+        "id": "effect_add_user_message",
+        "type": "state_mutation",
+        "path": "chat.history.{{args.conversation_id}}",
+        "append_value": {
+          "id": "{{generate_uuid()}}",
+          "sender": "user",
+          "content": "{{args.message}}",
+          "timestamp": "{{get_timestamp()}}"
+        }
+      },
+      {
+        "id": "effect_scroll_to_bottom_after_user_msg",
+        "type": "ui_scroll_to",
+        "selector": ".chat-history-container",
+        "options": { "block": "end" }
+      },
+      {
+        "id": "effect_set_typing_indicator",
+        "type": "state_mutation",
+        "path": "chat.is_typing",
+        "value": "true"
+      },
+      {
+        "id": "effect_llm_query",
+        "type": "llm_query",
+        "prompt_template": "{{args.message}}",
+        "model_template": "{{state.config.default_llm_model}}",
+        "context_template": "{{state.chat.history[args.conversation_id]}}",
+        "on_response_contract_name_template": "process_llm_response",
+        "on_error_contract_name_template": "handle_llm_query_error"
+      }
+    ]
+  },
+
+  {
+    "id": "contract_process_llm_response",
+    "timestamp": "2025-06-04T02:12:23Z",
+    "type": "contract_definition",
+    "name": "process_llm_response",
+    "version": "1.0.0-loglineos",
+    "description": "Processes the LLM response and adds it to chat history.",
+    "created_by": "LogLineOS Kernel",
+    "args_schema": {
+      "response_content": { "type": "string", "required": true },
+      "conversation_id": { "type": "string", "required": true }
+    },
+    "effects": [
+      {
+        "id": "effect_unset_typing_indicator",
+        "type": "state_mutation",
+        "path": "chat.is_typing",
+        "value": "false"
+      },
+      {
+        "id": "effect_add_bot_message",
+        "type": "state_mutation",
+        "path": "chat.history.{{args.conversation_id}}",
+        "append_value": {
+          "id": "{{generate_uuid()}}",
+          "sender": "bot",
+          "content": "{{args.response_content}}",
+          "timestamp": "{{get_timestamp()}}"
+        }
+      },
+      {
+        "id": "effect_scroll_to_bottom_after_bot_msg",
+        "type": "ui_scroll_to",
+        "selector": ".chat-history-container",
+        "options": { "block": "end" }
+      },
+      {
+        "id": "effect_haptic_feedback_message",
+        "type": "ui_haptic_feedback",
+        "pattern": "short_impact"
+      },
+      {
+        "id": "effect_play_message_received_sound",
+        "type": "ui_play_sound",
+        "sound_path": "/sounds/message_received.mp3"
+      }
+    ]
+  },
+
+  {
+    "id": "contract_handle_llm_query_error",
+    "timestamp": "2025-06-04T02:12:28Z",
+    "type": "contract_definition",
+    "name": "handle_llm_query_error",
+    "version": "1.0.0-loglineos",
+    "description": "Handles errors from LLM queries.",
+    "created_by": "LogLineOS Kernel",
+    "args_schema": {
+      "error_message": { "type": "string", "required": true }
+    },
+    "effects": [
+      {
+        "id": "effect_unset_typing_indicator_on_error",
+        "type": "state_mutation",
+        "path": "chat.is_typing",
+        "value": "false"
+      },
+      {
+        "id": "effect_add_error_message_to_chat",
+        "type": "state_mutation",
+        "path": "chat.history.{{args.conversation_id}}",
+        "append_value": {
+          "id": "{{generate_uuid()}}",
+          "sender": "system",
+          "content": "Ops! Erro do LLM: {{args.error_message}}",
+          "timestamp": "{{get_timestamp()}}",
+          "is_error": true
+        }
+      },
+      {
+        "id": "effect_haptic_feedback_error",
+        "type": "ui_haptic_feedback",
+        "pattern": "long_impact"
+      },
+      {
+        "id": "effect_play_error_sound",
+        "type": "ui_play_sound",
+        "sound_path": "/sounds/error.mp3"
+      }
+    ]
+  },
+
+  {
+    "id": "contract_user_typing_debounce",
+    "timestamp": "2025-06-04T02:12:33Z",
+    "type": "contract_definition",
+    "name": "user_typing_debounce",
+    "version": "1.0.0-loglineos",
+    "description": "Manages the user typing indicator with debounce logic.",
+    "created_by": "LogLineOS Kernel",
+    "args_schema": {},
+    "effects": [
+      {
+        "id": "effect_set_typing_true",
+        "type": "state_mutation",
+        "path": "chat.is_typing",
+        "value": "true"
+      },
+      {
+        "id": "effect_debounce_typing_timeout",
+        "type": "kernel_action",
+        "action_type": "invoke_native_debounce_timeout",
+        "timeout_ms": 3000,
+        "on_timeout_contract_name": "reset_typing_indicator",
+        "debounce_id": "chat_typing"
+      }
+    ]
+  },
+
+  {
+    "id": "contract_reset_typing_indicator",
+    "timestamp": "2025-06-04T02:12:38Z",
+    "type": "contract_definition",
+    "name": "reset_typing_indicator",
+    "version": "1.0.0-loglineos",
+    "description": "Resets the chat typing indicator after debounce.",
+    "created_by": "LogLineOS Kernel",
+    "args_schema": {},
+    "effects": [
+      {
+        "id": "effect_set_typing_false",
+        "type": "state_mutation",
+        "path": "chat.is_typing",
+        "value": "false"
+      }
+    ]
+  },
+
+  {
+    "id": "contract_input_key_press",
+    "timestamp": "2025-06-04T02:12:43Z",
+    "type": "contract_definition",
+    "name": "input_key_press",
+    "version": "1.0.0-loglineos",
+    "description": "Handles keypress events for input fields, especially 'Enter' to send.",
+    "created_by": "LogLineOS Kernel",
+    "args_schema": {
+      "key": { "type": "string", "required": true },
+      "target_value": { "type": "string", "required": true },
+      "contract_on_enter": { "type": "string", "required": false }
+    },
+    "effects": [
+      {
+        "type": "contract_call",
+        "contract_name": "{{args.contract_on_enter}}",
+        "args": { "message": "{{args.target_value}}" },
+        "when": "args.key === 'Enter' && args.contract_on_enter"
+      }
+    ]
+  }
+]
\ No newline at end of file
diff --git a/flipapp_runtime_full/contracts/infra/fix_structure.logline b/flipapp_runtime_full/contracts/infra/fix_structure.logline
new file mode 100644
index 0000000000000000000000000000000000000000..88937e49abf1f6bb736da465015997a046b3da3d
--- /dev/null
+++ b/flipapp_runtime_full/contracts/infra/fix_structure.logline
@@ -0,0 +1,40 @@
+- type: commit
+  name: "fix_structure"
+  description: "Reorganiza o repositório com base na curadoria técnica e simbólica"
+  steps:
+
+    # 🔥 Remover estruturas obsoletas
+    - run: rm -rf flipapp_with_whatsapp_merged
+    - run: rm -rf rust_vm
+    - run: rm -f remove_infrastructure.sh
+    - run: rm -f llm_proxy_hard.py
+    - run: rm -f validate_ui_schema.logline
+    - run: rm -f install_flip.sh
+
+    # 🧹 Consolidar e mover arquivos relevantes
+    - run: mv cleanup/* scripts/
+    - run: rm -rf cleanup
+
+    # 💎 Criar pastas para organização futura
+    - run: mkdir -p legacy archive system_docs
+
+    # 📦 Arquivar o que precisa de revisão
+    - run: mv flip-cli archive/
+    - run: mv cleanup_deprecated.sh archive/
+    - run: mv todo_bunker.txt archive/
+    - run: mv deployment archive/
+    - run: mv dev archive/
+    - run: mv examples archive/
+
+    # 🧠 Mover itens ambíguos para legacy
+    - run: mv app_boot.logline legacy/
+    - run: mv flip_rebuild_pack.sh legacy/
+
+    # 💖 Normalizar o estado final
+    - run: mv flipapp_ui_final_canonic.logline ui/flipapp_ui.logline
+
+    # 🔐 Garantir que todos arquivos tenham permissão correta
+    - run: chmod -R 755 scripts ui runtime engine cli
+
+    # ✅ Mensagem de sucesso
+    - echo: "Estrutura reorganizada com sucesso. Todos os arquivos essenciais preservados."
\ No newline at end of file
diff --git a/flipapp_runtime_full/contracts/llm_response_contract.logline b/flipapp_runtime_full/contracts/llm_response_contract.logline
new file mode 100644
index 0000000000000000000000000000000000000000..947159266c8b1aad82189e21e26e3b362d791218
--- /dev/null
+++ b/flipapp_runtime_full/contracts/llm_response_contract.logline
@@ -0,0 +1,13 @@
+# llm_response_contract.logline
+# Contrato que processa llm_response e aplica resultado ao estado
+
+- type: contract
+  name: apply_completion_result
+  args:
+    text: "string"
+  effects:
+    - type: state_mutation
+      path: "state.completion_output"
+      value: "{{args.text}}"
+    - type: audit_log_info
+      message: "LLM completou com entrada: {{args.text | slice(0,20)}}"
diff --git a/flipapp_runtime_full/engine/animation_engine.logline b/flipapp_runtime_full/engine/animation_engine.logline
new file mode 100644
index 0000000000000000000000000000000000000000..584dcfc96dd17034d65032436154d128c7f5b4e9
--- /dev/null
+++ b/flipapp_runtime_full/engine/animation_engine.logline
@@ -0,0 +1,88 @@
+[
+  {
+    "id": "animation_engine_meta_def",
+    "timestamp": "2024-06-08T15:00:00Z",
+    "type": "engine_definition",
+    "name": "Animation & Feedback Engine",
+    "description": "Defines rules for animating UI elements and providing sensory feedback.",
+    "version": "1.0.0-loglineos"
+  },
+
+  {
+    "id": "rule_ui_animate_element",
+    "timestamp": "2024-06-08T15:00:05Z",
+    "type": "execution_rule",
+    "description": "Executes a UI animation on a DOM element. Respects prefers-reduced-motion.",
+    "match": {
+      "type": "ui_animate"
+    },
+    "kernel_action": {
+      "action_type": "invoke_native_dom_animation",
+      "selector_template": "{{span.selector}}",
+      "keyframes_template": "{{span.keyframes}}",
+      "options_template": "{{span.options}}",
+      "prefers_reduced_motion_template": "{{state.ui.prefers_reduced_motion}}"
+    },
+    "output_mapping": {
+      "animation_id": "status.animation_id",
+      "execution_status": "status.execution_status"
+    },
+    "audit_event_type": "ui_animation_execution"
+  },
+
+  {
+    "id": "rule_ui_scroll_to_view",
+    "timestamp": "2024-06-08T15:00:10Z",
+    "type": "execution_rule",
+    "description": "Scrolls a DOM element into view.",
+    "match": {
+      "type": "ui_scroll_to"
+    },
+    "kernel_action": {
+      "action_type": "invoke_native_dom_scroll",
+      "selector_template": "{{span.selector}}",
+      "options_template": "{{span.options}}",
+      "prefers_reduced_motion_template": "{{state.ui.prefers_reduced_motion}}"
+    },
+    "output_mapping": {
+      "status": "scroll.status"
+    },
+    "audit_event_type": "ui_scroll_execution"
+  },
+
+  {
+    "id": "rule_ui_play_sound",
+    "timestamp": "2024-06-08T15:00:15Z",
+    "type": "execution_rule",
+    "description": "Plays an audio file via the native audio system.",
+    "match": {
+      "type": "ui_play_sound"
+    },
+    "kernel_action": {
+      "action_type": "invoke_native_audio_playback",
+      "sound_path_template": "{{span.sound_path}}"
+    },
+    "output_mapping": {
+      "status": "audio.status"
+    },
+    "audit_event_type": "audio_playback_event"
+  },
+
+  {
+    "id": "rule_ui_haptic_feedback",
+    "timestamp": "2024-06-08T15:00:20Z",
+    "type": "execution_rule",
+    "description": "Triggers haptic feedback (vibration) on supported devices.",
+    "match": {
+      "type": "ui_haptic_feedback"
+    },
+    "kernel_action": {
+      "action_type": "invoke_native_haptic_feedback",
+      "pattern_template": "{{span.pattern}}"
+    },
+    "output_mapping": {
+      "status": "haptic.status"
+    },
+    "audit_event_type": "haptic_feedback_event"
+  }
+]
\ No newline at end of file
diff --git a/flipapp_runtime_full/engine/contract_execution_rules.logline b/flipapp_runtime_full/engine/contract_execution_rules.logline
new file mode 100644
index 0000000000000000000000000000000000000000..a4b0ba565d04314a1790050c81c692a258a6b04b
--- /dev/null
+++ b/flipapp_runtime_full/engine/contract_execution_rules.logline
@@ -0,0 +1,122 @@
+[
+  {
+    "id": "contract_engine_def",
+    "timestamp": "2025-06-04T02:11:43Z",
+    "type": "engine_definition",
+    "name": "Contract Execution Engine",
+    "description": "Defines rules for executing and auditing LogLine contracts.",
+    "version": "1.0.0-loglineos"
+  },
+
+  {
+    "id": "rule_execute_contract_call",
+    "timestamp": "2025-06-04T02:11:48Z",
+    "type": "execution_rule",
+    "description": "Orchestrates the full lifecycle of a contract execution.",
+    "match": {
+      "type": "contract_call"
+    },
+    "kernel_action": {
+      "action_type": "orchestrate_contract_execution",
+      "contract_name_template": "{{span.contract_name}}",
+      "args_template": "{{span.args}}",
+      "context_template": "{{span.context}}",
+      "originating_span_id_template": "{{span.id}}"
+    },
+    "output_mapping": {
+      "status": "contract_execution.status",
+      "result": "contract_execution.result",
+      "error": "contract_execution.error"
+    },
+    "audit_event_type": "contract_flow_execution"
+  },
+
+  {
+    "id": "rule_contract_arg_validation",
+    "timestamp": "2025-06-04T02:11:53Z",
+    "type": "execution_rule",
+    "description": "Performs argument validation for a contract.",
+    "match": {
+      "type": "internal_validate_contract_args"
+    },
+    "kernel_action": {
+      "action_type": "invoke_native_schema_validator",
+      "schema_template": "{{span.validation_schema}}",
+      "data_to_validate_template": "{{span.data_to_validate}}"
+    },
+    "output_mapping": {
+      "is_valid": "validation.status",
+      "errors": "validation.messages"
+    },
+    "audit_event_type": "contract_argument_validation"
+  },
+
+  {
+    "id": "rule_contract_middleware_execution",
+    "timestamp": "2025-06-04T02:11:58Z",
+    "type": "execution_rule",
+    "description": "Executes a list of contract middleware effects.",
+    "match": {
+      "type": "internal_execute_contract_middleware"
+    },
+    "kernel_action": {
+      "action_type": "execute_logline_effects_list",
+      "effects_list_template": "{{span.effects_list}}",
+      "context_template": "{{span.execution_context}}"
+    },
+    "audit_event_type": "contract_middleware_execution"
+  },
+
+  {
+    "id": "rule_contract_effects_execution",
+    "timestamp": "2025-06-04T02:12:03Z",
+    "type": "execution_rule",
+    "description": "Executes the main effects list of a contract.",
+    "match": {
+      "type": "internal_execute_contract_effects"
+    },
+    "kernel_action": {
+      "action_type": "execute_logline_effects_list",
+      "effects_list_template": "{{span.effects_list}}",
+      "context_template": "{{span.execution_context}}"
+    },
+    "audit_event_type": "contract_effects_execution"
+  },
+
+  {
+    "id": "rule_contract_rate_limit_check",
+    "timestamp": "2025-06-04T02:12:08Z",
+    "type": "execution_rule",
+    "description": "Checks and applies rate limiting for a contract.",
+    "match": {
+      "type": "internal_contract_rate_limit_check"
+    },
+    "kernel_action": {
+      "action_type": "invoke_native_rate_limiter",
+      "contract_name_template": "{{span.contract_name}}",
+      "rate_limit_config_template": "{{span.rate_limit_config}}"
+    },
+    "output_mapping": {
+      "allowed": "rate_limit.allowed",
+      "remaining": "rate_limit.remaining",
+      "reset_time": "rate_limit.reset_time"
+    },
+    "audit_event_type": "contract_rate_limit_check"
+  },
+
+  {
+    "id": "rule_contract_log_history",
+    "timestamp": "2025-06-04T02:12:13Z",
+    "type": "execution_rule",
+    "description": "Logs a contract execution to the internal history for analytics and rate limiting.",
+    "match": {
+      "type": "internal_contract_log_history"
+    },
+    "kernel_action": {
+      "action_type": "invoke_native_contract_history_recorder",
+      "execution_details_template": "{{span.details}}"
+    },
+    "output_mapping": {},
+    "audit_event_type": "contract_history_record"
+  }
+]
\ No newline at end of file
diff --git a/flipapp_runtime_full/engine/gesture_engine.logline b/flipapp_runtime_full/engine/gesture_engine.logline
new file mode 100644
index 0000000000000000000000000000000000000000..781836a0bd0891040795cb9d481d6fabcfea11e0
--- /dev/null
+++ b/flipapp_runtime_full/engine/gesture_engine.logline
@@ -0,0 +1,72 @@
+[
+  {
+    "id": "gesture_engine_meta_def",
+    "timestamp": "2024-06-08T16:00:00Z",
+    "type": "engine_definition",
+    "name": "Gesture Engine",
+    "description": "Defines rules for detecting and reacting to user gestures.",
+    "version": "1.0.0-loglineos"
+  },
+
+  {
+    "id": "rule_register_swipe_gesture",
+    "timestamp": "2024-06-08T16:00:05Z",
+    "type": "execution_rule",
+    "description": "Registers a swipe gesture listener on a DOM element for touch and mouse.",
+    "match": {
+      "type": "gesture_register_swipe"
+    },
+    "kernel_action": {
+      "action_type": "invoke_native_gesture_register_swipe",
+      "element_selector_template": "{{span.selector}}",
+      "direction_template": "{{span.direction}}",
+      "threshold_template": "{{span.threshold}}",
+      "callback_contract_name_template": "{{span.on_swipe_contract}}"
+    },
+    "output_mapping": {
+      "listener_id": "gesture_registration.listener_id",
+      "status": "gesture_registration.status"
+    },
+    "audit_event_type": "gesture_swipe_registration"
+  },
+
+  {
+    "id": "rule_register_tap_gesture",
+    "timestamp": "2024-06-08T16:00:10Z",
+    "type": "execution_rule",
+    "description": "Registers a tap gesture listener on a DOM element.",
+    "match": {
+      "type": "gesture_register_tap"
+    },
+    "kernel_action": {
+      "action_type": "invoke_native_gesture_register_tap",
+      "element_selector_template": "{{span.selector}}",
+      "max_duration_ms_template": "{{span.max_duration_ms | default(300)}}",
+      "max_movement_px_template": "{{span.max_movement_px | default(10)}}",
+      "callback_contract_name_template": "{{span.on_tap_contract}}"
+    },
+    "output_mapping": {
+      "listener_id": "gesture_registration.listener_id",
+      "status": "gesture_registration.status"
+    },
+    "audit_event_type": "gesture_tap_registration"
+  },
+
+  {
+    "id": "rule_unregister_gesture_listener",
+    "timestamp": "2024-06-08T16:00:15Z",
+    "type": "execution_rule",
+    "description": "Unregisters a gesture listener by its ID.",
+    "match": {
+      "type": "gesture_unregister"
+    },
+    "kernel_action": {
+      "action_type": "invoke_native_gesture_unregister",
+      "listener_id_template": "{{span.listener_id}}"
+    },
+    "output_mapping": {
+      "status": "gesture_status.status"
+    },
+    "audit_event_type": "gesture_unregistration"
+  }
+]
\ No newline at end of file
diff --git a/flipapp_runtime_full/engine/simulation_rules.logline b/flipapp_runtime_full/engine/simulation_rules.logline
new file mode 100644
index 0000000000000000000000000000000000000000..df20d671145535d39108f42d46455144e38eaf43
--- /dev/null
+++ b/flipapp_runtime_full/engine/simulation_rules.logline
@@ -0,0 +1,239 @@
+[
+  {
+    "id": "simulation_engine_def",
+    "timestamp": "2025-06-04T02:21:30Z",
+    "type": "engine_definition",
+    "name": "Simulation Engine",
+    "description": "Defines rules for running LogLineOS in dry-run mode.",
+    "version": "1.0.0-loglineos"
+  },
+
+  {
+    "id": "contract_noop_effect",
+    "timestamp": "2025-06-04T02:21:35Z",
+    "type": "contract_definition",
+    "name": "noop_effect",
+    "version": "1.0.0-loglineos",
+    "description": "A no-operation effect. Used to nullify real actions in simulation mode.",
+    "created_by": "LogLineOS Kernel",
+    "args_schema": {
+      "original_action_type": { "type": "string", "required": true },
+      "original_span_data": { "type": "object", "required": false }
+    },
+    "effects": [
+      {
+        "type": "audit_log",
+        "level": "debug",
+        "message": "(SIMULATED) Original action {{args.original_action_type}} was skipped.",
+        "context": { 
+          "simulation_mode": true,
+          "original_span": "{{args.original_span_data}}"
+        }
+      }
+    ]
+  },
+
+  {
+    "id": "rule_intercept_simulated_state_mutation",
+    "timestamp": "2025-06-04T02:21:40Z",
+    "type": "execution_rule",
+    "description": "Intercepts state mutations when simulation mode is active.",
+    "priority": 1000,
+    "match": {
+      "type": "state_mutation",
+      "when_expression": "state.execution.simulate_mode === true"
+    },
+    "kernel_action": {
+      "action_type": "execute_logline_effects_list",
+      "effects_list_template": [
+        {
+          "type": "contract_call",
+          "contract_name": "noop_effect",
+          "args": { 
+            "original_action_type": "state_mutation",
+            "original_span_data": {
+              "path": "{{span.path}}",
+              "value": "{{span.value}}",
+              "operation": "{{span.operation | default('set')}}"
+            }
+          }
+        },
+        {
+          "type": "audit_log",
+          "level": "info",
+          "message": "(SIMULATED) State mutation: {{span.path}} = {{span.value}}",
+          "context": { "simulation_preview": true }
+        }
+      ],
+      "context_template": {}
+    },
+    "stop_execution": true,
+    "audit_event_type": "simulated_state_mutation"
+  },
+
+  {
+    "id": "rule_intercept_simulated_http_request",
+    "timestamp": "2025-06-04T02:21:45Z",
+    "type": "execution_rule",
+    "description": "Intercepts HTTP requests when simulation mode is active.",
+    "priority": 1000,
+    "match": {
+      "type": "http_request",
+      "when_expression": "state.execution.simulate_mode === true"
+    },
+    "kernel_action": {
+      "action_type": "execute_logline_effects_list",
+      "effects_list_template": [
+        {
+          "type": "contract_call",
+          "contract_name": "noop_effect",
+          "args": { 
+            "original_action_type": "http_request",
+            "original_span_data": {
+              "method": "{{span.method}}",
+              "url": "{{span.url_template}}",
+              "headers": "{{span.headers}}"
+            }
+          }
+        },
+        {
+          "type": "audit_log",
+          "level": "info",
+          "message": "(SIMULATED) HTTP {{span.method}} request to: {{span.url_template}}",
+          "context": { "simulation_preview": true }
+        }
+      ],
+      "context_template": {}
+    },
+    "stop_execution": true,
+    "audit_event_type": "simulated_http_request"
+  },
+
+  {
+    "id": "rule_intercept_simulated_ui_actions",
+    "timestamp": "2025-06-04T02:21:50Z",
+    "type": "execution_rule",
+    "description": "Intercepts UI actions when simulation mode is active.",
+    "priority": 1000,
+    "match": {
+      "type_pattern": "ui_*",
+      "when_expression": "state.execution.simulate_mode === true"
+    },
+    "kernel_action": {
+      "action_type": "execute_logline_effects_list",
+      "effects_list_template": [
+        {
+          "type": "contract_call",
+          "contract_name": "noop_effect",
+          "args": { 
+            "original_action_type": "{{span.type}}",
+            "original_span_data": "{{span}}"
+          }
+        },
+        {
+          "type": "audit_log",
+          "level": "info",
+          "message": "(SIMULATED) UI action: {{span.type}}",
+          "context": { "simulation_preview": true }
+        }
+      ],
+      "context_template": {}
+    },
+    "stop_execution": true,
+    "audit_event_type": "simulated_ui_action"
+  },
+
+  {
+    "id": "contract_enable_simulation_mode",
+    "timestamp": "2025-06-04T02:21:55Z",
+    "type": "contract_definition",
+    "name": "enable_simulation_mode",
+    "version": "1.0.0-loglineos",
+    "description": "Enables global simulation mode for the LogLineOS instance.",
+    "created_by": "LogLineOS Kernel",
+    "args_schema": {
+      "simulation_level": { "type": "string", "enum": ["basic", "full", "verbose"], "default": "full" }
+    },
+    "effects": [
+      {
+        "type": "state_mutation",
+        "path": "execution.simulate_mode",
+        "value": true
+      },
+      {
+        "type": "state_mutation",
+        "path": "execution.simulation_level",
+        "value": "{{args.simulation_level}}"
+      },
+      {
+        "type": "state_mutation",
+        "path": "execution.simulation_start_time",
+        "value": "{{get_timestamp()}}"
+      },
+      {
+        "type": "audit_log",
+        "level": "warn",
+        "message": "🔍 SIMULATION MODE ENABLED - Level: {{args.simulation_level}}",
+        "context": { "simulation_activated": true }
+      }
+    ]
+  },
+
+  {
+    "id": "contract_disable_simulation_mode",
+    "timestamp": "2025-06-04T02:22:00Z",
+    "type": "contract_definition",
+    "name": "disable_simulation_mode",
+    "version": "1.0.0-loglineos",
+    "description": "Disables simulation mode and returns to normal execution.",
+    "created_by": "LogLineOS Kernel",
+    "args_schema": {},
+    "effects": [
+      {
+        "type": "state_mutation",
+        "path": "execution.simulate_mode",
+        "value": false
+      },
+      {
+        "type": "state_mutation",
+        "path": "execution.simulation_end_time",
+        "value": "{{get_timestamp()}}"
+      },
+      {
+        "type": "audit_log",
+        "level": "info",
+        "message": "✅ Simulation mode disabled - returning to normal execution",
+        "context": { 
+          "simulation_deactivated": true,
+          "duration": "{{state.execution.simulation_end_time - state.execution.simulation_start_time}}"
+        }
+      }
+    ]
+  },
+
+  {
+    "id": "contract_simulation_summary",
+    "timestamp": "2025-06-04T02:22:05Z",
+    "type": "contract_definition",
+    "name": "simulation_summary",
+    "version": "1.0.0-loglineos",
+    "description": "Generates a summary report of actions that would have been executed in simulation mode.",
+    "created_by": "LogLineOS Kernel",
+    "args_schema": {},
+    "effects": [
+      {
+        "type": "kernel_action",
+        "action_type": "generate_simulation_report",
+        "filter_audit_events": ["simulated_state_mutation", "simulated_http_request", "simulated_ui_action"],
+        "since_timestamp": "{{state.execution.simulation_start_time}}",
+        "response_path": "temp.simulation_report"
+      },
+      {
+        "type": "audit_log",
+        "level": "info",
+        "message": "📊 Simulation Summary Generated",
+        "context": "{{temp.simulation_report}}"
+      }
+    ]
+  }
+]
\ No newline at end of file
diff --git a/flipapp_runtime_full/engine/span_validator.logline b/flipapp_runtime_full/engine/span_validator.logline
new file mode 100644
index 0000000000000000000000000000000000000000..c1939d792e6a0751f1cd612e275ececd521a2700
--- /dev/null
+++ b/flipapp_runtime_full/engine/span_validator.logline
@@ -0,0 +1,119 @@
+[
+  {
+    "id": "span_validator_engine_def",
+    "timestamp": "2025-06-04T02:14:39Z",
+    "type": "engine_definition",
+    "name": "Span Validator Engine",
+    "description": "Defines rules for validating incoming LogLine spans against schemas.",
+    "version": "1.0.0-loglineos"
+  },
+
+  {
+    "id": "schema_base_span",
+    "timestamp": "2025-06-04T02:14:44Z",
+    "type": "schema_definition",
+    "name": "base_span_schema",
+    "properties": {
+      "id": { "type": "string", "required": true, "pattern": "^[a-zA-Z0-9_-]+$" },
+      "timestamp": { "type": "string", "required": true, "pattern": "^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}Z$" },
+      "type": { "type": "string", "required": true, "pattern": "^[a-zA-Z_]+$" }
+    }
+  },
+
+  {
+    "id": "schema_ui_animate",
+    "timestamp": "2025-06-04T02:14:49Z",
+    "type": "schema_definition",
+    "name": "ui_animate_schema",
+    "extends": "base_span_schema",
+    "properties": {
+      "selector": { "type": "string", "required": true },
+      "keyframes": { "type": "array", "required": true, "min_length": 2 },
+      "options": { "type": "object", "required": false }
+    }
+  },
+
+  {
+    "id": "rule_validate_incoming_span",
+    "timestamp": "2025-06-04T02:14:54Z",
+    "type": "execution_rule",
+    "description": "Validates an incoming span against its defined schema before execution.",
+    "match": {
+      "type": "internal_span_validation_request"
+    },
+    "kernel_action": {
+      "action_type": "invoke_native_schema_validator",
+      "schema_name_template": "get_schema_for_span_type({{span.span_to_validate.type}})",
+      "data_to_validate_template": "{{span.span_to_validate}}"
+    },
+    "output_mapping": {
+      "is_valid": "validation_result.status",
+      "errors": "validation_result.messages"
+    },
+    "audit_event_type": "span_validation_execution"
+  },
+
+  {
+    "id": "contract_get_schema_for_span_type",
+    "timestamp": "2025-06-04T02:14:59Z",
+    "type": "contract_definition",
+    "name": "get_schema_for_span_type",
+    "version": "1.0.0-loglineos",
+    "description": "Returns the schema definition for a given span type.",
+    "created_by": "LogLineOS Kernel",
+    "args_schema": {
+      "span_type": { "type": "string", "required": true }
+    },
+    "effects": [
+      {
+        "type": "switch",
+        "expression": "args.span_type",
+        "cases": [
+          {
+            "value": "ui_animate",
+            "effects": [
+              { "type": "return_value", "value": "{{lookup_schema('ui_animate_schema')}}" }
+            ]
+          },
+          {
+            "value": "llm_query",
+            "effects": [
+              { "type": "return_value", "value": "{{lookup_schema('llm_query_schema')}}" }
+            ]
+          }
+        ],
+        "default_effects": [
+          { "type": "return_value", "value": "{{lookup_schema('base_span_schema')}}" }
+        ]
+      }
+    ]
+  },
+
+  {
+    "id": "contract_handle_invalid_span_fallback",
+    "timestamp": "2025-06-04T02:15:04Z",
+    "type": "contract_definition",
+    "name": "handle_invalid_span_fallback",
+    "version": "1.0.0-loglineos",
+    "description": "Fallback contract executed when a span fails validation.",
+    "created_by": "LogLineOS Kernel",
+    "args_schema": {
+      "span_id": { "type": "string", "required": true },
+      "span_type": { "type": "string", "required": true },
+      "validation_errors": { "type": "array", "required": true }
+    },
+    "effects": [
+      {
+        "type": "audit_log",
+        "level": "error",
+        "message": "Invalid span received: {{args.span_type}} (ID: {{args.span_id}})",
+        "context": { "errors": "{{args.validation_errors}}" }
+      },
+      {
+        "type": "ui_notification",
+        "level": "error",
+        "content_template": "Erro de span: '{{args.span_type}}' inválido. Detalhes no log de auditoria."
+      }
+    ]
+  }
+]
\ No newline at end of file
diff --git a/flipapp_runtime_full/engine/ui_error_handler.logline b/flipapp_runtime_full/engine/ui_error_handler.logline
new file mode 100644
index 0000000000000000000000000000000000000000..edb78207e5aac3182cafe65aae3bd4bc94f59dc3
--- /dev/null
+++ b/flipapp_runtime_full/engine/ui_error_handler.logline
@@ -0,0 +1,91 @@
+[
+  {
+    "id": "ui_error_handler_engine_def",
+    "timestamp": "2025-06-04T02:15:59Z",
+    "type": "engine_definition",
+    "name": "UI Error Handler Engine",
+    "description": "Defines rules for handling UI errors and displaying user-friendly feedback.",
+    "version": "1.0.0-loglineos"
+  },
+
+  {
+    "id": "rule_ui_notification",
+    "timestamp": "2025-06-04T02:16:04Z",
+    "type": "execution_rule",
+    "description": "Displays a transient notification message on the UI.",
+    "match": {
+      "type": "ui_notification"
+    },
+    "kernel_action": {
+      "action_type": "invoke_native_ui_notification",
+      "level_template": "{{span.level}}",
+      "content_template": "{{span.content_template}}",
+      "duration_ms_template": "{{span.duration_ms | default(5000)}}"
+    },
+    "audit_event_type": "ui_notification_event"
+  },
+
+  {
+    "id": "rule_ui_error_overlay",
+    "timestamp": "2025-06-04T02:16:09Z",
+    "type": "execution_rule",
+    "description": "Displays a critical full-screen error overlay.",
+    "match": {
+      "type": "ui_error_overlay"
+    },
+    "kernel_action": {
+      "action_type": "invoke_native_ui_error_overlay",
+      "message_template": "{{span.message}}",
+      "details_template": "{{span.details}}"
+    },
+    "audit_event_type": "ui_error_overlay_display"
+  },
+
+  {
+    "id": "rule_kernel_error_handler",
+    "timestamp": "2025-06-04T02:16:14Z",
+    "type": "execution_rule",
+    "description": "Global error handler for the Kernel, displays critical errors.",
+    "match": {
+      "type": "kernel_error_event"
+    },
+    "kernel_action": {
+      "action_type": "execute_logline_effects_list",
+      "effects_list_template": [
+        {
+          "type": "ui_error_overlay",
+          "message": "Erro inesperado do sistema: {{span.error.message}}",
+          "details": "{{span.error.stack}}"
+        },
+        {
+          "type": "audit_log",
+          "level": "critical",
+          "message": "Global Kernel Error: {{span.error.message}}",
+          "context": { "stack": "{{span.error.stack}}" }
+        }
+      ],
+      "context_template": {}
+    },
+    "audit_event_type": "kernel_error_handled"
+  },
+
+  {
+    "id": "rule_ui_render_from_ast",
+    "timestamp": "2025-06-04T02:16:19Z",
+    "type": "execution_rule",
+    "description": "Renders UI from a parsed LogLine AST into the DOM.",
+    "match": {
+      "type": "ui_render_from_ast"
+    },
+    "kernel_action": {
+      "action_type": "invoke_native_ui_render_from_ast",
+      "ast_template": "{{span.ast_template}}",
+      "target_selector_template": "{{span.target_selector}}"
+    },
+    "output_mapping": {
+      "render_status": "ui_render.status",
+      "components_rendered": "ui_render.components_count"
+    },
+    "audit_event_type": "ui_render_from_ast_execution"
+  }
+]
\ No newline at end of file
diff --git a/flipapp_runtime_full/flip_hardening.logline b/flipapp_runtime_full/flip_hardening.logline
new file mode 100644
index 0000000000000000000000000000000000000000..25a825508d7e43d1e07a6183607f0d020b9bc148
--- /dev/null
+++ b/flipapp_runtime_full/flip_hardening.logline
@@ -0,0 +1,76 @@
+# flip_hardening.logline
+# Span com todo o plano de refatoração e hardening do FlipApp LogLineOS
+- type: refactor_task
+  id: "task_fallback_llm_local"
+  description: "Implementar fallback local no llm_proxy com modelo WASM ou resposta default"
+  status: "pending"
+
+- type: refactor_task
+  id: "task_cache_semantic"
+  description: "Adicionar cache semântico local por hash de prompt no llm_proxy"
+  status: "pending"
+
+- type: refactor_task
+  id: "task_llm_timeout_retry"
+  description: "Adicionar timeout, retry e captura de exceções no llm_proxy_hard"
+  status: "pending"
+
+- type: refactor_task
+  id: "task_logger_integration"
+  description: "Remover utils/logger.js ou integrar com audit.logline"
+  status: "pending"
+
+- type: refactor_task
+  id: "task_add_tests_logline"
+  description: "Adicionar testes .logline para expect_state, expect_effect e expect_error em contratos e regras"
+  status: "pending"
+
+- type: refactor_task
+  id: "task_contract_versioning"
+  description: "Adicionar version, description e created_by em todos os contratos, com validação de tipos"
+  status: "pending"
+
+- type: refactor_task
+  id: "task_ui_schema_validation"
+  description: "Criar schemas para UI e validar na carga de arquivos .logline, adicionar fallback visual"
+  status: "pending"
+
+- type: refactor_task
+  id: "task_rotate_logs_archive"
+  description: "Adicionar contratos rotate_logs, archive_spans e expire_state, configurar via cron local"
+  status: "pending"
+
+- type: refactor_task
+  id: "task_add_cli_flip"
+  description: "Criar binário 'flip' com comandos run, logs, audit e test"
+  status: "pending"
+
+- type: refactor_task
+  id: "task_consolidate_config"
+  description: "Consolidar configuração em config.logline e gerar load_env.sh, remover env-config.sh"
+  status: "pending"
+
+- type: refactor_task
+  id: "task_remove_docker"
+  description: "Remover Dockerfile, docker-compose.yml e pasta docker/, atualizar README.md"
+  status: "pending"
+
+- type: refactor_task
+  id: "task_span_validation"
+  description: "Adicionar span_validator.logline com validação sintática e contratos de fallback para spans inválidos"
+  status: "pending"
+
+- type: refactor_task
+  id: "task_llm_isolation"
+  description: "Adicionar origin, context_id ou user_id em llm_request.logline para rastreabilidade"
+  status: "pending"
+
+- type: refactor_task
+  id: "task_contract_audit"
+  description: "Adicionar audit_contract_change.logline para registro de alterações de contrato e changelog simbólico"
+  status: "pending"
+
+- type: refactor_task
+  id: "task_simulation_mode"
+  description: "Implementar modo 'simulate' em LogLine para spans sem efeitos colaterais"
+  status: "pending"
diff --git a/flipapp_runtime_full/flip_ui.logline b/flipapp_runtime_full/flip_ui.logline
new file mode 100644
index 0000000000000000000000000000000000000000..98cac5cf94cc9ab16cbcd7d2d9e36765d494f78d
--- /dev/null
+++ b/flipapp_runtime_full/flip_ui.logline
@@ -0,0 +1 @@
+{"type": "doc_comment", "source": "dan", "content": "🪞 O painel esquerdo, chamado Espelho, mostra sempre o histórico ChatGPT (por agora),\ncom três seções visíveis fixas: pessoal, trabalho e conta corrente. As sessões são do tipo `chat_card`.\n\n🪞 O conteúdo principal fica no painel da direita, com entradas e saídas do Chat (ou WhatsApp, ou contrato novo). Cada linha é um `span_card`.\n\n🎨 O layout é sempre em dois painéis: sidebar esquerda (fixa) e painel direito com conteúdo dinâmico, que muda entre três opções (chat, WhatsApp, contrato novo). Isso já está no `flipapp_ui.logline`.\n\n🎨 Tokens de estilo são definidos no `tokens_ui.logline`, como `padding_md`, `font_title`, `bg_dark`, `text_soft` etc.\n\n🎨 Os cards são sempre `radius_md`, `shadow_lg`, `bg_panel`, `gap_sm`, com padding fixo.\n\n🎨 Há transições suaves entre spans, com efeito de fade e anim_card_slide (entrada à direita).\n\n🎨 O layout é mobile-first, com breakpoints para o segundo painel (em telas menores ele vira um botão de alternância com ícone de menu).\n\n🎨 A navegação é feita com `router.push` entre as três telas. A seleção atual é marcada com `state=selected`.\n\n🎨 O modo imersivo (fullscreen) pode ser ativado com `view_mode=immersive`, que oculta o painel esquerdo.\n\n🎨 O som, vibração e toque podem ser aplicados com spans do tipo `fx_sound`, `fx_haptic`, `fx_shake` com parâmetros declarativos.\n\n🎨 O estilo geral é premium: dark elegante, com fundo escuro (#111), cards translúcidos (#222 com blur), textos brancos suaves (#eee), espaçamento generoso, animações responsivas e uma aura simbólica sofisticada. É como se fosse um app do futuro."}
\ No newline at end of file
diff --git a/flipapp_runtime_full/install_vm.logline b/flipapp_runtime_full/install_vm.logline
new file mode 100644
index 0000000000000000000000000000000000000000..c53896f2cf6e3fa7de2e7eaf6760790264b8fc3f
--- /dev/null
+++ b/flipapp_runtime_full/install_vm.logline
@@ -0,0 +1,46 @@
+# install_vm.logline
+# Roteiro auditável para instalação de um ambiente LogLineOS.
+
+- type: installation_step
+  id: "check_prerequisites"
+  description: "Verify if Node.js (for toolchain) and Rust/wasm-pack are installed."
+  commands:
+    - "node --version"
+    - "rustc --version"
+    - "wasm-pack --version"
+  expected_output: "version strings"
+  on_success_contract: "log_info(message='Prerequisites met')"
+  on_failure_contract: "log_error(message='Missing prerequisites. Please install Node.js and Rust/wasm-pack.'). Then fail_install_contract."
+
+- type: installation_step
+  id: "build_wasm_module"
+  description: "Compile the Rust WASM module."
+  commands:
+    - "cd rust_vm && wasm-pack build --target web --release"
+  on_success_contract: "log_info(message='WASM module built')"
+  on_failure_contract: "log_error(message='WASM build failed'). Then fail_install_contract."
+
+- type: installation_step
+  id: "install_frontend_dependencies"
+  description: "Install frontend JavaScript dependencies."
+  commands:
+    - "npm install"
+  on_success_contract: "log_info(message='Frontend dependencies installed')"
+  on_failure_contract: "log_error(message='Frontend dependencies install failed'). Then fail_install_contract."
+
+- type: installation_step
+  id: "start_local_http_server"
+  description: "Start the local HTTP server to serve the public/ directory."
+  commands:
+    - "logline run scripts/serve.py.logline --port 8080 --root public"
+  on_success_contract: "log_info(message='Local HTTP server running')"
+  on_failure_contract: "log_error(message='Failed to start local HTTP server'). Then fail_install_contract."
+
+- type: contract
+  name: fail_install_contract
+  effects:
+    - type: ui_notification
+      level: "critical"
+      content: "Instalação do LogLineOS falhou! Verifique os logs."
+    - type: audit_log_error
+      message: "Fatal installation failure."
diff --git a/flipapp_runtime_full/logline.yml b/flipapp_runtime_full/logline.yml
new file mode 100644
index 0000000000000000000000000000000000000000..21d6847e93a8c6518d98762a811627cc609e352a
--- /dev/null
+++ b/flipapp_runtime_full/logline.yml
@@ -0,0 +1,17 @@
+project:
+  name: "FlipApp LogLineOS"
+  version: "1.0.0-pure"
+  description: "A flagship application demonstrating 100% LogLineOS paradigm."
+  authors: ["Your Name"]
+  license: "MIT"
+  entry_point: "/ui/flipapp_ui.logline"
+  default_state: "/spans/state.jsonl"
+  default_audit_log: "/spans/audit.jsonl"
+  default_chat_history: "/spans/chat_history.jsonl"
+
+execution:
+  mode: "browser"
+  initial_contracts:
+    - "log_info(message='LogLineOS started')"
+    - "initialize_session"
+    - "load_ui"
diff --git a/flipapp_runtime_full/observer_bot.logline b/flipapp_runtime_full/observer_bot.logline
new file mode 100644
index 0000000000000000000000000000000000000000..ab5c71942fc99e97239fb6d7bd2dc2f352366a94
--- /dev/null
+++ b/flipapp_runtime_full/observer_bot.logline
@@ -0,0 +1,169 @@
+[
+  {
+    "id": "observer_bot_engine_def",
+    "timestamp": "2025-06-04T10:05:46Z",
+    "type": "engine_definition",
+    "name": "Observer Bot Engine",
+    "description": "Monitora mensagens WhatsApp e executa pipeline de processamento automático",
+    "version": "1.0.0-loglineos"
+  },
+
+  {
+    "id": "continuous_whatsapp_observer",
+    "timestamp": "2025-06-04T10:05:46Z",
+    "type": "continuous_rule",
+    "description": "Observa mensagens WebSocket do WhatsApp",
+    "trigger": {
+      "type": "websocket_message",
+      "url": "wss://api.flipapp/whatsapp/{{state.user_id}}"
+    },
+    "effects": [
+      {
+        "type": "audit_log",
+        "level": "info",
+        "message": "Nova mensagem WhatsApp recebida",
+        "context": {
+          "conversation_id": "{{message.conversation_id}}",
+          "sender": "{{message.sender}}",
+          "content_preview": "{{message.text | truncate(50)}}"
+        }
+      },
+      {
+        "type": "when",
+        "when": "!state.messages[message.conversation_id]",
+        "effects": [
+          {
+            "type": "contract_call",
+            "contract_name": "load_chat_history",
+            "args": {
+              "conversation_id": "{{message.conversation_id}}"
+            }
+          }
+        ]
+      },
+      {
+        "type": "contract_call",
+        "contract_name": "detect_language",
+        "args": {
+          "text": "{{message.text}}"
+        },
+        "response_path": "temp.detected_language"
+      },
+      {
+        "type": "when",
+        "when": "temp.detected_language.language != 'pt'",
+        "effects": [
+          {
+            "type": "contract_call",
+            "contract_name": "translate_message",
+            "args": {
+              "text": "{{message.text}}",
+              "source": "{{temp.detected_language.language}}",
+              "target": "pt"
+            },
+            "response_path": "temp.translated_message"
+          }
+        ]
+      },
+      {
+        "type": "when",
+        "when": "message.media_type == 'audio'",
+        "effects": [
+          {
+            "type": "contract_call",
+            "contract_name": "transcribe_audio",
+            "args": {
+              "audio": "{{message.content}}"
+            },
+            "response_path": "temp.transcription"
+          }
+        ]
+      },
+      {
+        "type": "kernel_action",
+        "action_type": "execute_function",
+        "function_name": "compute_risk",
+        "args": {
+          "text": "{{message.text}}"
+        },
+        "response_path": "temp.risk_level"
+      },
+      {
+        "type": "contract_call",
+        "contract_name": "raise_alert",
+        "args": {
+          "message": "{{message.text}}",
+          "level": "{{temp.risk_level}}"
+        }
+      },
+      {
+        "type": "state_mutation",
+        "path": "message_counter",
+        "operation": "increment",
+        "value": 1
+      },
+      {
+        "type": "when",
+        "when": "state.message_counter > state.config.daily_msg_limit",
+        "effects": [
+          {
+            "type": "contract_call",
+            "contract_name": "raise_alert",
+            "args": {
+              "level": 9,
+              "message": "Daily message limit exceeded"
+            }
+          }
+        ]
+      },
+      {
+        "type": "contract_call",
+        "contract_name": "receive_whatsapp_message",
+        "args": {
+          "conversation_id": "{{message.conversation_id}}",
+          "sender": "{{message.sender}}",
+          "content": "{{message.text}}",
+          "media_type": "{{message.media_type}}",
+          "timestamp": "{{timestamp_now}}",
+          "translated": "{{temp.translated_message.translated || ''}}",
+          "transcription": "{{temp.transcription.transcription || ''}}"
+        }
+      }
+    ]
+  },
+
+  {
+    "id": "daily_cleanup_scheduler",
+    "timestamp": "2025-06-04T10:05:46Z",
+    "type": "continuous_rule",
+    "description": "Executa limpeza diária à meia-noite",
+    "trigger": {
+      "type": "scheduler",
+      "cron": "0 0 * * *"
+    },
+    "effects": [
+      {
+        "type": "contract_call",
+        "contract_name": "rotate_logs"
+      },
+      {
+        "type": "contract_call",
+        "contract_name": "archive_spans"
+      },
+      {
+        "type": "contract_call",
+        "contract_name": "expire_state"
+      },
+      {
+        "type": "state_mutation",
+        "path": "message_counter",
+        "value": 0
+      },
+      {
+        "type": "audit_log",
+        "level": "info",
+        "message": "Daily cleanup completed"
+      }
+    ]
+  }
+]
\ No newline at end of file
diff --git a/flipapp_runtime_full/package.json b/flipapp_runtime_full/package.json
new file mode 100644
index 0000000000000000000000000000000000000000..aa4ff02f3aa13f5b559152581f3faacb160e0f55
--- /dev/null
+++ b/flipapp_runtime_full/package.json
@@ -0,0 +1,17 @@
+--- a/package.json
++++ b/package.json
+@@ -13,11 +13,6 @@
+     "build:wasm:release": "cd rust_vm && wasm-pack build --target web --release",
+     "validate:logline": "node scripts/validate-logline.js ui/",
+     "docker:build": "docker build -t flipapp .",
+-    "docker:run": "docker run -p 80:80 flipapp",
+-    "docker:compose": "docker-compose up",
+-    "clean": "rimraf dist/ public/wasm/"
+-  },
+-  "dependencies": {
++    "clean": "rimraf dist/ public/wasm/"
++  },
++  "dependencies": {
+     "dompurify": "3.0.6",
+     "markdown-it": "13.0.1"
+   },
\ No newline at end of file
diff --git a/flipapp_runtime_full/public/index.html b/flipapp_runtime_full/public/index.html
new file mode 100644
index 0000000000000000000000000000000000000000..0967c9c87a2f00037a51575ce95a0311c70c229a
--- /dev/null
+++ b/flipapp_runtime_full/public/index.html
@@ -0,0 +1,38 @@
+<!DOCTYPE html>
+<html lang="pt-br">
+<head>
+  <meta charset="UTF-8">
+  <meta name="viewport" content="width=device-width, initial-scale=1.0">
+  <title>Loading LogLineOS...</title>
+  <!-- Styles for initial loading only -->
+  <style>
+    body { margin: 0; font-family: sans-serif; display: flex; justify-content: center; align-items: center; height: 100vh; background-color: #f0f4f8; color: #212121; }
+    .loading-container { text-align: center; }
+    .loading-spinner { width: 40px; height: 40px; border: 4px solid rgba(0, 0, 0, 0.1); border-radius: 50%; border-top-color: #1a8cff; animation: spin 1s linear infinite; margin: 0 auto 15px; }
+    @keyframes spin { to { transform: rotate(360deg); } }
+  </style>
+</head>
+<body>
+  <div id="loglineos-root">
+    <div class="loading-container">
+      <div class="loading-spinner"></div>
+      <p>Iniciando LogLineOS...</p>
+    </div>
+  </div>
+  <!-- O Executor (binário WASM) é carregado aqui.
+       Exemplo para Go/WASM: -->
+  <script src="/path/to/wasm_exec.js"></script> <!-- Boilerplate para Go/WASM -->
+  <script>
+    // Carrega o binário WASM do Executor
+    const go = new Go(); // Instância do runtime Go
+    WebAssembly.instantiateStreaming(fetch("/executor.wasm"), go.importObject).then((result) => {
+      go.run(result.instance);
+      // Uma vez que o Executor está rodando, ele assume o controle,
+      // lendo 'logline.yml' e iniciando a aplicação LogLineOS.
+    }).catch(err => {
+      console.error("Error loading LogLineOS Executor WASM:", err);
+      document.getElementById('loglineos-root').innerHTML = "<h1>Error</h1><p>Failed to load LogLineOS. Check console.</p>";
+    });
+  </script>
+</body>
+</html>
\ No newline at end of file
diff --git a/flipapp_runtime_full/public/manifest.json b/flipapp_runtime_full/public/manifest.json
new file mode 100644
index 0000000000000000000000000000000000000000..cc4bbf93386dba918d5b98d7e0a4a0b653303f41
--- /dev/null
+++ b/flipapp_runtime_full/public/manifest.json
@@ -0,0 +1,73 @@
+{
+  "name": "FlipApp",
+  "short_name": "FlipApp",
+  "description": "Plataforma conversacional com UI declarativa e integração WASM",
+  "display": "standalone",
+  "start_url": "/",
+  "background_color": "#ffffff",
+  "theme_color": "#1a8cff",
+  "orientation": "portrait-primary",
+  "categories": ["productivity", "utilities"],
+  "lang": "pt-BR",
+  "dir": "ltr",
+  "icons": [
+    {
+      "src": "/icons/192.png",
+      "sizes": "192x192",
+      "type": "image/png",
+      "purpose": "any maskable"
+    },
+    {
+      "src": "/icons/512.png",
+      "sizes": "512x512",
+      "type": "image/png",
+      "purpose": "any maskable"
+    }
+  ],
+  "shortcuts": [
+    {
+      "name": "Chat",
+      "url": "/?tab=chat",
+      "description": "Abrir chat",
+      "icons": [{ "src": "/icons/chat.png", "sizes": "96x96" }]
+    },
+    {
+      "name": "Espelho",
+      "url": "/?tab=espelho",
+      "description": "Ver espelho",
+      "icons": [{ "src": "/icons/espelho.png", "sizes": "96x96" }]
+    }
+  ],
+  "screenshots": [
+    {
+      "src": "/images/screenshot1.png",
+      "sizes": "1280x720",
+      "type": "image/png",
+      "label": "Tela principal do FlipApp"
+    },
+    {
+      "src": "/images/screenshot2.png",
+      "sizes": "1280x720",
+      "type": "image/png",
+      "label": "Chat no FlipApp"
+    }
+  ],
+  "related_applications": [],
+  "prefer_related_applications": false,
+  "share_target": {
+    "action": "/share-target",
+    "method": "POST",
+    "enctype": "multipart/form-data",
+    "params": {
+      "title": "title",
+      "text": "text",
+      "url": "url",
+      "files": [
+        {
+          "name": "media",
+          "accept": ["image/*", "video/*"]
+        }
+      ]
+    }
+  }
+}
\ No newline at end of file
diff --git a/flipapp_runtime_full/public/service-worker.js b/flipapp_runtime_full/public/service-worker.js
new file mode 100644
index 0000000000000000000000000000000000000000..e97eaa406e31af98e1dc5be248187327685f02ff
--- /dev/null
+++ b/flipapp_runtime_full/public/service-worker.js
@@ -0,0 +1,7 @@
+@@ const ASSET_URLS = [
+-  '/ui/chat_ui.logline',
+-  '/ui/components/button.logline',
++  '/ui/chat_ui.logline',
++  '/ui/whatsapp_ui.logline',
++  '/ui/whatsapp_styles.logline',
++  '/ui/components/button.logline',
\ No newline at end of file
diff --git a/flipapp_runtime_full/public/src/core/animation-system.js b/flipapp_runtime_full/public/src/core/animation-system.js
new file mode 100644
index 0000000000000000000000000000000000000000..138018c1015acb3e8e0a865fbe3c4cbef871cdab
--- /dev/null
+++ b/flipapp_runtime_full/public/src/core/animation-system.js
@@ -0,0 +1,49 @@
+// public/src/core/animation-system.js
+/**
+ * Sistema de animações e feedback sensível à acessibilidade
+ */
+export class AnimationSystem {
+  constructor() {
+    this.prefersReducedMotion = window.matchMedia("(prefers-reduced-motion: reduce)").matches;
+  }
+
+  animateElement(target, animation) {
+    const element = typeof target === 'string' ? document.querySelector(target) : target;
+    if (!element) return;
+
+    if (this.prefersReducedMotion) {
+      // Aplica estilo final somente se movimento for reduzido
+      if (animation.keyframes && animation.keyframes.length) {
+        Object.assign(element.style, animation.keyframes[animation.keyframes.length - 1]);
+      }
+      return;
+    }
+
+    element.animate(
+      animation.keyframes,
+      {
+        duration: animation.duration || 300,
+        easing: animation.easing || 'ease-in-out',
+        fill: 'forwards'
+      }
+    );
+  }
+
+  scrollIntoView(selector, options = {}) {
+    const el = typeof selector === 'string' ? document.querySelector(selector) : selector;
+    if (el) el.scrollIntoView({ behavior: this.prefersReducedMotion ? 'auto' : (options.behavior || 'smooth'), block: options.block || 'nearest' });
+  }
+
+  playSound(sound) {
+    // Caminho relativo a /sounds/
+    const audio = new Audio(`/sounds/${sound}.mp3`);
+    audio.play().catch(() => {});
+  }
+
+  hapticImpact(pattern = 10) {
+    if (navigator.vibrate) navigator.vibrate(pattern);
+  }
+}
+
+// Singleton
+export const animationSystem = new AnimationSystem();
\ No newline at end of file
diff --git a/flipapp_runtime_full/public/src/core/config.js b/flipapp_runtime_full/public/src/core/config.js
new file mode 100644
index 0000000000000000000000000000000000000000..c4aac15519e62d208a9d76d314207831b7b6f86c
--- /dev/null
+++ b/flipapp_runtime_full/public/src/core/config.js
@@ -0,0 +1,32 @@
+// public/src/core/config.js
+/**
+ * Configuração global do FlipApp, injetada via window.APP_CONFIG ou .env
+ */
+const defaultConfig = {
+  apiBaseUrl: 'http://localhost:8000/api',
+  whatsappWsUrl: 'ws://localhost:3001/ws/messages',
+  espelhoWsUrl: 'ws://localhost:8080/ws/espelho',
+  reconnectInterval: 5000,
+  environment: 'development',
+  version: '1.0.0'
+};
+
+function loadConfig() {
+  if (typeof window !== 'undefined' && window.APP_CONFIG) {
+    return { ...defaultConfig, ...window.APP_CONFIG };
+  }
+  if (typeof process !== 'undefined' && process.env) {
+    return {
+      ...defaultConfig,
+      apiBaseUrl: process.env.API_BASE_URL || defaultConfig.apiBaseUrl,
+      whatsappWsUrl: process.env.WHATSAPP_WS_URL || defaultConfig.whatsappWsUrl,
+      espelhoWsUrl: process.env.ESPELHO_WS_URL || defaultConfig.espelhoWsUrl,
+      reconnectInterval: process.env.RECONNECT_INTERVAL ? +process.env.RECONNECT_INTERVAL : defaultConfig.reconnectInterval,
+      environment: process.env.NODE_ENV || defaultConfig.environment,
+      version: process.env.FLIPAPP_VERSION || defaultConfig.version
+    };
+  }
+  return defaultConfig;
+}
+
+export const appConfig = loadConfig();
\ No newline at end of file
diff --git a/flipapp_runtime_full/public/src/core/contracts.js b/flipapp_runtime_full/public/src/core/contracts.js
new file mode 100644
index 0000000000000000000000000000000000000000..9b2f84d94e349659ae4fb1843b98ecf93e227f6e
--- /dev/null
+++ b/flipapp_runtime_full/public/src/core/contracts.js
@@ -0,0 +1,614 @@
+// public/src/core/contracts.js
+/**
+ * Sistema de contratos refinado com middleware, validação e auditoria
+ */
+import { stateManager } from './state.js';
+import { indexedDBBridge } from '../wasm/indexeddb-bridge.js';
+
+export class ContractSystem {
+  constructor() {
+    this.contracts = new Map();
+    this.middleware = [];
+    this.validationRules = new Map();
+    this.executionHistory = [];
+    this.maxHistorySize = 1000;
+    
+    // Registra contratos padrão
+    this.registerDefaultContracts();
+  }
+
+  /**
+   * Registra um novo contrato
+   * @param {string} name - Nome do contrato
+   * @param {function} handler - Função do contrato
+   * @param {Object} options - Opções (validação, middleware, etc.)
+   */
+  register(name, handler, options = {}) {
+    if (typeof handler !== 'function') {
+      throw new Error(`Handler do contrato ${name} deve ser uma função`);
+    }
+
+    this.contracts.set(name, {
+      handler,
+      options: {
+        requiresAuth: false,
+        rateLimit: null,
+        validation: null,
+        async: false,
+        ...options
+      }
+    });
+
+    if (options.validation) {
+      this.validationRules.set(name, options.validation);
+    }
+
+    console.debug(`[Contracts] Contrato registrado: ${name}`);
+  }
+
+  /**
+   * Remove um contrato
+   * @param {string} name - Nome do contrato
+   */
+  unregister(name) {
+    this.contracts.delete(name);
+    this.validationRules.delete(name);
+    console.debug(`[Contracts] Contrato removido: ${name}`);
+  }
+
+  /**
+   * Executa um contrato
+   * @param {string} name - Nome do contrato
+   * @param {Object} params - Parâmetros do contrato
+   * @param {Object} context - Contexto de execução
+   * @returns {Promise<*>} Resultado da execução
+   */
+  async execute(name, params = {}, context = {}) {
+    const startTime = performance.now();
+    const executionId = `${name}_${Date.now()}_${Math.random().toString(36).substr(2, 9)}`;
+
+    try {
+      // Verifica se contrato existe
+      if (!this.contracts.has(name)) {
+        throw new ContractError(`Contrato não encontrado: ${name}`, 'CONTRACT_NOT_FOUND');
+      }
+
+      const contract = this.contracts.get(name);
+      const { handler, options } = contract;
+
+      // Execução de middleware "before"
+      const middlewareContext = { name, params, context, executionId };
+      await this.executeMiddleware('before', middlewareContext);
+
+      // Validação de parâmetros
+      if (options.validation) {
+        this.validateParams(name, params);
+      }
+
+      // Verificação de rate limiting
+      if (options.rateLimit) {
+        await this.checkRateLimit(name, options.rateLimit);
+      }
+
+      // Execução do contrato
+      let result;
+      if (options.async) {
+        result = await handler(params, context, this);
+      } else {
+        result = handler(params, context, this);
+      }
+
+      const endTime = performance.now();
+      const executionTime = endTime - startTime;
+
+      // Log de auditoria
+      const auditSpan = {
+        id: executionId,
+        type: 'contract_executed',
+        contract: name,
+        params: this.sanitizeParams(params),
+        result: this.sanitizeResult(result),
+        execution_time_ms: executionTime,
+        timestamp: new Date().toISOString(),
+        session_id: stateManager.getState('session.id'),
+        user_id: context.userId || 'anonymous',
+        success: true
+      };
+
+      // Salva span de auditoria
+      stateManager.appendState('audit.spans', auditSpan);
+      this.addToHistory(auditSpan);
+
+      // Persiste no IndexedDB se disponível
+      try {
+        await indexedDBBridge.addSpan(auditSpan);
+      } catch (error) {
+        console.warn('[Contracts] Falha ao persistir span de auditoria:', error);
+      }
+
+      // Execução de middleware "after"
+      middlewareContext.result = result;
+      middlewareContext.executionTime = executionTime;
+      await this.executeMiddleware('after', middlewareContext);
+
+      console.debug(`[Contracts] Contrato ${name} executado com sucesso em ${executionTime.toFixed(2)}ms`);
+      return result;
+
+    } catch (error) {
+      const endTime = performance.now();
+      const executionTime = endTime - startTime;
+
+      // Log de erro
+      const errorSpan = {
+        id: executionId,
+        type: 'contract_execution_error',
+        contract: name,
+        params: this.sanitizeParams(params),
+        error: error.message,
+        error_type: error.constructor.name,
+        execution_time_ms: executionTime,
+        timestamp: new Date().toISOString(),
+        session_id: stateManager.getState('session.id'),
+        user_id: context.userId || 'anonymous',
+        success: false
+      };
+
+      stateManager.appendState('audit.spans', errorSpan);
+      this.addToHistory(errorSpan);
+
+      // Execução de middleware "error"
+      const middlewareContext = { name, params, context, executionId, error };
+      await this.executeMiddleware('error', middlewareContext);
+
+      console.error(`[Contracts] Erro no contrato ${name}:`, error);
+      throw error;
+    }
+  }
+
+  /**
+   * Adiciona middleware
+   * @param {string} phase - Fase (before, after, error)
+   * @param {function} handler - Handler do middleware
+   */
+  addMiddleware(phase, handler) {
+    if (!['before', 'after', 'error'].includes(phase)) {
+      throw new Error(`Fase de middleware inválida: ${phase}`);
+    }
+
+    if (!this.middleware[phase]) {
+      this.middleware[phase] = [];
+    }
+
+    this.middleware[phase].push(handler);
+  }
+
+  /**
+   * Executa middleware para uma fase
+   * @param {string} phase - Fase do middleware
+   * @param {Object} context - Contexto
+   */
+  async executeMiddleware(phase, context) {
+    const handlers = this.middleware[phase] || [];
+    
+    for (const handler of handlers) {
+      try {
+        await handler(context);
+      } catch (error) {
+        console.error(`[Contracts] Erro em middleware ${phase}:`, error);
+        if (phase === 'before') {
+          throw error; // Para middleware before, falha para a execução
+        }
+      }
+    }
+  }
+
+  /**
+   * Valida parâmetros do contrato
+   * @param {string} contractName - Nome do contrato
+   * @param {Object} params - Parâmetros a validar
+   */
+  validateParams(contractName, params) {
+    const rules = this.validationRules.get(contractName);
+    if (!rules) return;
+
+    const errors = [];
+
+    Object.entries(rules).forEach(([field, rule]) => {
+      const value = params[field];
+
+      if (rule.required && (value === undefined || value === null || value === '')) {
+        errors.push(`Campo obrigatório: ${field}`);
+        return;
+      }
+
+      if (value !== undefined && value !== null) {
+        if (rule.type && typeof value !== rule.type) {
+          errors.push(`Campo ${field} deve ser do tipo ${rule.type}`);
+        }
+
+        if (rule.minLength && value.length < rule.minLength) {
+          errors.push(`Campo ${field} deve ter pelo menos ${rule.minLength} caracteres`);
+        }
+
+        if (rule.maxLength && value.length > rule.maxLength) {
+          errors.push(`Campo ${field} deve ter no máximo ${rule.maxLength} caracteres`);
+        }
+
+        if (rule.pattern && !rule.pattern.test(value)) {
+          errors.push(`Campo ${field} não atende ao padrão exigido`);
+        }
+
+        if (rule.validator && !rule.validator(value)) {
+          errors.push(`Campo ${field} não passou na validação customizada`);
+        }
+      }
+    });
+
+    if (errors.length > 0) {
+      throw new ValidationError(`Erro de validação: ${errors.join(', ')}`, errors);
+    }
+  }
+
+  /**
+   * Verifica rate limiting
+   * @param {string} contractName - Nome do contrato
+   * @param {Object} rateLimit - Configuração de rate limit
+   */
+  async checkRateLimit(contractName, rateLimit) {
+    const { maxCalls, windowMs } = rateLimit;
+    const now = Date.now();
+    const windowStart = now - windowMs;
+
+    // Filtra execuções recentes
+    const recentExecutions = this.executionHistory.filter(
+      execution => execution.contract === contractName && 
+                  new Date(execution.timestamp).getTime() > windowStart
+    );
+
+    if (recentExecutions.length >= maxCalls) {
+      throw new RateLimitError(`Rate limit excedido para contrato ${contractName}: ${maxCalls} calls/${windowMs}ms`);
+    }
+  }
+
+  /**
+   * Adiciona execução ao histórico
+   * @param {Object} execution - Dados da execução
+   */
+  addToHistory(execution) {
+    this.executionHistory.push(execution);
+    
+    // Mantém tamanho do histórico
+    if (this.executionHistory.length > this.maxHistorySize) {
+      this.executionHistory = this.executionHistory.slice(-this.maxHistorySize);
+    }
+  }
+
+  /**
+   * Sanitiza parâmetros para log (remove dados sensíveis)
+   * @param {Object} params - Parâmetros originais
+   * @returns {Object} Parâmetros sanitizados
+   */
+  sanitizeParams(params) {
+    const sanitized = { ...params };
+    const sensitiveFields = ['password', 'token', 'apiKey', 'secret', 'private'];
+    
+    Object.keys(sanitized).forEach(key => {
+      if (sensitiveFields.some(field => key.toLowerCase().includes(field))) {
+        sanitized[key] = '[REDACTED]';
+      }
+    });
+
+    return sanitized;
+  }
+
+  /**
+   * Sanitiza resultado para log
+   * @param {*} result - Resultado original
+   * @returns {*} Resultado sanitizado
+   */
+  sanitizeResult(result) {
+    if (typeof result === 'object' && result !== null) {
+      const sanitized = { ...result };
+      delete sanitized.password;
+      delete sanitized.token;
+      delete sanitized.apiKey;
+      return sanitized;
+    }
+    return result;
+  }
+
+  /**
+   * Registra contratos padrão do sistema
+   */
+  registerDefaultContracts() {
+    // Contrato de envio de mensagem
+    this.register('sendMessage', async (params, context) => {
+      const sessionId = stateManager.getState('session.id');
+      const messageContent = params.message || stateManager.getState('chat.messageInput');
+      
+      if (!messageContent?.trim()) {
+        throw new ValidationError('Mensagem não pode estar vazia');
+      }
+
+      // Limpa input
+      stateManager.setState('chat.messageInput', '');
+
+      // Adiciona mensagem do usuário
+      const userMessage = {
+        id: `user_${Date.now()}`,
+        type: 'message',
+        sender: 'user',
+        content: messageContent.trim(),
+        timestamp: new Date().toISOString()
+      };
+      
+      stateManager.appendState('chat.history', userMessage);
+
+      // Indica que bot está digitando
+      stateManager.setState('chat.isTyping', true);
+
+      try {
+        // Chama LLM via WASM
+        let llmResponse = '';
+        if (typeof window.wasm_process_prompt !== 'undefined') {
+          llmResponse = await window.wasm_process_prompt(sessionId, messageContent);
+        } else {
+          llmResponse = `Resposta simulada para: "${messageContent}"`;
+        }
+
+        // Simula delay de processamento
+        await new Promise(resolve => setTimeout(resolve, 1000));
+
+        // Para digitação e adiciona resposta
+        stateManager.setState('chat.isTyping', false);
+        
+        const botMessage = {
+          id: `bot_${Date.now()}`,
+          type: 'message',
+          sender: 'bot',
+          content: llmResponse,
+          timestamp: new Date().toISOString()
+        };
+        
+        stateManager.appendState('chat.history', botMessage);
+
+        // Persiste histórico
+        const fullHistory = stateManager.getState('chat.history');
+        try {
+          await indexedDBBridge.saveChatHistory(sessionId, fullHistory);
+        } catch (error) {
+          console.warn('[Contracts] Falha ao persistir chat:', error);
+        }
+
+        return { userMessage, botMessage };
+
+      } catch (error) {
+        stateManager.setState('chat.isTyping', false);
+        throw error;
+      }
+    }, {
+      async: true,
+      validation: {
+        message: { type: 'string', maxLength: 1000 }
+      },
+      rateLimit: { maxCalls: 10, windowMs: 60000 } // 10 mensagens por minuto
+    });
+
+    // Contrato de mudança de aba
+    this.register('switchTab', (params) => {
+      const { tabId } = params;
+      if (!tabId) {
+        throw new ValidationError('tabId é obrigatório');
+      }
+
+      stateManager.setState('activeTab', tabId);
+      
+      return { previousTab: stateManager.getState('activeTab'), newTab: tabId };
+    }, {
+      validation: {
+        tabId: { type: 'string', required: true }
+      }
+    });
+
+    // Contrato de commit de pagamento
+    this.register('commitPaymentContract', async (params) => {
+      const contractData = {
+        amount: params.amount,
+        recipient: params.recipient,
+        currency: params.currency || 'BRL',
+        description: params.description || '',
+        timestamp: new Date().toISOString()
+      };
+
+      // Serializa para WASM
+      const contractYaml = `
+type: payment_request
+amount: ${contractData.amount}
+recipient: ${contractData.recipient}
+currency: ${contractData.currency}
+description: "${contractData.description}"
+timestamp: ${contractData.timestamp}
+      `.trim();
+
+      let commitResult;
+      try {
+        if (typeof window.wasm_commit_contract !== 'undefined') {
+          commitResult = await window.wasm_commit_contract(contractYaml);
+        } else {
+          commitResult = JSON.stringify({ 
+            status: "simulated_ok", 
+            span_id: `sim-${Date.now()}`,
+            message: "Contrato simulado processado"
+          });
+        }
+
+        const result = JSON.parse(commitResult);
+        stateManager.setState('payment.lastStatus', result.status);
+
+        // Persiste contrato
+        try {
+          await indexedDBBridge.saveContract({
+            id: result.span_id,
+            session_id: stateManager.getState('session.id'),
+            type: 'payment_request',
+            data: contractData,
+            status: result.status,
+            created_at: new Date().toISOString()
+          });
+        } catch (error) {
+          console.warn('[Contracts] Falha ao persistir contrato:', error);
+        }
+
+        return result;
+
+      } catch (error) {
+        stateManager.setState('payment.lastStatus', 'error');
+        throw new ContractError(`Falha ao processar contrato: ${error.message}`);
+      }
+    }, {
+      async: true,
+      validation: {
+        amount: { type: 'number', required: true, validator: (v) => v > 0 },
+        recipient: { type: 'string', required: true, minLength: 3 },
+        currency: { type: 'string', pattern: /^[A-Z]{3}$/ }
+      }
+    });
+
+    // Contrato de debounce para digitação
+    this.register('userTypingDebounce', (() => {
+      let timeoutId = null;
+      
+      return () => {
+        // Cancela timeout anterior
+        if (timeoutId) {
+          clearTimeout(timeoutId);
+        }
+
+        stateManager.setState('chat.isTyping', true);
+
+        // Define novo timeout
+        timeoutId = setTimeout(() => {
+          stateManager.setState('chat.isTyping', false);
+          timeoutId = null;
+        }, 3000);
+      };
+    })(), {
+      rateLimit: { maxCalls: 20, windowMs: 10000 } // 20 calls por 10 segundos
+    });
+
+    // Contrato de debug de estado
+    this.register('logState', (params) => {
+      const path = params.path || '*';
+      const state = stateManager.getState(path);
+      console.log(`[Debug] Estado em '${path}':`, state);
+      return state;
+    });
+
+    // Contrato de limpeza de cache
+    this.register('clearCache', async () => {
+      try {
+        // Limpa localStorage
+        localStorage.clear();
+        
+        // Limpa IndexedDB
+        await indexedDBBridge.cleanup(0); // Remove tudo
+        
+        // Reseta estado
+        stateManager.resetState();
+        
+        console.log('[Contracts] Cache limpo com sucesso');
+        return { success: true, message: 'Cache limpo' };
+      } catch (error) {
+        throw new ContractError(`Falha ao limpar cache: ${error.message}`);
+      }
+    }, { async: true });
+  }
+
+  /**
+   * Lista contratos registrados
+   * @returns {Array} Lista de nomes de contratos
+   */
+  listContracts() {
+    return Array.from(this.contracts.keys());
+  }
+
+  /**
+   * Obtém informações de um contrato
+   * @param {string} name - Nome do contrato
+   * @returns {Object|null} Informações do contrato
+   */
+  getContractInfo(name) {
+    const contract = this.contracts.get(name);
+    if (!contract) return null;
+
+    return {
+      name,
+      options: contract.options,
+      validation: this.validationRules.get(name) || null
+    };
+  }
+
+  /**
+   * Obtém estatísticas de execução
+   * @param {string} contractName - Nome do contrato (opcional)
+   * @returns {Object} Estatísticas
+   */
+  getExecutionStats(contractName = null) {
+    let executions = this.executionHistory;
+    
+    if (contractName) {
+      executions = executions.filter(e => e.contract === contractName);
+    }
+
+    const successful = executions.filter(e => e.success);
+    const failed = executions.filter(e => !e.success);
+    const avgExecutionTime = executions.reduce((sum, e) => sum + e.execution_time_ms, 0) / executions.length || 0;
+
+    return {
+      total: executions.length,
+      successful: successful.length,
+      failed: failed.length,
+      successRate: executions.length > 0 ? (successful.length / executions.length) * 100 : 0,
+      avgExecutionTime: Math.round(avgExecutionTime * 100) / 100,
+      contracts: contractName ? [contractName] : [...new Set(executions.map(e => e.contract))]
+    };
+  }
+}
+
+// Classes de erro customizadas
+export class ContractError extends Error {
+  constructor(message, code = 'CONTRACT_ERROR') {
+    super(message);
+    this.name = 'ContractError';
+    this.code = code;
+  }
+}
+
+export class ValidationError extends Error {
+  constructor(message, fields = []) {
+    super(message);
+    this.name = 'ValidationError';
+    this.fields = fields;
+  }
+}
+
+export class RateLimitError extends Error {
+  constructor(message) {
+    super(message);
+    this.name = 'RateLimitError';
+  }
+}
+
+// Singleton global
+export const contractSystem = new ContractSystem();
+
+// Middleware de desenvolvimento
+if (process.env.NODE_ENV === 'development') {
+  contractSystem.addMiddleware('before', (context) => {
+    console.debug(`[Contracts] Executando: ${context.name}`, context.params);
+  });
+
+  contractSystem.addMiddleware('after', (context) => {
+    console.debug(`[Contracts] Concluído: ${context.name} (${context.executionTime.toFixed(2)}ms)`);
+  });
+}
\ No newline at end of file
diff --git a/flipapp_runtime_full/public/src/core/gesture-system.js b/flipapp_runtime_full/public/src/core/gesture-system.js
new file mode 100644
index 0000000000000000000000000000000000000000..a30d8c31e7b684227e82acfa473bcd465d1366ad
--- /dev/null
+++ b/flipapp_runtime_full/public/src/core/gesture-system.js
@@ -0,0 +1,59 @@
+// public/src/core/gesture-system.js
+/**
+ * Sistema de gestos (swipe, tap) para desktop e mobile
+ */
+export class GestureSystem {
+  registerSwipe(element, direction, threshold, callback) {
+    let startX, startY, isSwiping = false;
+    const thresholdPx = threshold * window.innerWidth;
+
+    // Touch
+    element.addEventListener('touchstart', (e) => {
+      startX = e.touches[0].clientX;
+      startY = e.touches[0].clientY;
+      isSwiping = true;
+    }, { passive: true });
+
+    element.addEventListener('touchmove', (e) => {
+      if (!isSwiping) return;
+      e.preventDefault();
+    }, { passive: false });
+
+    element.addEventListener('touchend', (e) => {
+      if (!isSwiping) return;
+      this._checkSwipe(e.changedTouches[0].clientX, e.changedTouches[0].clientY, startX, startY, direction, thresholdPx, callback);
+      isSwiping = false;
+    }, { passive: true });
+
+    // Mouse
+    let mouseMoveHandler, mouseUpHandler;
+    element.addEventListener('mousedown', (e) => {
+      if (e.button !== 0) return;
+      startX = e.clientX;
+      startY = e.clientY;
+      isSwiping = true;
+      mouseMoveHandler = () => {};
+      mouseUpHandler = (upEvent) => {
+        if (!isSwiping) return;
+        this._checkSwipe(upEvent.clientX, upEvent.clientY, startX, startY, direction, thresholdPx, callback);
+        isSwiping = false;
+        document.removeEventListener('mousemove', mouseMoveHandler);
+        document.removeEventListener('mouseup', mouseUpHandler);
+      };
+      document.addEventListener('mousemove', mouseMoveHandler);
+      document.addEventListener('mouseup', mouseUpHandler);
+    });
+  }
+
+  _checkSwipe(endX, endY, startX, startY, direction, thresholdPx, callback) {
+    const diffX = endX - startX;
+    const diffY = endY - startY;
+    if (Math.abs(diffX) > Math.abs(diffY)) {
+      if (direction === 'left' && diffX < -thresholdPx) callback();
+      else if (direction === 'right' && diffX > thresholdPx) callback();
+    }
+  }
+}
+
+// Singleton
+export const gestureSystem = new GestureSystem();
\ No newline at end of file
diff --git a/flipapp_runtime_full/public/src/core/parser.js b/flipapp_runtime_full/public/src/core/parser.js
new file mode 100644
index 0000000000000000000000000000000000000000..f14f24323296f9c952489ceed4e83e25ddf3e55e
--- /dev/null
+++ b/flipapp_runtime_full/public/src/core/parser.js
@@ -0,0 +1,275 @@
+// public/src/core/parser.js
+/**
+ * Parser avançado para LogLine DSL com suporte a componentes e validação
+ */
+export class LogLineParser {
+  constructor(componentRegistry = new Map()) {
+    this.componentRegistry = componentRegistry;
+    this.parseErrors = [];
+  }
+
+  /**
+   * Analisa código LogLine e retorna AST
+   * @param {string} source - Código LogLine
+   * @param {string} fileName - Nome do arquivo (para debug)
+   * @returns {Object} AST com blocos e metadados
+   */
+  parse(source, fileName = 'unknown') {
+    this.parseErrors = [];
+    const blocks = [];
+    const lines = source.split('\n').map((line, index) => ({ content: line, number: index + 1 }));
+    const contextStack = [];
+
+    for (const { content: line, number: lineNumber } of lines) {
+      if (!line.trim() || line.trim().startsWith('#')) continue; // Ignora vazias e comentários
+
+      try {
+        this._parseLine(line, lineNumber, blocks, contextStack, fileName);
+      } catch (error) {
+        this.parseErrors.push({
+          line: lineNumber,
+          message: error.message,
+          file: fileName
+        });
+      }
+    }
+
+    // Valida AST final
+    this._validateAST(blocks, fileName);
+
+    return {
+      blocks,
+      errors: this.parseErrors,
+      metadata: {
+        fileName,
+        parsedAt: new Date().toISOString(),
+        componentCount: this._countComponents(blocks)
+      }
+    };
+  }
+
+  _parseLine(line, lineNumber, blocks, contextStack, fileName) {
+    const lineIndent = line.search(/\S|$/);
+
+    if (line.trim().startsWith('- type:')) {
+      this._parseBlockDeclaration(line, lineIndent, lineNumber, blocks, contextStack, fileName);
+    } else if (line.trim().startsWith('- component:')) {
+      this._parseComponentDeclaration(line, lineIndent, lineNumber, blocks, contextStack, fileName);
+    } else if (contextStack.length > 0) {
+      this._parseProperty(line, lineIndent, lineNumber, contextStack, fileName);
+    } else {
+      throw new Error(`Linha órfã sem contexto de bloco: ${line.trim()}`);
+    }
+  }
+
+  _parseBlockDeclaration(line, lineIndent, lineNumber, blocks, contextStack, fileName) {
+    const blockType = line.split(':')[1].trim();
+    
+    if (!blockType) {
+      throw new Error('Tipo de bloco vazio');
+    }
+
+    const newBlock = {
+      type: blockType,
+      properties: {},
+      children: [],
+      metadata: {
+        line: lineNumber,
+        file: fileName,
+        indent: lineIndent
+      }
+    };
+
+    // Desempilha contextos com indentação menor ou igual
+    while (contextStack.length > 0 && lineIndent <= contextStack[contextStack.length - 1].indentLevel) {
+      contextStack.pop();
+    }
+
+    if (contextStack.length > 0) {
+      contextStack[contextStack.length - 1].block.children.push(newBlock);
+    } else {
+      blocks.push(newBlock);
+    }
+
+    contextStack.push({
+      block: newBlock,
+      indentLevel: lineIndent
+    });
+  }
+
+  _parseComponentDeclaration(line, lineIndent, lineNumber, blocks, contextStack, fileName) {
+    const componentName = line.split(':')[1].trim();
+    
+    if (!this.componentRegistry.has(componentName)) {
+      throw new Error(`Componente não registrado: ${componentName}`);
+    }
+
+    const componentTemplate = this.componentRegistry.get(componentName);
+    const componentBlock = {
+      type: 'component',
+      componentName,
+      properties: {},
+      children: [],
+      template: componentTemplate,
+      metadata: {
+        line: lineNumber,
+        file: fileName,
+        indent: lineIndent
+      }
+    };
+
+    // Mesmo fluxo de aninhamento que blocos normais
+    while (contextStack.length > 0 && lineIndent <= contextStack[contextStack.length - 1].indentLevel) {
+      contextStack.pop();
+    }
+
+    if (contextStack.length > 0) {
+      contextStack[contextStack.length - 1].block.children.push(componentBlock);
+    } else {
+      blocks.push(componentBlock);
+    }
+
+    contextStack.push({
+      block: componentBlock,
+      indentLevel: lineIndent
+    });
+  }
+
+  _parseProperty(line, lineIndent, lineNumber, contextStack, fileName) {
+    const currentContext = contextStack[contextStack.length - 1];
+    
+    if (lineIndent <= currentContext.indentLevel) {
+      throw new Error('Propriedade deve estar mais indentada que seu bloco pai');
+    }
+
+    if (!line.includes(':')) {
+      throw new Error('Propriedade deve conter ":"');
+    }
+
+    const [key, ...valueParts] = line.trim().split(':');
+    const trimmedKey = key.trim();
+    
+    if (!trimmedKey) {
+      throw new Error('Chave de propriedade vazia');
+    }
+
+    let value = valueParts.join(':').trim();
+    
+    // Remove aspas externas se presentes
+    if ((value.startsWith('"') && value.endsWith('"')) || 
+        (value.startsWith("'") && value.endsWith("'"))) {
+      value = value.slice(1, -1);
+    }
+
+    // Validação de propriedades críticas
+    this._validateProperty(trimmedKey, value, currentContext.block, lineNumber, fileName);
+
+    currentContext.block.properties[trimmedKey] = value;
+  }
+
+  _validateProperty(key, value, block, lineNumber, fileName) {
+    const validations = {
+      'bind': (v) => {
+        if (!v.includes('|') && !v.match(/^[a-zA-Z_][a-zA-Z0-9_.]*$/)) {
+          throw new Error(`Propriedade 'bind' inválida: ${v}`);
+        }
+      },
+      'on': (v) => {
+        if (!v.match(/^[a-zA-Z]+\s*:\s*[a-zA-Z_][a-zA-Z0-9_]*(\([^)]*\))?(\s*;\s*[a-zA-Z]+\s*:\s*[a-zA-Z_][a-zA-Z0-9_]*(\([^)]*\))?)*$/)) {
+          throw new Error(`Propriedade 'on' com sintaxe inválida: ${v}`);
+        }
+      },
+      'when': (v) => {
+        // Validação básica de expressão
+        if (v.includes('eval') || v.includes('function') || v.includes('=>')) {
+          throw new Error(`Propriedade 'when' contém código perigoso: ${v}`);
+        }
+      }
+    };
+
+    if (validations[key]) {
+      try {
+        validations[key](value);
+      } catch (error) {
+        throw new Error(`${error.message} (linha ${lineNumber})`);
+      }
+    }
+  }
+
+  _validateAST(blocks, fileName) {
+    // Validações globais da AST
+    this._checkForCircularReferences(blocks);
+    this._validateComponentUsage(blocks);
+  }
+
+  _checkForCircularReferences(blocks, visited = new Set()) {
+    blocks.forEach(block => {
+      if (block.type === 'component') {
+        if (visited.has(block.componentName)) {
+          this.parseErrors.push({
+            message: `Referência circular detectada no componente: ${block.componentName}`,
+            file: block.metadata.file,
+            line: block.metadata.line
+          });
+          return;
+        }
+        
+        visited.add(block.componentName);
+        if (block.template && block.template.blocks) {
+          this._checkForCircularReferences(block.template.blocks, new Set(visited));
+        }
+        visited.delete(block.componentName);
+      }
+      
+      if (block.children) {
+        this._checkForCircularReferences(block.children, visited);
+      }
+    });
+  }
+
+  _validateComponentUsage(blocks) {
+    blocks.forEach(block => {
+      if (block.type === 'component' && !this.componentRegistry.has(block.componentName)) {
+        this.parseErrors.push({
+          message: `Componente não registrado: ${block.componentName}`,
+          file: block.metadata.file,
+          line: block.metadata.line
+        });
+      }
+      
+      if (block.children) {
+        this._validateComponentUsage(block.children);
+      }
+    });
+  }
+
+  _countComponents(blocks) {
+    let count = 0;
+    blocks.forEach(block => {
+      if (block.type === 'component') count++;
+      if (block.children) count += this._countComponents(block.children);
+    });
+    return count;
+  }
+
+  /**
+   * Registra um componente reutilizável
+   * @param {string} name - Nome do componente
+   * @param {Object} template - Template AST do componente
+   */
+  registerComponent(name, template) {
+    this.componentRegistry.set(name, template);
+  }
+
+  /**
+   * Obtém erros de parsing formatados
+   * @returns {string} Relatório de erros
+   */
+  getErrorReport() {
+    if (this.parseErrors.length === 0) return 'Nenhum erro encontrado.';
+    
+    return this.parseErrors.map(error => 
+      `${error.file}:${error.line} - ${error.message}`
+    ).join('\n');
+  }
+}
\ No newline at end of file
diff --git a/flipapp_runtime_full/public/src/core/renderer.js b/flipapp_runtime_full/public/src/core/renderer.js
new file mode 100644
index 0000000000000000000000000000000000000000..b80d71d0cad0368087d5249105bef6dd5248355d
--- /dev/null
+++ b/flipapp_runtime_full/public/src/core/renderer.js
@@ -0,0 +1,758 @@
+// public/src/core/renderer.js
+/**
+ * Renderer otimizado com Virtual DOM e diffing para performance
+ */
+import { expressionEngine } from './expression-engine.js';
+import { stateManager } from './state.js';
+
+export class VirtualNode {
+  constructor(type, props = {}, children = [], key = null) {
+    this.type = type;
+    this.props = props;
+    this.children = children;
+    this.key = key;
+    this.ref = null; // Referência ao DOM real
+  }
+}
+
+export class LogLineRenderer {
+  constructor() {
+    this.componentRegistry = new Map();
+    this.mountedComponents = new Map(); // Rastreamento de componentes montados
+    this.eventListeners = new WeakMap(); // Rastreamento de listeners para cleanup
+    this.subscriptions = new WeakMap(); // Rastreamento de subscrições de estado
+  }
+
+  /**
+   * Renderiza AST para Virtual DOM
+   * @param {Object} ast - AST do LogLine
+   * @param {Object} handlers - Handlers de contrato
+   * @param {Object} context - Contexto de dados local
+   * @returns {VirtualNode} Árvore Virtual DOM
+   */
+  render(ast, handlers = {}, context = {}) {
+    if (!ast || !ast.blocks) {
+      console.warn('[Renderer] AST inválida recebida');
+      return new VirtualNode('div', {}, []);
+    }
+
+    const children = ast.blocks.map(block => this.renderBlock(block, handlers, context));
+    return new VirtualNode('div', { class: 'logline-root' }, children);
+  }
+
+  /**
+   * Renderiza um bloco individual
+   * @param {Object} block - Bloco LogLine
+   * @param {Object} handlers - Handlers de contrato
+   * @param {Object} context - Contexto local
+   * @returns {VirtualNode} Nó virtual
+   */
+  renderBlock(block, handlers, context) {
+    if (!block || !block.type) {
+      console.warn('[Renderer] Bloco inválido:', block);
+      return new VirtualNode('div', {}, []);
+    }
+
+    // Processa componentes personalizados
+    if (block.type === 'component') {
+      return this.renderComponent(block, handlers, context);
+    }
+
+    // Processa loops
+    if (block.type === 'loop') {
+      return this.renderLoop(block, handlers, context);
+    }
+
+    // Cria o nó virtual
+    const elementType = this.getElementTag(block.type);
+    const props = this.processProperties(block.properties || {}, context);
+    
+    // Processa condições 'when'
+    if (props.when !== undefined) {
+      const condition = expressionEngine.evaluate(props.when, context);
+      if (!condition) {
+        return new VirtualNode('template', { style: 'display: none' }, []);
+      }
+      delete props.when; // Remove para não passar ao DOM
+    }
+
+    // Renderiza filhos
+    const children = (block.children || []).map(child => 
+      this.renderBlock(child, handlers, context)
+    );
+
+    const vnode = new VirtualNode(elementType, props, children, block.metadata?.key);
+    
+    // Adiciona metadados para debugging
+    if (process.env.NODE_ENV === 'development') {
+      vnode.loglineBlock = block;
+    }
+
+    return vnode;
+  }
+
+  /**
+   * Renderiza um componente reutilizável
+   * @param {Object} block - Bloco de componente
+   * @param {Object} handlers - Handlers
+   * @param {Object} context - Contexto
+   * @returns {VirtualNode} Nó virtual do componente
+   */
+  renderComponent(block, handlers, context) {
+    const { componentName, properties = {}, template } = block;
+    
+    if (!this.componentRegistry.has(componentName) && !template) {
+      console.error(`[Renderer] Componente não registrado: ${componentName}`);
+      return new VirtualNode('div', { 
+        class: 'error-component',
+        'data-error': `Componente não encontrado: ${componentName}`
+      }, []);
+    }
+
+    const componentTemplate = template || this.componentRegistry.get(componentName);
+    
+    // Cria contexto de props para o componente
+    const componentContext = {
+      ...context,
+      props: this.processProperties(properties, context)
+    };
+
+    // Renderiza template do componente
+    if (componentTemplate.blocks) {
+      const children = componentTemplate.blocks.map(childBlock => 
+        this.renderBlock(childBlock, handlers, componentContext)
+      );
+      
+      return new VirtualNode('div', { 
+        class: `component-${componentName}`,
+        'data-component': componentName
+      }, children);
+    }
+
+    return new VirtualNode('div', {}, []);
+  }
+
+  /**
+   * Renderiza um loop sobre dados
+   * @param {Object} block - Bloco de loop
+   * @param {Object} handlers - Handlers
+   * @param {Object} context - Contexto
+   * @returns {VirtualNode} Container com itens do loop
+   */
+  renderLoop(block, handlers, context) {
+    const { properties = {}, children = [] } = block;
+    const dataPath = properties.data;
+    
+    if (!dataPath) {
+      console.warn('[Renderer] Loop sem propriedade data:', block);
+      return new VirtualNode('div', {}, []);
+    }
+
+    // Resolve dados do loop
+    const loopData = expressionEngine.evaluate(dataPath, context);
+    
+    if (!Array.isArray(loopData)) {
+      console.warn(`[Renderer] Dados de loop não são array: ${dataPath}`, loopData);
+      return new VirtualNode('div', {}, []);
+    }
+
+    // Renderiza cada item
+    const loopItems = loopData.map((item, index) => {
+      const itemContext = {
+        ...context,
+        item,
+        index,
+        $index: index,
+        $first: index === 0,
+        $last: index === loopData.length - 1,
+        $even: index % 2 === 0,
+        $odd: index % 2 === 1
+      };
+
+      // Renderiza filhos para este item
+      const itemChildren = children.map(child => 
+        this.renderBlock(child, handlers, itemContext)
+      );
+
+      // Usa key único para otimização de diffing
+      const key = item.id || item.key || `loop-item-${index}`;
+      
+      return new VirtualNode('div', { 
+        class: 'loop-item',
+        'data-loop-index': index
+      }, itemChildren, key);
+    });
+
+    return new VirtualNode('div', { 
+      class: 'loop-container',
+      'data-loop-source': dataPath
+    }, loopItems);
+  }
+
+  /**
+   * Processa propriedades do bloco
+   * @param {Object} properties - Propriedades brutas
+   * @param {Object} context - Contexto para interpolação
+   * @returns {Object} Propriedades processadas
+   */
+  processProperties(properties, context) {
+    const processed = {};
+    
+    Object.entries(properties).forEach(([key, value]) => {
+      if (typeof value === 'string') {
+        processed[key] = expressionEngine.interpolate(value, context);
+      } else {
+        processed[key] = value;
+      }
+    });
+
+    return processed;
+  }
+
+  /**
+   * Monta Virtual DOM no DOM real com diffing otimizado
+   * @param {VirtualNode} vnode - Nó virtual
+   * @param {Element} container - Container DOM
+   * @param {Object} handlers - Handlers de evento
+   */
+  mount(vnode, container, handlers = {}) {
+    const element = this.createDOMElement(vnode, handlers);
+    
+    // Limpa container e monta
+    this.unmountAll(container);
+    container.appendChild(element);
+    
+    // Registra elemento para tracking
+    vnode.ref = element;
+    
+    return element;
+  }
+
+  /**
+   * Atualiza DOM usando algoritmo de diffing
+   * @param {VirtualNode} oldVNode - Árvore virtual anterior
+   * @param {VirtualNode} newVNode - Nova árvore virtual
+   * @param {Element} container - Container DOM
+   * @param {Object} handlers - Handlers
+   */
+  patch(oldVNode, newVNode, container, handlers = {}) {
+    if (!oldVNode) {
+      return this.mount(newVNode, container, handlers);
+    }
+
+    if (!newVNode) {
+      this.unmount(oldVNode);
+      return null;
+    }
+
+    if (this.shouldReplace(oldVNode, newVNode)) {
+      const newElement = this.createDOMElement(newVNode, handlers);
+      const oldElement = oldVNode.ref;
+      
+      if (oldElement && oldElement.parentNode) {
+        this.unmount(oldVNode);
+        oldElement.parentNode.replaceChild(newElement, oldElement);
+      }
+      
+      newVNode.ref = newElement;
+      return newElement;
+    }
+
+    // Atualiza nó existente
+    this.updateElement(oldVNode, newVNode, handlers);
+    this.patchChildren(oldVNode, newVNode, handlers);
+    
+    newVNode.ref = oldVNode.ref;
+    return newVNode.ref;
+  }
+
+  /**
+   * Verifica se deve substituir elemento completamente
+   * @param {VirtualNode} oldVNode - Nó antigo
+   * @param {VirtualNode} newVNode - Nó novo
+   * @returns {boolean} True se deve substituir
+   */
+  shouldReplace(oldVNode, newVNode) {
+    return oldVNode.type !== newVNode.type || 
+           oldVNode.key !== newVNode.key;
+  }
+
+  /**
+   * Atualiza propriedades de um elemento existente
+   * @param {VirtualNode} oldVNode - Nó antigo
+   * @param {VirtualNode} newVNode - Nó novo
+   * @param {Object} handlers - Handlers
+   */
+  updateElement(oldVNode, newVNode, handlers) {
+    const element = oldVNode.ref;
+    if (!element) return;
+
+    // Remove propriedades antigas
+    Object.keys(oldVNode.props).forEach(key => {
+      if (!(key in newVNode.props)) {
+        this.removeProp(element, key, oldVNode.props[key]);
+      }
+    });
+
+    // Adiciona/atualiza propriedades novas
+    Object.entries(newVNode.props).forEach(([key, value]) => {
+      if (oldVNode.props[key] !== value) {
+        this.setProp(element, key, value, handlers);
+      }
+    });
+  }
+
+  /**
+   * Faz patch dos filhos usando algoritmo otimizado
+   * @param {VirtualNode} oldVNode - Nó pai antigo
+   * @param {VirtualNode} newVNode - Nó pai novo
+   * @param {Object} handlers - Handlers
+   */
+  patchChildren(oldVNode, newVNode, handlers) {
+    const element = oldVNode.ref;
+    const oldChildren = oldVNode.children || [];
+    const newChildren = newVNode.children || [];
+
+    // Algoritmo simples de diffing por chave
+    const keyedOld = new Map();
+    const keyedNew = new Map();
+    
+    oldChildren.forEach((child, index) => {
+      if (child.key) {
+        keyedOld.set(child.key, { child, index });
+      }
+    });
+    
+    newChildren.forEach((child, index) => {
+      if (child.key) {
+        keyedNew.set(child.key, { child, index });
+      }
+    });
+
+    // Patches por chave primeiro
+    for (const [key, { child: newChild, index: newIndex }] of keyedNew) {
+      if (keyedOld.has(key)) {
+        const { child: oldChild, index: oldIndex } = keyedOld.get(key);
+        this.patch(oldChild, newChild, element, handlers);
+        
+        // Move elemento se necessário
+        if (newIndex !== oldIndex) {
+          const childElement = newChild.ref || oldChild.ref;
+          if (childElement && element.children[newIndex] !== childElement) {
+            element.insertBefore(childElement, element.children[newIndex] || null);
+          }
+        }
+      } else {
+        // Novo elemento com chave
+        const newElement = this.createDOMElement(newChild, handlers);
+        element.insertBefore(newElement, element.children[newIndex] || null);
+        newChild.ref = newElement;
+      }
+    }
+
+    // Remove elementos com chave que não existem mais
+    for (const [key, { child: oldChild }] of keyedOld) {
+      if (!keyedNew.has(key)) {
+        this.unmount(oldChild);
+      }
+    }
+
+    // Patch elementos sem chave (fallback simples)
+    const maxLength = Math.max(oldChildren.length, newChildren.length);
+    for (let i = 0; i < maxLength; i++) {
+      const oldChild = oldChildren[i];
+      const newChild = newChildren[i];
+      
+      if (oldChild && !oldChild.key && newChild && !newChild.key) {
+        this.patch(oldChild, newChild, element, handlers);
+      } else if (!oldChild && newChild && !newChild.key) {
+        const newElement = this.createDOMElement(newChild, handlers);
+        element.appendChild(newElement);
+        newChild.ref = newElement;
+      } else if (oldChild && !oldChild.key && !newChild) {
+        this.unmount(oldChild);
+      }
+    }
+  }
+
+  /**
+   * Cria elemento DOM a partir de VirtualNode
+   * @param {VirtualNode} vnode - Nó virtual
+   * @param {Object} handlers - Handlers de evento
+   * @returns {Element} Elemento DOM
+   */
+  createDOMElement(vnode, handlers) {
+    if (vnode.type === 'text') {
+      const text = vnode.props.content || '';
+      return document.createTextNode(text);
+    }
+
+    const element = document.createElement(vnode.type);
+    
+    // Aplica propriedades
+    Object.entries(vnode.props).forEach(([key, value]) => {
+      this.setProp(element, key, value, handlers);
+    });
+
+    // Cria filhos
+    vnode.children.forEach(child => {
+      const childElement = this.createDOMElement(child, handlers);
+      element.appendChild(childElement);
+      child.ref = childElement;
+    });
+
+    return element;
+  }
+
+  /**
+   * Define propriedade no elemento DOM
+   * @param {Element} element - Elemento DOM
+   * @param {string} key - Nome da propriedade
+   * @param {*} value - Valor da propriedade
+   * @param {Object} handlers - Handlers de evento
+   */
+  setProp(element, key, value, handlers) {
+    switch (key) {
+      case 'style':
+        if (typeof value === 'string') {
+          element.style.cssText = value;
+        } else if (typeof value === 'object') {
+          Object.assign(element.style, value);
+        }
+        break;
+        
+      case 'class':
+      case 'className':
+        element.className = value;
+        break;
+        
+      case 'on':
+        this.applyEventHandlers(element, value, handlers);
+        break;
+        
+      case 'bind':
+        this.applyDataBinding(element, value, handlers);
+        break;
+        
+      case 'content':
+        if (element.tagName === 'INPUT' || element.tagName === 'TEXTAREA') {
+          element.value = value;
+        } else {
+          element.textContent = value;
+        }
+        break;
+        
+      case 'innerHTML':
+        // Sanitização básica - em produção usar DOMPurify
+        if (typeof window.DOMPurify !== 'undefined') {
+          element.innerHTML = window.DOMPurify.sanitize(value);
+        } else {
+          console.warn('[Renderer] DOMPurify não disponível, pulando innerHTML');
+        }
+        break;
+        
+      default:
+        // Propriedades HTML padrão
+        try {
+          if (value === null || value === undefined || value === false) {
+            element.removeAttribute(key);
+          } else if (value === true) {
+            element.setAttribute(key, '');
+          } else {
+            element.setAttribute(key, String(value));
+          }
+        } catch (error) {
+          console.warn(`[Renderer] Erro ao definir propriedade ${key}:`, error);
+        }
+        break;
+    }
+  }
+
+  /**
+   * Remove propriedade do elemento DOM
+   * @param {Element} element - Elemento DOM
+   * @param {string} key - Nome da propriedade
+   * @param {*} value - Valor antigo da propriedade
+   */
+  removeProp(element, key, value) {
+    switch (key) {
+      case 'style':
+        element.style.cssText = '';
+        break;
+        
+      case 'class':
+      case 'className':
+        element.className = '';
+        break;
+        
+      case 'on':
+        this.removeEventHandlers(element);
+        break;
+        
+      case 'bind':
+        this.removeDataBinding(element);
+        break;
+        
+      default:
+        element.removeAttribute(key);
+        break;
+    }
+  }
+
+  /**
+   * Aplica handlers de evento
+   * @param {Element} element - Elemento DOM
+   * @param {string} handlersString - String de handlers
+   * @param {Object} contractHandlers - Handlers de contrato
+   */
+  applyEventHandlers(element, handlersString, contractHandlers) {
+    if (!handlersString || !contractHandlers.onContractCall) return;
+
+    // Remove listeners antigos
+    this.removeEventHandlers(element);
+
+    const handlers = this.parseEventHandlers(handlersString);
+    const elementListeners = [];
+
+    handlers.forEach(({ event, contract, params }) => {
+      const listener = (e) => {
+        e.preventDefault();
+        contractHandlers.onContractCall(contract, {
+          ...params,
+          event: e,
+          element: element,
+          state: stateManager.getState('*')
+        });
+      };
+
+      element.addEventListener(event, listener);
+      elementListeners.push({ event, listener });
+    });
+
+    // Armazena listeners para cleanup
+    this.eventListeners.set(element, elementListeners);
+  }
+
+  /**
+   * Remove handlers de evento
+   * @param {Element} element - Elemento DOM
+   */
+  removeEventHandlers(element) {
+    const listeners = this.eventListeners.get(element);
+    if (listeners) {
+      listeners.forEach(({ event, listener }) => {
+        element.removeEventListener(event, listener);
+      });
+      this.eventListeners.delete(element);
+    }
+  }
+
+  /**
+   * Aplica binding de dados
+   * @param {Element} element - Elemento DOM
+   * @param {string} binding - String de binding
+   * @param {Object} handlers - Handlers
+   */
+  applyDataBinding(element, binding, handlers) {
+    const [statePath, contract] = binding.split('|').map(s => s.trim());
+
+    if (element.tagName === 'INPUT' || element.tagName === 'TEXTAREA') {
+      // Sincroniza valor inicial
+      element.value = stateManager.getState(statePath) || '';
+
+      // Listener para mudanças
+      const listener = (e) => {
+        stateManager.setState(statePath, e.target.value);
+        
+        // Executa contrato se especificado
+        if (contract && handlers.onContractCall) {
+          handlers.onContractCall(contract, {
+            value: e.target.value,
+            element: element,
+            statePath: statePath
+          });
+        }
+      };
+
+      element.addEventListener('input', listener);
+
+      // Subscrição para mudanças externas do estado
+      const subscription = stateManager.subscribe(statePath, (value) => {
+        if (element.value !== value) {
+          element.value = value || '';
+        }
+      });
+
+      // Armazena para cleanup
+      this.subscriptions.set(element, { listener, subscription });
+    }
+  }
+
+  /**
+   * Remove binding de dados
+   * @param {Element} element - Elemento DOM
+   */
+  removeDataBinding(element) {
+    const binding = this.subscriptions.get(element);
+    if (binding) {
+      if (binding.listener) {
+        element.removeEventListener('input', binding.listener);
+      }
+      if (binding.subscription) {
+        binding.subscription(); // Chama função de cleanup
+      }
+      this.subscriptions.delete(element);
+    }
+  }
+
+  /**
+   * Faz parse de string de event handlers
+   * @param {string} handlersString - String de handlers
+   * @returns {Array} Array de handlers parseados
+   */
+  parseEventHandlers(handlersString) {
+    const handlers = [];
+    
+    handlersString.split(';').forEach(handlerEntry => {
+      const parts = handlerEntry.trim().split(':');
+      if (parts.length < 2) return;
+
+      const event = parts[0].trim();
+      const contractCall = parts.slice(1).join(':').trim();
+
+      let contract = contractCall;
+      let params = {};
+
+      // Parse de parâmetros: contract(param1='value', param2='another')
+      const paramMatch = contractCall.match(/^([a-zA-Z0-9_]+)\((.*)\)$/);
+      if (paramMatch) {
+        contract = paramMatch[1];
+        const rawParamsString = paramMatch[2];
+        
+        // Parse simples de parâmetros
+        rawParamsString.split(',').forEach(paramPair => {
+          const [key, rawValue] = paramPair.split('=').map(s => s.trim());
+          if (key && rawValue) {
+            params[key] = rawValue.replace(/^['"]|['"]$/g, '');
+          }
+        });
+      }
+
+      handlers.push({ event, contract, params });
+    });
+
+    return handlers;
+  }
+
+  /**
+   * Desmonta elemento e faz cleanup
+   * @param {VirtualNode} vnode - Nó virtual a ser desmontado
+   */
+  unmount(vnode) {
+    if (!vnode.ref) return;
+
+    // Cleanup de eventos e bindings
+    this.removeEventHandlers(vnode.ref);
+    this.removeDataBinding(vnode.ref);
+
+    // Desmonta filhos recursivamente
+    if (vnode.children) {
+      vnode.children.forEach(child => this.unmount(child));
+    }
+
+    // Remove do DOM
+    if (vnode.ref.parentNode) {
+      vnode.ref.parentNode.removeChild(vnode.ref);
+    }
+
+    vnode.ref = null;
+  }
+
+  /**
+   * Desmonta todos os elementos de um container
+   * @param {Element} container - Container a ser limpo
+   */
+  unmountAll(container) {
+    while (container.firstChild) {
+      // Cleanup recursivo se necessário
+      this.cleanupElement(container.firstChild);
+      container.removeChild(container.firstChild);
+    }
+  }
+
+  /**
+   * Faz cleanup de um elemento DOM
+   * @param {Element} element - Elemento a ser limpo
+   */
+  cleanupElement(element) {
+    // Remove listeners
+    this.eventListeners.delete(element);
+    this.subscriptions.delete(element);
+
+    // Cleanup recursivo dos filhos
+    if (element.children) {
+      Array.from(element.children).forEach(child => this.cleanupElement(child));
+    }
+  }
+
+  /**
+   * Obtém tag HTML para tipo de bloco
+   * @param {string} blockType - Tipo do bloco LogLine
+   * @returns {string} Tag HTML
+   */
+  getElementTag(blockType) {
+    const typeMap = {
+      container: 'div',
+      text: 'p',
+      markdown: 'div',
+      button: 'button',
+      input: 'input',
+      textarea: 'textarea',
+      image: 'img',
+      link: 'a',
+      list: 'ul',
+      listItem: 'li',
+      header: 'h1',
+      section: 'section',
+      article: 'article',
+      nav: 'nav',
+      footer: 'footer',
+      span: 'span',
+      strong: 'strong',
+      em: 'em',
+      loop: 'div',
+      when: 'div',
+      component: 'div'
+    };
+    
+    return typeMap[blockType] || 'div';
+  }
+
+  /**
+   * Registra componente reutilizável
+   * @param {string} name - Nome do componente
+   * @param {Object} template - Template do componente
+   */
+  registerComponent(name, template) {
+    this.componentRegistry.set(name, template);
+  }
+
+  /**
+   * Remove componente registrado
+   * @param {string} name - Nome do componente
+   */
+  unregisterComponent(name) {
+    this.componentRegistry.delete(name);
+  }
+
+  /**
+   * Lista componentes registrados
+   * @returns {Array} Lista de nomes de componentes
+   */
+  getRegisteredComponents() {
+    return Array.from(this.componentRegistry.keys());
+  }
+}
+
+// Singleton global
+export const renderer = new LogLineRenderer();
\ No newline at end of file
diff --git a/flipapp_runtime_full/public/src/core/state.js b/flipapp_runtime_full/public/src/core/state.js
new file mode 100644
index 0000000000000000000000000000000000000000..3d63132579df4d5f5c2bad4ef2e28ffb4d98d659
--- /dev/null
+++ b/flipapp_runtime_full/public/src/core/state.js
@@ -0,0 +1,273 @@
+// public/src/core/state.js
+/**
+ * Sistema de gerenciamento de estado reativo centralizado
+ * Implementa padrão Observer para notificações automáticas de mudanças
+ */
+export class StateManager {
+  constructor() {
+    this.state = {};
+    this.observers = new Map(); // path -> Set<callback>
+    this.middlewares = []; // Para logging, debugging, etc.
+  }
+
+  /**
+   * Define um valor no estado com notificação reativa
+   * @param {string} path - Caminho dot-notation (ex: 'user.profile.name')
+   * @param {*} value - Valor a ser definido
+   */
+  setState(path, value) {
+    const oldValue = this.getState(path);
+    if (oldValue === value) return; // Evita re-renders desnecessários
+
+    this._setNestedValue(path, value);
+    
+    // Executa middlewares
+    this.middlewares.forEach(middleware => {
+      middleware({ type: 'SET_STATE', path, value, oldValue });
+    });
+
+    this.notifyObservers(path, value);
+    this._persistCriticalState(path, value);
+  }
+
+  /**
+   * Obtém valor do estado
+   * @param {string} path - Caminho para a propriedade
+   * @returns {*} Valor ou undefined
+   */
+  getState(path) {
+    if (path === '*') return this.state;
+    return this._getNestedValue(path);
+  }
+
+  /**
+   * Adiciona item a um array no estado
+   * @param {string} path - Caminho para o array
+   * @param {*} value - Item a ser adicionado
+   */
+  appendState(path, value) {
+    const currentArray = this.getState(path);
+    if (Array.isArray(currentArray)) {
+      this.setState(path, [...currentArray, value]);
+    } else {
+      console.warn(`[StateManager] Tentativa de append em não-array: ${path}`);
+      this.setState(path, [value]);
+    }
+  }
+
+  /**
+   * Remove item de array por índice ou condição
+   * @param {string} path - Caminho para o array
+   * @param {number|function} condition - Índice ou função de filtro
+   */
+  removeFromState(path, condition) {
+    const currentArray = this.getState(path);
+    if (!Array.isArray(currentArray)) return;
+
+    let newArray;
+    if (typeof condition === 'number') {
+      newArray = currentArray.filter((_, index) => index !== condition);
+    } else if (typeof condition === 'function') {
+      newArray = currentArray.filter(item => !condition(item));
+    }
+    
+    if (newArray) {
+      this.setState(path, newArray);
+    }
+  }
+
+  /**
+   * Atualiza múltiplos caminhos atomicamente
+   * @param {Object} updates - Objeto com path: value
+   */
+  batchUpdate(updates) {
+    const oldValues = {};
+    
+    // Aplica todas as mudanças sem notificar
+    Object.entries(updates).forEach(([path, value]) => {
+      oldValues[path] = this.getState(path);
+      this._setNestedValue(path, value);
+    });
+
+    // Executa middlewares para cada mudança
+    Object.entries(updates).forEach(([path, value]) => {
+      this.middlewares.forEach(middleware => {
+        middleware({ type: 'BATCH_UPDATE', path, value, oldValue: oldValues[path] });
+      });
+    });
+
+    // Notifica observadores
+    Object.entries(updates).forEach(([path, value]) => {
+      this.notifyObservers(path, value);
+      this._persistCriticalState(path, value);
+    });
+  }
+
+  /**
+   * Subscreve a mudanças de estado
+   * @param {string} path - Caminho a observar
+   * @param {function} callback - Função chamada na mudança
+   * @returns {function} Função de cleanup para remover observador
+   */
+  subscribe(path, callback) {
+    if (!this.observers.has(path)) {
+      this.observers.set(path, new Set());
+    }
+    this.observers.get(path).add(callback);
+
+    // Chama callback imediatamente com valor atual
+    callback(this.getState(path), path);
+
+    // Retorna função de cleanup
+    return () => this.unsubscribe(path, callback);
+  }
+
+  /**
+   * Remove subscrição
+   * @param {string} path - Caminho observado
+   * @param {function} callback - Callback a ser removido
+   */
+  unsubscribe(path, callback) {
+    const observers = this.observers.get(path);
+    if (observers) {
+      observers.delete(callback);
+      if (observers.size === 0) {
+        this.observers.delete(path);
+      }
+    }
+  }
+
+  /**
+   * Notifica observadores de mudanças
+   * @param {string} changedPath - Caminho que mudou
+   * @param {*} newValue - Novo valor
+   */
+  notifyObservers(changedPath, newValue) {
+    // Notifica observadores do caminho exato
+    this._notifyPathObservers(changedPath, newValue);
+
+    // Notifica observadores de caminhos ancestrais
+    let currentPath = changedPath;
+    while (currentPath.includes('.')) {
+      currentPath = currentPath.substring(0, currentPath.lastIndexOf('.'));
+      this._notifyPathObservers(currentPath, this.getState(currentPath));
+    }
+
+    // Notifica observador global
+    this._notifyPathObservers('*', this.state);
+  }
+
+  _notifyPathObservers(path, value) {
+    const observers = this.observers.get(path);
+    if (observers) {
+      observers.forEach(callback => {
+        try {
+          callback(value, path);
+        } catch (error) {
+          console.error(`[StateManager] Erro em observer para ${path}:`, error);
+        }
+      });
+    }
+  }
+
+  /**
+   * Adiciona middleware para interceptar mudanças de estado
+   * @param {function} middleware - Função middleware
+   */
+  addMiddleware(middleware) {
+    this.middlewares.push(middleware);
+  }
+
+  /**
+   * Limpa todo o estado
+   */
+  resetState() {
+    this.state = {};
+    this.observers.clear();
+    localStorage.clear();
+    console.warn('[StateManager] Estado resetado completamente');
+  }
+
+  // Métodos internos
+  _setNestedValue(path, value) {
+    const parts = path.split('.');
+    let current = this.state;
+    
+    for (let i = 0; i < parts.length - 1; i++) {
+      if (!current[parts[i]] || typeof current[parts[i]] !== 'object') {
+        current[parts[i]] = {};
+      }
+      current = current[parts[i]];
+    }
+    
+    current[parts[parts.length - 1]] = value;
+  }
+
+  _getNestedValue(path) {
+    const parts = path.split('.');
+    let current = this.state;
+    
+    for (const part of parts) {
+      if (current === null || typeof current !== 'object' || !current.hasOwnProperty(part)) {
+        return undefined;
+      }
+      current = current[part];
+    }
+    
+    return current;
+  }
+
+  /**
+   * Persiste estados críticos no localStorage
+   * @param {string} path - Caminho alterado
+   * @param {*} value - Novo valor
+   */
+  _persistCriticalState(path, value) {
+    const criticalPaths = ['chat.history', 'user.preferences', 'session.id'];
+    
+    for (const criticalPath of criticalPaths) {
+      if (path.startsWith(criticalPath)) {
+        const sessionId = this.getState('session.id');
+        const key = `flipapp_${criticalPath.replace(/\./g, '_')}_${sessionId || 'default'}`;
+        
+        try {
+          localStorage.setItem(key, JSON.stringify(this.getState(criticalPath)));
+        } catch (error) {
+          console.error(`[StateManager] Falha ao persistir ${criticalPath}:`, error);
+        }
+        break;
+      }
+    }
+  }
+
+  /**
+   * Carrega estados persistidos
+   */
+  loadPersistedState() {
+    const sessionId = this.getState('session.id');
+    const criticalPaths = ['chat.history', 'user.preferences'];
+    
+    criticalPaths.forEach(path => {
+      const key = `flipapp_${path.replace(/\./g, '_')}_${sessionId || 'default'}`;
+      const stored = localStorage.getItem(key);
+      
+      if (stored) {
+        try {
+          this.setState(path, JSON.parse(stored));
+        } catch (error) {
+          console.error(`[StateManager] Falha ao carregar ${path}:`, error);
+        }
+      }
+    });
+  }
+}
+
+// Singleton global
+export const stateManager = new StateManager();
+
+// Middleware de logging para desenvolvimento
+if (process.env.NODE_ENV === 'development') {
+  stateManager.addMiddleware(({ type, path, value, oldValue }) => {
+    console.debug(`[State] ${type}: ${path}`, { oldValue, newValue: value });
+  });
+}
\ No newline at end of file
diff --git a/flipapp_runtime_full/public/src/core/websocket-manager.js b/flipapp_runtime_full/public/src/core/websocket-manager.js
new file mode 100644
index 0000000000000000000000000000000000000000..ccac6708b22cd4eea916d15aa5e72a945699d6e5
--- /dev/null
+++ b/flipapp_runtime_full/public/src/core/websocket-manager.js
@@ -0,0 +1,121 @@
+// public/src/core/websocket-manager.js
+/**
+ * Gerenciamento robusto de WebSockets com reconexão exponencial, eventos e integração reativa
+ */
+
+import { stateManager } from './state.js';
+
+export class WebSocketManager {
+  constructor(url, options = {}) {
+    this.url = url;
+    this.protocols = options.protocols || [];
+    this.label = options.label || 'ws';
+    this.reconnectBaseDelay = options.reconnectBaseDelay || 1000; // ms
+    this.reconnectMaxDelay = options.reconnectMaxDelay || 15000;
+    this.reconnectAttempts = 0;
+    this.handlers = new Map(); // type -> [handler]
+    this.ws = null;
+    this.connected = false;
+    this.forcedClose = false;
+    this.connectionStatePath = options.connectionStatePath || `connection.${this.label}`;
+    this._setupState();
+    this.connect();
+  }
+
+  _setupState() {
+    stateManager.setState(this.connectionStatePath, {
+      status: 'connecting',
+      attempts: 0,
+      lastError: null
+    });
+  }
+
+  connect() {
+    this.ws = new window.WebSocket(this.url, this.protocols);
+    this.connected = false;
+    this.forcedClose = false;
+
+    this.ws.onopen = () => {
+      this.connected = true;
+      this.reconnectAttempts = 0;
+      stateManager.setState(this.connectionStatePath, { status: 'connected', attempts: this.reconnectAttempts, lastError: null });
+      this._emit('open');
+    };
+
+    this.ws.onmessage = (event) => {
+      let data;
+      try {
+        data = JSON.parse(event.data);
+      } catch {
+        data = event.data;
+      }
+      this._emit('message', data);
+      if (data && data.type) {
+        this._emit(data.type, data);
+      }
+    };
+
+    this.ws.onerror = (event) => {
+      stateManager.setState(this.connectionStatePath, { status: 'error', attempts: this.reconnectAttempts, lastError: event });
+      this._emit('error', event);
+    };
+
+    this.ws.onclose = (event) => {
+      this.connected = false;
+      stateManager.setState(this.connectionStatePath, { status: 'disconnected', attempts: this.reconnectAttempts, lastError: event });
+      this._emit('close', event);
+
+      if (!this.forcedClose) {
+        this._scheduleReconnect();
+      }
+    };
+  }
+
+  send(data) {
+    if (this.connected && this.ws.readyState === window.WebSocket.OPEN) {
+      this.ws.send(typeof data === 'string' ? data : JSON.stringify(data));
+    } else {
+      throw new Error(`[WebSocketManager] Não conectado`);
+    }
+  }
+
+  close() {
+    this.forcedClose = true;
+    if (this.ws) this.ws.close();
+  }
+
+  _scheduleReconnect() {
+    this.reconnectAttempts += 1;
+    const delay = Math.min(this.reconnectBaseDelay * (2 ** this.reconnectAttempts), this.reconnectMaxDelay);
+    stateManager.setState(this.connectionStatePath, { status: 'reconnecting', attempts: this.reconnectAttempts, lastError: null });
+    setTimeout(() => {
+      this.connect();
+    }, delay);
+  }
+
+  on(type, handler) {
+    if (!this.handlers.has(type)) this.handlers.set(type, []);
+    this.handlers.get(type).push(handler);
+  }
+
+  off(type, handler) {
+    if (!this.handlers.has(type)) return;
+    const arr = this.handlers.get(type);
+    const idx = arr.indexOf(handler);
+    if (idx >= 0) arr.splice(idx, 1);
+  }
+
+  _emit(type, ...args) {
+    if (this.handlers.has(type)) {
+      this.handlers.get(type).forEach(fn => {
+        try { fn(...args); } catch (e) { console.error(`[WebSocketManager] Handler ${type} erro:`, e); }
+      });
+    }
+  }
+}
+
+// Exemplo de uso (na inicialização do runtime):
+// const wsEspelho = new WebSocketManager('ws://localhost:8080/ws/espelho', { label: 'espelho' });
+// wsEspelho.on('message', msg => stateManager.appendState('espelho.spans', msg));
+// wsEspelho.on('open', () => stateManager.setState('espelho.status', 'ok'));
+// wsEspelho.on('close', () => stateManager.setState('espelho.status', 'down'));
\ No newline at end of file
diff --git a/flipapp_runtime_full/public/src/main.js b/flipapp_runtime_full/public/src/main.js
new file mode 100644
index 0000000000000000000000000000000000000000..82f9ee0dc45c97cc6f5eb7a14b7a5815063a7ebd
--- /dev/null
+++ b/flipapp_runtime_full/public/src/main.js
@@ -0,0 +1,365 @@
+// public/src/main.js
+// Ponto de entrada principal da aplicação FlipApp
+import { appConfig } from './core/config.js';
+import { stateManager } from './core/state.js';
+import { renderer } from './core/renderer.js';
+import { parser } from './core/parser.js';
+import { contractSystem } from './core/contracts.js';
+import { WebSocketManager } from './core/websocket-manager.js';
+import { setupComponents } from './ui/components/index.js';
+import { expressionEngine } from './core/expression-engine.js';
+import { indexedDBBridge } from './wasm/indexeddb-bridge.js';
+import { logger } from './utils/logger.js';
+
+/**
+ * Classe principal que inicializa e orquestra o FlipApp
+ */
+class FlipApp {
+  constructor() {
+    this.initialized = false;
+    this.version = appConfig.version;
+    this.environment = appConfig.environment;
+    this.websockets = new Map();
+
+    logger.info(`FlipApp ${this.version} iniciando (${this.environment})...`);
+  }
+
+  /**
+   * Inicializa a aplicação
+   */
+  async init() {
+    if (this.initialized) return;
+    
+    try {
+      // Inicializa IndexedDB
+      await indexedDBBridge.init();
+      logger.info('IndexedDB inicializado');
+
+      // Registra Service Worker para PWA/offline
+      await this.registerServiceWorker();
+
+      // Inicializa componentes de UI LogLine reutilizáveis
+      setupComponents(parser, renderer);
+      
+      // Inicializa sessão WASM
+      await this.initWasmSession();
+      
+      // Carrega estado persistente
+      stateManager.loadPersistedState();
+      
+      // Configura WebSockets
+      this.setupWebSockets();
+
+      // Configura handlers de contrato
+      this.setupContractHandlers();
+
+      // Carrega e renderiza UI LogLine
+      await this.loadAndRenderUI();
+
+      this.initialized = true;
+      logger.info('FlipApp inicializado com sucesso');
+      
+      // Dispara evento de inicialização
+      window.dispatchEvent(new CustomEvent('flipapp:ready'));
+    } catch (error) {
+      logger.error('Erro ao inicializar FlipApp:', error);
+      this.showErrorScreen(error);
+    }
+  }
+
+  /**
+   * Registra Service Worker
+   */
+  async registerServiceWorker() {
+    if ('serviceWorker' in navigator) {
+      try {
+        const registration = await navigator.serviceWorker.register('/service-worker.js');
+        logger.info('Service Worker registrado:', registration);
+        
+        // Escuta por atualizações
+        registration.addEventListener('updatefound', () => {
+          const newWorker = registration.installing;
+          newWorker.addEventListener('statechange', () => {
+            if (newWorker.state === 'installed' && navigator.serviceWorker.controller) {
+              logger.info('Nova versão disponível!');
+              this.notifyAppUpdate();
+            }
+          });
+        });
+      } catch (error) {
+        logger.error('Erro ao registrar Service Worker:', error);
+      }
+    }
+  }
+
+  /**
+   * Configura conexões WebSocket
+   */
+  setupWebSockets() {
+    try {
+      // Espelho WebSocket
+      const espelhoWs = new WebSocketManager(appConfig.espelhoWsUrl, { 
+        label: 'espelho',
+        reconnectBaseDelay: 2000,
+      });
+      
+      espelhoWs.on('message', message => {
+        if (message && typeof message === 'object') {
+          stateManager.appendState('espelho.spans', message);
+        }
+      });
+      
+      espelhoWs.on('open', () => {
+        logger.info('WebSocket Espelho conectado');
+        stateManager.setState('espelho.status', 'online');
+      });
+      
+      espelhoWs.on('close', () => {
+        stateManager.setState('espelho.status', 'offline');
+      });
+      
+      this.websockets.set('espelho', espelhoWs);
+      
+      // WhatsApp WebSocket
+      const whatsappWs = new WebSocketManager(appConfig.whatsappWsUrl, {
+        label: 'whatsapp',
+        reconnectBaseDelay: 3000,
+      });
+      
+      whatsappWs.on('message', message => {
+        if (message && typeof message === 'object') {
+          stateManager.appendState('whatsapp.messages', message);
+          
+          // Reproduz som se estiver em outra aba
+          if (stateManager.getState('activeTab') !== 'whatsapp') {
+            try {
+              new Audio('/sounds/notification.mp3').play();
+            } catch (e) {
+              // Alguns navegadores bloqueiam áudio sem interação do usuário
+            }
+          }
+        }
+      });
+      
+      whatsappWs.on('open', () => {
+        logger.info('WebSocket WhatsApp conectado');
+        stateManager.setState('whatsapp.status', 'online');
+      });
+      
+      whatsappWs.on('close', () => {
+        stateManager.setState('whatsapp.status', 'offline');
+      });
+      
+      this.websockets.set('whatsapp', whatsappWs);
+      
+    } catch (error) {
+      logger.error('Erro ao configurar WebSockets:', error);
+    }
+  }
+
+  /**
+   * Inicializa sessão WASM
+   */
+  async initWasmSession() {
+    try {
+      // Obtém ID de sessão existente ou cria um novo
+      let sessionId = localStorage.getItem('flipapp_session_id');
+      
+      if (!sessionId) {
+        sessionId = await window.wasm_init_session();
+        localStorage.setItem('flipapp_session_id', sessionId);
+        logger.info('Nova sessão WASM criada:', sessionId);
+      } else {
+        // Verifica se sessão ainda existe
+        try {
+          const healthCheck = await window.wasm_health_check();
+          const health = JSON.parse(healthCheck);
+          
+          if (health.status !== 'healthy') {
+            // Recria sessão se não estiver saudável
+            sessionId = await window.wasm_init_session();
+            localStorage.setItem('flipapp_session_id', sessionId);
+            logger.info('Sessão WASM recriada:', sessionId);
+          }
+        } catch (e) {
+          // Recria sessão em caso de erro
+          sessionId = await window.wasm_init_session();
+          localStorage.setItem('flipapp_session_id', sessionId);
+          logger.info('Sessão WASM recriada após erro:', sessionId);
+        }
+      }
+      
+      // Salva ID no estado
+      stateManager.setState('session.id', sessionId);
+      
+      // Carrega spans da sessão
+      const spansJson = await window.wasm_get_spans(sessionId);
+      const spans = JSON.parse(spansJson);
+      stateManager.setState('audit.spans', spans);
+      
+      logger.info('Sessão WASM inicializada:', sessionId);
+    } catch (error) {
+      logger.error('Erro ao inicializar sessão WASM:', error);
+      throw new Error('Falha ao inicializar módulo WASM. Verifique se o arquivo .wasm está disponível.');
+    }
+  }
+
+  /**
+   * Configura handlers para contratos
+   */
+  setupContractHandlers() {
+    stateManager.setState('contract.handlers', {
+      onContractCall: (name, params) => {
+        logger.debug(`Contrato chamado: ${name}`, params);
+        
+        contractSystem.execute(name, params)
+          .then(result => {
+            logger.debug(`Contrato executado: ${name}`, result);
+          })
+          .catch(error => {
+            logger.error(`Erro no contrato ${name}:`, error);
+            // Mostra notificação de erro se apropriado
+            if (error.name !== 'ValidationError' && error.name !== 'RateLimitError') {
+              this.showErrorNotification(error.message);
+            }
+          });
+      }
+    });
+  }
+
+  /**
+   * Carrega e renderiza a UI LogLine
+   */
+  async loadAndRenderUI() {
+    try {
+      // Define estado inicial da UI
+      stateManager.setState('activeTab', 'chat');
+      stateManager.setState('ui.loading', true);
+      stateManager.setState('chat.isTyping', false);
+      stateManager.setState('chat.history', []);
+      stateManager.setState('chat.messageInput', '');
+      
+      // Carrega arquivo LogLine principal
+      const response = await fetch('/ui/flipapp_ui.logline');
+      
+      if (!response.ok) {
+        throw new Error(`Falha ao carregar UI: ${response.status} ${response.statusText}`);
+      }
+      
+      const logLineSource = await response.text();
+      
+      // Parse da DSL LogLine
+      const ast = parser.parse(logLineSource, 'flipapp_ui.logline');
+      
+      if (ast.errors && ast.errors.length > 0) {
+        logger.error('Erros no parse da LogLine UI:', ast.errors);
+        throw new Error(`Erro de sintaxe na UI LogLine: ${ast.errors[0].message}`);
+      }
+      
+      // Renderiza árvore Virtual DOM
+      const vnode = renderer.render(ast, stateManager.getState('contract.handlers'), stateManager.getState('*'));
+      
+      // Monta no DOM
+      const rootElement = document.getElementById('root');
+      renderer.mount(vnode, rootElement, stateManager.getState('contract.handlers'));
+      
+      // Estado inicial
+      stateManager.setState('ui.loading', false);
+      stateManager.setState('ui.rendered', true);
+      
+      logger.info('UI LogLine renderizada com sucesso');
+    } catch (error) {
+      logger.error('Erro ao renderizar UI:', error);
+      throw error;
+    }
+  }
+  
+  /**
+   * Mostra tela de erro fatal
+   */
+  showErrorScreen(error) {
+    const rootElement = document.getElementById('root');
+    
+    if (!rootElement) return;
+    
+    rootElement.innerHTML = `
+      <div style="text-align:center; padding:2rem; font-family:system-ui,sans-serif;">
+        <h1 style="color:#d32f2f;">Erro ao inicializar FlipApp</h1>
+        <p style="margin:1rem 0;">${error.message || 'Ocorreu um erro inesperado.'}</p>
+        <button onclick="location.reload()" style="padding:0.5rem 1rem; background:#2196f3; color:white; border:none; border-radius:4px; cursor:pointer;">
+          Recarregar aplicação
+        </button>
+        <p style="margin-top:2rem; font-size:0.8rem; color:#666;">
+          Versão ${this.version} | ${this.environment}
+        </p>
+      </div>
+    `;
+  }
+  
+  /**
+   * Mostra notificação de atualização disponível
+   */
+  notifyAppUpdate() {
+    const notification = document.createElement('div');
+    notification.className = 'app-update-notification';
+    notification.innerHTML = `
+      <p>Nova versão disponível!</p>
+      <button id="update-button">Atualizar agora</button>
+    `;
+    
+    document.body.appendChild(notification);
+    
+    document.getElementById('update-button').addEventListener('click', () => {
+      window.location.reload();
+    });
+  }
+  
+  /**
+   * Mostra notificação de erro
+   */
+  showErrorNotification(message) {
+    const notification = document.createElement('div');
+    notification.className = 'error-notification';
+    notification.textContent = message;
+    
+    document.body.appendChild(notification);
+    
+    // Remove após 5 segundos
+    setTimeout(() => {
+      notification.classList.add('hide');
+      setTimeout(() => {
+        notification.remove();
+      }, 300);
+    }, 5000);
+  }
+  
+  /**
+   * Finaliza a aplicação
+   */
+  cleanup() {
+    // Fecha WebSockets
+    this.websockets.forEach((ws) => {
+      ws.close();
+    });
+    
+    // Fecha IndexedDB
+    indexedDBBridge.close();
+    
+    logger.info('FlipApp finalizado');
+  }
+}
+
+// Inicializa app
+window.addEventListener('DOMContentLoaded', () => {
+  window.app = new FlipApp();
+  window.app.init().catch(error => {
+    console.error('Erro fatal:', error);
+  });
+});
+
+// Cleanup na saída
+window.addEventListener('beforeunload', () => {
+  if (window.app) {
+    window.app.cleanup();
+  }
+});
\ No newline at end of file
diff --git a/flipapp_runtime_full/public/src/services/api-client.js b/flipapp_runtime_full/public/src/services/api-client.js
new file mode 100644
index 0000000000000000000000000000000000000000..73b348365300517bf94477f09877c50d2b122a2d
--- /dev/null
+++ b/flipapp_runtime_full/public/src/services/api-client.js
@@ -0,0 +1,229 @@
+// public/src/services/api-client.js
+/**
+ * Cliente HTTP para APIs com cache, retry e validação
+ */
+import { logger } from '../utils/logger.js';
+import { appConfig } from '../core/config.js';
+
+export class ApiClient {
+  constructor(options = {}) {
+    this.baseUrl = options.baseUrl || appConfig.apiBaseUrl;
+    this.defaultHeaders = options.headers || {};
+    this.timeout = options.timeout || 20000;
+    this.retries = options.retries || 1;
+    
+    this.cache = {
+      data: new Map(),
+      ttl: options.cacheTtl || 60000 // Default 1 min
+    };
+  }
+  
+  /**
+   * Faz request GET
+   * @param {string} endpoint Caminho da API
+   * @param {Object} options Opções do request
+   * @returns {Promise<Object>} Resposta da API
+   */
+  async get(endpoint, options = {}) {
+    return this.request('GET', endpoint, null, options);
+  }
+  
+  /**
+   * Faz request POST
+   * @param {string} endpoint Caminho da API
+   * @param {Object} data Dados para enviar
+   * @param {Object} options Opções do request
+   * @returns {Promise<Object>} Resposta da API
+   */
+  async post(endpoint, data, options = {}) {
+    return this.request('POST', endpoint, data, options);
+  }
+  
+  /**
+   * Faz request PUT
+   * @param {string} endpoint Caminho da API
+   * @param {Object} data Dados para enviar
+   * @param {Object} options Opções do request
+   * @returns {Promise<Object>} Resposta da API
+   */
+  async put(endpoint, data, options = {}) {
+    return this.request('PUT', endpoint, data, options);
+  }
+  
+  /**
+   * Faz request DELETE
+   * @param {string} endpoint Caminho da API
+   * @param {Object} options Opções do request
+   * @returns {Promise<Object>} Resposta da API
+   */
+  async delete(endpoint, options = {}) {
+    return this.request('DELETE', endpoint, null, options);
+  }
+  
+  /**
+   * Método genérico para requests HTTP
+   */
+  async request(method, endpoint, data = null, options = {}) {
+    const url = `${this.baseUrl}${endpoint}`;
+    const cacheKey = `${method}:${url}:${JSON.stringify(data || {})}`;
+    
+    // Tenta do cache para GET
+    if (method === 'GET' && options.useCache !== false) {
+      const cached = this.getFromCache(cacheKey);
+      if (cached) {
+        return cached;
+      }
+    }
+    
+    const requestOptions = {
+      method,
+      headers: {
+        'Content-Type': 'application/json',
+        ...this.defaultHeaders,
+        ...options.headers
+      },
+      credentials: options.credentials || 'same-origin'
+    };
+    
+    // Adiciona corpo para métodos não GET
+    if (data && method !== 'GET') {
+      requestOptions.body = JSON.stringify(data);
+    }
+    
+    // Retry logic
+    let attempts = 0;
+    let error;
+    
+    while (attempts <= this.retries) {
+      try {
+        const controller = new AbortController();
+        const timeoutId = setTimeout(() => controller.abort(), this.timeout);
+        requestOptions.signal = controller.signal;
+        
+        const response = await fetch(url, requestOptions);
+        clearTimeout(timeoutId);
+        
+        // Valida resposta HTTP
+        if (!response.ok) {
+          let errorData;
+          try {
+            errorData = await response.json();
+          } catch (e) {
+            errorData = { message: response.statusText };
+          }
+          
+          throw new ApiError(
+            errorData.message || `HTTP error ${response.status}`,
+            response.status,
+            errorData
+          );
+        }
+        
+        // Parse JSON response
+        let responseData;
+        const contentType = response.headers.get('content-type');
+        
+        if (contentType && contentType.includes('application/json')) {
+          responseData = await response.json();
+        } else {
+          responseData = await response.text();
+        }
+        
+        // Adiciona ao cache para GET
+        if (method === 'GET' && options.useCache !== false) {
+          this.addToCache(cacheKey, responseData);
+        }
+        
+        return responseData;
+        
+      } catch (e) {
+        attempts++;
+        error = e;
+        
+        if (e.name === 'AbortError') {
+          logger.warn(`[API] Timeout em request para ${endpoint}`);
+        } else {
+          logger.warn(`[API] Erro em request para ${endpoint} (tentativa ${attempts}/${this.retries + 1}):`, e);
+        }
+        
+        // Não faz retry para certos casos
+        if (e instanceof ApiError && [400, 401, 403, 404].includes(e.status)) {
+          break;
+        }
+        
+        // Espera antes de retry
+        if (attempts <= this.retries) {
+          const delay = Math.min(1000 * (2 ** (attempts - 1)), 10000);
+          await new Promise(resolve => setTimeout(resolve, delay));
+        }
+      }
+    }
+    
+    // Se chegou aqui, todas as tentativas falharam
+    throw error || new Error(`Falha em request para ${endpoint}`);
+  }
+  
+  /**
+   * Recupera do cache
+   */
+  getFromCache(key) {
+    const cacheItem = this.cache.data.get(key);
+    if (!cacheItem) return null;
+    
+    const now = Date.now();
+    if (now - cacheItem.timestamp < this.cache.ttl) {
+      return cacheItem.data;
+    }
+    
+    // Expirou
+    this.cache.data.delete(key);
+    return null;
+  }
+  
+  /**
+   * Adiciona ao cache
+   */
+  addToCache(key, data) {
+    this.cache.data.set(key, {
+      data,
+      timestamp: Date.now()
+    });
+    
+    // Limita tamanho do cache
+    if (this.cache.data.size > 100) {
+      const oldestKey = [...this.cache.data.entries()]
+        .reduce((oldest, [key, item]) => {
+          if (!oldest || item.timestamp < oldest.timestamp) {
+            return { key, timestamp: item.timestamp };
+          }
+          return oldest;
+        }, null).key;
+      
+      if (oldestKey) {
+        this.cache.data.delete(oldestKey);
+      }
+    }
+  }
+  
+  /**
+   * Limpa cache
+   */
+  clearCache() {
+    this.cache.data.clear();
+  }
+}
+
+/**
+ * Erro customizado para API
+ */
+export class ApiError extends Error {
+  constructor(message, status, data = {}) {
+    super(message);
+    this.name = 'ApiError';
+    this.status = status;
+    this.data = data;
+  }
+}
+
+// Singleton
+export const apiClient = new ApiClient();
\ No newline at end of file
diff --git a/flipapp_runtime_full/public/src/services/llm-client.js b/flipapp_runtime_full/public/src/services/llm-client.js
new file mode 100644
index 0000000000000000000000000000000000000000..d1f08f80a6ea39e0b2f7c1540e0da9e7aab48fe8
--- /dev/null
+++ b/flipapp_runtime_full/public/src/services/llm-client.js
@@ -0,0 +1,138 @@
+// public/src/services/llm-client.js
+/**
+ * Cliente para API de LLM real (para substituir simulação WASM)
+ */
+import { logger } from '../utils/logger.js';
+import { appConfig } from '../core/config.js';
+
+export class LlmClient {
+  constructor(options = {}) {
+    this.baseUrl = options.baseUrl || appConfig.llmApiUrl || 'http://localhost:3000/api/llm';
+    this.timeout = options.timeout || 30000; // 30 segundos
+    this.retries = options.retries || 2;
+    this.model = options.model || 'gpt-4-turbo';
+    this.cache = new Map();
+  }
+
+  /**
+   * Processa um prompt via API LLM
+   * @param {string} prompt Texto do prompt
+   * @param {Object} options Opções adicionais
+   * @returns {Promise<string>} Resposta do LLM
+   */
+  async processPrompt(prompt, options = {}) {
+    const sessionId = options.sessionId || 'default';
+    const reqOptions = {
+      model: options.model || this.model,
+      max_tokens: options.maxTokens || 1024,
+      temperature: options.temperature || 0.7
+    };
+    
+    // Simulação de modo offline para dev
+    if (options.offlineModeSimulation) {
+      logger.info(`[LLM] Modo offline: simulando resposta para "${prompt.substring(0, 30)}..."`);
+      await new Promise(resolve => setTimeout(resolve, 1000));
+      return `[OFFLINE MODE] Resposta simulada para: "${prompt.substring(0, 30)}..."`;
+    }
+    
+    // Checa cache, se aplicável
+    if (options.useCache !== false) {
+      const cacheKey = `${prompt}|${JSON.stringify(reqOptions)}`;
+      if (this.cache.has(cacheKey)) {
+        logger.info(`[LLM] Resposta recuperada do cache para "${prompt.substring(0, 30)}..."`);
+        return this.cache.get(cacheKey);
+      }
+    }
+
+    // Prepara payload
+    const payload = {
+      prompt,
+      options: reqOptions,
+      session_id: sessionId,
+      context: options.context || []
+    };
+    
+    // Faz request
+    let attempts = 0;
+    let error;
+    
+    while (attempts <= this.retries) {
+      try {
+        const controller = new AbortController();
+        const timeoutId = setTimeout(() => controller.abort(), this.timeout);
+        
+        const response = await fetch(`${this.baseUrl}/process`, {
+          method: 'POST',
+          headers: {
+            'Content-Type': 'application/json',
+            ...(options.apiKey ? { 'Authorization': `Bearer ${options.apiKey}` } : {})
+          },
+          body: JSON.stringify(payload),
+          signal: controller.signal
+        });
+        
+        clearTimeout(timeoutId);
+        
+        if (!response.ok) {
+          throw new Error(`HTTP error: ${response.status} ${response.statusText}`);
+        }
+        
+        const data = await response.json();
+        
+        if (!data.content) {
+          throw new Error('Resposta do LLM não contém conteúdo');
+        }
+        
+        // Adiciona ao cache se aplicável
+        if (options.useCache !== false) {
+          const cacheKey = `${prompt}|${JSON.stringify(reqOptions)}`;
+          this.cache.set(cacheKey, data.content);
+          
+          // Limita tamanho do cache
+          if (this.cache.size > 100) {
+            // Remove entrada mais antiga
+            const firstKey = this.cache.keys().next().value;
+            this.cache.delete(firstKey);
+          }
+        }
+        
+        // Rastreamento de uso/custos
+        if (data.usage) {
+          logger.debug('[LLM] Uso de tokens:', data.usage);
+        }
+        
+        return data.content;
+        
+      } catch (e) {
+        attempts++;
+        error = e;
+        
+        if (e.name === 'AbortError') {
+          logger.warn('[LLM] Timeout ao processar prompt');
+        } else {
+          logger.warn(`[LLM] Erro ao processar prompt (tentativa ${attempts}/${this.retries + 1}):`, e);
+        }
+        
+        // Espera antes de retry com backoff exponencial
+        if (attempts <= this.retries) {
+          const delay = Math.min(1000 * (2 ** (attempts - 1)), 10000);
+          await new Promise(resolve => setTimeout(resolve, delay));
+        }
+      }
+    }
+    
+    // Se chegou aqui, todas as tentativas falharam
+    throw error || new Error('Falha ao processar prompt');
+  }
+  
+  /**
+   * Cancela operações pendentes
+   */
+  cancelPending() {
+    // Implementação depende do mecanismo usado para controle de requests
+    logger.info('[LLM] Cancelando operações pendentes');
+  }
+}
+
+// Singleton
+export const llmClient = new LlmClient();
\ No newline at end of file
diff --git a/flipapp_runtime_full/public/src/ui/components/index.js b/flipapp_runtime_full/public/src/ui/components/index.js
new file mode 100644
index 0000000000000000000000000000000000000000..0f7a1e4d57de13a60957842618a8f38f20e8f880
--- /dev/null
+++ b/flipapp_runtime_full/public/src/ui/components/index.js
@@ -0,0 +1,147 @@
+// public/src/ui/components/index.js
+/**
+ * Registro de componentes reutilizáveis da UI LogLine
+ */
+import { logger } from '../../utils/logger.js';
+
+// Componente Button
+const buttonComponent = `
+- type: button
+  class: "{{props.className || 'btn'}}"
+  style: "{{props.style}}"
+  disabled: "{{props.disabled}}"
+  content: "{{props.label || props.content || ''}}"
+  on: "click: {{props.onClick || 'noop'}}"
+`;
+
+// Componente Card
+const cardComponent = `
+- type: container
+  class: "card {{props.className || ''}}"
+  style: "border-radius:8px; box-shadow:0 2px 8px rgba(0,0,0,0.1); padding:1rem; margin:1rem 0; {{props.style}}"
+  children:
+    - type: when
+      when: "props.title"
+      children:
+        - type: container
+          class: "card-header"
+          style: "margin-bottom:0.5rem; font-weight:bold"
+          content: "{{props.title}}"
+    - type: container
+      class: "card-content"
+      children:
+        - type: text
+          content: "{{props.content}}"
+          style: "margin:0"
+`;
+
+// Componente Input
+const inputComponent = `
+- type: container
+  class: "input-wrapper {{props.className || ''}}"
+  style: "{{props.style}}"
+  children:
+    - type: when
+      when: "props.label"
+      children:
+        - type: text
+          content: "{{props.label}}"
+          style: "display:block; margin-bottom:0.25rem; font-size:0.9rem"
+    - type: input
+      placeholder: "{{props.placeholder || ''}}"
+      bind: "{{props.bind || ''}}"
+      type: "{{props.type || 'text'}}"
+      style: "width:100%; padding:0.5rem; border:1px solid #ccc; border-radius:4px"
+      disabled: "{{props.disabled}}"
+`;
+
+// Componente Badge
+const badgeComponent = `
+- type: container
+  class: "badge {{props.className || ''}}"
+  style: "display:inline-block; padding:0.25rem 0.5rem; border-radius:16px; font-size:0.75rem; background:{{props.color || '#e0e0e0'}}; color:{{props.textColor || '#333'}}; {{props.style}}"
+  children:
+    - type: text
+      content: "{{props.content || props.text || ''}}"
+      style: "margin:0"
+`;
+
+// Componente Avatar
+const avatarComponent = `
+- type: container
+  class: "avatar {{props.className || ''}}"
+  style: "width:{{props.size || '40px'}}; height:{{props.size || '40px'}}; border-radius:50%; overflow:hidden; {{props.style}}"
+  children:
+    - type: image
+      src: "{{props.src || ''}}"
+      alt: "{{props.alt || 'Avatar'}}"
+      style: "width:100%; height:100%; object-fit:cover"
+`;
+
+// Componente Loader
+const loaderComponent = `
+- type: container
+  class: "loader {{props.className || ''}}"
+  style: "display:inline-block; width:{{props.size || '24px'}}; height:{{props.size || '24px'}}; border:3px solid rgba(0,0,0,0.1); border-radius:50%; border-top-color:{{props.color || '#3498db'}}; animation:spin 1s linear infinite; {{props.style}}"
+`;
+
+// Componente Modal
+const modalComponent = `
+- type: container
+  class: "modal {{props.className || ''}}"
+  style: "position:fixed; top:0; left:0; width:100%; height:100%; display:{{props.visible ? 'flex' : 'none'}}; align-items:center; justify-content:center; background:rgba(0,0,0,0.5); z-index:1000; {{props.style}}"
+  children:
+    - type: container
+      class: "modal-content"
+      style: "background:white; border-radius:8px; padding:1rem; max-width:90%; width:{{props.width || '400px'}}; max-height:90vh; overflow-y:auto"
+      children:
+        - type: container
+          class: "modal-header"
+          style: "display:flex; justify-content:space-between; align-items:center; margin-bottom:1rem"
+          children:
+            - type: text
+              content: "{{props.title || 'Modal'}}"
+              style: "font-weight:bold; margin:0"
+            - type: button
+              content: "×"
+              style: "background:none; border:none; font-size:1.5rem; cursor:pointer"
+              on: "click: {{props.onClose || ''}}"
+        - type: container
+          class: "modal-body"
+          children:
+            - type: text
+              content: "{{props.content || ''}}"
+`;
+
+// Lista de todos os componentes a registrar
+const components = {
+  Button: buttonComponent,
+  Card: cardComponent,
+  Input: inputComponent,
+  Badge: badgeComponent,
+  Avatar: avatarComponent,
+  Loader: loaderComponent,
+  Modal: modalComponent
+};
+
+/**
+ * Registra todos os componentes no parser e renderer
+ */
+export function setupComponents(parser, renderer) {
+  Object.entries(components).forEach(([name, template]) => {
+    try {
+      const ast = parser.parse(template, `component_${name.toLowerCase()}.logline`);
+      renderer.registerComponent(name, ast);
+      logger.debug(`Componente '${name}' registrado com sucesso`);
+    } catch (error) {
+      logger.error(`Erro ao registrar componente '${name}':`, error);
+    }
+  });
+  
+  logger.info(`${Object.keys(components).length} componentes LogLine registrados`);
+}
+
+/**
+ * Exporta componentes individuais para uso direto
+ */
+export const componentTemplates = components;
\ No newline at end of file
diff --git a/flipapp_runtime_full/public/src/utils/logger.js b/flipapp_runtime_full/public/src/utils/logger.js
new file mode 100644
index 0000000000000000000000000000000000000000..2b00f5a920f8628dee28975e03b6c5e8fb82d6b6
--- /dev/null
+++ b/flipapp_runtime_full/public/src/utils/logger.js
@@ -0,0 +1,151 @@
+// public/src/utils/logger.js
+/**
+ * Sistema de logging centralizado com níveis e formatação
+ */
+export class Logger {
+  constructor(options = {}) {
+    this.options = {
+      level: options.level || (process.env.NODE_ENV === 'production' ? 'warn' : 'debug'),
+      prefix: options.prefix || 'FlipApp',
+      enableTimestamps: options.enableTimestamps !== undefined ? options.enableTimestamps : true,
+      enableStackTrace: options.enableStackTrace !== undefined ? options.enableStackTrace : false,
+      persist: options.persist !== undefined ? options.persist : false,
+      consoleOutput: options.consoleOutput !== undefined ? options.consoleOutput : true
+    };
+    
+    this.levels = {
+      debug: 0,
+      info: 1,
+      warn: 2,
+      error: 3
+    };
+    
+    this.persistedLogs = [];
+    this.maxPersistedLogs = options.maxPersistedLogs || 1000;
+  }
+  
+  /**
+   * Log no nível debug
+   * @param {...any} args Argumentos para log
+   */
+  debug(...args) {
+    this._log('debug', ...args);
+  }
+  
+  /**
+   * Log no nível info
+   * @param {...any} args Argumentos para log
+   */
+  info(...args) {
+    this._log('info', ...args);
+  }
+  
+  /**
+   * Log no nível warn
+   * @param {...any} args Argumentos para log
+   */
+  warn(...args) {
+    this._log('warn', ...args);
+  }
+  
+  /**
+   * Log no nível error
+   * @param {...any} args Argumentos para log
+   */
+  error(...args) {
+    this._log('error', ...args);
+  }
+  
+  /**
+   * Método interno para processamento do log
+   * @param {string} level Nível do log
+   * @param {...any} args Argumentos para log
+   */
+  _log(level, ...args) {
+    if (this.levels[level] < this.levels[this.options.level]) return;
+    
+    const timestamp = this.options.enableTimestamps ? new Date().toISOString() : '';
+    const prefix = this.options.prefix ? `[${this.options.prefix}]` : '';
+    const logLevel = `[${level.toUpperCase()}]`;
+    
+    // Prep mensagem para console
+    const consoleArgs = [];
+    if (timestamp) consoleArgs.push(`${timestamp}`);
+    consoleArgs.push(`${prefix}${logLevel}`);
+    consoleArgs.push(...args);
+    
+    // Log para console
+    if (this.options.consoleOutput) {
+      switch (level) {
+        case 'debug':
+          console.debug(...consoleArgs);
+          break;
+        case 'info':
+          console.info(...consoleArgs);
+          break;
+        case 'warn':
+          console.warn(...consoleArgs);
+          break;
+        case 'error':
+          console.error(...consoleArgs);
+          break;
+      }
+    }
+    
+    // Persistir log se configurado
+    if (this.options.persist) {
+      const logEntry = {
+        timestamp: new Date().toISOString(),
+        level,
+        message: args.map(arg => 
+          typeof arg === 'object' ? JSON.stringify(arg) : String(arg)
+        ).join(' ')
+      };
+      
+      this.persistedLogs.push(logEntry);
+      
+      // Limita tamanho do buffer de logs
+      if (this.persistedLogs.length > this.maxPersistedLogs) {
+        this.persistedLogs.shift();
+      }
+    }
+    
+    // Adiciona stack trace para erros se configurado
+    if (level === 'error' && this.options.enableStackTrace) {
+      console.groupCollapsed('Stack trace:');
+      console.trace();
+      console.groupEnd();
+    }
+  }
+  
+  /**
+   * Exporta logs persistidos
+   * @returns {Array} Array de logs persistidos
+   */
+  export() {
+    return this.persistedLogs;
+  }
+  
+  /**
+   * Limpa logs persistidos
+   */
+  clear() {
+    this.persistedLogs = [];
+  }
+  
+  /**
+   * Define nível de log
+   * @param {string} level Nível de log ('debug', 'info', 'warn', 'error')
+   */
+  setLevel(level) {
+    if (this.levels[level] !== undefined) {
+      this.options.level = level;
+    }
+  }
+}
+
+// Logger singleton
+export const logger = new Logger({
+  level: process.env.NODE_ENV === 'production' ? 'warn' : 'debug',
+  persist: process.env.NODE_ENV === 'development'
+});
\ No newline at end of file
diff --git a/flipapp_runtime_full/public/src/wasm/indexeddb-bridge.js b/flipapp_runtime_full/public/src/wasm/indexeddb-bridge.js
new file mode 100644
index 0000000000000000000000000000000000000000..ea866f59a5ccb0b31347d02466e56f3fd24235f7
--- /dev/null
+++ b/flipapp_runtime_full/public/src/wasm/indexeddb-bridge.js
@@ -0,0 +1,450 @@
+// public/src/wasm/indexeddb-bridge.js
+/**
+ * Bridge entre Rust/WASM e IndexedDB para persistência
+ */
+export class IndexedDBBridge {
+  constructor(dbName = 'FlipAppDB', version = 1) {
+    this.dbName = dbName;
+    this.version = version;
+    this.db = null;
+    this.isReady = false;
+    this.initPromise = null;
+  }
+
+  /**
+   * Inicializa conexão com IndexedDB
+   * @returns {Promise<IDBDatabase>} Database instance
+   */
+  async init() {
+    if (this.initPromise) {
+      return this.initPromise;
+    }
+
+    this.initPromise = new Promise((resolve, reject) => {
+      const request = indexedDB.open(this.dbName, this.version);
+
+      request.onerror = () => {
+        console.error('[IndexedDB] Erro ao abrir database:', request.error);
+        reject(request.error);
+      };
+
+      request.onsuccess = () => {
+        this.db = request.result;
+        this.isReady = true;
+        console.log(`[IndexedDB] Database ${this.dbName} aberta com sucesso`);
+        resolve(this.db);
+      };
+
+      request.onupgradeneeded = (event) => {
+        const db = event.target.result;
+        this.createStores(db, event.oldVersion);
+      };
+    });
+
+    return this.initPromise;
+  }
+
+  /**
+   * Cria object stores necessárias
+   * @param {IDBDatabase} db - Database instance
+   * @param {number} oldVersion - Versão anterior
+   */
+  createStores(db, oldVersion) {
+    console.log(`[IndexedDB] Upgrade de ${oldVersion} para ${this.version}`);
+
+    // Store para spans/eventos
+    if (!db.objectStoreNames.contains('spans')) {
+      const spanStore = db.createObjectStore('spans', { keyPath: 'id' });
+      spanStore.createIndex('session_id', 'session_id', { unique: false });
+      spanStore.createIndex('timestamp', 'timestamp', { unique: false });
+      spanStore.createIndex('type', 'type', { unique: false });
+    }
+
+    // Store para sessões
+    if (!db.objectStoreNames.contains('sessions')) {
+      const sessionStore = db.createObjectStore('sessions', { keyPath: 'id' });
+      sessionStore.createIndex('created_at', 'created_at', { unique: false });
+    }
+
+    // Store para contratos
+    if (!db.objectStoreNames.contains('contracts')) {
+      const contractStore = db.createObjectStore('contracts', { keyPath: 'id' });
+      contractStore.createIndex('session_id', 'session_id', { unique: false });
+      contractStore.createIndex('status', 'status', { unique: false });
+    }
+
+    // Store para cache de chat
+    if (!db.objectStoreNames.contains('chat_history')) {
+      const chatStore = db.createObjectStore('chat_history', { keyPath: 'id' });
+      chatStore.createIndex('session_id', 'session_id', { unique: false });
+      chatStore.createIndex('timestamp', 'timestamp', { unique: false });
+    }
+  }
+
+  /**
+   * Adiciona span ao IndexedDB
+   * @param {Object} span - Objeto span
+   * @returns {Promise<string>} ID do span adicionado
+   */
+  async addSpan(span) {
+    await this.ensureReady();
+    
+    const transaction = this.db.transaction(['spans'], 'readwrite');
+    const store = transaction.objectStore('spans');
+    
+    // Adiciona timestamp se não existir
+    if (!span.timestamp) {
+      span.timestamp = new Date().toISOString();
+    }
+
+    return new Promise((resolve, reject) => {
+      const request = store.add(span);
+      
+      request.onsuccess = () => {
+        console.debug(`[IndexedDB] Span adicionado: ${span.id}`);
+        resolve(span.id);
+      };
+      
+      request.onerror = () => {
+        console.error('[IndexedDB] Erro ao adicionar span:', request.error);
+        reject(request.error);
+      };
+    });
+  }
+
+  /**
+   * Recupera spans por sessão
+   * @param {string} sessionId - ID da sessão
+   * @param {Object} options - Opções de filtro
+   * @returns {Promise<Array>} Array de spans
+   */
+  async getSpans(sessionId, options = {}) {
+    await this.ensureReady();
+    
+    const transaction = this.db.transaction(['spans'], 'readonly');
+    const store = transaction.objectStore('spans');
+    const index = store.index('session_id');
+    
+    return new Promise((resolve, reject) => {
+      const request = index.getAll(sessionId);
+      
+      request.onsuccess = () => {
+        let spans = request.result;
+        
+        // Aplica filtros
+        if (options.type) {
+          spans = spans.filter(span => span.type === options.type);
+        }
+        
+        if (options.since) {
+          const sinceDate = new Date(options.since);
+          spans = spans.filter(span => new Date(span.timestamp) >= sinceDate);
+        }
+        
+        if (options.limit) {
+          spans = spans.slice(-options.limit); // Últimos N spans
+        }
+        
+        // Ordena por timestamp
+        spans.sort((a, b) => new Date(a.timestamp) - new Date(b.timestamp));
+        
+        resolve(spans);
+      };
+      
+      request.onerror = () => {
+        console.error('[IndexedDB] Erro ao recuperar spans:', request.error);
+        reject(request.error);
+      };
+    });
+  }
+
+  /**
+   * Salva sessão
+   * @param {Object} session - Dados da sessão
+   * @returns {Promise<string>} ID da sessão
+   */
+  async saveSession(session) {
+    await this.ensureReady();
+    
+    const transaction = this.db.transaction(['sessions'], 'readwrite');
+    const store = transaction.objectStore('sessions');
+    
+    if (!session.created_at) {
+      session.created_at = new Date().toISOString();
+    }
+    session.updated_at = new Date().toISOString();
+
+    return new Promise((resolve, reject) => {
+      const request = store.put(session);
+      
+      request.onsuccess = () => {
+        console.debug(`[IndexedDB] Sessão salva: ${session.id}`);
+        resolve(session.id);
+      };
+      
+      request.onerror = () => {
+        console.error('[IndexedDB] Erro ao salvar sessão:', request.error);
+        reject(request.error);
+      };
+    });
+  }
+
+  /**
+   * Recupera sessão por ID
+   * @param {string} sessionId - ID da sessão
+   * @returns {Promise<Object|null>} Dados da sessão
+   */
+  async getSession(sessionId) {
+    await this.ensureReady();
+    
+    const transaction = this.db.transaction(['sessions'], 'readonly');
+    const store = transaction.objectStore('sessions');
+    
+    return new Promise((resolve, reject) => {
+      const request = store.get(sessionId);
+      
+      request.onsuccess = () => {
+        resolve(request.result || null);
+      };
+      
+      request.onerror = () => {
+        console.error('[IndexedDB] Erro ao recuperar sessão:', request.error);
+        reject(request.error);
+      };
+    });
+  }
+
+  /**
+   * Salva contrato
+   * @param {Object} contract - Dados do contrato
+   * @returns {Promise<string>} ID do contrato
+   */
+  async saveContract(contract) {
+    await this.ensureReady();
+    
+    const transaction = this.db.transaction(['contracts'], 'readwrite');
+    const store = transaction.objectStore('contracts');
+    
+    if (!contract.created_at) {
+      contract.created_at = new Date().toISOString();
+    }
+
+    return new Promise((resolve, reject) => {
+      const request = store.add(contract);
+      
+      request.onsuccess = () => {
+        console.debug(`[IndexedDB] Contrato salvo: ${contract.id}`);
+        resolve(contract.id);
+      };
+      
+      request.onerror = () => {
+        console.error('[IndexedDB] Erro ao salvar contrato:', request.error);
+        reject(request.error);
+      };
+    });
+  }
+
+  /**
+   * Salva histórico de chat
+   * @param {string} sessionId - ID da sessão
+   * @param {Array} messages - Array de mensagens
+   * @returns {Promise<void>}
+   */
+  async saveChatHistory(sessionId, messages) {
+    await this.ensureReady();
+    
+    const transaction = this.db.transaction(['chat_history'], 'readwrite');
+    const store = transaction.objectStore('chat_history');
+    
+    // Remove histórico anterior desta sessão
+    const index = store.index('session_id');
+    const deleteRequest = index.openCursor(sessionId);
+    
+    return new Promise((resolve, reject) => {
+      deleteRequest.onsuccess = (event) => {
+        const cursor = event.target.result;
+        if (cursor) {
+          cursor.delete();
+          cursor.continue();
+        } else {
+          // Adiciona novo histórico
+          messages.forEach((message, index) => {
+            store.add({
+              id: `${sessionId}_${message.id || index}`,
+              session_id: sessionId,
+              ...message,
+              timestamp: message.timestamp || new Date().toISOString()
+            });
+          });
+          
+          resolve();
+        }
+      };
+      
+      deleteRequest.onerror = () => {
+        console.error('[IndexedDB] Erro ao salvar chat:', deleteRequest.error);
+        reject(deleteRequest.error);
+      };
+    });
+  }
+
+  /**
+   * Recupera histórico de chat
+   * @param {string} sessionId - ID da sessão
+   * @returns {Promise<Array>} Array de mensagens
+   */
+  async getChatHistory(sessionId) {
+    await this.ensureReady();
+    
+    const transaction = this.db.transaction(['chat_history'], 'readonly');
+    const store = transaction.objectStore('chat_history');
+    const index = store.index('session_id');
+    
+    return new Promise((resolve, reject) => {
+      const request = index.getAll(sessionId);
+      
+      request.onsuccess = () => {
+        const messages = request.result;
+        messages.sort((a, b) => new Date(a.timestamp) - new Date(b.timestamp));
+        resolve(messages);
+      };
+      
+      request.onerror = () => {
+        console.error('[IndexedDB] Erro ao recuperar chat:', request.error);
+        reject(request.error);
+      };
+    });
+  }
+
+  /**
+   * Limpa dados antigos (manutenção)
+   * @param {number} daysOld - Dias de idade para limpeza
+   * @returns {Promise<number>} Número de registros removidos
+   */
+  async cleanup(daysOld = 30) {
+    await this.ensureReady();
+    
+    const cutoffDate = new Date();
+    cutoffDate.setDate(cutoffDate.getDate() - daysOld);
+    const cutoffISO = cutoffDate.toISOString();
+    
+    let totalDeleted = 0;
+    const stores = ['spans', 'sessions', 'contracts', 'chat_history'];
+    
+    for (const storeName of stores) {
+      const transaction = this.db.transaction([storeName], 'readwrite');
+      const store = transaction.objectStore('stores');
+      const index = store.index('timestamp') || store.index('created_at');
+      
+      if (index) {
+        const range = IDBKeyRange.upperBound(cutoffISO);
+        const request = index.openCursor(range);
+        
+        await new Promise((resolve) => {
+          request.onsuccess = (event) => {
+            const cursor = event.target.result;
+            if (cursor) {
+              cursor.delete();
+              totalDeleted++;
+              cursor.continue();
+            } else {
+              resolve();
+            }
+          };
+        });
+      }
+    }
+    
+    console.log(`[IndexedDB] Limpeza concluída: ${totalDeleted} registros removidos`);
+    return totalDeleted;
+  }
+
+  /**
+   * Garante que o DB está pronto
+   * @returns {Promise<void>}
+   */
+  async ensureReady() {
+    if (!this.isReady) {
+      await this.init();
+    }
+  }
+
+  /**
+   * Fecha conexão com o banco
+   */
+  close() {
+    if (this.db) {
+      this.db.close();
+      this.db = null;
+      this.isReady = false;
+      this.initPromise = null;
+    }
+  }
+
+  /**
+   * Remove database completamente
+   * @returns {Promise<void>}
+   */
+  async deleteDatabase() {
+    this.close();
+    
+    return new Promise((resolve, reject) => {
+      const deleteRequest = indexedDB.deleteDatabase(this.dbName);
+      
+      deleteRequest.onsuccess = () => {
+        console.log(`[IndexedDB] Database ${this.dbName} removida`);
+        resolve();
+      };
+      
+      deleteRequest.onerror = () => {
+        console.error('[IndexedDB] Erro ao remover database:', deleteRequest.error);
+        reject(deleteRequest.error);
+      };
+    });
+  }
+}
+
+// Singleton global
+export const indexedDBBridge = new IndexedDBBridge();
+
+// Funções utilitárias para usar no WASM
+window.idb_add_span = async (spanJson) => {
+  try {
+    const span = JSON.parse(spanJson);
+    return await indexedDBBridge.addSpan(span);
+  } catch (error) {
+    console.error('[IndexedDB Bridge] Erro em idb_add_span:', error);
+    throw error;
+  }
+};
+
+window.idb_get_spans = async (sessionId, optionsJson = '{}') => {
+  try {
+    const options = JSON.parse(optionsJson);
+    const spans = await indexedDBBridge.getSpans(sessionId, options);
+    return JSON.stringify(spans);
+  } catch (error) {
+    console.error('[IndexedDB Bridge] Erro em idb_get_spans:', error);
+    throw error;
+  }
+};
+
+window.idb_save_session = async (sessionJson) => {
+  try {
+    const session = JSON.parse(sessionJson);
+    return await indexedDBBridge.saveSession(session);
+  } catch (error) {
+    console.error('[IndexedDB Bridge] Erro em idb_save_session:', error);
+    throw error;
+  }
+};
+
+window.idb_save_contract = async (contractJson) => {
+  try {
+    const contract = JSON.parse(contractJson);
+    return await indexedDBBridge.saveContract(contract);
+  } catch (error) {
+    console.error('[IndexedDB Bridge] Erro em idb_save_contract:', error);
+    throw error;
+  }
+};
\ No newline at end of file
diff --git a/flipapp_runtime_full/public/styles/main.css b/flipapp_runtime_full/public/styles/main.css
new file mode 100644
index 0000000000000000000000000000000000000000..9bed7ce1d00b51fce7165c11d53107137b60115a
--- /dev/null
+++ b/flipapp_runtime_full/public/styles/main.css
@@ -0,0 +1,296 @@
+/* public/styles/main.css (continuação) */
+  display: inline-flex;
+  align-items: center;
+  justify-content: center;
+  padding: 0.5rem 1rem;
+  border-radius: var(--border-radius);
+  font-weight: 500;
+  text-align: center;
+  cursor: pointer;
+  transition: background-color var(--transition-fast), transform var(--transition-fast);
+  border: none;
+  background-color: #e0e0e0;
+  color: var(--text-primary);
+}
+
+.btn:hover {
+  background-color: #d0d0d0;
+}
+
+.btn:active {
+  transform: translateY(1px);
+}
+
+.btn-primary {
+  background-color: var(--primary);
+  color: white;
+}
+
+.btn-primary:hover {
+  background-color: var(--primary-dark);
+}
+
+.btn-secondary {
+  background-color: var(--secondary);
+  color: white;
+}
+
+.btn-danger {
+  background-color: var(--error);
+  color: white;
+}
+
+.btn-disabled, .btn[disabled] {
+  opacity: 0.6;
+  cursor: not-allowed;
+  pointer-events: none;
+}
+
+.btn-loading {
+  position: relative;
+  color: transparent;
+}
+
+/* Cartões */
+.card {
+  background-color: var(--surface);
+  border-radius: var(--border-radius-lg);
+  box-shadow: var(--shadow-sm);
+  overflow: hidden;
+  transition: box-shadow var(--transition-medium);
+}
+
+.card:hover {
+  box-shadow: var(--shadow-md);
+}
+
+/* Formulários */
+input, textarea, select {
+  width: 100%;
+  padding: 0.75rem;
+  border: 1px solid var(--border);
+  border-radius: var(--border-radius);
+  background-color: var(--surface);
+  transition: border-color var(--transition-fast), box-shadow var(--transition-fast);
+}
+
+input:focus, textarea:focus, select:focus {
+  outline: none;
+  border-color: var(--primary);
+  box-shadow: 0 0 0 2px var(--primary-light);
+}
+
+label {
+  display: block;
+  margin-bottom: var(--spacing-xs);
+  color: var(--text-secondary);
+}
+
+/* Tabs */
+.tabs {
+  display: flex;
+  border-bottom: 1px solid var(--border);
+}
+
+.tab-item {
+  padding: var(--spacing-md) var(--spacing-lg);
+  cursor: pointer;
+  transition: all var(--transition-fast);
+}
+
+.tab-item:hover {
+  background-color: rgba(0, 0, 0, 0.05);
+}
+
+.tab-item.active {
+  border-bottom: 2px solid var(--primary);
+  color: var(--primary);
+  font-weight: 500;
+}
+
+/* Badges */
+.badge {
+  display: inline-flex;
+  align-items: center;
+  padding: 0.25rem 0.5rem;
+  border-radius: 16px;
+  font-size: var(--font-size-sm);
+  font-weight: 500;
+}
+
+/* Loaders */
+.loading-container {
+  display: flex;
+  flex-direction: column;
+  align-items: center;
+  justify-content: center;
+  height: 100vh;
+}
+
+.loading-spinner {
+  width: 40px;
+  height: 40px;
+  border: 3px solid rgba(0, 0, 0, 0.1);
+  border-radius: 50%;
+  border-top-color: var(--primary);
+  animation: spin 1s linear infinite;
+  margin-bottom: var(--spacing-md);
+}
+
+@keyframes spin {
+  to { transform: rotate(360deg); }
+}
+
+/* Typing indicator */
+.typing-indicator {
+  display: inline-flex;
+  align-items: center;
+}
+
+.typing-indicator::after {
+  content: '...';
+  width: 1.5em;
+  overflow: hidden;
+  animation: ellipsis 1.5s infinite;
+}
+
+@keyframes ellipsis {
+  0% { width: 0; }
+  33% { width: 0.5em; }
+  66% { width: 1em; }
+  100% { width: 1.5em; }
+}
+
+/* Chat styles */
+.message-bubble {
+  position: relative;
+  max-width: 80%;
+}
+
+.message-user .message-bubble {
+  background-color: var(--primary-light);
+  margin-left: auto;
+}
+
+.message-bot .message-bubble {
+  background-color: var(--surface);
+}
+
+.chat-history-container {
+  max-height: 400px;
+  overflow-y: auto;
+  padding: var(--spacing-md);
+  background-color: #f9f9f9;
+}
+
+/* Notificações */
+.error-notification, .app-update-notification {
+  position: fixed;
+  bottom: 20px;
+  right: 20px;
+  background-color: white;
+  padding: var(--spacing-md);
+  border-radius: var(--border-radius);
+  box-shadow: var(--shadow-md);
+  z-index: 1000;
+  animation: slide-up 0.3s ease;
+}
+
+.error-notification {
+  background-color: var(--error);
+  color: white;
+}
+
+.app-update-notification {
+  display: flex;
+  align-items: center;
+  justify-content: space-between;
+  gap: var(--spacing-md);
+}
+
+@keyframes slide-up {
+  from { transform: translateY(100%); opacity: 0; }
+  to { transform: translateY(0); opacity: 1; }
+}
+
+.offline-banner {
+  background-color: var(--warning);
+  color: white;
+  text-align: center;
+  padding: var(--spacing-sm);
+}
+
+.offline-banner.hidden {
+  display: none;
+}
+
+/* Responsividade */
+@media (max-width: 768px) {
+  .tabs {
+    overflow-x: auto;
+    white-space: nowrap;
+  }
+  
+  .chat-input {
+    flex-direction: column;
+  }
+}
+
+/* Acessibilidade - preferência por movimento reduzido */
+@media (prefers-reduced-motion: reduce) {
+  * {
+    animation-duration: 0.01ms !important;
+    animation-iteration-count: 1 !important;
+    transition-duration: 0.01ms !important;
+    scroll-behavior: auto !important;
+  }
+}
+
+/* High contrast mode */
+@media (prefers-contrast: high) {
+  :root {
+    --primary: #0000ff;
+    --secondary: #008000;
+    --error: #ff0000;
+    --warning: #cc6600;
+    --background: #ffffff;
+    --text-primary: #000000;
+    --border: #000000;
+  }
+}
+
+/* Dark mode */
+@media (prefers-color-scheme: dark) {
+  :root {
+    --primary: #4db8ff;
+    --primary-dark: #007acc;
+    --primary-light: #1c313a;
+    --background: #121212;
+    --surface: #1e1e1e;
+    --text-primary: #ffffff;
+    --text-secondary: #aaaaaa;
+    --border: #333333;
+  }
+  
+  .card {
+    background-color: #252525;
+  }
+  
+  input, textarea, select {
+    background-color: #333333;
+    color: var(--text-primary);
+    border-color: #444444;
+  }
+  
+  .message-bot .message-bubble {
+    background-color: #252525;
+  }
+  
+  .message-user .message-bubble {
+    background-color: #1c313a;
+  }
+  
+  .offline-banner {
+    background-color: #664400;
+  }
+}
\ No newline at end of file
diff --git a/flipapp_runtime_full/runtime.html b/flipapp_runtime_full/runtime.html
new file mode 100644
index 0000000000000000000000000000000000000000..f143773c4db7c2e58780045c764e39a9176a4c74
--- /dev/null
+++ b/flipapp_runtime_full/runtime.html
@@ -0,0 +1,12 @@
+<!DOCTYPE html>
+<html>
+<head>
+  <meta charset="UTF-8" />
+  <title>FlipApp Runtime</title>
+  <link rel="stylesheet" href="style.css">
+</head>
+<body>
+  <div id="app"></div>
+  <script src="app.js"></script>
+</body>
+</html>
\ No newline at end of file
diff --git a/flipapp_runtime_full/runtime/config.logline b/flipapp_runtime_full/runtime/config.logline
new file mode 100644
index 0000000000000000000000000000000000000000..013b4a0d6554f4183e5753c879155872e1273a47
--- /dev/null
+++ b/flipapp_runtime_full/runtime/config.logline
@@ -0,0 +1,8 @@
+@@
+   {
+-    "key": "app_entry_point",
+-    "value": "/ui/flipapp_ui.logline",
++    "key": "app_entry_point",
++    "value": "/ui/whatsapp_ui.logline",
+     "description": "Main UI entry point file path"
+   },
\ No newline at end of file
diff --git a/flipapp_runtime_full/runtime/expression_engine.logline b/flipapp_runtime_full/runtime/expression_engine.logline
new file mode 100644
index 0000000000000000000000000000000000000000..7b4878d50495322a7785dc405fa6698286f9d7e4
--- /dev/null
+++ b/flipapp_runtime_full/runtime/expression_engine.logline
@@ -0,0 +1,446 @@
+# runtime/expression_engine.logline
+# Definição declarativa do motor de expressão seguro para o LogLineOS.
+
+- type: token
+  name: WHITESPACE
+  pattern: "\\s+"
+  skip: true
+
+- type: token
+  name: NUMBER
+  pattern: "[0-9]+(\\.[0-9]+)?"
+  capture: true
+
+- type: token
+  name: STRING
+  pattern: "\\"([^\\"\\]|\\\\.)*\\""
+  capture: true
+
+- type: token
+  name: IDENTIFIER
+  pattern: "[A-Za-z_][A-Za-z0-9_]*"
+  capture: true
+
+- type: token
+  name: PLUS
+  pattern: "\\+"
+
+- type: token
+  name: MINUS
+  pattern: "-"
+
+- type: token
+  name: ASTERISK
+  pattern: "\\*"
+
+- type: token
+  name: SLASH
+  pattern: "/"
+
+- type: token
+  name: LPAREN
+  pattern: "\\("
+
+- type: token
+  name: RPAREN
+  pattern: "\\)"
+
+- type: token
+  name: AND
+  pattern: "and\\b"
+
+- type: token
+  name: OR
+  pattern: "or\\b"
+
+- type: token
+  name: EQ
+  pattern: "=="
+
+- type: token
+  name: NEQ
+  pattern: "!="
+
+- type: token
+  name: GT
+  pattern: ">"
+
+- type: token
+  name: LT
+  pattern: "<"
+
+- type: token
+  name: GTE
+  pattern: ">="
+
+- type: token
+  name: LTE
+  pattern: "<="
+
+
+# Regras de gramática e AST
+- type: rule
+  name: LiteralNumber
+  pattern:
+    - type: token
+      name: NUMBER
+  emit:
+    type: ast
+    node_type: NumberLiteral
+    transform: |
+      function(token) {
+        return { type: "NumberLiteral", value: parseFloat(token.text) };
+      }
+
+- type: rule
+  name: LiteralString
+  pattern:
+    - type: token
+      name: STRING
+  emit:
+    type: ast
+    node_type: StringLiteral
+    transform: |
+      function(token) {
+        let text = token.text.slice(1, -1);
+        return { type: "StringLiteral", value: JSON.parse(`"${text}"`) };
+      }
+
+- type: rule
+  name: Identifier
+  pattern:
+    - type: token
+      name: IDENTIFIER
+  emit:
+    type: ast
+    node_type: Identifier
+    transform: |
+      function(token) {
+        return { type: "Identifier", name: token.text };
+      }
+
+- type: rule
+  name: GroupExpression
+  pattern:
+    - type: token
+      name: LPAREN
+    - type: rule
+      name: Expression
+    - type: token
+      name: RPAREN
+  emit:
+    type: ast
+    node_type: GroupExpression
+    transform: |
+      function(_lp, exprNode, _rp) {
+        return exprNode;
+      }
+
+- type: rule
+  name: Term
+  choice:
+    - type: rule
+      name: LiteralNumber
+    - type: rule
+      name: LiteralString
+    - type: rule
+      name: Identifier
+    - type: rule
+      name: GroupExpression
+
+- type: rule
+  name: Factor
+  pattern:
+    - type: rule
+      name: Term
+    - repeat:
+        - pattern:
+            - type: choice
+              options:
+                - type: token
+                  name: ASTERISK
+                - type: token
+                  name: SLASH
+            - type: rule
+              name: Term
+  emit:
+    type: ast
+    node_type: BinaryExpression
+    transform: |
+      function(firstTerm, rest) {
+        return rest.reduce((leftNode, [opToken, rightTerm]) => ({
+          type: "BinaryExpression",
+          operator: opToken.text,
+          left: leftNode,
+          right: rightTerm
+        }), firstTerm);
+      }
+
+- type: rule
+  name: AddExpr
+  pattern:
+    - type: rule
+      name: Factor
+    - repeat:
+        - pattern:
+            - type: choice
+              options:
+                - type: token
+                  name: PLUS
+                - type: token
+                  name: MINUS
+            - type: rule
+              name: Factor
+  emit:
+    type: ast
+    node_type: BinaryExpression
+    transform: |
+      function(firstFactor, rest) {
+        return rest.reduce((leftNode, [opToken, rightFactor]) => ({
+          type: "BinaryExpression",
+          operator: opToken.text,
+          left: leftNode,
+          right: rightFactor
+        }), firstFactor);
+      }
+
+- type: rule
+  name: CompareExpr
+  pattern:
+    - type: rule
+      name: AddExpr
+    - optional:
+        - pattern:
+            - type: choice
+              options:
+                - type: token
+                  name: GT
+                - type: token
+                  name: LT
+                - type: token
+                  name: GTE
+                - type: token
+                  name: LTE
+                - type: token
+                  name: EQ
+                - type: token
+                  name: NEQ
+            - type: rule
+              name: AddExpr
+  emit:
+    type: ast
+    node_type: BinaryExpression
+    transform: |
+      function(leftAdd, maybeOpRight) {
+        if (!maybeOpRight) return leftAdd;
+        let [opToken, rightAdd] = maybeOpRight;
+        return {
+          type: "BinaryExpression",
+          operator: opToken.text,
+          left: leftAdd,
+          right: rightAdd
+        };
+      }
+
+- type: rule
+  name: LogicExpr
+  pattern:
+    - type: rule
+      name: CompareExpr
+    - repeat:
+        - pattern:
+            - type: choice
+              options:
+                - type: token
+                  name: AND
+                - type: token
+                  name: OR
+            - type: rule
+              name: CompareExpr
+  emit:
+    type: ast
+    node_type: LogicalExpression
+    transform: |
+      function(firstCompare, rest) {
+        return rest.reduce((leftNode, [opToken, rightCompare]) => ({
+          type: "LogicalExpression",
+          operator: opToken.text,
+          left: leftNode,
+          right: rightCompare
+        }), firstCompare);
+      }
+
+- type: rule
+  name: Expression
+  pattern:
+    - type: rule
+      name: LogicExpr
+  emit:
+    type: ast
+    node_type: Expression
+    transform: |
+      function(node) {
+        return node;
+      }
+
+- type: contract
+  name: EvalNumberLiteral
+  args:
+    - name: astNode
+      type: NumberLiteral
+  returns: Number
+  body: |
+    return astNode.value;
+
+- type: contract
+  name: EvalStringLiteral
+  args:
+    - name: astNode
+      type: StringLiteral
+  returns: String
+  body: |
+    return astNode.value;
+
+- type: contract
+  name: EvalIdentifier
+  args:
+    - name: astNode
+      type: Identifier
+    - name: context
+      type: Object
+  returns: Any
+  body: |
+    let varName = astNode.name;
+    if (!context.hasOwnProperty(varName)) {
+      throw Error(`Variável não definida: ${varName}`);
+    }
+    return context[varName];
+
+- type: contract
+  name: EvalBinaryExpression
+  args:
+    - name: astNode
+      type: BinaryExpression
+    - name: context
+      type: Object
+  returns: Any
+  body: |
+    let leftVal = callContract("EvalExpression", [astNode.left, context]);
+    let rightVal = callContract("EvalExpression", [astNode.right, context]);
+    switch (astNode.operator) {
+      case "+":
+        return leftVal + rightVal;
+      case "-":
+        return leftVal - rightVal;
+      case "*":
+        return leftVal * rightVal;
+      case "/":
+        if (rightVal === 0) throw Error("Divisão por zero");
+        return leftVal / rightVal;
+      case "==":
+        return leftVal === rightVal;
+      case "!=":
+        return leftVal !== rightVal;
+      case ">":
+        return leftVal > rightVal;
+      case "<":
+        return leftVal < rightVal;
+      case ">=":
+        return leftVal >= rightVal;
+      case "<=":
+        return leftVal <= rightVal;
+      default:
+        throw Error(`Operador não suportado: ${astNode.operator}`);
+    }
+
+- type: contract
+  name: EvalLogicalExpression
+  args:
+    - name: astNode
+      type: LogicalExpression
+    - name: context
+      type: Object
+  returns: Boolean
+  body: |
+    let leftVal = callContract("EvalExpression", [astNode.left, context]);
+    switch (astNode.operator) {
+      case "and":
+        if (!leftVal) return false;
+        return callContract("EvalExpression", [astNode.right, context]);
+      case "or":
+        if (leftVal) return true;
+        return callContract("EvalExpression", [astNode.right, context]);
+      default:
+        throw Error(`Operador lógico não suportado: ${astNode.operator}`);
+    }
+
+- type: contract
+  name: EvalExpression
+  args:
+    - name: astNode
+      type: Expression
+    - name: context
+      type: Object
+  returns: Any
+  body: |
+    switch (astNode.type) {
+      case "NumberLiteral":
+        return callContract("EvalNumberLiteral", [astNode]);
+      case "StringLiteral":
+        return callContract("EvalStringLiteral", [astNode]);
+      case "Identifier":
+        return callContract("EvalIdentifier", [astNode, context]);
+      case "BinaryExpression":
+        return callContract("EvalBinaryExpression", [astNode, context]);
+      case "LogicalExpression":
+        return callContract("EvalLogicalExpression", [astNode, context]);
+      case "GroupExpression":
+        return callContract("EvalExpression", [astNode, context]);
+      default:
+        throw Error(`Tipo de nó AST não reconhecido: ${astNode.type}`);
+    }
+
+- type: span
+  id: Lexer
+  inputs:
+    - name: inputText
+      type: String
+  outputs:
+    - name: tokens
+      type: Array<Token>
+  body: |
+    return scanTokens(inputText, [
+      "WHITESPACE", "NUMBER", "STRING", "IDENTIFIER", 
+      "PLUS", "MINUS", "ASTERISK", "SLASH", 
+      "LPAREN", "RPAREN", "AND", "OR", 
+      "EQ", "NEQ", "GT", "LT", "GTE", "LTE"
+    ]);
+
+- type: span
+  id: Parser
+  inputs:
+    - name: tokens
+      type: Array<Token>
+    - name: startRule
+      type: String
+  outputs:
+    - name: ast
+      type: Expression
+  body: |
+    return parseTokens(tokens, startRule);
+
+- type: span
+  id: TokenizeAndEval
+  inputs:
+    - name: inputText
+      type: String
+    - name: context
+      type: Object
+  outputs:
+    - name: result
+      type: Any
+  body: |
+    let tokens = callSpan("Lexer", [inputText]);
+    let astTree = callSpan("Parser", [tokens, "Expression"]);
+    let value = callContract("EvalExpression", [astTree, context]);
+    return { result: value };
diff --git a/flipapp_runtime_full/runtime/llm_config.logline b/flipapp_runtime_full/runtime/llm_config.logline
new file mode 100644
index 0000000000000000000000000000000000000000..fad0154aa3600323755037afa6e7bbac51aa6f53
--- /dev/null
+++ b/flipapp_runtime_full/runtime/llm_config.logline
@@ -0,0 +1,14 @@
+# llm_config.logline
+# Configuração para o LLM: URLs e chave de API
+
+- type: config_param
+  key: "llm_proxy_url"
+  value: "http://localhost:8000/api/llm"
+
+- type: config_param
+  key: "default_llm_model"
+  value: "mistral-7b"
+
+- type: config_param
+  key: "openai_api_key"
+  value: "{{ env.OPENAI_API_KEY }}"
diff --git a/flipapp_runtime_full/runtime/llm_request.logline b/flipapp_runtime_full/runtime/llm_request.logline
new file mode 100644
index 0000000000000000000000000000000000000000..da943a4aa6798476f3d48c47f0d848bfd44f05da
--- /dev/null
+++ b/flipapp_runtime_full/runtime/llm_request.logline
@@ -0,0 +1,10 @@
+# llm_request.logline
+# Declaração de um span para requisição LLM
+
+- type: llm_request
+  id: "{{generate_uuid()}}"
+  prompt: "{{state.input_text}}"
+  model: "mistral-7b"
+  max_tokens: 200
+  temperature: 0.7
+  on_response: "apply_completion_result"
diff --git a/flipapp_runtime_full/runtime/span_execution_rules.logline b/flipapp_runtime_full/runtime/span_execution_rules.logline
new file mode 100644
index 0000000000000000000000000000000000000000..b7df5eef09f907146c05a6678032532066852568
--- /dev/null
+++ b/flipapp_runtime_full/runtime/span_execution_rules.logline
@@ -0,0 +1,15 @@
+# span_execution_rules.logline
+# Regras declarativas para execução de spans, incluindo LLM
+
+- type: execution_rule
+  for_span_type: llm_request
+  description: "Regra para processar llm_request, chamando o backend apropriado"
+  execution_logic:
+    - type: kernel_action
+      action: "call_llm_backend"
+      args:
+        prompt: "{{span.prompt}}"
+        model: "{{span.model}}"
+        temperature: "{{span.temperature}}"
+        max_tokens: "{{span.max_tokens}}"
+      response_contract: "{{span.on_response}}"
diff --git a/flipapp_runtime_full/spans/hardening_spans.jsonl b/flipapp_runtime_full/spans/hardening_spans.jsonl
new file mode 100644
index 0000000000000000000000000000000000000000..3d2b1ab6f2c541e1bd338d433c785c9cdfa7765f
--- /dev/null
+++ b/flipapp_runtime_full/spans/hardening_spans.jsonl
@@ -0,0 +1,2 @@
+{"type": "refactor_task", "created_by": "dan.amarilho", "created_at": "2025-06-04T00:00:00Z", "status": "pending", "id": "task_llm_fallback_local", "title": "Adicionar fallback local ao llm_proxy_hard.py", "description": "Implementar mecanismo para usar WASM local (mistral.wasm) antes de chamar LLM remoto. Incluir cache sem\u00e2ntico por hash de prompt."}
+{"type": "refactor_task", "created_by": "dan.amarilho", "created_at": "2025-06-04T00:00:00Z", "status": "pending", "id": "task_llm_timeout_retry", "title": "Configurar timeout e retry no llm_proxy_hard.py", "description": "Adicionar `timeout`, `retry` e tratamento de exce\u00e7\u00f5es no proxy LLM. Garantir logs claros em stderr e resposta padronizada em JSON."}
diff --git a/flipapp_runtime_full/src/core/expression-engine.js b/flipapp_runtime_full/src/core/expression-engine.js
new file mode 100644
index 0000000000000000000000000000000000000000..491ba3d9e3d1ab2fc323bd8f8adf839a140a94ae
--- /dev/null
+++ b/flipapp_runtime_full/src/core/expression-engine.js
@@ -0,0 +1,162 @@
+// src/core/expression-engine.js
+/**
+ * Motor de expressão seguro para substituir eval()
+ * @module core/expression-engine
+ */
+
+/**
+ * Erros específicos do motor de expressão
+ */
+export class ExpressionError extends Error {
+  constructor(message, position = null, code = 'EXPRESSION_ERROR') {
+    super(message);
+    this.name = 'ExpressionError';
+    this.position = position;
+    this.code = code;
+  }
+}
+
+/**
+ * Motor de expressão seguro para substituir eval()
+ */
+export class ExpressionEngine {
+  constructor() {
+    // Operadores permitidos
+    this.allowedOperators = [
+      '===', '!==', '==', '!=', '>', '<', '>=', '<=',
+      '&&', '||', '!', '+', '-', '*', '/', '%',
+      '?', ':', '(', ')', '[', ']', '.', ','
+    ];
+    
+    // Funções seguras permitidas
+    this.allowedFunctions = new Map([
+      ['includes', (arr, item) => Array.isArray(arr) ? arr.includes(item) : String(arr).includes(item)],
+      ['length', (obj) => obj?.length || 0],
+      ['toString', (obj) => String(obj || '')],
+      ['toLowerCase', (str) => String(str || '').toLowerCase()],
+      ['toUpperCase', (str) => String(str || '').toUpperCase()],
+      ['trim', (str) => String(str || '').trim()],
+      ['parseInt', (str) => parseInt(str, 10) || 0],
+      ['parseFloat', (str) => parseFloat(str) || 0],
+      ['max', (...args) => Math.max(...args)],
+      ['min', (...args) => Math.min(...args)],
+      ['startsWith', (str, search) => String(str || '').startsWith(search || '')],
+      ['endsWith', (str, search) => String(str || '').endsWith(search || '')],
+      ['replace', (str, search, replaceValue) => String(str || '').replace(search || '', replaceValue || '')],
+      ['join', (arr, separator) => Array.isArray(arr) ? arr.join(separator || '') : ''],
+      ['slice', (arr, start, end) => Array.isArray(arr) ? arr.slice(start, end) : 
+                                     typeof arr === 'string' ? arr.slice(start, end) : []],
+      ['map', (arr, fn) => Array.isArray(arr) && typeof fn === 'function' ? 
+                          arr.map(item => fn(item)) : []],
+      ['filter', (arr, fn) => Array.isArray(arr) && typeof fn === 'function' ? 
+                             arr.filter(item => fn(item)) : []]
+    ]);
+    
+    // Cache de expressões já avaliadas para desempenho
+    this.cache = new Map();
+    this.cacheHits = 0;
+    this.cacheMisses = 0;
+  }
+
+  /**
+   * Avalia uma expressão de forma segura
+   * @param {string} expression - Expressão a ser avaliada
+   * @param {Object} context - Contexto de variáveis
+   * @returns {*} Resultado da avaliação
+   * @throws {ExpressionError} Em caso de erro de sintaxe ou expressão proibida
+   */
+  evaluate(expression, context = {}) {
+    if (typeof expression !== 'string') {
+      return expression;
+    }
+
+    // Expressões vazias retornam undefined
+    expression = expression.trim();
+    if (!expression) {
+      return undefined;
+    }
+
+    // Verifica expressões proibidas
+    this._validateSafety(expression);
+
+    // Verifica cache (apenas para expressões pequenas)
+    if (expression.length < 100) {
+      const cacheKey = `${expression}::${JSON.stringify(context)}`;
+      if (this.cache.has(cacheKey)) {
+        this.cacheHits++;
+        return this.cache.get(cacheKey);
+      }
+      this.cacheMisses++;
+    }
+
+    try {
+      const tokens = this._tokenize(expression);
+      const ast = this._parse(tokens);
+      const result = this._evaluate(ast, context);
+      
+      // Adiciona ao cache se a expressão for pequena
+      if (expression.length < 100) {
+        const cacheKey = `${expression}::${JSON.stringify(context)}`;
+        if (this.cache.size > 1000) { // Limite de tamanho do cache
+          this.cache.clear();
+        }
+        this.cache.set(cacheKey, result);
+      }
+      
+      return result;
+    } catch (error) {
+      if (error instanceof ExpressionError) {
+        throw error;
+      }
+      throw new ExpressionError(`Erro ao avaliar expressão: ${expression}. ${error.message}`);
+    }
+  }
+
+  /**
+   * Interpola variáveis em uma string template
+   * @param {string} template - String com {{variable}} placeholders
+   * @param {Object} context - Contexto de variáveis
+   * @returns {string} String interpolada
+   */
+  interpolate(template, context = {}) {
+    if (typeof template !== 'string') {
+      return String(template || '');
+    }
+
+    return template.replace(/\{\{([^}]+)\}\}/g, (match, expression) => {
+      try {
+        const result = this.evaluate(expression.trim(), context);
+        return result !== undefined && result !== null ? String(result) : '';
+      } catch (error) {
+        console.error(`[ExpressionEngine] Erro na interpolação: ${expression}`, error);
+        return ''; // Retorna string vazia em caso de erro
+      }
+    });
+  }
+
+  /**
+   * Valida se uma expressão é segura
+   * @param {string} expression - Expressão a validar
+   * @throws {ExpressionError} Se a expressão contiver construções perigosas
+   * @private
+   */
+  _validateSafety(expression) {
+    const forbiddenPatterns = [
+      { pattern: /eval\s*\(/, message: 'Uso de eval() não permitido' },
+      { pattern: /Function\s*\(/, message: 'Construtor Function não permitido' },
+      { pattern: /setTimeout|setInterval/, message: 'Funções de timer não permitidas' },
+      { pattern: /new\s+Function/, message: 'Construtor Function não permitido' },
+      { pattern: /\bwindow\b|\bdocument\b|\blocation\b/, message: 'Acesso a objetos globais não permitido' },
+      { pattern: /\bObject\b\s*\.\s*\b(defineProperty|__proto__|constructor|prototype)\b/, message: 'Acesso a propriedades de Object não permitido' },
+      { pattern: /__proto__|prototype|constructor/, message: 'Acesso a propriedades internas não permitido' }
+    ];
+
+    for (const { pattern, message } of forbiddenPatterns) {
+      if (pattern.test(expression)) {
+        throw new ExpressionError(message, null, 'SECURITY_VIOLATION');
+      }
+    }
+  }
+
+  /**
+   
\ No newline at end of file
diff --git a/flipapp_runtime_full/src/core/state.js b/flipapp_runtime_full/src/core/state.js
new file mode 100644
index 0000000000000000000000000000000000000000..fef9e2a7684d9ef0d6e93d5276f20d8e5a0bfc6d
--- /dev/null
+++ b/flipapp_runtime_full/src/core/state.js
@@ -0,0 +1,293 @@
+// src/core/state.js
+/**
+ * Sistema de gerenciamento de estado reativo centralizado
+ * @module core/state
+ */
+
+/**
+ * Gerenciador de estado com padrão Observer
+ */
+export class StateManager {
+  constructor() {
+    this.state = {};
+    this.observers = new Map(); // path -> Set<callback>
+    this.middlewares = [];
+    this.localStorage = typeof window !== 'undefined' ? window.localStorage : null;
+    this.sessionId = null;
+  }
+
+  /**
+   * Define um valor no estado com notificação reativa
+   * @param {string} path - Caminho dot-notation (ex: 'user.profile.name')
+   * @param {*} value - Valor a ser definido
+   */
+  setState(path, value) {
+    const oldValue = this.getState(path);
+    if (JSON.stringify(oldValue) === JSON.stringify(value)) return; // Evita re-renders desnecessários
+
+    this._setNestedValue(path, value);
+    
+    // Executa middlewares
+    this.middlewares.forEach(middleware => {
+      try {
+        middleware({ type: 'SET_STATE', path, value, oldValue });
+      } catch (e) {
+        console.error('Middleware error:', e);
+      }
+    });
+
+    this.notifyObservers(path, value);
+    this._persistCriticalState(path, value);
+  }
+
+  /**
+   * Obtém valor do estado
+   * @param {string} path - Caminho para a propriedade
+   * @returns {*} Valor ou undefined
+   */
+  getState(path) {
+    if (!path || path === '*') return this.state;
+    return this._getNestedValue(path);
+  }
+
+  /**
+   * Adiciona item a um array no estado
+   * @param {string} path - Caminho para o array
+   * @param {*} value - Item a ser adicionado
+   */
+  appendState(path, value) {
+    const currentArray = this.getState(path);
+    if (Array.isArray(currentArray)) {
+      this.setState(path, [...currentArray, value]);
+    } else {
+      console.warn(`[StateManager] Tentativa de append em não-array: ${path}`);
+      this.setState(path, [value]);
+    }
+  }
+
+  /**
+   * Remove item de array por índice ou condição
+   * @param {string} path - Caminho para o array
+   * @param {number|function} condition - Índice ou função de filtro
+   */
+  removeFromState(path, condition) {
+    const currentArray = this.getState(path);
+    if (!Array.isArray(currentArray)) return;
+
+    let newArray;
+    if (typeof condition === 'number') {
+      newArray = currentArray.filter((_, index) => index !== condition);
+    } else if (typeof condition === 'function') {
+      newArray = currentArray.filter(item => !condition(item));
+    }
+    
+    if (newArray) {
+      this.setState(path, newArray);
+    }
+  }
+
+  /**
+   * Subscreve a mudanças de estado
+   * @param {string} path - Caminho a observar
+   * @param {function} callback - Função chamada na mudança
+   * @returns {function} Função de cleanup para remover observador
+   */
+  subscribe(path, callback) {
+    if (!this.observers.has(path)) {
+      this.observers.set(path, new Set());
+    }
+    this.observers.get(path).add(callback);
+
+    // Chama callback imediatamente com valor atual
+    callback(this.getState(path), path);
+
+    // Retorna função de cleanup
+    return () => this.unsubscribe(path, callback);
+  }
+
+  /**
+   * Remove subscrição
+   * @param {string} path - Caminho observado
+   * @param {function} callback - Callback a ser removido
+   */
+  unsubscribe(path, callback) {
+    const observers = this.observers.get(path);
+    if (observers) {
+      observers.delete(callback);
+      if (observers.size === 0) {
+        this.observers.delete(path);
+      }
+    }
+  }
+
+  /**
+   * Notifica observadores de mudanças
+   * @param {string} changedPath - Caminho que mudou
+   * @param {*} newValue - Novo valor
+   */
+  notifyObservers(changedPath, newValue) {
+    // Notifica observadores do caminho exato
+    this._notifyPathObservers(changedPath, newValue);
+
+    // Notifica observadores de caminhos ancestrais
+    let currentPath = changedPath;
+    while (currentPath.includes('.')) {
+      currentPath = currentPath.substring(0, currentPath.lastIndexOf('.'));
+      this._notifyPathObservers(currentPath, this.getState(currentPath));
+    }
+
+    // Notifica observador global
+    this._notifyPathObservers('*', this.state);
+  }
+
+  /**
+   * Adiciona middleware para interceptar mudanças de estado
+   * @param {function} middleware - Função middleware(event)
+   */
+  addMiddleware(middleware) {
+    if (typeof middleware !== 'function') {
+      throw new Error('Middleware deve ser uma função');
+    }
+    this.middlewares.push(middleware);
+    return () => {
+      this.middlewares = this.middlewares.filter(m => m !== middleware);
+    };
+  }
+
+  /**
+   * Inicializa o estado com dados persistidos se disponíveis
+   * @param {string} sessionId - ID da sessão
+   */
+  init(sessionId) {
+    this.sessionId = sessionId;
+    this.setState('session.id', sessionId);
+    this._loadPersistedState();
+  }
+
+  /**
+   * Limpa todo o estado
+   */
+  resetState() {
+    this.state = {};
+    this.observers.clear();
+    if (this.localStorage) {
+      // Limpa apenas as entradas do localStorage relacionadas à sessão atual
+      if (this.sessionId) {
+        Object.keys(this.localStorage).forEach(key => {
+          if (key.includes(this.sessionId)) {
+            this.localStorage.removeItem(key);
+          }
+        });
+      }
+    }
+    console.warn('[StateManager] Estado resetado');
+  }
+
+  // Métodos internos
+  _setNestedValue(path, value) {
+    if (!path) return;
+    
+    const parts = path.split('.');
+    let current = this.state;
+    
+    for (let i = 0; i < parts.length - 1; i++) {
+      if (current[parts[i]] === undefined) {
+        current[parts[i]] = {};
+      } else if (typeof current[parts[i]] !== 'object' || current[parts[i]] === null) {
+        // Converte valores não-objeto para objeto vazio
+        current[parts[i]] = {};
+      }
+      current = current[parts[i]];
+    }
+    
+    current[parts[parts.length - 1]] = value;
+  }
+
+  _getNestedValue(path) {
+    if (!path) return undefined;
+    
+    const parts = path.split('.');
+    let current = this.state;
+    
+    for (const part of parts) {
+      if (current === null || typeof current !== 'object' || !(part in current)) {
+        return undefined;
+      }
+      current = current[part];
+    }
+    
+    return current;
+  }
+
+  _notifyPathObservers(path, value) {
+    const observers = this.observers.get(path);
+    if (observers) {
+      observers.forEach(callback => {
+        try {
+          callback(value, path);
+        } catch (error) {
+          console.error(`[StateManager] Erro em observer para ${path}:`, error);
+        }
+      });
+    }
+  }
+
+  /**
+   * Persiste estados críticos no localStorage
+   */
+  _persistCriticalState(path, value) {
+    if (!this.localStorage || !this.sessionId) return;
+
+    // Lista de paths que devem ser persistidos
+    const criticalPaths = [
+      'chat.history', 
+      'user.preferences',
+      'espelho.spans',
+      'activeTab'
+    ];
+    
+    for (const criticalPath of criticalPaths) {
+      if (path === criticalPath || path.startsWith(`${criticalPath}.`)) {
+        const key = `flipapp_${criticalPath.replace(/\./g, '_')}_${this.sessionId}`;
+        try {
+          this.localStorage.setItem(key, JSON.stringify(this.getState(criticalPath)));
+        } catch (error) {
+          console.error(`[StateManager] Falha ao persistir ${criticalPath}:`, error);
+        }
+        break;
+      }
+    }
+  }
+
+  /**
+   * Carrega estados persistidos
+   */
+  _loadPersistedState() {
+    if (!this.localStorage || !this.sessionId) return;
+
+    const criticalPaths = [
+      'chat.history',
+      'user.preferences',
+      'espelho.spans',
+      'activeTab'
+    ];
+    
+    criticalPaths.forEach(path => {
+      const key = `flipapp_${path.replace(/\./g, '_')}_${this.sessionId}`;
+      const stored = this.localStorage.getItem(key);
+      
+      if (stored) {
+        try {
+          const parsedValue = JSON.parse(stored);
+          this.setState(path, parsedValue);
+          console.debug(`[StateManager] Estado "${path}" carregado de localStorage`);
+        } catch (error) {
+          console.error(`[StateManager] Falha ao carregar ${path}:`, error);
+        }
+      }
+    });
+  }
+}
+
+// Singleton global
+export const stateManager = new StateManager();
\ No newline at end of file
diff --git a/flipapp_runtime_full/state.logline b/flipapp_runtime_full/state.logline
new file mode 100644
index 0000000000000000000000000000000000000000..8b3b3fad31634673f313d9f0fa031ad3c2a753ef
--- /dev/null
+++ b/flipapp_runtime_full/state.logline
@@ -0,0 +1,249 @@
+# state.logline
+# Estado global da aplicação FlipApp LogLineOS + WhatsApp
+# Autor: danvoulez
+# Data: 2025-06-04
+
+# Estado do usuário e sessão
+user_id: ""
+session_id: ""
+app_version: "1.0.0-loglineos"
+last_activity: ""
+
+# Estado do chat WhatsApp
+conversations: {}
+  # conversation_id: {
+  #   name: string,
+  #   unread: int,
+  #   last_message: string,
+  #   last_timestamp: string,
+  #   avatar_url: string,
+  #   status: "active"|"archived"|"blocked",
+  #   ghost_sale_status: null|"tracking"|"completed"|"aborted"
+  # }
+
+active_conversation: ""
+input_message: ""
+messages: {}
+  # conversation_id: [
+  #   {
+  #     id: string,
+  #     sender: "user"|"bot"|"system",
+  #     content: string,
+  #     timestamp: string,
+  #     status: "sent"|"delivered"|"read",
+  #     type: "text"|"audio"|"image"|"document",
+  #     media_url: string|null,
+  #     transcription: string|null,
+  #     translation: string|null,
+  #     show_original: boolean,
+  #     feedback_tag: null|"good"|"bad_translation"|"ignored"
+  #   }
+  # ]
+
+# Estados de erro e loading
+error: null
+loading: false
+typing_indicator: false
+
+# Sistema Ghost Sale
+ghost_sales: {}
+  # conversation_id: {
+  #   id: string,
+  #   status: "init"|"tracking"|"committed"|"aborted",
+  #   start_time: string,
+  #   end_time: string|null,
+  #   events: [
+  #     {
+  #       type: "incentive"|"pressure"|"commitment"|"abandonment",
+  #       timestamp: string,
+  #       context: object
+  #     }
+  #   ],
+  #   penalty_score: int,
+  #   agent_id: string
+  # }
+
+# Sistema de Auditoria e Julgamento
+affairs: []
+  # [
+  #   {
+  #     id: string,
+  #     type: "ghost_sale"|"agent_misconduct"|"system_violation",
+  #     status: "pending"|"investigating"|"resolved"|"dismissed",
+  #     created_at: string,
+  #     evidence: object,
+  #     involved_parties: [string],
+  #     resolution: object|null
+  #   }
+  # ]
+
+judgements: []
+  # [
+  #   {
+  #     id: string,
+  #     type: "fail_sale"|"negligence"|"misconduct",
+  #     target: string, # agent_id
+  #     timestamp: string,
+  #     evidence: object,
+  #     status: "pending"|"confirmed"|"appealed"|"dismissed",
+  #     penalty_applied: object|null
+  #   }
+  # ]
+
+# Intervenções do Agente Simbólico
+agent_interventions: []
+  # [
+  #   {
+  #     id: string,
+  #     conversation_id: string,
+  #     context_hash: string,
+  #     policy_used: string,
+  #     analysis: object,
+  #     action: "suggest"|"alert"|"intervene"|"escalate",
+  #     timestamp: string,
+  #     result: object|null
+  #   }
+  # ]
+
+# Sistema de Alertas
+alerts: []
+  # [
+  #   {
+  #     id: string,
+  #     level: int, # 1-10
+  #     type: "risk"|"capacity"|"system"|"agent",
+  #     message: string,
+  #     conversation_id: string|null,
+  #     timestamp: string,
+  #     acknowledged: boolean,
+  #     action_taken: string|null
+  #   }
+  # ]
+
+# Resumos de Conversas
+summaries: {}
+  # conversation_id: {
+  #   content: string,
+  #   timestamp: string,
+  #   tokens_used: int,
+  #   confidence: float
+  # }
+
+# Sistema de Treinamento
+training_samples: []
+  # [
+  #   {
+  #     id: string,
+  #     message_id: string,
+  #     tag: "good"|"bad_translation"|"ignored",
+  #     timestamp: string,
+  #     context: object,
+  #     original_content: string,
+  #     corrected_content: string|null
+  #   }
+  # ]
+
+# Métricas e Capacidade do Sistema
+system_load: []
+  # [
+  #   {
+  #     timestamp: string,
+  #     cpu_usage: float,
+  #     memory_usage: float,
+  #     active_conversations: int,
+  #     messages_per_minute: int,
+  #     llm_calls_per_minute: int
+  #   }
+  # ] # Manter apenas últimas 1440 entradas (24h)
+
+message_counter: 0
+daily_message_limit_reached: false
+
+# Gestão de Agentes
+agents: {}
+  # agent_id: {
+  #   name: string,
+  #   status: "active"|"suspended"|"training"|"inactive",
+  #   penalty_score: int,
+  #   last_activity: string,
+  #   assigned_conversations: [string],
+  #   performance_metrics: {
+  #     successful_sales: int,
+  #     ghost_sales_detected: int,
+  #     interventions_made: int,
+  #     accuracy_score: float
+  #   }
+  # }
+
+# Configurações Dinâmicas
+config:
+  llm_model: "gpt-4"
+  max_conversations: 50
+  daily_msg_limit: 1000
+  risk_threshold: 7
+  auto_translate: true
+  haptic_feedback: true
+  sound_effects: true
+  dark_mode: false
+  language: "pt-BR"
+  api_endpoints:
+    whatsapp_api: "https://api.whatsapp.com/v1"
+    llm_proxy: "http://localhost:8001"
+    translation_api: "builtin"
+    transcription_api: "builtin"
+  timeouts:
+    http_request: 30000
+    llm_response: 60000
+    file_upload: 120000
+
+# Políticas de Agentes
+agent_policies:
+  default_agent_policy:
+    analysis_model: "gpt-4"
+    prompt_template: "Analise a conversa e identifique possíveis vendas fantasma ou comportamentos suspeitos."
+    risk_threshold: 5
+    intervention_rules:
+      - condition: "ghost_sale_detected"
+        action: "suggest"
+      - condition: "high_risk_behavior"
+        action: "alert"
+      - condition: "critical_violation"
+        action: "escalate"
+  
+  strict_agent_policy:
+    analysis_model: "gpt-4"
+    prompt_template: "Aplique análise rigorosa para detectar qualquer irregularidade na conversa."
+    risk_threshold: 3
+    intervention_rules:
+      - condition: "any_suspicion"
+        action: "intervene"
+
+# Estados temporários (limpeza automática)
+temp:
+  ui_loading: false
+  last_llm_call: null
+  pending_uploads: []
+  websocket_status: "disconnected"
+  simulation_mode: false
+
+# Conexões ativas
+connections:
+  whatsapp_ws: 
+    status: "disconnected"
+    last_ping: null
+    reconnect_attempts: 0
+  llm_proxy:
+    status: "disconnected"
+    last_request: null
+    queue_size: 0
+
+# Dados de UI específicos
+ui:
+  sidebar_collapsed: false
+  active_tab: "chat"
+  theme: "whatsapp_default"
+  zoom_level: 1.0
+  notifications_enabled: true
+  scroll_positions: {}
+  selected_messages: []
+  context_menu_open: false
\ No newline at end of file
diff --git a/flipapp_runtime_full/style.css b/flipapp_runtime_full/style.css
new file mode 100644
index 0000000000000000000000000000000000000000..b609093b262a7b094d5cf81ade3a92264c9d83d3
--- /dev/null
+++ b/flipapp_runtime_full/style.css
@@ -0,0 +1 @@
+body { font-family: sans-serif; background: #111; color: #eee; margin: 0; padding: 2rem; }
\ No newline at end of file
diff --git a/flipapp_runtime_full/ui/app_boot.logline b/flipapp_runtime_full/ui/app_boot.logline
new file mode 100644
index 0000000000000000000000000000000000000000..5b0ae499ac6630923794ed19d50c3d438b4c430c
--- /dev/null
+++ b/flipapp_runtime_full/ui/app_boot.logline
@@ -0,0 +1,87 @@
+[
+  {
+    "id": "contract_load_main_ui",
+    "timestamp": "2025-06-04T02:15:34Z",
+    "type": "contract_definition",
+    "name": "load_main_ui",
+    "version": "1.0.0-loglineos",
+    "description": "Loads and renders the main UI for the application.",
+    "created_by": "LogLineOS Kernel",
+    "args_schema": {},
+    "effects": [
+      {
+        "id": "effect_show_loading_skeleton",
+        "type": "ui_display_skeleton",
+        "selector": "#app_root",
+        "skeleton_type": "full_screen_app_layout"
+      },
+      {
+        "id": "effect_fetch_main_ui_logline",
+        "type": "http_request",
+        "method": "GET",
+        "url_template": "{{state.config.app_entry_point}}",
+        "response_path": "temp.main_ui_logline_source",
+        "error_next_span_id": "handle_ui_load_error"
+      },
+      {
+        "id": "effect_parse_main_ui_logline",
+        "type": "kernel_action",
+        "action_type": "parse_logline_source",
+        "source_template": "{{temp.main_ui_logline_source.body}}",
+        "response_path": "temp.main_ui_ast",
+        "error_next_span_id": "handle_ui_parse_error"
+      },
+      {
+        "id": "effect_validate_ui_schema",
+        "type": "contract_call",
+        "contract_name": "ui_schema_validation",
+        "args": {
+          "file_path": "{{state.config.app_entry_point}}",
+          "selector": "#app_root"
+        },
+        "error_next_span_id": "handle_ui_schema_error"
+      },
+      {
+        "id": "effect_render_main_ui",
+        "type": "ui_render_from_ast",
+        "ast_template": "{{temp.main_ui_ast}}",
+        "target_selector": "#app_root"
+      },
+      {
+        "id": "effect_hide_loading_skeleton",
+        "type": "ui_hide_skeleton",
+        "selector": "#app_root"
+      },
+      {
+        "type": "audit_log",
+        "level": "info",
+        "message": "Main UI loaded and rendered."
+      }
+    ]
+  },
+  {
+    "id": "rule_ui_display_skeleton",
+    "timestamp": "2025-06-04T02:15:39Z",
+    "type": "execution_rule",
+    "description": "Displays a skeleton loading screen on the UI.",
+    "match": { "type": "ui_display_skeleton" },
+    "kernel_action": {
+      "action_type": "invoke_native_ui_display_skeleton",
+      "selector_template": "{{span.selector}}",
+      "skeleton_type_template": "{{span.skeleton_type}}"
+    },
+    "audit_event_type": "ui_display_skeleton_event"
+  },
+  {
+    "id": "rule_ui_hide_skeleton",
+    "timestamp": "2025-06-04T02:15:44Z",
+    "type": "execution_rule",
+    "description": "Hides a skeleton loading screen.",
+    "match": { "type": "ui_hide_skeleton" },
+    "kernel_action": {
+      "action_type": "invoke_native_ui_hide_skeleton",
+      "selector_template": "{{span.selector}}"
+    },
+    "audit_event_type": "ui_hide_skeleton_event"
+  }
+]
\ No newline at end of file
diff --git a/flipapp_runtime_full/ui/chat_ui.logline b/flipapp_runtime_full/ui/chat_ui.logline
new file mode 100644
index 0000000000000000000000000000000000000000..a8093ed812c7c645e755a31174674bfa27565edb
--- /dev/null
+++ b/flipapp_runtime_full/ui/chat_ui.logline
@@ -0,0 +1,83 @@
+# ui/chat_ui.logline
+# UI da aba de chat
+- type: container
+  style: "display:flex; flex-direction:column; height:100%; max-width:800px; margin:0 auto"
+  children:
+    - type: container
+      class: "chat-header"
+      style: "padding:1rem; border-bottom:1px solid #eee; display:flex; justify-content:space-between; align-items:center"
+      children:
+        - type: text
+          content: "💬 Chat"
+          style: "font-size:1.2rem; font-weight:bold; margin:0"
+        - type: container
+          style: "display:flex; align-items:center"
+          children:
+            - type: when
+              when: "connection['espelho'] && connection['espelho'].status === 'connected'"
+              children:
+                - type: Badge
+                  content: "Online"
+                  color: "#e6f7e9"
+                  textColor: "#1b873b"
+                  style: "margin-right:0.5rem"
+            - type: when
+              when: "connection['espelho'] && connection['espelho'].status !== 'connected'"
+              children:
+                - type: Badge
+                  content: "Offline"
+                  color: "#fff5f5"
+                  textColor: "#dd3333"
+                  style: "margin-right:0.5rem"
+    - type: container
+      class: "chat-history-container"
+      style: "flex:1; overflow-y:auto; padding:1rem; background:#f9f9f9"
+      children:
+        # Welcome message on empty chat
+        - type: when
+          when: "!chat.history || chat.history.length === 0"
+          children:
+            - type: container
+              style: "text-align:center; margin:2rem 0"
+              children:
+                - type: text
+                  content: "👋 Bem-vindo ao FlipApp!"
+                  style: "font-size:1.5rem; font-weight:bold; margin-bottom:1rem"
+                - type: text
+                  content: "Envie uma mensagem para começar a conversa."
+                  style: "color:#666"
+        # Chat history loop
+        - type: loop
+          data: "chat.history"
+          children:
+            - type: ChatMessage
+              sender: "{{item.sender}}"
+              content: "{{item.content}}"
+              timestamp: "{{item.timestamp}}"
+    # Typing indicator
+    - type: when
+      when: "chat.isTyping"
+      children:
+        - type: container
+          style: "padding:0.5rem 1rem; font-style:italic; color:#888"
+          children:
+            - type: text
+              content: "Bot está digitando..."
+              style: "display:inline-block"
+              class: "typing-indicator"
+    # Input area
+    - type: container
+      class: "chat-input"
+      style: "padding:1rem; border-top:1px solid #eee; display:flex; gap:0.5rem"
+      children:
+        - type: input
+          bind: "chat.messageInput | userTypingDebounce"
+          placeholder: "Digite sua mensagem..."
+          style: "flex:1; padding:0.75rem; border:1px solid #ddd; border-radius:4px; font-size:1rem"
+          on: "keypress: inputKeyPress"
+        - type: Button
+          content: "Enviar"
+          variant: "primary"
+          onClick: "sendMessage"
+          disabled: "{{!chat.messageInput || chat.messageInput.trim() === ''}}"
+          style: "background:#1a8cff; color:white"
\ No newline at end of file
diff --git a/flipapp_runtime_full/ui/components/button.logline b/flipapp_runtime_full/ui/components/button.logline
new file mode 100644
index 0000000000000000000000000000000000000000..7e5990bd912f28616425544d8bcaa1b0176a4b2a
--- /dev/null
+++ b/flipapp_runtime_full/ui/components/button.logline
@@ -0,0 +1,21 @@
+# ui/components/button.logline
+# Componente Button avançado com estados
+- type: button
+  class: "btn {{props.variant ? 'btn-' + props.variant : ''}} {{props.className || ''}} {{props.loading ? 'btn-loading' : ''}} {{props.disabled ? 'btn-disabled' : ''}}"
+  style: "position:relative; padding:0.5rem 1rem; border:none; border-radius:4px; font-size:1rem; cursor:pointer; transition:background 0.2s ease; {{props.style}}"
+  disabled: "{{props.disabled || props.loading || false}}"
+  content: "{{props.loading ? '' : props.label || props.content || ''}}"
+  on: "click: {{props.loading || props.disabled ? 'noop' : (props.onClick || 'noop')}}"
+  children:
+    - type: when
+      when: "props.loading"
+      children:
+        - type: container
+          class: "loader"
+          style: "display:inline-block; width:16px; height:16px; border:2px solid rgba(255,255,255,0.3); border-radius:50%; border-top-color:white; animation:spin 1s linear infinite"
+    - type: when
+      when: "props.icon && !props.loading"
+      children:
+        - type: container
+          style: "margin-right:0.5rem"
+          content: "{{props.icon}}"
\ No newline at end of file
diff --git a/flipapp_runtime_full/ui/components/chat-message.logline b/flipapp_runtime_full/ui/components/chat-message.logline
new file mode 100644
index 0000000000000000000000000000000000000000..b3f85b442a7f511c14ba38880abb39c5e589a538
--- /dev/null
+++ b/flipapp_runtime_full/ui/components/chat-message.logline
@@ -0,0 +1,25 @@
+# ui/components/chat-message.logline
+# Componente para mensagens de chat com diferentes estilos para usuário e bot
+- type: container
+  class: "message {{props.sender === 'user' ? 'message-user' : 'message-bot'}} {{props.className || ''}}"
+  style: "display:flex; margin:0.5rem 0; {{props.sender === 'user' ? 'flex-direction:row-reverse' : 'flex-direction:row'}}; {{props.style}}"
+  children:
+    - type: container
+      class: "message-avatar"
+      style: "width:32px; height:32px; flex-shrink:0; border-radius:50%; margin:0 0.5rem; background:{{props.sender === 'user' ? '#e3f2fd' : '#f0f4c3'}}; display:flex; align-items:center; justify-content:center"
+      children:
+        - type: text
+          content: "{{props.sender === 'user' ? '👤' : '🤖'}}"
+    - type: container
+      class: "message-bubble"
+      style: "padding:0.75rem; border-radius:8px; max-width:80%; background:{{props.sender === 'user' ? '#e3f2fd' : '#f0f4c3'}}; {{props.sender === 'user' ? 'margin-right:0.5rem' : 'margin-left:0.5rem'}}"
+      children:
+        - type: markdown
+          content: "{{props.content || ''}}"
+          style: "margin:0"
+    - type: container
+      class: "message-time"
+      style: "font-size:0.7rem; color:#999; align-self:flex-end; margin-bottom:0.25rem; {{props.sender === 'user' ? 'margin-right:0.25rem' : 'margin-left:0.25rem'}}"
+      children:
+        - type: text
+          content: "{{props.timestamp ? (new Date(props.timestamp)).toLocaleTimeString([], {hour: '2-digit', minute:'2-digit'}) : ''}}"
\ No newline at end of file
diff --git a/flipapp_runtime_full/ui/components/error_overlay.logline b/flipapp_runtime_full/ui/components/error_overlay.logline
new file mode 100644
index 0000000000000000000000000000000000000000..43b2f694ab04948e0e51eddd647e944ec808480f
--- /dev/null
+++ b/flipapp_runtime_full/ui/components/error_overlay.logline
@@ -0,0 +1,58 @@
+[
+  {
+    "id": "error_overlay_component_def",
+    "timestamp": "2025-06-04T02:15:49Z",
+    "type": "component_definition",
+    "name": "ErrorOverlay",
+    "version": "1.0.0-loglineos",
+    "description": "Displays a full-screen error overlay for critical UI failures.",
+    "properties": {
+      "message": { "type": "string", "required": true },
+      "details": { "type": "string", "required": false }
+    },
+    "template": [
+      {
+        "type": "container",
+        "class": "error-overlay",
+        "style": "position: fixed; top: 0; left: 0; width: 100%; height: 100%; background: rgba(255, 0, 0, 0.8); color: white; display: flex; flex-direction: column; justify-content: center; align-items: center; z-index: 9999;"
+      },
+      {
+        "type": "text",
+        "content": "🚨 Erro Crítico do LogLineOS 🚨",
+        "style": "font-size: 2em; margin-bottom: 1em;"
+      },
+      {
+        "type": "text",
+        "content": "{{props.message}}",
+        "style": "font-size: 1.2em; text-align: center; max-width: 80%; margin-bottom: 1em;"
+      },
+      {
+        "type": "when",
+        "when": "props.details",
+        "children": [
+          { "type": "text", "content": "{{props.details}}", "style": "font-size: 0.9em; max-width: 80%; margin-bottom: 2em; font-family: monospace;" }
+        ]
+      },
+      {
+        "type": "button",
+        "content": "Recarregar Aplicação",
+        "on": "click: reload_application",
+        "style": "padding: 1em 2em; background: #fff; color: #d32f2f; border: none; border-radius: 5px; cursor: pointer;"
+      }
+    ]
+  },
+  {
+    "id": "contract_reload_application",
+    "timestamp": "2025-06-04T02:15:54Z",
+    "type": "contract_definition",
+    "name": "reload_application",
+    "version": "1.0.0-loglineos",
+    "description": "Reloads the entire LogLineOS application in the browser.",
+    "created_by": "LogLineOS Kernel",
+    "args_schema": {},
+    "effects": [
+      { "type": "audit_log", "level": "warn", "message": "Application reload initiated by user action." },
+      { "type": "kernel_action", "action_type": "invoke_native_browser_reload" }
+    ]
+  }
+]
\ No newline at end of file
diff --git a/flipapp_runtime_full/ui/espelho_ui.logline b/flipapp_runtime_full/ui/espelho_ui.logline
new file mode 100644
index 0000000000000000000000000000000000000000..7644c59c7be11e2a319679b52c82476d12c92c2c
--- /dev/null
+++ b/flipapp_runtime_full/ui/espelho_ui.logline
@@ -0,0 +1,463 @@
+# espelho_ui.logline
+# Interface lateral para exibição de spans do LogLine
+# Autor: Dan (LogLineOS)
+# Data: 2025-06-03
+
+- type: container
+  class: "espelho-panel"
+  token:
+    display: "flex"
+    flex-direction: "column"
+    height: "100%"
+    background: "{{token('color.surface')}}"
+    border-right: "1px solid {{token('color.border')}}"
+  children:
+    # Cabeçalho com título e status de conexão
+    - type: container
+      class: "espelho-header"
+      token:
+        padding: "{{token('layout.spacing.md')}}"
+        border-bottom: "1px solid {{token('color.border')}}"
+        display: "flex"
+        justify-content: "space-between"
+        align-items: "center"
+      children:
+        # Título
+        - type: text
+          class: "espelho-title"
+          token:
+            font-size: "{{token('typography.font.size.lg')}}"
+            font-weight: "{{token('typography.font.weight.bold')}}"
+            color: "{{token('color.text')}}"
+          content: "🪞 Espelho"
+        
+        # Indicador de status de conexão
+        - type: container
+          class: "connection-status"
+          token:
+            display: "flex"
+            align-items: "center"
+          children:
+            - type: container
+              class: "status-indicator"
+              token:
+                width: "8px"
+                height: "8px"
+                border-radius: "{{token('layout.radius.circle')}}"
+                background: "{{connection['espelho'] && connection['espelho'].status === 'connected' ? token('color.success') : token('color.alert')}}"
+                margin-right: "{{token('layout.spacing.xs')}}"
+            - type: text
+              class: "status-text"
+              token:
+                font-size: "{{token('typography.font.size.xs')}}"
+                color: "{{token('color.text.secondary')}}"
+              content: "{{connection['espelho'] && connection['espelho'].status === 'connected' ? 'Online' : 'Offline'}}"
+    
+    # Filtros
+    - type: container
+      class: "espelho-filters"
+      token:
+        padding: "{{token('layout.spacing.md')}}"
+        border-bottom: "1px solid {{token('color.border')}}"
+        background: "{{token('color.surface.alt')}}"
+      children:
+        # Campo de busca
+        - type: container
+          class: "search-container"
+          token:
+            margin-bottom: "{{token('layout.spacing.sm')}}"
+            position: "relative"
+          children:
+            - type: input
+              class: "search-input"
+              placeholder: "Buscar em spans..."
+              bind: "espelho.filters.search | updateFilters"
+              token:
+                width: "100%"
+                padding: "{{token('layout.spacing.sm')}}"
+                padding-left: "{{token('layout.spacing.xl')}}"
+                border: "1px solid {{token('color.border')}}"
+                border-radius: "{{token('layout.radius.sm')}}"
+                font-size: "{{token('typography.font.size.sm')}}"
+                background: "{{token('color.surface')}}"
+                color: "{{token('color.text')}}"
+            - type: container
+              class: "search-icon"
+              content: "🔍"
+              token:
+                position: "absolute"
+                left: "{{token('layout.spacing.sm')}}"
+                top: "50%"
+                transform: "translateY(-50%)"
+                font-size: "{{token('typography.font.size.sm')}}"
+                color: "{{token('color.text.secondary')}}"
+        
+        # Filtros de tipo e categoria
+        - type: container
+          class: "filter-section"
+          token:
+            display: "flex"
+            flex-direction: "column"
+            gap: "{{token('layout.spacing.sm')}}"
+          children:
+            # Filtros de tipo
+            - type: container
+              class: "filter-row"
+              token:
+                display: "flex"
+                flex-wrap: "wrap"
+                gap: "{{token('layout.spacing.xs')}}"
+                margin-bottom: "{{token('layout.spacing.xs')}}"
+              children:
+                - type: text
+                  class: "filter-label"
+                  token:
+                    font-size: "{{token('typography.font.size.xs')}}"
+                    font-weight: "{{token('typography.font.weight.medium')}}"
+                    color: "{{token('color.text.secondary')}}"
+                    margin-right: "{{token('layout.spacing.xs')}}"
+                    flex-basis: "100%"
+                  content: "Filtrar por tipo:"
+                
+                # Botões de filtro de tipo dinâmicos
+                - type: loop
+                  data: "getAvailableTypes()"
+                  children:
+                    - type: container
+                      class: "filter-type-button {{isTypeSelected(item) ? 'selected' : ''}}"
+                      on: "click: toggleTypeFilter(type={{item}})"
+                      token:
+                        padding: "{{token('layout.spacing.xs')}} {{token('layout.spacing.sm')}}"
+                        border-radius: "{{token('layout.radius.sm')}}"
+                        font-size: "{{token('typography.font.size.xs')}}"
+                        cursor: "pointer"
+                        background: "{{isTypeSelected(item) ? token('color.primary.light') : token('color.surface')}}"
+                        border: "1px solid {{isTypeSelected(item) ? token('color.primary') : token('color.border')}}"
+                        color: "{{isTypeSelected(item) ? token('color.primary.dark') : token('color.text')}}"
+                        display: "flex"
+                        align-items: "center"
+                        gap: "{{token('layout.spacing.xs')}}"
+                        transition: "{{token('motion.transition.fast')}}"
+                      children:
+                        - type: container
+                          content: "{{getIconForType(item)}}"
+                        - type: text
+                          content: "{{item}}"
+            
+            # Filtros de categoria
+            - type: container
+              class: "filter-row"
+              token:
+                display: "flex"
+                flex-wrap: "wrap"
+                gap: "{{token('layout.spacing.xs')}}"
+              children:
+                - type: text
+                  class: "filter-label"
+                  token:
+                    font-size: "{{token('typography.font.size.xs')}}"
+                    font-weight: "{{token('typography.font.weight.medium')}}"
+                    color: "{{token('color.text.secondary')}}"
+                    margin-right: "{{token('layout.spacing.xs')}}"
+                    flex-basis: "100%"
+                  content: "Filtrar por categoria:"
+                
+                # Botões de filtro de categoria dinâmicos
+                - type: loop
+                  data: "getAvailableCategories()"
+                  children:
+                    - type: container
+                      class: "filter-category-button {{isCategorySelected(item) ? 'selected' : ''}}"
+                      on: "click: toggleCategoryFilter(category={{item}})"
+                      token:
+                        padding: "{{token('layout.spacing.xs')}} {{token('layout.spacing.sm')}}"
+                        border-radius: "{{token('layout.radius.sm')}}"
+                        font-size: "{{token('typography.font.size.xs')}}"
+                        cursor: "pointer"
+                        background: "{{isCategorySelected(item) ? getCategoryColor(item) : token('color.surface')}}"
+                        border: "1px solid {{isCategorySelected(item) ? getCategoryColorBorder(item) : token('color.border')}}"
+                        color: "{{token('color.text')}}"
+                        display: "flex"
+                        align-items: "center"
+                        gap: "{{token('layout.spacing.xs')}}"
+                        transition: "{{token('motion.transition.fast')}}"
+                      children:
+                        - type: container
+                          content: "{{getCategoryIcon(item)}}"
+                        - type: text
+                          content: "{{item}}"
+
+    # Área principal com spans filtrados
+    - type: container
+      class: "espelho-content"
+      token:
+        flex: "1"
+        overflow-y: "auto"
+        padding: "{{token('layout.spacing.md')}}"
+      children:
+        # Estado vazio quando não há spans
+        - type: when
+          when: "!hasFilteredSpans()"
+          children:
+            - type: container
+              class: "empty-state"
+              token:
+                display: "flex"
+                flex-direction: "column"
+                align-items: "center"
+                justify-content: "center"
+                padding: "{{token('layout.spacing.xl')}}"
+                text-align: "center"
+              children:
+                - type: container
+                  class: "empty-icon"
+                  content: "🔍"
+                  token:
+                    font-size: "{{token('typography.font.size.xxl')}}"
+                    margin-bottom: "{{token('layout.spacing.md')}}"
+                    opacity: "0.5"
+                - type: text
+                  class: "empty-title"
+                  token:
+                    font-size: "{{token('typography.font.size.md')}}"
+                    font-weight: "{{token('typography.font.weight.medium')}}"
+                    color: "{{token('color.text')}}"
+                    margin-bottom: "{{token('layout.spacing.sm')}}"
+                  content: "{{isFiltering() ? 'Nenhum span corresponde aos filtros' : 'Nenhum span disponível'}}"
+                - type: text
+                  class: "empty-subtitle"
+                  token:
+                    font-size: "{{token('typography.font.size.sm')}}"
+                    color: "{{token('color.text.secondary')}}"
+                  content: "{{isFiltering() ? 'Tente alterar seus filtros' : connection['espelho'] && connection['espelho'].status !== 'connected' ? 'Verifique sua conexão' : 'Aguardando novos dados'}}"
+                - type: when
+                  when: "isFiltering()"
+                  children:
+                    - type: container
+                      class: "reset-filters-button"
+                      on: "click: resetFilters"
+                      token:
+                        margin-top: "{{token('layout.spacing.md')}}"
+                        padding: "{{token('layout.spacing.sm')}} {{token('layout.spacing.md')}}"
+                        background: "{{token('color.primary')}}"
+                        color: "white"
+                        border-radius: "{{token('layout.radius.sm')}}"
+                        cursor: "pointer"
+                        font-size: "{{token('typography.font.size.sm')}}"
+                      children:
+                        - type: text
+                          content: "Limpar filtros"
+        
+        # Loop para exibir spans filtrados
+        - type: when
+          when: "hasFilteredSpans()"
+          children:
+            - type: loop
+              data: "getFilteredSpans()"
+              children:
+                - type: include
+                  source: "span_card.logline"
+                  props:
+                    item: "{{item}}"
+                    expanded: "{{isSpanExpanded(item.span_id)}}"
+
+# Funções auxiliares para espelho_ui
+- type: function
+  name: getAvailableTypes
+  code: |
+    function() {
+      const spans = stateManager.getState('espelho.spans') || [];
+      const types = new Set();
+      spans.forEach(span => {
+        if (span.type) types.add(span.type);
+      });
+      return Array.from(types);
+    }
+
+- type: function
+  name: getAvailableCategories
+  code: |
+    function() {
+      const spans = stateManager.getState('espelho.spans') || [];
+      const categories = new Set();
+      spans.forEach(span => {
+        if (span.category) categories.add(span.category);
+      });
+      return Array.from(categories);
+    }
+
+- type: function
+  name: getFilteredSpans
+  code: |
+    function() {
+      const spans = stateManager.getState('espelho.spans') || [];
+      const filters = stateManager.getState('espelho.filters') || {};
+      
+      // Aplicar filtro de busca
+      let filtered = spans;
+      
+      if (filters.search) {
+        const searchLower = filters.search.toLowerCase();
+        filtered = filtered.filter(span => 
+          (span.content && span.content.toLowerCase().includes(searchLower)) ||
+          (span.type && span.type.toLowerCase().includes(searchLower)) ||
+          (span.category && span.category.toLowerCase().includes(searchLower))
+        );
+      }
+      
+      // Aplicar filtros de tipo
+      if (filters.types && filters.types.length > 0) {
+        filtered = filtered.filter(span => 
+          span.type && filters.types.includes(span.type)
+        );
+      }
+      
+      // Aplicar filtros de categoria
+      if (filters.categories && filters.categories.length > 0) {
+        filtered = filtered.filter(span => 
+          span.category && filters.categories.includes(span.category)
+        );
+      }
+      
+      // Ordenar por timestamp decrescente
+      return filtered.sort((a, b) => {
+        const dateA = a.timestamp ? new Date(a.timestamp).getTime() : 0;
+        const dateB = b.timestamp ? new Date(b.timestamp).getTime() : 0;
+        return dateB - dateA;
+      });
+    }
+
+- type: function
+  name: isTypeSelected
+  code: |
+    function(type) {
+      const filters = stateManager.getState('espelho.filters') || {};
+      return filters.types && filters.types.includes(type);
+    }
+
+- type: function
+  name: isCategorySelected
+  code: |
+    function(category) {
+      const filters = stateManager.getState('espelho.filters') || {};
+      return filters.categories && filters.categories.includes(category);
+    }
+
+- type: function
+  name: toggleTypeFilter
+  code: |
+    function(type) {
+      const filters = stateManager.getState('espelho.filters') || {};
+      const types = filters.types || [];
+      
+      const index = types.indexOf(type);
+      if (index === -1) {
+        // Adiciona o tipo
+        stateManager.setState('espelho.filters.types', [...types, type]);
+        feedback.vibration(token('feedback.vibration.short'));
+      } else {
+        // Remove o tipo
+        const newTypes = [...types];
+        newTypes.splice(index, 1);
+        stateManager.setState('espelho.filters.types', newTypes);
+      }
+    }
+
+- type: function
+  name: toggleCategoryFilter
+  code: |
+    function(category) {
+      const filters = stateManager.getState('espelho.filters') || {};
+      const categories = filters.categories || [];
+      
+      const index = categories.indexOf(category);
+      if (index === -1) {
+        // Adiciona a categoria
+        stateManager.setState('espelho.filters.categories', [...categories, category]);
+        feedback.vibration(token('feedback.vibration.short'));
+      } else {
+        // Remove a categoria
+        const newCategories = [...categories];
+        newCategories.splice(index, 1);
+        stateManager.setState('espelho.filters.categories', newCategories);
+      }
+    }
+
+- type: function
+  name: resetFilters
+  code: |
+    function() {
+      stateManager.setState('espelho.filters', {});
+      feedback.vibration(token('feedback.vibration.medium'));
+    }
+
+- type: function
+  name: hasFilteredSpans
+  code: |
+    function() {
+      return getFilteredSpans().length > 0;
+    }
+
+- type: function
+  name: isFiltering
+  code: |
+    function() {
+      const filters = stateManager.getState('espelho.filters') || {};
+      return !!(filters.search || (filters.types && filters.types.length) || (filters.categories && filters.categories.length));
+    }
+
+- type: function
+  name: isSpanExpanded
+  code: |
+    function(spanId) {
+      const expandedSpans = stateManager.getState('espelho.expandedSpans') || {};
+      return !!expandedSpans[spanId];
+    }
+
+- type: function
+  name: toggleExpand
+  code: |
+    function(spanId) {
+      if (!spanId) return;
+      
+      const expandedSpans = stateManager.getState('espelho.expandedSpans') || {};
+      const isExpanded = !!expandedSpans[spanId];
+      
+      stateManager.setState(`espelho.expandedSpans.${spanId}`, !isExpanded);
+      feedback.vibration(token('feedback.vibration.short'));
+      
+      if (!isExpanded) {
+        // Play animation or sound when expanding
+        feedback.sound(token('feedback.sound.message'));
+      }
+    }
+
+- type: function
+  name: updateFilters
+  code: |
+    function() {
+      // Debounce para atualizações de filtro de busca
+      if (this._searchTimeout) {
+        clearTimeout(this._searchTimeout);
+      }
+      
+      this._searchTimeout = setTimeout(() => {
+        const searchValue = stateManager.getState('espelho.filters.search');
+        if (searchValue && searchValue.trim().length > 0) {
+          feedback.vibration(token('feedback.vibration.short'));
+        }
+      }, 300);
+    }
+
+- type: function
+  name: getCategoryColorBorder
+  code: |
+    function(category) {
+      const colorMap = {
+        'pessoal': 'color.primary',
+        'trabalho': 'color.success',
+        'conta corrente': 'color.info',
+        'estudos': 'color.warning'
+      };
+      return token(colorMap[category.toLowerCase()] || 'color.border');
+    }
\ No newline at end of file
diff --git a/flipapp_runtime_full/ui/flipapp_ui.logline b/flipapp_runtime_full/ui/flipapp_ui.logline
new file mode 100644
index 0000000000000000000000000000000000000000..07abc742e617f7114c1ceb9bb6e8a0f36d95c0ea
--- /dev/null
+++ b/flipapp_runtime_full/ui/flipapp_ui.logline
@@ -0,0 +1,267 @@
+# flipapp_ui.logline
+# Interface principal do FlipApp com layout dual e navegação simbólica
+# Autor: Dan (LogLineOS)
+# Data: 2025-06-03
+
+- type: container
+  class: "flipapp-root"
+  token:
+    height: "100vh"
+    display: "flex"
+    flex-direction: "column"
+    background: "{{token('color.background')}}"
+    font-family: "{{token('typography.font.family')}}"
+    color: "{{token('color.text')}}"
+  children:
+    # Header principal
+    - type: container
+      class: "flipapp-header"
+      token:
+        background: "{{token('color.primary')}}"
+        color: "white"
+        padding: "{{token('layout.spacing.md')}}"
+        box-shadow: "{{token('shadow.base')}}"
+        z-index: "10"
+      children:
+        - type: container
+          class: "header-content"
+          token:
+            display: "flex"
+            justify-content: "space-between"
+            align-items: "center"
+            max-width: "{{token('layout.container.max-width')}}"
+            margin: "0 auto"
+            width: "100%"
+          children:
+            # Logo e título
+            - type: container
+              class: "app-branding"
+              on: "click: navigateHome"
+              token:
+                display: "flex"
+                align-items: "center"
+                cursor: "pointer"
+              children:
+                - type: container
+                  class: "app-logo"
+                  content: "🤖"
+                  token:
+                    font-size: "{{token('typography.font.size.xl')}}"
+                    margin-right: "{{token('layout.spacing.sm')}}"
+                - type: text
+                  class: "app-title"
+                  content: "FlipApp"
+                  token:
+                    font-weight: "{{token('typography.font.weight.bold')}}"
+                    font-size: "{{token('typography.font.size.lg')}}"
+            
+            # Indicadores de status de conexão
+            - type: container
+              class: "connection-status-group"
+              token:
+                display: "flex"
+                gap: "{{token('layout.spacing.md')}}"
+              children:
+                # Status do Espelho
+                - type: container
+                  class: "connection-badge"
+                  token:
+                    display: "flex"
+                    align-items: "center"
+                    gap: "{{token('layout.spacing.xs')}}"
+                    padding: "{{token('layout.spacing.xs')}} {{token('layout.spacing.sm')}}"
+                    border-radius: "{{token('layout.radius.sm')}}"
+                    background: "{{connection['espelho'] && connection['espelho'].status === 'connected' ? 'rgba(255,255,255,0.2)' : 'rgba(255,255,255,0.1)'}}"
+                    font-size: "{{token('typography.font.size.xs')}}"
+                  children:
+                    - type: container
+                      class: "status-dot"
+                      token:
+                        width: "8px"
+                        height: "8px"
+                        border-radius: "{{token('layout.radius.circle')}}"
+                        background: "{{connection['espelho'] && connection['espelho'].status === 'connected' ? token('color.success') : token('color.alert')}}"
+                    - type: text
+                      content: "Espelho"
+                
+                # Status do WhatsApp
+                - type: container
+                  class: "connection-badge"
+                  token:
+                    display: "flex"
+                    align-items: "center"
+                    gap: "{{token('layout.spacing.xs')}}"
+                    padding: "{{token('layout.spacing.xs')}} {{token('layout.spacing.sm')}}"
+                    border-radius: "{{token('layout.radius.sm')}}"
+                    background: "{{connection['whatsapp'] && connection['whatsapp'].status === 'connected' ? 'rgba(255,255,255,0.2)' : 'rgba(255,255,255,0.1)'}}"
+                    font-size: "{{token('typography.font.size.xs')}}"
+                  children:
+                    - type: container
+                      class: "status-dot"
+                      token:
+                        width: "8px"
+                        height: "8px"
+                        border-radius: "{{token('layout.radius.circle')}}"
+                        background: "{{connection['whatsapp'] && connection['whatsapp'].status === 'connected' ? token('color.success') : token('color.alert')}}"
+                    - type: text
+                      content: "WhatsApp"
+    
+    # Navegação por tabs
+    - type: container
+      class: "nav-tabs"
+      token:
+        display: "flex"
+        background: "{{token('color.surface')}}"
+        border-bottom: "1px solid {{token('color.border')}}"
+        overflow-x: "auto"
+        padding: "0 {{token('layout.spacing.md')}}"
+        scroll-behavior: "smooth"
+      children:
+        # Container para centralizar tabs
+        - type: container
+          token:
+            display: "flex"
+            max-width: "{{token('layout.container.max-width')}}"
+            margin: "0 auto"
+            width: "100%"
+          children:
+            # Tab: Chat
+            - type: container
+              class: "nav-tab {{activeTab === 'chat' ? 'active' : ''}}"
+              on: "click: switchTab(tabId='chat')"
+              token:
+                padding: "{{token('layout.spacing.md')}}"
+                cursor: "pointer"
+                border-bottom: "2px solid {{activeTab === 'chat' ? token('color.primary') : 'transparent'}}"
+                transition: "{{token('motion.transition.fast')}}"
+                color: "{{activeTab === 'chat' ? token('color.primary') : token('color.text')}}"
+                display: "flex"
+                align-items: "center"
+                gap: "{{token('layout.spacing.xs')}}"
+              children:
+                - type: container
+                  content: "💬"
+                - type: text
+                  content: "Chat"
+            
+            # Tab: Espelho
+            - type: container
+              class: "nav-tab {{activeTab === 'espelho' ? 'active' : ''}}"
+              on: "click: switchTab(tabId='espelho')"
+              token:
+                padding: "{{token('layout.spacing.md')}}"
+                cursor: "pointer"
+                border-bottom: "2px solid {{activeTab === 'espelho' ? token('color.primary') : 'transparent'}}"
+                transition: "{{token('motion.transition.fast')}}"
+                color: "{{activeTab === 'espelho' ? token('color.primary') : token('color.text')}}"
+                display: "flex"
+                align-items: "center"
+                gap: "{{token('layout.spacing.xs')}}"
+              children:
+                - type: container
+                  content: "🪞"
+                - type: text
+                  content: "Espelho"
+            
+            # Tab: WhatsApp
+            - type: container
+              class: "nav-tab {{activeTab === 'whatsapp' ? 'active' : ''}}"
+              on: "click: switchTab(tabId='whatsapp')"
+              token:
+                padding: "{{token('layout.spacing.md')}}"
+                cursor: "pointer"
+                border-bottom: "2px solid {{activeTab === 'whatsapp' ? token('color.primary') : 'transparent'}}"
+                transition: "{{token('motion.transition.fast')}}"
+                color: "{{activeTab === 'whatsapp' ? token('color.primary') : token('color.text')}}"
+                display: "flex"
+                align-items: "center"
+                gap: "{{token('layout.spacing.xs')}}"
+              children:
+                - type: container
+                  content: "📱"
+                - type: text
+                  content: "WhatsApp"
+            
+            # Tab: New
+            - type: container
+              class: "nav-tab {{activeTab === 'new' ? 'active' : ''}}"
+              on: "click: switchTab(tabId='new')"
+              token:
+                padding: "{{token('layout.spacing.md')}}"
+                cursor: "pointer"
+                border-bottom: "2px solid {{activeTab === 'new' ? token('color.primary') : 'transparent'}}"
+                transition: "{{token('motion.transition.fast')}}"
+                color: "{{activeTab === 'new' ? token('color.primary') : token('color.text')}}"
+                display: "flex"
+                align-items: "center"
+                gap: "{{token('layout.spacing.xs')}}"
+              children:
+                - type: container
+                  content: "✨"
+                - type: text
+                  content: "Novo"
+    
+    # Layout principal com conteúdo
+    - type: container
+      class: "main-content"
+      token:
+        display: "flex"
+        flex: "1"
+        overflow: "hidden"
+      children:
+        # Sidebar: Espelho (sempre visível em telas grandes)
+        - type: container
+          class: "sidebar-espelho"
+          token:
+            width: "300px"
+            flex-shrink: "0"
+            display: "{{isDesktop() ? 'flex' : 'none'}}"
+            flex-direction: "column"
+            border-right: "1px solid {{token('color.border')}}"
+            background: "{{token('color.surface')}}"
+            overflow: "hidden"
+            "media[{{token('layout.breakpoint.lg')}}]":
+              display: "flex"
+          children:
+            - type: include
+              source: "espelho_ui.logline"
+              props: {}
+        
+        # Conteúdo principal (direito) baseado na tab ativa
+        - type: container
+          class: "main-panel"
+          token:
+            flex: "1"
+            display: "flex"
+            flex-direction: "column"
+            overflow: "hidden"
+          children:
+            # Tab: Chat
+            - type: when
+              when: "activeTab === 'chat'"
+              children:
+                - type: include
+                  source: "chat_ui.logline"
+                  props: {}
+            
+            # Tab: Espelho (visível apenas em telas pequenas)
+            - type: when
+              when: "activeTab === 'espelho' && !isDesktop()"
+              children:
+                - type: include
+                  source: "espelho_ui.logline"
+                  props: {}
+            
+            # Tab: Espelho (placeholder em telas grandes)
+            - type: when
+              when: "activeTab === 'espelho' && isDesktop()"
+              children:
+                - type: container
+                  class: "info-placeholder"
+                  token:
+                    display: "flex"
+                    flex-direction: "column"
+                    align-items: "center"
+                    justify-content: "center"
+                    height: "100%"
+                    padding: "{{token('layout.spacing.xl')}}
\ No newline at end of file
diff --git a/flipapp_runtime_full/ui/span_card.logline b/flipapp_runtime_full/ui/span_card.logline
new file mode 100644
index 0000000000000000000000000000000000000000..3edd05859113e7e80196434618dc1474068e7463
--- /dev/null
+++ b/flipapp_runtime_full/ui/span_card.logline
@@ -0,0 +1,206 @@
+# span_card.logline
+# Componente visual canônico para exibição de spans
+# Autor: Dan (LogLineOS)
+# Data: 2025-06-03
+
+- type: container
+  class: "span-card"
+  token: 
+    background: "{{category ? 'color.surface' : 'color.surface.alt'}}"
+    border: "1px solid {{token('color.border')}}"
+    radius: "{{token('layout.radius.md')}}"
+    shadow: "{{token('shadow.base')}}"
+    padding: "{{token('layout.spacing.md')}}"
+    margin-bottom: "{{token('layout.spacing.sm')}}"
+    transition: "{{token('motion.transition.normal')}}"
+  animation: "{{token('motion.animation.fade-in')}}"
+  bind: "expanded | expandContract"
+  on: "click: toggleExpand(span_id={{item.span_id || 'null'}})"
+  children:
+    # Cabeçalho do span com ícone de tipo, categoria e timestamp
+    - type: container
+      class: "span-card-header"
+      token:
+        display: "flex"
+        justify-content: "space-between"
+        align-items: "center"
+        margin-bottom: "{{expanded ? token('layout.spacing.sm') : '0'}}"
+      children:
+        # Ícone e tipo
+        - type: container
+          token:
+            display: "flex"
+            align-items: "center"
+          children:
+            - type: container
+              class: "span-type-icon"
+              token:
+                font-size: "{{token('typography.font.size.lg')}}"
+                margin-right: "{{token('layout.spacing.sm')}}"
+              content: "{{getIconForType(item.type)}}"
+            - type: text
+              class: "span-type"
+              token:
+                font-weight: "{{token('typography.font.weight.medium')}}"
+                color: "{{token('color.text')}}"
+                margin-right: "{{token('layout.spacing.md')}}"
+              content: "{{item.type || 'span'}}"
+            # Badge de categoria se existir
+            - type: when
+              when: "item.category"
+              children:
+                - type: container
+                  class: "span-category-badge"
+                  token:
+                    display: "inline-flex"
+                    align-items: "center"
+                    background: "{{getCategoryColor(item.category)}}"
+                    padding: "{{token('layout.spacing.xs')}} {{token('layout.spacing.sm')}}"
+                    border-radius: "{{token('layout.radius.sm')}}"
+                    font-size: "{{token('typography.font.size.xs')}}"
+                  children:
+                    - type: container
+                      token:
+                        margin-right: "{{token('layout.spacing.xs')}}"
+                      content: "{{getCategoryIcon(item.category)}}"
+                    - type: text
+                      token:
+                        color: "{{token('color.text')}}"
+                      content: "{{item.category}}"
+        # Timestamp
+        - type: text
+          class: "span-timestamp"
+          token:
+            font-size: "{{token('typography.font.size.xs')}}"
+            color: "{{token('color.text.secondary')}}"
+          content: "{{formatTimestamp(item.timestamp)}}"
+
+    # Conteúdo principal (sempre visível)
+    - type: container
+      class: "span-content-preview"
+      token:
+        margin-top: "{{token('layout.spacing.sm')}}"
+        overflow: "hidden"
+        text-overflow: "ellipsis"
+        white-space: "{{expanded ? 'normal' : 'nowrap'}}"
+      children:
+        - type: text
+          class: "span-content"
+          token:
+            font-size: "{{token('typography.font.size.sm')}}"
+            color: "{{token('color.text')}}"
+            line-height: "{{token('typography.line.normal')}}"
+          content: "{{item.content || 'Sem conteúdo'}}"
+
+    # Detalhes expandidos (visíveis apenas quando expandido)
+    - type: when
+      when: "expanded"
+      children:
+        - type: container
+          class: "span-details"
+          token:
+            margin-top: "{{token('layout.spacing.md')}}"
+            padding-top: "{{token('layout.spacing.sm')}}"
+            border-top: "1px solid {{token('color.border')}}"
+          children:
+            # ID do span
+            - type: container
+              class: "span-detail-row"
+              token:
+                display: "flex"
+                margin-bottom: "{{token('layout.spacing.xs')}}"
+                font-size: "{{token('typography.font.size.xs')}}"
+              children:
+                - type: text
+                  token:
+                    font-weight: "{{token('typography.font.weight.medium')}}"
+                    color: "{{token('color.text.secondary')}}"
+                    margin-right: "{{token('layout.spacing.sm')}}"
+                  content: "ID:"
+                - type: text
+                  class: "span-id"
+                  token:
+                    color: "{{token('color.text.secondary')}}"
+                    font-family: "monospace"
+                  content: "{{item.span_id || 'undefined'}}"
+
+            # Efeitos (se houver)
+            - type: when
+              when: "item.effects && item.effects.length > 0"
+              children:
+                - type: container
+                  class: "span-effects"
+                  token:
+                    margin-top: "{{token('layout.spacing.sm')}}"
+                  children:
+                    - type: text
+                      token:
+                        font-size: "{{token('typography.font.size.xs')}}"
+                        font-weight: "{{token('typography.font.weight.medium')}}"
+                        color: "{{token('color.text.secondary')}}"
+                        margin-bottom: "{{token('layout.spacing.xs')}}"
+                      content: "Efeitos:"
+                    - type: loop
+                      data: "item.effects"
+                      children:
+                        - type: container
+                          class: "span-effect-item"
+                          token:
+                            display: "flex"
+                            align-items: "center"
+                            font-size: "{{token('typography.font.size.xs')}}"
+                            padding: "{{token('layout.spacing.xs')}}"
+                            margin-bottom: "{{token('layout.spacing.xs')}}"
+                            background: "{{token('color.surface.alt')}}"
+                            border-radius: "{{token('layout.radius.xs')}}"
+                          children:
+                            - type: text
+                              token:
+                                color: "{{token('color.info')}}"
+                              content: "→ "
+                            - type: text
+                              content: "{{item.type}}: {{item.target || 'sistema'}}"
+
+# Funções auxiliares para o componente
+- type: function
+  name: getIconForType
+  code: |
+    function(type) {
+      if (!type) return token('icons.type.default');
+      return token('icons.type.' + type.toLowerCase()) || token('icons.type.default');
+    }
+
+- type: function
+  name: getCategoryIcon
+  code: |
+    function(category) {
+      if (!category) return token('icons.category.default');
+      return token('icons.category.' + category.toLowerCase()) || token('icons.category.default');
+    }
+
+- type: function
+  name: getCategoryColor
+  code: |
+    function(category) {
+      const colorMap = {
+        'pessoal': 'color.primary.light',
+        'trabalho': 'color.success',
+        'conta corrente': 'color.info',
+        'estudos': 'color.warning'
+      };
+      const baseColor = colorMap[category.toLowerCase()] || 'color.surface.alt';
+      return `${token(baseColor)}40`; // Adiciona transparência
+    }
+
+- type: function
+  name: formatTimestamp
+  code: |
+    function(timestamp) {
+      if (!timestamp) return '';
+      try {
+        const date = new Date(timestamp);
+        return date.toLocaleString();
+      } catch (e) {
+        return timestamp;
+      }
+    }
\ No newline at end of file
diff --git a/flipapp_runtime_full/ui/tokens_ui.logline b/flipapp_runtime_full/ui/tokens_ui.logline
new file mode 100644
index 0000000000000000000000000000000000000000..1714d693576dd4b0f47ab46c1bcecc1ce6e5bf07
--- /dev/null
+++ b/flipapp_runtime_full/ui/tokens_ui.logline
@@ -0,0 +1,191 @@
+# tokens_ui.logline
+# Design tokens canônicos para identidade visual do FlipApp
+# Autor: Dan (LogLineOS)
+# Data: 2025-06-03
+
+- type: token
+  token: theme
+  values:
+    light:
+      color:
+        primary: "#1a8cff"
+        primary.light: "#e3f2fd"
+        primary.dark: "#0066cc"
+        surface: "#ffffff"
+        surface.alt: "#f8f9fa"
+        background: "#f0f4f8"
+        text: "#212121"
+        text.secondary: "#757575"
+        border: "#e0e0e0"
+        bot: "#f0f4c3"
+        user: "#e3f2fd"
+        alert: "#f44336"
+        success: "#4caf50"
+        warning: "#ff9800"
+        info: "#2196f3"
+      shadow:
+        base: "0 2px 5px rgba(0,0,0,0.08)"
+        elevated: "0 4px 10px rgba(0,0,0,0.12)"
+        focus: "0 0 0 2px rgba(26,140,255,0.4)"
+    dark:
+      color:
+        primary: "#4db8ff"
+        primary.light: "#1c313a"
+        primary.dark: "#007acc"
+        surface: "#1e1e1e"
+        surface.alt: "#252525"
+        background: "#121212"
+        text: "#ffffff"
+        text.secondary: "#aaaaaa"
+        border: "#333333"
+        bot: "#3a3c2e"
+        user: "#1c313a"
+        alert: "#f44336"
+        success: "#5cb860"
+        warning: "#ffb74d"
+        info: "#64b5f6"
+      shadow:
+        base: "0 2px 5px rgba(0,0,0,0.2)"
+        elevated: "0 4px 10px rgba(0,0,0,0.4)"
+        focus: "0 0 0 2px rgba(77,184,255,0.4)"
+    voulezvous:
+      color:
+        primary: "#9c27b0"
+        primary.light: "#f3e5f5"
+        primary.dark: "#6a0080"
+        surface: "#ffffff"
+        surface.alt: "#faf8fc"
+        background: "#f5f0f7"
+        text: "#37294d"
+        text.secondary: "#7e6b98"
+        border: "#e6dbef"
+        bot: "#ede7f6"
+        user: "#f3e5f5"
+        alert: "#e91e63"
+        success: "#66bb6a"
+        warning: "#ffab40"
+        info: "#7e57c2"
+      shadow:
+        base: "0 2px 5px rgba(156,39,176,0.1)"
+        elevated: "0 4px 10px rgba(156,39,176,0.2)"
+        focus: "0 0 0 2px rgba(156,39,176,0.3)"
+    cam4:
+      color:
+        primary: "#ff5722"
+        primary.light: "#ffccbc"
+        primary.dark: "#c41c00"
+        surface: "#121212"
+        surface.alt: "#1d1d1d"
+        background: "#0a0a0a"
+        text: "#ffffff"
+        text.secondary: "#b3b3b3"
+        border: "#2c2c2c"
+        bot: "#331400"
+        user: "#541f00"
+        alert: "#ff1744"
+        success: "#00e676"
+        warning: "#ffab00"
+        info: "#00b0ff"
+      shadow:
+        base: "0 2px 5px rgba(255,87,34,0.2)"
+        elevated: "0 4px 10px rgba(255,87,34,0.4)"
+        focus: "0 0 0 2px rgba(255,87,34,0.5)"
+
+- type: token
+  token: typography
+  values:
+    font:
+      family: "system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif"
+      size:
+        xs: "0.75rem"
+        sm: "0.875rem"
+        md: "1rem"
+        lg: "1.25rem"
+        xl: "1.5rem"
+        xxl: "2rem"
+      weight:
+        regular: "400"
+        medium: "500"
+        bold: "700"
+      line:
+        tight: "1.2"
+        normal: "1.5"
+        loose: "1.8"
+
+- type: token
+  token: layout
+  values:
+    spacing:
+      xs: "0.25rem"
+      sm: "0.5rem"
+      md: "1rem"
+      lg: "1.5rem"
+      xl: "2rem"
+      xxl: "3rem"
+    radius:
+      xs: "2px"
+      sm: "4px"
+      md: "8px"
+      lg: "16px"
+      circle: "50%"
+    breakpoint:
+      sm: "640px"
+      md: "768px"
+      lg: "1024px"
+      xl: "1280px"
+    container:
+      padding: "1rem"
+      max-width: "1200px"
+
+- type: token
+  token: motion
+  values:
+    animation:
+      fade-in: "opacity 0.3s ease-in-out"
+      slide-in: "transform 0.3s ease-out"
+      scale: "transform 0.2s ease-in-out"
+      typing: "ellipsis 1.5s infinite"
+    transition:
+      fast: "0.15s ease"
+      normal: "0.3s ease"
+      slow: "0.5s ease"
+    keyframes:
+      pulse: [{ opacity: 0.6 }, { opacity: 1 }, { opacity: 0.6 }]
+      ellipsis: [{ width: 0 }, { width: "0.5em" }, { width: "1em" }, { width: "1.5em" }]
+      bounce: [{ transform: "translateY(0)" }, { transform: "translateY(-5px)" }, { transform: "translateY(0)" }]
+
+- type: token
+  token: feedback
+  values:
+    vibration:
+      short: "10ms"
+      medium: "50ms"
+      long: "100ms"
+      pattern: "10ms 30ms 10ms"
+    sound:
+      message: "/sounds/message.mp3"
+      notification: "/sounds/notification.mp3"
+      success: "/sounds/success.mp3"
+      error: "/sounds/error.mp3"
+
+- type: token
+  token: icons
+  values:
+    type:
+      register: "📝"
+      commit: "✅"
+      event: "🔔"
+      alert: "⚠️"
+      message: "💬"
+      log: "📋"
+      action: "⚡"
+      payment: "💰"
+      transfer: "🔄"
+      default: "📄"
+    category:
+      pessoal: "👤"
+      trabalho: "💼"
+      "conta corrente": "🏦"
+      estudos: "📚"
+      saude: "❤️"
+      default: "📁"
\ No newline at end of file
diff --git a/flipapp_runtime_full/ui/whatsapp_styles.logline b/flipapp_runtime_full/ui/whatsapp_styles.logline
new file mode 100644
index 0000000000000000000000000000000000000000..9d94afa94d68cbd22432d2b5978e40b1bb77e50b
--- /dev/null
+++ b/flipapp_runtime_full/ui/whatsapp_styles.logline
@@ -0,0 +1,32 @@
+# Design tokens para identidade visual do WhatsApp Web
+# Autor: danvoulez
+# Data: 2025-06-04
+
+- type: token
+  token: whatsapp_theme
+  values:
+    default:
+      color:
+        background: "#ECE5DD"
+        sidebar_bg: "#111b21"
+        sidebar_text: "#ffffff"
+        header_bg: "#128C7E"
+        header_text: "#ffffff"
+        bubble_incoming: "#FFFFFF"
+        bubble_outgoing: "#DCF8C6"
+        text_primary: "#303030"
+        text_secondary: "#717171"
+        accent: "#34B7F1"
+        border: "#E0E0E0"
+      layout:
+        sidebar_width: "300px"
+        header_height: "56px"
+        bubble_radius: "8px"
+        spacing_md: "1rem"
+      typography:
+        font_family: "Helvetica, Arial, sans-serif"
+        font_size_base: "14px"
+        font_size_sm: "12px"
+      shadow:
+        bubble: "0 1px 2px rgba(0,0,0,0.1)"
+        sidebar: "0 2px 4px rgba(0,0,0,0.2)"
\ No newline at end of file
diff --git a/flipapp_runtime_full/ui/whatsapp_ui.logline b/flipapp_runtime_full/ui/whatsapp_ui.logline
new file mode 100644
index 0000000000000000000000000000000000000000..83f87da65c2fea016b78d8f58bcf8be6ff4d57af
--- /dev/null
+++ b/flipapp_runtime_full/ui/whatsapp_ui.logline
@@ -0,0 +1,82 @@
+# Interface principal do módulo WhatsApp Web
+# Autor: danvoulez
+# Data: 2025-06-04
+
+- type: container
+  class: "whatsapp-root"
+  token: 
+    display: "flex"
+    height: "100vh"
+    font-family: "{{token('whatsapp_theme.typography.font_family')}}"
+    background: "{{token('whatsapp_theme.color.background')}}"
+  children:
+    # Sidebar de conversas
+    - type: container
+      class: "whatsapp-sidebar"
+      token:
+        width: "{{token('whatsapp_theme.layout.sidebar_width')}}"
+        background: "{{token('whatsapp_theme.color.sidebar_bg')}}"
+        color: "{{token('whatsapp_theme.color.sidebar_text')}}"
+        overflow-y: "auto"
+    # Área de chat
+    - type: container
+      class: "whatsapp-chat-area"
+      token:
+        flex: "1"
+        display: "flex"
+        flex-direction: "column"
+    # Header do chat ativo
+    - type: container
+      class: "whatsapp-header"
+      token:
+        height: "{{token('whatsapp_theme.layout.header_height')}}"
+        background: "{{token('whatsapp_theme.color.header_bg')}}"
+        color: "{{token('whatsapp_theme.color.header_text')}}"
+        display: "flex"
+        align-items: "center"
+        padding: "0 {{token('whatsapp_theme.layout.spacing_md')}}"
+      children:
+        - type: text
+          class: "chat-with-name"
+          content: "{{state.conversations[state.active_conversation].name}}"
+          token:
+            font-size: "16px"
+            font-weight: "bold"
+    # Histórico de mensagens
+    - type: container
+      class: "whatsapp-messages"
+      token:
+        flex: "1"
+        overflow-y: "auto"
+        padding: "{{token('whatsapp_theme.layout.spacing_md')}}"
+      children:
+        - type: loop
+          data: "state.messages[state.active_conversation]"
+          children:
+            - type: ChatMessage
+              sender: "{{item.sender}}"
+              content: "{{item.content}}"
+              timestamp: "{{item.timestamp}}"
+              class: "{{item.sender === 'user' ? 'outgoing' : 'incoming'}}"
+    # Área de digitação
+    - type: container
+      class: "whatsapp-input-area"
+      token:
+        display: "flex"
+        padding: "{{token('whatsapp_theme.layout.spacing_md')}}"
+        border-top: "1px solid {{token('whatsapp_theme.color.border')}}"
+      children:
+        - type: input
+          bind: "state.input_message | userTypingDebounce"
+          placeholder: "Digite uma mensagem..."
+          token:
+            flex: "1"
+            padding: "0.5rem"
+            border: "1px solid {{token('whatsapp_theme.color.border')}}"
+            border-radius: "50px"
+            font-size: "{{token('whatsapp_theme.typography.font_size_base')}}"
+        - type: Button
+          content: "Enviar"
+          variant: "primary"
+          onClick: "sendMessage"
+          disabled: "{{!state.input_message || state.input_message.trim() === ''}}"
\ No newline at end of file
diff --git a/flipapp_runtime_full/vite.config.js b/flipapp_runtime_full/vite.config.js
new file mode 100644
index 0000000000000000000000000000000000000000..06f7435229f5af5429b7ebf8993cbe9ee8099b4a
--- /dev/null
+++ b/flipapp_runtime_full/vite.config.js
@@ -0,0 +1,109 @@
+// vite.config.js
+import { defineConfig, loadEnv } from 'vite';
+import { resolve } from 'path';
+
+export default defineConfig(({ mode }) => {
+  // Carrega variáveis de ambiente com base no modo
+  const env = loadEnv(mode, process.cwd(), '');
+  
+  return {
+    // Define diretório base para build
+    base: '/',
+    
+    // Aliasing para facilitar importações
+    resolve: {
+      alias: {
+        '@': resolve(__dirname, 'src'),
+        '@core': resolve(__dirname, 'src/core'),
+        '@ui': resolve(__dirname, 'src/ui'),
+        '@wasm': resolve(__dirname, 'src/wasm'),
+        '@services': resolve(__dirname, 'src/services'),
+        '@utils': resolve(__dirname, 'src/utils')
+      }
+    },
+    
+    // Define entrada principal
+    build: {
+      outDir: 'dist',
+      sourcemap: mode !== 'production',
+      rollupOptions: {
+        input: {
+          main: resolve(__dirname, 'index.html'),
+          'service-worker': resolve(__dirname, 'src/service-worker.js')
+        },
+        output: {
+          manualChunks: {
+            vendor: ['src/vendor/markdown-it.js', 'src/vendor/dompurify.js'],
+            wasm: ['src/wasm/indexeddb-bridge.js']
+          }
+        }
+      },
+      minify: mode === 'production',
+      target: 'es2018'
+    },
+    
+    // Servidor de desenvolvimento
+    server: {
+      port: 3000,
+      strictPort: true,
+      proxy: {
+        // Proxy requests para APIs e WebSockets
+        '/api': {
+          target: env.API_BASE_URL || 'http://localhost:8000',
+          changeOrigin: true,
+          secure: false
+        },
+        '/ws': {
+          target: env.WS_BASE_URL || 'ws://localhost:8080',
+          ws: true,
+          changeOrigin: true
+        }
+      }
+    },
+    
+    // Plugins para otimização
+    plugins: [
+      // Plugin para substituir variáveis de ambiente
+      {
+        name: 'html-env-vars',
+        transformIndexHtml: {
+          enforce: 'pre',
+          transform(html) {
+            return html.replace(/\{\{([^}]+)\}\}/g, (match, p1) => {
+              return env[p1] || '';
+            });
+          }
+        }
+      },
+      // Helper para Service Worker
+      {
+        name: 'service-worker',
+        apply: 'build',
+        enforce: 'post',
+        generateBundle(options, bundle) {
+          // Adiciona versão de build ao SW para facilitar invalidação
+          const buildVersion = new Date().toISOString().replace(/[^0-9]/g, '');
+          const html = bundle['index.html'];
+          
+          if (html) {
+            html.source = html.source.toString()
+              .replace(/\{\{BUILD_VERSION\}\}/g, buildVersion);
+          }
+        }
+      }
+    ],
+    
+    // Optimizações para produção
+    optimizeDeps: {
+      exclude: ['wasm/rust_vm']
+    },
+    
+    // Variáveis de ambiente expostas ao cliente
+    define: {
+      'process.env.NODE_ENV': JSON.stringify(mode),
+      'process.env.API_BASE_URL': JSON.stringify(env.API_BASE_URL),
+      'process.env.WS_BASE_URL': JSON.stringify(env.WS_BASE_URL),
+      'process.env.FLIPAPP_VERSION': JSON.stringify(env.FLIPAPP_VERSION || '1.0.0'),
+    }
+  };
+});
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/contracts/load_chat_history.logline b/flipapp_runtime_full/whatsapp/contracts/load_chat_history.logline
new file mode 100644
index 0000000000000000000000000000000000000000..c5e8f7391b091a1713e73abf874e8ae9a51716bc
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/contracts/load_chat_history.logline
@@ -0,0 +1,19 @@
+contract:
+  name: "load_chat_history"
+  args:
+    conversation_id: string
+  effects:
+    - type: "http_request"
+      method: "GET"
+      url: "/api/history/{{args.conversation_id}}"
+      on_success:
+        contract: "load_chat_history_success"
+        args_mapping:
+          conversation_id: "args.conversation_id"
+          messages: "response.body"
+      on_failure:
+        - type: "state_mutation"
+          path: "state.error"
+          value: "Failed to load history"
+        - type: "audit_log_error"
+          message: "Failed loading history for {{args.conversation_id}}"
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/contracts/load_chat_history_success.logline b/flipapp_runtime_full/whatsapp/contracts/load_chat_history_success.logline
new file mode 100644
index 0000000000000000000000000000000000000000..8ffad6bb9fe17db2624de0b4466016ef87c9f839
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/contracts/load_chat_history_success.logline
@@ -0,0 +1,11 @@
+contract:
+  name: "load_chat_history_success"
+  args:
+    conversation_id: string
+    messages: list
+  effects:
+    - type: "state_merge"
+      path: "messages.{{args.conversation_id}}"
+      value: "{{args.messages}}"
+    - type: "audit_log_info"
+      message: "Loaded {{args.messages.length}} messages for conversation {{args.conversation_id}}"
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/contracts/mark_message_read.logline b/flipapp_runtime_full/whatsapp/contracts/mark_message_read.logline
new file mode 100644
index 0000000000000000000000000000000000000000..8b385f86ac44b5d450bcd26b69964811942387f5
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/contracts/mark_message_read.logline
@@ -0,0 +1,19 @@
+contract:
+  name: "mark_message_read"
+  args:
+    conversation_id: string
+    message_id: string
+  effects:
+    - type: "state_mutation"
+      path: "messages.{{args.conversation_id}}[?id=='{{args.message_id}}'].status"
+      value: "read"
+    - type: "http_request"
+      method: "POST"
+      url: "/api/mark_read"
+      body:
+        message_id: "{{args.message_id}}"
+    - type: "state_mutation"
+      path: "conversations[{{args.conversation_id}}].unread"
+      value: 0
+    - type: "audit_log_info"
+      message: "✅ Message {{args.message_id}} marked as read"
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/contracts/observe_whatsapp_socket.logline b/flipapp_runtime_full/whatsapp/contracts/observe_whatsapp_socket.logline
new file mode 100644
index 0000000000000000000000000000000000000000..348d93cf84681f07f2908e5993c8faacee034df3
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/contracts/observe_whatsapp_socket.logline
@@ -0,0 +1,24 @@
+contract:
+  name: "observe_whatsapp_socket"
+  args:
+    url: string
+  effects:
+    - type: "websocket_listen"
+      url: "{{args.url}}"
+      on_message:
+        contract: "receive_whatsapp_message"
+        args_mapping:
+          conversation_id: "payload.conversation_id"
+          sender: "payload.sender"
+          content: "payload.content"
+          timestamp: "payload.timestamp"
+          media_url: "payload.media_url"
+      on_error:
+        - type: "state_mutation"
+          path: "state.error"
+          value: "WebSocket connection failed"
+        - type: "audit_log_error"
+          message: "WebSocket error: {{error}}"
+      on_reconnect:
+        - type: "audit_log_info"
+          message: "WebSocket reconnected"
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/contracts/receive_whatsapp_message.logline b/flipapp_runtime_full/whatsapp/contracts/receive_whatsapp_message.logline
new file mode 100644
index 0000000000000000000000000000000000000000..601a602948c70124344be1d0fbec4b3c611e6d60
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/contracts/receive_whatsapp_message.logline
@@ -0,0 +1,29 @@
+contract:
+  name: "receive_whatsapp_message"
+  args:
+    conversation_id: string
+    sender: string
+    content: string
+    timestamp: string
+    media_url: optional string
+  effects:
+    - type: "state_append"
+      path: "messages.{{args.conversation_id}}"
+      value:
+        id: "{{generate_uuid}}"
+        conversation_id: "{{args.conversation_id}}"
+        sender: "{{args.sender}}"
+        content: "{{args.content}}"
+        media_url: "{{args.media_url}}"
+        timestamp: "{{args.timestamp}}"
+        status: "unread"
+    - type: "haptic"
+      pattern: "short"
+    - when:
+        condition: "{{state.active_conversation != args.conversation_id}}"
+        effects:
+          - type: "state_mutation"
+            path: "conversations[{{args.conversation_id}}].unread"
+            value: "{{state.conversations[args.conversation_id].unread + 1}}"
+    - type: "audit_log_info"
+      message: "📥 Message received from {{args.sender}} in conversation {{args.conversation_id}}"
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/contracts/send_media_whatsapp.logline b/flipapp_runtime_full/whatsapp/contracts/send_media_whatsapp.logline
new file mode 100644
index 0000000000000000000000000000000000000000..eb678a1d2326e34bd0be1e66c298882e572c1f45
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/contracts/send_media_whatsapp.logline
@@ -0,0 +1,37 @@
+contract:
+  name: "send_media_whatsapp"
+  args:
+    conversation_id: string
+    media_type: "image|audio|document"
+    file_path: string
+  effects:
+    - type: "http_request"
+      method: "POST"
+      url: "/api/upload_media"
+      body:
+        file: "{{args.file_path}}"
+        type: "{{args.media_type}}"
+      on_success:
+        - type: "state_append"
+          path: "messages.{{args.conversation_id}}"
+          value:
+            id: "{{generate_uuid}}"
+            conversation_id: "{{args.conversation_id}}"
+            sender: "{{state.user_id}}"
+            media_url: "{{response.body.url}}"
+            media_type: "{{args.media_type}}"
+            timestamp: "{{timestamp_now}}"
+            status: "sent"
+        - type: "contract_call"
+          contract: "send_whatsapp_message"
+          args:
+            conversation_id: "{{args.conversation_id}}"
+            content: "{{args.media_type}} sent"
+      on_failure:
+        - type: "state_mutation"
+          path: "state.error"
+          value: "Failed to upload media"
+        - type: "haptic"
+          pattern: "long"
+    - type: "audit_log_info"
+      message: "📤 Media sent to {{args.conversation_id}} (type: {{args.media_type}})"
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/contracts/send_whatsapp_message.logline b/flipapp_runtime_full/whatsapp/contracts/send_whatsapp_message.logline
new file mode 100644
index 0000000000000000000000000000000000000000..4c3b64506ba0c74eb80310bcdd95fe8efa1b9992
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/contracts/send_whatsapp_message.logline
@@ -0,0 +1,34 @@
+contract:
+  name: "send_whatsapp_message"
+  args:
+    conversation_id: string
+  effects:
+    - type: "http_request"
+      method: "POST"
+      url: "/api/send_whatsapp"
+      body:
+        conversation_id: "{{args.conversation_id}}"
+        sender: "{{state.user_id}}"
+        content: "{{state.input_message}}"
+        timestamp: "{{timestamp_now}}"
+      on_success:
+        - type: "state_append"
+          path: "messages.{{args.conversation_id}}"
+          value:
+            id: "{{generate_uuid}}"
+            conversation_id: "{{args.conversation_id}}"
+            sender: "{{state.user_id}}"
+            content: "{{state.input_message}}"
+            timestamp: "{{timestamp_now}}"
+            status: "sent"
+        - type: "state_mutation"
+          path: "state.input_message"
+          value: ""
+      on_failure:
+        - type: "state_mutation"
+          path: "state.error"
+          value: "Failed to send message"
+        - type: "haptic"
+          pattern: "long"
+    - type: "audit_log_info"
+      message: "📤 Message sent to conversation {{args.conversation_id}} by {{state.user_id}}"
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/ui/whatsapp_ui.logline b/flipapp_runtime_full/whatsapp/ui/whatsapp_ui.logline
new file mode 100644
index 0000000000000000000000000000000000000000..36cbcec0c29fddee8c75701a2c4fd803267bfd7f
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/ui/whatsapp_ui.logline
@@ -0,0 +1,150 @@
+ui:
+  - container:
+      id: "whatsapp_root"
+      style:
+        flexDirection: "column"
+        height: "100vh"
+        backgroundColor: "#f0f2f5"
+      children:
+        # Header
+        - container:
+            id: "header"
+            style:
+              padding: "12px"
+              backgroundColor: "#128C7E"
+              alignItems: "center"
+            children:
+              - text:
+                  content: "WhatsApp"
+                  style:
+                    color: "white"
+                    fontSize: "20px"
+                    fontWeight: "bold"
+
+        # Conversation List
+        - container:
+            id: "conversation_list"
+            style:
+              flex: 1
+              overflowY: "scroll"
+            children:
+              - loop:
+                  items: "{{state.conversations}}"
+                  as: "convo"
+                  children:
+                    - container:
+                        onTap: "set_active_conversation(id='{{convo.id}}')"
+                        style:
+                          padding: "12px"
+                          borderBottom: "1px solid #eee"
+                        children:
+                          - text:
+                              content: "{{convo.name}}"
+                          - when:
+                              condition: "{{convo.unread}} > 0"
+                              children:
+                                - text:
+                                    content: "({{convo.unread}})"
+                                    style:
+                                      color: "#128C7E"
+
+        # Messages Panel
+        - container:
+            id: "messages_panel"
+            style:
+              flex: 3
+              display: "{{state.active_conversation ? 'flex' : 'none'}}"
+              flexDirection: "column"
+            children:
+              # Messages History
+              - container:
+                  id: "messages_history"
+                  style:
+                    flex: 1
+                    padding: "10px"
+                    overflowY: "scroll"
+                  children:
+                    - loop:
+                        items: "{{state.messages[state.active_conversation]}}"
+                        as: "msg"
+                        children:
+                          - container:
+                              style:
+                                alignSelf: "{{msg.sender == state.user_id ? 'flex-end' : 'flex-start'}}"
+                                backgroundColor: "{{msg.sender == state.user_id ? '#DCF8C6' : 'white'}}"
+                                borderRadius: "8px"
+                                padding: "8px"
+                                marginBottom: "8px"
+                                maxWidth: "70%"
+                              children:
+                                - text:
+                                    content: "{{msg.sender}}:"
+                                    style:
+                                      fontWeight: "bold"
+                                - text:
+                                    content: "{{msg.content}}"
+                                - when:
+                                    condition: "{{msg.media_url}}"
+                                    children:
+                                      - image:
+                                          src: "{{msg.media_url}}"
+                                          style:
+                                            maxWidth: "200px"
+                                            marginTop: "8px"
+                                - when:
+                                    condition: "{{msg.status == 'unread'}}"
+                                    children:
+                                      - text:
+                                          content: "⬤"
+                                          style:
+                                            color: "#128C7E"
+                                            fontSize: "8px"
+                                            alignSelf: "flex-end"
+
+              # Input Area
+              - container:
+                  id: "input_area"
+                  style:
+                    flexDirection: "row"
+                    padding: "10px"
+                    borderTop: "1px solid #ddd"
+                  children:
+                    - input:
+                        id: "message_input"
+                        bind: "state.input_message"
+                        placeholder: "Type a message..."
+                        style:
+                          flex: 1
+                          padding: "10px"
+                          border: "1px solid #ddd"
+                          borderRadius: "20px"
+                          marginRight: "10px"
+                    - button:
+                        id: "send_button"
+                        content: "📤"
+                        onTap: "send_whatsapp_message(conversation_id='{{state.active_conversation}}')"
+                    - button:
+                        id: "attach_button"
+                        content: "📎"
+                        onTap: "open_media_picker()"
+
+        # Fallback UI
+        - when:
+            condition: "{{state.error}}"
+            children:
+              - container:
+                  style:
+                    position: "absolute"
+                    bottom: "20px"
+                    right: "20px"
+                    backgroundColor: "#ff6b6b"
+                    padding: "10px"
+                    borderRadius: "8px"
+                  children:
+                    - text:
+                        content: "Error: {{state.error}}"
+                        style:
+                          color: "white"
+                    - button:
+                        content: "Retry"
+                        onTap: "reconnect_whatsapp()"
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/whatsapp/config/agent_policy.logline b/flipapp_runtime_full/whatsapp/whatsapp/config/agent_policy.logline
new file mode 100644
index 0000000000000000000000000000000000000000..a9ac427b36659b229ea555f937a71df4cf1470c4
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/whatsapp/config/agent_policy.logline
@@ -0,0 +1,13 @@
+policy default_agent_policy {
+  analysis_model: "local/gemma-2b",
+  prompt_template: """
+    ANALYZE_CONVERSATION:
+    CONTEXT: {{CONTEXT}}
+    POLICIES:
+      - Intervene if risk_score > 7
+      - Suggest if ghost_sale active
+      - Alert if keywords: [urgente, socorro, acidente]
+    OUTPUT_FORMAT: json
+  """,
+  risk_threshold: 7
+}
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/whatsapp/config/capacity_manager.logline b/flipapp_runtime_full/whatsapp/whatsapp/config/capacity_manager.logline
new file mode 100644
index 0000000000000000000000000000000000000000..dee4735016b054b45c6276f379649e52449ec7b2
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/whatsapp/config/capacity_manager.logline
@@ -0,0 +1,18 @@
+contract manage_capacity {
+  continuous every 60s {
+    let load = {
+      active_conv: count(state.active_conversations),
+      daily_msg: state.message_counter
+    }
+    state.system_load.push(load)
+    if load.active_conv > 10 {
+      call raise_alert(
+        level: 8,
+        message: "Capacity exceeded! Active conversations: " + load.active_conv
+      )
+    }
+    if is_midnight() {
+      state.message_counter = 0
+    }
+  }
+}
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/whatsapp/core/agent_intervene.logline b/flipapp_runtime_full/whatsapp/whatsapp/core/agent_intervene.logline
new file mode 100644
index 0000000000000000000000000000000000000000..d6d76c64104e6c7b50517bc5cac1e5b8b8b07f9b
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/whatsapp/core/agent_intervene.logline
@@ -0,0 +1,32 @@
+contract agent_intervene {
+  input {
+    conversation_id: string,
+    context: [message],
+    policy: ref("agent_policy")
+  }
+  output {
+    action: "suggest" | "alert" | "summarize" | "none",
+    content?: string,
+    urgency: int
+  }
+  effect {
+    let policy = state.agent_policies[input.policy]
+    let analysis = llm_call(
+      model: policy.analysis_model,
+      prompt: policy.prompt_template.replace("{{CONTEXT}}", input.context)
+    )
+    output = parse_json(analysis)
+    if output.action == "suggest" {
+      call ghost_sale_track(
+        conversation_id: input.conversation_id,
+        event_type: "incentive"
+      )
+    }
+    state.agent_interventions.push({
+      conversation: input.conversation_id,
+      action: output.action,
+      timestamp: now(),
+      context_hash: hash(input.context)
+    })
+  }
+}
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/whatsapp/core/detect_language.logline b/flipapp_runtime_full/whatsapp/whatsapp/core/detect_language.logline
new file mode 100644
index 0000000000000000000000000000000000000000..83ee0fa1d7c52537998cd2b70448739a689993eb
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/whatsapp/core/detect_language.logline
@@ -0,0 +1,12 @@
+contract detect_language {
+  input { text: string }
+  output { language: string, confidence: float }
+  effect {
+    output = builtin_lang_detect(input.text) 
+      ?: llm_call(
+          model: state.config.llm_model,
+          prompt: "DETECT_LANG:{{input.text}}",
+          format: "json"
+        )
+  }
+}
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/whatsapp/core/ghost_sale_abort.logline b/flipapp_runtime_full/whatsapp/whatsapp/core/ghost_sale_abort.logline
new file mode 100644
index 0000000000000000000000000000000000000000..c3755146fa5885ffc62affd998575e0439a4bc96
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/whatsapp/core/ghost_sale_abort.logline
@@ -0,0 +1,28 @@
+contract ghost_sale_abort {
+  input { 
+    conversation_id: string,
+    reason: "timeout" | "high_hesitation" | "no_response"
+  }
+  effect {
+    let sale = state.ghost_sales[input.conversation_id]
+    if sale {
+      sale.status = "aborted"
+      sale.end_reason = input.reason
+      if input.reason == "no_response" {
+        state.affairs.push({
+          type: "negligence",
+          agent: sale.last_agent,
+          conversation: input.conversation_id,
+          timestamp: now(),
+          penalty: calculate_penalty(sale)
+        })
+        call raise_judgement(
+          type: "fail_sale",
+          target: sale.last_agent,
+          evidence: sale
+        )
+      }
+      audit_log("GHOST_SALE_ABORT", sale)
+    }
+  }
+}
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/whatsapp/core/ghost_sale_init.logline b/flipapp_runtime_full/whatsapp/whatsapp/core/ghost_sale_init.logline
new file mode 100644
index 0000000000000000000000000000000000000000..cf70f57deaba2a18d46b742a55f2731ab058f4ed
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/whatsapp/core/ghost_sale_init.logline
@@ -0,0 +1,18 @@
+contract ghost_sale_init {
+  input { message_id: string }
+  effect {
+    let conv_id = get_conversation(input.message_id)
+    if !state.ghost_sales[conv_id] {
+      state.ghost_sales[conv_id] = {
+        id: generate_id(),
+        status: "initiated",
+        created: now(),
+        last_update: now(),
+        messages: [input.message_id],
+        incentive_log: [],
+        hesitation_count: 0
+      }
+      audit_log("GHOST_SALE_INIT", state.ghost_sales[conv_id])
+    }
+  }
+}
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/whatsapp/core/ghost_sale_track.logline b/flipapp_runtime_full/whatsapp/whatsapp/core/ghost_sale_track.logline
new file mode 100644
index 0000000000000000000000000000000000000000..af5877aa95dc737ddfa1303b4329fecd31257af6
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/whatsapp/core/ghost_sale_track.logline
@@ -0,0 +1,30 @@
+contract ghost_sale_track {
+  input { 
+    conversation_id: string,
+    event_type: "incentive" | "hesitation" | "progress"
+  }
+  effect {
+    let sale = state.ghost_sales[input.conversation_id]
+    if sale {
+      sale.last_update = now()
+      if input.event_type == "incentive" {
+        sale.incentive_log.push({
+          timestamp: now(),
+          agent: current_agent(),
+          method: "auto_suggest"
+        })
+      }
+      if input.event_type == "hesitation" {
+        sale.hesitation_count += 1
+        if sale.hesitation_count > 3 {
+          call ghost_sale_abort(conversation_id: input.conversation_id, reason: "high_hesitation")
+        }
+      }
+      if input.event_type == "progress" {
+        sale.status = "progressing"
+        sale.hesitation_count = 0
+      }
+      audit_log("GHOST_SALE_UPDATE", sale)
+    }
+  }
+}
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/whatsapp/core/raise_alert.logline b/flipapp_runtime_full/whatsapp/whatsapp/core/raise_alert.logline
new file mode 100644
index 0000000000000000000000000000000000000000..77f82ce51c590a3852a04e6fdcce4b137ca7d627
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/whatsapp/core/raise_alert.logline
@@ -0,0 +1,16 @@
+contract raise_alert {
+  input { message: string, level: int }
+  effect {
+    if input.level >= 8 {
+      state.alerts.push({
+        timestamp: now(),
+        message: input.message,
+        level: input.level
+      })
+      call agent_intervene(
+        conversation_id: current_conversation(),
+        context: last_messages(5)
+      )
+    }
+  }
+}
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/whatsapp/core/raise_judgement.logline b/flipapp_runtime_full/whatsapp/whatsapp/core/raise_judgement.logline
new file mode 100644
index 0000000000000000000000000000000000000000..8b8176c4c6c50a77473ab146daaf07a9d2cc192b
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/whatsapp/core/raise_judgement.logline
@@ -0,0 +1,23 @@
+contract raise_judgement {
+  input {
+    type: "fail_sale" | "negligence",
+    target: agent_id,
+    evidence: any
+  }
+  effect {
+    state.judgements.push({
+      id: generate_id(),
+      type: input.type,
+      target: input.target,
+      timestamp: now(),
+      evidence: input.evidence,
+      status: "pending"
+    })
+    if input.type == "negligence" {
+      state.agents[input.target].penalty_score += 10
+      if state.agents[input.target].penalty_score > 50 {
+        suspend_agent(input.target)
+      }
+    }
+  }
+}
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/whatsapp/core/summarize_conversation.logline b/flipapp_runtime_full/whatsapp/whatsapp/core/summarize_conversation.logline
new file mode 100644
index 0000000000000000000000000000000000000000..9dea77a78bcb5995febe62767775bc374406717b
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/whatsapp/core/summarize_conversation.logline
@@ -0,0 +1,14 @@
+contract summarize_conversation {
+  input { conversation_id: string }
+  output { summary: string }
+  effect {
+    let messages = get_last_messages(input.conversation_id, 40)
+    let prompt = "SUMMARIZE_CONVERSATION: " + messages
+    output.summary = llm_call(
+      model: state.config.llm_model,
+      prompt: prompt
+    )
+    state.summaries[input.conversation_id] = output.summary
+    audit_log("CONVERSATION_SUMMARY", {conversation_id: input.conversation_id, summary: output.summary})
+  }
+}
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/whatsapp/core/transcribe_audio.logline b/flipapp_runtime_full/whatsapp/whatsapp/core/transcribe_audio.logline
new file mode 100644
index 0000000000000000000000000000000000000000..882c1be1ba36429750c4fd108711e53c0e71b5ce
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/whatsapp/core/transcribe_audio.logline
@@ -0,0 +1,12 @@
+contract transcribe_audio {
+  input { audio: base64 }
+  output { transcription: string }
+  effect {
+    output.transcription = builtin_speech_to_text(input.audio)
+      ?: http_post(
+          url: "https://api.whisper.local/transcribe",
+          body: input.audio,
+          headers: {"Content-Type": "audio/ogg"}
+        ).text
+  }
+}
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/whatsapp/core/translate_message.logline b/flipapp_runtime_full/whatsapp/whatsapp/core/translate_message.logline
new file mode 100644
index 0000000000000000000000000000000000000000..848e6a8483c339f8d4c98f6db1d0c15c0206fbe0
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/whatsapp/core/translate_message.logline
@@ -0,0 +1,11 @@
+contract translate_message {
+  input { text: string, source: string, target: "pt" }
+  output { translated: string }
+  effect {
+    output.translated = builtin_translate(input.text, input.source, input.target)
+      ?: llm_call(
+          model: state.config.llm_model,
+          prompt: "TRANSLATE:{{input.source}}>{{input.target}}:{{input.text}}"
+        )
+  }
+}
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/whatsapp/observer_bot.logline b/flipapp_runtime_full/whatsapp/whatsapp/observer_bot.logline
new file mode 100644
index 0000000000000000000000000000000000000000..0999f0df6d3de218fd340920f8322e141439ab11
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/whatsapp/observer_bot.logline
@@ -0,0 +1,22 @@
+contract observe_whatsapp_socket {
+  continuous on message_received(msg) {
+    call ghost_sale_init(message_id: msg.id)
+    let lang = call detect_language(text: msg.text)
+    let translated = call translate_message({
+      text: msg.text,
+      source: lang.language,
+      target: "pt"
+    })
+    if msg.media_type == "audio" {
+      msg.transcription = call transcribe_audio(msg.content)
+    }
+    // Detecção de emergência
+    let emergency = ["urgente", "socorro", "acidente"].some(word => msg.text.includes(word))
+    if emergency {
+      call raise_alert({ message: msg.text, level: 9 })
+    }
+    if unread_count(msg.conversation_id) > 20 {
+      call summarize_conversation({ conversation_id: msg.conversation_id })
+    }
+  }
+}
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/whatsapp/state.logline b/flipapp_runtime_full/whatsapp/whatsapp/state.logline
new file mode 100644
index 0000000000000000000000000000000000000000..72723d9896fb1345fc6f640e04bff24e1c5459c0
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/whatsapp/state.logline
@@ -0,0 +1,16 @@
+state {
+  ghost_sales: map<conversation_id, ghost_sale> = {}
+  affairs: []affair = []
+  judgements: []judgement = []
+  agent_interventions: []intervention = []
+  system_load: []load_metric = []
+  message_counter: int = 0
+  agents: map<agent_id, agent> = {
+    "default": { penalty_score: 0, status: "active" }
+  }
+  config: {
+    llm_model: "local/gemma-2b",
+    max_conversations: 10,
+    daily_msg_limit: 150
+  }
+}
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/whatsapp/tests.logline b/flipapp_runtime_full/whatsapp/whatsapp/tests.logline
new file mode 100644
index 0000000000000000000000000000000000000000..530c3f53192be4b3ade014b6b50bc2fb7ea9f0f4
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/whatsapp/tests.logline
@@ -0,0 +1,55 @@
+test "Ghost sale abandonment triggers negligence affair" {
+  simulate websocket_event(
+    type: "message",
+    data: {text: "Quero comprar urgente!", sender: "+551199999999"}
+  )
+  advance_time(24h)
+  assert state.ghost_sales["conv_123"].status == "aborted"
+  assert state.affairs[0].type == "negligence"
+  assert state.judgements[0].type == "fail_sale"
+}
+
+test "Agent intervention on high risk" {
+  simulate websocket_event(
+    type: "message",
+    data: {text: "Socorro! Acidente grave!", sender: "+551188888888"}
+  )
+  assert called_contracts.includes("agent_intervene")
+  assert last_output("agent_intervene").urgency >= 9
+}
+
+test "Capacity limit enforcement" {
+  for i in 1..15 {
+    state.active_conversations.push("conv_" + i)
+  }
+  call manage_capacity()
+  assert state.alerts.last().level == 8
+}
+
+test "Translated message appears" {
+  simulate websocket_event(
+    type: "message",
+    data: {text: "Hello, I want to buy!", sender: "+5511999990000"}
+  )
+  let msg = state.conversations["conv_123"].messages.last()
+  assert msg.translated != undefined
+}
+
+test "Feedback updates model" {
+  let sample = {
+    message_id: "msg_123",
+    tag: "translation_error"
+  }
+  call log_user_feedback(sample)
+  call train_llm(state.training_samples.last(10))
+  assert state.config.llm_model.version > "1.2"
+}
+
+test "Emergency message triggers alert and summary" {
+  simulate websocket_event(
+    type: "message",
+    data: {text: "Acidente! Preciso de ajuda urgente", sender: "+5511999990000"}
+  )
+  assert state.alerts.last().level == 9
+  assert state.summaries["conv_123"] != undefined
+}
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/whatsapp/training/log_user_feedback.logline b/flipapp_runtime_full/whatsapp/whatsapp/training/log_user_feedback.logline
new file mode 100644
index 0000000000000000000000000000000000000000..50df1f5f430757b0d69791c1d52123848018e72f
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/whatsapp/training/log_user_feedback.logline
@@ -0,0 +1,11 @@
+contract log_user_feedback {
+  input { message_id: string, tag: string }
+  effect {
+    state.training_samples.push({
+      id: input.message_id,
+      tag: input.tag,
+      timestamp: now(),
+      conversation: current_context()
+    })
+  }
+}
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/whatsapp/training/prepare_feedback_data.logline b/flipapp_runtime_full/whatsapp/whatsapp/training/prepare_feedback_data.logline
new file mode 100644
index 0000000000000000000000000000000000000000..028fabe1e6436b58d5d5e7475a956c6e8f878f73
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/whatsapp/training/prepare_feedback_data.logline
@@ -0,0 +1,9 @@
+contract prepare_feedback_data {
+  output { data: array }
+  effect {
+    output.data = state.training_samples.map(s => ({
+      input: s,
+      context: get_message_context(s.message_id)
+    }))
+  }
+}
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/whatsapp/training/train_llm.logline b/flipapp_runtime_full/whatsapp/whatsapp/training/train_llm.logline
new file mode 100644
index 0000000000000000000000000000000000000000..774d9c87fbfb3138827ed3c53eb8187a93d4ab69
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/whatsapp/training/train_llm.logline
@@ -0,0 +1,12 @@
+contract train_llm {
+  input { samples: [feedback] }
+  output { model_checkpoint: string }
+  effect {
+    let dataset = call prepare_feedback_data(input.samples)
+    output.model_checkpoint = builtin_finetune(
+      base_model: state.config.llm_model,
+      data: dataset
+    )
+    state.config.llm_model = output.model_checkpoint
+  }
+}
\ No newline at end of file
diff --git a/flipapp_runtime_full/whatsapp/whatsapp/ui.logline b/flipapp_runtime_full/whatsapp/whatsapp/ui.logline
new file mode 100644
index 0000000000000000000000000000000000000000..f8bf95ac57a96b94f97f074c4870bdcfeef675c8
--- /dev/null
+++ b/flipapp_runtime_full/whatsapp/whatsapp/ui.logline
@@ -0,0 +1,48 @@
+ui component conversation_item {
+  props: conversation
+  view {
+    card {
+      if state.ghost_sales[props.id]?.status == "initiated" {
+        badge(color: "blue", icon: "ghost", text: "Venda Fantasma Ativa")
+      }
+      if state.ghost_sales[props.id]?.status == "aborted" {
+        badge(color: "red", icon: "warning", text: "ABANDONO: " + state.ghost_sales[props.id].end_reason)
+      }
+      for message in props.messages {
+        component whatsapp_message(message: message)
+      }
+      button_group {
+        button("Venda Concluída", on_click: { call ghost_sale_commit(props.id) })
+        button("Ignorado", on_click: { call log_user_feedback(props.id, "ignored") })
+      }
+    }
+  }
+}
+
+ui component whatsapp_message {
+  props: message
+  state { show_original: false }
+  view {
+    card {
+      toggle_button(
+        label: state.show_original ? "Ver Tradução" : "Ver Original",
+        on_click: { state.show_original = !state.show_original }
+      )
+      text_block(
+        content: state.show_original 
+          ? props.message.original 
+          : props.message.translated,
+        highlight: props.from_agent ? "verde_escuro" : "padrao"
+      )
+      if props.message.transcription {
+        expandable_section("Transcrição") {
+          text_block(props.message.transcription)
+        }
+      }
+      button_group {
+        button("Boa resposta", on_click: call log_user_feedback(props.id, "positive"))
+        button("Tradução errada", on_click: call log_user_feedback(props.id, "translation_error"))
+      }
+    }
+  }
+}
\ No newline at end of file
diff --git a/terreno_completo_FINAL/.terreno_version b/terreno_completo_FINAL/.terreno_version
new file mode 100644
index 0000000000000000000000000000000000000000..ea32fced6fd5a3a11c8d37a46450fc8a9651da4c
--- /dev/null
+++ b/terreno_completo_FINAL/.terreno_version
@@ -0,0 +1 @@
+v1.0.0-terreno
\ No newline at end of file
diff --git a/terreno_completo_FINAL/README.md b/terreno_completo_FINAL/README.md
new file mode 100644
index 0000000000000000000000000000000000000000..32fe55ac5035aeeced71f5a7bf8ff66956f214f5
--- /dev/null
+++ b/terreno_completo_FINAL/README.md
@@ -0,0 +1,29 @@
+# TERRENO LogLineOS
+
+Este pacote contém a implementação simbólica completa do sistema operacional declarativo LogLineOS,
+incluindo:
+
+- Executor LLM com cache auditável
+- Autenticação simbólica via flat-file
+- Frontend simbólico em HTML
+- Contratos `.logline` reais e simulados
+- Scripts auxiliares e código-fonte em Go, Rust e Python
+- Suporte a spans, replay e visualização simbólica
+
+## Estrutura
+
+- `contracts/`: contratos organizacionais, cognitivos e de sistema
+- `bin/`: executores (`logline`, `watch`, `archive`, `cognitive`)
+- `spans/`: rastreamento completo das execuções
+- `ui/`: visualização simbólica e auditável
+- `scripts/`: utilitários de instalação, cache, migração
+- `src/`: código-fonte modular por linguagem
+
+## Execução
+
+```bash
+logline run contracts/system/build_vm.logline
+systemctl start logline-watch
+```
+
+Acesse o frontend em `http://localhost:8080`.
diff --git a/terreno_completo_FINAL/auth/tokens.txt b/terreno_completo_FINAL/auth/tokens.txt
new file mode 100644
index 0000000000000000000000000000000000000000..13b9b6bbd0ec3cf2eb1e8959ca531f6cd6cfbf97
--- /dev/null
+++ b/terreno_completo_FINAL/auth/tokens.txt
@@ -0,0 +1,2 @@
+admin_token|*|never
+user_token|read,write|2025-12-31
diff --git a/terreno_completo_FINAL/bin/logline b/terreno_completo_FINAL/bin/logline
new file mode 100755
index 0000000000000000000000000000000000000000..2f2e79949d28b5d7b97e6e8fc559b8d0955629f0
--- /dev/null
+++ b/terreno_completo_FINAL/bin/logline
@@ -0,0 +1,2 @@
+#!/bin/bash
+echo 'run'
\ No newline at end of file
diff --git a/terreno_completo_FINAL/bin/logline-archive b/terreno_completo_FINAL/bin/logline-archive
new file mode 100755
index 0000000000000000000000000000000000000000..e0ef23fd17b86f60f8fa80702370d4d456f9e0f7
--- /dev/null
+++ b/terreno_completo_FINAL/bin/logline-archive
@@ -0,0 +1,2 @@
+#!/bin/bash
+echo 'archive'
\ No newline at end of file
diff --git a/terreno_completo_FINAL/bin/logline-cognitive b/terreno_completo_FINAL/bin/logline-cognitive
new file mode 100755
index 0000000000000000000000000000000000000000..d6829ce85d07ea0943242d1225a68a7fda779477
--- /dev/null
+++ b/terreno_completo_FINAL/bin/logline-cognitive
@@ -0,0 +1,2 @@
+#!/bin/bash
+echo 'cognitive'
\ No newline at end of file
diff --git a/terreno_completo_FINAL/bin/logline-watch b/terreno_completo_FINAL/bin/logline-watch
new file mode 100755
index 0000000000000000000000000000000000000000..b526887f91571fcdb2a153ef9ce694ed7e3d3d5e
--- /dev/null
+++ b/terreno_completo_FINAL/bin/logline-watch
@@ -0,0 +1,2 @@
+#!/bin/bash
+echo 'watch'
\ No newline at end of file
diff --git a/terreno_completo_FINAL/contracts/llm/gerar_comando_shell.logline b/terreno_completo_FINAL/contracts/llm/gerar_comando_shell.logline
new file mode 100644
index 0000000000000000000000000000000000000000..ee1e7b8e12d6700ae0e333918d4c4ec1a7b7c0c1
--- /dev/null
+++ b/terreno_completo_FINAL/contracts/llm/gerar_comando_shell.logline
@@ -0,0 +1 @@
+# gerar_comando_shell.logline (placeholder)
\ No newline at end of file
diff --git a/terreno_completo_FINAL/contracts/llm/resumir_terreno.logline b/terreno_completo_FINAL/contracts/llm/resumir_terreno.logline
new file mode 100644
index 0000000000000000000000000000000000000000..addfa2efe076ba692166ca761d5476286cd5160e
--- /dev/null
+++ b/terreno_completo_FINAL/contracts/llm/resumir_terreno.logline
@@ -0,0 +1 @@
+# resumir_terreno.logline (placeholder)
\ No newline at end of file
diff --git a/terreno_completo_FINAL/contracts/llm/simulate_llm.logline b/terreno_completo_FINAL/contracts/llm/simulate_llm.logline
new file mode 100644
index 0000000000000000000000000000000000000000..d20af1b8cce35286b300bc9db65e7150dcdac914
--- /dev/null
+++ b/terreno_completo_FINAL/contracts/llm/simulate_llm.logline
@@ -0,0 +1 @@
+# simulate_llm.logline (placeholder)
\ No newline at end of file
diff --git a/terreno_completo_FINAL/contracts/misc/._contract.flipapp.logline b/terreno_completo_FINAL/contracts/misc/._contract.flipapp.logline
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/contracts/misc/._dashboard.logline b/terreno_completo_FINAL/contracts/misc/._dashboard.logline
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/contracts/misc/._deployment.yaml b/terreno_completo_FINAL/contracts/misc/._deployment.yaml
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/contracts/misc/._github-actions.yml b/terreno_completo_FINAL/contracts/misc/._github-actions.yml
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/contracts/misc/._gitlab-ci.yml b/terreno_completo_FINAL/contracts/misc/._gitlab-ci.yml
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/contracts/misc/._k8s_deploy.logline b/terreno_completo_FINAL/contracts/misc/._k8s_deploy.logline
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/contracts/misc/._login.logline b/terreno_completo_FINAL/contracts/misc/._login.logline
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/contracts/misc/._login_flow.logline b/terreno_completo_FINAL/contracts/misc/._login_flow.logline
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/contracts/misc/._loki-config.yaml b/terreno_completo_FINAL/contracts/misc/._loki-config.yaml
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/contracts/misc/._onboarding_user.logline b/terreno_completo_FINAL/contracts/misc/._onboarding_user.logline
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/contracts/misc/._prometheus.yml b/terreno_completo_FINAL/contracts/misc/._prometheus.yml
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/contracts/misc/._query_contracts.logline b/terreno_completo_FINAL/contracts/misc/._query_contracts.logline
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/contracts/misc/._splash.logline b/terreno_completo_FINAL/contracts/misc/._splash.logline
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/contracts/misc/._summary.logline b/terreno_completo_FINAL/contracts/misc/._summary.logline
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/contracts/misc/._whatsapp_flow.logline b/terreno_completo_FINAL/contracts/misc/._whatsapp_flow.logline
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/contracts/misc/contract.flipapp.logline b/terreno_completo_FINAL/contracts/misc/contract.flipapp.logline
new file mode 100644
index 0000000000000000000000000000000000000000..4f5ab5ae561c399a6c5df5d5d6f765fd37306b89
--- /dev/null
+++ b/terreno_completo_FINAL/contracts/misc/contract.flipapp.logline
@@ -0,0 +1,23 @@
+# Contrato-mãe do FlipApp/Minicontratos
+- contract: flipapp
+  who: system
+  did: iniciar_ambiente
+  payload:
+    descricao: "Ambiente institucional VoulezVous"
+    data_inicio: "2025-06-02"
+
+- contract: onboarding
+  who: cliente
+  did: enviar_mensagem
+  payload:
+    texto: "Bem-vindo ao sistema!"
+
+- did: registrar_presenca
+  who: cliente
+  when: "2025-06-02T09:00:00Z"
+  affair: onboarding
+
+- did: enviar_audio
+  payload:
+    duracao: 12
+    origem: whatsapp
diff --git a/terreno_completo_FINAL/contracts/misc/dashboard.logline b/terreno_completo_FINAL/contracts/misc/dashboard.logline
new file mode 100644
index 0000000000000000000000000000000000000000..575f8a242e898514780063337adf7d94bd57863f
--- /dev/null
+++ b/terreno_completo_FINAL/contracts/misc/dashboard.logline
@@ -0,0 +1,32 @@
+- id: "ui.dashboard"
+  do:
+    - type: ui.route
+      path: "/dashboard"
+      title: "Dashboard"
+    - type: ui.render
+      blocks:
+        - component: "h2"
+          props:
+            text: "Contratos Ativos"
+            class: "text-2xl font-bold mb-4"
+        - component: "table"
+          props:
+            class: "min-w-full bg-white shadow rounded-lg"
+        - component: "tbody"
+          children:
+            - component: "tr"
+              children:
+                - component: "td"
+                  props:
+                    text: "Contrato 1"
+                - component: "td"
+                  props:
+                    text: "Ativo"
+            - component: "tr"
+              children:
+                - component: "td"
+                  props:
+                    text: "Contrato 2"
+                - component: "td"
+                  props:
+                    text: "Ativo"
diff --git a/terreno_completo_FINAL/contracts/misc/deployment.yaml b/terreno_completo_FINAL/contracts/misc/deployment.yaml
new file mode 100644
index 0000000000000000000000000000000000000000..0caed9849d2691be23ac5e359b36f4e89528601a
--- /dev/null
+++ b/terreno_completo_FINAL/contracts/misc/deployment.yaml
@@ -0,0 +1,38 @@
+apiVersion: apps/v1
+kind: Deployment
+metadata:
+  name: minicontratos-deployment
+spec:
+  replicas: 2
+  selector:
+    matchLabels:
+      app: minicontratos
+  template:
+    metadata:
+      labels:
+        app: minicontratos
+    spec:
+      containers:
+        - name: minicontratos
+          image: myregistry/minicontratos:latest
+          ports:
+            - containerPort: 8080
+          env:
+            - name: OPENAI_API_KEY
+              valueFrom:
+                secretKeyRef:
+                  name: openai-secret
+                  key: api_key
+---
+apiVersion: v1
+kind: Service
+metadata:
+  name: minicontratos-svc
+spec:
+  selector:
+    app: minicontratos
+  ports:
+    - protocol: TCP
+      port: 80
+      targetPort: 8080
+  type: ClusterIP
diff --git a/terreno_completo_FINAL/contracts/misc/file_watcher.logline b/terreno_completo_FINAL/contracts/misc/file_watcher.logline
new file mode 100644
index 0000000000000000000000000000000000000000..1eb7931837cc60fb58e8633baaf0584ee09553b6
--- /dev/null
+++ b/terreno_completo_FINAL/contracts/misc/file_watcher.logline
@@ -0,0 +1,43 @@
+# workflow: file_watcher.logline
+# Descrição: Workflow para demonstrar o monitoramento contínuo de um arquivo ou webhook.
+
+name: "Monitor de Arquivo e Webhook"
+description: "Monitora um arquivo simulado e um endpoint de webhook para padrões específicos."
+
+tasks:
+  - id: monitorar_status_servico
+    type: observe
+    spec:
+      source: "/var/log/my_service_status.log" # Caminho simulado de um arquivo de log
+      pattern: "SERVICE_READY|ALL_OK" # Padrão que indica que o serviço está pronto
+      interval: "5s" # Verifica a cada 5 segundos
+      action:
+        type: "alert"
+        target: "ops@voulezvous.tv"
+        message: "Status de serviço alterado: $(pattern) detectado em $(source)."
+    hooks:
+      - "on_block_success"
+      - "on_block_error" # Se o monitor falhar (e.g. erro de leitura)
+
+  - id: receber_evento_deploy
+    type: observe
+    spec:
+      source: "https://api.example.com/deploy/status" # URL conceitual
+      pattern: "DEPLOY_COMPLETE|ROLLBACK_SUCCESS" # Padrão no corpo do webhook (ou query params)
+      on_event: "webhook" # Espera um evento de webhook
+      listen_path: "/webhooks/deploy-status" # O caminho que o servidor LogLine vai escutar
+      action:
+        type: "trigger"
+        target: "pos_deploy_action" # Bloco a ser disparado (abaixo)
+        message: "Evento de deploy recebido e processado!"
+    hooks:
+      - "on_block_success"
+
+  - id: pos_deploy_action # Bloco a ser disparado pelo webhook
+    type: mechanical
+    spec:
+      command: "echo"
+      args: ["Ação pós-deploy disparada! Evento: {{.receber_evento_deploy.output.raw_event_content | json}}"] # <-- Captura e mostra o payload do webhook
+    when: "false" # Este bloco só é ativado por trigger, não roda em fluxo normal
+    hooks:
+      - "on_block_success"
\ No newline at end of file
diff --git a/terreno_completo_FINAL/contracts/misc/github-actions.yml b/terreno_completo_FINAL/contracts/misc/github-actions.yml
new file mode 100644
index 0000000000000000000000000000000000000000..775e3e0aecc8a21c12db9d9e3c541ea04eb47e4c
--- /dev/null
+++ b/terreno_completo_FINAL/contracts/misc/github-actions.yml
@@ -0,0 +1,13 @@
+name: CI
+on:
+  push:
+    branches: [ main ]
+jobs:
+  build:
+    runs-on: ubuntu-latest
+    steps:
+      - uses: actions/checkout@v2
+      - name: Build Backend
+        run: cd backend && cargo build --release
+      - name: Build Frontend
+        run: cd frontend && ./build.sh
diff --git a/terreno_completo_FINAL/contracts/misc/gitlab-ci.yml b/terreno_completo_FINAL/contracts/misc/gitlab-ci.yml
new file mode 100644
index 0000000000000000000000000000000000000000..051b3f2437b3ebe6df7046ce48963a7185613220
--- /dev/null
+++ b/terreno_completo_FINAL/contracts/misc/gitlab-ci.yml
@@ -0,0 +1,14 @@
+stages:
+  - build
+  - test
+  - deploy
+
+build_backend:
+  stage: build
+  script:
+    - cd backend && cargo build --release
+
+build_frontend:
+  stage: build
+  script:
+    - cd frontend && ./build.sh
diff --git a/terreno_completo_FINAL/contracts/misc/iterative_process.logline b/terreno_completo_FINAL/contracts/misc/iterative_process.logline
new file mode 100644
index 0000000000000000000000000000000000000000..f794ec282a0e286977a4db71cf223775bb137eb7
--- /dev/null
+++ b/terreno_completo_FINAL/contracts/misc/iterative_process.logline
@@ -0,0 +1,77 @@
+# workflow: iterative_process.logline
+# Descrição: Demonstra o uso de blocos 'loop' e 'await'.
+
+name: "Processo Iterativo e Sincronizado"
+description: "Um workflow que repete uma ação e aguarda por um evento externo."
+
+tasks:
+  - id: inicio_processo
+    type: mechanical
+    spec:
+      command: "echo"
+      args: ["Iniciando processo iterativo..."]
+
+  - id: loop_tentativas_api
+    type: loop
+    spec:
+      count: 3 # Tentar 3 vezes
+      # until: "external_service_status.output.available == true" # Outra opção de loop (descomentar para testar)
+      max_duration: "30s" # Garante que o loop não dure mais que 30 segundos
+    tasks: # Blocos aninhados dentro do loop
+      - id: chamar_api_externa
+        type: mechanical
+        spec:
+          command: "sh"
+          args: ["-c", "if (( RANDOM % 2 == 0 )); then echo 'API OK!'; else echo 'API FALHOU!'; exit 1; fi"] # 50% de chance de falhar
+        retry:
+          max_attempts: 2
+          initial_delay: "1s"
+          backoff_factor: 1.5
+        fallback:
+          type: cognitive
+          spec:
+            llm:
+              goal: "Sugira rota alternativa para API falha"
+              input: "A API 'chamar_api_externa' falhou. Gerar um plano de ação para o operador."
+              trust: "medium"
+            action_on_output: # Simplesmente loga o disparo aqui
+              trigger_block_id: "log_acao_sugerida"
+      - id: log_acao_sugerida # Este bloco seria o 'alvo' do action_on_output
+        type: mechanical
+        spec:
+          command: "echo"
+          args: ["Ação sugerida pela IA: {{.llm_output.output}}"] # O LLMOutput viria do action_on_output
+        when: "false" # Este bloco só é ativado por trigger_block_id, não roda em fluxo normal
+      - id: registrar_tentativa
+        type: register
+        spec:
+          entity: "log_tentativa_api"
+          id: "tentativa-{{.loop_iteration}}-{{.chamar_api_externa.status}}-{{.Timestamp | printf \"%.8s\"}}" # <-- NOVO: Usa variáveis de loop e histórico
+          data:
+            loop_id: "loop_tentativas_api"
+            iteration: "{{.loop_iteration}}" # <-- NOVO: Passa a iteração
+            status_api: "{{.chamar_api_externa.status}}" # <-- NOVO: Captura status da API
+            api_output: "{{.chamar_api_externa.output.stdout}}" # <-- NOVO: Captura output da API
+            api_error: "{{.chamar_api_externa.output.error_details}}" # <-- NOVO: Captura detalhes do erro
+        when: "true" # Sempre registra
+
+  - id: signal_external_process
+    type: mechanical
+    spec:
+      command: "sleep"
+      args: ["3s"] # Simula um trabalho que leva 3 segundos
+    when: "loop_tentativas_api.status == 'success'"
+
+  - id: esperar_processo_externo
+    type: await
+    spec:
+      task_id: "signal_external_process"
+      condition: "status == 'success'"
+      timeout: "10s"
+
+  - id: processo_finalizado
+    type: mechanical
+    spec:
+      command: "echo"
+      args: ["Processo iterativo e sincronizado concluído!"]
+    when: "esperar_processo_externo.status == 'success'"
\ No newline at end of file
diff --git a/terreno_completo_FINAL/contracts/misc/k8s_deploy.logline b/terreno_completo_FINAL/contracts/misc/k8s_deploy.logline
new file mode 100644
index 0000000000000000000000000000000000000000..383cea27bf4263272d741202a00cd3d0f6dbd4b2
--- /dev/null
+++ b/terreno_completo_FINAL/contracts/misc/k8s_deploy.logline
@@ -0,0 +1,11 @@
+- id: "infra.k8s_deploy"
+  do:
+    - type: infra.k8s
+      action: "apply"
+      file: "deployment.yaml"
+      target: "k8s_result"
+    - type: ui.render
+      blocks:
+        - component: "div"
+          props:
+            text: "Resultado do deploy: {{k8s_result}}"
diff --git a/terreno_completo_FINAL/contracts/misc/llm_hook_workflow.logline b/terreno_completo_FINAL/contracts/misc/llm_hook_workflow.logline
new file mode 100644
index 0000000000000000000000000000000000000000..f6a6888da20f8fec50af550d45116875ad4c6a75
--- /dev/null
+++ b/terreno_completo_FINAL/contracts/misc/llm_hook_workflow.logline
@@ -0,0 +1,68 @@
+# workflow: llm_hook_workflow.logline
+# Descrição: Demonstra a integração com LLM e o sistema de hooks.
+
+name: "Processo com IA e Hooks"
+description: "Um workflow que usa inteligência artificial para resumir um texto e dispara hooks customizados."
+
+tasks:
+  - id: gerar_texto_para_ia
+    type: mechanical
+    spec:
+      command: "echo"
+      args: ["O LogLineOS é uma linguagem operacional de alto nível, desenhada para auditoria e rastreabilidade de processos. Cada bloco é um gesto organizacional auditável."]
+    hooks:
+      - "on_block_success" # Dispara hook global de sucesso
+      - "on_block_start"   # Dispara hook global de início
+      - "my_specific_hook" # Dispara um hook nomeado customizado (definido em logline.yml)
+
+  - id: resumir_texto_com_ia
+    type: cognitive
+    spec:
+      llm:
+        goal: "Resumir o texto sobre LogLineOS em uma frase curta, e classificar o sentimento como 'positivo' ou 'neutro' em formato JSON com campos 'summary' e 'sentiment'."
+        input: "Texto: O LogLineOS é uma linguagem operacional de alto nível, desenhada para auditoria e rastreabilidade de processos. Cada bloco é um gesto organizacional auditável. Ele permite simular, executar e auditar tudo. É uma ferramenta fantástica!"
+        trust: "high"
+        model: "gpt-3.5-turbo" # Pode sobrescrever o modelo padrão do logline.yml
+      action_on_output: # Ação baseada no output do LLM
+        trigger_block_id: "reagir_ao_sentimento_ia" # Bloco a ser disparado
+        condition: "llm_output.sentiment == 'positivo'" # Condição para disparar
+        map_output_to_input: # Mapeia campos do LLM output para input do bloco disparado
+          resumo: "{{.llm_output.summary}}"
+          sentimento: "{{.llm_output.sentiment}}"
+    when: "gerar_texto_para_ia.status == 'success'"
+    hooks:
+      - "on_block_success"
+      - "on_block_error"
+      - "my_specific_hook"
+
+  - id: reagir_ao_sentimento_ia # Bloco disparado pela ação cognitiva
+    type: mechanical
+    spec:
+      command: "echo"
+      # Para usar .resumo e .sentimento aqui, o executor precisaria injetar
+      # dinamicamente estes valores no 'spec' do bloco disparado antes da execução.
+      # Por enquanto, eles seriam parte do 'output' do span que disparou este.
+      args: ["IA reagiu! Saída LLM completa: {{.resumir_texto_com_ia.output.llm_output | json}}"] # <-- Exemplo que espera JSON
+    when: "false" # Este bloco só é ativado por trigger_block_id, não roda em fluxo normal
+    hooks:
+      - "on_block_success"
+
+  - id: registrar_resumo
+    type: register
+    spec:
+      entity: "resumo_ia"
+      id: "resumo_logline_ia"
+      data:
+        resumo_gerado: "{{.resumir_texto_com_ia.output.llm_output}}" # Este ainda pegaria a string bruta. Melhorar para JSON.
+        bloco_origem: "resumir_texto_com_ia"
+        timestamp_resumo: "{{.Timestamp}}"
+    when: "resumir_texto_com_ia.status == 'success'"
+    hooks:
+      - "on_block_success"
+
+  - id: finalizar_processo_ia
+    type: mechanical
+    spec:
+      command: "echo"
+      args: ["Processo com IA e hooks concluído!"]
+    when: "resumir_texto_com_ia.status == 'success' or registrar_resumo.status == 'success'"
\ No newline at end of file
diff --git a/terreno_completo_FINAL/contracts/misc/login.logline b/terreno_completo_FINAL/contracts/misc/login.logline
new file mode 100644
index 0000000000000000000000000000000000000000..dfa974ccb677c01e9ca53f5a3b1d7ab29032d04d
--- /dev/null
+++ b/terreno_completo_FINAL/contracts/misc/login.logline
@@ -0,0 +1,29 @@
+- id: "ui.login"
+  do:
+    - type: ui.route
+      path: "/login"
+      title: "Login"
+    - type: ui.render
+      blocks:
+        - component: "form"
+          props:
+            class: "flex flex-col gap-4 max-w-sm mx-auto mt-20"
+            action: "/api/auth/login"
+            method: "POST"
+        - component: "input"
+          props:
+            type: "email"
+            name: "email"
+            placeholder: "E-mail"
+            class: "border p-2 rounded"
+        - component: "input"
+          props:
+            type: "password"
+            name: "senha"
+            placeholder: "Senha"
+            class: "border p-2 rounded"
+        - component: "button"
+          props:
+            type: "submit"
+            text: "Entrar"
+            class: "bg-blue-600 text-white px-4 py-2 rounded"
diff --git a/terreno_completo_FINAL/contracts/misc/login_flow.logline b/terreno_completo_FINAL/contracts/misc/login_flow.logline
new file mode 100644
index 0000000000000000000000000000000000000000..9a83ec8e4f1e19af30b9564e4f943f4ec84fac7a
--- /dev/null
+++ b/terreno_completo_FINAL/contracts/misc/login_flow.logline
@@ -0,0 +1,12 @@
+- id: "auth.login_flow"
+  do:
+    - type: auth.login
+      input:
+        email: "{{email}}"
+        senha: "{{senha}}"
+      target: "login_result"
+    - type: ui.render
+      blocks:
+        - component: "div"
+          props:
+            text: "Resultado do login: {{login_result}}"
diff --git a/terreno_completo_FINAL/contracts/misc/logline.yml b/terreno_completo_FINAL/contracts/misc/logline.yml
new file mode 100644
index 0000000000000000000000000000000000000000..0969d4d1b0c8bcfdb9d850d4d7de58493ea6141a
--- /dev/null
+++ b/terreno_completo_FINAL/contracts/misc/logline.yml
@@ -0,0 +1,58 @@
+# config/logline.yml
+# Configuração global para LogLineOS
+
+llm:
+  provider: "mock" # ou "openai", "ollama"
+  openai:
+    api_key: "YOUR_OPENAI_API_KEY" # Substitua pela sua chave REAL se usar OpenAI
+    model: "gpt-3.5-turbo" # Exemplo: gpt-4o, gpt-3.5-turbo, etc.
+  ollama: # Exemplo de configuração para Ollama (se provider for "ollama")
+    base_url: "http://localhost:11434"
+    model: "llama3" # Exemplo: llama3, mistral
+
+hooks:
+  # Exemplo de hook global que roda após cada bloco, se for sucesso
+  on_block_success:
+    - name: "log_success"
+      type: "shell"
+      command: "echo"
+      args: ["LogLine Hook: Bloco {{.BlockID}} ({{.Type}}) SUCCESSO! Span: {{.SpanID}}"]
+  # Exemplo de hook global que roda após cada bloco, se falhar
+  on_block_error:
+    - name: "notify_error"
+      type: "shell"
+      command: "echo"
+      args: ["LogLine Hook: Bloco {{.BlockID}} ({{.Type}}) ERRO! Detalhes: {{.Error}}"]
+  # Exemplo de hook global que roda após cada bloco, se for pulado
+  on_block_skipped:
+    - name: "log_skipped"
+      type: "shell"
+      command: "echo"
+      args: ["LogLine Hook: Bloco {{.BlockID}} ({{.Type}}) PULADO! Razão: {{.Message}}"]
+  # Exemplo de hook global que roda no início do workflow
+  on_workflow_start:
+    - name: "log_workflow_start"
+      type: "shell"
+      command: "echo"
+      args: ["LogLine Hook: Workflow {{.WorkflowName}} INICIADO!"]
+  # Exemplo de hook global que roda no fim do workflow
+  on_workflow_end:
+    - name: "log_workflow_end"
+      type: "shell"
+      command: "echo"
+      args: ["LogLine Hook: Workflow {{.WorkflowName}} FINALIZADO com status: {{.Status}}!"]
+
+  # Exemplo de hook específico nomeado (usado no workflow llm_hook_workflow.logline)
+  my_specific_hook:
+    - name: "custom_process_shell"
+      type: "shell"
+      command: "echo"
+      args: ["Hook customizado 'my_specific_hook' disparado para bloco {{.BlockID}}!"]
+    # Removido exemplo HTTP para simplificar a dependência de um endpoint real para demo
+    # - name: "http_notification"
+    #   type: "http"
+    #   url: "http://localhost:8081/webhook/logline-hook" # Substitua pelo seu endpoint real para testar
+    #   method: "POST"
+    #   headers:
+    #     Content-Type: "application/json"
+    #   body: '{"event": "custom_hook_triggered", "block_id": "{{.BlockID}}", "status": "{{.Status}}", "span_id": "{{.SpanID}}"}'
\ No newline at end of file
diff --git a/terreno_completo_FINAL/contracts/misc/loki-config.yaml b/terreno_completo_FINAL/contracts/misc/loki-config.yaml
new file mode 100644
index 0000000000000000000000000000000000000000..017f446a93f698e240612b53c46dc0741bf3b590
--- /dev/null
+++ b/terreno_completo_FINAL/contracts/misc/loki-config.yaml
@@ -0,0 +1,3 @@
+auth_enabled: false
+server:
+  http_listen_port: 3100
diff --git a/terreno_completo_FINAL/contracts/misc/onboarding_ana.logline b/terreno_completo_FINAL/contracts/misc/onboarding_ana.logline
new file mode 100644
index 0000000000000000000000000000000000000000..e1490ae554cb6991d242b86c811195b9a3ee1b11
--- /dev/null
+++ b/terreno_completo_FINAL/contracts/misc/onboarding_ana.logline
@@ -0,0 +1,91 @@
+# workflow: onboarding_ana.logline
+# Descrição: Workflow de cadastro e formalização de uma nova performer.
+
+name: "Onboarding de Performer - Ana"
+description: "Registra nova pessoa, sua chegada, e finaliza o processo com um commit auditável."
+
+tasks:
+  - id: registrar_ana_performer
+    type: register
+    spec:
+      entity: "performer"
+      id: "ana_2025"
+      data:
+        nome_completo: "Ana Paula Silva"
+        email: "ana.paula@voulezvous.tv"
+        status_onboarding: "pendente_presenca"
+    hooks:
+      - "on_block_success" # Exemplo de hook específico para este bloco
+
+  - id: simular_comando_interno
+    type: mechanical
+    spec:
+      command: "echo"
+      args: ["Simulando um comando interno com sucesso!"]
+
+  - id: tentativa_de_conexao_falha
+    type: mechanical
+    spec:
+      command: "sh"
+      args: ["-c", "exit 1"] # <-- ESTE COMANDO IRÁ FALHAR!
+    retry:
+      max_attempts: 3
+      initial_delay: "500ms"
+      backoff_factor: 2.0
+    fallback:
+      type: cognitive
+      spec:
+        llm:
+          goal: "Analisar erro de conexão e sugerir alternativa humana."
+          input: "Falha ao conectar com o serviço X. Gerar um plano de ação para o operador."
+          trust: "medium"
+        action_on_output:
+          trigger_block_id: "notificar_operador_manual" # Este bloco será "disparado"
+          condition: "llm_output.confidence > 0.7" # Exemplo: disparar só se IA confia na sugestão
+    hooks:
+      - "on_block_error" # Disparar hook global de erro
+      - "on_block_fallback_success" # Disparar se o fallback for sucesso
+  
+  - id: notificar_operador_manual # Bloco alvo de action_on_output
+    type: mechanical
+    spec:
+      command: "echo"
+      args: ["Notificação do operador disparada! Sugestão da IA: {{.llm_output.output}}"]
+    when: "false" # Só é ativado via action_on_output, não em fluxo normal
+
+  - id: registrar_presenca_ana
+    type: affair
+    spec:
+      type: "presença"
+      from: "Dan"
+      to: "ana_2025"
+      note: "Ana chegou ao estúdio para o ensaio fotográfico inicial."
+      intensity: 95
+    when: "registrar_ana_performer.status == 'success'"
+
+  - id: verificar_documentacao_pendente
+    type: observe
+    spec:
+      source: "/vault/docs/ana_2025_contrato.pdf"
+      pattern: "ASSINATURA_OK|REVISAO_CONCLUIDA"
+      action:
+        type: "alert"
+        target: "legal@voulezvous.tv"
+        message: "Documentação de Ana (contrato) requer atenção: $(pattern) detectado em $(source)."
+    when: "registrar_presenca_ana.status == 'success'"
+
+  - id: selar_onboarding_performer
+    type: commit
+    spec:
+      goal: "Onboarding da Performer Ana concluído e formalizado."
+      inputs:
+        - "registrar_ana_performer"
+        - "registrar_presenca_ana"
+        - "verificar_documentacao_pendente"
+      signed_by:
+        - "Dan"
+        - "Departamento_RH_RH"
+      message: "Performer Ana Paula Silva foi registrada e sua presença formalizada. Contrato verificado. Pronta para iniciar atividades."
+    when: "verificar_documentacao_pendente.output.matched == true"
+    hooks:
+      - "on_block_success"
\ No newline at end of file
diff --git a/terreno_completo_FINAL/contracts/misc/onboarding_user.logline b/terreno_completo_FINAL/contracts/misc/onboarding_user.logline
new file mode 100644
index 0000000000000000000000000000000000000000..d7a1b1e1628efa629316eef72963e286ea394151
--- /dev/null
+++ b/terreno_completo_FINAL/contracts/misc/onboarding_user.logline
@@ -0,0 +1,18 @@
+- id: "workflows.onboarding_user"
+  do:
+    - type: db.query
+      query: "SELECT * FROM users WHERE email = '{{email}}'"
+      target: "user"
+    - type: llm.gateway
+      input:
+        prompt: "Crie uma mensagem de boas-vindas para {{user.name}}"
+      target: "welcome_msg"
+    - type: messaging.email
+      to: "{{user.email}}"
+      subject: "Bem-vindo!"
+      body: "{{welcome_msg}}"
+    - type: ui.render
+      blocks:
+        - component: "div"
+          props:
+            text: "Onboarding concluído para {{user.email}}"
diff --git a/terreno_completo_FINAL/contracts/misc/prometheus.yml b/terreno_completo_FINAL/contracts/misc/prometheus.yml
new file mode 100644
index 0000000000000000000000000000000000000000..a11fc6ca81d2ee8ed2cf4dd2d14410fcb36f616a
--- /dev/null
+++ b/terreno_completo_FINAL/contracts/misc/prometheus.yml
@@ -0,0 +1,6 @@
+global:
+  scrape_interval: 15s
+scrape_configs:
+  - job_name: 'minicontratos-backend'
+    static_configs:
+      - targets: ['backend:8080']
diff --git a/terreno_completo_FINAL/contracts/misc/query_contracts.logline b/terreno_completo_FINAL/contracts/misc/query_contracts.logline
new file mode 100644
index 0000000000000000000000000000000000000000..44c45709574322a306feedf9f4de08df770ba23f
--- /dev/null
+++ b/terreno_completo_FINAL/contracts/misc/query_contracts.logline
@@ -0,0 +1,10 @@
+- id: "db.query_contracts"
+  do:
+    - type: db.query
+      query: "SELECT id, title, status FROM contracts WHERE tenant_id = '{{tenant_id}}'"
+      target: "contracts"
+    - type: ui.render
+      blocks:
+        - component: "table"
+          props:
+            class: "min-w-full bg-white shadow rounded-lg"
diff --git a/terreno_completo_FINAL/contracts/misc/splash.logline b/terreno_completo_FINAL/contracts/misc/splash.logline
new file mode 100644
index 0000000000000000000000000000000000000000..2ef84d6b05820803984fa91a71b85277e0e6a51b
--- /dev/null
+++ b/terreno_completo_FINAL/contracts/misc/splash.logline
@@ -0,0 +1,19 @@
+- id: "ui.splash"
+  do:
+    - type: ui.route
+      path: "/"
+      title: "Minicontratos 2030"
+    - type: ui.render
+      blocks:
+        - component: "div"
+          props:
+            class: "flex flex-col items-center justify-center h-screen bg-slate-50"
+        - component: "h1"
+          props:
+            text: "Bem-vindo ao Minicontratos 2030"
+            class: "text-4xl font-bold text-slate-800 mb-6"
+        - component: "button"
+          props:
+            text: "Entrar"
+            class: "px-6 py-3 bg-blue-600 text-white rounded hover:bg-blue-700"
+            action: "/login"
diff --git a/terreno_completo_FINAL/contracts/misc/summary.logline b/terreno_completo_FINAL/contracts/misc/summary.logline
new file mode 100644
index 0000000000000000000000000000000000000000..dfc6ec16e7896ecf75732259e3354f75fb0959f7
--- /dev/null
+++ b/terreno_completo_FINAL/contracts/misc/summary.logline
@@ -0,0 +1,16 @@
+- id: "llm.summary"
+  do:
+    - type: llm.gateway
+      input:
+        prompt: "Resuma o contrato abaixo: {{contrato}}"
+        importance: "high"
+      props:
+        local_model: "mistral"
+        cloud_model: "gpt-4"
+        budget_usd: 0.10
+      target: "resumo"
+    - type: ui.render
+      blocks:
+        - component: "div"
+          props:
+            text: "Resumo gerado: {{resumo}}"
diff --git a/terreno_completo_FINAL/contracts/misc/whatsapp_flow.logline b/terreno_completo_FINAL/contracts/misc/whatsapp_flow.logline
new file mode 100644
index 0000000000000000000000000000000000000000..acf2d66b3678fa85e2d9bb879c713665ab0051d8
--- /dev/null
+++ b/terreno_completo_FINAL/contracts/misc/whatsapp_flow.logline
@@ -0,0 +1,12 @@
+- id: "messaging.whatsapp_flow"
+  do:
+    - type: messaging.whatsapp
+      conversation_id: "{{conversation_id}}"
+      sender: "{{sender}}"
+      content: "{{content}}"
+      target: "whatsapp_result"
+    - type: ui.render
+      blocks:
+        - component: "div"
+          props:
+            text: "Mensagem enviada: {{whatsapp_result}}"
diff --git a/terreno_completo_FINAL/contracts/system/auth.validate.logline b/terreno_completo_FINAL/contracts/system/auth.validate.logline
new file mode 100644
index 0000000000000000000000000000000000000000..9139936e94dee84b6211f6e4b0a17973d0311510
--- /dev/null
+++ b/terreno_completo_FINAL/contracts/system/auth.validate.logline
@@ -0,0 +1 @@
+# auth.validate.logline (placeholder)
\ No newline at end of file
diff --git a/terreno_completo_FINAL/contracts/system/build_vm.logline b/terreno_completo_FINAL/contracts/system/build_vm.logline
new file mode 100644
index 0000000000000000000000000000000000000000..9ad6a8e4388b397c6ed3902f08be74f70e65813f
--- /dev/null
+++ b/terreno_completo_FINAL/contracts/system/build_vm.logline
@@ -0,0 +1 @@
+# build_vm.logline (placeholder)
\ No newline at end of file
diff --git a/terreno_completo_FINAL/docs/._README.md b/terreno_completo_FINAL/docs/._README.md
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/docs/._architecture.md b/terreno_completo_FINAL/docs/._architecture.md
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/docs/._quickstart.md b/terreno_completo_FINAL/docs/._quickstart.md
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/docs/._span_reference.md b/terreno_completo_FINAL/docs/._span_reference.md
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/docs/README-INSTRUCOES.md b/terreno_completo_FINAL/docs/README-INSTRUCOES.md
new file mode 100644
index 0000000000000000000000000000000000000000..13211432f93670ce8899df3c585198b9ad3115a8
--- /dev/null
+++ b/terreno_completo_FINAL/docs/README-INSTRUCOES.md
@@ -0,0 +1,97 @@
+# 📦 LogLineOS: Instruções Finais de Instalação e Execução
+
+Este documento traz o passo a passo final, detalhado, para instalar dependências, compilar, executar e testar o LogLineOS em sua versão premium.
+
+---
+
+## 🛠️ Instalação e Build
+
+1. **Entre no diretório do projeto:**
+   ```
+   cd loglineos
+   ```
+
+2. **Instale TODAS as dependências Go:**
+   ```
+   go mod tidy
+   ```
+   - Se houver erros, revise os arquivos `.go` e repita.
+
+3. **Configure o `config/logline.yml`:**
+   - Se for usar OpenAI, coloque sua chave real no campo `api_key`.
+   - Se NÃO for usar OpenAI, altere `provider: "openai"` para `provider: "mock"` ou `provider: "ollama"`.
+   - Para Ollama, garanta que o Ollama esteja rodando localmente e um modelo carregado.
+
+4. **Inicialize o repositório Git (opcional, mas recomendado):**
+   ```
+   git init .
+   git config user.name "LogLineOS User"
+   git config user.email "user@loglineos.com"
+   echo "data/*" > .gitignore
+   echo "!data/snapshots/" >> .gitignore
+   git add .gitignore
+   git commit -m "Initial LogLineOS repository setup"
+   ```
+
+5. **Configure GPG (opcional, para assinatura):**
+   - Instale o GnuPG (`brew install gnupg` ou `sudo apt install gnupg`).
+   - Gere uma chave se necessário: `gpg --full-generate-key`.
+
+6. **Compile o LogLineOS:**
+   ```
+   go build -o logline
+   ```
+   - Se houver erros de build, revise os códigos e repita `go mod tidy`.
+
+7. **Limpe o diretório de dados (opcional):**
+   ```
+   rm -rf ./data
+   mkdir -p ./data
+   ```
+
+---
+
+## 🌟 Como Testar (Experiência Premium)
+
+### 1. **Execução de Workflow com LLM**
+   ```
+   rm -f ./data/spans.jsonl && ./logline run examples/workflows/llm_hook_workflow.logline
+   ```
+   - Observe hooks, chamadas LLM, disparos condicionais e o arquivo `data/spans.jsonl`.
+
+### 2. **Loops e Contexto Dinâmico**
+   ```
+   rm -f ./data/spans.jsonl && ./logline run examples/workflows/iterative_process.logline
+   ```
+   - Verifique spans dinâmicos e variáveis de loop em `data/spans.jsonl`.
+
+### 3. **Monitoramento Contínuo e Webhook**
+   ```
+   rm -f ./data/spans.jsonl && ./logline watch examples/workflows/file_watcher.logline
+   ```
+   - Em outro terminal:
+     ```
+     curl -X POST -H 'Content-Type: application/json' -d '{"status": "DEPLOY_COMPLETE", "version": "v1.0.0"}' http://localhost:8080/webhooks/deploy-status
+     ```
+   - Veja os eventos e spans gerados.
+
+### 4. **PromptPad - IDE Terminal**
+   ```
+   rm -f ./data/spans.jsonl && ./logline promptpad
+   ```
+   - Use comandos como `open`, `simulate`, `run`, `audit`, `commit`, `revert`, `replay`, `? <pergunta>` no PromptPad.
+   - Valide, edite e visualize grafos dos workflows, auditando tudo em tempo real.
+
+---
+
+## 💡 Dicas Finais
+
+- **Qualquer erro de dependência ou build:** revise os códigos, salve corretamente e repita `go mod tidy` e `go build`.
+- **Para logs e auditoria:** consulte sempre `data/spans.jsonl` e use o comando `logline audit` para relatórios.
+- **Explore todos os exemplos em `examples/workflows/`.**
+
+---
+
+**Parabéns! Você está com o LogLineOS pronto para operar em sua melhor versão — auditável, inteligente, visual, integrável e com uma IDE de terminal premium.**
+
+---
\ No newline at end of file
diff --git a/terreno_completo_FINAL/docs/README.md b/terreno_completo_FINAL/docs/README.md
new file mode 100644
index 0000000000000000000000000000000000000000..b9789bc2ed250f0f8180264e42ba7c0384b7c86d
--- /dev/null
+++ b/terreno_completo_FINAL/docs/README.md
@@ -0,0 +1,78 @@
+# FlipApp = Minicontratos
+
+O ambiente operacional oficial da VoulezVous, onde tudo é contrato, tudo é rastreável, e tudo acontece com semântica nativa LogLine.
+
+## Estrutura
+
+```
+flipapp/
+├── cmd/                   ← CLI e comandos locais
+├── internal/
+│   ├── contracts/         ← parser, executor, runner (LogLine core)
+│   ├── spans/             ← storage, timeline, JSONL e DB
+│   ├── auth/              ← JWT + tenants + users
+│   └── channels/          ← WhatsApp, OpenAI, Email...
+├── frontend/              ← FlipApp UI
+│   ├── PromptPad/         ← terminal IA + executor
+│   ├── ContractsView/     ← timeline e spans
+│   └── ChannelsPanel/     ← histórico por canal
+├── api/                   ← HTTP layer (Go), rotas REST
+├── config/                ← logline.yml, tenant settings
+├── migrations/
+├── scripts/
+└── README.md              ← manifesto institucional FlipApp = Minicontratos
+```
+
+## Telas
+
+- **LogLine**: REPL institucional, chat, executor de contratos, visualizador de spans.
+- **WhatsApp**: canal humano, cada conversa = contrato auditável, agente automático.
+- **New**: gerador rápido de contratos/registro, semeador de cultura institucional.
+
+## Princípios
+
+- Tudo é contrato LogLine.
+- Não há dashboards artificiais: toda consulta é feita via prompt.
+- Auditabilidade e rastreabilidade nativas.
+
+## Exemplos de Contrato LogLine
+
+```
+# Exemplo: registro de presença
+- did: registrar_presenca
+  who: cliente
+  when: "2025-06-02T09:00:00Z"
+  affair: onboarding
+```
+
+```
+# Exemplo: mensagem WhatsApp
+- did: enviar_audio
+  payload:
+    duracao: 12
+    origem: whatsapp
+```
+
+```
+# Exemplo: criar rotina de boas-vindas
+- contract: onboarding
+  who: cliente
+  did: enviar_mensagem
+  payload:
+    texto: "Bem-vindo ao sistema!"
+```
+
+## Rotas de API sugeridas
+
+- POST /api/logline
+- GET /api/logline/history?userId=&limit=
+- GET /api/whatsapp/conversations?tenantId=
+- GET /api/whatsapp/conversations/{conversationId}/messages?limit=&offset=
+- POST /api/whatsapp/send
+- GET /api/templates?tenantId=
+- POST /api/contracts
+- POST /api/contracts/{id}/execute
+
+## Como contribuir
+
+Abra issues, envie PRs ou sugira novos contratos LogLine para enriquecer o ecossistema institucional.
diff --git a/terreno_completo_FINAL/docs/architecture.md b/terreno_completo_FINAL/docs/architecture.md
new file mode 100644
index 0000000000000000000000000000000000000000..b591b8c212eb7dc4fb2ee0984a68573845f2e5ef
--- /dev/null
+++ b/terreno_completo_FINAL/docs/architecture.md
@@ -0,0 +1,3 @@
+# Minicontratos 2030 Architecture
+
+This document describes the architecture of the Minicontratos 2030 platform, including contracts, backend, frontend, infrastructure, and monitoring components.
diff --git a/terreno_completo_FINAL/docs/quickstart.md b/terreno_completo_FINAL/docs/quickstart.md
new file mode 100644
index 0000000000000000000000000000000000000000..22b55ecc99566c6ea6c7f831726c105b3816e41e
--- /dev/null
+++ b/terreno_completo_FINAL/docs/quickstart.md
@@ -0,0 +1,6 @@
+# Quickstart
+
+1. Clone the repository
+2. Build backend and frontend
+3. Run migrations
+4. Deploy using provided scripts
diff --git a/terreno_completo_FINAL/docs/span_reference.md b/terreno_completo_FINAL/docs/span_reference.md
new file mode 100644
index 0000000000000000000000000000000000000000..207335c3c69e8f2de4fc114734304017290553ca
--- /dev/null
+++ b/terreno_completo_FINAL/docs/span_reference.md
@@ -0,0 +1,3 @@
+# Span Reference
+
+This document describes the span model and its usage in Minicontratos 2030 contracts and backend.
diff --git a/terreno_completo_FINAL/llm_cache.jsonl b/terreno_completo_FINAL/llm_cache.jsonl
new file mode 100644
index 0000000000000000000000000000000000000000..e69de29bb2d1d6434b8b29ae775ad8c2e48c5391
diff --git a/terreno_completo_FINAL/logline.yml b/terreno_completo_FINAL/logline.yml
new file mode 100644
index 0000000000000000000000000000000000000000..2f4b026e9c56f5e39dbcd559c3a6070c44e05d15
--- /dev/null
+++ b/terreno_completo_FINAL/logline.yml
@@ -0,0 +1,8 @@
+fallbacks:
+  llm:
+    openai:
+      model: gpt-4
+      key: $OPENAI_KEY
+    local:
+      model: mistral
+      engine: ollama
diff --git a/terreno_completo_FINAL/manifesto.txt b/terreno_completo_FINAL/manifesto.txt
new file mode 100644
index 0000000000000000000000000000000000000000..4f08f860395418385f14eab534390c1741c530d4
--- /dev/null
+++ b/terreno_completo_FINAL/manifesto.txt
@@ -0,0 +1,135 @@
+e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855  llm_cache.jsonl
+3b7e7fefded2a10cfaed5eaeb71de7e5dbf40a07bd45e385180265d222cbe5c0  .terreno_version
+c44f7d8f9c8d15ae9c846f82b530c55a80c8bcb4c21050bbe45c1207a3b1f8eb  README.md
+9dc048798153ab8f7ba97ec7dd93b60a1dbc225b10fcf956ef70f3b3a8030fbe  logline.yml
+a25d04d1fd96d963e97e71122bd6f299890c99eb4206c59f01d8f59240a39d03  auth/tokens.txt
+2421a30978335fbbbd79f766b2ffc9f766a882e2541e314857c7cb201b4a5fef  bin/logline
+d5d44d65e6991a86791cf745789047af92a9d5af4fde94c60f038a6e574594d3  bin/logline-watch
+cf8f543b5c3518fde19b62fb6a15a13f0c05bd56d158e25bbbf290fc24da68f3  bin/logline-cognitive
+aeaef0f8c45a9ea4e2a7b41b25eb51b03abe3a9f8fa1436ea492dee0c5c25c33  bin/logline-archive
+896b06276bc74e43897e5b41827a9a35f9c9c6b21b6f67fc0e752f35cc3a95f5  contracts/system/build_vm.logline
+ffb94b20702b2be377add57096884bd25b3e8fc687bf62381ef1c362d8da6686  contracts/system/auth.validate.logline
+7f13e56f13194edc52e2c1fe46e340d03b42e383ac883c09c69ac4aabd2a77fc  contracts/llm/resumir_terreno.logline
+665e9640a449276ccbad4bfee60f4820b28d1072020367594d4101f006fa57ef  contracts/llm/gerar_comando_shell.logline
+83081cd110430ac2509802fe7439210882a7cb4eccde59e4e893aa74ebfcc1ab  contracts/llm/simulate_llm.logline
+9270b139475b7508368a60cb93df07ba3d59c25cd2b9618006e74b7c1f99078a  contracts/misc/github-actions.yml
+1f5b3e540aef780fbd3352cf104de968eb0ddd26209ccb155c7d9cd79b21b844  contracts/misc/gitlab-ci.yml
+c7d24fb01da93f35075ea6a550eefd5e5e0c2dd188c0424b1b83fbb42d02eaa6  contracts/misc/deployment.yaml
+ada49aa42eb8e9bc641bdbb9d20c797a61aa7d15dd8286036b3f5daf1805492f  contracts/misc/dashboard.logline
+8c1fddf6884932528d1fb71fbb6603f66002b9c2d52546bd7e0c1ac179947be7  contracts/misc/splash.logline
+6b467db781bcb949464996b36931b8f6ebae7b54acc4ab19e55ca437ba041078  contracts/misc/login.logline
+facb73f92ee4e8d50cf04e629bf5bad087bd31ab8dfddb7c9d467eadf053c74b  contracts/misc/k8s_deploy.logline
+b259e5396939c3fca08abc04912177b546a3b8e26e34743eb5485539e79537a7  contracts/misc/summary.logline
+5ac71e904aeebe72e040e3dabbfecbb2ff22f0295aebef53d142644347b370fd  contracts/misc/login_flow.logline
+1f098fcb64b1d1d69cc9446b293220022fafc7d5222648e379da8c0f46ccdfcc  contracts/misc/onboarding_user.logline
+61f123ca3414a4e34e8c6d55a08f0c2fc45ea967737f7e65bd3d234bfa285834  contracts/misc/query_contracts.logline
+bc3152b939c4832d49bff9b7bbe70c87b5f20d8393af4f30fdcb887a86282301  contracts/misc/whatsapp_flow.logline
+63dac241942a5b5110b7390eedb99fc4676ac0a0059b335dbf6fa4ec917981d0  contracts/misc/loki-config.yaml
+a0cdac1bae845d0bd0c4616f6c64935b2b0330a8122afd1a6f787405a734286c  contracts/misc/prometheus.yml
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  contracts/misc/._github-actions.yml
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  contracts/misc/._gitlab-ci.yml
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  contracts/misc/._deployment.yaml
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  contracts/misc/._dashboard.logline
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  contracts/misc/._splash.logline
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  contracts/misc/._login.logline
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  contracts/misc/._k8s_deploy.logline
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  contracts/misc/._summary.logline
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  contracts/misc/._login_flow.logline
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  contracts/misc/._onboarding_user.logline
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  contracts/misc/._query_contracts.logline
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  contracts/misc/._whatsapp_flow.logline
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  contracts/misc/._loki-config.yaml
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  contracts/misc/._prometheus.yml
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  contracts/misc/._contract.flipapp.logline
+c0272908aed7d09bd2beede23a377d90d20de71bdc6d75b8c06b99d0788fcd01  contracts/misc/contract.flipapp.logline
+c355d481b5c7c797b0045186ffa0915ec8505e3826fcc11d96ddc0f6f45eda47  contracts/misc/logline.yml
+8bc8fc77e49bbc985ade691da731566a910b0cc2f9b1f184dc46679098f267d7  contracts/misc/iterative_process.logline
+f387dfed7ca6953f48136d7bf057c629b0ceda0c49b5266738ae7d5a008a2435  contracts/misc/onboarding_ana.logline
+1f89cb2dbc4fd69e85acc3375609c9de9dae22b4654ffdfce1c773ec42230538  contracts/misc/file_watcher.logline
+c5e26e489effbb9967756985a7979f5fa8983dfa888d36fa054be294cd794d89  contracts/misc/llm_hook_workflow.logline
+f71aa00c1269ce31697ebc598cc7ad56c040cee8dd23c7b1760c36c9b834a9ae  ui/index.html
+48b6c931ca0d1a994744a28184de64fceba70fe9d544b6f642807ef487d25329  docs/architecture.md
+bb562ad006ae016a2bf1d21993d2f97c549b433004e5554f23c4b0950ba40658  docs/quickstart.md
+ca8d3cd219405229ec348526692cec6dc33d76e536f54827ef8d03d239bc061d  docs/span_reference.md
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  docs/._architecture.md
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  docs/._quickstart.md
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  docs/._span_reference.md
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  docs/._README.md
+c4d9023ea4812a7bfec42e513f7c62fe16ffd0c395f55578bc21cae7b2e6b3ee  docs/README.md
+66cb88eec284d04d7f675199e71da5279a6f799e28fe0307963bab88178c0c87  docs/README-INSTRUCOES.md
+97486253abb8ab736aa588f1775b11a5c771405af3c438c0179fe6adc006c881  scripts/install_flipapp.sh
+3428d0fb68f9c60a900ce4ae43ca91fb5c7d3cde6f8206878357720e39a826b5  scripts/migrate.sh
+f74b44ea0d28da47d7a5384455f554964e422978cbe583e5a850e8a62e1ef14d  scripts/deploy.sh
+849db217b157796761e3d493a3e7307b7c75d382eb5b0be4ff6d59be2a9caf63  scripts/rollback.sh
+1efd9a6f42e3ecdac6ad871be821885b8ef69e5fbf56e57d925f53416ed52854  scripts/build.sh
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  scripts/._install_flipapp.sh
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  scripts/._migrate.sh
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  scripts/._deploy.sh
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  scripts/._rollback.sh
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  scripts/._build.sh
+063acd8547994f9c96c9edae3457bdd11b2d11df6f9a6c04c82198a59f15e68f  src/go/errors.go
+9a0e99d31b1afc4fe997da092c95c8acf5e78909291f68e69f1a66120ffeff7b  src/go/config_loader.go
+56b206fdd83df48fe51efa479221effbd14519efd4b81fd675b76fcf36a042ee  src/go/llm_integrator.go
+831e545ab18d412e273023036b804508ef5bd5a9da474f2ed756ff9736314f8f  src/go/hooks.go
+efb192fe6c1d1591255c70b9b954887eb2ac5b24a1c06d44d9e3abe53469f2d4  src/go/ast.go
+61cb30ee9ad1561f2bbbba79711dcfb34b093178340de0632bcaa460e8c142ad  src/go/parser.go
+580ab8802e31e6df6ed979396fd08a3fb9756c8cdcc4ff1b282ad798ea6720a4  src/go/spans.go
+aea39c0218c5dffd0acf58db11154b7f7724f10bb1844835fc92873e8476db45  src/go/monitor.go
+a9587748f3b9422f0a0f708bc5fe9028df9c2752c7e34f01caaabd285f6097ca  src/go/planner.go
+9d9c56e577dd070dacf34f631229723bc130ee14034dbf7d813823f76ef1d747  src/go/snapshotter.go
+70d1f1db1c0498515b886d89144588a9814258ac523f75781321bd0492e2f7a2  src/go/ascii_graph.go
+4ea9feadb2ca515e06f5235b34b53e460cd985ca72ce48981b82d716c78bc3a8  src/go/format.go
+f291ed67ec587a237532af46334e91dfc8343277ceb7cd20b0ebe88b31669ae7  src/go/webserver.go
+546f7f73b1b0e234831f87828cdd7d4db08eed338df739068b0e72dc76776f52  src/go/promptpad_app.go
+9ea384e5421919b79779a80dca2a94dc6bbfb55b847d01bdd870631c144c8a13  src/rust/lib.rs
+8f8018f3d784fc26afa26206a0b1d6e28cd38fcf4b40c10d0c728ed36122f141  src/rust/handlers.rs
+2ba9ebecb5c6c5076ef2c5305e5b32fd8fdb6149e34dd6d12b3d74dd1c254cea  src/rust/logline_runner.rs
+c4a677d4c443b67d2193cda6ad99bfe56dfbfad8506202ec302443a474b0632b  src/rust/main.rs
+ff90e36d80ac78ac1b618d89a2bc26895b06fd64f9f6dd6b1167e139b7c9a8e8  src/rust/utils.rs
+9da84742f3949fc0e55a34b5a71ced3f93412f8abc441b86018c898e81a46a44  src/rust/time_utils.rs
+6b4ed048a39cf4e42d1abf33e16d9ed4702ee05f6d6b1ccecfda20bd04fb0a47  src/rust/validation.rs
+594d1e9fb83f9caa6d198550121375af4bd0d475c007b0cdf21986198f30bb54  src/rust/errors.rs
+c302da9c362d5ca535494ed977eda68897951885c028431b45036c08ad209059  src/rust/tenant.rs
+d989e24bed90d077f66b4bdc0dd3aaecef4e776ee65efe7b906c7006841fbafc  src/rust/user.rs
+9942bdb610fe06d82e8350ea8b95b79a4307d73785b30239749edbb11c2dcc66  src/rust/span.rs
+1845db5a31db0aaa1140b883c13fa674cdfa5b885314f89c1294c8bed2f0694e  src/rust/contract.rs
+cb950fcaa3a5305a170ff91ba84061a8762247c9f13a26eb470f8d42f2bcb872  src/rust/telemetry.rs
+d9261d871d479658eaf3feb56f66bf04b2e02db5c0a5e1eb2617b07b321adafb  src/rust/cache.rs
+5de4d669290dfa5f2de95deb7b2ac3bbfa0ed642821629f8fbb437fbe4743885  src/rust/db.rs
+47e1b29a13937e608fbabccc58d23033ea4b65302546516597573f108c9fa42f  src/rust/jwt.rs
+38bedb46599dc978dff00366d209538e2afe518db822f55d21dca8272b7ea158  src/rust/llm.rs
+1a42d9f7f4e933ff5b4c9050f4227b999490050af419a41d0b3b648054357bbb  src/rust/health.rs
+e5f4810fa15bf93300f8b0386d5667630c1a02564355732935036f5804940970  src/rust/whatsapp.rs
+845cf06307cf73e96f630f4e8721edaf89f413ea43369d24407f1bbb4eb40380  src/rust/auth.rs
+9ffa1d9cb7e43922751d506c45aaf29a2b01fa66bd8dbea75982340e7e49ae3d  src/rust/infra.rs
+2c116b47a7228af4f5d781116a8e4eeefaf775ce2b5c2673037e816974b11e21  src/rust/infra_service.rs
+938c3885ea19c0ae8af58949185e8a7c02456f3d7782ecc7072497a7ccda7c6d  src/rust/auth_service.rs
+e7572f375d644bf30508f27b8848a251e981b3572c46c67911bbba0c7ce93ba9  src/rust/db_service.rs
+76c5cae673320329fc1f6cc388ab50ee4822c446b1ac12d29011176a908eaeb8  src/rust/contract_service.rs
+22822bb4f74a76552aaf6af0d7453fe89ee384f72c5cee7d5902fe65479033f2  src/rust/llm_service.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._lib.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._handlers.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._logline_runner.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._main.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._utils.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._time_utils.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._validation.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._errors.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._tenant.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._user.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._span.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._contract.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._telemetry.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._cache.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._db.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._jwt.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._llm.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._health.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._whatsapp.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._auth.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._infra.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._infra_service.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._auth_service.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._db_service.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._contract_service.rs
+4c454a5284c3d744ce544381251fcc02c3510c274cb7ab1b223eebfa45142074  src/rust/._llm_service.rs
\ No newline at end of file
diff --git a/terreno_completo_FINAL/scripts/._build.sh b/terreno_completo_FINAL/scripts/._build.sh
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/scripts/._deploy.sh b/terreno_completo_FINAL/scripts/._deploy.sh
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/scripts/._install_flipapp.sh b/terreno_completo_FINAL/scripts/._install_flipapp.sh
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/scripts/._migrate.sh b/terreno_completo_FINAL/scripts/._migrate.sh
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/scripts/._rollback.sh b/terreno_completo_FINAL/scripts/._rollback.sh
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/scripts/build.sh b/terreno_completo_FINAL/scripts/build.sh
new file mode 100644
index 0000000000000000000000000000000000000000..761db533f14ce0956b26b80b20737eca09c44945
--- /dev/null
+++ b/terreno_completo_FINAL/scripts/build.sh
@@ -0,0 +1,11 @@
+#!/usr/bin/env bash
+set -e
+
+# 1. Compila o crate Rust para WASM usando wasm-pack
+cd frontend
+wasm-pack build --target web --out-dir static
+
+# 2. Ajusta permissões (se necessário) e volta
+cd ..
+
+echo "Frontend WASM compilado com sucesso em frontend/static/"
diff --git a/terreno_completo_FINAL/scripts/deploy.sh b/terreno_completo_FINAL/scripts/deploy.sh
new file mode 100644
index 0000000000000000000000000000000000000000..e6893832a17989ae5cc8ada164a851ec60badf7d
--- /dev/null
+++ b/terreno_completo_FINAL/scripts/deploy.sh
@@ -0,0 +1,3 @@
+#!/bin/bash
+# Deploy script for Minicontratos 2030
+kubectl apply -f ../k8s/deployment.yaml
diff --git a/terreno_completo_FINAL/scripts/install_flipapp.sh b/terreno_completo_FINAL/scripts/install_flipapp.sh
new file mode 100644
index 0000000000000000000000000000000000000000..0597d94f626cf90a25b330436b307640fac266da
--- /dev/null
+++ b/terreno_completo_FINAL/scripts/install_flipapp.sh
@@ -0,0 +1,386 @@
+#!/usr/bin/env bash
+set -euo pipefail
+
+# =============================================================================
+# Instalador “de primeira classe” do FlipApp / LogLineOS para VM Ubuntu/Debian
+# =============================================================================
+#
+# Este script faz tudo de forma automatizada:
+#   1. Verifica pré-requisitos (root, sistema suportado).
+#   2. Instala pacotes essenciais (git, build-essential, curl, Go, Rust, wasm-pack, Nginx, UFW).
+#   3. Cria usuário “flipapp” e estrutura de diretórios (/opt/flipapp, /etc/loglineos).
+#   4. Baixa ou extrai o código-fonte do FlipApp/LogLineOS (via Git ou arquivo ZIP local).
+#   5. Compila o back-end em Go e copia o binário para /opt/flipapp/bin/loglineos.
+#   6. Compila o front-end em Rust → gera WASM/JS e deposita em /opt/flipapp/static.
+#   7. Ajusta permissões, configura firewall (UFW) e Nginx (proxy reverso opcional).
+#   8. Cria e habilita serviço systemd para iniciar o FlipApp automaticamente.
+#
+# Uso:
+#   1) Copie este script para /opt/flipapp/install_flipapp.sh
+#   2) Edite as variáveis GIT_REPO (ou defina SOURCE_ZIP), se necessário.
+#   3) Torne executável: sudo chmod +x /opt/flipapp/install_flipapp.sh
+#   4) Execute: sudo /opt/flipapp/install_flipapp.sh
+#
+# Após a execução, o FlipApp estará disponível em http://<IP-da-VM>:8080
+#
+# Observação:
+#   • Se quiser usar um ZIP local (em vez de Git), defina SOURCE_ZIP=/caminho/para/arquivo.zip
+#   • Por padrão, o script define GIT_REPO="" e espera que SOURCE_ZIP seja informado ou que o
+#     repositório já esteja em /opt/flipapp/source. Ajuste conforme necessidade.
+# =============================================================================
+
+# ---------------------------- Variáveis Globais -----------------------------
+FLIPAPP_USER="flipapp"
+FLIPAPP_HOME="/opt/flipapp"
+CONFIG_DIR="/etc/loglineos"
+BIN_DIR="${FLIPAPP_HOME}/bin"
+SRC_DIR="${FLIPAPP_HOME}/source"
+STATIC_DIR="${FLIPAPP_HOME}/static"
+CONFIG_DEST="${FLIPAPP_HOME}/config"
+DATA_DIR="${FLIPAPP_HOME}/data"
+EXAMPLES_DIR="${FLIPAPP_HOME}/examples"
+SYSTEMD_UNIT="/etc/systemd/system/loglineos.service"
+REQUIRED_PORT="8080"
+
+# Escolha entre GIT_REPO (clonar) ou SOURCE_ZIP (extrair localmente).
+# Se preferir baixar de GitHub, preencha GIT_REPO. Caso contrário, deixe vazio e coloque um ZIP em /opt/flipapp/source.zip
+GIT_REPO=""  
+# Exemplo: GIT_REPO="https://github.com/seu-usuario/flipapp.git"
+SOURCE_ZIP="/opt/flipapp/source.zip"  # Se usar arquivo ZIP, coloque-o aqui
+
+# Cores para saída
+RED="$(printf '\033[0;31m')"
+GREEN="$(printf '\033[0;32m')"
+YELLOW="$(printf '\033[1;33m')"
+BLUE="$(printf '\033[0;34m')"
+NO_COLOR="$(printf '\033[0m')"
+
+# ----------------------------- Funções Auxiliares ---------------------------
+
+function info {
+    echo -e "${BLUE}[INFO]${NO_COLOR} $1"
+}
+
+function success {
+    echo -e "${GREEN}[OK]${NO_COLOR} $1"
+}
+
+function warning {
+    echo -e "${YELLOW}[AVISO]${NO_COLOR} $1"
+}
+
+function error_exit {
+    echo -e "${RED}[ERRO]${NO_COLOR} $1"
+    exit 1
+}
+
+# Verifica se o comando existe
+function check_command {
+    if ! command -v "$1" &> /dev/null; then
+        return 1
+    fi
+    return 0
+}
+
+# ------------------------ Verificações Iniciais -----------------------------
+
+# 1) Verificar se está rodando como root
+if [[ $EUID -ne 0 ]]; then
+    error_exit "Este script precisa ser executado como root (ou via sudo)."
+fi
+
+# 2) Verificar sistema operacional (Ubuntu/Debian)
+if ! grep -Ei "ubuntu|debian" /etc/os-release &> /dev/null; then
+    warning "Este script foi testado em Ubuntu/Debian. Se estiver usando outra distro, pode falhar."
+fi
+
+info "=============================="
+info "  Instalador do FlipApp/LogLineOS"
+info "=============================="
+
+# -------------------------- Instalação de Pacotes ---------------------------
+
+function install_base_packages {
+    info "Instalando pacotes essenciais: git, build-essential, curl, ufw..."
+    apt-get update -y
+    apt-get install -y git build-essential curl ufw
+    success "Pacotes básicos instalados."
+}
+
+function install_go {
+    if check_command go; then
+        warning "Go já está instalado ($(go version)). Pulando."
+    else
+        info "Instalando Go..."
+        apt-get install -y golang
+        if ! check_command go; then
+            error_exit "Falha ao instalar Go."
+        fi
+        success "Go instalado: $(go version)."
+    fi
+}
+
+function install_rust_and_wasm {
+    if check_command rustc && check_command cargo; then
+        warning "Rust já está instalado ($(rustc --version)). Pulando."
+    else
+        info "Instalando Rust (rustup)..."
+        curl https://sh.rustup.rs -sSf | bash -s -- -y
+        export PATH="/root/.cargo/bin:$PATH"
+        if ! check_command rustc || ! check_command cargo; then
+            error_exit "Falha ao instalar Rust."
+        fi
+        success "Rust instalado: $(rustc --version)"
+    fi
+
+    if check_command wasm-pack; then
+        warning "wasm-pack já está instalado ($(wasm-pack --version)). Pulando."
+    else
+        info "Instalando wasm-pack..."
+        cargo install wasm-pack
+        if ! check_command wasm-pack; then
+            error_exit "Falha ao instalar wasm-pack."
+        fi
+        success "wasm-pack instalado: $(wasm-pack --version)"
+    fi
+}
+
+function install_nginx {
+    if check_command nginx; then
+        warning "Nginx já está instalado ($(nginx -v 2>&1)). Pulando."
+    else
+        info "Instalando Nginx..."
+        apt-get install -y nginx
+        systemctl enable nginx
+        systemctl start nginx
+        success "Nginx instalado e iniciado."
+    fi
+}
+
+function configure_ufw {
+    info "Configurando firewall (UFW)..."
+    # Permitir SSH
+    ufw allow OpenSSH
+    # Permitir porta do FlipApp
+    ufw allow "${REQUIRED_PORT}/tcp"
+    # Habilitar UFW de forma não interativa
+    ufw --force enable
+    success "UFW configurado (SSH e porta ${REQUIRED_PORT} liberados)."
+}
+
+# ------------------- Criação de Usuário e Diretórios -----------------------
+
+function create_flipapp_user_and_dirs {
+    # 1) Criar usuário de sistema se não existir
+    if id -u "${FLIPAPP_USER}" &> /dev/null; then
+        warning "Usuário ${FLIPAPP_USER} já existe. Pulando criação."
+    else
+        info "Criando usuário de sistema '${FLIPAPP_USER}'..."
+        adduser --system --group --home "${FLIPAPP_HOME}" "${FLIPAPP_USER}"
+        success "Usuário ${FLIPAPP_USER} criado."
+    fi
+
+    # 2) Criar diretórios principais
+    info "Criando estrutura de diretórios em ${FLIPAPP_HOME} e ${CONFIG_DIR}..."
+    mkdir -p "${FLIPAPP_HOME}"/{bin,source,config,static,data,examples}
+    chown -R "${FLIPAPP_USER}:${FLIPAPP_USER}" "${FLIPAPP_HOME}"
+    chmod -R 750 "${FLIPAPP_HOME}"
+
+    mkdir -p "${CONFIG_DIR}"
+    chown root:root "${CONFIG_DIR}"
+    chmod 755 "${CONFIG_DIR}"
+    success "Diretórios criados e permissões ajustadas."
+}
+
+# -------------------------- Baixar / Extrair Código ------------------------
+
+function fetch_or_extract_source {
+    # Se GIT_REPO estiver definido, clonar; caso contrário, montar a partir de SOURCE_ZIP
+    if [[ -n "${GIT_REPO}" ]]; then
+        info "Clonando repositório Git em ${SRC_DIR}..."
+        mkdir -p "${SRC_DIR}"
+        chown "${FLIPAPP_USER}:${FLIPAPP_USER}" "${SRC_DIR}"
+        sudo -u "${FLIPAPP_USER}" git clone "${GIT_REPO}" "${SRC_DIR}"
+        success "Repositório clonado."
+    else
+        if [[ -f "${SOURCE_ZIP}" ]]; then
+            info "Extraindo arquivo ZIP ${SOURCE_ZIP} para ${SRC_DIR}..."
+            mkdir -p "${SRC_DIR}"
+            chown "${FLIPAPP_USER}:${FLIPAPP_USER}" "${SRC_DIR}"
+            sudo -u "${FLIPAPP_USER}" unzip -q "${SOURCE_ZIP}" -d "${SRC_DIR}"
+            success "Arquivo ZIP extraído."
+        else
+            warning "Nem GIT_REPO nem SOURCE_ZIP definidos / encontrados."
+            warning "Coloque o código-fonte em ${SRC_DIR} manualmente e pressione [Enter] para continuar..."
+            read -r
+            if [[ ! -d "${SRC_DIR}" ]]; then
+                error_exit "Pasta ${SRC_DIR} não existe. Saindo."
+            fi
+        fi
+    fi
+}
+
+# -------------------------- Compilação do Back-end -------------------------
+
+function build_backend_go {
+    info "Compilando o back-end Go (loglineos)..."
+    # Detectar o caminho do ponto de entrada (ex.: internal/ui/promptpad_app.go)
+    local entry_point
+    if [[ -f "${SRC_DIR}/internal/ui/promptpad_app.go" ]]; then
+        entry_point="${SRC_DIR}/internal/ui/promptpad_app.go"
+    else
+        # Tentar localizar qualquer main.go se a estrutura for diferente
+        entry_point="$(find "${SRC_DIR}" -maxdepth 3 -type f -name 'promptpad_app.go' | head -n 1)"
+        if [[ -z "${entry_point}" ]]; then
+            error_exit "Não encontrei promptpad_app.go em ${SRC_DIR}."
+        fi
+    fi
+
+    # Compilar usando Go
+    cd "$(dirname "${entry_point}")"
+    sudo -u "${FLIPAPP_USER}" go mod tidy
+    sudo -u "${FLIPAPP_USER}" go build -o "${FLIPAPP_HOME}/bin/loglineos" "${entry_point}"
+
+    if [[ ! -x "${FLIPAPP_HOME}/bin/loglineos" ]]; then
+        error_exit "Falha ao compilar loglineos."
+    fi
+    chown "${FLIPAPP_USER}:${FLIPAPP_USER}" "${FLIPAPP_HOME}/bin/loglineos"
+    chmod 750 "${FLIPAPP_HOME}/bin/loglineos"
+    success "Back-end Go compilado: ${FLIPAPP_HOME}/bin/loglineos."
+}
+
+# ------------------------ Configuração de Arquivos --------------------------
+
+function copy_configuration_files {
+    info "Copiando arquivos de configuração para ${CONFIG_DEST}..."
+    if [[ -d "${SRC_DIR}/config" ]]; then
+        cp -r "${SRC_DIR}/config/." "${CONFIG_DEST}/"
+        chown -R "${FLIPAPP_USER}:${FLIPAPP_USER}" "${CONFIG_DEST}"
+        chmod -R 640 "${CONFIG_DEST}"/*
+        success "Arquivos de configuração copiados."
+    else
+        warning "Pasta de configuração ${SRC_DIR}/config não encontrada. Pule esta etapa."
+    fi
+}
+
+# ------------------------- Compilação Front-end WASM ------------------------
+
+function build_frontend_wasm {
+    if [[ -d "${SRC_DIR}/frontend" ]]; then
+        info "Compilando front-end em Rust → WASM..."
+        cd "${SRC_DIR}/frontend"
+        # Garantir que o ambiente Rust esteja carregado
+        export PATH="/root/.cargo/bin:$PATH"
+        sudo -u "${FLIPAPP_USER}" wasm-pack build --release --target web --out-dir "${STATIC_DIR}"
+        chown -R "${FLIPAPP_USER}:${FLIPAPP_USER}" "${STATIC_DIR}"
+        chmod -R 750 "${STATIC_DIR}"
+        success "Front-end compilado (arquivos em ${STATIC_DIR})."
+    else
+        warning "Pasta frontend/ não encontrada em ${SRC_DIR}. Pule a compilação WASM."
+    fi
+}
+
+# ----------------------- Configuração do Serviço systemd --------------------
+
+function create_systemd_service {
+    info "Criando unit file systemd em ${SYSTEMD_UNIT}..."
+    cat > "${SYSTEMD_UNIT}" <<EOF
+[Unit]
+Description=LogLineOS – Motor de Execução de Spans (FlipApp)
+After=network.target
+
+[Service]
+Type=simple
+User=${FLIPAPP_USER}
+WorkingDirectory=${FLIPAPP_HOME}
+ExecStart=${FLIPAPP_HOME}/bin/loglineos --config ${CONFIG_DEST}/logline.yml
+Restart=on-failure
+LimitNOFILE=65536
+
+[Install]
+WantedBy=multi-user.target
+EOF
+
+    chmod 644 "${SYSTEMD_UNIT}"
+    systemctl daemon-reload
+    systemctl enable loglineos
+    success "Serviço systemd criado e habilitado (loglineos)."
+}
+
+# ----------------------------- Proxy Reverso Nginx --------------------------
+
+function configure_nginx_proxy {
+    local site_conf="/etc/nginx/sites-available/flipapp"
+    info "Configurando Nginx para proxy reverso em porta ${REQUIRED_PORT}..."
+
+    cat > "${site_conf}" <<EOF
+server {
+    listen 80;
+    server_name _;  # Pode trocar para domínio real
+
+    location / {
+        proxy_pass         http://127.0.0.1:${REQUIRED_PORT};
+        proxy_http_version 1.1;
+        proxy_set_header   Upgrade \$http_upgrade;
+        proxy_set_header   Connection keep-alive;
+        proxy_set_header   Host \$host;
+        proxy_cache_bypass \$http_upgrade;
+    }
+
+    # Servir arquivos estáticos (WASM/JS)
+    location /static/ {
+        alias ${STATIC_DIR}/;
+    }
+}
+EOF
+
+    ln -sf "${site_conf}" /etc/nginx/sites-enabled/flipapp
+    nginx -t
+    systemctl reload nginx
+    success "Nginx configurado (proxy para FlipApp em http://localhost:${REQUIRED_PORT})."
+}
+
+# --------------------------------- Main --------------------------------------
+
+# 1) Instalar pacotes base
+install_base_packages
+
+# 2) Instalar Go
+install_go
+
+# 3) Instalar Rust + wasm-pack
+install_rust_and_wasm
+
+# 4) Instalar Nginx (opcional)
+install_nginx
+
+# 5) Configurar firewall (UFW)
+configure_ufw
+
+# 6) Criar usuário e diretórios
+create_flipapp_user_and_dirs
+
+# 7) Obter código-fonte (Git ou ZIP ou manual)
+fetch_or_extract_source
+
+# 8) Compilar Back-end Go
+build_backend_go
+
+# 9) Copiar configurações
+copy_configuration_files
+
+# 10) Compilar Front-end WASM (Rust)
+build_frontend_wasm
+
+# 11) Criar serviço systemd
+create_systemd_service
+
+# 12) Configurar Nginx para proxy reverso
+configure_nginx_proxy
+
+echo
+success "======================================================"
+success "   FlipApp/LogLineOS instalado com sucesso!          "
+success "   Acesse: http://<IP-da-VM> ou http://<seu-dominio>   "
+success "======================================================"
+echo
diff --git a/terreno_completo_FINAL/scripts/migrate.sh b/terreno_completo_FINAL/scripts/migrate.sh
new file mode 100644
index 0000000000000000000000000000000000000000..1e1c7f19a5539df714b6153d90a1c5aea7304787
--- /dev/null
+++ b/terreno_completo_FINAL/scripts/migrate.sh
@@ -0,0 +1,4 @@
+#!/bin/bash
+# Migration script for Minicontratos 2030
+psql $DATABASE_URL -f ../../migrations/202501010001_create_users.sql
+# Add more migration files as needed
diff --git a/terreno_completo_FINAL/scripts/rollback.sh b/terreno_completo_FINAL/scripts/rollback.sh
new file mode 100644
index 0000000000000000000000000000000000000000..bec68f7c37a75014f0e24c19a3c153de520c685e
--- /dev/null
+++ b/terreno_completo_FINAL/scripts/rollback.sh
@@ -0,0 +1,3 @@
+#!/bin/bash
+# Rollback script for Minicontratos 2030
+kubectl rollout undo deployment/minicontratos-backend
diff --git a/terreno_completo_FINAL/src/go/ascii_graph.go b/terreno_completo_FINAL/src/go/ascii_graph.go
new file mode 100644
index 0000000000000000000000000000000000000000..7a7e302898ec9d61e161195145e8b8222f08cb41
--- /dev/null
+++ b/terreno_completo_FINAL/src/go/ascii_graph.go
@@ -0,0 +1,69 @@
+// internal/utils/ascii_graph.go
+package utils
+
+import (
+	"fmt"
+	"loglineos/internal/parser"
+	"strings"
+)
+
+// GenerateSimpleASCIIGraph gera uma representação pseudo-ASCII de um workflow.
+// É uma representação linear com indentação e setas, não um grafo visual completo.
+// Para um grafo DOT completo, use planner.BuildGraph.
+func GenerateSimpleASCIIGraph(workflow *parser.Workflow) string {
+	var sb strings.Builder
+	sb.WriteString(fmt.Sprintf("Workflow: %s\n", workflow.Name))
+	sb.WriteString(fmt.Sprintf("Descrição: %s\n", workflow.Description))
+	sb.WriteString("----------------------------------------\n")
+
+	var renderBlocks func(blocks []parser.Block, indent string)
+	renderBlocks = func(blocks []parser.Block, indent string) {
+		for i, block := range blocks {
+			prefix := "├──"
+			if i == len(blocks)-1 && len(block.Tasks) == 0 { // Último nó no nível atual
+				prefix = "└──"
+			}
+
+			// Exibe o bloco principal
+			sb.WriteString(fmt.Sprintf("%s%s %s (%s)\n", indent, prefix, block.ID, block.Type))
+
+			// Adiciona anotações de when, retry, fallback para clareza
+			detailsIndent := indent
+			if i < len(blocks)-1 || len(block.Tasks) > 0 { // Se não for o último, usa linha de continuidade
+				detailsIndent += "│   "
+			} else {
+				detailsIndent += "    "
+			}
+
+			if block.When != "" {
+				sb.WriteString(fmt.Sprintf("%sWhen: %s\n", detailsIndent, block.When))
+			}
+			if block.Retry != nil {
+				sb.WriteString(fmt.Sprintf("%sRetry: max_attempts=%d\n", detailsIndent, block.Retry.MaxAttempts))
+			}
+			if block.Fallback != nil {
+				sb.WriteString(fmt.Sprintf("%sFallback: type=%s\n", detailsIndent, block.Fallback.Type))
+			}
+            if len(block.Hooks) > 0 {
+                sb.WriteString(fmt.Sprintf("%sHooks: %s\n", detailsIndent, strings.Join(block.Hooks, ", ")))
+            }
+
+
+			// Se for um bloco 'loop', renderiza os blocos aninhados recursivamente
+			if block.Type == "loop" && block.Tasks != nil && len(block.Tasks) > 0 {
+				newIndent := indent
+				if i < len(blocks)-1 {
+					newIndent += "│   " // Linha vertical para continuidade
+				} else {
+					newIndent += "    " // Espaço para o último elemento
+				}
+				sb.WriteString(fmt.Sprintf("%s└── Tasks aninhadas:\n", indent))
+				renderBlocks(block.Tasks, newIndent) // Chama recursivamente
+			}
+		}
+	}
+
+	renderBlocks(workflow.Tasks, "")
+	sb.WriteString("----------------------------------------\n")
+	return sb.String()
+}
\ No newline at end of file
diff --git a/terreno_completo_FINAL/src/go/ast.go b/terreno_completo_FINAL/src/go/ast.go
new file mode 100644
index 0000000000000000000000000000000000000000..11f93b77f2e3079f6354749bccb6a6023c0f00a2
--- /dev/null
+++ b/terreno_completo_FINAL/src/go/ast.go
@@ -0,0 +1,112 @@
+// internal/parser/ast.go
+package parser
+
+import "encoding/json"
+
+// Workflow representa a estrutura de alto nível de um arquivo .logline
+type Workflow struct {
+	Name        string  `yaml:"name" json:"name"`
+	Description string  `yaml:"description" json:"description"`
+	Tasks       []Block `yaml:"tasks" json:"tasks"`
+}
+
+// Block representa um bloco genérico da DSL LogLine
+type Block struct {
+	ID      string          `yaml:"id" json:"id"`
+	Type    string          `yaml:"type" json:"type"`
+	Spec    json.RawMessage `yaml:"spec" json:"spec"` // Conteúdo do 'spec' como JSON bruto para unmarshal posterior
+	Fallback *FallbackConfig `yaml:"fallback,omitempty" json:"fallback,omitempty"`
+	Retry    *RetryPolicy    `yaml:"retry,omitempty" json:"retry,omitempty"`
+	When     string          `yaml:"when,omitempty" json:"when,omitempty"`
+	Tasks    []Block         `yaml:"tasks,omitempty" json:"-"` // json:"-" para evitar Marshal recursivo em JSON
+	Hooks    []string        `yaml:"hooks,omitempty" json:"hooks,omitempty"`
+}
+
+// FallbackConfig define a configuração para um bloco de fallback
+type FallbackConfig struct {
+    Type string          `yaml:"type" json:"type"` // "cognitive", "mechanical"
+    Spec json.RawMessage `yaml:"spec" json:"spec"`
+}
+
+// RetryPolicy define a política de retry para um bloco
+type RetryPolicy struct {
+	MaxAttempts   int    `yaml:"max_attempts" json:"max_attempts"`
+	InitialDelay  string `yaml:"initial_delay,omitempty" json:"initial_delay,omitempty"` // e.g., "1s", "500ms"
+	BackoffFactor float64 `yaml:"backoff_factor,omitempty" json:"backoff_factor,omitempty"` // e.g., 2.0 for exponential
+}
+
+type RegisterSpec struct {
+	Entity string                 `yaml:"entity" json:"entity"`
+	ID     string                 `yaml:"id" json:"id"`
+	Data   map[string]interface{} `yaml:"data" json:"data"`
+	Schema string                 `yaml:"schema,omitempty" json:"schema,omitempty"`
+}
+
+type AffairSpec struct {
+	Type     string `yaml:"type" json:"type"`
+	From     string `yaml:"from" json:"from"`
+	To       string `yaml:"to" json:"to"`
+	Note     string `yaml:"note,omitempty" json:"note,omitempty"`
+	Intensity int    `yaml:"intensity,omitempty" json:"intensity,omitempty"`
+}
+
+type ObserveSpec struct {
+	Source  string          `yaml:"source" json:"source"`
+	Pattern string          `yaml:"pattern" json:"pattern"`
+	Action  ObserveAction   `yaml:"action" json:"action"`
+	Interval string         `yaml:"interval,omitempty" json:"interval,omitempty"`
+	OnEvent string          `yaml:"on_event,omitempty" json:"on_event,omitempty"`
+	ListenPath string        `yaml:"listen_path,omitempty" json:"listen_path,omitempty"`
+}
+
+type ObserveAction struct {
+	Type   string `yaml:"type" json:"type"`
+	Target string `yaml:"target" json:"target"`
+	Message string `yaml:"message,omitempty" json:"message,omitempty"`
+}
+
+type CommitSpec struct {
+	Goal     string   `yaml:"goal" json:"goal"`
+	Inputs   []string `yaml:"inputs" json:"inputs"`
+	SignedBy []string `yaml:"signed_by,omitempty" json:"signed_by,omitempty"`
+	Message  string   `yaml:"message,omitempty" json:"message,omitempty"`
+}
+
+type MechanicalSpec struct {
+	Command string   `yaml:"command" json:"command"`
+	Args    []string `yaml:"args,omitempty" json:"args,omitempty"`
+	Timeout string   `yaml:"timeout,omitempty" json:"timeout,omitempty"`
+	Env     map[string]string `yaml:"env,omitempty" json:"env,omitempty"`
+}
+
+type CognitiveSpec struct {
+	LLM CognitiveLLMConfig `yaml:"llm" json:"llm"`
+	ActionOnOutput *CognitiveActionOnOutput `yaml:"action_on_output,omitempty" json:"action_on_output,omitempty"`
+}
+
+type CognitiveLLMConfig struct {
+	Goal  string `yaml:"goal" json:"goal"`
+	Input string `yaml:"input" json:"input"`
+	Trust string `yaml:"trust" json:"trust"`
+	Model string `yaml:"model,omitempty" json:"model,omitempty"`
+}
+
+// CognitiveActionOnOutput define a ação a ser tomada com a saída do LLM
+type CognitiveActionOnOutput struct {
+    TriggerBlockID string `yaml:"trigger_block_id" json:"trigger_block_id"`
+    Condition      string `yaml:"condition,omitempty" json:"condition,omitempty"`
+    MapOutputToInput map[string]string `yaml:"map_output_to_input,omitempty" json:"map_output_to_input,omitempty"`
+}
+
+type LoopSpec struct {
+	Count      *int   `yaml:"count,omitempty" json:"count,omitempty"`
+	Until      string `yaml:"until,omitempty" json:"until,omitempty"`
+	MaxDuration string `yaml:"max_duration,omitempty" json:"max_duration,omitempty"`
+}
+
+type AwaitSpec struct {
+	SpanID    string `yaml:"span_id,omitempty" json:"span_id,omitempty"`
+	TaskID    string `yaml:"task_id,omitempty" json:"task_id,omitempty"`
+	Condition string `yaml:"condition,omitempty" json:"condition,omitempty"`
+	Timeout   string `yaml:"timeout" json:"timeout"`
+}
\ No newline at end of file
diff --git a/terreno_completo_FINAL/src/go/config_loader.go b/terreno_completo_FINAL/src/go/config_loader.go
new file mode 100644
index 0000000000000000000000000000000000000000..9af7cad5d91a949af6fc1f6541a733c8786e9c64
--- /dev/null
+++ b/terreno_completo_FINAL/src/go/config_loader.go
@@ -0,0 +1,96 @@
+// internal/config_loader/config_loader.go
+package config_loader
+
+import (
+	"fmt"
+	"os"
+	"path/filepath"
+
+	"github.com/spf13/viper"
+)
+
+// Config representa a estrutura do arquivo logline.yml
+type Config struct {
+	LLM   LLMConfig   `mapstructure:"llm"`
+	Hooks HooksConfig `mapstructure:"hooks"`
+}
+
+// LLMConfig para configurações de LLM
+type LLMConfig struct {
+	Provider string       `mapstructure:"provider"`
+	OpenAI   OpenAIConfig `mapstructure:"openai"`
+	Ollama   OllamaConfig `mapstructure:"ollama"`
+}
+
+// OpenAIConfig para configurações da OpenAI
+type OpenAIConfig struct {
+	APIKey string `mapstructure:"api_key"`
+	Model  string `mapstructure:"model"`
+}
+
+// OllamaConfig para configurações do Ollama
+type OllamaConfig struct {
+    BaseURL string `mapstructure:"base_url"`
+    Model   string `mapstructure:"model"`
+}
+
+// HooksConfig para todas as definições de hooks
+type HooksConfig map[string][]HookDefinition
+
+// HookDefinition define um hook individual
+type HookDefinition struct {
+	Name    string            `mapstructure:"name"`
+	Type    string            `mapstructure:"type"` // "shell" ou "http"
+	Command string            `mapstructure:"command"`
+	Args    []string          `mapstructure:"args"`
+	URL     string            `mapstructure:"url"`
+	Method  string            `mapstructure:"method"`
+	Headers map[string]string `mapstructure:"headers"`
+	Body    string            `mapstructure:"body"`
+}
+
+// LoadConfig carrega a configuração do arquivo logline.yml
+func LoadConfig(configPath string) (*Config, error) {
+	v := viper.New()
+	v.SetConfigName("logline") // Nome do arquivo sem extensão
+	v.SetConfigType("yaml")    // Tipo de arquivo
+	
+    // Caminhos de busca padrão
+	v.AddConfigPath(configPath) // Caminho passado (ex: ./config)
+	v.AddConfigPath("./config") // Caminho padrão relativo ao executável
+	v.AddConfigPath("$HOME/.loglineos") // Caminho no diretório home do usuário
+	v.AddConfigPath(".") // Diretório atual
+
+	v.SetDefault("llm.provider", "mock") // Default para mock se não configurado
+	v.SetDefault("llm.openai.model", "gpt-3.5-turbo")
+    v.SetDefault("llm.ollama.base_url", "http://localhost:11434")
+    v.SetDefault("llm.ollama.model", "llama3")
+
+	if err := v.ReadInConfig(); err != nil {
+        if _, ok := err.(viper.ConfigFileNotFoundError); ok {
+            fmt.Fprintf(os.Stderr, "Configuração: arquivo logline.yml não encontrado. Usando defaults ou variáveis de ambiente.\n")
+        } else {
+            return nil, fmt.Errorf("config: erro ao ler arquivo de configuração: %w", err)
+        }
+	}
+
+	var config Config
+	if err := v.Unmarshal(&config); err != nil {
+		return nil, fmt.Errorf("config: erro ao fazer unmarshal da configuração: %w", err)
+	}
+
+    // Validações básicas
+    if config.LLM.Provider != "mock" && config.LLM.Provider != "openai" && config.LLM.Provider != "ollama" {
+        return nil, fmt.Errorf("config: provedor LLM '%s' inválido. Use 'openai', 'ollama' ou 'mock'.", config.LLM.Provider)
+    }
+    if config.LLM.Provider == "openai" && config.LLM.OpenAI.APIKey == "" {
+        fmt.Fprintf(os.Stderr, "ATENÇÃO: Provedor OpenAI selecionado, mas 'llm.openai.api_key' não configurado em logline.yml. Chamadas falharão.\n")
+    }
+
+	return &config, nil
+}
+
+// GetDefaultConfigPath retorna o caminho padrão para a configuração
+func GetDefaultConfigPath() string {
+    return filepath.Join(".", "config")
+}
\ No newline at end of file
diff --git a/terreno_completo_FINAL/src/go/errors.go b/terreno_completo_FINAL/src/go/errors.go
new file mode 100644
index 0000000000000000000000000000000000000000..0f1bc20e65a0525258aa63e7cad1304c8dc5caca
--- /dev/null
+++ b/terreno_completo_FINAL/src/go/errors.go
@@ -0,0 +1,80 @@
+// internal/errors/errors.go
+package errors
+
+import "fmt"
+
+// WorkflowError é uma interface para erros relacionados a workflows
+type WorkflowError interface {
+	error
+	WorkflowError() bool
+}
+
+// BlockExecutionError representa um erro ocorrido durante a execução de um bloco.
+type BlockExecutionError struct {
+	BlockID string
+	Type    string
+	Cause   error // A causa original do erro
+	Message string
+}
+
+func (e *BlockExecutionError) Error() string {
+	if e.Cause != nil {
+		return fmt.Sprintf("erro no bloco '%s' (%s): %s, causa: %v", e.BlockID, e.Type, e.Message, e.Cause)
+	}
+	return fmt.Sprintf("erro no bloco '%s' (%s): %s", e.BlockID, e.Type, e.Message)
+}
+
+func (e *BlockExecutionError) Unwrap() error {
+	return e.Cause
+}
+
+func (e *BlockExecutionError) WorkflowError() bool {
+	return true
+}
+
+// ValidationError representa um erro de validação (schema ou semântica).
+type ValidationError struct {
+	FilePath string
+	Errors   []string
+	Cause    error // Causa original, se houver
+}
+
+func (e *ValidationError) Error() string {
+	msg := fmt.Sprintf("erros de validação no arquivo '%s':\n", e.FilePath)
+	for _, err := range e.Errors {
+		msg += fmt.Sprintf("- %s\n", err)
+	}
+	if e.Cause != nil {
+		msg += fmt.Sprintf("causa: %v\n", e.Cause)
+	}
+	return msg
+}
+
+func (e *ValidationError) WorkflowError() bool {
+	return true
+}
+
+// NewBlockExecutionError cria uma nova BlockExecutionError
+func NewBlockExecutionError(blockID, blockType, message string, cause error) *BlockExecutionError {
+	return &BlockExecutionError{
+		BlockID: blockID,
+		Type:    blockType,
+		Cause:   cause,
+		Message: message,
+	}
+}
+
+// NewValidationError cria uma nova ValidationError
+func NewValidationError(filePath string, errors []string, cause error) *ValidationError {
+	return &ValidationError{
+		FilePath: filePath,
+		Errors:   errors,
+		Cause:    cause,
+	}
+}
+
+// IsWorkflowError verifica se um erro é do tipo WorkflowError
+func IsWorkflowError(err error) bool {
+	_, ok := err.(WorkflowError)
+	return ok
+}
\ No newline at end of file
diff --git a/terreno_completo_FINAL/src/go/format.go b/terreno_completo_FINAL/src/go/format.go
new file mode 100644
index 0000000000000000000000000000000000000000..1fc4b466f676b5cb934ef28a67fe34b533b56ab9
--- /dev/null
+++ b/terreno_completo_FINAL/src/go/format.go
@@ -0,0 +1,130 @@
+// internal/utils/format.go
+package utils
+
+import (
+	"encoding/csv"
+	"encoding/json"
+	"fmt"
+	"io"
+	"loglineos/internal/spans"
+	"os"
+	"text/tabwriter" // Para tabelas Markdown
+)
+
+// PrintSpansJSON imprime spans em formato JSON (indentado para leitura)
+func PrintSpansJSON(spans []spans.Span, w io.Writer) error {
+	encoder := json.NewEncoder(w)
+	encoder.SetIndent("", "  ") // Indenta para legibilidade
+	return encoder.Encode(spans)
+}
+
+// PrintSpansMarkdown imprime spans em formato de tabela Markdown
+func PrintSpansMarkdown(spans []spans.Span, w io.Writer) {
+	// Definir as colunas que queremos exibir na tabela Markdown
+	headers := []string{"SPAN_ID", "TASK_ID", "TYPE", "STATUS", "DURATION_MS", "ATTEMPT", "IS_FALLBACK", "TIMESTAMP", "MESSAGE"}
+	
+	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', tabwriter.Debug) // Debug mostra os limites das colunas
+	
+	// Imprime cabeçalho
+	for i, h := range headers {
+		fmt.Fprint(tw, h)
+		if i < len(headers)-1 {
+			fmt.Fprint(tw, "\t")
+		}
+	}
+	fmt.Fprintln(tw)
+
+	// Imprime separador
+	for i := range headers {
+		for j := 0; j < len(headers[i]); j++ {
+			fmt.Fprint(tw, "-")
+		}
+		if i < len(headers)-1 {
+			fmt.Fprint(tw, "\t")
+		}
+	}
+	fmt.Fprintln(tw)
+
+	// Imprime dados
+	for _, span := range spans {
+		fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%d\t%d\t%t\t%s\t%s\n",
+			shortenID(span.SpanID),
+			span.TaskID,
+			span.Type,
+			span.Status,
+			span.DurationMs,
+			span.Attempt,
+			span.IsFallback,
+			span.Timestamp,
+			span.Message,
+		)
+	}
+	tw.Flush() // Garante que tudo seja escrito
+}
+
+// PrintSpansCSV imprime spans em formato CSV
+func PrintSpansCSV(spans []spans.Span, w io.Writer) error {
+	writer := csv.NewWriter(w)
+
+	// Escreve cabeçalho CSV
+	headers := []string{"span_id", "task_id", "type", "status", "timestamp", "duration_ms", "source_workflow", "invoked_by", "attempt", "is_fallback", "message", "output_json", "context_json"}
+	if err := writer.Write(headers); err != nil {
+		return fmt.Errorf("falha ao escrever cabeçalho CSV: %w", err)
+	}
+
+	for _, span := range spans {
+		outputJSON, _ := json.Marshal(span.Output) // Converte output para JSON string
+		contextJSON, _ := json.Marshal(span.Context) // Converte context para JSON string
+		row := []string{
+			span.SpanID,
+			span.TaskID,
+			span.Type,
+			span.Status,
+			span.Timestamp,
+			fmt.Sprintf("%d", span.DurationMs),
+			span.SourceWorkflow,
+			span.InvokedBy,
+			fmt.Sprintf("%d", span.Attempt),
+			fmt.Sprintf("%t", span.IsFallback),
+			span.Message,
+			string(outputJSON),
+			string(contextJSON),
+		}
+		if err := writer.Write(row); err != nil {
+			return fmt.Errorf("falha ao escrever linha CSV: %w", err)
+		}
+	}
+
+	writer.Flush()
+	return nil
+}
+
+// PrintAuditMetrics imprime as métricas coletadas
+func PrintAuditMetrics(metrics spans.AuditMetrics, w io.Writer) {
+	fmt.Fprintln(w, "\n--- Métricas de Auditoria ---")
+	fmt.Fprintf(w, "Total de Spans: %d\n", metrics.TotalSpans)
+	fmt.Fprintln(w, "Spans por Status:")
+	for status, count := range metrics.SpansByStatus {
+		fmt.Fprintf(w, "  - %s: %d\n", status, count)
+	}
+	fmt.Fprintln(w, "Spans por Tipo:")
+	for typ, count := range metrics.SpansByType {
+		fmt.Fprintf(w, "  - %s: %d\n", typ, count)
+	}
+	if len(metrics.AvgDurationMsByType) > 0 {
+		fmt.Fprintln(w, "Duração Média (ms) por Tipo (Apenas Sucessos):")
+		for typ, avg := range metrics.AvgDurationMsByType {
+			fmt.Fprintf(w, "  - %s: %.2fms\n", typ, avg)
+		}
+	}
+	fmt.Fprintf(w, "Duração Total (ms) de Todos os Spans: %dms\n", metrics.TotalDurationMs)
+	fmt.Fprintln(w, "----------------------------")
+}
+
+// shortenID para exibir IDs mais curtos na tabela Markdown
+func shortenID(id string) string {
+	if len(id) > 8 {
+		return id[:8] + "..."
+	}
+	return id
+}
\ No newline at end of file
diff --git a/terreno_completo_FINAL/src/go/hooks.go b/terreno_completo_FINAL/src/go/hooks.go
new file mode 100644
index 0000000000000000000000000000000000000000..5d7c721313fb1fcaf3dc3af8a33b79b87545dbc2
--- /dev/null
+++ b/terreno_completo_FINAL/src/go/hooks.go
@@ -0,0 +1,211 @@
+// internal/hooks/hooks.go
+package hooks
+
+import (
+	"bytes"
+	"encoding/json"
+	"fmt"
+	"io"
+	"loglineos/internal/config_loader"
+	"loglineos/internal/spans" // Para SpanID e Writer
+	"net/http"
+	"os"
+	"os/exec"
+	"strings"
+	"text/template"
+	"time"
+
+	"github.com/google/uuid" // Para SpanID
+)
+
+// HookContext é o contexto passado para o template do hook e para o log
+type HookContext struct {
+	WorkflowName string                 `json:"workflow_name"`
+	BlockID      string                 `json:"block_id"`
+	Type         string                 `json:"type"`
+	Status       string                 `json:"status"`
+	SpanID       string                 `json:"span_id,omitempty"`
+	Timestamp    string                 `json:"timestamp,omitempty"`
+	DurationMs   int64                  `json:"duration_ms,omitempty"`
+	Message      string                 `json:"message"`
+	Error        string                 `json:"error,omitempty"` // Se for um hook de erro
+	Output       map[string]interface{} `json:"output,omitempty"` // Output do span
+	HookName     string                 `json:"hook_name"`       // Nome do hook que está sendo executado
+}
+
+// HooksManager gerencia a execução dos hooks
+type HooksManager struct {
+	HooksConfig config_loader.HooksConfig
+	SpanWriter *spans.JSONLWriter // Para spans de hooks (injetado via context)
+}
+
+// NewHooksManager cria uma nova instância de HooksManager
+func NewHooksManager(hooksConfig config_loader.HooksConfig, spanWriter *spans.JSONLWriter) *HooksManager {
+	return &HooksManager{
+		HooksConfig: hooksConfig,
+		SpanWriter: spanWriter,
+	}
+}
+
+// ExecuteHook executa um hook específico, se definido
+func (hm *HooksManager) ExecuteHook(hookKey string, ctx *HookContext) {
+	definitions, ok := hm.HooksConfig[hookKey]
+	if !ok || len(definitions) == 0 {
+		return // Nenhum hook definido para esta chave
+	}
+
+	fmt.Printf("   [HOOKS] Disparando hook '%s' para bloco '%s'...\n", hookKey, ctx.BlockID)
+
+	for _, def := range definitions {
+		hookSpan := spans.Span{
+			SpanID:         uuid.New().String(),
+			TaskID:         fmt.Sprintf("hook_%s_%s", ctx.BlockID, def.Name),
+			Type:           "hook",
+			Status:         "running",
+			Timestamp:      time.Now().UTC().Format(time.RFC3339),
+			SourceWorkflow: ctx.WorkflowName,
+			InvokedBy:      "logline_hook_manager",
+			Message:        fmt.Sprintf("Executando hook '%s' (tipo: %s)", def.Name, def.Type),
+			Context:        map[string]interface{}{"hook_key": hookKey, "hook_name": def.Name, "hook_type": def.Type},
+			Inputs:         map[string]interface{}{"block_context": ctx}, // Passa o contexto do bloco para o input do span
+		}
+        ctx.HookName = def.Name // Define o nome do hook no contexto para o template
+
+		startTime := time.Now()
+		var err error
+		var hookOutput map[string]interface{}
+
+		switch def.Type {
+		case "shell":
+			hookOutput, err = hm.executeShellHook(def, ctx)
+		case "http":
+			hookOutput, err = hm.executeHTTPHook(def, ctx)
+		default:
+			err = fmt.Errorf("tipo de hook '%s' não suportado", def.Type)
+		}
+
+		hookSpan.DurationMs = time.Since(startTime).Milliseconds()
+		hookSpan.Output = hookOutput
+		if err != nil {
+			hookSpan.Status = "error"
+			hookSpan.Message = fmt.Sprintf("Hook '%s' falhou: %v", def.Name, err)
+			fmt.Printf("   ❌ [HOOK %s] Falha no hook '%s': %v\n", hookKey, def.Name, err)
+		} else {
+			hookSpan.Status = "success"
+			hookSpan.Message = fmt.Sprintf("Hook '%s' executado com sucesso.", def.Name)
+			fmt.Printf("   🟢 [HOOK %s] Sucesso no hook '%s'.\n", hookKey, def.Name)
+		}
+		
+		if hm.SpanWriter != nil {
+			hm.SpanWriter.WriteSpan(hookSpan)
+		} else {
+            fmt.Fprintf(os.Stderr, "Hooks: SpanWriter não inicializado para hook '%s'. Não será auditado.\n", def.Name)
+        }
+	}
+}
+
+// executeShellHook executa um comando shell definido no hook
+func (hm *HooksManager) executeShellHook(def config_loader.HookDefinition, ctx *HookContext) (map[string]interface{}, error) {
+	command, err := renderTemplate(def.Command, ctx)
+	if err != nil {
+		return nil, fmt.Errorf("erro ao renderizar comando do hook: %w", err)
+	}
+
+	var args []string
+	for _, arg := range def.Args {
+		renderedArg, err := renderTemplate(arg, ctx)
+		if err != nil {
+			return nil, fmt.Errorf("erro ao renderizar argumento '%s' do hook: %w", arg, err)
+		}
+		args = append(args, renderedArg)
+	}
+
+	cmd := exec.Command(command, args...)
+	output, err := cmd.CombinedOutput()
+	if err != nil {
+		return map[string]interface{}{"stdout": string(output)}, fmt.Errorf("comando shell falhou: %w, output: %s", err, string(output))
+	}
+	return map[string]interface{}{"stdout": string(output)}, nil
+}
+
+// executeHTTPHook executa uma requisição HTTP definida no hook
+func (hm *HooksManager) executeHTTPHook(def config_loader.HookDefinition, ctx *HookContext) (map[string]interface{}, error) {
+	url, err := renderTemplate(def.URL, ctx)
+	if err != nil {
+		return nil, fmt.Errorf("erro ao renderizar URL do hook: %w", err)
+	}
+	method := def.Method
+	if method == "" {
+		method = "GET"
+	}
+
+	var bodyReader io.Reader
+	if def.Body != "" {
+		renderedBody, err := renderTemplate(def.Body, ctx)
+		if err != nil {
+			return nil, fmt.Errorf("erro ao renderizar corpo do hook: %w", err)
+		}
+		bodyReader = strings.NewReader(renderedBody)
+	}
+
+	req, err := http.NewRequest(method, url, bodyReader)
+	if err != nil {
+		return nil, fmt.Errorf("falha ao criar requisição HTTP: %w", err)
+	}
+
+	for k, v := range def.Headers {
+		renderedValue, err := renderTemplate(v, ctx)
+		if err != nil {
+			return nil, fmt.Errorf("erro ao renderizar cabeçalho '%s' do hook: %w", k, err)
+		}
+		req.Header.Set(k, renderedValue)
+	}
+
+	client := &http.Client{Timeout: 10 * time.Second}
+	resp, err := client.Do(req)
+	if err != nil {
+		return nil, fmt.Errorf("falha ao executar requisição HTTP: %w", err)
+	}
+	defer resp.Body.Close()
+
+	respBody, err := ioutil.ReadAll(resp.Body)
+	if err != nil {
+		return nil, fmt.Errorf("falha ao ler corpo da resposta HTTP: %w", err)
+	}
+
+	return map[string]interface{}{
+		"status_code": resp.StatusCode,
+		"response_body": string(respBody),
+	}, nil
+}
+
+// renderTemplate processa uma string com dados do HookContext
+func renderTemplate(text string, ctx *HookContext) (string, error) {
+	// Usar Funções JSON para permitir acesso a campos de Output complexos
+	funcMap := template.FuncMap{
+		"json": func(v interface{}) (string, error) {
+			b, err := json.Marshal(v)
+			if err != nil {
+				return "", err
+			}
+			return string(b), nil
+		},
+	}
+	
+	// Permite usar um template sem tags de ação se o texto não contiver '{{'
+	// Isso evita erros de parsing para strings simples que não são templates.
+	if !strings.Contains(text, "{{") && !strings.Contains(text, "}}") {
+		return text, nil
+	}
+
+	tmpl, err := template.New("hook_template").Funcs(funcMap).Parse(text)
+	if err != nil {
+		return "", fmt.Errorf("erro ao parsear template: %w", err)
+	}
+	var buf bytes.Buffer
+	err = tmpl.Execute(&buf, ctx)
+	if err != nil {
+		return "", fmt.Errorf("erro ao executar template: %w", err)
+	}
+	return buf.String(), nil
+}
\ No newline at end of file
diff --git a/terreno_completo_FINAL/src/go/llm_integrator.go b/terreno_completo_FINAL/src/go/llm_integrator.go
new file mode 100644
index 0000000000000000000000000000000000000000..e69fe0e3e946b64a6ca3fed1c2b2fe6989387190
--- /dev/null
+++ b/terreno_completo_FINAL/src/go/llm_integrator.go
@@ -0,0 +1,150 @@
+// internal/llm_integrator/llm_integrator.go
+package llm_integrator
+
+import (
+	"context"
+	"fmt"
+	"loglineos/internal/config_loader"
+	"time"
+
+	"github.com/sashabaranov/go-openai" // Para OpenAI
+	// Para uso futuro de JSON Schema, se a resposta do LLM precisar ser validada
+	// "github.com/sashabaranov/go-openai/jsonschema"
+	// "github.com/sashabaranov/go-openai/types" 
+
+    "github.com/ollama/ollama-go/ollama" // Para Ollama
+    "github.com/ollama/ollama-go/api"    // Para Ollama API
+
+)
+
+// LLMResponse representa a resposta de um LLM
+type LLMResponse struct {
+	Output      string                 `json:"output"`
+	TrustLevel  string                 `json:"trust_level"`
+	Provider    string                 `json:"provider"`
+	Model       string                 `json:"model"`
+	TokensUsed  int                    `json:"tokens_used,omitempty"`
+	RawResponse map[string]interface{} `json:"raw_response,omitempty"` // Para inspeção
+}
+
+// LLMIntegrator interface para diferentes provedores de LLM
+type LLMIntegrator interface {
+	Generate(ctx context.Context, prompt string, trustLevel string, modelOverride string) (*LLMResponse, error)
+}
+
+// NewLLMIntegrator cria uma instância do integrador LLM com base na configuração
+func NewLLMIntegrator(config *config_loader.LLMConfig) (LLMIntegrator, error) {
+	switch config.Provider {
+	case "openai":
+		if config.OpenAI.APIKey == "" {
+			return nil, fmt.Errorf("LLM: chave de API OpenAI não configurada")
+		}
+		return &OpenAIIntegrator{
+			client: openai.NewClient(config.OpenAI.APIKey),
+			defaultModel: config.OpenAI.Model,
+		}, nil
+	case "ollama":
+        if config.Ollama.BaseURL == "" {
+            return nil, fmt.Errorf("LLM: URL base do Ollama não configurada")
+        }
+        return &OllamaIntegrator{
+            client: ollama.NewClient(config.Ollama.BaseURL),
+            defaultModel: config.Ollama.Model,
+        }, nil
+	case "mock":
+		return &MockLLMIntegrator{}, nil
+	default:
+		return nil, fmt.Errorf("LLM: provedor '%s' não suportado", config.Provider)
+	}
+}
+
+// OpenAIIntegrator implementa LLMIntegrator para OpenAI
+type OpenAIIntegrator struct {
+	client       *openai.Client
+	defaultModel string
+}
+
+func (o *OpenAIIntegrator) Generate(ctx context.Context, prompt string, trustLevel string, modelOverride string) (*LLMResponse, error) {
+	model := o.defaultModel
+	if modelOverride != "" {
+		model = modelOverride
+	}
+
+	messages := []openai.ChatCompletionMessage{
+		{Role: openai.ChatMessageRoleSystem, Content: "Você é um assistente LogLineOS. Gero respostas concisas e úteis."},
+		{Role: openai.ChatMessageRoleUser, Content: prompt},
+	}
+
+	req := openai.ChatCompletionRequest{
+		Model:    model,
+		Messages: messages,
+		MaxTokens: 500, // Limite para evitar custos excessivos na demo
+	}
+
+	resp, err := o.client.CreateChatCompletion(ctx, req)
+	if err != nil {
+		return nil, fmt.Errorf("LLM OpenAI: falha ao gerar resposta: %w", err)
+	}
+
+	if len(resp.Choices) == 0 {
+		return nil, fmt.Errorf("LLM OpenAI: nenhuma escolha na resposta")
+	}
+
+	return &LLMResponse{
+		Output:      resp.Choices[0].Message.Content,
+		TrustLevel:  trustLevel,
+		Provider:    "openai",
+		Model:       model,
+		TokensUsed:  resp.Usage.TotalTokens,
+		RawResponse: map[string]interface{}{"usage": resp.Usage, "id": resp.ID},
+	}, nil
+}
+
+// OllamaIntegrator implementa LLMIntegrator para Ollama
+type OllamaIntegrator struct {
+    client *ollama.Client
+    defaultModel string
+}
+
+func (o *OllamaIntegrator) Generate(ctx context.Context, prompt string, trustLevel string, modelOverride string) (*LLMResponse, error) {
+    model := o.defaultModel
+    if modelOverride != "" {
+        model = modelOverride
+    }
+
+    req := &api.GenerateRequest{
+        Model:  model,
+        Prompt: prompt,
+        Stream: false, // Não usar stream para esta demo
+    }
+
+    resp, err := o.client.Generate(ctx, req)
+    if err != nil {
+        return nil, fmt.Errorf("LLM Ollama: falha ao gerar resposta: %w", err)
+    }
+
+    return &LLMResponse{
+        Output:      resp.Response,
+        TrustLevel:  trustLevel,
+        Provider:    "ollama",
+        Model:       model,
+        TokensUsed:  0, // Ollama API response doesn't directly provide tokens
+        RawResponse: map[string]interface{}{"done": resp.Done, "context": resp.Context},
+    }, nil
+}
+
+// MockLLMIntegrator para simulação
+type MockLLMIntegrator struct{}
+
+func (m *MockLLMIntegrator) Generate(ctx context.Context, prompt string, trustLevel string, modelOverride string) (*LLMResponse, error) {
+	// Simula um delay
+	time.Sleep(500 * time.Millisecond)
+	return &LLMResponse{
+		Output:      fmt.Sprintf("MockLLM: Resposta simulada para '%s' (Trust: %s)", prompt, trustLevel),
+		TrustLevel:  trustLevel,
+		Provider:    "mock",
+		Model:       "mock-model-v1",
+		TokensUsed:  100,
+		RawResponse: map[string]interface{}{"simulated": true},
+	}, nil
+}
\ No newline at end of file
diff --git a/terreno_completo_FINAL/src/go/monitor.go b/terreno_completo_FINAL/src/go/monitor.go
new file mode 100644
index 0000000000000000000000000000000000000000..bbb547f12d491555b7f6c7e63cbd2c7b5fba92fa
--- /dev/null
+++ b/terreno_completo_FINAL/src/go/monitor.go
@@ -0,0 +1,214 @@
+// internal/monitor/monitor.go
+package monitor
+
+import (
+	"context"
+	"encoding/json"
+	"fmt"
+	"io/ioutil"
+	errorsInternal "loglineos/internal/errors"
+	"loglineos/internal/executor"
+	"loglineos/internal/parser"
+	"loglineos/internal/spans"
+	"loglineos/internal/utils" // Para WebhookPayload
+	"os"
+	"regexp"
+	"time"
+
+	"github.com/google/uuid"
+)
+
+// Monitor representa um monitor contínuo para um bloco 'observe'
+type Monitor struct {
+	Block          parser.Block
+	ExecutionContext *executor.ExecutionContext
+	CancelFunc     context.CancelFunc
+	LastObservedContent string // Para `interval`
+	PatternRegex   *regexp.Regexp
+}
+
+// NewMonitor cria um novo monitor para um bloco 'observe'
+func NewMonitor(block parser.Block, execCtx *executor.ExecutionContext) (*Monitor, error) {
+	if block.Type != "observe" {
+		return nil, fmt.Errorf("monitor: apenas blocos 'observe' podem ser monitorados continuamente")
+	}
+
+	var spec parser.ObserveSpec
+	if err := json.Unmarshal(block.Spec, &spec); err != nil {
+		return nil, fmt.Errorf("monitor: erro ao parsear spec de observe para monitor: %w", err)
+	}
+
+	if spec.Interval == "" && spec.OnEvent == "" {
+		return nil, fmt.Errorf("monitor: bloco observe '%s' não tem 'interval' nem 'on_event' para monitoramento contínuo", block.ID)
+	}
+
+	re, err := regexp.Compile(spec.Pattern)
+	if err != nil {
+		return nil, fmt.Errorf("monitor: padrão regex inválido para bloco '%s': %w", block.ID, err)
+	}
+
+	if execCtx.OutputDir == "" {
+		execCtx.OutputDir = "./data"
+	}
+	if execCtx.SpanWriter == nil {
+		spanWriter, err := spans.NewJSONLWriter(execCtx.OutputDir + "/spans.jsonl")
+		if err != nil {
+			return nil, fmt.Errorf("monitor: falha ao criar span writer para monitor: %w", err)
+		}
+		execCtx.SpanWriter = spanWriter
+	}
+
+
+	return &Monitor{
+		Block:          block,
+		ExecutionContext: execCtx,
+		PatternRegex:   re,
+	}, nil
+}
+
+// Start inicia o monitoramento contínuo (polling) ou espera eventos
+func (m *Monitor) Start(ctx context.Context) {
+	var spec parser.ObserveSpec
+	json.Unmarshal(m.Block.Spec, &spec)
+
+	if spec.Interval != "" {
+		duration, err := time.ParseDuration(spec.Interval)
+		if err != nil {
+			fmt.Fprintf(os.Stderr, "Monitor '%s': erro ao parsear intervalo '%s': %v. Monitor não será iniciado.\n", m.Block.ID, spec.Interval, err)
+			return
+		}
+
+		ticker := time.NewTicker(duration)
+		defer ticker.Stop()
+
+		fmt.Printf("✅ Monitor '%s' (Tipo: %s) iniciado com intervalo: %v\n", m.Block.ID, m.Block.Type, duration)
+
+		for {
+			select {
+			case <-ctx.Done():
+				fmt.Printf("⚪ Monitor '%s' parado (contexto cancelado).\n", m.Block.ID)
+				return
+			case <-ticker.C:
+				m.performObservation(spec, nil) // Chamada de polling
+			}
+		}
+	} else if spec.OnEvent != "" {
+		fmt.Printf("✅ Monitor '%s' (Tipo: %s) esperando eventos de '%s'.\n", m.Block.ID, m.Block.Type, spec.OnEvent)
+		// Esta goroutine apenas espera pelo cancelamento, os eventos são tratados via HandleWebhookEvent
+		<-ctx.Done()
+		fmt.Printf("⚪ Monitor '%s' parado (contexto cancelado).\n", m.Block.ID)
+		return
+	}
+}
+
+// HandleWebhookEvent é chamado pelo WebhookServer quando um evento relevante é recebido
+func (m *Monitor) HandleWebhookEvent(payload utils.WebhookPayload) {
+	var spec parser.ObserveSpec
+	json.Unmarshal(m.Block.Spec, &spec) // Erro já tratado em NewMonitor
+
+	fmt.Printf("   [MONITOR %s] Recebido evento de webhook em '%s'...\n", m.Block.ID, payload.Path)
+
+	// Converte o corpo do webhook para string para a regex
+	payloadContent := string(payload.Body)
+	if payloadContent == "" && len(payload.Query) > 0 { // Se não tem corpo, pode ser via query params
+		// Converte query params para string para a regex
+		queryStr := ""
+		for k, v := range payload.Query {
+			queryStr += fmt.Sprintf("%s=%s&", k, strings.Join(v, ","))
+		}
+		payloadContent = queryStr
+	}
+
+	m.performObservation(spec, &payloadContent) // Passa o conteúdo do evento
+}
+
+
+// performObservation executa a lógica de observação para um tick/evento
+func (m *Monitor) performObservation(spec parser.ObserveSpec, eventContent *string) {
+	span := spans.Span{
+		SpanID:         uuid.New().String(),
+		TaskID:         m.Block.ID,
+		Type:           m.Block.Type,
+		Timestamp:      time.Now().UTC().Format(time.RFC3339),
+		SourceWorkflow: m.ExecutionContext.WorkflowName,
+		InvokedBy:      "logline_monitor",
+		Context:        map[string]interface{}{"interval": spec.Interval, "on_event": spec.OnEvent},
+	}
+
+	var currentContent string
+	if eventContent != nil { // Se veio de um evento (webhook)
+		currentContent = *eventContent
+		span.Context["event_type"] = spec.OnEvent
+		span.Context["event_source"] = spec.ListenPath
+		span.Output["raw_event_content"] = currentContent
+	} else { // Se veio de polling (interval)
+		currentContent = readSimulatedSource(spec.Source) // Lê conteúdo simulado
+	}
+	
+	isMatched := m.PatternRegex.MatchString(currentContent)
+	contentChanged := (currentContent != m.LastObservedContent)
+	m.LastObservedContent = currentContent // Atualiza último conteúdo (para polling)
+
+	span.Output["source"] = spec.Source
+	span.Output["pattern"] = spec.Pattern
+	span.Output["matched_pattern"] = isMatched
+	span.Output["content_changed"] = contentChanged
+	span.Output["current_content_hash"] = fmt.Sprintf("%x", HashString(currentContent))
+
+
+	if isMatched {
+		span.Status = "success"
+		span.Message = fmt.Sprintf("Padrão '%s' ENCONTRADO em '%s'. Conteúdo mudou: %t", spec.Pattern, spec.Source, contentChanged)
+		fmt.Printf("   %s [MONITOR %s] %s\n", getStatusEmoji(span.Status), m.Block.ID, span.Message)
+		// TODO: Disparar action (alert, trigger, commit) - Isso é complexo e precisa de um executor de actions
+		// Por enquanto, apenas o span é gerado. A lógica de trigger block é no executor, não aqui.
+	} else {
+		span.Status = "warning"
+		span.Message = fmt.Sprintf("Padrão '%s' NÃO ENCONTRADO em '%s'. Conteúdo mudou: %t", spec.Pattern, spec.Source, contentChanged)
+		fmt.Printf("   %s [MONITOR %s] %s\n", getStatusEmoji(span.Status), m.Block.ID, span.Message)
+	}
+
+	if err := m.ExecutionContext.SpanWriter.WriteSpan(span); err != nil {
+		fmt.Fprintf(os.Stderr, "Monitor '%s': erro ao escrever span: %v\n", m.Block.ID, err)
+	}
+	if m.ExecutionContext.SpanStream != nil {
+		m.ExecutionContext.SpanStream <- span
+	}
+}
+
+// readSimulatedSource simula a leitura de um arquivo/fonte externa
+var simulatedContentCounter int
+func readSimulatedSource(source string) string {
+	simulatedContentCounter++
+	if simulatedContentCounter%3 == 0 {
+		return fmt.Sprintf("Conteúdo simulado de %s com PADRÃO_ALVO (mudou em %s)", source, time.Now().Format("15:04:05"))
+	}
+	return fmt.Sprintf("Conteúdo simulado de %s sem padrao (stable %s)", source, time.Now().Format("15:04:05"))
+}
+
+// HashString retorna um hash simples de uma string
+func HashString(s string) uint32 {
+	h := uint32(0)
+	for i := 0; i < len(s); i++ {
+		h = (h << 5) - h + uint32(s[i])
+	}
+	return h
+}
+
+// getStatusEmoji (copiado do executor, para uso independente aqui)
+func getStatusEmoji(status string) string {
+	switch status {
+	case "success":
+		return "🟢"
+	case "error":
+		return "❌"
+	case "warning":
+		return "🟡"
+	case "draft":
+		return "📝"
+	case "skipped":
+		return "⚪"
+	default:
+		return "⚪"
+	}
+}
\ No newline at end of file
diff --git a/terreno_completo_FINAL/src/go/parser.go b/terreno_completo_FINAL/src/go/parser.go
new file mode 100644
index 0000000000000000000000000000000000000000..f59d47cf2d02ebc6b63a8c4414ec53ed62605df3
--- /dev/null
+++ b/terreno_completo_FINAL/src/go/parser.go
@@ -0,0 +1,48 @@
+// internal/parser/parser.go
+package parser
+
+import (
+	"fmt"
+	errorsInternal "loglineos/internal/errors"
+
+	"github.com/go-json-schema/go-json-schema-validator" // Para validação de schema
+	"gopkg.in/yaml.v3" // Para parsing de YAML
+)
+
+// ParseWorkflow lê um arquivo YAML/JSON e o converte para a estrutura Workflow
+func ParseWorkflow(content []byte) (*Workflow, error) {
+	var workflow Workflow
+	err := yaml.Unmarshal(content, &workflow)
+	if err != nil {
+		return nil, fmt.Errorf("erro ao fazer unmarshal do workflow: %w", err)
+	}
+
+	// TODO: Aqui entraria a validação semântica mais complexa
+	// por exemplo, verificar se IDs referenciados em 'inputs' existem,
+	// ou se condições 'when' são parsáveis (já feito no executor/planner).
+
+	return &workflow, nil
+}
+
+// ValidateContent valida o conteúdo do workflow contra o schema JSON
+// schemaContent: o conteúdo JSON do schema (logline.schema.json)
+// workflowContent: o conteúdo YAML/JSON do workflow a ser validado
+func ValidateContent(filePath string, schemaContent, workflowContent []byte) *errorsInternal.ValidationError {
+	v := gojsonschema.NewValidator()
+	schemaLoader := gojsonschema.NewBytesLoader(schemaContent)
+	
+	result, err := v.Validate(schemaLoader, gojsonschema.NewBytesLoader(workflowContent))
+	if err != nil {
+		return errorsInternal.NewValidationError(filePath, []string{"Erro interno ao processar schema ou documento"}, err)
+	}
+
+	if result.Valid() {
+		return nil // Retorna nil se válido
+	} else {
+		errMsgs := make([]string, len(result.Errors()))
+		for i, desc := range result.Errors() {
+			errMsgs[i] = desc.String() // string() dá uma descrição mais legível
+		}
+		return errorsInternal.NewValidationError(filePath, errMsgs, nil)
+	}
+}
\ No newline at end of file
diff --git a/terreno_completo_FINAL/src/go/planner.go b/terreno_completo_FINAL/src/go/planner.go
new file mode 100644
index 0000000000000000000000000000000000000000..38ead0cf8961e29b6ac24bb563a5767f92b878c3
--- /dev/null
+++ b/terreno_completo_FINAL/src/go/planner.go
@@ -0,0 +1,212 @@
+// internal/planner/planner.go
+package planner
+
+import (
+	"encoding/json" // Para unmarshal de spec
+	"fmt"
+	"loglineos/internal/parser"
+	"os"
+	"strings"
+
+	"gonum.org/v1/gonum/graph/encoding/dot"
+	"gonum.org/v1/gonum/graph/simple"
+)
+
+// Node representa um bloco no grafo do workflow
+type Node struct {
+	id   int64	// ID interno para o grafo
+	Block parser.Block
+	DOTID string // ID para o formato DOT (o ID do bloco)
+}
+
+// ID retorna o ID interno do nó
+func (n Node) ID() int64 {
+	return n.id
+}
+
+// DOTID retorna o ID para o formato DOT
+func (n Node) DOTID() string {
+	return n.DOTID
+}
+
+// Attributes de nó para Graphviz (cor, label, etc.)
+func (n Node) Attributes() []dot.Attribute {
+	label := fmt.Sprintf("%s\\n(%s)", n.Block.ID, n.Block.Type)
+	color := "black"
+	shape := "box"
+	fillcolor := "lightgrey" // Default fill color
+
+	switch n.Block.Type {
+	case "register":
+		color = "blue"
+		fillcolor = "lightblue"
+	case "affair":
+		color = "darkgreen"
+		fillcolor = "lightgreen"
+	case "observe":
+		color = "orange"
+		fillcolor = "lightyellow"
+	case "commit":
+		color = "purple"
+		fillcolor = "lavender"
+	case "mechanical":
+		color = "gray"
+		fillcolor = "lightgray"
+	case "cognitive":
+		color = "brown"
+		fillcolor = "peachpuff"
+	case "loop": // Loop blocks get a special shape
+		shape = "doubleoctagon"
+		color = "darkred"
+		fillcolor = "pink"
+	case "await": // Await blocks get a special shape
+		shape = "parallelogram"
+		color = "darkblue"
+		fillcolor = "lightcyan"
+	}
+
+	return []dot.Attribute{
+		{Key: "label", Value: dot.Quote(label)},
+		{Key: "color", Value: dot.Quote(color)},
+		{Key: "shape", Value: dot.Quote(shape)},
+		{Key: "style", Value: dot.Quote("filled")},
+		{Key: "fillcolor", Value: dot.Quote(fillcolor)},
+	}
+}
+
+// BuildGraph constrói um grafo Gonum a partir de um workflow LogLine
+func BuildGraph(workflow *parser.Workflow) (*simple.DirectedGraph, error) {
+	g := simple.NewDirectedGraph()
+	blockNodes := make(map[string]*Node)
+	var nodeID int64 = 0
+
+	// Função auxiliar para adicionar blocos recursivamente e criar nós no grafo
+	var addBlocksRecursively func(blocks []parser.Block)
+	addBlocksRecursively = func(blocks []parser.Block) {
+		for _, block := range blocks {
+			if _, exists := blockNodes[block.ID]; !exists {
+				node := &Node{id: nodeID, Block: block, DOTID: block.ID}
+				g.AddNode(node)
+				blockNodes[block.ID] = node
+				nodeID++
+			}
+
+			if block.Type == "loop" && block.Tasks != nil && len(block.Tasks) > 0 {
+				addBlocksRecursively(block.Tasks) // Adiciona blocos aninhados do loop
+			}
+		}
+	}
+	addBlocksRecursively(workflow.Tasks)
+
+
+	// Função auxiliar para processar dependências recursivamente e adicionar arestas
+	var processDependenciesRecursively func(blocks []parser.Block)
+	processDependenciesRecursively = func(blocks []parser.Block) {
+		for i, block := range blocks {
+			currentNode := blockNodes[block.ID]
+
+			// 1. Dependência sequencial implícita
+			if i > 0 {
+				prevBlock := blocks[i-1]
+				prevNode := blockNodes[prevBlock.ID]
+				if !g.HasEdgeFromTo(prevNode.ID(), currentNode.ID()) {
+					g.SetEdge(g.NewEdge(prevNode, currentNode))
+				}
+			}
+
+			// 2. Dependências explícitas (when, inputs)
+			if block.When != "" {
+				// Para 'when', tentamos inferir dependências por TaskID referenciados
+				// Esta é uma inferência básica; expressões complexas exigiriam análise de AST da expressão
+				for prevID, prevNode := range blockNodes {
+					if strings.Contains(block.When, prevID) {
+						if !g.HasEdgeFromTo(prevNode.ID(), currentNode.ID()) {
+							g.SetEdge(g.NewEdge(prevNode, currentNode))
+						}
+					}
+				}
+			}
+
+			// 3. Dependência para 'await' blocks
+			if block.Type == "await" {
+				var awaitSpec parser.AwaitSpec
+				if err := json.Unmarshal(block.Spec, &awaitSpec); err == nil {
+					targetID := awaitSpec.TaskID
+					if targetID == "" {
+						// Se await for por span_id, e não task_id, é mais difícil
+						// inferir no planner sem histórico. Pulamos por agora.
+					}
+					if targetID != "" {
+						if targetNode, ok := blockNodes[targetID]; ok {
+							if !g.HasEdgeFromTo(targetNode.ID(), currentNode.ID()) {
+								g.SetEdge(g.NewEdge(targetNode, currentNode))
+							}
+						}
+					}
+				}
+			}
+
+			// 4. Para blocos 'commit', eles dependem de seus 'inputs'
+			if block.Type == "commit" {
+				var commitSpec parser.CommitSpec
+				if err := json.Unmarshal(block.Spec, &commitSpec); err == nil {
+					for _, inputID := range commitSpec.Inputs {
+						if inputNode, ok := blockNodes[inputID]; ok {
+							if !g.HasEdgeFromTo(inputNode.ID(), currentNode.ID()) {
+								g.SetEdge(g.NewEdge(inputNode, currentNode))
+							}
+						}
+					}
+				}
+			}
+
+            // 5. Para blocos 'cognitive' com 'action_on_output'
+            if block.Type == "cognitive" {
+                var cognitiveSpec parser.CognitiveSpec
+                if err := json.Unmarshal(block.Spec, &cognitiveSpec); err == nil && cognitiveSpec.ActionOnOutput != nil {
+                    targetID := cognitiveSpec.ActionOnOutput.TriggerBlockID
+                    if targetNode, ok := blockNodes[targetID]; ok {
+                        if !g.HasEdgeFromTo(currentNode.ID(), targetNode.ID()) {
+                            g.SetEdge(g.NewEdge(currentNode, targetNode)) // A aresta vai do cognitive para o bloco disparado
+                        }
+                    }
+                }
+            }
+
+
+			// 6. Se for um bloco 'loop', seus blocos aninhados também têm dependências
+			if block.Type == "loop" && block.Tasks != nil && len(block.Tasks) > 0 {
+				loopNode := blockNodes[block.ID]
+				for _, innerBlock := range block.Tasks {
+					innerNode := blockNodes[innerBlock.ID]
+					// O loop block "contém" os inner blocks, mas também há dependência explícita
+					// do loop para o primeiro bloco aninhado para visualização de fluxo.
+					if !g.HasEdgeFromTo(loopNode.ID(), innerNode.ID()) {
+						g.SetEdge(g.NewEdge(loopNode, innerNode))
+					}
+				}
+				processDependenciesRecursively(block.Tasks) // Chamada recursiva para dependências aninhadas
+			}
+		}
+	}
+
+	// Começa o processamento de dependências da lista de blocos raiz
+	processDependenciesRecursively(workflow.Tasks)
+
+	return g, nil
+}
+
+// GenerateDOTFile gera um arquivo .dot a partir do grafo
+func GenerateDOTFile(graph *simple.DirectedGraph, filePath string) error {
+	file, err := os.Create(filePath)
+	if err != nil {
+		return fmt.Errorf("planner: falha ao criar arquivo DOT: %w", err)
+	}
+	defer file.Close()
+
+	_, err = dot.Marshal(graph, graph.Name(), "", "", file)
+	if err != nil {
+		return fmt.Errorf("planner: falha ao Marshal grafo para DOT: %w", err)
+	}
+	return nil
+}
\ No newline at end of file
diff --git a/terreno_completo_FINAL/src/go/promptpad_app.go b/terreno_completo_FINAL/src/go/promptpad_app.go
new file mode 100644
index 0000000000000000000000000000000000000000..950dee165789d8837fb45dc45c77fea8c95e4e06
--- /dev/null
+++ b/terreno_completo_FINAL/src/go/promptpad_app.go
@@ -0,0 +1,646 @@
+// internal/ui/promptpad_app.go
+package ui
+
+import (
+	"bytes"
+	"context"
+	"fmt"
+	"io/ioutil"
+	"loglineos/cmd"
+	errorsInternal "loglineos/internal/errors"
+	"loglineos/internal/executor"
+	"loglineos/internal/parser"
+	"loglineos/internal/spans"
+	"loglineos/internal/utils"
+	"os"
+	"path/filepath"
+	"regexp"
+	"strings"
+	"time"
+
+	"github.com/rivo/tview"
+)
+
+// PromptPadApp contém a lógica da aplicação TUI
+type PromptPadApp struct {
+	app           *tview.Application
+	pages         *tview.Pages
+	inputField    *tview.InputField
+	outputView    *tview.TextView
+	workflowEditor *tview.TextArea
+	graphView     *tview.TextView
+	statusText    *tview.TextView
+	
+	currentWorkflowPath string
+	loglineSchemaContent []byte
+	spansOutputDir      string
+	
+	spanStream    chan spans.Span
+}
+
+// NewPromptPadApp cria e inicializa a aplicação PromptPad
+func NewPromptPadApp(schemaContent []byte, spansOutputDir string) *PromptPadApp {
+	app := tview.NewApplication()
+
+	// Componentes da UI
+	inputField := tview.NewInputField().
+		SetLabel(">>> ").
+		SetFieldWidth(0).
+		SetBorder(true).
+		SetTitle(" Comando ")
+
+	outputView := tview.NewTextView().
+		SetDynamicColors(true).
+		SetScrollable(true).
+		SetWordWrap(true).
+		SetMaxBuffer(1024 * 50)
+
+	workflowEditor := tview.NewTextArea().
+		SetPlaceholder("Cole seu workflow LogLine aqui ou use 'open <arquivo.logline>'").
+		SetBorder(true).
+		SetTitle(" Workflow Editor (F2 Salvar, F3 Validar, F4 Grafo) ")
+
+	graphView := tview.NewTextView().
+		SetDynamicColors(true).
+		SetScrollable(true).
+		SetWordWrap(false).
+		SetBorder(true).
+		SetTitle(" Workflow Grafo ")
+
+	statusText := tview.NewTextView().
+		SetDynamicColors(true).
+		SetText("Bem-vindo ao LogLineOS PromptPad!").SetTextAlign(tview.AlignCenter)
+
+	// Layout principal
+	mainLayout := tview.NewFlex().SetDirection(tview.FlexRow).
+		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
+			AddItem(workflowEditor, 0, 3, false).
+			AddItem(graphView, 0, 1, false),
+			0, 4, false).
+		AddItem(inputField, 1, 1, true).
+		AddItem(statusText, 1, 1, false)
+
+	pages := tview.NewPages().
+		AddPage("main", mainLayout, true, true)
+
+	p := &PromptPadApp{
+		app:           app,
+		pages:         pages,
+		inputField:    inputField,
+		outputView:    outputView,
+		workflowEditor: workflowEditor,
+		graphView:     graphView,
+		statusText:    statusText,
+		
+		loglineSchemaContent: schemaContent,
+		spansOutputDir:      spansOutputDir,
+		
+		spanStream:    make(chan spans.Span, 100),
+	}
+
+	inputField.SetDoneFunc(func(key tview.Key) {
+		if key == tview.KeyEnter {
+			command := strings.TrimSpace(inputField.GetText())
+			inputField.SetText("")
+			p.processCommand(command)
+		}
+	})
+
+	app.SetInputCapture(func(event *tview.Event) *tview.Event {
+		if p.app.GetFocus() == p.workflowEditor {
+			switch event.Key() {
+			case tview.KeyF2: // F2 para Salvar
+				p.saveWorkflow()
+				return nil
+			case tview.KeyF3: // F3 para Validar
+				p.validateCurrentWorkflow()
+				return nil
+			case tview.KeyF4: // F4 para Gerar Grafo
+				p.generateWorkflowGraph()
+				return nil
+			}
+		}
+		return event
+	})
+
+	// Syntax highlighting e validação contínua (debounce)
+	validationTimer := time.NewTimer(time.Millisecond * 200)
+	validationTimer.Stop()
+	go func() {
+		for range validationTimer.C {
+			p.syntaxHighlightAndValidate()
+		}
+	}()
+
+	workflowEditor.SetChangedFunc(func() {
+		validationTimer.Stop()
+		validationTimer.Reset(time.Millisecond * 200)
+	})
+
+	go p.listenForSpans()
+
+	return p
+}
+
+// Run inicia a aplicação TUI
+func (p *PromptPadApp) Run() error {
+	return p.app.SetRoot(p.pages, true).Run()
+}
+
+// processCommand interpreta e executa comandos do PromptPad
+func (p *PromptPadApp) processCommand(command string) {
+	p.outputView.Clear()
+	p.statusText.SetText(fmt.Sprintf("[blue]Executando: '%s'...", command)).SetTextAlign(tview.AlignCenter)
+	p.app.Draw()
+
+	parts := strings.Fields(command)
+	if len(parts) == 0 {
+		p.statusText.SetText("[red]Comando vazio.").SetTextAlign(tview.AlignCenter)
+		return
+	}
+
+	ctx := context.Background()
+
+	switch parts[0] {
+	case "open":
+		if len(parts) < 2 {
+			p.outputView.SetText("[red]Erro: 'open' requer um caminho de arquivo.")
+			p.statusText.SetText("[red]Uso: open <caminho_do_arquivo.logline>").SetTextAlign(tview.AlignCenter)
+			return
+		}
+		filePath := parts[1]
+		p.openWorkflow(filePath)
+	case "run":
+		workflowPath := p.currentWorkflowPath
+		if len(parts) >= 2 {
+			workflowPath = parts[1]
+		}
+		if workflowPath == "" {
+			p.outputView.SetText("[red]Erro: Nenhum workflow carregado para 'run'. Use 'open' ou 'run <caminho>'.")
+			p.statusText.SetText("[red]Falha ao executar.").SetTextAlign(tview.AlignCenter)
+			return
+		}
+		p.saveAndExecute(ctx, workflowPath, false)
+	case "simulate":
+		workflowPath := p.currentWorkflowPath
+		if len(parts) >= 2 {
+			workflowPath = parts[1]
+		}
+		if workflowPath == "" {
+			p.outputView.SetText("[red]Erro: Nenhum workflow carregado para 'simulate'. Use 'open' ou 'simulate <caminho>'.")
+			p.statusText.SetText("[red]Falha ao simular.").SetTextAlign(tview.AlignCenter)
+			return
+		}
+		p.saveAndExecute(ctx, workflowPath, true)
+	case "audit":
+		args := parts[1:]
+		var buf bytes.Buffer
+		err := cmd.RunAuditCommand("", strings.Join(args, " "), "", "markdown", "", false, &buf)
+		p.outputView.WriteString(buf.String())
+		if err != nil {
+			p.outputView.WriteString(fmt.Sprintf("[red]❌ Erro na auditoria: %v\n", err))
+			p.statusText.SetText("[red]Auditoria falhou.").SetTextAlign(tview.AlignCenter)
+		} else {
+			p.statusText.SetText("[green]Auditoria concluída.").SetTextAlign(tview.AlignCenter)
+		}
+		p.outputView.ScrollToEnd()
+	case "commit":
+		if p.currentWorkflowPath == "" {
+			p.outputView.SetText("[red]Erro: Nenhum workflow aberto para commit.")
+			p.statusText.SetText("[red]Falha ao commit.").SetTextAlign(tview.AlignCenter)
+			return
+		}
+		goal := "Manual commit via PromptPad"
+		message := "Commit manual do estado atual."
+		inputs := []string{}
+		
+		go func() {
+			err := cmd.RunCommitCommand(goal, inputs, []string{"PromptPad_User"}, message, p.spansOutputDir, false)
+			p.app.QueueUpdateDraw(func() {
+				if err != nil {
+					p.outputView.WriteString(fmt.Sprintf("[red]❌ Erro no commit: %v\n", err))
+					p.statusText.SetText("[red]Commit falhou.").SetTextAlign(tview.AlignCenter)
+				} else {
+					p.outputView.WriteString("[green]✅ Commit realizado com sucesso!\n")
+					p.statusText.SetText("[green]Commit concluído.").SetTextAlign(tview.AlignCenter)
+				}
+				p.outputView.ScrollToEnd()
+			})
+		}()
+	case "revert":
+		p.outputView.Clear()
+		p.outputView.WriteString("[yellow]Comando 'revert' invocado.\n")
+		p.outputView.WriteString("Para reverter para um estado anterior, você precisará usar o Git.\n")
+		p.outputView.WriteString("1. Use 'logline audit' para encontrar o 'snapshot_id' do ponto que deseja reverter.\n")
+		p.outputView.WriteString("2. Use o hash do commit Git associado ao snapshot (se disponível no metadata.json).\n")
+		p.outputView.WriteString("3. No seu terminal, fora do PromptPad, execute:\n")
+		p.outputView.WriteString("   [blue]git checkout <hash_do_commit_do_snapshot>\n")
+		p.outputView.WriteString("   [blue]OU: git checkout $(logline audit --filter 'snapshot_id == \"<id_do_snapshot>\"' --format json | jq -r '.[0].git_hash')\n")
+		p.outputView.WriteString("[yellow]AVISO: Reverter um repositório Git pode sobrescrever arquivos. Salve suas mudanças antes!\n")
+		p.statusText.SetText("[yellow]Instruções de reversão exibidas.").SetTextAlign(tview.AlignCenter)
+		p.outputView.ScrollToEnd()
+	case "replay":
+		go p.replaySpansInTUI()
+	case "?": // <-- NOVO: AI Assistant
+		p.handleAIQuery(strings.Join(parts[1:], " "))
+	case "clear":
+		p.outputView.Clear()
+		p.statusText.SetText("Saída limpa.").SetTextAlign(tview.AlignCenter)
+	case "exit":
+		p.app.Stop()
+	default:
+		p.outputView.SetText(fmt.Sprintf("[red]Comando desconhecido: '%s'", parts[0]))
+		p.statusText.SetText("[red]Comando inválido.").SetTextAlign(tview.AlignCenter)
+	}
+}
+
+// saveAndExecute é uma função auxiliar para salvar e executar/simular
+func (p *PromptPadApp) saveAndExecute(ctx context.Context, workflowPath string, dryRun bool) {
+	if p.currentWorkflowPath != "" && p.workflowEditor.GetText() != "" {
+		if err := ioutil.WriteFile(p.currentWorkflowPath, []byte(p.workflowEditor.GetText()), 0644); err != nil {
+			p.outputView.WriteString(fmt.Sprintf("[red]Erro ao salvar antes de rodar: %v\n", err))
+			p.statusText.SetText("[red]Falha ao salvar antes de rodar.").SetTextAlign(tview.AlignCenter)
+			return
+		}
+	}
+
+	p.outputView.Clear()
+	action := "Executando"
+	if dryRun { action = "Simulando" }
+	p.outputView.WriteString(fmt.Sprintf("%s workflow: %s\n", action, workflowPath))
+	p.statusText.SetText(fmt.Sprintf("[blue]%s '%s'...", action, workflowPath)).SetTextAlign(tview.AlignCenter)
+
+	close(p.spanStream)
+	p.spanStream = make(chan spans.Span, 100)
+	go p.listenForSpans()
+
+	workflowContent, err := ioutil.ReadFile(workflowPath)
+	if err != nil {
+		p.outputView.WriteString(fmt.Sprintf("[red]Erro ao ler arquivo: %v\n", err))
+		p.statusText.SetText(fmt.Sprintf("[red]%s '%s' falhou.", action, workflowPath)).SetTextAlign(tview.AlignCenter)
+		return
+	}
+	validationError := parser.ValidateContent(workflowPath, p.loglineSchemaContent, workflowContent)
+	if validationError != nil {
+		p.outputView.WriteString(fmt.Sprintf("[red]Workflow tem erros de validação:\n%s\n", validationError.Error()))
+		p.statusText.SetText(fmt.Sprintf("[red]%s '%s' falhou devido a erros de validação.", action, workflowPath)).SetTextAlign(tview.AlignCenter)
+		return
+	}
+
+	workflow, err := parser.ParseWorkflow(workflowContent)
+	if err != nil {
+		p.outputView.WriteString(fmt.Sprintf("[red]Erro ao parsear workflow: %v\n", err))
+		p.statusText.SetText(fmt.Sprintf("[red]%s '%s' falhou.", action, workflowPath)).SetTextAlign(tview.AlignCenter)
+		return
+	}
+
+	execContext := &executor.ExecutionContext{
+		DryRun:     dryRun,
+		SpanWriter: nil,
+		SpanStream: p.spanStream,
+		OutputDir:  p.spansOutputDir,
+		EntityStore: make(map[string]executor.EntityRecord),
+		WorkflowName: workflow.Name,
+		HooksManager: cmd.GetHooksManager(), // Pega o HooksManager global
+		LLMIntegrator: cmd.GetLLMIntegrator(), // Pega o LLMIntegrator global
+	}
+
+	// NOVO: Preenche o SpanWriter no HooksManager global para esta execução
+	if execContext.HooksManager != nil {
+		spanWriter, err := spans.NewJSONLWriter(execContext.OutputDir + "/spans.jsonl")
+		if err != nil {
+			p.outputView.WriteString(fmt.Sprintf("[red]Erro ao inicializar span writer para hooks: %v\n", err))
+		} else {
+			execContext.HooksManager.SpanWriter = spanWriter
+			defer func() {
+				if err := spanWriter.Close(); err != nil {
+					fmt.Fprintf(os.Stderr, "Erro ao fechar SpanWriter de hooks: %v\n", err)
+				}
+			}()
+		}
+	}
+
+
+	err = executor.ExecuteWorkflow(workflow, execContext)
+	
+	p.app.QueueUpdateDraw(func() {
+		if err != nil {
+			p.outputView.WriteString(fmt.Sprintf("[red]❌ Erro na %s: %v\n", strings.ToLower(action), err))
+			p.statusText.SetText(fmt.Sprintf("[red]%s '%s' falhou.", action, workflowPath)).SetTextAlign(tview.AlignCenter)
+		} else {
+			p.outputView.WriteString(fmt.Sprintf("[green]✅ %s concluída.\n", action))
+			p.statusText.SetText(fmt.Sprintf("[green]%s '%s' concluída com sucesso.", action, workflowPath)).SetTextAlign(tview.AlignCenter)
+		}
+		p.outputView.ScrollToEnd()
+	})
+}
+
+
+// openWorkflow carrega e exibe o conteúdo de um workflow no editor, e o valida
+func (p *PromptPadApp) openWorkflow(filePath string) {
+	content, err := cmd.LoadFileContent(filePath)
+	if err != nil {
+		p.outputView.SetText(fmt.Sprintf("[red]Erro ao abrir workflow: %v", err))
+		p.statusText.SetText("[red]Falha ao carregar workflow.").SetTextAlign(tview.AlignCenter)
+		return
+	}
+	p.workflowEditor.SetText(string(content), false)
+	p.currentWorkflowPath = filePath
+	p.statusText.SetText(fmt.Sprintf("[green]Workflow '%s' carregado.", filePath)).SetTextAlign(tview.AlignCenter)
+	p.syntaxHighlightAndValidate()
+	p.generateWorkflowGraph()
+}
+
+// saveWorkflow salva o conteúdo atual do editor no arquivo
+func (p *PromptPadApp) saveWorkflow() {
+	if p.currentWorkflowPath == "" {
+		p.statusText.SetText("[red]Nenhum workflow aberto para salvar.").SetTextAlign(tview.AlignCenter)
+		return
+	}
+	content := []byte(p.workflowEditor.GetText())
+	err := ioutil.WriteFile(p.currentWorkflowPath, content, 0644)
+	if err != nil {
+		p.statusText.SetText(fmt.Sprintf("[red]Erro ao salvar: %v", err)).SetTextAlign(tview.AlignCenter)
+		return
+	}
+	p.statusText.SetText(fmt.Sprintf("[green]Workflow '%s' salvo com sucesso.", p.currentWorkflowPath)).SetTextAlign(tview.AlignCenter)
+	p.syntaxHighlightAndValidate()
+	p.generateWorkflowGraph()
+}
+
+// validateCurrentWorkflow valida o conteúdo atual do editor em tempo real
+func (p *PromptPadApp) validateCurrentWorkflow() {
+	if p.currentWorkflowPath == "" {
+		p.statusText.SetText("[yellow]Nenhum workflow carregado para validação.").SetTextAlign(tview.AlignCenter)
+		return
+	}
+
+	content := []byte(p.workflowEditor.GetText())
+	validationError := parser.ValidateContent(p.currentWorkflowPath, p.loglineSchemaContent, content)
+
+	p.app.QueueUpdateDraw(func() {
+		p.outputView.Clear()
+		if validationError != nil {
+			p.outputView.SetTitle("[red] Erros de Validação! ")
+			p.outputView.WriteString(fmt.Sprintf("[red]❌ %s\n", validationError.Error()))
+			p.statusText.SetText(fmt.Sprintf("[red]Workflow inválido.").SetTextAlign(tview.AlignCenter))
+		} else {
+			p.outputView.SetTitle("[green] Validação ")
+			p.outputView.WriteString("[green]✅ Workflow válido de acordo com o schema.\n")
+			p.statusText.SetText("[green]Workflow válido.").SetTextAlign(tview.AlignCenter)
+		}
+		p.outputView.ScrollToBeginning()
+	})
+}
+
+
+// syntaxHighlightAndValidate combina highlight de sintaxe e validação (refatorado para ser mais rápido)
+func (p *PromptPadApp) syntaxHighlightAndValidate() {
+	if p.currentWorkflowPath == "" {
+		return
+	}
+
+	content := p.workflowEditor.GetText() // Obtém o texto UMA VEZ
+
+	// Aplicar highlight
+	highlightedText := p.applySyntaxHighlight(content)
+
+	// Validar (agora usa o novo método de parser.ValidateContent)
+	validationError := parser.ValidateContent(p.currentWorkflowPath, p.loglineSchemaContent, []byte(content))
+
+	p.app.QueueUpdateDraw(func() {
+		// Atualiza o texto do editor com o highlight ANTES de exibir erros
+		p.workflowEditor.SetText(highlightedText, false)
+
+		// Exibir erros/status da validação
+		p.outputView.Clear()
+		if validationError != nil {
+			p.outputView.SetTitle("[red] Erros de Validação! ")
+			p.outputView.WriteString(fmt.Sprintf("[red]❌ %s\n", validationError.Error()))
+			p.statusText.SetText(fmt.Sprintf("[red]Workflow inválido.").SetTextAlign(tview.AlignCenter))
+		} else {
+			p.outputView.SetTitle("[green] Validação ")
+			p.outputView.WriteString("[green]✅ Workflow válido de acordo com o schema.\n")
+			p.statusText.SetText("[green]Workflow válido.").SetTextAlign(tview.AlignCenter)
+		}
+		p.outputView.ScrollToBeginning()
+	})
+}
+
+
+// applySyntaxHighlight aplica um highlight de sintaxe básico ao texto do workflow
+func (p *PromptPadApp) applySyntaxHighlight(text string) string {
+	var sb strings.Builder
+	lines := strings.Split(text, "\n")
+	
+	// Regexes para tipos e palavras-chave
+	keywordRegex := regexp.MustCompile(`\b(name|description|tasks|id|type|spec|fallback|retry|when|inputs|command|args|entity|data|goal|from|to|note|pattern|source|action|message|signed_by|max_attempts|interval|on_event|count|until|timeout|llm|model|trust|input|listen_path|trigger_block_id|condition|map_output_to_input|hooks)\b:`)
+	typeRegex := regexp.MustCompile(`type:\s*["']?(register|affair|observe|commit|mechanical|cognitive|loop|await|alert|trigger)["']?`)
+	stringValRegex := regexp.MustCompile(`(".*?")`)
+	numberValRegex := regexp.MustCompile(`\b(\d+(\.\d+)?)\b`)
+
+	for _, line := range lines {
+		// Aplicar cores em ordem, do mais geral para o mais específico
+		// Primeiro, a cor padrão (branca ou do tema)
+		coloredLine := "[white]" + line + "[white]" // Reset cor no final da linha
+
+		// Tipos de blocos (amarelo)
+		coloredLine = typeRegex.ReplaceAllString(coloredLine, "type: [yellow]$1[white]")
+
+		// Palavras-chave (laranja)
+		coloredLine = keywordRegex.ReplaceAllString(coloredLine, "[orange]$1[white]:")
+
+		// Valores de string (verde)
+		coloredLine = stringValRegex.ReplaceAllString(coloredLine, "[green]$1[white]")
+		
+		// Valores numéricos (azul claro) - pode conflitar com strings, ordem importa
+		coloredLine = numberValRegex.ReplaceAllString(coloredLine, "[aqua]$1[white]")
+
+		sb.WriteString(coloredLine)
+		sb.WriteString("\n")
+	}
+	return sb.String()
+}
+
+
+// generateWorkflowGraph gera e exibe a representação ASCII do grafo no painel `graphView`
+func (p *PromptPadApp) generateWorkflowGraph() {
+	if p.currentWorkflowPath == "" {
+		p.graphView.SetText("[yellow]Nenhum workflow aberto para gerar grafo.").SetTextAlign(tview.AlignCenter)
+		return
+	}
+
+	content := []byte(p.workflowEditor.GetText())
+	workflow, err := parser.ParseWorkflow(content)
+	if err != nil {
+		p.app.QueueUpdateDraw(func() {
+			p.graphView.SetText(fmt.Sprintf("[red]Erro ao parsear workflow para grafo: %v", err)).SetTextAlign(tview.AlignCenter)
+			p.statusText.SetText("[red]Erro ao gerar grafo.").SetTextAlign(tview.AlignCenter)
+		})
+		return
+	}
+
+	asciiGraph := utils.GenerateSimpleASCIIGraph(workflow)
+	p.app.QueueUpdateDraw(func() {
+		p.graphView.SetText(asciiGraph)
+		p.graphView.ScrollToBeginning()
+		p.graphView.SetTitle(fmt.Sprintf(" Workflow Grafo (%s) ", filepath.Base(p.currentWorkflowPath)))
+		p.statusText.SetText("[green]Grafo do workflow gerado.").SetTextAlign(tview.AlignCenter)
+	})
+}
+
+
+// replaySpansInTUI reproduz os spans em câmera lenta
+func (p *PromptPadApp) replaySpansInTUI() {
+	p.outputView.Clear()
+	p.outputView.WriteString("[yellow]Iniciando Replay de Spans... (Pressione Ctrl+C para parar fora do PromptPad se for stuck)\n")
+	p.statusText.SetText("[blue]Replay ativo.').").SetTextAlign(tview.AlignCenter)
+
+	spanFilePath := filepath.Join(p.spansOutputDir, "spans.jsonl")
+	spansToReplay, err := spans.ReadAllSpans(spanFilePath)
+	if err != nil {
+		p.outputView.WriteString(fmt.Sprintf("[red]Erro ao ler spans para replay: %v\n", err))
+		p.statusText.SetText("[red]Replay falhou.").SetTextAlign(tview.AlignCenter)
+		return
+	}
+	if len(spansToReplay) == 0 {
+		p.outputView.WriteString("[yellow]Nenhum span encontrado para replay. Execute um workflow primeiro.\n")
+		p.statusText.SetText("[yellow]Replay sem spans.").SetTextAlign(tview.AlignCenter)
+		return
+	}
+
+	spansToReplay, _ = spans.SortSpans(spansToReplay, "timestamp", "asc")
+
+	go func() {
+		for i, span := range spansToReplay {
+			p.app.QueueUpdateDraw(func() {
+				statusEmoji := getStatusEmoji(span.Status)
+				outputLine := fmt.Sprintf("[%d/%d] %s [%s] %s: %s (Duração: %dms",
+					i+1, len(spansToReplay), statusEmoji, span.Status, span.TaskID, span.Message, span.DurationMs)
+				if span.Attempt > 1 {
+					outputLine += fmt.Sprintf(", Tentativa: %d", span.Attempt)
+				}
+				if span.IsFallback {
+					outputLine += ", Fallback!"
+				}
+				outputLine += ")\n"
+				p.outputView.WriteString(outputLine)
+				if span.Status == "error" {
+					p.outputView.WriteString(fmt.Sprintf("  Detalhes do Erro: %v\n", span.Output))
+				}
+				p.outputView.ScrollToEnd()
+			})
+			time.Sleep(500 * time.Millisecond) // Pausa para cada span
+		}
+		p.app.QueueUpdateDraw(func() {
+			p.outputView.WriteString("[green]Replay concluído.\n")
+			p.statusText.SetText("[green]Replay concluído.").SetTextAlign(tview.AlignCenter)
+		})
+	}()
+}
+
+// handleAIQuery processa uma pergunta para o assistente de IA
+func (p *PromptPadApp) handleAIQuery(query string) {
+	p.outputView.Clear()
+	p.outputView.WriteString(fmt.Sprintf("[blue]🧠 Assistente IA: '%s'\n", query))
+	p.statusText.SetText("[blue]Consultando assistente IA...").SetTextAlign(tview.AlignCenter)
+
+	llmIntegrator := cmd.GetLLMIntegrator()
+	if llmIntegrator == nil {
+		p.outputView.WriteString("[red]Erro: LLM Integrator não está disponível. Verifique sua configuração de logline.yml.\n")
+		p.statusText.SetText("[red]Assistente IA indisponível.").SetTextAlign(tview.AlignCenter)
+		return
+	}
+
+	// Constrói um prompt baseado na query e no contexto atual (spans recentes, workflow)
+	contextualPrompt := p.buildContextualPromptForAI(query)
+
+	go func() {
+		resp, err := llmIntegrator.Generate(context.Background(), contextualPrompt, "medium", "")
+		p.app.QueueUpdateDraw(func() {
+			if err != nil {
+				p.outputView.WriteString(fmt.Sprintf("[red]❌ Erro ao consultar IA: %v\n", err))
+				p.statusText.SetText("[red]Consulta IA falhou.").SetTextAlign(tview.AlignCenter)
+			} else {
+				p.outputView.WriteString(fmt.Sprintf("[green]✅ Resposta da IA:\n%s\n", resp.Output))
+				p.statusView.SetText("[green]Resposta da IA recebida.").SetTextAlign(tview.AlignCenter)
+			}
+			p.outputView.ScrollToEnd()
+		})
+	}()
+}
+
+// buildContextualPromptForAI constrói um prompt com base na query e no contexto do PromptPad
+func (p *PromptPadApp) buildContextualPromptForAI(query string) string {
+	var sb strings.Builder
+	sb.WriteString("Você é um assistente LogLineOS no PromptPad. Responda à pergunta do usuário. Inclua informações sobre o workflow atual e spans recentes, se relevantes.\n\n")
+	sb.WriteString("PERGUNTA DO USUÁRIO: " + query + "\n\n")
+
+	// Adicionar o workflow atual, se houver
+	if p.currentWorkflowPath != "" {
+		sb.WriteString("WORKFLOW ATUAL:\n")
+		sb.WriteString("Caminho: " + p.currentWorkflowPath + "\n")
+		sb.WriteString("Conteúdo:\n```yaml\n" + p.workflowEditor.GetText() + "\n```\n\n")
+	}
+
+	// Adicionar spans recentes (últimos 10, por exemplo)
+	spanFilePath := filepath.Join(p.spansOutputDir, "spans.jsonl")
+	recentSpans, err := spans.ReadAllSpans(spanFilePath)
+	if err == nil && len(recentSpans) > 0 {
+		sb.WriteString("SPANS RECENTES (Últimos 10):\n")
+		if len(recentSpans) > 10 {
+			recentSpans = recentSpans[len(recentSpans)-10:] // Pega os 10 últimos
+		}
+		for _, s := range recentSpans {
+			spanJSON, _ := json.Marshal(s) // Converte span para JSON
+			sb.WriteString(fmt.Sprintf("- %s\n", string(spanJSON)))
+		}
+		sb.WriteString("\n")
+	}
+
+	sb.WriteString("Por favor, seja conciso e direto ao ponto. Use o contexto fornecido para sua resposta.")
+	return sb.String()
+}
+
+
+// listenForSpans escuta o canal de spans e os adiciona ao outputView
+func (p *PromptPadApp) listenForSpans() {
+	for span := range p.spanStream {
+		p.app.QueueUpdateDraw(func() {
+			statusEmoji := getStatusEmoji(span.Status)
+			outputLine := fmt.Sprintf("%s [%s] %s: %s (Duração: %dms",
+				statusEmoji, span.Status, span.TaskID, span.Message, span.DurationMs)
+			if span.Attempt > 1 {
+				outputLine += fmt.Sprintf(", Tentativa: %d", span.Attempt)
+			}
+			if span.IsFallback {
+				outputLine += ", Fallback!"
+			}
+			outputLine += ")\n"
+			p.outputView.WriteString(outputLine)
+			if span.Status == "error" {
+				p.outputView.WriteString(fmt.Sprintf("  Detalhes do Erro: %v\n", span.Output))
+			}
+			p.outputView.ScrollToEnd()
+		})
+	}
+}
+
+
+// Funções utilitárias (repetidas do executor para este pacote, idealmente centralizadas)
+func getStatusEmoji(status string) string {
+	switch status {
+	case "success":
+		return "🟢"
+	case "error":
+		return "❌"
+	case "warning":
+		return "🟡"
+	case "draft":
+		return "📝"
+	case "skipped":
+		return "⚪"
+	default:
+		return "⚪"
+	}
+}
\ No newline at end of file
diff --git a/terreno_completo_FINAL/src/go/snapshotter.go b/terreno_completo_FINAL/src/go/snapshotter.go
new file mode 100644
index 0000000000000000000000000000000000000000..97603b4ee8d4044422fb3e15b94bc2690c82dcad
--- /dev/null
+++ b/terreno_completo_FINAL/src/go/snapshotter.go
@@ -0,0 +1,215 @@
+// internal/snapshotter/snapshotter.go
+package snapshotter
+
+import (
+	"encoding/json"
+	"fmt"
+	"io/ioutil"
+	"log"
+	errorsInternal "loglineos/internal/errors"
+	"loglineos/internal/executor" // Importa o executor para EntityRecord
+	"loglineos/internal/spans"
+	"os"
+	"os/exec"
+	"path/filepath"
+	"strings"
+	"time"
+
+	"github.com/go-git/go-git/v5"
+	"github.com/go-git/go-git/v5/plumbing/object"
+	"github.com/google/uuid"
+)
+
+// SnapshotMetadata contém informações sobre o snapshot
+type SnapshotMetadata struct {
+	SnapshotID  string                 `json:"snapshot_id"`
+	Timestamp   string                 `json:"timestamp"`
+	Goal        string                 `json:"goal"`
+	Message     string                 `json:"message,omitempty"`
+	Inputs      []string               `json:"inputs_committed,omitempty"`
+	SignedBy    []string               `json:"signed_by,omitempty"`
+	SpansCount  int                    `json:"spans_count"`
+	EntitiesCount int                  `json:"entities_count"`
+	GitHash     string                 `json:"git_hash,omitempty"`
+	GitBranch   string                 `json:"git_branch,omitempty"`
+	Signed      bool                   `json:"signed,omitempty"`
+	Signature   string                 `json:"signature_file,omitempty"` // Caminho para o arquivo .asc ou base64 da assinatura
+	Context     map[string]interface{} `json:"context,omitempty"`
+}
+
+// CreateSnapshot gera um novo snapshot do estado atual
+func CreateSnapshot(
+	goal string,
+	inputs []string,
+	signedBy []string,
+	message string,
+	globalSpansFilePath string, // Caminho para o arquivo spans.jsonl principal
+	entityStore map[string]executor.EntityRecord, // EntityStore do executor
+	baseOutputDir string, // Diretório base onde os snapshots serão criados (e.g., ./data)
+	signGPG bool, // Flag para assinar com GPG
+) (string, error) {
+	
+	// 1. Gerar ID e caminho para o snapshot
+	snapshotID := uuid.New().String()
+	snapshotDir := filepath.Join(baseOutputDir, "snapshots", snapshotID)
+
+	if err := os.MkdirAll(snapshotDir, 0755); err != nil {
+		return "", fmt.Errorf("snapshotter: falha ao criar diretório de snapshot '%s': %w", snapshotDir, err)
+	}
+	fmt.Printf("   Criado diretório de snapshot: %s\n", snapshotDir)
+
+	// 2. Coletar e filtrar spans relevantes
+	allSpans, err := spans.ReadAllSpans(globalSpansFilePath)
+	if err != nil {
+		return "", fmt.Errorf("snapshotter: falha ao ler spans globais: %w", err)
+	}
+
+	filteredSpans := []spans.Span{}
+	
+	inputMap := make(map[string]bool)
+	for _, id := range inputs {
+		inputMap[id] = true
+	}
+
+	for _, s := range allSpans {
+		if inputMap[s.TaskID] || inputMap[s.SpanID] {
+			filteredSpans = append(filteredSpans, s)
+		}
+	}
+	
+	spansSnapshotPath := filepath.Join(snapshotDir, "spans_snapshot.jsonl")
+	writer, err := spans.NewJSONLWriter(spansSnapshotPath)
+	if err != nil {
+		return "", fmt.Errorf("snapshotter: falha ao criar escritor de spans para snapshot: %w", err)
+	}
+	defer writer.Close()
+	for _, s := range filteredSpans {
+		if err := writer.WriteSpan(s); err != nil {
+			return "", fmt.Errorf("snapshotter: falha ao escrever span no snapshot: %w", err)
+		}
+	}
+	fmt.Printf("   Salvos %d spans relevantes em: %s\n", len(filteredSpans), spansSnapshotPath)
+
+	// Salva as entidades registradas do EntityStore
+	entitiesSlice := make([]executor.EntityRecord, 0, len(entityStore))
+	for _, entity := range entityStore {
+		entitiesSlice = append(entitiesSlice, entity)
+	}
+
+	entitiesSnapshotPath := filepath.Join(snapshotDir, "entities_snapshot.json")
+	entitiesData, err := json.MarshalIndent(entitiesSlice, "", "  ")
+	if err != nil {
+		return "", fmt.Errorf("snapshotter: falha ao serializar entidades: %w", err)
+	}
+	if err := ioutil.WriteFile(entitiesSnapshotPath, entitiesData, 0644); err != nil {
+		return "", fmt.Errorf("snapshotter: falha ao escrever entidades no snapshot: %w", err)
+	}
+	fmt.Printf("   Salvas %d entidades registradas em: %s\n", len(entitiesSlice), entitiesSnapshotPath)
+
+
+	// 3. Gerar metadados do snapshot
+	metadata := SnapshotMetadata{
+		SnapshotID:  snapshotID,
+		Timestamp:   time.Now().UTC().Format(time.RFC3339),
+		Goal:        goal,
+		Message:     message,
+		Inputs:      inputs,
+		SignedBy:    signedBy,
+		SpansCount:  len(filteredSpans),
+		EntitiesCount: len(entitiesSlice),
+		Context: map[string]interface{}{
+			"logline_version": "v0.4", // Atualiza a versão
+		},
+	}
+	metadataPath := filepath.Join(snapshotDir, "metadata.json")
+	metadataBytes, err := json.MarshalIndent(metadata, "", "  ")
+	if err != nil {
+		return "", fmt.Errorf("snapshotter: falha ao serializar metadados do snapshot: %w", err)
+	}
+	if err := ioutil.WriteFile(metadataPath, metadataBytes, 0644); err != nil {
+		return "", fmt.Errorf("snapshotter: falha ao escrever metadados do snapshot: %w", err)
+	}
+	fmt.Printf("   Metadados do snapshot salvos em: %s\n", metadataPath)
+
+
+	// 4. Integração com Git Real (`go-git`)
+	repoPath := filepath.Clean(filepath.Join(baseOutputDir, ".."))
+
+	r, err := git.PlainOpen(repoPath)
+	if err != nil {
+		fmt.Printf("   [INFO] Repositório Git não encontrado ou erro ao abrir (%v). Pulando integração Git.\n", err)
+	} else {
+		fmt.Println("   [INFO] Integrando com Git REAL...")
+		w, err := r.Worktree()
+		if err != nil {
+			return "", fmt.Errorf("snapshotter: falha ao obter Worktree Git: %w", err)
+		}
+
+		relSnapshotPath, err := filepath.Rel(repoPath, snapshotDir)
+		if err != nil {
+			return "", fmt.Errorf("snapshotter: falha ao obter caminho relativo do snapshot: %w", err)
+		}
+		_, err = w.Add(relSnapshotPath)
+		if err != nil {
+			return "", fmt.Errorf("snapshotter: falha ao adicionar snapshot ao Git staging: %w", err)
+		}
+		fmt.Printf("   Git: Adicionado '%s' ao staging.\n", relSnapshotPath)
+
+		commitMsg := fmt.Sprintf("LogLine Commit: %s (Snapshot ID: %s)", goal, snapshotID)
+		if message != "" {
+			commitMsg = fmt.Sprintf("%s\n\n%s", commitMsg, message)
+		}
+
+		authorName := "LogLineOS Automated Commit"
+		authorEmail := "loglineos@example.com"
+		if len(signedBy) > 0 {
+			authorName = strings.Join(signedBy, ", ") + " (via LogLineOS)"
+		}
+
+		commitHash, err := w.Commit(commitMsg, &git.CommitOptions{
+			Author: &object.Signature{
+				Name:  authorName,
+				Email: authorEmail,
+				When:  time.Now(),
+			},
+		})
+		if err != nil {
+			return "", fmt.Errorf("snapshotter: falha ao criar commit Git: %w", err)
+		}
+		fmt.Printf("   Git: Commit criado com sucesso: %s\n", commitHash.String())
+
+		metadata.GitHash = commitHash.String()
+		head, err := r.Head()
+		if err == nil {
+			metadata.GitBranch = head.Name().Short()
+		}
+		metadataBytesUpdated, err := json.MarshalIndent(metadata, "", "  ")
+		if err == nil {
+			ioutil.WriteFile(metadataPath, metadataBytesUpdated, 0644)
+		}
+	}
+
+	// 5. Assinatura Digital GPG Real (via exec.Command)
+	if signGPG {
+		fmt.Println("   [INFO] Tentando assinatura GPG REAL...")
+		cmd := exec.Command("gpg", "--batch", "--yes", "--armor", "--detach-sign", "--output", metadataPath+".asc", metadataPath)
+		cmd.Stderr = os.Stderr
+		cmd.Stdout = os.Stdout
+		
+		err := cmd.Run()
+		if err != nil {
+			log.Printf("⚠️ GPG_SIGN_FAIL: %v. Verifique se o GPG está instalado, sua chave está configurada e a senha pode ser fornecida (ou use --batch).\n", err)
+			metadata.Signed = false
+		} else {
+			metadata.Signed = true
+			metadata.Signature = filepath.Base(metadataPath) + ".asc"
+			fmt.Println("   GPG: Metadados do snapshot GPG-assinados com sucesso.")
+			metadataBytesUpdated, err := json.MarshalIndent(metadata, "", "  ")
+			if err == nil {
+				ioutil.WriteFile(metadataPath, metadataBytesUpdated, 0644)
+			}
+		}
+	}
+
+	return snapshotID, nil
+}
\ No newline at end of file
diff --git a/terreno_completo_FINAL/src/go/spans.go b/terreno_completo_FINAL/src/go/spans.go
new file mode 100644
index 0000000000000000000000000000000000000000..23bf08bf134b1b1d4a87f9dcf0fcaa277290bc0f
--- /dev/null
+++ b/terreno_completo_FINAL/src/go/spans.go
@@ -0,0 +1,258 @@
+// internal/spans/spans.go
+package spans
+
+import (
+	"encoding/json"
+	"fmt"
+	"io"
+	"os"
+	"sort"
+	"sync"
+	"time"
+
+	"github.com/Knetic/govaluate"
+)
+
+// Span representa um evento auditável no LogLineOS
+type Span struct {
+	SpanID       string                 `json:"span_id"`        // ID único do span
+	TaskID       string                 `json:"task_id"`        // ID do bloco no workflow
+	Type         string                 `json:"type"`           // Tipo do bloco (e.g., "register", "mechanical")
+	Status       string                 `json:"status"`         // Status da execução (success, error, warning, draft, skipped)
+	Timestamp    string                 `json:"timestamp"`      // Timestamp da execução (ISO 8601)
+	DurationMs   int64                  `json:"duration_ms"`    // Duração em milissegundos
+	SourceWorkflow string               `json:"source_workflow,omitempty"` // Nome do workflow de origem
+	InvokedBy    string                 `json:"invoked_by,omitempty"` // Quem invocou (usuário, sistema)
+	Attempt      int                    `json:"attempt,omitempty"` // Número da tentativa (para retries)
+	IsFallback   bool                   `json:"is_fallback,omitempty"` // Indica se este span é de um fallback
+	Inputs       map[string]interface{} `json:"inputs,omitempty"`       // Entradas ou dependências do bloco
+	Output       map[string]interface{} `json:"output,omitempty"`       // Saída do bloco
+	Context      map[string]interface{} `json:"context,omitempty"`      // Contexto adicional (versão, ambiente)
+	Message      string                 `json:"message,omitempty"`      // Mensagem resumida para o span
+}
+
+// JSONLWriter para escrever spans em formato JSON Lines
+type JSONLWriter struct {
+	file *os.File
+	mu   sync.Mutex // Protege a escrita concorrente
+}
+
+// NewJSONLWriter cria um novo escritor de spans JSONL
+func NewJSONLWriter(filePath string) (*JSONLWriter, error) {
+	// Usar O_APPEND|O_CREATE|O_WRONLY para adicionar ao final ou criar se não existe
+	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
+	if err != nil {
+		return nil, fmt.Errorf("spans: falha ao abrir arquivo '%s' para escrita: %w", filePath, err)
+	}
+	return &JSONLWriter{file: file}, nil
+}
+
+// WriteSpan serializa um Span para JSON e o escreve no arquivo, seguido por uma nova linha.
+func (w *JSONLWriter) WriteSpan(span Span) error {
+	w.mu.Lock()
+	defer w.mu.Unlock()
+
+	data, err := json.Marshal(span)
+	if err != nil {
+		return fmt.Errorf("spans: falha ao serializar span '%s': %w", span.SpanID, err)
+	}
+
+	if _, err := w.file.Write(data); err != nil {
+		return fmt.Errorf("spans: falha ao escrever dados do span '%s': %w", span.SpanID, err)
+	}
+	if _, err := w.file.WriteString("\n"); err != nil { // JSONL requires newline
+		return fmt.Errorf("spans: falha ao escrever newline para span '%s': %w", span.SpanID, err)
+	}
+	return nil
+}
+
+// Close fecha o arquivo de spans.
+func (w *JSONLWriter) Close() error {
+	if w.file == nil {
+		return nil // Já fechado ou nunca aberto
+	}
+	err := w.file.Close()
+	if err != nil {
+		return fmt.Errorf("spans: falha ao fechar arquivo de spans: %w", err)
+	}
+	w.file = nil // Evita fechamento duplo
+	return nil
+}
+
+// ReadAllSpans lê todos os spans de um arquivo JSONL
+func ReadAllSpans(filePath string) ([]Span, error) {
+	file, err := os.Open(filePath)
+	if err != nil {
+		if os.IsNotExist(err) {
+			return []Span{}, nil // Retorna slice vazio se o arquivo não existe, sem erro
+		}
+		return nil, fmt.Errorf("spans: falha ao abrir arquivo '%s' para leitura: %w", filePath, err)
+	}
+	defer file.Close()
+
+	var spans []Span
+	decoder := json.NewDecoder(file)
+
+	for {
+		var span Span
+		if err := decoder.Decode(&span); err == io.EOF {
+			break // Fim do arquivo
+		} else if err != nil {
+			return nil, fmt.Errorf("spans: falha ao decodificar linha em '%s': %w", filePath, err)
+		}
+		spans = append(spans, span)
+	}
+	return spans, nil
+}
+
+// FilterSpans filtra uma lista de spans com base em uma expressão.
+// Ex: "status == 'error' && duration_ms > 100"
+func FilterSpans(spans []Span, filterExpr string) ([]Span, error) {
+	if filterExpr == "" {
+		return spans, nil // Sem filtro, retorna todos os spans
+	}
+
+	expression, err := govaluate.NewEvaluableExpression(filterExpr)
+	if err != nil {
+		return nil, fmt.Errorf("spans: expressão de filtro '%s' inválida: %w", filterExpr, err)
+	}
+
+	filtered := []Span{}
+	for _, span := range spans {
+		// Converte o span para um map[string]interface{} para avaliação
+		spanMap, err := convertSpanToMap(span)
+		if err != nil {
+			return nil, fmt.Errorf("spans: erro ao converter span '%s' para map para avaliação: %w", span.SpanID, err)
+		}
+
+		result, err := expression.Evaluate(spanMap)
+		if err != nil {
+			// Erros de avaliação podem ocorrer se um campo não existir na expressão.
+			// Consideramos que o span não corresponde ao filtro neste caso.
+			continue
+		}
+
+		if boolResult, ok := result.(bool); ok && boolResult {
+			filtered = append(filtered, span)
+		}
+	}
+	return filtered, nil
+}
+
+// convertSpanToMap converte uma struct Span para um map[string]interface{} para govaluate
+func convertSpanToMap(span Span) (map[string]interface{}, error) {
+	// A maneira mais fácil de fazer isso é serializar e desserializar
+	// No futuro, considere usar reflexão para evitar o overhead de JSON
+	data, err := json.Marshal(span)
+	if err != nil {
+		return nil, fmt.Errorf("spans: falha ao serializar span para map de avaliação: %w", err)
+	}
+	var m map[string]interface{}
+	err = json.Unmarshal(data, &m)
+	if err != nil {
+		return nil, fmt.Errorf("spans: falha ao desserializar span para map de avaliação: %w", err)
+	}
+	return m, nil
+}
+
+
+// SortSpans ordena uma lista de spans por um campo e ordem (asc/desc)
+func SortSpans(spans []Span, sortBy string, order string) ([]Span, error) {
+	if sortBy == "" {
+		return spans, nil
+	}
+
+	sortedSpans := make([]Span, len(spans))
+	copy(sortedSpans, spans)
+
+	isAsc := (order == "asc" || order == "")
+
+	sort.Slice(sortedSpans, func(i, j int) bool {
+		s1 := sortedSpans[i]
+		s2 := sortedSpans[j]
+
+		var less bool
+		switch sortBy {
+		case "timestamp":
+			t1, err1 := time.Parse(time.RFC3339, s1.Timestamp)
+			t2, err2 := time.Parse(time.RFC3339, s2.Timestamp)
+			if err1 == nil && err2 == nil {
+				less = t1.Before(t2)
+			} else {
+				// Fallback para comparação de string se parse falhar
+				less = s1.Timestamp < s2.Timestamp
+			}
+		case "duration_ms":
+			less = s1.DurationMs < s2.DurationMs
+		case "type":
+			less = s1.Type < s2.Type
+		case "status":
+			less = s1.Status < s2.Status
+		case "task_id":
+			less = s1.TaskID < s2.TaskID
+		case "attempt":
+			less = s1.Attempt < s2.Attempt
+		case "is_fallback":
+			// false (não é fallback) vem antes de true (é fallback) em ordem ascendente
+			less = !s1.IsFallback && s2.IsFallback
+		case "span_id":
+			less = s1.SpanID < s2.SpanID
+		default:
+			// No caso de campos aninhados ou complexos, seria necessário mais lógica (reflexão, etc.)
+			// Para simplicidade, vamos usar fmt.Sprintf para tentar comparar, o que não é ideal para todos os tipos.
+			// Um sistema premium teria um `GetField(span, fieldPath string) (interface{}, error)` robusto.
+			less = fmt.Sprintf("%v", s1) < fmt.Sprintf("%v", s2)
+		}
+
+		if !isAsc {
+			return !less // Inverte para ordem descendente
+		}
+		return less
+	})
+
+	return sortedSpans, nil
+}
+
+// AuditMetrics guarda métricas agregadas dos spans
+type AuditMetrics struct {
+	TotalSpans    int                       `json:"total_spans"`
+	SpansByStatus map[string]int            `json:"spans_by_status"`
+	SpansByType   map[string]int            `json:"spans_by_type"`
+	AvgDurationMsByType map[string]float64    `json:"avg_duration_ms_by_type"`
+	TotalDurationMs int64                     `json:"total_duration_ms"`
+}
+
+// CollectMetrics coleta métricas básicas de uma lista de spans
+func CollectMetrics(spans []Span) AuditMetrics {
+	metrics := AuditMetrics{
+		TotalSpans:    len(spans),
+		SpansByStatus: make(map[string]int),
+		SpansByType:   make(map[string]int),
+		AvgDurationMsByType: make(map[string]float64),
+		TotalDurationMs: 0,
+	}
+
+	typeDurations := make(map[string][]int64)
+
+	for _, span := range spans {
+		metrics.SpansByStatus[span.Status]++
+		metrics.SpansByType[span.Type]++
+		metrics.TotalDurationMs += span.DurationMs
+
+		if span.Status == "success" {
+			typeDurations[span.Type] = append(typeDurations[span.Type], span.DurationMs)
+		}
+	}
+
+	for spanType, durations := range typeDurations {
+		if len(durations) > 0 {
+			var sum int64
+			for _, d := range durations {
+				sum += d
+			}
+			metrics.AvgDurationMsByType[spanType] = float64(sum) / float64(len(durations))
+		}
+	}
+
+	return metrics
+}
\ No newline at end of file
diff --git a/terreno_completo_FINAL/src/go/webserver.go b/terreno_completo_FINAL/src/go/webserver.go
new file mode 100644
index 0000000000000000000000000000000000000000..cfdaa10908a56f1658bcf2b8fe36789a257ea51b
--- /dev/null
+++ b/terreno_completo_FINAL/src/go/webserver.go
@@ -0,0 +1,120 @@
+// internal/utils/webserver.go
+package utils
+
+import (
+	"context"
+	"encoding/json"
+	"fmt"
+	"io/ioutil"
+	"net/http"
+	"sync"
+	"time"
+)
+
+// WebhookPayload representa a estrutura básica de um payload de webhook recebido
+type WebhookPayload struct {
+	Timestamp string                 `json:"timestamp"`
+	Method    string                 `json:"method"`
+	Path      string                 `json:"path"`
+	Headers   map[string][]string    `json:"headers"`
+	Body      json.RawMessage        `json:"body"`
+	Query     map[string][]string    `json:"query"`
+}
+
+// WebhookEvent representa um evento de webhook com a URL que o disparou
+type WebhookEvent struct {
+	URL     string
+	Payload WebhookPayload
+}
+
+// WebhookServer gerencia um servidor HTTP para receber webhooks
+type WebhookServer struct {
+	server      *http.Server
+	eventCh     chan WebhookEvent
+	listenAddr  string
+	routes      map[string]struct{} // Caminhos que estamos monitorando
+	mu          sync.Mutex          // Protege o mapa de rotas
+}
+
+// NewWebhookServer cria uma nova instância de WebhookServer
+func NewWebhookServer(listenAddr string) *WebhookServer {
+	return &WebhookServer{
+		eventCh:    make(chan WebhookEvent, 100), // Buffer de eventos
+		listenAddr: listenAddr,
+		routes:     make(map[string]struct{}),
+	}
+}
+
+// AddRoute adiciona uma rota (caminho) para o servidor monitorar
+func (ws *WebhookServer) AddRoute(path string) {
+	ws.mu.Lock()
+	defer ws.mu.Unlock()
+	ws.routes[path] = struct{}{}
+	fmt.Printf("Webserver: Monitorando rota de webhook: %s%s\n", ws.listenAddr, path)
+}
+
+// Start inicia o servidor HTTP em uma goroutine
+func (ws *WebhookServer) Start() {
+	mux := http.NewServeMux()
+	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
+		ws.mu.Lock()
+		_, isMonitoredRoute := ws.routes[r.URL.Path]
+		ws.mu.Unlock()
+
+		if !isMonitoredRoute {
+			http.NotFound(w, r)
+			return
+		}
+
+		fmt.Printf("Webserver: Recebido webhook em %s %s\n", r.Method, r.URL.Path)
+
+		body, err := ioutil.ReadAll(r.Body)
+		if err != nil {
+			http.Error(w, "Erro ao ler corpo da requisição", http.StatusInternalServerError)
+			return
+		}
+
+		payload := WebhookPayload{
+			Timestamp: time.Now().UTC().Format(time.RFC3339),
+			Method:    r.Method,
+			Path:      r.URL.Path,
+			Headers:   r.Header,
+			Body:      body,
+			Query:     r.URL.Query(),
+		}
+
+		// Envia o evento para o canal
+		select {
+		case ws.eventCh <- WebhookEvent{URL: r.URL.String(), Payload: payload}:
+			w.WriteHeader(http.StatusOK)
+			fmt.Fprint(w, "Webhook recebido com sucesso!")
+		case <-time.After(1 * time.Second): // Timeout para não bloquear o handler
+			w.WriteHeader(http.StatusServiceUnavailable)
+			fmt.Fprint(w, "Servidor ocupado, tente novamente.")
+		}
+	})
+
+	ws.server = &http.Server{
+		Addr:    ws.listenAddr,
+		Handler: mux,
+	}
+
+	go func() {
+		fmt.Printf("Webserver: Iniciando servidor de webhook em %s\n", ws.listenAddr)
+		if err := ws.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
+			fmt.Fprintf(os.Stderr, "Webserver: Erro ao iniciar servidor HTTP: %v\n", err)
+		}
+	}()
+}
+
+// Stop encerra o servidor HTTP
+func (ws *WebhookServer) Stop(ctx context.Context) error {
+	fmt.Println("Webserver: Encerrando servidor de webhook...")
+	close(ws.eventCh) // Fecha o canal para sinalizar o fim dos eventos
+	return ws.server.Shutdown(ctx)
+}
+
+// Events retorna o canal de eventos de webhook
+func (ws *WebhookServer) Events() <-chan WebhookEvent {
+	return ws.eventCh
+}
\ No newline at end of file
diff --git a/terreno_completo_FINAL/src/rust/._auth.rs b/terreno_completo_FINAL/src/rust/._auth.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/._auth_service.rs b/terreno_completo_FINAL/src/rust/._auth_service.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/._cache.rs b/terreno_completo_FINAL/src/rust/._cache.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/._contract.rs b/terreno_completo_FINAL/src/rust/._contract.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/._contract_service.rs b/terreno_completo_FINAL/src/rust/._contract_service.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/._db.rs b/terreno_completo_FINAL/src/rust/._db.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/._db_service.rs b/terreno_completo_FINAL/src/rust/._db_service.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/._errors.rs b/terreno_completo_FINAL/src/rust/._errors.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/._handlers.rs b/terreno_completo_FINAL/src/rust/._handlers.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/._health.rs b/terreno_completo_FINAL/src/rust/._health.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/._infra.rs b/terreno_completo_FINAL/src/rust/._infra.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/._infra_service.rs b/terreno_completo_FINAL/src/rust/._infra_service.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/._jwt.rs b/terreno_completo_FINAL/src/rust/._jwt.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/._lib.rs b/terreno_completo_FINAL/src/rust/._lib.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/._llm.rs b/terreno_completo_FINAL/src/rust/._llm.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/._llm_service.rs b/terreno_completo_FINAL/src/rust/._llm_service.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/._logline_runner.rs b/terreno_completo_FINAL/src/rust/._logline_runner.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/._main.rs b/terreno_completo_FINAL/src/rust/._main.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/._span.rs b/terreno_completo_FINAL/src/rust/._span.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/._telemetry.rs b/terreno_completo_FINAL/src/rust/._telemetry.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/._tenant.rs b/terreno_completo_FINAL/src/rust/._tenant.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/._time_utils.rs b/terreno_completo_FINAL/src/rust/._time_utils.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/._user.rs b/terreno_completo_FINAL/src/rust/._user.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/._utils.rs b/terreno_completo_FINAL/src/rust/._utils.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/._validation.rs b/terreno_completo_FINAL/src/rust/._validation.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/._whatsapp.rs b/terreno_completo_FINAL/src/rust/._whatsapp.rs
new file mode 100644
index 0000000000000000000000000000000000000000..20ae59c33f77fb064d7b1d7e3cbe1431087a2fe2
GIT binary patch
literal 163
zcmZQz6=P>$Vqox1Ojhs@R)|o50+1L3ClDI}aUl?c_=|y<2;dkJ5(HHS(lG;wxzV&S
oBE&_L^K<nQ3kq^l^$Lpe%Tn_a^O92;7?~J89ORzWqUEs}063!)-2eap

literal 0
HcmV?d00001

diff --git a/terreno_completo_FINAL/src/rust/auth.rs b/terreno_completo_FINAL/src/rust/auth.rs
new file mode 100644
index 0000000000000000000000000000000000000000..5a0255aaa8b7f46cd1725c18fd53a20e70530888
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/auth.rs
@@ -0,0 +1,8 @@
+// Handlers de autenticação (login, refresh, revoke)
+// Exemplo mínimo
+use actix_web::{post, web, HttpResponse, Responder};
+
+#[post("/api/auth/login")]
+pub async fn login_api() -> impl Responder {
+    HttpResponse::Ok().body("Login handler")
+}
diff --git a/terreno_completo_FINAL/src/rust/auth_service.rs b/terreno_completo_FINAL/src/rust/auth_service.rs
new file mode 100644
index 0000000000000000000000000000000000000000..76f560f1fb1d06edeb0af8b11c2a4bb8c58392a0
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/auth_service.rs
@@ -0,0 +1 @@
+// Serviço de autenticação: hash/sal de senhas, geração e validação de JWT
diff --git a/terreno_completo_FINAL/src/rust/cache.rs b/terreno_completo_FINAL/src/rust/cache.rs
new file mode 100644
index 0000000000000000000000000000000000000000..17aad169d7229f3b4a6e62b5eb75c4a1892aba90
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/cache.rs
@@ -0,0 +1 @@
+// Redis por tenant, TTL, contadores de cotas LLM
diff --git a/terreno_completo_FINAL/src/rust/contract.rs b/terreno_completo_FINAL/src/rust/contract.rs
new file mode 100644
index 0000000000000000000000000000000000000000..9c6ec4aca3181891ec990e7dd01e363e9fef620e
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/contract.rs
@@ -0,0 +1,7 @@
+// Handlers de contratos (CRUD, execução, versionamento)
+use actix_web::{get, post, web, HttpResponse, Responder};
+
+#[get("/api/contracts/list")]
+pub async fn list() -> impl Responder {
+    HttpResponse::Ok().body("List contracts handler")
+}
diff --git a/terreno_completo_FINAL/src/rust/contract_service.rs b/terreno_completo_FINAL/src/rust/contract_service.rs
new file mode 100644
index 0000000000000000000000000000000000000000..454df1cea2cc0af5fac8f5a923666aec43e40589
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/contract_service.rs
@@ -0,0 +1 @@
+// Serviço de contratos: validação de YAML, persistência, versionamento
diff --git a/terreno_completo_FINAL/src/rust/db.rs b/terreno_completo_FINAL/src/rust/db.rs
new file mode 100644
index 0000000000000000000000000000000000000000..3b593b4f72853a70a3c290f501c299be3b980c42
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/db.rs
@@ -0,0 +1 @@
+// Configuração do pool Postgres, RLS, migrações automáticas
diff --git a/terreno_completo_FINAL/src/rust/db_service.rs b/terreno_completo_FINAL/src/rust/db_service.rs
new file mode 100644
index 0000000000000000000000000000000000000000..69bd44327c07996a4efdcae2e79425f6c16e7934
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/db_service.rs
@@ -0,0 +1 @@
+// Serviço de banco de dados: pool SQLx, transações, migrações
diff --git a/terreno_completo_FINAL/src/rust/errors.rs b/terreno_completo_FINAL/src/rust/errors.rs
new file mode 100644
index 0000000000000000000000000000000000000000..ca7fd859feee0b75e39f13aa9d0eafe40cb2e7bf
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/errors.rs
@@ -0,0 +1 @@
+// Enum AppError, Result<T, AppError>, conversão para HttpResponse
diff --git a/terreno_completo_FINAL/src/rust/handlers.rs b/terreno_completo_FINAL/src/rust/handlers.rs
new file mode 100644
index 0000000000000000000000000000000000000000..9f02e98c5165796fca92895e7226b6a3aba08f6a
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/handlers.rs
@@ -0,0 +1,87 @@
+use actix_web::{get, post, web, HttpResponse, Responder};
+use sqlx::SqlitePool;
+use crate::AppState;
+use serde::Deserialize;
+use std::time::SystemTime;
+use crate::logline_runner;
+use crate::utils::record_span;
+
+// Página de login: redireciona para o frontend WASM que irá buscar splash.logline
+#[get("/login")]
+pub async fn login_page() -> impl Responder {
+    // Simplesmente devolve o index.html do WASM; o WASM vai carregar /contracts/ui/login.logline
+    HttpResponse::Ok()
+        .content_type("text/html")
+        .body(include_str!("../frontend/static/index.html"))
+}
+
+// API de login: recebe JSON (ou form) e responde com fragmento HTML
+#[derive(Deserialize)]
+pub struct LoginData {
+    email: String,
+    senha: String,
+}
+
+#[post("/api/auth/login")]
+pub async fn login_api(
+    data: web::Json<LoginData>,
+    state: web::Data<AppState>,
+) -> impl Responder {
+    // Gera span de auditoria
+    record_span("auth.login", &format!("user:{}", data.email)).await;
+
+    // Simulação de autenticação (substituir por consulta real a DB)
+    if data.email.contains("@") && data.senha.len() >= 6 {
+        let html = r#"
+        <div class="p-6 bg-green-100 rounded">
+          <h2 class="text-xl font-semibold mb-4">Login bem-sucedido!</h2>
+          <a href="/dashboard" class="text-blue-600 hover:underline">Ir para o dashboard</a>
+        </div>
+        "#;
+        HttpResponse::Ok().content_type("text/html").body(html)
+    } else {
+        let html = r#"
+        <div class="p-6 bg-red-100 rounded">
+          <h2 class="text-xl font-semibold mb-4 text-red-700">Credenciais inválidas</h2>
+          <a href="/login" class="text-blue-600 hover:underline">Tentar novamente</a>
+        </div>
+        "#;
+        HttpResponse::Ok().content_type("text/html").body(html)
+    }
+}
+
+// Listar contratos ativos (JSON ou HTML fragment para HTMX)
+#[get("/api/contracts/list")]
+pub async fn contracts_list(state: web::Data<AppState>) -> impl Responder {
+    let pool = state.db.lock().await;
+    let rows = sqlx::query!("SELECT id, title, status FROM contracts")
+        .fetch_all(&*pool)
+        .await
+        .unwrap_or_else(|_| vec![]);
+
+    // Gera fragmento HTML da tabela
+    let mut html = String::new();
+    html.push_str(r#"<table class="min-w-full bg-gray-50 shadow overflow-hidden rounded-lg">"#);
+    html.push_str(r#"<thead><tr><th class="px-4 py-2 border-b">ID</th><th class="px-4 py-2 border-b">Título</th><th class="px-4 py-2 border-b">Status</th></tr></thead><tbody>"#);
+    for row in rows {
+        html.push_str(&format!(
+            r#"<tr><td class="px-4 py-2 border-b">{}</td><td class="px-4 py-2 border-b">{}</td><td class="px-4 py-2 border-b">{}</td></tr>"#,
+            row.id, row.title, row.status
+        ));
+    }
+    html.push_str("</tbody></table>");
+    HttpResponse::Ok().content_type("text/html").body(html)
+}
+
+// Handler LLM: recebe prompt via JSON e devolve JSON com “conteúdo”
+#[derive(Deserialize)]
+pub struct LlmRequest {
+    prompt: String,
+}
+
+#[post("/api/llm/summary")]
+pub async fn llm_summary(req: web::Json<LlmRequest>) -> impl Responder {
+    // Executa Span llm.gateway via função auxiliar
+    let summary = logline_runner::run_llm_gateway(&req.prompt).await.unwrap_or_else(|e| e);
+    HttpResponse::Ok().json(serde_json::json!({ "summary": summary }))
+}
diff --git a/terreno_completo_FINAL/src/rust/health.rs b/terreno_completo_FINAL/src/rust/health.rs
new file mode 100644
index 0000000000000000000000000000000000000000..c4dcaaa051b08b024cb4f12f7bb431f2efa16d8b
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/health.rs
@@ -0,0 +1,7 @@
+// Handler de health/liveness/readiness
+use actix_web::{get, HttpResponse, Responder};
+
+#[get("/health")]
+pub async fn health() -> impl Responder {
+    HttpResponse::Ok().body("OK")
+}
diff --git a/terreno_completo_FINAL/src/rust/infra.rs b/terreno_completo_FINAL/src/rust/infra.rs
new file mode 100644
index 0000000000000000000000000000000000000000..93d0d119543b4521855033a2caee4777300944ea
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/infra.rs
@@ -0,0 +1,7 @@
+// Handlers de infraestrutura (executar spans de infra)
+use actix_web::{post, web, HttpResponse, Responder};
+
+#[post("/api/infra/deploy")]
+pub async fn deploy() -> impl Responder {
+    HttpResponse::Ok().body("Infra deploy handler")
+}
diff --git a/terreno_completo_FINAL/src/rust/infra_service.rs b/terreno_completo_FINAL/src/rust/infra_service.rs
new file mode 100644
index 0000000000000000000000000000000000000000..b612acee504885d308e562a0f55b77892e2bc350
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/infra_service.rs
@@ -0,0 +1 @@
+// Serviço de infraestrutura: client k8s, AWS SDK, alertas Prometheus
diff --git a/terreno_completo_FINAL/src/rust/jwt.rs b/terreno_completo_FINAL/src/rust/jwt.rs
new file mode 100644
index 0000000000000000000000000000000000000000..a8d818ed2f61616cd4da1787a8dd5062d3c0c0a3
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/jwt.rs
@@ -0,0 +1 @@
+// Chave pública/privada, algoritmos HS256/RS256
diff --git a/terreno_completo_FINAL/src/rust/lib.rs b/terreno_completo_FINAL/src/rust/lib.rs
new file mode 100644
index 0000000000000000000000000000000000000000..c9d73990da81a73dc365f283e1e70c43cc8ff886
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/lib.rs
@@ -0,0 +1,123 @@
+use wasm_bindgen::prelude::*;
+use web_sys::{Document, Element};
+use serde::Deserialize;
+use serde_yaml::Value as YamlValue;
+use std::collections::BTreeMap;
+
+// Define a estrutura básica de um span ui.render
+#[derive(Debug, Deserialize)]
+struct Span {
+    id: Option<String>,
+    #[serde(rename = "type")]
+    span_type: String,
+    blocks: Option<Vec<Block>>,
+}
+
+#[derive(Debug, Deserialize)]
+struct Block {
+    component: String,
+    #[serde(default)]
+    props: BTreeMap<String, YamlValue>,
+    #[serde(default)]
+    children: Option<Vec<YamlValue>>, // simplificado, para strings ou blocos aninhados
+}
+
+#[wasm_bindgen(start)]
+pub fn run() {
+    // Permite erros de panic aparecerem no console do navegador
+    console_error_panic_hook::set_once();
+}
+
+// Função exposta ao JS para renderizar um contrato .logline no <div id="app">
+#[wasm_bindgen]
+pub fn render_contract(logline_text: &str) -> Result<(), JsValue> {
+    let contract: Vec<Span> = serde_yaml::from_str(logline_text)
+        .map_err(|e| JsValue::from_str(&format!("YAML parse error: {}", e)))?;
+    let window = web_sys::window().unwrap();
+    let document: Document = window.document().unwrap();
+    let app_div: Element = document.get_element_by_id("app")
+        .ok_or_else(|| JsValue::from_str("Elemento #app não encontrado"))?;
+
+    // Limpa conteúdo
+    app_div.set_inner_html("");
+
+    for span in contract {
+        if span.span_type == "ui.render" {
+            if let Some(blocks) = span.blocks {
+                for block in blocks {
+                    let child_el = create_element_from_block(&document, &block)?;
+                    app_div.append_child(&child_el)?;
+                }
+            }
+        }
+    }
+    Ok(())
+}
+
+// Converte um Block em um Element DOM
+fn create_element_from_block(document: &Document, block: &Block) -> Result<Element, JsValue> {
+    let tag = &block.component;
+    let el = document.create_element(tag)
+        .map_err(|e| JsValue::from_str(&format!("Falha ao criar <{}>: {:?}", tag, e)))?;
+
+    // Atribui props (class, text, style, onclick, etc.)
+    for (key, val) in &block.props {
+        match key.as_str() {
+            "class" => {
+                if let Some(s) = val.as_str() {
+                    el.set_attribute("class", s)?;
+                }
+            }
+            "text" => {
+                if let Some(s) = val.as_str() {
+                    el.set_text_content(Some(s));
+                }
+            }
+            "style" => {
+                // Suporta um mapa YAML de estilos
+                if let Some(style_map) = val.as_mapping() {
+                    let mut style_str = String::new();
+                    for (k, v) in style_map {
+                        if let (Some(ks), Some(vs)) = (k.as_str(), v.as_str()) {
+                            style_str.push_str(&format!("{}:{};", ks, vs));
+                        }
+                    }
+                    el.set_attribute("style", &style_str)?;
+                }
+            }
+            "action" => {
+                // Trata ui.click: atribui evento onclick que gera span ui.click
+                if let Some(action_name) = val.as_str() {
+                    let action = action_name.to_string();
+                    let closure = Closure::wrap(Box::new(move || {
+                        // Gera um span de clique; aqui só logamos no console
+                        web_sys::console::log_1(&JsValue::from_str(
+                            &format!(r#"{{\"type\":\"ui.click\",\"action\":\"{}\",\"timestamp\":\"{}\"}}"#,
+                                action,
+                                js_sys::Date::new_0().to_iso_string()
+                            )
+                        ));
+                    }) as Box<dyn FnMut()>);
+                    el.add_event_listener_with_callback("click", closure.as_ref().unchecked_ref())?;
+                    closure.forget();
+                }
+            }
+            "href" => {
+                if let Some(s) = val.as_str() {
+                    el.set_attribute("href", s)?;
+                }
+            }
+            _ => {
+                // Outras props genéricas: data-*
+                if let Some(s) = val.as_str() {
+                    el.set_attribute(key, s)?;
+                }
+            }
+        }
+    }
+
+    // Caso haja children aninhados (HTML bruto ou novos blocks em YAML),
+    // poderíamos parsear recursivamente. Para simplicidade, ignoramos aqui.
+
+    Ok(el)
+}
diff --git a/terreno_completo_FINAL/src/rust/llm.rs b/terreno_completo_FINAL/src/rust/llm.rs
new file mode 100644
index 0000000000000000000000000000000000000000..32b61fc16e9413c926bfd5396cfe378891ba677b
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/llm.rs
@@ -0,0 +1,7 @@
+// Handlers de LLM (summarize, plan)
+use actix_web::{post, web, HttpResponse, Responder};
+
+#[post("/api/llm/summary")]
+pub async fn summary() -> impl Responder {
+    HttpResponse::Ok().body("LLM summary handler")
+}
diff --git a/terreno_completo_FINAL/src/rust/llm_service.rs b/terreno_completo_FINAL/src/rust/llm_service.rs
new file mode 100644
index 0000000000000000000000000000000000000000..7d7c3a83141a53d5aff1777826a6255007803249
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/llm_service.rs
@@ -0,0 +1 @@
+// Serviço de LLM: roteamento local/cloud, controle de quotas, monitoramento de custos
diff --git a/terreno_completo_FINAL/src/rust/logline_runner.rs b/terreno_completo_FINAL/src/rust/logline_runner.rs
new file mode 100644
index 0000000000000000000000000000000000000000..da92953f8abd75ac37e0c188ee0a015b38df0cf9
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/logline_runner.rs
@@ -0,0 +1,52 @@
+use std::time::SystemTime;
+use tokio::process::Command;
+use serde_json::Value;
+
+// Função que simula “llm.gateway”: tenta local Mistral, senão cloud
+pub async fn run_llm_gateway(prompt: &str) -> Result<String, String> {
+    // Primeiro, tenta LLM local (Mistral via CLI)
+    let output = Command::new("ollama")
+        .arg("run")
+        .arg("mistral")
+        .arg(prompt)
+        .output()
+        .await
+        .map_err(|e| format!("Erro ao executar LLM local: {}", e))?;
+
+    if output.status.success() {
+        let resp = String::from_utf8_lossy(&output.stdout).into_owned();
+        // Grava span llm.response
+        record_span("llm.route", "mistral").await;
+        record_span("llm.response", &resp).await;
+        return Ok(resp);
+    }
+    // Se falhou, chama API OpenAI (curl)
+    record_span("llm.fallback", "true").await;
+    let api_key = std::env::var("OPENAI_API_KEY").unwrap_or_default();
+    let body = serde_json::json!({
+      "model": "gpt-4",
+      "messages": [{ "role": "user", "content": prompt }],
+      "temperature": 0.3
+    });
+    let resp = reqwest::Client::new()
+        .post("https://api.openai.com/v1/chat/completions")
+        .bearer_auth(api_key)
+        .json(&body)
+        .send()
+        .await
+        .map_err(|e| format!("Erro OpenAI: {}", e))?;
+
+    let data: Value = resp.json().await.map_err(|e| format!("Parse JSON falhou: {}", e))?;
+    let content = data["choices"][0]["message"]["content"].as_str().unwrap_or("").to_string();
+    record_span("llm.route", "gpt-4").await;
+    record_span("llm.response", &content).await;
+    Ok(content)
+}
+
+// Função auxiliar para gravar um span genérico em JSONL (por simplicidade, grava no console)
+pub async fn record_span(span_type: &str, message: &str) {
+    let now = SystemTime::now();
+    let timestamp = chrono::DateTime::<chrono::Utc>::from(now).to_rfc3339();
+    let span = format!(r#"{{"timestamp":"{}","type":"{}","message":"{}"}}"#, timestamp, span_type, message);
+    println!("{}", span); // Em produção, grave em arquivo .jsonl ou serviço de logs
+}
diff --git a/terreno_completo_FINAL/src/rust/main.rs b/terreno_completo_FINAL/src/rust/main.rs
new file mode 100644
index 0000000000000000000000000000000000000000..e0184d6fe05ead9b6c4a481ee30f27de95209c90
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/main.rs
@@ -0,0 +1,54 @@
+use actix_files::Files;
+use actix_web::{
+    get, post,
+    web::{self, Data},
+    App, HttpResponse, HttpServer, Responder,
+};
+use std::sync::Arc;
+use tokio::sync::Mutex;
+use serde::Deserialize;
+use sqlx::{Pool, Sqlite};
+use std::time::SystemTime;
+
+mod handlers;
+mod logline_runner;
+mod utils;
+
+#[actix_web::main]
+async fn main() -> std::io::Result<()> {
+    // Inicializa logger
+    env_logger::init();
+
+    // Cria pool SQLite (para protótipo; em produção use Postgres)
+    let pool = Pool::<Sqlite>::connect("sqlite://minicontratos.db").await.unwrap();
+
+    // Se quiser rodar migrações (usando sqlx-cli)
+    // sqlx::migrate!("./migrations").run(&pool).await.unwrap();
+
+    // Estado compartilhado
+    let app_state = Data::new(AppState {
+        db: Arc::new(Mutex::new(pool)),
+    });
+
+    HttpServer::new(move || {
+        App::new()
+            .app_data(app_state.clone())
+            // Serve arquivos estáticos do frontend (WASM, CSS, HTML)
+            .service(Files::new("/static", "./frontend/static").show_files_listing())
+            // Endpoints da UI (via WASM)
+            .service(Files::new("/contracts/ui", "./contracts/ui"))
+            // Endpoints:
+            .service(handlers::login_page)
+            .service(handlers::login_api)
+            .service(handlers::contracts_list)
+            .service(handlers::llm_summary)
+    })
+    .bind(("127.0.0.1", 8080))?
+    .run()
+    .await
+}
+
+pub struct AppState {
+    // Pool de DB protegido por Mutex para concorrência simples
+    db: Arc<Mutex<Pool<Sqlite>>>,
+}
diff --git a/terreno_completo_FINAL/src/rust/span.rs b/terreno_completo_FINAL/src/rust/span.rs
new file mode 100644
index 0000000000000000000000000000000000000000..b0114732383463e3b824e77bdc71916d373dfa9a
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/span.rs
@@ -0,0 +1 @@
+// Model de span
diff --git a/terreno_completo_FINAL/src/rust/telemetry.rs b/terreno_completo_FINAL/src/rust/telemetry.rs
new file mode 100644
index 0000000000000000000000000000000000000000..e4fdb67907ec3540c05a850fc822cdc8a5a9d1c5
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/telemetry.rs
@@ -0,0 +1 @@
+// OpenTelemetry: tracer, exporter OTLP, métricas customizadas
diff --git a/terreno_completo_FINAL/src/rust/tenant.rs b/terreno_completo_FINAL/src/rust/tenant.rs
new file mode 100644
index 0000000000000000000000000000000000000000..4ab7fbb7c54e474678240d68b2322538039a51b4
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/tenant.rs
@@ -0,0 +1 @@
+// Model de tenant
diff --git a/terreno_completo_FINAL/src/rust/time_utils.rs b/terreno_completo_FINAL/src/rust/time_utils.rs
new file mode 100644
index 0000000000000000000000000000000000000000..b628c3dda1dd47e809dcc06afc3003fddb131fbe
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/time_utils.rs
@@ -0,0 +1 @@
+// Funções de timestamp, formatação
diff --git a/terreno_completo_FINAL/src/rust/user.rs b/terreno_completo_FINAL/src/rust/user.rs
new file mode 100644
index 0000000000000000000000000000000000000000..b5298a8b6d8f3f1b6795faa506898686991b40b5
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/user.rs
@@ -0,0 +1 @@
+// Model de usuário
diff --git a/terreno_completo_FINAL/src/rust/utils.rs b/terreno_completo_FINAL/src/rust/utils.rs
new file mode 100644
index 0000000000000000000000000000000000000000..347397dd881deb37ee4b02161bf38fb76851cccc
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/utils.rs
@@ -0,0 +1 @@
+// Qualquer utilitário comum, ex.: geração de resposta HTML a partir de spans ui.render do backend
diff --git a/terreno_completo_FINAL/src/rust/validation.rs b/terreno_completo_FINAL/src/rust/validation.rs
new file mode 100644
index 0000000000000000000000000000000000000000..3cd03156e17128ea3691f451108034d998352cdc
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/validation.rs
@@ -0,0 +1 @@
+// Validações customizadas, esquema YAML/JSON
diff --git a/terreno_completo_FINAL/src/rust/whatsapp.rs b/terreno_completo_FINAL/src/rust/whatsapp.rs
new file mode 100644
index 0000000000000000000000000000000000000000..458ad7dcd63313d3aa878f890a686067246b301f
--- /dev/null
+++ b/terreno_completo_FINAL/src/rust/whatsapp.rs
@@ -0,0 +1,7 @@
+// Handlers de WhatsApp (listar conversas, enviar mensagem)
+use actix_web::{post, web, HttpResponse, Responder};
+
+#[post("/api/whatsapp/send")]
+pub async fn send() -> impl Responder {
+    HttpResponse::Ok().body("WhatsApp send handler")
+}
diff --git a/terreno_completo_FINAL/ui/index.html b/terreno_completo_FINAL/ui/index.html
new file mode 100644
index 0000000000000000000000000000000000000000..772fbfd94c826004f2806b01dd6adbc4cbdd7beb
--- /dev/null
+++ b/terreno_completo_FINAL/ui/index.html
@@ -0,0 +1 @@
+<html><body>Frontend simbólico</body></html>
\ No newline at end of file
 
EOF
)