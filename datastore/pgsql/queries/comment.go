package queries

func addCommentQueries() {
	queries["insert_comment"] = `
    INSERT INTO comments (post_id, comment_id, op_id, body)
           VALUES (:post_id, :comment_id, :op_id, :body)`

	queries["get_comment_by_id"] = `
	SELECT * FROM comments WHERE id = $1;`

	queries["get_post_comments"] = `
    WITH RECURSIVE cte (id, body, op_id, path, comment_id, depth)  AS (
    SELECT id,
        body,
        op_id,
        array[id] AS path,
        comment_id,
        1 AS depth
    FROM comments
    WHERE comment_id = 0 AND post_id = $1

    UNION ALL

    SELECT comments.id,
        comments.body,
        comments.op_id,
        cte.path || comments.id,
        comments.comment_id,
        cte.depth + 1 AS depth
    FROM comments
    JOIN cte ON comments.comment_id = cte.id
    )
    SELECT id, body, op_id, comment_id, path, depth FROM cte
    ORDER BY path;`

	queries["get_user_comments"] = `
    WITH RECURSIVE cte (id, body, op_id, path, comment_id, depth)  AS (
    SELECT id,
        body,
        op_id,
        array[id] AS path,
        comment_id,
        1 AS depth
    FROM comments
    WHERE comment_id = 0 AND op_id = $1

    UNION ALL

    SELECT comments.id,
        comments.body,
        comments.op_id,
        cte.path || comments.id,
        comments.comment_id,
        cte.depth + 1 AS depth
    FROM comments
    JOIN cte ON comments.comment_id = cte.id
    )
    SELECT id, body, op_id, comment_id, path, depth FROM cte
    ORDER BY path;`
}
