package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	tests := []struct {
		playQuerystring string
		statusCode      int
		contentType     string
	}{
		{playQuerystring: "/", statusCode: 500, contentType: ""},
		{playQuerystring: "/?play=0", statusCode: 200, contentType: "application/json"},
		{playQuerystring: "/?play=1", statusCode: 200, contentType: "application/json"},
		{playQuerystring: "/?play=2", statusCode: 200, contentType: "application/json"},
		{playQuerystring: "/?play=3", statusCode: 500, contentType: ""},
	}
	for _, playData := range tests {
		// Set up a new request.[
		req, err := http.NewRequest("GET", playData.playQuerystring, nil)
		if err != nil {
			t.Fatal(err)
		}
		// Our API expects a form body, so set the content-type header to make sure it's treated as one.
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()

		http.HandlerFunc(IndexHandler).ServeHTTP(rr, req)

		statusCode := rr.Result().StatusCode
		if statusCode != playData.statusCode {
			t.Errorf("Expected Status Code (%v), Got Status Code(%v)", playData.statusCode, statusCode)
		}

		contentType := rr.Result().Header.Get("Content-Type")
		if contentType != playData.contentType {
			t.Errorf("Expected Content Type (%v), Got Content Type (%v)", playData.contentType, contentType)
		}
	}
}
