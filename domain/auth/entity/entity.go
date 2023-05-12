package entity

import (
	"time"

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

func GetAccountSearcheables() []string {
    return []string{"id", "name", "email", "createdAt", "updatedAt"}
}
