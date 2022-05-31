package parser

type MetaData struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	SiteName    string   `json:"name"`
	Domain      string   `json:"domain"`
	URL         string   `json:"url"`
	SiteType    string   `json:"type"`
	Images      []string `json:"images"`
	Favicons    []string `json:"favicons"`
}

type Link struct {
	URL    string
	Domain string
}

// Instantiate a new metadata object.
func NewMetaData() *MetaData {
	return &MetaData{
		Title:       "",
		Description: "",
		Domain:      "",
		URL:         "",
		SiteName:    "",
		SiteType:    "",
		Images:      []string{},
		Favicons:    []string{},
	}
}

// Instantiate a new link object.
func NewLink(link, domain string) *Link {
	return &Link{
		URL:    link,
		Domain: domain,
	}
}
