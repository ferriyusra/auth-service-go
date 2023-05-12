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

func MobileAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("authorization")

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
			return []byte(os.Getenv("AUTH_JWT_SECRET")), nil
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

		userID, name, email, err := getUserFromJwtToken(claims)
		if err != nil {
			util.Error(w, http.StatusUnauthorized, nil, "Invalid token: "+err.Error())
			return
		}

		if userID == 0 {
			util.Error(w, http.StatusUnauthorized, nil, "Invalid token: no user ID")
			return
		}

		if name == "" {
			util.Error(w, http.StatusUnauthorized, nil, "Invalid token: no user name")
			return
		}

		if email == "" {
			util.Error(w, http.StatusUnauthorized, nil, "Invalid token: no user email")
			return
		}

		user := &entity.AuthenticatedUser{
			ID:   userID,
			Name: name,
			Email: email,
		}

		ctx := context.WithValue(r.Context(), "user", user)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func getUserFromJwtToken(js map[string]interface{}) (int64, string, string, error) {

	var userID int64
	name := ""
	email := ""

	data, ok := js["data"].(map[string]interface{})
	if !ok {
		return 0, "", "", errors.New("Invalid Token")
	}

	if val, ok := data["userId"]; ok {
		// interface{} (for JSON numbers) will be converted to float64
		userID = int64(val.(float64))

		// return userID, nil
	}

	if val, ok := data["Name"]; ok {
		name = val.(string)
	}

	return userID, name, email, nil
}
