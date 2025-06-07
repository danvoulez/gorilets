# Usage

Run a UI script:
```bash
logline run src/logline/ui/ui_flip.logline
```

To initialize the entire frontend via LogLine:
```
logline run src/logline/ui/frontend_runtime.logline
```

Simulate infrastructure provisioning:
```bash
LOGLINE_MODE=simulate logline run src/logline/infra/install_vm.logline
```
