# TERRENO LogLineOS

Este pacote contém a implementação simbólica completa do sistema operacional declarativo LogLineOS,
incluindo:

- Executor LLM com cache auditável
- Autenticação simbólica via flat-file
- Frontend simbólico em HTML
- Contratos `.logline` reais e simulados
- Scripts auxiliares e código-fonte em Go, Rust e Python
- Suporte a spans, replay e visualização simbólica

## Estrutura

- `contracts/`: contratos organizacionais, cognitivos e de sistema
- `bin/`: executores (`logline`, `watch`, `archive`, `cognitive`)
- `spans/`: rastreamento completo das execuções
- `ui/`: visualização simbólica e auditável
- `scripts/`: utilitários de instalação, cache, migração
- `src/`: código-fonte modular por linguagem

## Execução

```bash
logline run contracts/system/build_vm.logline
systemctl start logline-watch
```

Acesse o frontend em `http://localhost:8080`.
