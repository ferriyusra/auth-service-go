package util

import (
	"context"
	"testing"

	"skegsTech/auth-service-go/domain/authenticated-user/entity"

	"github.com/google/uuid"
)

func TestGetAuthenticatedUserID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		userId := int64(100)
		user := &entity.AuthenticatedUser{
			ID: userId,
		}

		ctx := context.WithValue(context.TODO(), "user", user)

		res, _ := GetAuthenticatedUserID(ctx)

		if res != userId {
			t.Errorf(`got: %d, want: %d`, res, userId)
		}
	})

	t.Run("error", func(t *testing.T) {
		res, err := GetAuthenticatedUserID(context.TODO()) // no user passed

		if res != 0 {
			t.Errorf(`got: %d, want: %d`, res, 0)
		}
		if err == nil {
			t.Errorf("expecting an error")
		}
	})
}

// func TestGetAuthenticatedUserRole(t *testing.T) {
// 	t.Run("success", func(t *testing.T) {
// 		userId := int64(100)
// 		name := "user"
// 		email := "user@gmail.com"

// 		user := &entity.AuthenticatedUser{
// 			ID:   userId,
// 			Name: name,
// 			Email: email,
// 		}

// 		ctx := context.WithValue(context.TODO(), "user", user)

// 		res, _ := GetAuthenticatedUserRole(ctx)

// 		if res != name {
// 			t.Errorf(`got: %s, want: %s`, res, name)
// 		}
// 	})

// 	t.Run("error", func(t *testing.T) {
// 		res, err := GetAuthenticatedUserRole(context.TODO()) // no user passed

// 		if res != "" {
// 			t.Errorf(`got: %s, want: empty string`, res)
// 		}
// 		if err == nil {
// 			t.Errorf("expecting an error")
// 		}
// 	})
// }

func TestCheckUUIDsIsUniq(t *testing.T) {
	tests := []struct {
		list []uuid.UUID
		want bool
	}{
		{list: []uuid.UUID{
			uuid.MustParse("4bd31e5a-3b17-4575-98f1-73740b6ca51a"),
			uuid.MustParse("ae3f46ba-922f-405d-a082-eed0cf457fff"),
			uuid.MustParse("8270766d-a076-43b0-b31a-6813c7c40352"),
		}, want: true},
		{list: []uuid.UUID{
			uuid.MustParse("4bd31e5a-3b17-4575-98f1-73740b6ca51a"),
			uuid.MustParse("ae3f46ba-922f-405d-a082-eed0cf457fff"),
			uuid.MustParse("ae3f46ba-922f-405d-a082-eed0cf457fff"),
		}, want: false},
	}

	for i, tt := range tests {
		res := CheckUUIDsIsUniq(tt.list)

		if res != tt.want {
			t.Errorf(`test #%d - got: %t, want: %t`, i+1, res, tt.want)
		}
	}
}
