# Architecture

```
[src/logline/*.logline]
      ↓ parser Lark → AST
      ↓
[python runtime]
      ↓ database
      ↓
[WebSocket]
      ↓
[frontend/index.html + JS]
```
