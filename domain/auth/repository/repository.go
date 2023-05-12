package repository

import (
	"context"
	"database/sql"

	"skegsTech/auth-service-go/domain/auth/entity"
)

type userRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Register(ctx context.Context, user *entity.User) (res *entity.User, err error) {
	
	sql := `INSERT INTO users (unique_id, name, email, password)
			VALUES ($1, $2, $3, $4)
			RETURNING id, unique_id, name, email, password, created_at, updated_at, deleted_at`

	row := r.db.QueryRow(sql, user.UniqueId, user.Name, user.Email, user.Password)

	if err := row.Scan(
		&user.Id,
		&user.UniqueId,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt); err != nil {
		return nil, err
	}

	return user , nil
}
