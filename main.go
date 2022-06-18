package parser

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Return the metadata of the page or error.
func ParseLink(p Parameters) (MetaData, error) {
	// Get and Set the default parameters.
	p = setDefaultParameters(p)

	// Start the http client.
	client := &http.Client{
		Timeout: time.Second * time.Duration(p.Timeout),
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if p.AllowRedirect {
				return nil
			}
			return errors.New("redirect not allowed. Pass `AllowRedirect: true` if you want to allow redirects")
		},
	}

	// Validate the URL.
	domain, err := ValidateLink(p.URL)

	if err != nil {
		panic(err)
	}

	// Set the URL and domain to parse.
	linkData := NewLink(p.URL, domain)

	start := time.Now()
	metaData := NewMetaData()

	log("✅ Valid URL provided.")
	log("✅ Generated metadata template.")

	// Fetch the html from the url.
	req, err := http.NewRequest("GET", p.URL, http.NoBody)

	// Add the headers.
	for k, v := range p.Headers {
		req.Header.Set(k, v)
	}

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
