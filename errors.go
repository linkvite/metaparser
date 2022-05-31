package parser

import (
	"errors"
	"fmt"
)

var errorMessage = errors.New("❌ Failed to parse the url. ")

// Make the error function available to the package.
func PrintError(why string) error {
	return fmt.Errorf("%w Reason: %s", errorMessage, why)
}

// Handle creating a new metadata object if there's an error.
func returnResultWithError(why string, metaData *MetaData, linkData *Link) MetaData {
	logError(why)

	// Generate a blank data.
	metaData.setDefaultValue(linkData, linkData.Domain)

	return *metaData
}

// Log error.
func logError(word string) {
	log(fmt.Sprintf("❌ Failed to get the url. Reason: %s", word))
	log("🤧 Generating blank data.")
}
