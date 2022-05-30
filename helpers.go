package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
)

//Covert struct to json
func ToJson(v interface{}) string {
	data, _ := json.MarshalIndent(v, "", "\t")
	return string(data)
}

//Log error.
func logError(err error) {
	log(fmt.Sprintf("❌ Failed to get the url. This is because of: %s", err))
	log("🤧 Generating blank data.")
}

//Log a message.
func log(msg interface{}) {
	fmt.Printf("%s\n", msg)
	fmt.Println("================================================================")
}

//Handle creating a new meta data object if there's an error.
func returnResultWithError(err error, metaData *MetaData, link, domain string) MetaData {
	logError(err)

	//generate a blank data.
	metaData.setDefaultValue(link, domain, domain)
	return *metaData
}

//Get the link from the user.
func getLink() string {
	var link string
	fmt.Scanln(&link)
	fmt.Println("================================================================")
	return link
}

//Validate the link
//
//Return the domain name if valid or error if invalid.
func validateLink(link string) (string, error) {
	//parse the url, also validates the url.
	u, err := url.Parse(link)

	if err != nil {
		badRequest := errors.New("Please provide a valid url.")
		return "", badRequest
	}

	//if the scheme is not http or https, exit.
	if u.Scheme != "http" && u.Scheme != "https" {
		badRequest := errors.New("The url must be a http or https url.")
		return "", badRequest
	}

	//split the hostname to get only the domain
	host := u.Hostname()
	hostSplit := strings.Split(host, ".")

	//if the hostname is less than 2, exit
	if len(hostSplit) < 2 {
		badRequest := errors.New("Please provide a valid url.")
		return "", badRequest
	}

	//get the domain
	domain := hostSplit[len(hostSplit)-2] + "." + hostSplit[len(hostSplit)-1]

	return domain, nil
}

//Get the base url or a link.
func getBaseUrl(link string) string {
	u, err := url.Parse(link)

	if err != nil {
		panic(err)
	}

	baseURL := u.Scheme + "://" + u.Host
	return baseURL
}

//Resolve urls by adding the root url to the passed string if necessary.
func resolveUrl(baseUrl, data string) string {
	if data == "" {
		return ""
	}

	//if the url is relative, then add the baseURL to it.
	if strings.HasPrefix(data, "/") {
		return baseUrl + data
	}

	return data
}

//Set value. if the value is empty, then set the default value.
func setValue(value, fallBack string) string {
	if value == "" {
		return fallBack
	}

	return value
}
