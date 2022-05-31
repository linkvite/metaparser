package parser

import (
	"fmt"
	"net/url"
	"strings"
)

// Get the link from the user.
func GetLink() string {
	log("👋 Enter the url of the web page 👇")

	var link string

	fmt.Scanln(&link)

	return link
}

// Validate the link
//
// Return the domain name if valid or error if invalid.
func ValidateLink(link string) (l string, e error) {
	// Parse the url, also validates the url.
	u, err := url.Parse(link)

	if err != nil {
		return "", PrintError("please provide a valid url")
	}

	// If the scheme is not http or https, exit.
	if u.Scheme != "http" && u.Scheme != "https" {
		return "", PrintError("the url must be of scheme http or https")
	}

	// Split the hostname to get only the domain.
	host := u.Hostname()
	hostSplit := strings.Split(host, ".")

	// If the hostname is less than 2, exit.
	minLength := 2
	if len(hostSplit) < minLength {
		return "", PrintError("please provide a valid url")
	}

	// Get the domain.
	domain := hostSplit[len(hostSplit)-2] + "." + hostSplit[len(hostSplit)-1]

	return domain, nil
}

// Get the base url of a link.
func getBaseURL(link string) string {
	u, err := url.Parse(link)

	if err != nil {
		message := PrintError(err.Error())
		panic(message)
	}

	baseURL := u.Scheme + "://" + u.Host

	return baseURL
}

// Resolve urls by adding the root url to the passed string if necessary.
func resolveURL(baseURL, data string) string {
	if data == "" {
		return ""
	}

	// If it starts with "//", remove it.
	if strings.HasPrefix(data, "//") {
		return strings.Replace(data, "//", "", 1)
	}

	// If the url is relative, then add the baseURL to it.
	if strings.HasPrefix(data, "/") {
		return baseURL + data
	}

	return data
}
