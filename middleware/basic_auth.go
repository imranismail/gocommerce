package middleware

import "net/http"

func BasicAuth(username string, password string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
			usernameInput, passwordInput, ok := req.BasicAuth()

			if ok && username == usernameInput && password == passwordInput {
				next.ServeHTTP(writer, req)
			} else {
				writer.Header().Set("www-authenticate", "Basic realm=\"imranismail\"")
				writer.WriteHeader(http.StatusUnauthorized)
				writer.Write([]byte("Unauthorized"))
			}
		})
	}
}
