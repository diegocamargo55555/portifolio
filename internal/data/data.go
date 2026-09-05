package data

import "github.com/diegocamargo55555/portfolio/internal/models"

// GetData retorna todos os dados do portfólio no idioma especificado ("pt" ou "en")
func GetData(lang string) models.PageData {
	if lang == "en" {
		return getEnglishData()
	}
	return getPortugueseData()
}

func getPortugueseData() models.PageData {
	return models.PageData{
		CurrentLang: "pt",
		T: models.Translations{
			NavAbout:               "Sobre",
			NavProjects:            "Projetos",
			NavSkills:              "Habilidades",
			NavHomelab:             "Homelab",
			NavContact:             "Contato",
			HeroBadge:              "Disponível para novos projetos e oportunidades",
			HeroTitlePrefix:        "Olá, eu sou",
			HeroTitleHighlight:     "Diego Camargo",
			HeroSubtitle:           "Engenheiro de Software & Desenvolvedor Full-Stack focado em Golang, Clean Architecture, React e Infraestrutura Homelab com Docker no Ubuntu Server.",
			BtnProjects:            "Ver Projetos",
			BtnContact:             "Entrar em Contato",
			BtnLiveDemo:            "Acessar Demo",
			BtnSourceCode:          "Código Fonte",
			BtnArchDeepDive:        "Ver Arquitetura",
			SectionProjectsTitle:   "Projetos em Destaque",
			SectionProjectsSub:     "Aplicações completas rodando em produção no meu Homelab com arquitetura resiliente.",
			SectionSkillsTitle:     "Stack Tecnológica",
			SectionSkillsSub:       "Tecnologias, linguagens e ferramentas aplicadas diariamente na resolução de problemas complexos.",
			SectionHomelabTitle:    "Arquitetura Homelab no Notebook",
			SectionHomelabSub:      "Como este portfólio e os projetos rodam simultaneamente em um notebook Ubuntu Server via Docker e Cloudflare Tunnel.",
			SectionAboutTitle:      "Sobre Mim",
			SectionAboutSub:        "Engenharia de software focada em código limpo, confiabilidade e performance.",
			SectionContactTitle:    "Vamos Conversar?",
			SectionContactSub:      "Estou aberto a oportunidades profissionais, parcerias e projetos desafiadores.",
			CopyEmail:              "Copiar E-mail",
			EmailCopied:            "E-mail copiado com sucesso!",
			CloseModal:             "Fechar",
			ArchitectureModalTitle: "Detalhamento de Engenharia & Arquitetura",
			LiveContainerStatus:    "Status do Container",
			SelfHostedBadge:        "Self-Hosted no Ubuntu Server",
		},
		Profile: models.Profile{
			Name:             "Diego Camargo",
			Title:            "Engenheiro de Software & Desenvolvedor Full-Stack",
			Subtitle:         "Especialista em Golang, Clean Architecture, Bancos Relacionais e Containers",
			BioShort:         "Desenvolvedor apaixonado por construir sistemas robustos, eficientes e escaláveis utilizando Golang no backend e React no frontend.",
			BioLong:          "Sou desenvolvedor focado em soluções backend de alta performance com Golang e Clean Architecture, interfaces modernas e responsivas com React/TypeScript, e automação de infraestrutura com Docker e Linux (Ubuntu Server). Tenho grande apreço por precisão de dados (como cálculos monetários de ponto fixo), caching distribuído com Redis, modelagem relacional em PostgreSQL e self-hosting.",
			Email:            "diegocamargo55555@gmail.com",
			GitHub:           "diegocamargo55555",
			GitHubURL:        "https://github.com/diegocamargo55555",
			LinkedIn:         "https://linkedin.com/in/",
			Location:         "Brasil",
			AvailableForWork: true,
			KeyHighlights: []string{
				"Backend robusto com Golang, Gin e Clean Architecture",
				"Modelagem e otimização em PostgreSQL e Redis",
				"Consumo e agregação de APIs financeiras em tempo real",
				"Infraestrutura e orquestração Homelab em Linux/Docker",
			},
		},
		Projects: []models.Project{
			{
				ID:                  "finhub",
				Slug:                "finhub",
				Title:               "FinHub (inv)",
				Category:            "Finanças & Investimentos",
				Tagline:             "Plataforma completa de controle financeiro pessoal e consolidação de investimentos.",
				Description:         "Sistema completo de gestão patrimonial com cotações de mercado em tempo real (Brapi/CoinGecko), cálculo financeiro com precisão arbitrária e cache distribuído.",
				LongDescription:     "O FinHub é uma plataforma de alta precisão desenhada para solucionar as dores do investidor moderno. Ele unifica controle de despesas e receitas com a custódia e rentabilidade de carteiras de ações B3, fundos imobiliários (FIIs) e criptomoedas. Implementa Clean Architecture no backend Go com camadas estritas de Domain, Repository e Use Cases.",
				ArchitectureSummary: "Golang (Gin + GORM) em Clean Architecture, React + TypeScript + Vite, PostgreSQL 16 e Redis 7.",
				KeyFeatures: []string{
					"Precisão Monetária Arbitrária: Uso de shopspring/decimal para eliminar qualquer erro de arredondamento IEEE 754 em transações e dividendos.",
					"Cotações de Mercado em Tempo Real: Integração com APIs externas (Brapi e CoinGecko) com camada de cache inteligente em Redis para mitigar rate-limits.",
					"Autenticação e Segurança: Tokens JWT (Access + Refresh) com hashing Bcrypt e proteção CSRF/CORS.",
					"Auto-Migrate & Seeds: Estrutura do PostgreSQL versionada e populada de forma transparente na inicialização.",
				},
				TechStack:           []string{"Golang", "Gin", "GORM", "PostgreSQL 16", "Redis 7", "React", "TypeScript", "Tailwind CSS", "Docker"},
				GitHubURL:           "https://github.com/diegocamargo55555/inv",
				LiveURL:             "https://finhub.diegocamargo.dev",
				Subdomain:           "finhub",
				InternalPort:        "3001",
				ContainerName:       "inv_frontend / inv_backend",
				BadgeColor:          "emerald",
				ArchitectureDetails: []string{
					"Clean Architecture (Camadas desacopladas: Domain -> Usecase -> Repository -> Handler)",
					"shopspring/decimal garantindo fidelidade contábil estrita",
					"PostgreSQL 16 com relacionamentos consistentes e índices otimizados",
					"Redis 7 com TTL dinâmico para cache de cotações B3 e Cripto",
				},
			},
			{
				ID:                  "manga-wrapper",
				Slug:                "manga-wrapper",
				Title:               "Manga Wrapper",
				Category:            "Consumo de Conteúdo & APIs",
				Tagline:             "Rastreador inteligente de leitura e wrapper de agregação de mangás.",
				Description:         "Aplicação focada em experiência do usuário para catalogar leituras, acompanhar lançamentos de capítulos e manter histórico de progresso em tempo real.",
				LongDescription:     "O Manga Wrapper foi criado para organizar e centralizar o consumo de mangás e quadrinhos digitais. Possui uma API veloz construída em Golang com persistência relacional em PostgreSQL e interface SPA moderna em React + Vite, com controle de estado fluido e autenticação segura de usuários.",
				ArchitectureSummary: "Golang REST API, React + Vite SPA, PostgreSQL 15, Docker & Hot-Reload com Air.",
				KeyFeatures: []string{
					"Acompanhamento de Leitura: Marcadores de capítulos lidos e alertas de atualização.",
					"API RESTful em Go: Endpoints rápidos e leves para consulta de biblioteca e metadados.",
					"Autenticação JWT: Sessões criptografadas com tokens de autenticação sem estado.",
					"Interface Otimizada para Mobile e Desktop: Leitura agradável e navegação fluida.",
				},
				TechStack:           []string{"Golang", "Gin", "PostgreSQL 15", "React", "Vite", "JWT", "Docker"},
				GitHubURL:           "https://github.com/diegocamargo55555/manga-wrapper",
				LiveURL:             "https://manga.diegocamargo.dev",
				Subdomain:           "manga",
				InternalPort:        "3002",
				ContainerName:       "manga_linker_frontend / manga_linker_backend",
				BadgeColor:          "cyan",
				ArchitectureDetails: []string{
					"API REST em Go com arquitetura orientada a serviços",
					"Banco relacional PostgreSQL 15 com integridade referencial",
					"Frontend React modular com Vite para carregamento ultrarrápido",
					"Containerização com Docker Compose e ambiente de desenvolvimento isolado",
				},
			},
		},
		Skills: []models.SkillCategory{
			{
				Title:       "Backend & Engenharia",
				Description: "Construção de APIs escaláveis, alta concorrência e arquitetura limpa.",
				Icon:        "server",
				Skills: []models.TechItem{
					{Name: "Golang", Icon: "go", Level: "Avançado"},
					{Name: "Gin Gonic", Icon: "gin", Level: "Avançado"},
					{Name: "Clean Architecture", Icon: "layers", Level: "Avançado"},
					{Name: "GORM & SQL Nativo", Icon: "database", Level: "Avançado"},
					{Name: "APIs RESTful & JWT", Icon: "shield", Level: "Avançado"},
					{Name: "Precisão Decimal / Finanças", Icon: "calculator", Level: "Especialista"},
				},
			},
			{
				Title:       "Frontend & Interface",
				Description: "Desenvolvimento de interfaces modernas, reativas e com foco no usuário.",
				Icon:        "layout",
				Skills: []models.TechItem{
					{Name: "React", Icon: "react", Level: "Avançado"},
					{Name: "TypeScript", Icon: "ts", Level: "Intermediário/Avançado"},
					{Name: "Tailwind CSS", Icon: "tailwind", Level: "Avançado"},
					{Name: "Vite", Icon: "vite", Level: "Avançado"},
					{Name: "HTML5 & CSS Moderno", Icon: "code", Level: "Avançado"},
				},
			},
			{
				Title:       "Bancos de Dados & Cache",
				Description: "Modelagem, persistência relacional e estratégias de caching.",
				Icon:        "database",
				Skills: []models.TechItem{
					{Name: "PostgreSQL", Icon: "postgres", Level: "Avançado"},
					{Name: "Redis Cache", Icon: "redis", Level: "Avançado"},
					{Name: "Data Modeling & Migrations", Icon: "table", Level: "Avançado"},
					{Name: "Transações ACID", Icon: "lock", Level: "Avançado"},
				},
			},
			{
				Title:       "DevOps & Homelab",
				Description: "Infraestrutura Linux, containers Docker e redes seguras.",
				Icon:        "cpu",
				Skills: []models.TechItem{
					{Name: "Docker & Docker Compose", Icon: "docker", Level: "Avançado"},
					{Name: "Ubuntu Server (Linux)", Icon: "terminal", Level: "Avançado"},
					{Name: "Cloudflare Zero Trust / Tunnels", Icon: "cloud", Level: "Avançado"},
					{Name: "Nginx Reverse Proxy", Icon: "network", Level: "Intermediário/Avançado"},
					{Name: "Git & GitHub Workflows", Icon: "git", Level: "Avançado"},
				},
			},
		},
		HomelabServices: []models.HomelabService{
			{
				Name:          "Portfólio Go",
				Role:          "Showcase Profissional & Gateway",
				ContainerName: "portfolio_go",
				Stack:         "Golang 1.22 + Gin + embed.FS",
				Port:          "8080 (Interna)",
				Subdomain:     "portfolio.diegocamargo.dev",
				Status:        "online",
				RAMUsage:      "~12 MB",
			},
			{
				Name:          "FinHub (inv)",
				Role:          "Gestão Financeira & Investimentos",
				ContainerName: "inv_frontend / inv_backend",
				Stack:         "Go + React + Postgres 16 + Redis 7",
				Port:          "3001 (Interna)",
				Subdomain:     "finhub.diegocamargo.dev",
				Status:        "running",
				RAMUsage:      "~85 MB",
			},
			{
				Name:          "Manga Wrapper",
				Role:          "Rastreador & Leitor de Mangás",
				ContainerName: "manga_frontend / manga_backend",
				Stack:         "Go + React + Postgres 15",
				Port:          "3002 (Interna)",
				Subdomain:     "manga.diegocamargo.dev",
				Status:        "running",
				RAMUsage:      "~45 MB",
			},
			{
				Name:          "Cloudflare Tunnel",
				Role:          "Ingress Seguro & SSL Gratuito",
				ContainerName: "cloudflared",
				Stack:         "Cloudflare Daemon (Sem Portas Abertas)",
				Port:          "Saída TLS 443",
				Subdomain:     "*.diegocamargo.dev",
				Status:        "online",
				RAMUsage:      "~22 MB",
			},
		},
	}
}

func getEnglishData() models.PageData {
	pt := getPortugueseData()
	pt.CurrentLang = "en"
	pt.T = models.Translations{
		NavAbout:               "About",
		NavProjects:            "Projects",
		NavSkills:              "Skills",
		NavHomelab:             "Homelab",
		NavContact:             "Contact",
		HeroBadge:              "Available for new opportunities and engineering projects",
		HeroTitlePrefix:        "Hi, I am",
		HeroTitleHighlight:     "Diego Camargo",
		HeroSubtitle:           "Software Engineer & Full-Stack Developer specializing in Golang, Clean Architecture, React, and Homelab Infrastructure on Ubuntu Server.",
		BtnProjects:            "View Projects",
		BtnContact:             "Get in Touch",
		BtnLiveDemo:            "Live Demo",
		BtnSourceCode:          "Source Code",
		BtnArchDeepDive:        "Architecture Deep Dive",
		SectionProjectsTitle:   "Featured Engineering Projects",
		SectionProjectsSub:     "Full-stack production applications running simultaneously on my self-hosted homelab.",
		SectionSkillsTitle:     "Technical Stack",
		SectionSkillsSub:       "Core languages, frameworks, databases, and DevOps tools I use to solve complex problems.",
		SectionHomelabTitle:    "Ubuntu Server Homelab Architecture",
		SectionHomelabSub:      "How this portfolio and projects run harmoniously on an Ubuntu Server notebook using Docker and Cloudflare Tunnel.",
		SectionAboutTitle:      "About Me",
		SectionAboutSub:        "Driven by clean code, high reliability, data precision, and performance.",
		SectionContactTitle:    "Let's Connect",
		SectionContactSub:      "Feel free to reach out for software engineering roles, collaborations, or tech talks.",
		CopyEmail:              "Copy Email",
		EmailCopied:            "Email copied to clipboard!",
		CloseModal:             "Close",
		ArchitectureModalTitle: "Architecture & Engineering Deep Dive",
		LiveContainerStatus:    "Container Status",
		SelfHostedBadge:        "Self-Hosted on Ubuntu Server",
	}

	pt.Profile.Title = "Software Engineer & Full-Stack Developer"
	pt.Profile.Subtitle = "Specialized in Golang, Clean Architecture, Relational Databases & Containers"
	pt.Profile.BioShort = "Passionate software engineer building robust, scalable and clean backend services in Go paired with modern React frontends."
	pt.Profile.BioLong = "I am a software engineer focused on high-performance backend systems with Golang and Clean Architecture, modern and responsive web applications with React/TypeScript, and infrastructure automation using Docker and Linux (Ubuntu Server). I care deeply about numeric precision in monetary operations, distributed caching with Redis, relational data integrity in PostgreSQL, and self-hosted environments."
	pt.Profile.KeyHighlights = []string{
		"High-performance backend systems with Go, Gin & Clean Architecture",
		"Advanced data modeling and indexing in PostgreSQL & Redis caching",
		"Real-time market data ingestion and arbitrary-precision financial math",
		"Homelab infrastructure and container orchestration on Linux/Docker",
	}

	pt.Projects[0].Tagline = "Comprehensive personal finance management and investment portfolio consolidation platform."
	pt.Projects[0].Description = "Full-featured wealth management system with real-time market data feeds (Brapi/CoinGecko), fixed-point monetary math, and distributed caching."
	pt.Projects[0].LongDescription = "FinHub is an arbitrary-precision wealth tracking engine built to solve portfolio consolidation challenges. It handles multi-asset tracking (stocks, REITs, crypto) with strict adherence to Clean Architecture across Domain, Repository, and Usecase layers."
	pt.Projects[0].KeyFeatures = []string{
		"Arbitrary Monetary Precision: Powered by shopspring/decimal to eliminate IEEE-754 floating-point inaccuracies.",
		"Real-time Market Quotes: Brapi and CoinGecko integrations with smart Redis TTL caching to prevent rate-limit throttling.",
		"Secure Authentication: JWT access and refresh token lifecycle paired with Bcrypt salted hashing.",
		"Automated Migrations & Seeding: Schema evolution managed gracefully on boot.",
	}

	pt.Projects[1].Tagline = "Smart reading tracker and manga aggregation wrapper."
	pt.Projects[1].Description = "A streamlined application designed for comic tracking, reading status synchronization, and catalog exploration."
	pt.Projects[1].LongDescription = "Manga Wrapper organizes digital comic reading logs. It features a Golang REST API backend with PostgreSQL persistence and a fast React + Vite frontend."
	pt.Projects[1].KeyFeatures = []string{
		"Progress Synchronization: Bookmark management and release notifications.",
		"Performant Go REST API: Lightweight and responsive endpoints for library management.",
		"Stateless JWT Security: Token-based authentication across sessions.",
		"Responsive UI: Tailored for seamless reading on both desktop and mobile screens.",
	}

	return pt
}
