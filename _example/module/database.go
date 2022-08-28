package module

import (
	"github.com/couchbase/gocb/v2"
	"github.com/knightspore/rss-reader-app/backend/vo"
)

// New Example

func (s *Server) Get(c string, id string) (*gocb.GetResult, error) {
	col := s.DB.Collection(c)
	result, err := col.Get(id, nil)
	if err != nil {
		return nil,  err
	}
	return result, nil
}

func(s *Server) Upsert(c string, id string, j []byte) error {
		col := s.DB.Collection(c)
		_, err := col.Upsert(id, j, nil)
		if err != nil {
			return err
		}
		return nil
}


// Old

func (s *Server) UserGet(id string) (vo.User, error) {

	var user vo.User

	result, err := s.Get("users", id)
	if err != nil {
		return user, err
	}

	err = result.Content(&user)
	if err != nil {
		return user, err
	}

	return user, nil

}

func(s *Server) UserUpsert(user vo.User) error {

		col := s.DB.Collection("users")
		_, err := col.Upsert(user.ID, user, nil)
		if err != nil {
			return err
		}

		return nil

}


func (s *Server) SubscriptionsGet(ids []string) ([]vo.Subscription, error) {

	var subs []vo.Subscription

	for _, id := range ids {
		var sub vo.Subscription
		col := s.DB.Collection("subscriptions")
		result, err := col.Get(id, nil)
		if err != nil {
			return subs, err
		}
		err = result.Content(&sub)
		if err != nil {
			return subs, err
		}
		subs = append(subs, sub)
	}

	return subs, nil

}

func(s *Server) SubscriptionCreate(url string, title string) error {

		sub, err := vo.NewSubscription(url, title)
		if err != nil {
			return err
		}

		col := s.DB.Collection("subscriptions")
		_, err = col.Upsert(sub.ID, sub, nil)
		if err != nil {
			return err
		}

		return nil

}
