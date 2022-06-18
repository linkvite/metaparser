package parser

import (
	"reflect"
	"testing"
)

func TestParseLink(t *testing.T) {
	want1 := MetaData{
		Title:       "Linkvite",
		Description: "The most convenient way to save and organize your bookmarks, collaborate with others, and much more.",
		Domain:      "linkvite.io",
		URL:         "https://linkvite.io",
		SiteName:    "Linkvite",
		SiteType:    "website",
		Images:      []string{"https://assets.linkvite.io/covers/cover.png"},
		Favicons:    []string{"https://linkvite.io/favicon.ico"},
	}

	want2 := MetaData{
		Title:       "sample-domain.com",
		Description: "",
		Domain:      "sample-domain.com",
		URL:         "http://sample-domain.com",
		SiteName:    "sample-domain.com",
		SiteType:    "website",
		Images:      []string{},
		Favicons:    []string{},
	}

	tests := []struct {
		URL  string
		want MetaData
	}{
		{
			URL:  "https://linkvite.io",
			want: want1,
		},
		{
			URL:  "http://sample-domain.com",
			want: want2,
		},
	}

	for _, test := range tests {
		got, _ := ParseLink(Parameters{
			URL: test.URL,
		})

		if !reflect.DeepEqual(got, test.want) {
			t.Error("Expected ", test.want, " got ", got)
		}
	}
}
