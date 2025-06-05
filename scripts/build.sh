#!/usr/bin/env bash
set -e

# 1. Compila o crate Rust para WASM usando wasm-pack
cd frontend
wasm-pack build --target web --out-dir static

# 2. Ajusta permissões (se necessário) e volta
cd ..

echo "Frontend WASM compilado com sucesso em frontend/static/"
