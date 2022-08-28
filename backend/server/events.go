package server

import (
	"encoding/json"
	"io"
)

type UserEvent struct {
	ID string `json:"id"`
}

func NewUserEvent(body io.ReadCloser) (UserEvent, error) {
	var e UserEvent
	err := json.NewDecoder(body).Decode(&e)
	if err != nil {
		return e, err
	}
	return e, nil
}

type SubscriptionEvent struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	URL    string `json:"url"`
	UserID string `json:"userId"`
}

func NewSubscriptionEvent(body io.ReadCloser) (SubscriptionEvent, error) {
	var e SubscriptionEvent
	err := json.NewDecoder(body).Decode(&e)
	if err != nil {
		return e, err
	}
	return e, nil
}

type ArticleEvent struct {
	ID       string `json:"id"`
	URL      string `json:"url"`
	UserID   string `json:"userId"`
	ParentID string `json:"parent"`
}

func NewArticleEvent(body io.ReadCloser) (ArticleEvent, error) {
	var e ArticleEvent
	err := json.NewDecoder(body).Decode(&e)
	if err != nil {
		return e, err
	}
	return e, nil
}
