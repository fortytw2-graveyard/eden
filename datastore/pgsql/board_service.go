package pgsql

import (
	"github.com/fortytw2/eden/datastore/pgsql/queries"
	"github.com/fortytw2/eden/datastore/pgsql/types"
	"github.com/fortytw2/eden/model"
	"github.com/jmoiron/sqlx"
)

// BoardService is a BoardService backed by Postgres
type BoardService struct {
	db *sqlx.DB
}

// NewBoardService returns a new BoardService from the db handle
func NewBoardService(db *sqlx.DB) *BoardService {
	return &BoardService{
		db: db,
	}
}

// CreateBoard adds a new board to the datastore
func (bs *BoardService) CreateBoard(b *model.Board) (err error) {
	dbb := types.NewPGBoard(b)
	_, err = bs.db.NamedQuery(queries.Get("insert_board"), dbb)
	return err
}

// UpdateBoard just updates a board
func (bs *BoardService) UpdateBoard(b *model.Board) (err error) {
	dbb := types.NewPGBoard(b)
	_, err = bs.db.NamedQuery(queries.Get("update_board"), dbb)
	return
}

// GetBoards returns paginated lists of boards
func (bs *BoardService) GetBoards(page int) (boards []*model.Board, err error) {
	var rows *sqlx.Rows
	rows, err = bs.db.Queryx(queries.Get("get_all_boards"), page*50)
	if err != nil {
		return
	}

	for rows.Next() {
		var board types.PGBoard
		err = rows.StructScan(&board)
		if err != nil {
			return
		}

		boards = append(boards, types.NewModelBoard(&board))
	}
	return
}

// GetBoardByName returns the board information
func (bs *BoardService) GetBoardByName(boardName string) (b *model.Board, err error) {
	row := bs.db.QueryRowx(queries.Get("get_board_by_name"), boardName)

	var scanBoard types.PGBoard
	err = row.StructScan(&scanBoard)
	if err != nil {
		return
	}
	b = types.NewModelBoard(&scanBoard)

	return
}

// GetBoardByID returns the board information
func (bs *BoardService) GetBoardByID(boardID int) (b *model.Board, err error) {
	row := bs.db.QueryRowx(queries.Get("get_board_by_id"), boardID)

	var scanBoard types.PGBoard
	err = row.StructScan(&scanBoard)
	if err != nil {
		return
	}
	b = types.NewModelBoard(&scanBoard)

	return
}

// DeleteBoard deletes the board - returns error if it fails
func (bs *BoardService) DeleteBoard(boardID int) (err error) {
	_, err = bs.db.Queryx(queries.Get("delete_board_by_id"), boardID)
	return
}
