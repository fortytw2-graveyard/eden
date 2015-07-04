package pgsql

import (
	"testing"

	"github.com/fortytw2/eden/model"
)

func TestCommentService(t *testing.T) {
	ds, err := NewDatastore()
	if err != nil {
		t.Fatalf("NewDatastore failed, %s", err)
	}

	board := &model.Board{
		Name:    "AskMeNothing",
		Creator: "fortytw2",
		Mods:    []string{"mod 1", "mod 2"},
	}

	// if this errors due to Unique constraint we're OK
	_ = ds.CreateBoard(board)

	b, err := ds.GetBoardByName("AskMeNothing")
	if err != nil {
		t.Errorf("something went wrong getting a board for comment service, %s", err)
	}

	u, err := model.CreateUser("yoko", "yoko@jedi.org", "iminlovewiththecoco")
	if err != nil {
		t.Errorf("model.CreateUser returned error %s", err)
	}

	err = ds.CreateUser(u)
	if err != nil {
		t.Errorf("create user returned error %s", err)
	}

	u, err = ds.GetUserByUsername("yoko")
	if err != nil {
		t.Errorf("get user returned error %s", err)
	}

	err = ds.CreatePost(&model.Post{
		Board: b.ID,
		OpID:  u.ID,
		Title: "whoa dude",
		Link:  "this is crazy",
		Body:  "oh my god",
	})
	if err != nil {
		t.Errorf("error creating post, %s", err)
	}

	ps, err := ds.GetUserPosts(u.ID, 0)
	if err != nil {
		t.Errorf("getting a users posts doesn't work... , %s", err)
	}

	comments := []*model.Comment{
		model.NewComment(ps[0].ID, 0, u.ID, "root comment #1"),
		model.NewComment(ps[0].ID, 0, u.ID, "root comment #2"),
		model.NewComment(ps[0].ID, 0, u.ID, "root comment #3"),
	}

	for _, c := range comments {
		err = ds.CreateComment(c)
		if err != nil {
			t.Errorf("error creating comment, %s", err)
		}
	}

	cs, err := ds.GetPostComments(ps[0].ID)
	if err != nil {
		t.Errorf("error getting comments on a post, %s", err)
	}

	if cs == nil {
		t.Errorf("no comments returned for post")
	}

	if len(cs) != 3 && cs != nil {
		t.Errorf("should be 3 comments for post ID - %d, instead %d", ps[0].ID, len(cs))
	}

	err = ds.CreateComment(model.NewComment(ps[0].ID, cs[0].ID, u.ID, "CHILD COMMENT #1 CHECK IT"))
	if err != nil {
		t.Errorf("error creating comment, %s", err)
	}

	cs, err = ds.GetPostComments(ps[0].ID)
	if err != nil {
		t.Errorf("error getting comments on a post, %s", err)
	}

	if cs == nil {
		t.Errorf("no comments returned for post")
	}

}
