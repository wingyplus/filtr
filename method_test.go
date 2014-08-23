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
	if recorder.Body.String() != "Hello Method" {
		t.Errorf("expect \"Hello Method\" but was \"%s\"", recorder.Body.String())
	}
}

type testCase struct {
	method string
	h      http.Handler
}

func fn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Method")
}

func TestMethodOK(t *testing.T) {
	var testCases = []testCase{
		testCase{"GET", GET(fn)},
		testCase{"POST", POST(fn)},
	}

	for _, tc := range testCases {
		req, _ := http.NewRequest(tc.method, "/hello", nil)
		recorder := httptest.NewRecorder()

		tc.h.ServeHTTP(recorder, req)

		testMethodOK(t, recorder)
	}
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
		methods              = []string{"POST", "PUT", "DELETE", "HEAD"}
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

func Test_POST_Method_Not_Allow(t *testing.T) {
	var (
		methods              = []string{"GET", "PUT", "DELETE", "HEAD"}
		handler http.Handler = POST(func(w http.ResponseWriter, r *http.Request) {
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
