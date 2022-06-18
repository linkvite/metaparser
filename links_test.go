package parser

import "testing"

func TestValidateLink(t *testing.T) {
	got, _ := ValidateLink("https://app.linkvite.io")
	want := "linkvite.io"

	if want != got {
		t.Error("Expected ", want, " got ", got)
	}
}

func TestGetBaseURL(t *testing.T) {
	got := getBaseURL("https://linkvite.io/kayode")
	want := "https://linkvite.io"

	if want != got {
		t.Error("Expected ", want, " got ", got)
	}
}

func TestResolveURL(t *testing.T) {
	tests := []struct {
		link string
		data string
		want string
	}{
		{
			link: "https://linkvite.io",
			data: "//linkvite.io",
			want: "linkvite.io",
		},
		{
			link: "https://linkvite.io",
			data: "/kayode",
			want: "https://linkvite.io/kayode",
		},
		{
			link: "https://linkvite.io",
			data: "",
			want: "",
		},
		{
			link: "https://linkvite.io/kayode",
			data: "https://linkvite.io/kayode",
			want: "https://linkvite.io/kayode",
		},
	}

	for _, test := range tests {
		got := resolveURL(test.link, test.data)
		if test.want != got {
			t.Error("Expected ", test.want, " got ", got)
		}
	}
}
