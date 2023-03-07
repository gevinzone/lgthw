package suite

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type sqlSuite struct {
	suite.Suite

	db *sql.DB

	driver string
	dsn    string
}

func (s *sqlSuite) TearDownTest() {
	_, err := s.db.Exec("DELETE FROM test_model;")
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *sqlSuite) SetupSuite() {
	db, err := sql.Open(s.driver, s.dsn)
	s.Require().NoError(err)
	s.db = db
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err = s.db.ExecContext(ctx, `
CREATE TABLE IF NOT EXISTS test_model(
    id INTEGER PRIMARY KEY,
    first_name TEXT NOT NULL,
    age INTEGER,
    last_name TEXT NOT NULL
)
`)
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *sqlSuite) TearDownSuite() {
	_ = s.db.Close()
}

func TestSQLite(t *testing.T) {
	suite.Run(t, &sqlSuite{
		driver: "sqlite3",
		dsn:    "file:test.db?cache=shared&mode=memory",
	})
}

type TestModel struct {
	Id        int64 `eorm:"auto_increment,primary_key"`
	FirstName string
	Age       int8
	LastName  *sql.NullString
}
