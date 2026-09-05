package handlers

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/diegocamargo55555/portfolio/internal/config"
	"github.com/diegocamargo55555/portfolio/internal/data"
	"github.com/gin-gonic/gin"
)

var startTime = time.Now()

type Handler struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Handler {
	return &Handler{cfg: cfg}
}

// resolveLanguage identifica o idioma preferido do usuário
func (h *Handler) resolveLanguage(c *gin.Context) string {
	// 1. Query parameter ?lang=en ou ?lang=pt
	if qLang := c.Query("lang"); qLang != "" {
		qLang = strings.ToLower(qLang)
		if qLang == "en" || qLang == "pt" {
			c.SetCookie("portfolio_lang", qLang, 365*24*3600, "/", "", false, false)
			return qLang
		}
	}

	// 2. Cookie salvo
	if cookieLang, err := c.Cookie("portfolio_lang"); err == nil {
		cookieLang = strings.ToLower(cookieLang)
		if cookieLang == "en" || cookieLang == "pt" {
			return cookieLang
		}
	}

	// 3. Header Accept-Language do navegador
	accept := strings.ToLower(c.GetHeader("Accept-Language"))
	if strings.Contains(accept, "en") && !strings.Contains(accept, "pt") {
		return "en"
	}

	return h.cfg.DefaultLang
}

// Index renderiza a página inicial do portfólio
func (h *Handler) Index(c *gin.Context) {
	lang := h.resolveLanguage(c)
	pageData := data.GetData(lang)

	// Customiza domínios baseado na variável de ambiente
	for i := range pageData.Projects {
		pageData.Projects[i].LiveURL = fmt.Sprintf("https://%s.%s", pageData.Projects[i].Subdomain, h.cfg.Domain)
	}

	pageData.HostInfo.OS = "Ubuntu Server (Linux x86_64)"
	pageData.HostInfo.Runtime = "Golang 1.22+ (Multi-stage Scratch/Alpine)"
	pageData.HostInfo.Engine = "Docker Compose & Bridge Network"
	pageData.HostInfo.Tunnel = "Cloudflare Zero Trust Tunnel"

	c.HTML(http.StatusOK, "index.html", pageData)
}

// SetLanguage atualiza o cookie de idioma e redireciona
func (h *Handler) SetLanguage(c *gin.Context) {
	lang := c.Param("lang")
	if lang != "en" && lang != "pt" {
		lang = "pt"
	}

	c.SetCookie("portfolio_lang", lang, 365*24*3600, "/", "", false, false)

	referer := c.Request.Referer()
	if referer == "" {
		referer = "/"
	}
	c.Redirect(http.StatusFound, referer)
}

// HealthCheck responde o status operacional do servidor Go
func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"service":   "portfolio_go",
		"uptime":    time.Since(startTime).String(),
		"runtime":   "golang",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

// ProjectDetail retorna os dados detalhados de um projeto em JSON
func (h *Handler) ProjectDetail(c *gin.Context) {
	slug := c.Param("slug")
	lang := h.resolveLanguage(c)
	pageData := data.GetData(lang)

	for _, p := range pageData.Projects {
		if p.Slug == slug {
			c.JSON(http.StatusOK, p)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
}

// HomelabStatus verifica a conectividade dos containers na rede homelab
func (h *Handler) HomelabStatus(c *gin.Context) {
	services := []gin.H{
		checkService("portfolio_go", "8080", "Portfólio Go", true),
		checkService("inv_backend", "8080", "FinHub Backend", false),
		checkService("inv_frontend", "80", "FinHub Frontend", false),
		checkService("manga_linker_backend", "8080", "Manga Backend", false),
		checkService("manga_linker_frontend", "5173", "Manga Frontend", false),
	}

	c.JSON(http.StatusOK, gin.H{
		"host":       "Ubuntu Server Homelab",
		"checked_at": time.Now().Format("15:04:05"),
		"services":   services,
	})
}

func checkService(host, port, displayName string, isSelf bool) gin.H {
	if isSelf {
		return gin.H{
			"name":    displayName,
			"target":  fmt.Sprintf("%s:%s", host, port),
			"status":  "online",
			"latency": "< 1ms",
		}
	}

	// Tenta conexão TCP com timeout de 150ms
	address := fmt.Sprintf("%s:%s", host, port)
	start := time.Now()
	conn, err := net.DialTimeout("tcp", address, 150*time.Millisecond)
	if err != nil {
		return gin.H{
			"name":    displayName,
			"target":  address,
			"status":  "configured", // Container configurado na rede Docker
			"latency": "isolated",
		}
	}
	defer conn.Close()
	latency := time.Since(start).Round(time.Millisecond).String()

	return gin.H{
		"name":    displayName,
		"target":  address,
		"status":  "online",
		"latency": latency,
	}
}
