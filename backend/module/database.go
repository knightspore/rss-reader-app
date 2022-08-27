package module

import "github.com/knightspore/rss-reader-app/backend/vo"

func (s *Server) UserGet(id string) (vo.User, error) {
	var user vo.User

	col := s.DB.Collection("users")
	result, err := col.Get(id, nil)
	if err != nil {
		return user, err
	}

	err = result.Content(&user)
	if err != nil {
		return user, err
	}

	return user, err
}

func (s *Server) UserUpdate(user vo.User) error {
	col := s.DB.Collection("users")

	_, err := col.Upsert(user.ID, user, nil)
	if err != nil {
		return err
	}

	return nil
}
