package queries

func addBoardQueries() {
	// inserts a new board
	queries["insert_board"] = `
    INSERT INTO boards (name, creator_name, mod_names, summary)
         VALUES (:name, :creator_name, :mod_names, :summary)`

	queries["update_board"] = `
    UPDATE boards SET name = :name,
                  creator_name = :creator_name,
                  mod_names = :mod_names,
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

	queries["get_all_boards"] = `
    SELECT * FROM boards LIMIT 50 OFFSET $1;`
}
