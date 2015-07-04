package queries

func addCommentQueries() {
	queries["insert_comment"] = `
    INSERT INTO comments (id, post_id, comment_id, op_id, op_name, op_admin, body)
           VALUES (:id, :post_id, :comment_id, :op_id, :op_name, :op_admin, :body)`

	queries["get_comment_tree"] = `
    WITH RECURSIVE cte (id, post_id, comment_id, op_id, op_name, op_admin, body, depth) AS (
    SELECT id,
        post_id,
        comment_id,
        array[id] AS path,
        op_id,
        op_name,
        op_admin,
        body,
        1 AS depth
    FROM comments
    WHERE comment_id IS NULL

    UNION ALL

    SELECT comments.id,
        comment.post_id,
        comments.body,
        comments.op_id,
        comments.op_name,
        comments.op_admin,
        cte.path || comments.id,
        comments.comment_id,
        cte.depth + 1 AS depth
    FROM comments
        JOIN cte ON comments.comment_id = cte.id
    )
    SELECT id, post_id, comment_id, op_id, op_name, op_admin, body, path, depth FROM cte
    ORDER BY path;`
}

// id          SERIAL PRIMARY KEY,
// post        INT references posts (id),
// comment     INT references comments (id),
//
// op          INT references users (id),
// op_name     TEXT,
// op_admin    BOOLEAN,
// body        TEXT,
