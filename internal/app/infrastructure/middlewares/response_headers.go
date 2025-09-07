/*
The Middlewares handles the middlewares in the request - response pipeline
*/
package middlewares

import "net/http"

var responseHeaders map[string]string = map[string]string{
	"Content-Type": "application/json",
}

// AddHeaderMiddleware add the response header to response.
func AddHeaderMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Add the custom header to the response
			for k, v := range responseHeaders {
				w.Header().Set(k, v)
			}
			// Call the next handler in the chain
			next.ServeHTTP(w, r)
		})
	}
}
