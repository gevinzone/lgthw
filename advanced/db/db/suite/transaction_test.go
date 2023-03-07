package suite

import (
	"context"
	"database/sql"

	"time"
)

func (s *sqlSuite) TestTransaction() {
	db := s.db
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	s.Require().NoError(err)
	res, err := tx.ExecContext(ctx, "INSERT INTO `test_model`(`id`, `first_name`, `age`, `last_name`) VALUES (1, 'Tom', 18, 'Jerry')")
	s.Require().NoError(err)
	affected, err := res.RowsAffected()
	s.Require().NoError(err)
	s.Require().Equal(int64(1), affected)
	err = tx.Commit()
	s.Require().NoError(err)
}
