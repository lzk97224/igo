package iurl

import "net/url"

// GetFullUrl 拼接url
func GetFullUrl(baseUrl string, path ...string) string {
	fullUrl, err := url.JoinPath(baseUrl, path...)
	if err != nil {
		return ""
	}
	return fullUrl
}
