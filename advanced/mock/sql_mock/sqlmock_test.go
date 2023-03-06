package sqlmock

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestSqlMock(t *testing.T) {
	type TestModel struct {
		Id        int64
		FirstName string
		Age       int8
		LastName  *sql.NullString
	}

	db, mock, err := sqlmock.New()
	defer db.Close()
	require.NoError(t, err)

	mockRows := sqlmock.NewRows([]string{"id", "first_name"})
	mockRows.AddRow(1, "Tom")
	mockRows.AddRow(2, "Jerry")
	mock.ExpectQuery("SELECT id, first_name from `table`.*").WillReturnRows(mockRows)

	mockRows = sqlmock.NewRows([]string{"id", "first_name", "last_name"})
	mockRows.AddRow(1, "TomCat", "Yu")
	mockRows.AddRow(2, "JerryMouse", "Yu")
	mock.ExpectQuery("SELECT id, first_name from `table`.*").WillReturnRows(mockRows)

	mock.ExpectQuery("SELECT id from `user`.*").WillReturnError(errors.New("mock error"))

	rows, err := db.QueryContext(context.Background(), "SELECT id, first_name from `table` WHERE code='user'")
	require.NoError(t, err)
	for rows.Next() {
		tm := TestModel{}
		err = rows.Scan(&tm.Id, &tm.FirstName)
		require.NoError(t, err)
		log.Println(tm)
	}

	rows, err = db.QueryContext(context.Background(), "SELECT id, first_name from `table` WHERE code='user'")
	require.NoError(t, err)
	for rows.Next() {
		tm := &TestModel{}
		err = rows.Scan(&tm.Id, &tm.FirstName, &tm.LastName)
		require.NoError(t, err)
		log.Println(tm)
	}

	_, err = db.QueryContext(context.Background(), "SELECT id from `user` WHERE id=1")
	require.Error(t, err)

}
