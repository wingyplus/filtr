package filtr

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testMethodOK(t *testing.T, recorder *httptest.ResponseRecorder) {
	if recorder.Code != http.StatusOK {
		t.Errorf("expeect HTTP 200 status code but was %d", recorder.Code)
	}
	if recorder.Body.String() != "Hello GET Method" {
		t.Errorf("expect \"Hello GET Method\" but was \"%s\"", recorder.Body.String())
	}
}

func Test_GET_Method_OK(t *testing.T) {
	var handler http.Handler = GET(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello GET Method")
	})

	req, _ := http.NewRequest("GET", "/hello", nil)
	res := httptest.NewRecorder()

	handler.ServeHTTP(res, req)

	testMethodOK(t, res)
}

func testMethodNotAllowed(t *testing.T, recorder *httptest.ResponseRecorder) {
	if recorder.Code != http.StatusMethodNotAllowed {
		t.Errorf("expeect HTTP 405 status code but was %d", recorder.Code)
	}

	if recorder.Body.String() != "Method Not Allowed\n" {
		t.Errorf("expeect \"Method Not Allowed\" but was \"%s\"", recorder.Body.String())
	}
}

func Test_GET_Method_Not_Allow(t *testing.T) {
	var (
		methods = []string{"POST", "PUT", "DELETE", "HEAD"}
		handler http.Handler = GET(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello GET Method")
		})
	)

	for _, m := range methods {
		req, _ := http.NewRequest(m, "/hello", nil)
		res := httptest.NewRecorder()

		handler.ServeHTTP(res, req)

		testMethodNotAllowed(t, res)
	}
}

func Test_POST_Method_OK(t *testing.T) {
	var handler http.Handler = POST(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello GET Method")
	})

	req, _ := http.NewRequest("POST", "/hello", nil)
	res := httptest.NewRecorder()

	handler.ServeHTTP(res, req)

	testMethodOK(t, res)
}

