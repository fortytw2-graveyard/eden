package queries

func addPostQueries() {
	queries["insert_post"] = `
    INSERT INTO posts (board_id, op_id, link, title, body)
            VALUES(:board_id, :op_id, :link, :title, :body)`

	queries["get_post_by_id"] = `
    SELECT * FROM posts WHERE id = $1;`

	queries["get_board_posts_by_id"] = `
    SELECT * FROM posts WHERE board_id = $1
    ORDER BY created_at DESC
    LIMIT 50 OFFSET $2;`

	queries["get_board_posts_by_name"] = `
    SELECT p.id
    , p.board_id
    , p.op_id
    , p.link
    , p.title
    , p.body
    , p.created_at
    FROM posts p
    LEFT JOIN boards b
        ON (b.name = $1)
    WHERE board_id = b.id
    ORDER BY p.created_at DESC
    LIMIT 50 OFFSET $2;`

	queries["get_user_posts"] = `
    SELECT * FROM posts WHERE op_id = $1 LIMIT 50 OFFSET $2;`
}
