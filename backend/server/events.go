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
	ID     string `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	URL    string `json:"url,omitempty"`
	UserID string `json:"userId,omitempty"`
}

func NewSubscriptionEvents(body io.ReadCloser) ([]SubscriptionEvent, error) {
	var e []SubscriptionEvent
	err := json.NewDecoder(body).Decode(&e)
	if err != nil {
		return e, err
	}
	return e, nil
}

type ArticleEvent struct {
	ID       string `json:"id,omitempty"`
	URL      string `json:"url,omitempty"`
	UserID   string `json:"userId,omitempty"`
	ParentID string `json:"parent,omitempty"`
}

func NewArticleEvent(body io.ReadCloser) (ArticleEvent, error) {
	var e ArticleEvent
	err := json.NewDecoder(body).Decode(&e)
	if err != nil {
		return e, err
	}
	return e, nil
}
