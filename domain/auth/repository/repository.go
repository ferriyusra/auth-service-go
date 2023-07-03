package repository

import (
	"context"
	"database/sql"

	"skegsTech/auth-service-go/domain/auth/entity"
	"skegsTech/auth-service-go/myerror"
)

type userRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) (res *entity.User, err error) {
	
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

func (r *userRepository) Get(ctx context.Context, email string) (*entity.User, error) {
	
	sql := `SELECT u.id, u.unique_id, u.name, u.email, u.password, u.created_at, u.updated_at FROM users u WHERE u.email = $1`

	row := r.db.QueryRow(sql, email)
	
	user := entity.User{}
	if err := row.Scan(
		&user.Id,
		&user.UniqueId,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt); err != nil {
			return nil, myerror.ErrRecordNotFound
	}

	return &user , nil
}

func (r *userRepository) GetById(ctx context.Context, id int64) (*entity.User, error) {
	
	sql := `SELECT u.id, u.unique_id, u.name, u.email, u.password, u.created_at, u.updated_at FROM users u WHERE u.id = $1 AND u.deleted_at IS NULL`

	row := r.db.QueryRow(sql, id)
	
	user := entity.User{}
	if err := row.Scan(
		&user.Id,
		&user.UniqueId,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt); err != nil {
			return nil, myerror.ErrRecordNotFound
	}

	return &user , nil
}
