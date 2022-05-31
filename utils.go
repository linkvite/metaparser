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
