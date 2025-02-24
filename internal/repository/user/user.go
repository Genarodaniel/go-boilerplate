package user

import (
	"context"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Save(ctx context.Context, user UserDTO) error {
	var query = `
		INSERT INTO tab_user(
			name,
			email,
			client_id,
			client_secret
		)
		VALUES(
			$1,
			$2,
			$3,
			$4
		)
	`
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		user.Name,
		user.Email,
		user.ClientID,
		user.ClientSecret,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, userID string) (UserDTO, error) {
	var user UserDTO
	var query = `
		SELECT
			id,
			name,
			email,
			client_id,
			client_secret
		FROM tab_user
		WHERE id = $1
	`
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return user, err
	}

	defer stmt.Close()

	result := stmt.QueryRowContext(ctx, userID)
	err = result.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.ClientID,
		&user.ClientSecret,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) GetByClientID(ctx context.Context, clientID string) (UserDTO, error) {
	var user UserDTO
	var query = `
		SELECT
			id,
			name,
			email,
			client_id,
			client_secret
		FROM tab_user
		WHERE client_id = $1
	`
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return user, err
	}

	defer stmt.Close()

	result := stmt.QueryRowContext(ctx, clientID)
	err = result.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.ClientID,
		&user.ClientSecret,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}
