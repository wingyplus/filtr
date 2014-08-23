package filtr

import "net/http"

func allowedMethod(method string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		} else {
			h.ServeHTTP(w, r)
		}
	})
}

func GET(f func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return allowedMethod("GET", http.HandlerFunc(f))
}

func POST(f func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return allowedMethod("POST", http.HandlerFunc(f))
}

func PUT(f func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return allowedMethod("PUT", http.HandlerFunc(f))
}
