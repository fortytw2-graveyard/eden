package model

import "time"

// A PostFilter is used to query posts
type PostFilter struct {
}

// Post is a peice of content with either a link or a body
type Post struct {
	ID    int    `json:"id"`
	OP    *User  `json:"op"`
	Votes int    `json:"votes" db:"-"`
	Title string `json:"title"`
	Link  string `json:"link,omitempty"`
	Body  string `json:"body,omitempty"`

	CreatedAt time.Time `json:"created_at"`
}
