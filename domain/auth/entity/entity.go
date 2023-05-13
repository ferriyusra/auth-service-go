package entity

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type User struct {
    Id          int            `json:"id"`
    UniqueId    uuid.UUID      `json:"uniqueId"`
    Name        string         `json:"name"`
    Email       string         `json:"email"`
    Password    string         `json:"password"`
    CreatedAt   time.Time      `json:"createdAt"`
    UpdatedAt   time.Time      `json:"updatedAt"`
    DeletedAt   *time.Time     `json:"deletedAt"`
}

type Claims struct {
    Id string `json:"id"`
    jwt.StandardClaims
}

func GetAccountSearcheables() []string {
    return []string{"id", "name", "email", "createdAt", "updatedAt"}
}
