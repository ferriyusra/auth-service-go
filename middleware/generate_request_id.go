package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

func GenerateRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uuid := uuid.New()
		r = r.WithContext(context.WithValue(r.Context(), "requestID", uuid))
		next.ServeHTTP(w, r)
	})
}
