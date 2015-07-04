package queries

func addPostQueries() {
	queries["insert_post"] = `
    INSERT INTO posts (board, op, link, title, body)
            VALUES(:board, :op, :link, :title, :body)`

	queries["get_post_by_id"] = `
    SELECT * FROM posts WHERE id = $1;`

	queries["get_board_posts_by_id"] = `
    SELECT * FROM posts WHERE board = $1
    ORDER BY created_at DESC
    LIMIT 50 OFFSET $2;`

	queries["get_board_posts_by_name"] = `
    SELECT p.id
    , p.board
    , p.op
    , p.link
    , p.title
    , p.body
    , p.created_at
    FROM posts p
    LEFT JOIN boards b
        ON (b.name = $1 AND b.id = p.board)
    WHERE board = b.id
    ORDER BY p.created_at DESC
    LIMIT 50 OFFSET $2;`

	queries["get_user_posts"] = `
    SELECT * FROM posts WHERE op = $1 LIMIT 50 OFFSET $2;`
}
