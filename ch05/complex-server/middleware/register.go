package middleware

import (
	"net/http"

	"github.com/YoungsoonLee/practical-go/ch05/complex-server/config"
)

// RegisterMiddleware ...
func RegisterMiddleware(mux *http.ServeMux, c config.AppConfig) http.Handler {
	return loggingMiddleware(panicMiddleware(mux, c), c)
}
