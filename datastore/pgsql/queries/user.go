package queries

func addUserQueries() {
	// inserts a new user
	queries["insert_user"] = `
  INSERT INTO users (username, email, passwordhash, confirmed)
         VALUES (:username, :email, :passwordhash, :confirmed)`

	// update_user updates the user given by id
	queries["update_user"] = `
  UPDATE users SET username = :username,
                  passwordhash = :passwordhash,
                  email = :email,
                  confirmed = :confirmed
                  WHERE id = :id`

	// get a user by id
	queries["get_user_by_id"] = `
  SELECT * FROM users WHERE id = $1;`

	// get a user by username
	queries["get_user_by_username"] = `
  SELECT * FROM users WHERE username = $1;`
}
