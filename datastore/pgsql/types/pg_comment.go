package types

import "github.com/fortytw2/eden/model"

// PGComment handles serialization and deserialization of comments from PG
// using sneaky Disqus style SQL recursive WITH
type PGComment struct {
	ID        int
	Body      string
	OpID      int `db:"op_id"`
	CommentID int `db:"comment_id"`
	Path      IntArray
	Depth     int
}

// GetModelComment returns a *model.Comment from a PGComment
func (p *PGComment) GetModelComment() *model.Comment {
	return &model.Comment{
		ID:        p.ID,
		Body:      p.Body,
		UserID:    p.OpID,
		CommentID: p.CommentID,
		Children:  make([]*model.Comment, 0),
	}
}

// AssembleCommentTree pulls together a proper tree-structured array of comments
// from the query result PGComment structs
func AssembleCommentTree(rows []PGComment) (comments []*model.Comment) {
	for _, row := range rows {
		if row.Depth == 1 {
			comments = append(comments, row.GetModelComment())
		}
		if row.Depth == 2 {
			id := getParentNode(row.Path[0]-1, comments)
			comments[id].Children = append(comments[id].Children, row.GetModelComment())
		}
		if row.Depth == 3 {
			id1 := getParentNode(row.Path[0]-1, comments)
			id2 := getParentNode(row.Path[1]-1, comments)
			comments[id1].Children[id2].Children = append(comments[id1].Children[id2].Children, row.GetModelComment())
		}
		if row.Depth == 4 {
			id1 := getParentNode(row.Path[0]-1, comments)
			id2 := getParentNode(row.Path[1]-1, comments)
			id3 := getParentNode(row.Path[2]-1, comments)
			comments[id1].Children[id2].Children[id3].Children = append(comments[id1].Children[id2].Children[id3].Children, row.GetModelComment())
		}
		if row.Depth == 5 {
			id1 := getParentNode(row.Path[0]-1, comments)
			id2 := getParentNode(row.Path[1]-1, comments)
			id3 := getParentNode(row.Path[2]-1, comments)
			id4 := getParentNode(row.Path[3]-1, comments)
			comments[id1].Children[id2].Children[id3].Children[id4].Children = append(comments[id1].Children[id2].Children[id3].Children[id4].Children, row.GetModelComment())
		}
	}
	return
}

func getParentNode(id int, cs []*model.Comment) int {
	for i, node := range cs {
		if node.ID == id {
			return i
		}
	}
	return 0

}
