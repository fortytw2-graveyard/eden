package model

// Comment is either a root or child comment on a post
type Comment struct {
	PostID int

	children []Comment
}
