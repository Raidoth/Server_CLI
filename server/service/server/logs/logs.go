package logs

import (
	"log"
	"net/http"
)

func LoggingHTTP(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println(r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}
