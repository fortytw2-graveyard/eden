package model

// Comment is either a root or child comment on a post
type Comment struct {
	ID     int
	PostID int
	Votes  int `json:"votes" db:"-"`

	children []Comment
}
