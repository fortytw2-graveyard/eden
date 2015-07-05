package pgsql

import (
	"log"

	"github.com/fortytw2/eden/datastore/pgsql/queries"
	"github.com/fortytw2/eden/datastore/pgsql/types"
	"github.com/fortytw2/eden/model"
	"github.com/jmoiron/sqlx"
)

// CommentService is a CommentService backed by Postgres
type CommentService struct {
	db *sqlx.DB
}

// NewCommentService returns a new CommentService from the db handle
func NewCommentService(db *sqlx.DB) *CommentService {
	return &CommentService{
		db: db,
	}
}

// CreateComment creates a comment on another comment
func (cs *CommentService) CreateComment(c *model.Comment) (err error) {
	_, err = cs.db.NamedQuery(queries.Get("insert_comment"), c)
	return
}

// GetComment returns a comment along with its full tree
func (cs *CommentService) GetComment(commentID int) (c *model.Comment, err error) {
	row := cs.db.QueryRowx(queries.Get("get_comment_by_id"), commentID)

	var scanC model.Comment
	err = row.StructScan(&scanC)
	if err != nil {
		return
	}
	c = &scanC

	return
}

type pgComment struct {
	ID        int
	Body      string
	OpID      int `db:"op_id"`
	CommentID int `db:"comment_id"`
	Path      types.IntArray
	Depth     int
}

// GetPostComments returns all root level comments along with their children for a Post`
func (cs *CommentService) GetPostComments(postID int) (comments []*model.Comment, err error) {
	var rows *sqlx.Rows
	rows, err = cs.db.Queryx(queries.Get("get_post_comments"), postID)
	if err != nil {
		return
	}

	var pgcr []types.PGComment
	for rows.Next() {
		var pgc types.PGComment
		err = rows.StructScan(&pgc)
		if err != nil {
			return
		}
		pgcr = append(pgcr, pgc)
	}

	comments = types.AssembleCommentTree(pgcr)

	return
}

// GetUserComments returns all of a users past comments
func (cs *CommentService) GetUserComments(userID int) (comments []*model.Comment, err error) {
	var rows *sqlx.Rows
	rows, err = cs.db.Queryx(queries.Get("get_user_comments"), userID)
	if err != nil {
		return
	}

	var pgcr []types.PGComment
	for rows.Next() {
		var pgc types.PGComment
		err = rows.StructScan(&pgc)
		if err != nil {
			return
		}
		log.Printf("%+v", pgc)
		pgcr = append(pgcr, pgc)
	}

	comments = types.AssembleCommentTree(pgcr)

	return
}

// DeleteComment removes a comment
func (cs *CommentService) DeleteComment(commentID int) (err error) {
	return
}
