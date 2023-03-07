package suite

import "context"

func (s *sqlSuite) TestPrepareStatement() {
	stmt, err := s.db.Prepare("SELECT * FROM `test_model` WHERE `id` = ?")
	s.Require().NoError(err)
	_, err = stmt.QueryContext(context.Background(), 1)
	s.Require().NoError(err)
	_, err = stmt.QueryContext(context.Background(), 2)
	s.Require().NoError(err)

	err = stmt.Close()
	s.Require().NoError(err)
}
