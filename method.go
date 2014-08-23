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

func GET(h http.Handler) http.Handler {
	return allowedMethod("GET", h)
}

func POST(h http.Handler) http.Handler {
	return allowedMethod("POST", h)
}

func PUT(h http.Handler) http.Handler {
	return allowedMethod("PUT", h)
}

func DELETE(h http.Handler) http.Handler {
	return allowedMethod("DELETE", h)
}
