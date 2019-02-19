package model

type Post struct {
	ID     int    `json:"id"`
	Body   string `json:"body"`
	UserID string `json:"user_id"`
}
