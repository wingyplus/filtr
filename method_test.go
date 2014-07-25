package filtr

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GET_Method_OK(t *testing.T) {
	var handler http.Handler = GET(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello GET Method")
	})

	req, _ := http.NewRequest("GET", "/hello", nil)
	res := httptest.NewRecorder()

	handler.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("expeect HTTP 200 status code but was %d", res.Code)
	}
	if res.Body.String() != "Hello GET Method" {
		t.Errorf("expect \"Hello GET Method\" but was \"%s\"", res.Body.String())
	}
}

func Test_GET_Method_Not_Allow(t *testing.T) {
	var handler http.Handler = GET(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello GET Method")
	})

	req, _ := http.NewRequest("POST", "/hello", nil)
	res := httptest.NewRecorder()

	handler.ServeHTTP(res, req)

	if res.Code != http.StatusMethodNotAllowed {
		t.Errorf("expeect HTTP 405 status code but was %d", res.Code)
	}

	if res.Body.String() != "Method Not Allowed\n" {
		t.Errorf("expeect \"Method Not Allowed\" but was \"%s\"", res.Body.String())
	}
}
