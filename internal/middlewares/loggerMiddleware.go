package middlewares

import (
	"log"
	"net/http"
)

func LoggerMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		// логируем запрос
		handler.ServeHTTP(w, r)
	})
}
