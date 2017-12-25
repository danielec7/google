package browser

import (
	"net/url"
)

func buildUrl(query string) string {
	u := &url.URL{
		Scheme: "https",
		Host: "www.google.com",
		Path: "/search",
	}
	q := u.Query()
	q.Set("q", query)
	u.RawQuery = q.Encode()

	return u.String()
}

func Start(query string) error {
	u := buildUrl(query)

	return open(u).Start()
}
