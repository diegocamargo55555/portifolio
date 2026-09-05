package models

// Profile representa os dados do desenvolvedor
type Profile struct {
	Name             string   `json:"name"`
	Title            string   `json:"title"`
	Subtitle         string   `json:"subtitle"`
	BioShort         string   `json:"bio_short"`
	BioLong          string   `json:"bio_long"`
	Email            string   `json:"email"`
	GitHub           string   `json:"github"`
	GitHubURL        string   `json:"github_url"`
	LinkedIn         string   `json:"linkedin"`
	Location         string   `json:"location"`
	AvailableForWork bool     `json:"available_for_work"`
	KeyHighlights    []string `json:"key_highlights"`
}

// TechItem representa uma tecnologia ou ferramenta
type TechItem struct {
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Level string `json:"level"`
}

// SkillCategory agrupa habilidades por especialidade
type SkillCategory struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Icon        string     `json:"icon"`
	Skills      []TechItem `json:"skills"`
}

// Project representa um projeto em destaque
type Project struct {
	ID                  string   `json:"id"`
	Slug                string   `json:"slug"`
	Title               string   `json:"title"`
	Category            string   `json:"category"`
	Tagline             string   `json:"tagline"`
	Description         string   `json:"description"`
	LongDescription     string   `json:"long_description"`
	ArchitectureSummary string   `json:"architecture_summary"`
	KeyFeatures         []string `json:"key_features"`
	TechStack           []string `json:"tech_stack"`
	GitHubURL           string   `json:"github_url"`
	LiveURL             string   `json:"live_url"`
	Subdomain           string   `json:"subdomain"`
	InternalPort        string   `json:"internal_port"`
	ContainerName       string   `json:"container_name"`
	BadgeColor          string   `json:"badge_color"`
	ArchitectureDetails []string `json:"architecture_details"`
}

// HomelabService representa um serviço ativo no Homelab Ubuntu Server
type HomelabService struct {
	Name          string `json:"name"`
	Role          string `json:"role"`
	ContainerName string `json:"container_name"`
	Stack         string `json:"stack"`
	Port          string `json:"port"`
	Subdomain     string `json:"subdomain"`
	Status        string `json:"status"` // "online", "running", "isolated"
	RAMUsage      string `json:"ram_usage"`
}

// Translations contém os textos da interface do usuário
type Translations struct {
	NavAbout               string `json:"nav_about"`
	NavProjects            string `json:"nav_projects"`
	NavSkills              string `json:"nav_skills"`
	NavHomelab             string `json:"nav_homelab"`
	NavContact             string `json:"nav_contact"`
	HeroBadge              string `json:"hero_badge"`
	HeroTitlePrefix        string `json:"hero_title_prefix"`
	HeroTitleHighlight     string `json:"hero_title_highlight"`
	HeroSubtitle           string `json:"hero_subtitle"`
	BtnProjects            string `json:"btn_projects"`
	BtnContact             string `json:"btn_contact"`
	BtnLiveDemo            string `json:"btn_live_demo"`
	BtnSourceCode          string `json:"btn_source_code"`
	BtnArchDeepDive        string `json:"btn_arch_deep_dive"`
	SectionProjectsTitle   string `json:"section_projects_title"`
	SectionProjectsSub     string `json:"section_projects_sub"`
	SectionSkillsTitle     string `json:"section_skills_title"`
	SectionSkillsSub       string `json:"section_skills_sub"`
	SectionHomelabTitle    string `json:"section_homelab_title"`
	SectionHomelabSub      string `json:"section_homelab_sub"`
	SectionAboutTitle      string `json:"section_about_title"`
	SectionAboutSub        string `json:"section_about_sub"`
	SectionContactTitle    string `json:"section_contact_title"`
	SectionContactSub      string `json:"section_contact_sub"`
	CopyEmail              string `json:"copy_email"`
	EmailCopied            string `json:"email_copied"`
	CloseModal             string `json:"close_modal"`
	ArchitectureModalTitle string `json:"architecture_modal_title"`
	LiveContainerStatus    string `json:"live_container_status"`
	SelfHostedBadge        string `json:"self_hosted_badge"`
}

// PageData estrutura todos os dados enviados ao template HTML
type PageData struct {
	CurrentLang     string           `json:"current_lang"`
	T               Translations     `json:"t"`
	Profile         Profile          `json:"profile"`
	Projects        []Project        `json:"projects"`
	Skills          []SkillCategory  `json:"skills"`
	HomelabServices []HomelabService `json:"homelab_services"`
	HostInfo        struct {
		OS      string `json:"os"`
		Runtime string `json:"runtime"`
		Engine  string `json:"engine"`
		Tunnel  string `json:"tunnel"`
	} `json:"host_info"`
}
