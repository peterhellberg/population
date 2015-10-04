package population

import (
	"net/http"
	"net/http/httptest"
)

func testServerAndClientWithJSON(code int, body []byte) (*httptest.Server, *Client) {
	return testServerAndClient(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(body)
	})
}

func testServerAndClient(fn http.HandlerFunc) (*httptest.Server, *Client) {
	var (
		ts = httptest.NewServer(http.HandlerFunc(fn))
		c  = NewClient()
	)

	c.BaseURL = ts.URL

	return ts, c
}
