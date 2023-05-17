package util

import (
	"context"
	"errors"

	"skegsTech/auth-service-go/domain/authenticated-user/entity"

	"github.com/google/uuid"
)

func GetAuthenticatedUserID(ctx context.Context) (int64, error) {
	user, ok := ctx.Value("user").(*entity.AuthenticatedUser)
	if !ok {
		return 0, errors.New("Invalid authenticated user")
	}

	return user.Id, nil
}

// func GetAuthenticatedUserRole(ctx context.Context) (string, error) {
// 	user, ok := ctx.Value("user").(*entity.AuthenticatedUser)
// 	if !ok {
// 		return "", errors.New("Invalid authenticated user")
// 	}

// 	return user.Role, nil
// }

func CheckUUIDsIsUniq(ids []uuid.UUID) bool {
	m := make(map[string]bool)

	for _, id := range ids {
		idStr := id.String()
		if _, exist := m[idStr]; !exist {
			m[idStr] = true
		} else {
			return false
		}
	}

	return true
}
