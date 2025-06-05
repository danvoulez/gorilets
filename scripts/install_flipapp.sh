#!/usr/bin/env bash
set -euo pipefail

# =============================================================================
# Instalador “de primeira classe” do FlipApp / LogLineOS para VM Ubuntu/Debian
# =============================================================================
#
# Este script faz tudo de forma automatizada:
#   1. Verifica pré-requisitos (root, sistema suportado).
#   2. Instala pacotes essenciais (git, build-essential, curl, Go, Rust, wasm-pack, Nginx, UFW).
#   3. Cria usuário “flipapp” e estrutura de diretórios (/opt/flipapp, /etc/loglineos).
#   4. Baixa ou extrai o código-fonte do FlipApp/LogLineOS (via Git ou arquivo ZIP local).
#   5. Compila o back-end em Go e copia o binário para /opt/flipapp/bin/loglineos.
#   6. Compila o front-end em Rust → gera WASM/JS e deposita em /opt/flipapp/static.
#   7. Ajusta permissões, configura firewall (UFW) e Nginx (proxy reverso opcional).
#   8. Cria e habilita serviço systemd para iniciar o FlipApp automaticamente.
#
# Uso:
#   1) Copie este script para /opt/flipapp/install_flipapp.sh
#   2) Edite as variáveis GIT_REPO (ou defina SOURCE_ZIP), se necessário.
#   3) Torne executável: sudo chmod +x /opt/flipapp/install_flipapp.sh
#   4) Execute: sudo /opt/flipapp/install_flipapp.sh
#
# Após a execução, o FlipApp estará disponível em http://<IP-da-VM>:8080
#
# Observação:
#   • Se quiser usar um ZIP local (em vez de Git), defina SOURCE_ZIP=/caminho/para/arquivo.zip
#   • Por padrão, o script define GIT_REPO="" e espera que SOURCE_ZIP seja informado ou que o
#     repositório já esteja em /opt/flipapp/source. Ajuste conforme necessidade.
# =============================================================================

# ---------------------------- Variáveis Globais -----------------------------
FLIPAPP_USER="flipapp"
FLIPAPP_HOME="/opt/flipapp"
CONFIG_DIR="/etc/loglineos"
BIN_DIR="${FLIPAPP_HOME}/bin"
SRC_DIR="${FLIPAPP_HOME}/source"
STATIC_DIR="${FLIPAPP_HOME}/static"
CONFIG_DEST="${FLIPAPP_HOME}/config"
DATA_DIR="${FLIPAPP_HOME}/data"
EXAMPLES_DIR="${FLIPAPP_HOME}/examples"
SYSTEMD_UNIT="/etc/systemd/system/loglineos.service"
REQUIRED_PORT="8080"

# Escolha entre GIT_REPO (clonar) ou SOURCE_ZIP (extrair localmente).
# Se preferir baixar de GitHub, preencha GIT_REPO. Caso contrário, deixe vazio e coloque um ZIP em /opt/flipapp/source.zip
GIT_REPO=""  
# Exemplo: GIT_REPO="https://github.com/seu-usuario/flipapp.git"
SOURCE_ZIP="/opt/flipapp/source.zip"  # Se usar arquivo ZIP, coloque-o aqui

# Cores para saída
RED="$(printf '\033[0;31m')"
GREEN="$(printf '\033[0;32m')"
YELLOW="$(printf '\033[1;33m')"
BLUE="$(printf '\033[0;34m')"
NO_COLOR="$(printf '\033[0m')"

# ----------------------------- Funções Auxiliares ---------------------------

function info {
    echo -e "${BLUE}[INFO]${NO_COLOR} $1"
}

function success {
    echo -e "${GREEN}[OK]${NO_COLOR} $1"
}

function warning {
    echo -e "${YELLOW}[AVISO]${NO_COLOR} $1"
}

function error_exit {
    echo -e "${RED}[ERRO]${NO_COLOR} $1"
    exit 1
}

# Verifica se o comando existe
function check_command {
    if ! command -v "$1" &> /dev/null; then
        return 1
    fi
    return 0
}

# ------------------------ Verificações Iniciais -----------------------------

# 1) Verificar se está rodando como root
if [[ $EUID -ne 0 ]]; then
    error_exit "Este script precisa ser executado como root (ou via sudo)."
fi

# 2) Verificar sistema operacional (Ubuntu/Debian)
if ! grep -Ei "ubuntu|debian" /etc/os-release &> /dev/null; then
    warning "Este script foi testado em Ubuntu/Debian. Se estiver usando outra distro, pode falhar."
fi

info "=============================="
info "  Instalador do FlipApp/LogLineOS"
info "=============================="

# -------------------------- Instalação de Pacotes ---------------------------

function install_base_packages {
    info "Instalando pacotes essenciais: git, build-essential, curl, ufw..."
    apt-get update -y
    apt-get install -y git build-essential curl ufw
    success "Pacotes básicos instalados."
}

function install_go {
    if check_command go; then
        warning "Go já está instalado ($(go version)). Pulando."
    else
        info "Instalando Go..."
        apt-get install -y golang
        if ! check_command go; then
            error_exit "Falha ao instalar Go."
        fi
        success "Go instalado: $(go version)."
    fi
}

function install_rust_and_wasm {
    if check_command rustc && check_command cargo; then
        warning "Rust já está instalado ($(rustc --version)). Pulando."
    else
        info "Instalando Rust (rustup)..."
        curl https://sh.rustup.rs -sSf | bash -s -- -y
        export PATH="/root/.cargo/bin:$PATH"
        if ! check_command rustc || ! check_command cargo; then
            error_exit "Falha ao instalar Rust."
        fi
        success "Rust instalado: $(rustc --version)"
    fi

    if check_command wasm-pack; then
        warning "wasm-pack já está instalado ($(wasm-pack --version)). Pulando."
    else
        info "Instalando wasm-pack..."
        cargo install wasm-pack
        if ! check_command wasm-pack; then
            error_exit "Falha ao instalar wasm-pack."
        fi
        success "wasm-pack instalado: $(wasm-pack --version)"
    fi
}

function install_nginx {
    if check_command nginx; then
        warning "Nginx já está instalado ($(nginx -v 2>&1)). Pulando."
    else
        info "Instalando Nginx..."
        apt-get install -y nginx
        systemctl enable nginx
        systemctl start nginx
        success "Nginx instalado e iniciado."
    fi
}

function configure_ufw {
    info "Configurando firewall (UFW)..."
    # Permitir SSH
    ufw allow OpenSSH
    # Permitir porta do FlipApp
    ufw allow "${REQUIRED_PORT}/tcp"
    # Habilitar UFW de forma não interativa
    ufw --force enable
    success "UFW configurado (SSH e porta ${REQUIRED_PORT} liberados)."
}

# ------------------- Criação de Usuário e Diretórios -----------------------

function create_flipapp_user_and_dirs {
    # 1) Criar usuário de sistema se não existir
    if id -u "${FLIPAPP_USER}" &> /dev/null; then
        warning "Usuário ${FLIPAPP_USER} já existe. Pulando criação."
    else
        info "Criando usuário de sistema '${FLIPAPP_USER}'..."
        adduser --system --group --home "${FLIPAPP_HOME}" "${FLIPAPP_USER}"
        success "Usuário ${FLIPAPP_USER} criado."
    fi

    # 2) Criar diretórios principais
    info "Criando estrutura de diretórios em ${FLIPAPP_HOME} e ${CONFIG_DIR}..."
    mkdir -p "${FLIPAPP_HOME}"/{bin,source,config,static,data,examples}
    chown -R "${FLIPAPP_USER}:${FLIPAPP_USER}" "${FLIPAPP_HOME}"
    chmod -R 750 "${FLIPAPP_HOME}"

    mkdir -p "${CONFIG_DIR}"
    chown root:root "${CONFIG_DIR}"
    chmod 755 "${CONFIG_DIR}"
    success "Diretórios criados e permissões ajustadas."
}

# -------------------------- Baixar / Extrair Código ------------------------

function fetch_or_extract_source {
    # Se GIT_REPO estiver definido, clonar; caso contrário, montar a partir de SOURCE_ZIP
    if [[ -n "${GIT_REPO}" ]]; then
        info "Clonando repositório Git em ${SRC_DIR}..."
        mkdir -p "${SRC_DIR}"
        chown "${FLIPAPP_USER}:${FLIPAPP_USER}" "${SRC_DIR}"
        sudo -u "${FLIPAPP_USER}" git clone "${GIT_REPO}" "${SRC_DIR}"
        success "Repositório clonado."
    else
        if [[ -f "${SOURCE_ZIP}" ]]; then
            info "Extraindo arquivo ZIP ${SOURCE_ZIP} para ${SRC_DIR}..."
            mkdir -p "${SRC_DIR}"
            chown "${FLIPAPP_USER}:${FLIPAPP_USER}" "${SRC_DIR}"
            sudo -u "${FLIPAPP_USER}" unzip -q "${SOURCE_ZIP}" -d "${SRC_DIR}"
            success "Arquivo ZIP extraído."
        else
            warning "Nem GIT_REPO nem SOURCE_ZIP definidos / encontrados."
            warning "Coloque o código-fonte em ${SRC_DIR} manualmente e pressione [Enter] para continuar..."
            read -r
            if [[ ! -d "${SRC_DIR}" ]]; then
                error_exit "Pasta ${SRC_DIR} não existe. Saindo."
            fi
        fi
    fi
}

# -------------------------- Compilação do Back-end -------------------------

function build_backend_go {
    info "Compilando o back-end Go (loglineos)..."
    # Detectar o caminho do ponto de entrada (ex.: internal/ui/promptpad_app.go)
    local entry_point
    if [[ -f "${SRC_DIR}/internal/ui/promptpad_app.go" ]]; then
        entry_point="${SRC_DIR}/internal/ui/promptpad_app.go"
    else
        # Tentar localizar qualquer main.go se a estrutura for diferente
        entry_point="$(find "${SRC_DIR}" -maxdepth 3 -type f -name 'promptpad_app.go' | head -n 1)"
        if [[ -z "${entry_point}" ]]; then
            error_exit "Não encontrei promptpad_app.go em ${SRC_DIR}."
        fi
    fi

    # Compilar usando Go
    cd "$(dirname "${entry_point}")"
    sudo -u "${FLIPAPP_USER}" go mod tidy
    sudo -u "${FLIPAPP_USER}" go build -o "${FLIPAPP_HOME}/bin/loglineos" "${entry_point}"

    if [[ ! -x "${FLIPAPP_HOME}/bin/loglineos" ]]; then
        error_exit "Falha ao compilar loglineos."
    fi
    chown "${FLIPAPP_USER}:${FLIPAPP_USER}" "${FLIPAPP_HOME}/bin/loglineos"
    chmod 750 "${FLIPAPP_HOME}/bin/loglineos"
    success "Back-end Go compilado: ${FLIPAPP_HOME}/bin/loglineos."
}

# ------------------------ Configuração de Arquivos --------------------------

function copy_configuration_files {
    info "Copiando arquivos de configuração para ${CONFIG_DEST}..."
    if [[ -d "${SRC_DIR}/config" ]]; then
        cp -r "${SRC_DIR}/config/." "${CONFIG_DEST}/"
        chown -R "${FLIPAPP_USER}:${FLIPAPP_USER}" "${CONFIG_DEST}"
        chmod -R 640 "${CONFIG_DEST}"/*
        success "Arquivos de configuração copiados."
    else
        warning "Pasta de configuração ${SRC_DIR}/config não encontrada. Pule esta etapa."
    fi
}

# ------------------------- Compilação Front-end WASM ------------------------

function build_frontend_wasm {
    if [[ -d "${SRC_DIR}/frontend" ]]; then
        info "Compilando front-end em Rust → WASM..."
        cd "${SRC_DIR}/frontend"
        # Garantir que o ambiente Rust esteja carregado
        export PATH="/root/.cargo/bin:$PATH"
        sudo -u "${FLIPAPP_USER}" wasm-pack build --release --target web --out-dir "${STATIC_DIR}"
        chown -R "${FLIPAPP_USER}:${FLIPAPP_USER}" "${STATIC_DIR}"
        chmod -R 750 "${STATIC_DIR}"
        success "Front-end compilado (arquivos em ${STATIC_DIR})."
    else
        warning "Pasta frontend/ não encontrada em ${SRC_DIR}. Pule a compilação WASM."
    fi
}

# ----------------------- Configuração do Serviço systemd --------------------

function create_systemd_service {
    info "Criando unit file systemd em ${SYSTEMD_UNIT}..."
    cat > "${SYSTEMD_UNIT}" <<EOF
[Unit]
Description=LogLineOS – Motor de Execução de Spans (FlipApp)
After=network.target

[Service]
Type=simple
User=${FLIPAPP_USER}
WorkingDirectory=${FLIPAPP_HOME}
ExecStart=${FLIPAPP_HOME}/bin/loglineos --config ${CONFIG_DEST}/logline.yml
Restart=on-failure
LimitNOFILE=65536

[Install]
WantedBy=multi-user.target
EOF

    chmod 644 "${SYSTEMD_UNIT}"
    systemctl daemon-reload
    systemctl enable loglineos
    success "Serviço systemd criado e habilitado (loglineos)."
}

# ----------------------------- Proxy Reverso Nginx --------------------------

function configure_nginx_proxy {
    local site_conf="/etc/nginx/sites-available/flipapp"
    info "Configurando Nginx para proxy reverso em porta ${REQUIRED_PORT}..."

    cat > "${site_conf}" <<EOF
server {
    listen 80;
    server_name _;  # Pode trocar para domínio real

    location / {
        proxy_pass         http://127.0.0.1:${REQUIRED_PORT};
        proxy_http_version 1.1;
        proxy_set_header   Upgrade \$http_upgrade;
        proxy_set_header   Connection keep-alive;
        proxy_set_header   Host \$host;
        proxy_cache_bypass \$http_upgrade;
    }

    # Servir arquivos estáticos (WASM/JS)
    location /static/ {
        alias ${STATIC_DIR}/;
    }
}
EOF

    ln -sf "${site_conf}" /etc/nginx/sites-enabled/flipapp
    nginx -t
    systemctl reload nginx
    success "Nginx configurado (proxy para FlipApp em http://localhost:${REQUIRED_PORT})."
}

# --------------------------------- Main --------------------------------------

# 1) Instalar pacotes base
install_base_packages

# 2) Instalar Go
install_go

# 3) Instalar Rust + wasm-pack
install_rust_and_wasm

# 4) Instalar Nginx (opcional)
install_nginx

# 5) Configurar firewall (UFW)
configure_ufw

# 6) Criar usuário e diretórios
create_flipapp_user_and_dirs

# 7) Obter código-fonte (Git ou ZIP ou manual)
fetch_or_extract_source

# 8) Compilar Back-end Go
build_backend_go

# 9) Copiar configurações
copy_configuration_files

# 10) Compilar Front-end WASM (Rust)
build_frontend_wasm

# 11) Criar serviço systemd
create_systemd_service

# 12) Configurar Nginx para proxy reverso
configure_nginx_proxy

echo
success "======================================================"
success "   FlipApp/LogLineOS instalado com sucesso!          "
success "   Acesse: http://<IP-da-VM> ou http://<seu-dominio>   "
success "======================================================"
echo
