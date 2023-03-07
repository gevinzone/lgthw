package suite

import (
	"context"
	"github.com/stretchr/testify/assert"
	"time"
)

func (s *sqlSuite) TestCRUD() {
	t := s.T()
	db := s.db
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 或者 Exec(xxx)
	res, err := db.ExecContext(ctx, "INSERT INTO `test_model`(`id`, `first_name`, `age`, `last_name`) VALUES (1, 'Tom', 18, 'Jerry')")
	s.Require().NoError(err)
	affected, err := res.RowsAffected()
	s.Require().NoError(err)
	s.Require().Equal(int64(1), affected)

	rows, err := db.QueryContext(context.Background(),
		"SELECT `id`, `first_name`,`age`, `last_name` FROM `test_model` LIMIT ?", 1)
	s.Require().NoError(err)
	for rows.Next() {
		tm := &TestModel{}
		err = rows.Scan(&tm.Id, &tm.FirstName, &tm.Age, &tm.LastName)
		// 常见错误，缺了指针
		// err = rows.Scan(tm.Id, tm.FirstName, tm.Age, tm.LastName)
		if err != nil {
			rows.Close()
			t.Fatal(err)
		}
		assert.Equal(t, "Tom", tm.FirstName)
	}
	rows.Close()

	// 或者 Exec(xxx)
	res, err = db.ExecContext(ctx, "UPDATE `test_model` SET `first_name` = 'changed' WHERE `id` = ?", 1)
	s.Require().NoError(err)
	affected, err = res.RowsAffected()
	s.Require().NoError(err)
	if affected != 1 {
		t.Fatal(err)
	}

	row := db.QueryRowContext(context.Background(), "SELECT `id`, `first_name`,`age`, `last_name` FROM `test_model` LIMIT 1")
	if row.Err() != nil {
		t.Fatal(row.Err())
	}
	tm := &TestModel{}

	err = row.Scan(&tm.Id, &tm.FirstName, &tm.Age, &tm.LastName)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "changed", tm.FirstName)
}
