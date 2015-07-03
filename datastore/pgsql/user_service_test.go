package pgsql

import (
	"testing"

	"github.com/fortytw2/eden/model"
)

func TestUserService(t *testing.T) {
	ds, err := NewDatastore()
	if err != nil {
		t.Fatalf("NewDatastore failed, %s", err)
	}

	u, err := model.CreateUser("luke", "luke@jedi.org", "iminlovewithmysister")
	if err != nil {
		t.Errorf("model.CreateUser returned error %s", err)
	}

	err = ds.CreateUser(u)
	if err != nil {
		t.Errorf("create user returned error %s", err)
	}

	u, err = ds.GetUserByUsername("luke")
	if err != nil {
		t.Errorf("get user returned error %s", err)
	}

	u.Username = "dark luke"
	err = ds.UpdateUser(u)
	if err != nil {
		t.Errorf("update user returned error %s", err)
	}

	u, err = ds.GetUserByUsername("dark luke")
	if err != nil {
		t.Errorf("get user %s returned error %s", u.Username, err)
	}

	u, err = ds.GetUserByID(u.ID)
	if err != nil {
		t.Errorf("get user by ID %s returned error %s", u.Username, err)
	}

	return
}
