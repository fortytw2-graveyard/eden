package model

import "testing"

// TestCreateUser generates a new user and checks to see that everything is OK
func TestCreateUser(t *testing.T) {
	u, err := CreateUser("luke", "luke@jedicouncil.org", "embraceTheDarkSide")
	if err != nil {
		t.Errorf("create user returned an error, %s", err)
	}

	if u.Username != "luke" {
		t.Errorf("something went wrong, %s is not luke", u.Username)
	}

	if u.Email != "luke@jedicouncil.org" {
		t.Errorf("something went wrong, %s is not luke@jedicouncil.org", u.Email)
	}
}
