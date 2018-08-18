package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var EmailSignupPOSTTests = []struct {
	email    string // input
	wantCode int    // expected result
}{
	{"bobby@bobby.com", http.StatusOK},
	{"chickenwing@hello.com", http.StatusOK},
	{"", http.StatusBadRequest},
	{"test_string", http.StatusBadRequest},
}

func TestEmailSignupPOST(t *testing.T) {
	// Request body.
	for _, tt := range EmailSignupPOSTTests {
		str := "email=" + tt.email
		reader := strings.NewReader(str)

		// Create a request to pass to our handler.
		req, err := http.NewRequest("POST", "/email-signup", reader)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(EmailSignupPOST)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(resp, req)

		// Check the status code is what we expect.
		if status := resp.Code; status != tt.wantCode {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}
}
