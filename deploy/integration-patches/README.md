# Como Integrar FinHub e Manga Wrapper sem Conflito de Portas

O Docker Compose possui um mecanismo nativo onde qualquer arquivo chamado `docker-compose.override.yml` presente na mesma pasta do `docker-compose.yml` tem suas configurações mescladas automaticamente.

Isso significa que **você não precisa alterar os arquivos originais versionados no git** de nenhum dos repositórios! Basta copiar os arquivos de override.

---

## 1. Integrando o FinHub (`inv`)

1. Vá até a pasta onde você clonou o repositório `inv`:
   ```bash
   cd /home/seu-usuario/homelab/inv
   ```
2. Copie o arquivo de override para dentro da pasta:
   ```bash
   cp /home/seu-usuario/homelab/Portifolio/deploy/integration-patches/inv-docker-compose.override.yml ./docker-compose.override.yml
   ```
3. Suba o FinHub:
   ```bash
   docker compose up -d
   ```

### O que mudou:
- Frontend FinHub: acessível localmente em `http://localhost:3001` (ou na rede local do notebook).
- Backend FinHub: acessível localmente em `http://localhost:8081`.
- Postgres FinHub: porta `5433` (livre).
- Redis FinHub: porta `6380` (livre).
- Conectado à rede `homelab_network` para o Cloudflare Tunnel.

---

## 2. Integrando o Manga Wrapper (`manga-wrapper`)

1. Vá até a pasta onde você clonou o repositório `manga-wrapper`:
   ```bash
   cd /home/seu-usuario/homelab/manga-wrapper
   ```
2. Copie o arquivo de override para dentro da pasta:
   ```bash
   cp /home/seu-usuario/homelab/Portifolio/deploy/integration-patches/manga-docker-compose.override.yml ./docker-compose.override.yml
   ```
3. Suba o Manga Wrapper:
   ```bash
   docker compose up -d
   ```

### O que mudou:
- Frontend Manga: acessível localmente em `http://localhost:3002`.
- Backend Manga: acessível localmente em `http://localhost:8082`.
- Postgres Manga: porta `5434` (livre, sem conflito com o FinHub).
- Conectado à rede `homelab_network` para o Cloudflare Tunnel.

---

## 3. Resumo das Portas Locais no Notebook

| Aplicação | Serviço | Porta Local no Host | Nome Interno no Docker (`homelab_network`) |
| :--- | :--- | :--- | :--- |
| **Portfólio Go** | Web + API | `:8080` (ou `:80`) | `portfolio_go:8080` |
| **FinHub** | Frontend | `:3001` | `inv_frontend:80` |
| **FinHub** | Backend Go | `:8081` | `inv_backend:8080` |
| **FinHub** | Postgres | `:5433` | `inv_postgres:5432` |
| **FinHub** | Redis | `:6380` | `inv_redis:6379` |
| **Manga Wrapper** | Frontend | `:3002` | `manga_linker_frontend:5173` |
| **Manga Wrapper** | Backend Go | `:8082` | `manga_linker_backend:8080` |
| **Manga Wrapper** | Postgres | `:5434` | `manga_linker_db:5432` |
| **Cloudflare Tunnel** | Daemon | *Nenhuma (Saída pura)* | `homelab_cloudflared` |

Nenhum container conflita com outro, e todos rodam com máxima estabilidade!
