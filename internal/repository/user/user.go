package user

import (
	"context"
	"database/sql"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		DB: db,
	}
}

func (r *OrderRepository) Save(ctx context.Context, user User) (string, error) {
	var id string
	var query = `
		INSERT INTO tab_user(
			name,
			email
		)
		VALUES(
			$1,
			$2
		)
		RETURNING id
	`
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return id, err
	}

	defer stmt.Close()

	result, err := stmt.QueryContext(ctx,
		user.Name,
		user.Email,
	)
	if err != nil {
		return id, err
	}

	result.Next()
	err = result.Scan(
		&id,
	)

	return id, nil
}
