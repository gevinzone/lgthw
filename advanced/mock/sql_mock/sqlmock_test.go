package sqlmock

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestSqlMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	require.NoError(t, err)

	// 第一个mock
	mockRows := sqlmock.NewRows([]string{"id", "first_name"})
	mockRows.AddRow(1, "Tom")
	mockRows.AddRow(2, "Jerry")
	mock.ExpectQuery("SELECT id, first_name from `table`.*").WillReturnRows(mockRows)

	// 第二个mock
	mockRows = sqlmock.NewRows([]string{"id", "first_name", "last_name"})
	mockRows.AddRow(1, "TomCat", "Yu")
	mockRows.AddRow(2, "JerryMouse", "Yu")
	mock.ExpectQuery("SELECT id, first_name from `table`.*").WillReturnRows(mockRows)

	// 第三个mock
	mock.ExpectQuery("SELECT id from `user`.*").WillReturnError(errors.New("mock error"))

	// 对应第一个mock query
	rows, err := db.QueryContext(context.Background(), "SELECT id, first_name from `table` WHERE code='user'")
	require.NoError(t, err)
	for rows.Next() {
		tm := TestModel{}
		err = rows.Scan(&tm.Id, &tm.FirstName)
		require.NoError(t, err)
		log.Println(tm)
	}

	// 对应二个mock query
	rows, err = db.QueryContext(context.Background(), "SELECT id, first_name from `table` WHERE code='user'")
	require.NoError(t, err)
	for rows.Next() {
		tm := &TestModel{}
		err = rows.Scan(&tm.Id, &tm.FirstName, &tm.LastName)
		require.NoError(t, err)
		log.Println(tm)
	}

	// 对应第三个mock query
	_, err = db.QueryContext(context.Background(), "SELECT id from `user` WHERE id=1")
	require.Error(t, err)

}

func TestSqlMockWithCases(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = db.Close() }()
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name           string
		query          string
		mockErr        error
		mockReturnRows *sqlmock.Rows
		wantErr        error
		wantVal        TestModel
	}{
		{
			name:  "get data",
			query: "SELECT id, first_name from `table` WHERE code='user'",
			mockReturnRows: func() *sqlmock.Rows {
				mockRows := sqlmock.NewRows([]string{"id", "first_name"})
				mockRows.AddRow(1, "Tom")
				return mockRows
			}(),
			wantVal: TestModel{
				Id:        1,
				FirstName: "Tom",
				Age:       0,
				LastName:  nil,
			},
		},
		{
			name:  "get data 2",
			query: "SELECT id, first_name from `table` WHERE code='user'",
			mockReturnRows: func() *sqlmock.Rows {
				mockRows := sqlmock.NewRows([]string{"id", "first_name"})
				mockRows.AddRow(2, "Jerry")
				return mockRows
			}(),
			wantVal: TestModel{
				Id:        2,
				FirstName: "Jerry",
				Age:       0,
				LastName:  nil,
			},
		},
		{
			name:    "query error",
			query:   "SELECT id, first_name from `table` WHERE code='user'",
			mockErr: errors.New("unexpected"),
			wantVal: TestModel{
				Id:        2,
				FirstName: "Jerry",
				Age:       0,
				LastName:  nil,
			},
		},
		{
			name:  "no row",
			query: "SELECT id, first_name from `table` WHERE code='user'",
			mockReturnRows: func() *sqlmock.Rows {
				mockRows := sqlmock.NewRows([]string{"id", "first_name"})
				return mockRows
			}(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expectedQuery := mock.ExpectQuery("SELECT.*")
			if tc.mockErr != nil {
				expectedQuery.WillReturnError(tc.mockErr)
			} else {
				expectedQuery.WillReturnRows(tc.mockReturnRows)
			}

			rows, er := db.QueryContext(context.Background(), tc.query)
			assert.Equal(t, tc.mockErr, er)
			if er != nil {
				return
			}
			if !rows.Next() {
				return
			}
			tm := TestModel{}
			er = rows.Scan(&tm.Id, &tm.FirstName)
			assert.Equal(t, tc.wantErr, er)
			if er != nil {
				return
			}
			assert.Equal(t, tc.wantVal, tm)
		})
	}
}

func TestSqlInsertMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	require.NoError(t, err)

	mockResult := sqlmock.NewResult(1, 1)
	mock.ExpectExec("INSERT INTO `test_model`.*").WillReturnResult(mockResult)

	mockResult = sqlmock.NewResult(2, 1)
	mock.ExpectExec("INSERT INTO `test_model`.*").WillReturnResult(mockResult)

	// 对应第一个mock result
	res, err := db.ExecContext(context.Background(), "INSERT INTO `test_model` (`id`, `first_name`, `age`, `last_name`) VALUES (1, 'Tom', 18, 'Jerry')")
	require.NoError(t, err)
	affected, err := res.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), affected)
	lastId, err := res.LastInsertId()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), lastId)

	// 对应第二个mock result
	res, err = db.ExecContext(context.Background(), "INSERT INTO `test_model` (`id`, `first_name`, `age`, `last_name`) VALUES (1, 'Tom', 18, 'Jerry')")
	require.NoError(t, err)
	affected, err = res.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), affected)
	lastId, err = res.LastInsertId()
	assert.NoError(t, err)
	assert.Equal(t, int64(2), lastId)

}

func TestSqlTransactionCommitMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	require.NoError(t, err)

	mockRows := sqlmock.NewRows([]string{"id", "first_name"})
	mockRows.AddRow(1, "Tom")
	mockRows.AddRow(2, "Jerry")
	mockResult := sqlmock.NewResult(1, 1)

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT id, first_name from `table`.*").WillReturnRows(mockRows)
	mock.ExpectExec("INSERT INTO `test_model`.*").WillReturnResult(mockResult)
	mock.ExpectCommit()

	err = bz(db)
	require.NoError(t, err)
}

func TestSqlTransactionRollbackMock1(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	require.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT id, first_name from `table`.*").WillReturnError(errors.New("unexpected"))
	mock.ExpectRollback()

	err = bz(db)
	require.Error(t, err)
}

func TestSqlTransactionRollbackMock2(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	require.NoError(t, err)

	mockRows := sqlmock.NewRows([]string{"id", "first_name"})
	mockRows.AddRow(1, "Tom")
	mockRows.AddRow(2, "Jerry")

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT id, first_name from `table`.*").WillReturnRows(mockRows)
	mock.ExpectExec("INSERT INTO `test_model`.*").WillReturnError(errors.New("fails to insert row"))
	mock.ExpectRollback()

	err = bz(db)
	require.Error(t, err)
}

func bz(db *sql.DB) (err error) {
	tx, err := db.Begin()
	defer func() {
		commitOrRollback(tx, err)
	}()

	if err != nil {
		return err
	}

	tm, err := bz1(tx)
	if err != nil {
		return
	}

	return bz2(tx, tm)
}

func commitOrRollback(tx *sql.Tx, err error) {
	if err != nil {
		log.Println("roll back: ", tx.Rollback())
		return
	}
	log.Println("commit ", tx.Commit())
}

func bz1(tx *sql.Tx) (*TestModel, error) {
	rows, err := tx.QueryContext(context.Background(), "SELECT id, first_name from `table` WHERE code='user'")
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, errors.New("no row found")
	}
	tm := &TestModel{}
	err = rows.Scan(&tm.Id, &tm.FirstName)
	if err != nil {
		return nil, err
	}
	return tm, nil
}

func bz2(tx *sql.Tx, tm *TestModel) error {
	res, err := tx.ExecContext(context.Background(), fmt.Sprintf("INSERT INTO `test_model` (`id`, `first_name`) VALUES (%d, '%s')", tm.Id+1, tm.FirstName))
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	_, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

type TestModel struct {
	Id        int64
	FirstName string
	Age       int8
	LastName  *sql.NullString
}
