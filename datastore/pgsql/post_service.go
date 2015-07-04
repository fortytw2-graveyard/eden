package pgsql

import (
	"github.com/fortytw2/eden/datastore/pgsql/queries"
	"github.com/fortytw2/eden/model"
	"github.com/jmoiron/sqlx"
)

// PostService is a PostService backed by Postgres
type PostService struct {
	db *sqlx.DB
}

// NewPostService returns a new PostService from the db handle
func NewPostService(db *sqlx.DB) *PostService {
	return &PostService{
		db: db,
	}
}

// CreatePost adds a new post to the database
func (ps *PostService) CreatePost(p *model.Post) (err error) {
	_, err = ps.db.NamedQuery(queries.Get("insert_post"), p)
	return
}

// GetPost returns a single post, without comments, by ID
func (ps *PostService) GetPost(postID int) (post *model.Post, err error) {
	row := ps.db.QueryRowx(queries.Get("get_post_by_id"), postID)

	var scanPost model.Post
	err = row.StructScan(&scanPost)
	if err != nil {
		return
	}
	post = &scanPost

	return
}

// GetBoardPostsByName returns the filtered posts for a board by BoardName
func (ps *PostService) GetBoardPostsByName(boardName string, page int) (posts []*model.Post, err error) {
	var rows *sqlx.Rows
	rows, err = ps.db.Queryx(queries.Get("get_board_posts_by_name"), boardName, page*50)
	if err != nil {
		return
	}

	for rows.Next() {
		var post model.Post
		err = rows.StructScan(&post)
		if err != nil {
			return
		}

		posts = append(posts, &post)
	}
	return
}

// GetBoardPostsByID returns the filtered posts for a board by the ID of the board
func (ps *PostService) GetBoardPostsByID(boardID, page int) (posts []*model.Post, err error) {
	var rows *sqlx.Rows
	rows, err = ps.db.Queryx(queries.Get("get_board_posts_by_id"), boardID, page*50)
	if err != nil {
		return
	}

	for rows.Next() {
		var post model.Post
		err = rows.StructScan(&post)
		if err != nil {
			return
		}

		posts = append(posts, &post)
	}
	return
}

// GetUserPosts returns all posts owned by a certain user
func (ps *PostService) GetUserPosts(userID, page int) (posts []*model.Post, err error) {
	var rows *sqlx.Rows
	rows, err = ps.db.Queryx(queries.Get("get_user_posts"), userID, page*50)
	if err != nil {
		return
	}

	for rows.Next() {
		var post model.Post
		err = rows.StructScan(&post)
		if err != nil {
			return
		}

		posts = append(posts, &post)
	}
	return
}

// DeletePost marks a post as deleted - deleted posts are still searchable though
func (ps *PostService) DeletePost(postID int) (err error) {
	return
}
