package types

import (
	"time"

	"github.com/fortytw2/eden/model"
)

// PGBoard functions to wrap the mods array up so it fits into PG
type PGBoard struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	Creator   string      `json:"creator" db:"creator_name"`
	Mods      StringArray `json:"mods" db:"mod_names"`
	Summary   string      `json:"summary"`
	Deleted   bool        `json:"deleted"`
	Approved  bool        `json:"-"`
	CreatedAt time.Time   `json:"created_at" db:"created_at"`
}

// NewPGBoard returns a PGBoard from a *model.Board
func NewPGBoard(b *model.Board) *PGBoard {
	return &PGBoard{
		ID:        b.ID,
		Name:      b.Name,
		Creator:   b.Creator,
		Mods:      StringArray(b.Mods),
		Summary:   b.Summary,
		Deleted:   b.Deleted,
		Approved:  b.Approved,
		CreatedAt: b.CreatedAt,
	}
}

// NewModelBoard converts a pGBoard into a *model.Board
func NewModelBoard(dbb *PGBoard) *model.Board {
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
