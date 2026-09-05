# Guia Completo: Cloudflare Tunnel para Homelab no Ubuntu Server

O **Cloudflare Tunnel (`cloudflared`)** é a forma mais moderna, segura e gratuita de expor aplicações do seu servidor doméstico (notebook) para a internet sem abrir portas no roteador de casa e sem depender de IP público estático.

---

## 1. Por que usar o Cloudflare Tunnel?

- **Zero Port Forwarding**: Não é necessário configurar redirecionamento de portas (NAT) no roteador da sua operadora (Claro, Vivo, etc.).
- **Funciona atrás de CGNAT**: Mesmo se o seu provedor usar CGNAT (IP compartilhado), o túnel funciona perfeitamente pois estabelece uma conexão de saída criptografada com a borda da Cloudflare.
- **HTTPS / SSL Automático e Grátis**: Certificados SSL gerenciados e renovados automaticamente pela Cloudflare.
- **Proteção contra DDoS**: O IP real da sua casa nunca é revelado para os visitantes.

---

## 2. Passo a Passo de Configuração

### Passo 1: Adicionar seu Domínio na Cloudflare
1. Crie uma conta gratuita em [cloudflare.com](https://www.cloudflare.com) (caso ainda não tenha).
2. Adicione seu domínio próprio (ex: `diegocamargo.dev` ou similar).
3. Aponte os nameservers do seu registrador (ex: Registro.br, Hostinger, GoDaddy) para os nameservers indicados pela Cloudflare.

### Passo 2: Criar o Túnel no Cloudflare Zero Trust
1. No painel da Cloudflare, clique no menu lateral em **Zero Trust** (ou acesse [one.dash.cloudflare.com](https://one.dash.cloudflare.com)).
2. Vá em **Networks** -> **Tunnels**.
3. Clique em **Add a tunnel** e selecione o tipo **Cloudflare (cloudflared)**.
4. Dê um nome para o túnel (ex: `homelab-notebook`).
5. Na tela seguinte de instalação, escolha o ambiente **Docker**.
6. A Cloudflare exibirá um comando contendo um token alfanumérico longo. Copie apenas o valor do token após `--token`. Exemplo:
   ```text
   eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
   ```

### Passo 3: Configurar os Subdomínios (Public Hostnames)
Na aba **Public Hostnames** do seu túnel na Cloudflare, adicione 3 rotas:

| Subdomínio | Domínio | Tipo | URL do Serviço (Interna Docker) |
| :--- | :--- | :--- | :--- |
| `portfolio` | `seu-dominio.com` | `HTTP` | `portfolio_go:8080` |
| `finhub` | `seu-dominio.com` | `HTTP` | `inv_frontend:80` |
| `manga` | `seu-dominio.com` | `HTTP` | `manga_linker_frontend:5173` |

> 💡 **Nota Importante**: Como o container do túnel (`homelab_cloudflared`) estará na mesma rede Docker (`homelab_network`) que os outros serviços, ele se comunica diretamente pelo nome do container (`portfolio_go`, `inv_frontend`, `manga_linker_frontend`), sem precisar de portas expostas no host!

---

## 3. Subir o Container do Túnel no Ubuntu Server

1. No seu notebook Ubuntu Server, entre na pasta do portfólio:
   ```bash
   cd /home/seu-usuario/homelab/Portifolio/deploy
   ```
2. Crie um arquivo `.env` para o túnel ou exporte a variável com o token copiado no Passo 2:
   ```bash
   echo "CLOUDFLARE_TUNNEL_TOKEN=seu_token_aqui" > .env
   ```
3. Suba o container:
   ```bash
   docker compose -f docker-compose.tunnel.yml up -d
   ```
4. Verifique os logs para confirmar a conexão:
   ```bash
   docker logs homelab_cloudflared
   ```
   Você verá a mensagem confirmando conexão estabelecida com múltiplos data centers da Cloudflare.

Acesse agora no navegador do seu celular ou computador:
- `https://portfolio.seu-dominio.com`
- `https://finhub.seu-dominio.com`
- `https://manga.seu-dominio.com`

Seus projetos já estão no ar com HTTPS seguro e de alta performance!
