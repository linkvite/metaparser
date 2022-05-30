package parser

type MetaData struct {
	SiteName string   `json:"name"`
	Title    string   `json:"title"`
	Desc     string   `json:"description"`
	Domain   string   `json:"domain"`
	Url      string   `json:"url"`
	SiteType string   `json:"type"`
	Images   []string `json:"images"`
	Favicons []string `json:"favicons"`
}
