package server

import (
	"log"
	"net/http"
	"time"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.RequestURI)

		defer func(startedAt time.Time) {
			log.Println(r.RequestURI, time.Since(startedAt))
		}(time.Now())

		next.ServeHTTP(w, r)
	})
}