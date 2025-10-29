package parser

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type Domain string

	testCases := []struct {
		name     string
		domain   string
		expected *Result
	}{
		{
			name:   "Simple Domain",
			domain: "www.google.com",
			expected: &Result{
				Subdomain: "www",
				Domain:    "google",
				Suffix:    "com",
				IsPrivate: false,
			},
		},
		{
			name:   "Multi-part Suffix",
			domain: "forums.bbc.co.uk",
			expected: &Result{
				Subdomain: "forums",
				Domain:    "bbc",
				Suffix:    "co.uk",
				IsPrivate: false,
			},
		},
		{
			name:   "Private Suffix",
			domain: "my-project.github.io",
			expected: &Result{
				Subdomain: "",
				Domain:    "my-project",
				Suffix:    "github.io",
				IsPrivate: true,
			},
		},
		{
			name:   "IP Address",
			domain: "192.168.1.1",
			expected: &Result{
				Subdomain: "",
				Domain:    "192.168.1.1",
				Suffix:    "",
				IsPrivate: false,
			},
		},
		{
			name:     "Empty String",
			domain:   "",
			expected: nil,
		},
		{
			name:   "Public Suffix itself",
			domain: "co.uk",
			expected: &Result{
				Subdomain: "",
				Domain:    "co.uk",
				Suffix:    "",
				IsPrivate: false,
			},
		},
		{
			name:   "Custom string type",
			domain: string(Domain("www.example.org")),
			expected: &Result{
				Subdomain: "www",
				Domain:    "example",
				Suffix:    "org",
				IsPrivate: false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test with standard string
			result := Parse(tc.domain)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Parse(%q) got %+v, want %+v", tc.domain, result, tc.expected)
			}

			// Test with generic type
			customDomain := Domain(tc.domain)
			genericResult := Parse(customDomain)
			if !reflect.DeepEqual(genericResult, tc.expected) {
				t.Errorf("Generic Parse(%q) got %+v, want %+v", tc.domain, genericResult, tc.expected)
			}
		})
	}
}
