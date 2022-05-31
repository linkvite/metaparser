package parser

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// Update the default metadata based on the type of tag.
func (m *MetaData) setValueByType(tag, value, domain string) {
	switch tag {
	case "og:site_name":
		m.SiteName = setValue(value, domain)

	case "og:type":
		m.SiteType = setValue(value, "website")
	}
}

// Set default values based on the link provided.
func (m *MetaData) setDefaultValue(linkData *Link, title string) {
	m.Title = title
	m.URL = linkData.URL
	m.Domain = linkData.Domain
	m.SiteName = linkData.Domain
	m.SiteType = "website"
}

// Generate metadata from the page.
func (m *MetaData) generateMetaData(doc *goquery.Document, linkData *Link) {
	title, description := getTitleAndDescription(doc)

	// Set the default values.
	m.setDefaultValue(linkData, title)

	// Set the description.
	m.Description = description

	// Find the favicons.
	m.Favicons = getFavicons(doc, linkData.URL)

	// Set the images.
	m.Images = getImages(doc, linkData.URL)

	// Rest of the meta tags.
	tagsRemaining := []string{
		"og:type",
		"og:site_name",
	}

	for _, tag := range tagsRemaining {
		value := getMetaContent(doc, tag)
		m.setValueByType(tag, value, linkData.Domain)
	}
}

// Get the favicons from the page.
func getFavicons(doc *goquery.Document, link string) []string {
	var favicons []string

	baseURL := getBaseURL(link)
	iconSelectors := []string{
		"icon",
		"mask-icon",
		"shortcut icon",
		"apple-touch-icon",
	}

	// Parse the doc for the icons.
	for _, iconSelector := range iconSelectors {
		doc.Find(fmt.Sprintf("link[rel='%s']", iconSelector)).
			Each(func(i int, s *goquery.Selection) {
				icon := s.AttrOr("href", "")
				if icon != "" {
					favicons = append(favicons, resolveURL(baseURL, icon))
				}
			})
	}

	return favicons
}

// Get a meta tag content by property name.
// Example:
//
// image := getMetaContent(doc, "og:image").
func getMetaContent(doc *goquery.Document, data string) string {
	// Find tag content by property.
	property := doc.Find(fmt.Sprintf("meta[property='%s']", data)).AttrOr("content", "")

	if property != "" {
		return property
	}

	// Find tag content by name.
	return doc.Find(fmt.Sprintf("meta[name='%s']", data)).AttrOr("content", "")
}

// Get title and description from the page.
func getTitleAndDescription(doc *goquery.Document) (t, d string) {
	// Default.
	title := getMetaContent(doc, "og:title")
	desc := getMetaContent(doc, "og:description")

	// Fallbacks.
	if title == "" {
		title = doc.Find("title").Text()
	}

	if desc == "" {
		desc = getMetaContent(doc, "description")
	}

	return title, desc
}

// Get images from the page.
func getImages(doc *goquery.Document, link string) []string {
	var images []string

	baseURL := getBaseURL(link)
	options := []string{
		"og:image",
		"twitter:image",
		"msapplication-TileImage",
	}

	defaultImage := doc.Find(fmt.Sprintf("link[rel='%s']", "image_src")).AttrOr("href", "")

	resolvedImages := []string{
		resolveURL(baseURL, defaultImage),
	}

	for _, option := range options {
		image := getMetaContent(doc, option)
		resolvedImages = append(resolvedImages, resolveURL(baseURL, image))
	}

	// Only add non-empty images.
	for _, image := range resolvedImages {
		if image != "" {
			images = append(images, image)
		}
	}

	return images
}
