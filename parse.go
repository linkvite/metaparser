package main

import (
	"github.com/PuerkitoBio/goquery"
)

type MetaData struct {
	SiteName string   `json:"name"`
	Title    string   `json:"title"`
	Desc     string   `json:"description"`
	Domain   string   `json:"domain"`
	Image    string   `json:"image"`
	Url      string   `json:"url"`
	SiteType string   `json:"type"`
	Keywords []string `json:"keywords"`
}

//Instantiate a new meta data object.
func NewMetaData() *MetaData {
	return &MetaData{
		SiteName: "",
		Title:    "",
		Desc:     "",
		Domain:   "",
		Image:    "",
		Url:      "",
		SiteType: "",
		Keywords: []string{},
	}
}

//Update the default meta data based on the type of tag.
func (m *MetaData) setValueByType(tag, value string) {
	switch tag {
	case "description":
		m.Desc = value
	case "og:site_name":
		m.SiteName = value
	case "og:title":
		m.Title = value
	case "og:description":
		m.Desc = value
	case "og:image":
		m.Image = value
	case "og:url":
		m.Url = value
	case "og:type":
		m.SiteType = value
	case "keywords":
		m.Keywords = append(m.Keywords, value)
	}

	return
}

//set default values based on the link provided.
func (m *MetaData) setDefaultValue(url, domain, title string) {
	m.Url = url
	m.Title = title
	m.Domain = domain
	m.SiteName = domain
	m.SiteType = "website"

	return
}

//generate meta data from the page.
func (m *MetaData) generateMetaData(doc *goquery.Document, url, domain string) {
	title := doc.Find("title").Text()
	m.setDefaultValue(url, domain, title)

	//find the meta tags.
	metaTags := doc.Find("meta").Nodes

	//the header tags we only need.
	tagsToCheck := []string{
		"og:site_name",
		"og:title",
		"og:description",
		"og:image",
		"og:url",
		"og:type",
		"description",
		"keywords",
		"title",
	}

	for _, tag := range metaTags {
		arr := tag.Attr

		for j := 0; j < len(arr); j++ {
			//if the key is content, move to beginning of array
			if arr[j].Key == "content" {
				arr[0], arr[j] = arr[j], arr[0]
			} else if arr[j].Key == "name" || arr[j].Key == "property" {
				arr[1], arr[j] = arr[j], arr[1]
			} else {
				continue
			}
		}

		//if the array is less that 2, move on.
		if len(arr) < 2 {
			continue
		}

		//only check for needed meta tags
		for _, tagToCheck := range tagsToCheck {
			if tag.Attr[1].Val == tagToCheck {
				m.setValueByType(tagToCheck, tag.Attr[0].Val)
			}
		}
	}
}
