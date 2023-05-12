package middleware

import (
    "log"
    "net/http"
    "time"
)

func RequestLogger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        next.ServeHTTP(w, r)

        log.Printf("%s\t%s\t%s\t", r.Method, r.RequestURI, time.Since(start))
    })
}