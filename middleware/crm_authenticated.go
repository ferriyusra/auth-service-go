package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"skegsTech/auth-service-go/domain/authenticated-user/entity"
	"skegsTech/auth-service-go/util"

	"github.com/golang-jwt/jwt"
)

func CrmAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("authorization")

		if authorizationHeader == "" {
			authorizationHeader = r.URL.Query().Get("authorization")
		}

		if authorizationHeader == "" {
			util.Error(w, http.StatusUnauthorized, nil, "An authorization header is required")
			return
		}

		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) != 2 {
			util.Error(w, http.StatusUnauthorized, nil, "An authorization header is required")
			return
		}

		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte(os.Getenv("AUTH_JWT_SECRET_CRM")), nil
		})
		if err != nil {
			util.Error(w, http.StatusUnauthorized, nil, "Invalid token: "+err.Error())
			return
		}

		if !token.Valid {
			util.Error(w, http.StatusUnauthorized, nil, "Invalid token")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			util.Error(w, http.StatusUnauthorized, nil, "Invalid token: "+err.Error())
			return
		}

		userID, err := getUserIdFromJwt(claims)
		if err != nil {
			util.Error(w, http.StatusUnauthorized, nil, "Invalid token: "+err.Error())
			return
		}

		if userID == 0 {
			util.Error(w, http.StatusUnauthorized, nil, "Invalid token: no user ID")
			return
		}

		user := &entity.AuthenticatedUser{
			ID: userID,
		}

		ctx := context.WithValue(r.Context(), "user", user)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func getUserIdFromJwt(js map[string]interface{}) (int64, error) {
	data, ok := js["data"].(map[string]interface{})
	if !ok {
		return 0, errors.New("Invalid CRM Token")
	}

	if val, ok := data["userId"]; ok {
		// interface{} (for JSON numbers) will be converted to float64
		userID := int64(val.(float64))

		return userID, nil
	}

	return 0, errors.New("Invalid CRM Token")
}
