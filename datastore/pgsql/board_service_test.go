package pgsql

import (
	"testing"

	"github.com/fortytw2/eden/model"
)

func TestBoardService(t *testing.T) {
	ds, err := NewDatastore()
	if err != nil {
		t.Fatalf("NewDatastore failed, %s", err)
	}

	board := &model.Board{
		Name:    "AskMeAnything",
		Creator: "fortytw2",
		Mods:    []string{"mod 1", "mod 2"},
	}

	err = ds.CreateBoard(board)
	if err != nil {
		t.Errorf("something went wrong inserting a board, %s", err)
	}

	board.Mods = []string{"mod 1"}
	err = ds.UpdateBoard(board)
	if err != nil {
		t.Errorf("something went wrong updating a board, %s", err)
	}

	b, err := ds.GetBoardByName("AskMeAnything")
	if err != nil {
		t.Errorf("something went wrong getting a board by username, %s", err)
	}
	if b.Creator != "fortytw2" {
		t.Errorf("found board was not created by fortytw2, instead %s", b.Creator)
	}

	if len(b.Mods) > 1 {
		t.Errorf("found board had more than 1 moderator, instead %s", b.Creator)
	}

	b, err = ds.GetBoardByID(b.ID)
	if err != nil {
		t.Errorf("something went wrong getting a board by id, %s", err)
	}
	if b.Creator != "fortytw2" {
		t.Errorf("found board was not created by fortytw2, instead %s", b.Creator)
	}

	board = &model.Board{
		Name:    "AskMeAnything2",
		Creator: "fortytw2",
		Mods:    []string{"mod 1", "mod 2"},
	}

	err = ds.CreateBoard(board)
	if err != nil {
		t.Errorf("something went wrong inserting a board, %s", err)
	}

	boards, err := ds.GetBoards(0)
	if err != nil {
		t.Errorf("something went wrong getting every board, %s", err)
	}

	if len(boards) != 2 {
		t.Error("apparently we don't have two boards...")
	}

	err = ds.DeleteBoard(b.ID)
	if err != nil {
		t.Errorf("something went wrong deleting a board, %s", err)
	}

	_, err = ds.GetBoardByID(b.ID)
	if err == nil {
		t.Errorf("apparently we can get boards that are deleted, err -, %s", err)
	}

}
