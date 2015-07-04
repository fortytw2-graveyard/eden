package pgsql

import (
	"testing"

	"github.com/fortytw2/eden/model"
)

func TestPostService(t *testing.T) {
	ds, err := NewDatastore()
	if err != nil {
		t.Fatalf("NewDatastore failed, %s", err)
	}

	board := &model.Board{
		Name:    "AskMeAnything",
		Creator: "fortytw2",
		Mods:    []string{"mod 1", "mod 2"},
	}

	// if this errors due to Unique constraint we're OK
	_ = ds.CreateBoard(board)

	b, err := ds.GetBoardByName("AskMeAnything")
	if err != nil {
		t.Errorf("something went wrong getting a board for post service, %s", err)
	}

	u, err := model.CreateUser("yolo", "yolo@jedi.org", "iminlovewiththecoco")
	if err != nil {
		t.Errorf("model.CreateUser returned error %s", err)
	}

	err = ds.CreateUser(u)
	if err != nil {
		t.Errorf("create user returned error %s", err)
	}

	u, err = ds.GetUserByUsername("yolo")
	if err != nil {
		t.Errorf("get user returned error %s", err)
	}

	posts := []*model.Post{
		&model.Post{
			Board: b.ID,
			OpID:  u.ID,
			Title: "reddit dies",
			Link:  "ellen.pao",
			Body:  "dis crazy",
		},
		&model.Post{
			Board: b.ID,
			OpID:  u.ID,
			Title: "reddit dies",
			Link:  "ellen.pao",
			Body:  "dis crazy",
		},
		&model.Post{
			Board: b.ID,
			OpID:  u.ID,
			Title: "reddit dies",
			Link:  "ellen.pao",
			Body:  "dis crazy",
		},
		&model.Post{
			Board: b.ID,
			OpID:  u.ID,
			Title: "reddit dies",
			Link:  "ellen.pao",
			Body:  "dis crazy",
		},
	}

	for _, post := range posts {
		err = ds.CreatePost(post)
		if err != nil {
			t.Errorf("creating a post doesn't seem to work right %s", err)
		}
	}

	ps, err := ds.GetUserPosts(u.ID, 0)
	if err != nil {
		t.Errorf("getting a users posts doesn't work... , %s", err)
	}

	if len(ps) != 4 {
		t.Error("something is wrong with the number of posts User1 has...")
	}

	ps, err = ds.GetBoardPostsByID(b.ID, 0)
	if err != nil {
		t.Errorf("getting board posts by id doesn't work... , %s", err)
	}

	if len(ps) != 4 {
		t.Error("something is wrong with the number of posts the board has...")
	}

	ps, err = ds.GetBoardPostsByName(b.Name, 0)
	if err != nil {
		t.Errorf("getting a boards posts by name doesn't work... , %s", err)
	}

	if len(ps) != 4 {
		t.Error("something is wrong with the number of posts the board has...")
	}

}
