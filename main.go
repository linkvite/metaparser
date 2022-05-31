package parser

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Return the metadata of the page or error.
func ParseLink(link string) (MetaData, error) {
	// Start the http client.
	client := &http.Client{}

	// Validate the link.
	domain, err := ValidateLink(link)

	if err != nil {
		panic(err)
	}

	// Set the URL and domain to parse.
	linkData := NewLink(link, domain)

	start := time.Now()
	metaData := NewMetaData()

	log("✅ Valid URL provided.")
	log("✅ Generated metadata template.")

	// Fetch the html from the url.
	req, err := http.NewRequest("GET", link, http.NoBody)

	// Add the twitterbot header to access many websites.
	req.Header.Set("User-Agent", "Twitterbot/1.0")

	// Timeout after 10 seconds.
	const secondsTimeout = 10
	client.Timeout = time.Second * secondsTimeout

	if err != nil {
		result := returnResultWithError(err.Error(), metaData, linkData)

		return result, nil
	}

	// Parse the response.
	resp, err := client.Do(req)

	if err != nil {
		result := returnResultWithError(err.Error(), metaData, linkData)

		return result, nil
	}

	// Close after the request is done.
	defer resp.Body.Close()

	validStatusCode := 200
	if resp.StatusCode != validStatusCode {
		message := fmt.Sprintf(resp.Status)
		result := returnResultWithError(message, metaData, linkData)

		return result, nil
	}

	// Parse the html document.
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		result := returnResultWithError(err.Error(), metaData, linkData)

		return result, nil
	}

	// Update the metadata.
	log("⏳ Updating metadata from html document...")
	metaData.generateMetaData(doc, linkData)

	log("✅ Updated metadata from html document.")

	end := time.Now()
	elapsed := end.Sub(start)
	log(fmt.Sprintf("⏱  Total time taken: %d milliseconds.", elapsed.Milliseconds()))

	return *metaData, nil
}
