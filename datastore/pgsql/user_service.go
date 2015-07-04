package pgsql

import (
	"github.com/fortytw2/eden/datastore/pgsql/queries"
	"github.com/fortytw2/eden/model"
	"github.com/jmoiron/sqlx"
)

// UserService is a User Service backed by Postgres
type UserService struct {
	db *sqlx.DB
}

// NewUserService returns a new User Service from the given database handle
func NewUserService(db *sqlx.DB) *UserService {
	return &UserService{
		db: db,
	}
}

// CreateUser adds a user to the datastore
func (us *UserService) CreateUser(u *model.User) (err error) {
	_, err = us.db.NamedQuery(queries.Get("insert_user"), u)
	return
}

// UpdateUser updates a database user by ID
func (us *UserService) UpdateUser(u *model.User) (err error) {
	_, err = us.db.NamedQuery(queries.Get("update_user"), u)
	return
}

// GetUserByID returns a user by their id
func (us *UserService) GetUserByID(id int) (u *model.User, err error) {
	row := us.db.QueryRowx(queries.Get("get_user_by_id"), id)

	var scanUser model.User
	err = row.StructScan(&scanUser)
	if err != nil {
		return
	}
	u = &scanUser

	return
}

// GetUserByUsername returns a user by their username
func (us *UserService) GetUserByUsername(username string) (u *model.User, err error) {
	row := us.db.QueryRowx(queries.Get("get_user_by_username"), username)

	var scanUser model.User
	err = row.StructScan(&scanUser)
	if err != nil {
		return
	}
	u = &scanUser

	return
}
