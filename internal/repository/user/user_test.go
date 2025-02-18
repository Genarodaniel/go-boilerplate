package user

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestOrderRepository_Save(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewOrderRepository(db)

	user := UserDTO{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	mock.ExpectPrepare("INSERT INTO tab_user").
		ExpectQuery().
		WithArgs(user.Name, user.Email).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))

	id, err := repo.Save(context.Background(), user)
	assert.NoError(t, err)
	assert.Equal(t, "1", id)
}

func TestOrderRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewOrderRepository(db)

	userID := "1"
	expectedUser := UserDTO{
		ID:    userID,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	mock.ExpectPrepare("SELECT id, name, email FROM tab_user WHERE id = \\$1").
		ExpectQuery().
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(expectedUser.ID, expectedUser.Name, expectedUser.Email))

	user, err := repo.GetByID(context.Background(), userID)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}
