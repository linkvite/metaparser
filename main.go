package main

import (
	"fmt"
	"strings"
	"time"

	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func parseLink(c *gin.Context) {
	link, err := getLink(c)

	if err != nil {
		serverError(err.Error(), c)
		return
	}

	//start the http client.
	client := &http.Client{}

	//parse the url, also validates the url.
	u, err := url.Parse(link)

	if err != nil {
		serverError("Please provide a valid url.", c)
		return
	}

	println("The link is: " + link)

	//if the scheme is not http or https, exit.
	if u.Scheme != "http" && u.Scheme != "https" {
		serverError("The url must be a http or https url", c)
		return
	}

	//split the hostname to get only the domain
	host := u.Hostname()
	hostSplit := strings.Split(host, ".")

	//if the hostname is less than 2, exit
	if len(hostSplit) < 2 {
		serverError("Please provide a valid url.", c)
		return
	}

	metaData := NewMetaData()

	//get the domain
	domain := hostSplit[len(hostSplit)-2] + "." + hostSplit[len(hostSplit)-1]

	//Record the execution time in milliseconds.
	start := time.Now()
	req, err := http.NewRequest("GET", link, nil)

	//add the twitterbot header to access many websites.
	req.Header.Set("User-Agent", "Twitterbot/1.0")

	//add a request timeout
	client.Timeout = time.Second * 10

	if err != nil {
		err := err.Error()
		returnResultWithError(err, metaData, link, domain, start, c)
		return
	}

	//parse the response
	resp, err := client.Do(req)

	if err != nil {
		err := err.Error()
		returnResultWithError(err, metaData, link, domain, start, c)
		return
	}

	//close after the request is done.
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err := fmt.Sprintf("%s", resp.Status)
		returnResultWithError(err, metaData, link, domain, start, c)
		return
	}

	//parse the html document
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		err := err.Error()
		returnResultWithError(err, metaData, link, domain, start, c)
		return
	}

	//update the meta data.
	metaData.generateMetaData(doc, link, domain)

	returnResult(metaData, start, c)

	return
}

func main() {
	router := gin.Default()
	router.GET("/api/parse", parseLink)
	router.Run("localhost:8080")
}
