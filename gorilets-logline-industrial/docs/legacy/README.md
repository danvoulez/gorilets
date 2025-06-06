# FlipApp = Minicontratos

O ambiente operacional oficial da VoulezVous, onde tudo é contrato, tudo é rastreável, e tudo acontece com semântica nativa LogLine.

## Estrutura

```
flipapp/
├── cmd/                   ← CLI e comandos locais
├── internal/
│   ├── contracts/         ← parser, executor, runner (LogLine core)
│   ├── spans/             ← storage, timeline, JSONL e DB
│   ├── auth/              ← JWT + tenants + users
│   └── channels/          ← WhatsApp, OpenAI, Email...
├── frontend/              ← FlipApp UI
│   ├── PromptPad/         ← terminal IA + executor
│   ├── ContractsView/     ← timeline e spans
│   └── ChannelsPanel/     ← histórico por canal
├── api/                   ← HTTP layer (Go), rotas REST
├── config/                ← logline.yml, tenant settings
├── migrations/
├── scripts/
└── README.md              ← manifesto institucional FlipApp = Minicontratos
```

## Telas

- **LogLine**: REPL institucional, chat, executor de contratos, visualizador de spans.
- **WhatsApp**: canal humano, cada conversa = contrato auditável, agente automático.
- **New**: gerador rápido de contratos/registro, semeador de cultura institucional.

## Princípios

- Tudo é contrato LogLine.
- Não há dashboards artificiais: toda consulta é feita via prompt.
- Auditabilidade e rastreabilidade nativas.

## Exemplos de Contrato LogLine

```
# Exemplo: registro de presença
- did: registrar_presenca
  who: cliente
  when: "2025-06-02T09:00:00Z"
  affair: onboarding
```

```
# Exemplo: mensagem WhatsApp
- did: enviar_audio
  payload:
    duracao: 12
    origem: whatsapp
```

```
# Exemplo: criar rotina de boas-vindas
- contract: onboarding
  who: cliente
  did: enviar_mensagem
  payload:
    texto: "Bem-vindo ao sistema!"
```

## Rotas de API sugeridas

- POST /api/logline
- GET /api/logline/history?userId=&limit=
- GET /api/whatsapp/conversations?tenantId=
- GET /api/whatsapp/conversations/{conversationId}/messages?limit=&offset=
- POST /api/whatsapp/send
- GET /api/templates?tenantId=
- POST /api/contracts
- POST /api/contracts/{id}/execute

## Como contribuir

Abra issues, envie PRs ou sugira novos contratos LogLine para enriquecer o ecossistema institucional.
