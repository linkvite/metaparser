package metaParser

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func ParseLink() (result MetaData, badRequest error) {
	log("👋 Enter the url of the web page 👇")

	//get the link from the user.
	link := getLink()

	//start the http client.
	client := &http.Client{}

	//validate the link
	domain, err := validateLink(link)

	if err != nil {
		return result, err
	}

	start := time.Now()
	metaData := NewMetaData()
	log("✅ Valid URL provided.")
	log("✅ Generated metadata template.")

	//fetch the html from the url
	req, err := http.NewRequest("GET", link, nil)

	//add the twitterbot header to access many websites.
	req.Header.Set("User-Agent", "Twitterbot/1.0")

	//add a request timeout
	client.Timeout = time.Second * 10

	if err != nil {
		result := returnResultWithError(err, metaData, link, domain)
		return result, nil
	}

	//parse the response
	resp, err := client.Do(req)

	if err != nil {
		result := returnResultWithError(err, metaData, link, domain)
		return result, nil
	}

	//close after the request is done.
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err := errors.New(resp.Status)
		result := returnResultWithError(err, metaData, link, domain)
		return result, nil
	}

	//parse the html document
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		result := returnResultWithError(err, metaData, link, domain)
		return result, nil
	}

	//update the metadata
	log("⏳ Updating metadata from html document...")
	metaData.generateMetaData(doc, link, domain)
	log("✅ Updated metadata from html document.")

	end := time.Now()
	elapsed := end.Sub(start)
	log(fmt.Sprintf("⏱  Total time taken: %d milliseconds.", elapsed.Milliseconds()))

	return *metaData, nil
}
