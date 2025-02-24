package user_test

import (
	"context"
	"errors"
	"testing"

	"go-boilerplate/internal/repository/user"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Save(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := user.NewUserRepository(db)

	t.Run("successful save", func(t *testing.T) {
		mock.ExpectPrepare("INSERT INTO tab_user").
			ExpectExec().
			WithArgs("name", "email", "client_id", "client_secret").
			WillReturnResult(sqlmock.NewResult(1, 1))

		err = repo.Save(context.Background(), user.UserDTO{
			Name:         "name",
			Email:        "email",
			ClientID:     "client_id",
			ClientSecret: "client_secret",
		})
		assert.NoError(t, err)
	})

	t.Run("prepare error", func(t *testing.T) {
		mock.ExpectPrepare("INSERT INTO tab_user").
			WillReturnError(errors.New("prepare error"))

		err = repo.Save(context.Background(), user.UserDTO{})
		assert.Error(t, err)
		assert.Equal(t, "prepare error", err.Error())
	})

	t.Run("query error", func(t *testing.T) {
		mock.ExpectPrepare("INSERT INTO tab_user").
			ExpectExec().
			WithArgs("name", "email", "client_id", "client_secret").
			WillReturnError(errors.New("query error"))

		err = repo.Save(context.Background(), user.UserDTO{
			Name:         "name",
			Email:        "email",
			ClientID:     "client_id",
			ClientSecret: "client_secret",
		})
		assert.Error(t, err)
		assert.Equal(t, "query error", err.Error())
	})
}

func TestUserRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := user.NewUserRepository(db)

	t.Run("successful get by id", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "email", "client_id", "client_secret"}).
			AddRow("1", "name", "email", "client_id", "client_secret")

		mock.ExpectPrepare("SELECT id, name, email, client_id, client_secret FROM tab_user WHERE id = \\$1").
			ExpectQuery().
			WithArgs("1").
			WillReturnRows(rows)

		result, err := repo.GetByID(context.Background(), "1")
		assert.NoError(t, err)
		assert.Equal(t, "1", result.ID)
		assert.Equal(t, "name", result.Name)
		assert.Equal(t, "email", result.Email)
		assert.Equal(t, "client_id", result.ClientID)
		assert.Equal(t, "client_secret", result.ClientSecret)
	})

	t.Run("prepare error", func(t *testing.T) {
		mock.ExpectPrepare("SELECT id, name, email, client_id, client_secret FROM tab_user WHERE id = \\$1").
			WillReturnError(errors.New("prepare error"))

		_, err := repo.GetByID(context.Background(), "1")
		assert.Error(t, err)
		assert.Equal(t, "prepare error", err.Error())
	})

	t.Run("query error", func(t *testing.T) {
		mock.ExpectPrepare("SELECT id, name, email, client_id, client_secret FROM tab_user WHERE id = \\$1").
			ExpectQuery().
			WithArgs("1").
			WillReturnError(errors.New("query error"))

		_, err := repo.GetByID(context.Background(), "1")
		assert.Error(t, err)
		assert.Equal(t, "query error", err.Error())
	})

	t.Run("scan error", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "email", "client_id", "client_secret"}).
			AddRow(nil, nil, nil, nil, nil)

		mock.ExpectPrepare("SELECT id, name, email, client_id, client_secret FROM tab_user WHERE id = \\$1").
			ExpectQuery().
			WithArgs("1").
			WillReturnRows(rows)

		_, err := repo.GetByID(context.Background(), "1")
		assert.Error(t, err)
	})
}

func TestUserRepository_GetByClientID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := user.NewUserRepository(db)

	t.Run("successful get by client id", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "email", "client_id", "client_secret"}).
			AddRow("1", "name", "email", "client_id", "client_secret")

		mock.ExpectPrepare("SELECT id, name, email, client_id, client_secret FROM tab_user WHERE client_id = \\$1").
			ExpectQuery().
			WithArgs("client_id").
			WillReturnRows(rows)

		result, err := repo.GetByClientID(context.Background(), "client_id")
		assert.NoError(t, err)
		assert.Equal(t, "1", result.ID)
		assert.Equal(t, "name", result.Name)
		assert.Equal(t, "email", result.Email)
		assert.Equal(t, "client_id", result.ClientID)
		assert.Equal(t, "client_secret", result.ClientSecret)
	})

	t.Run("prepare error", func(t *testing.T) {
		mock.ExpectPrepare("SELECT id, name, email, client_id, client_secret FROM tab_user WHERE client_id = \\$1").
			WillReturnError(errors.New("prepare error"))

		_, err := repo.GetByClientID(context.Background(), "client_id")
		assert.Error(t, err)
		assert.Equal(t, "prepare error", err.Error())
	})

	t.Run("query error", func(t *testing.T) {
		mock.ExpectPrepare("SELECT id, name, email, client_id, client_secret FROM tab_user WHERE client_id = \\$1").
			ExpectQuery().
			WithArgs("client_id").
			WillReturnError(errors.New("query error"))

		_, err := repo.GetByClientID(context.Background(), "client_id")
		assert.Error(t, err)
		assert.Equal(t, "query error", err.Error())
	})

	t.Run("scan error", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "email", "client_id", "client_secret"}).
			AddRow(nil, nil, nil, nil, nil)

		mock.ExpectPrepare("SELECT id, name, email, client_id, client_secret FROM tab_user WHERE client_id = \\$1").
			ExpectQuery().
			WithArgs("client_id").
			WillReturnRows(rows)

		_, err := repo.GetByClientID(context.Background(), "client_id")
		assert.Error(t, err)
	})
}
