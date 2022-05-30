package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

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

//Instantiate a new meta data object.
func NewMetaData() *MetaData {
	return &MetaData{
		SiteName: "",
		Title:    "",
		Desc:     "",
		Domain:   "",
		Url:      "",
		SiteType: "",
		Images:   []string{},
		Favicons: []string{},
	}
}

//Update the default meta data based on the type of tag.
func (m *MetaData) setValueByType(tag, value, domain string) {
	switch tag {
	case "og:site_name":
		m.SiteName = setValue(value, domain)

	case "og:type":
		m.SiteType = setValue(value, "website")
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
	title, description := getTitleAndDescription(doc)

	//set the default values.
	m.setDefaultValue(url, domain, title)

	//set the description.
	m.Desc = description

	//find the favicons.
	m.Favicons = getFavicons(doc, url)

	//set the images.
	m.Images = getImages(doc, url)

	//the res of the tags.
	tagsRemaining := []string{
		"og:type",
		"og:site_name",
	}

	for _, tag := range tagsRemaining {
		value := getMetaContent(doc, tag)
		m.setValueByType(tag, value, domain)
	}
}

//get the favicons from the page.
func getFavicons(doc *goquery.Document, link string) []string {
	baseURL := getBaseUrl(link)

	var favicons []string
	iconSelectors := []string{
		"shortcut icon",
		"apple-touch-icon",
	}

	//other icons like apple touch icons.
	for _, iconSelector := range iconSelectors {
		icon := doc.Find(fmt.Sprintf("link[rel='%s']", iconSelector)).AttrOr("href", "")

		if icon != "" {
			favicons = append(favicons, resolveUrl(baseURL, icon))
		}
	}

	//regular icons.
	normalIcons := doc.Find(fmt.Sprintf("link[rel='%s']", "icon")).
		Map(func(i int, s *goquery.Selection) string {
			return resolveUrl(baseURL, s.AttrOr("href", ""))
		})

	favicons = append(favicons, normalIcons...)

	return favicons
}

//get a meta tag content by property name.
// Example:
//
// image := getMetaContent(doc, "og:image")
func getMetaContent(doc *goquery.Document, data string) string {
	//find tag content by property.
	property := doc.Find(fmt.Sprintf("meta[property='%s']", data)).AttrOr("content", "")

	if property != "" {
		return property
	}

	//find tag content by name.
	return doc.Find(fmt.Sprintf("meta[name='%s']", data)).AttrOr("content", "")
}

//get title and description from the page.
func getTitleAndDescription(doc *goquery.Document) (string, string) {
	//default
	title := getMetaContent(doc, "og:title")
	desc := getMetaContent(doc, "og:description")

	//fallbacks
	if title == "" {
		title = doc.Find("title").Text()
	}

	if desc == "" {
		desc = getMetaContent(doc, "description")
	}

	return title, desc
}

//get images from the page.
func getImages(doc *goquery.Document, link string) []string {
	baseURL := getBaseUrl(link)
	var images []string

	defaultImage := doc.Find(fmt.Sprintf("link[rel='%s']", "image_src")).AttrOr("href", "")
	otherImage := getMetaContent(doc, "og:image")

	resolvedImages := []string{
		resolveUrl(baseURL, defaultImage),
		resolveUrl(baseURL, otherImage),
	}

	//only add non-empty images
	for _, image := range resolvedImages {
		if image != "" {
			images = append(images, image)
		}
	}

	return images
}
