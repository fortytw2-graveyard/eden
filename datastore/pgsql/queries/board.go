package queries

func addBoardQueries() {
	// inserts a new board
	queries["insert_board"] = `
  INSERT INTO boards (name, creator, mods, summary)
         VALUES (:name, :creator, :mods, :summary)`

	queries["update_board"] = `
	UPDATE boards SET name = :name,
                  creator = :creator,
                  mods = :mods,
                  summary = :summary,
									deleted = :deleted,
									approved = :approved
                  WHERE id = :id`

	queries["get_board_by_name"] = `
	SELECT * FROM boards WHERE name = $1;`

	queries["get_board_by_id"] = `
	SELECT * FROM boards WHERE id = $1;`

	queries["delete_board_by_id"] = `
	DELETE FROM boards WHERE id = $1;`
}
