package parser

import (
	"net"
	"strings"
)

// Result represents the parsed domain parts.
type Result struct {
	Subdomain string `json:"subdomain"`
	Domain    string `json:"domain"`
	Suffix    string `json:"suffix"`
	IsPrivate bool   `json:"isPrivate"`
}

// StringLike is a constraint that permits any type with an underlying type of string.
type StringLike interface {
	~string
}

// Parse takes a domain string and returns the parsed result.
// It returns nil if the domain is invalid or an empty string.
// This is a generic function that accepts any type with an underlying type of string.
func Parse[T StringLike](domain T) *Result {
	domainStr := string(domain)
	if domainStr == "" {
		return nil
	}

	normalizedDomain := strings.ToLower(domainStr)

	// Handle IP address
	if ip := net.ParseIP(normalizedDomain); ip != nil {
		return &Result{
			Domain: normalizedDomain,
		}
	}

	parts := strings.Split(normalizedDomain, ".")

	var longestSuffix string
	// Find the longest matching suffix.
	for i := 0; i < len(parts); i++ {
		potentialSuffix := strings.Join(parts[i:], ".")
		if _, exists := publicSuffixList[potentialSuffix]; exists {
			longestSuffix = potentialSuffix
			break
		}
	}

	if longestSuffix == "" {
		// If no public suffix is found, the last part is the TLD.
		if len(parts) < 2 {
			return &Result{
				Domain: normalizedDomain,
			}
		}
		suffix := parts[len(parts)-1]
		domainName := parts[len(parts)-2]
		subdomain := strings.Join(parts[:len(parts)-2], ".")
		return &Result{
			Subdomain: subdomain,
			Domain:    domainName,
			Suffix:    suffix,
			IsPrivate: false,
		}
	}

	suffixParts := strings.Split(longestSuffix, ".")
	domainIndex := len(parts) - len(suffixParts) - 1

	if domainIndex < 0 {
		// This happens if the domain is a public suffix itself, e.g., "co.uk"
		return &Result{
			Domain: normalizedDomain,
		}
	}

	domainPart := parts[domainIndex]
	subdomainPart := strings.Join(parts[:domainIndex], ".")
	isPrivate, _ := privateSuffixList[longestSuffix]

	return &Result{
		Subdomain: subdomainPart,
		Domain:    domainPart,
		Suffix:    longestSuffix,
		IsPrivate: isPrivate,
	}
}
