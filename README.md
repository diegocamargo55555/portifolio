# 🚀 Portfólio Go & Homelab Multi-Docker no Ubuntu Server

Repositório oficial do portfólio de engenharia de software de **Diego Camargo**, desenvolvido nativamente em **Golang** e projetado para orquestração em Homelab no **Ubuntu Server** (notebook) ao lado dos projetos **FinHub** e **Manga Wrapper** sem conflitos de portas.

---

## 🌟 Destaques do Projeto

- **100% Nativo em Golang**: Servidor de alta performance construído com o framework Gin e arquitetura limpa.
- **Tamanho Mínimo & Eficiência Extrema**:
  - Imagem Docker de apenas **~14.4MB**.
  - Consumo em execução de apenas **~8MB de RAM**.
  - Inicialização em menos de **1ms**.
- **Binário Estático Único com `embed.FS`**: Templates HTML, estilos e scripts embutidos diretamente no binário compilado.
- **Bilíngue Instantâneo (PT-BR / EN)**: Chaveador de idioma com persistência em cookie e suporte a recrutadores globais.
- **Design Dark Minimalista Premium**: Visual estilo Linear/Vercel/GitHub dark com toques em ciano e esmeralda, modais de arquitetura técnica e visualizador do Homelab em tempo real.
- **Orquestração Homelab & Cloudflare Tunnel**:
  - Isolamento completo de bancos de dados na rede Docker `homelab_network`.
  - Resolução dos choques de porta 5432 e 8080 do FinHub e Manga Wrapper.
  - Acesso público seguro sem abrir portas no roteador de casa.

---

## 🛠️ Stack Tecnológica

- **Linguagem**: Golang 1.22+
- **Web Framework**: Gin Gonic (`github.com/gin-gonic/gin`)
- **Assets & Templates**: Go `embed.FS` + `html/template`
- **Estilização**: Tailwind CSS + Custom CSS (Glassmorphism & Glow Effects)
- **Containerização**: Docker Multi-stage (Golang Alpine -> Alpine Runtime)
- **Orquestração**: Docker Compose & Docker Network Bridge
- **Ingress / Rede Segura**: Cloudflare Zero Trust Tunnel (`cloudflared`)

---

## 🚀 Como Executar Localmente

### Pré-requisitos
- Golang 1.22 ou superior instalado

### Execução Direta com Go
```bash
# Baixar dependências
go mod download

# Executar em modo desenvolvimento
go run main.go
```
O servidor estará acessível em: `http://localhost:8080`

Para testar o idioma inglês diretamente na URL:
`http://localhost:8080/?lang=en`

---

## 🐳 Como Executar com Docker

```bash
# Construir a imagem e subir o container
docker compose up -d --build

# Verificar logs
docker compose logs -f
```

---

## 📂 Estrutura do Repositório

```text
├── Dockerfile                  # Multi-stage Go build (~14.4MB final)
├── docker-compose.yml          # Compose do portfólio conectado à homelab_network
├── .env.example                # Configurações de ambiente
├── main.go                     # Ponto de entrada do servidor Go
├── go.mod / go.sum             # Módulo e dependências Go
├── internal/
│   ├── config/                 # Carregamento de variáveis de ambiente
│   ├── data/                   # Dados dos projetos e traduções PT-BR / EN
│   ├── handlers/               # Handlers HTTP (Renderização, Healthcheck, Homelab Ping)
│   └── models/                 # Modelos de dados TypeScript/Go
├── web/
│   ├── static/
│   │   ├── css/style.css       # Estilização Glassmorphism e tema dark
│   │   └── js/app.js           # Modais e atualização de status em tempo real
│   └── templates/
│       └── index.html          # Template HTML principal
└── deploy/
    ├── UBUNTU_SERVER_SETUP.md    # Guia para notebook (tampa fechada, auto-boot, UFW)
    ├── CLOUDFLARE_TUNNEL_GUIDE.md# Passo a passo de subdomínios e DNS no Cloudflare
    ├── docker-compose.tunnel.yml # Compose do Cloudflare Tunnel
    └── integration-patches/      # Overrides prontos para FinHub e Manga Wrapper
        ├── inv-docker-compose.override.yml
        ├── manga-docker-compose.override.yml
        └── README.md
```

---

## 📚 Guias de Infraestrutura e Homelab

- [Guia do Notebook Ubuntu Server](deploy/UBUNTU_SERVER_SETUP.md) (como desativar sleep com tampa fechada, auto-início no boot).
- [Guia do Cloudflare Tunnel](deploy/CLOUDFLARE_TUNNEL_GUIDE.md) (passo a passo para expor subdomínios com SSL gratuito).
- [Instruções de Integração de Portas](deploy/integration-patches/README.md) (resolução de conflitos de FinHub e Manga Wrapper).

---

## 👤 Desenvolvedor

**Diego Camargo**
- GitHub: [@diegocamargo55555](https://github.com/diegocamargo55555)
- E-mail: [diegocamargo55555@gmail.com](mailto:diegocamargo55555@gmail.com)
# portifolio
