package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//return a successfully parsed meta data.
func returnResult(metaData interface{}, start time.Time, c *gin.Context) {
	end := time.Now()
	elapsed := end.Sub(start)
	time := elapsed.Milliseconds()

	c.JSON(http.StatusOK, gin.H{
		"data":     metaData,
		"timeInMs": time,
	})

	return
}

//handle creating a new meta data object if there's an error.
func returnResultWithError(err string, metaData *MetaData, link, domain string, start time.Time, c *gin.Context) {

	//generate a blank data.
	metaData.setDefaultValue(link, domain, domain)
	returnResult(metaData, start, c)

	return
}

//return a server error.
func serverError(errorX string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": errorX,
		"data":  nil,
	})

	return
}

type LinkRequestBody struct {
	Link string `json:"link"`
}

//Get the link from query or body parameters.
func getLink(c *gin.Context) (url string, badRequest error) {
	var link string

	//check the query string.
	if c.Query("link") != "" {
		link = c.Query("link")

		return link, nil
	}

	if c.Request.Body != nil {

		//check the request body
		var requestBody LinkRequestBody

		body := c.Request.Body
		x, _ := ioutil.ReadAll(body)

		err := json.Unmarshal(x, &requestBody)

		if err != nil {
			badRequest = errors.New("Please provide a valid url.")
			return
		}

		link = requestBody.Link

		return link, nil
	}

	badRequest = errors.New("Please provide a valid url.")
	return
}
