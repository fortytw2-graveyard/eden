package pgsql

import (
	"errors"
	"time"

	"github.com/fortytw2/eden/datastore/pgsql/queries"
	"github.com/fortytw2/eden/datastore/pgsql/types"
	"github.com/fortytw2/eden/model"
	"github.com/jmoiron/sqlx"
)

var (
	errBoardNameTaken    = errors.New("board name already taken")
	errBoardSummaryEmpty = errors.New("board summary can't be empty")

	errModeratorNotFound = errors.New("cannot remove a mod that isn't a mod")
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
	dbb := newDBBoard(b)
	_, err = bs.db.NamedQuery(queries.Get("insert_board"), dbb)
	return err
}

// UpdateBoard just updates a board
func (bs *BoardService) UpdateBoard(b *model.Board) (err error) {
	dbb := newDBBoard(b)
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
		var board dbBoard
		err = rows.StructScan(&board)
		if err != nil {
			return
		}

		boards = append(boards, newModelBoard(&board))
	}
	return
}

// GetBoardByName returns the board information
func (bs *BoardService) GetBoardByName(boardName string) (b *model.Board, err error) {
	row := bs.db.QueryRowx(queries.Get("get_board_by_name"), boardName)

	var scanBoard dbBoard
	err = row.StructScan(&scanBoard)
	if err != nil {
		return
	}
	b = newModelBoard(&scanBoard)

	return
}

// GetBoardByID returns the board information
func (bs *BoardService) GetBoardByID(boardID int) (b *model.Board, err error) {
	row := bs.db.QueryRowx(queries.Get("get_board_by_id"), boardID)

	var scanBoard dbBoard
	err = row.StructScan(&scanBoard)
	if err != nil {
		return
	}
	b = newModelBoard(&scanBoard)

	return
}

// DeleteBoard deletes the board - returns error if it fails
func (bs *BoardService) DeleteBoard(boardID int) (err error) {
	_, err = bs.db.Queryx(queries.Get("delete_board_by_id"), boardID)
	return
}

// dbBoard functions to wrap the mods array up so it fits into PG
type dbBoard struct {
	ID        int               `json:"id"`
	Name      string            `json:"name"`
	Creator   string            `json:"creator"`
	Mods      types.StringArray `json:"mods"`
	Summary   string            `json:"summary"`
	Deleted   bool              `json:"deleted"`
	Approved  bool              `json:"-"`
	CreatedAt time.Time         `json:"created_at" db:"created_at"`
}

func newDBBoard(b *model.Board) *dbBoard {
	return &dbBoard{
		ID:        b.ID,
		Name:      b.Name,
		Creator:   b.Creator,
		Mods:      types.StringArray(b.Mods),
		Summary:   b.Summary,
		Deleted:   b.Deleted,
		Approved:  b.Approved,
		CreatedAt: b.CreatedAt,
	}
}

func newModelBoard(dbb *dbBoard) *model.Board {
	return &model.Board{
		ID:        dbb.ID,
		Name:      dbb.Name,
		Creator:   dbb.Creator,
		Mods:      []string(dbb.Mods),
		Summary:   dbb.Summary,
		Deleted:   dbb.Deleted,
		Approved:  dbb.Approved,
		CreatedAt: dbb.CreatedAt,
	}
}
