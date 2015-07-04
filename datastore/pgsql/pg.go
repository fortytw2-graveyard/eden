package pgsql

import (
	"fmt"
	"os"

	"github.com/fortytw2/eden/datastore"
	"github.com/jmoiron/sqlx"
	// import the postgres driver
	_ "github.com/lib/pq"
)

// NewDatastore creates a new datastore backed by PGSQL
func NewDatastore() (ds *datastore.Datastore, err error) {
	var db *sqlx.DB
	db, err = NewDBHandle()
	if err != nil {
		return
	}

	ds = &datastore.Datastore{
		UserService:    NewUserService(db),
		BoardService:   NewBoardService(db),
		PostService:    NewPostService(db),
		CommentService: NewCommentService(db),
	}

	return
}

// NewDBHandle returns the raw *sqlx.DB -> can be used to construct multi-backend
// datastores
func NewDBHandle() (db *sqlx.DB, err error) {
	db, err = sqlx.Connect("postgres", fmt.Sprintf("user=%s %s dbname=%s %s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_EXTRA")))
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}
	return
}
