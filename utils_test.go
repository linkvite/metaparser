package parser

import (
	"reflect"
	"testing"
)

func TestToJSON(t *testing.T) {
	type Sample struct {
		Name string `json:"name"`
	}

	param1 := Sample{
		Name: "https://linkvite.io",
	}

	param2 := Sample{
		Name: "http://sample-domain.com",
	}

	want1 :=
		`{
	"name": "https://linkvite.io"
}`

	want2 :=
		`{
	"name": "http://sample-domain.com"
}`

	tests := []struct {
		param Sample
		want  string
	}{
		{
			param: param1,
			want:  want1,
		},
		{
			param: param2,
			want:  want2,
		},
	}

	for _, test := range tests {
		got := ToJSON(test.param)
		if !reflect.DeepEqual(got, test.want) {
			t.Error("Expected ", test.want, " got ", got)
		}
	}
}

func TestSetValue(t *testing.T) {
	tests := []struct {
		value    string
		fallBack string
		want     string
	}{
		{
			value:    "",
			fallBack: "https://linkvite.io",
			want:     "https://linkvite.io",
		},
		{
			value:    "https://linkvite.io",
			fallBack: "https://app.linkvite.io",
			want:     "https://linkvite.io",
		},
	}

	for _, test := range tests {
		got := setValue(test.value, test.fallBack)
		if !reflect.DeepEqual(got, test.want) {
			t.Error("Expected ", test.want, " got ", got)
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		param []string
		want  []string
	}{
		{
			param: []string{"https://linkvite.io", "https://linkvite.io"},
			want:  []string{"https://linkvite.io"},
		},
		{
			param: []string{"apple", "banana", "apple", "orange", "banana"},
			want:  []string{"apple", "banana", "orange"},
		},
	}

	for _, test := range tests {
		got := removeDuplicates(test.param)
		if !reflect.DeepEqual(got, test.want) {
			t.Error("Expected ", test.want, " got ", got)
		}
	}
}

func TestSetDefaultParameters(t *testing.T) {
	sampleHeader := map[string]string{"User-Agent": "googlebot"}
	wantHeader := map[string]string{"User-Agent": "Twitterbot/1.0"}

	param1 := Parameters{
		URL: "https://linkvite.io",
	}

	param2 := Parameters{
		URL:     "https://linkvite.io",
		Timeout: 10,
	}

	param3 := Parameters{
		URL:     "https://linkvite.io",
		Timeout: 10,
		Headers: sampleHeader,
	}

	param4 := Parameters{
		URL:           "https://linkvite.io",
		Timeout:       10,
		Headers:       wantHeader,
		AllowRedirect: true,
	}

	want1 := Parameters{
		URL:           "https://linkvite.io",
		Timeout:       10,
		Headers:       sampleHeader,
		AllowRedirect: false,
	}

	want2 := Parameters{
		URL:           "https://linkvite.io",
		Timeout:       10,
		Headers:       wantHeader,
		AllowRedirect: false,
	}

	want3 := Parameters{
		URL:           "https://linkvite.io",
		Timeout:       10,
		Headers:       wantHeader,
		AllowRedirect: true,
	}

	tests := []struct {
		param Parameters
		want  Parameters
	}{
		{
			param: param1,
			want:  want2,
		},
		{
			param: param2,
			want:  want2,
		},
		{
			param: param3,
			want:  want1,
		},
		{
			param: param4,
			want:  want3,
		},
	}

	for _, test := range tests {
		got := setDefaultParameters(test.param)
		if !reflect.DeepEqual(got, test.want) {
			t.Error("Expected ", test.want, " got ", got)
		}
	}
}
