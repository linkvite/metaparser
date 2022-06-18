package parser

import (
	"encoding/json"
	"fmt"
)

// Covert struct to JSON.
func ToJSON(v interface{}) string {
	data, err := json.MarshalIndent(v, "", "\t")

	if err != nil {
		panic(err)
	}

	return string(data)
}

// Log a message or data.
func log(arg string) {
	fmt.Println(arg)
	fmt.Println("================================================================")
}

// If a value is empty, then set it to the default value.
func setValue(value, fallBack string) string {
	if value == "" {
		return fallBack
	}

	return value
}

// Remove duplicates from a slice.
func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] {
			continue
		}

		encountered[elements[v]] = true
		element := elements[v]
		result = append(result, element)
	}

	// Return the new slice.
	return result
}

// Set the default parameters.
func setDefaultParameters(p Parameters) Parameters {
	// Default Timeout after 10 seconds.
	timeout := 10
	if p.Timeout > 0 {
		timeout = p.Timeout
	}

	// Default AllowRedirect to false.
	allowRedirect := false
	if p.AllowRedirect {
		allowRedirect = true
	}

	headers := make(map[string]string)
	if p.Headers != nil {
		headers = p.Headers
	} else {
		// Add the twitterbot header to access many websites.
		headers["User-Agent"] = "Twitterbot/1.0"
	}

	URL := ""
	if p.URL != "" {
		URL = p.URL
	}

	return Parameters{
		URL:           URL,
		Timeout:       timeout,
		AllowRedirect: allowRedirect,
		Headers:       headers,
	}
}
