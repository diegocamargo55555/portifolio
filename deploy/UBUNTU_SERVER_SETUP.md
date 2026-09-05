# Guia de Configuração do Notebook Ubuntu Server (Homelab)

Este guia contém todas as etapas necessárias para transformar seu notebook com **Ubuntu Server** em um servidor Homelab 24/7 confiável, sem que ele durma ao fechar a tampa, e orquestrando o **Portfólio Go**, o **FinHub** e o **Manga Wrapper** via Docker.

---

## 1. Evitar Suspensão ao Fechar a Tampa do Notebook

Por padrão, distribuições Linux entram em modo de suspensão (*sleep*) ao detectar o fechamento da tampa. No Ubuntu Server, desative essa suspensão seguindo os passos:

1. Abra o arquivo de configuração do gerenciador de login:
   ```bash
   sudo nano /etc/systemd/logind.conf
   ```
2. Localize as seguintes linhas (remova o `#` se estiverem comentadas) e altere para `ignore`:
   ```ini
   HandleLidSwitch=ignore
   HandleLidSwitchExternalPower=ignore
   HandleLidSwitchDocked=ignore
   ```
3. Salve o arquivo (`Ctrl+O`, `Enter` e `Ctrl+X`).
4. Reinicie o serviço `systemd-logind` para aplicar as alterações imediatamente:
   ```bash
   sudo systemctl restart systemd-logind
   ```
> **Nota**: Agora você pode fechar a tampa do notebook e ele continuará funcionando ininterruptamente como servidor.

---

## 2. Garantir Auto-Início do Docker no Boot

Para garantir que o Docker e todos os containers subam automaticamente caso o notebook seja reiniciado ou após uma queda de energia:

```bash
sudo systemctl enable docker
sudo systemctl enable containerd
```

Como todos os nossos `docker-compose.yml` possuem a diretiva `restart: unless-stopped`, os containers subirão sozinhos no boot.

---

## 3. Criar a Rede Docker Compartilhada (`homelab_network`)

Todos os projetos do homelab comunicam-se através de uma rede Docker interna dedicada. Crie-a uma única vez no notebook:

```bash
docker network create homelab_network
```

---

## 4. Configuração do Firewall (UFW)

Para garantir segurança na sua rede local:

```bash
# Permite acesso SSH (troque para sua porta se não for 22)
sudo ufw allow 22/tcp

# Portas opcionais caso deseje acessar diretamente pelo IP local
sudo ufw allow 80/tcp
sudo ufw allow 8080/tcp
sudo ufw allow 3001/tcp
sudo ufw allow 3002/tcp

# Ativa o firewall
sudo ufw enable
sudo ufw status verbose
```

---

## 5. Estrutura de Pastas Recomendada no Servidor

Recomendamos organizar seus três projetos no servidor da seguinte forma:

```text
/home/seu-usuario/
├── homelab/
│   ├── Portifolio/            # Este repositório (Portfólio Go)
│   ├── inv/                   # Repositório FinHub
│   ├── manga-wrapper/         # Repositório Manga Wrapper
│   └── tunnel/                # Cloudflare Tunnel (deploy/docker-compose.tunnel.yml)
```

---

## 6. Inicialização dos Serviços

### Passo A: Subir o Portfólio Go
```bash
cd /home/seu-usuario/homelab/Portifolio
docker compose up -d --build
```

### Passo B: Subir o FinHub
```bash
cd /home/seu-usuario/homelab/inv
# (Aplicar o patch de portas explicado em deploy/integration-patches/inv-patch.md)
docker compose up -d --build
```

### Passo C: Subir o Manga Wrapper
```bash
cd /home/seu-usuario/homelab/manga-wrapper
# (Aplicar o patch de portas explicado em deploy/integration-patches/manga-patch.md)
docker compose up -d --build
```

### Passo D: Subir o Cloudflare Tunnel
```bash
cd /home/seu-usuario/homelab/Portifolio/deploy
docker compose -f docker-compose.tunnel.yml up -d
```

Pronto! Todos os containers estarão rodando de forma isolada, leve e sem conflito de portas no seu notebook.
