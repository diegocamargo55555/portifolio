package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"

	"github.com/diegocamargo55555/portfolio/internal/config"
	"github.com/diegocamargo55555/portfolio/internal/handlers"
	"github.com/gin-gonic/gin"
)

// Embed todos os templates e arquivos estáticos dentro do executável compilado
//
//go:embed web/templates/* web/static/*
var embeddedFS embed.FS

func main() {
	cfg := config.Load()

	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// Configuração de templates embutidos
	tmpl, err := template.ParseFS(embeddedFS, "web/templates/*.html")
	if err != nil {
		log.Fatalf("Falha crítica ao carregar templates embutidos: %v", err)
	}
	r.SetHTMLTemplate(tmpl)

	// Configuração de arquivos estáticos embutidos
	staticFS, err := fs.Sub(embeddedFS, "web/static")
	if err != nil {
		log.Fatalf("Falha crítica ao inicializar assets estáticos: %v", err)
	}
	r.StaticFS("/static", http.FS(staticFS))

	// Inicializa os handlers
	h := handlers.New(cfg)

	// Rotas do Portfólio
	r.GET("/", h.Index)
	r.HEAD("/", h.Index)
	r.GET("/set-lang/:lang", h.SetLanguage)
	r.GET("/api/health", h.HealthCheck)
	r.HEAD("/api/health", h.HealthCheck)
	r.GET("/api/projects/:slug", h.ProjectDetail)
	r.GET("/api/homelab-status", h.HomelabStatus)

	addr := fmt.Sprintf("0.0.0.0:%s", cfg.Port)
	log.Printf("🚀 Portfólio Go inicializado com sucesso em %s (Ambiente: %s)", addr, cfg.AppEnv)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Erro ao iniciar servidor HTTP: %v", err)
	}
}
