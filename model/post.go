package model

import "time"

// A PostFilter is used to query posts
type PostFilter struct {
}

// Post is a peice of content with either a link or a body
type Post struct {
	ID    int    `json:"id"`
	Board int    `json:"board" db:"board_id"`
	OpID  int    `json:"-" db:"op_id"`
	Votes int    `json:"votes"`
	Title string `json:"title"`
	Link  string `json:"link,omitempty"`
	Body  string `json:"body,omitempty"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
