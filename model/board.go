package model

import "time"

// Board is a source of related posts
type Board struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Creator   string    `json:"creator"`
	Mods      []string  `json:"mods"`
	Summary   string    `json:"summary"`
	Deleted   bool      `json:"deleted"`
	Approved  bool      `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}
