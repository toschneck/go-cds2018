package router

import (
	"testing"
	"net/http/httptest"
	"net/http"
)

func TestNewBLRouter(t *testing.T) {
	r := NewBLRouter()
	srv := httptest.NewServer(r)
	defer srv.Close()

	testCases := []struct {
		route          string
		expectedStatus int
	}{
		{"/home", http.StatusOK},
		{"/", http.StatusNotFound},
	}

	for _, c := range testCases {
		resp, err := http.Get(srv.URL + c.route)

		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != c.expectedStatus {
			t.Errorf("Status code is %v, but %v expected", resp.StatusCode, c.expectedStatus)
		}
	}
}
