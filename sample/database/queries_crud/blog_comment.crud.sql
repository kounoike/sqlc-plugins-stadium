-- Code generated by sqlc-crud. DO NOT EDIT.

-- name: InsertBlogComment :exec
INSERT INTO blog_comment (
	 guid, url, contents
) VALUES (
	 ?, ?, ?
);
-- name: GetBlogComment :one
SELECT * FROM blog_comment WHERE id = ?;


-- name: GetBlogCommentByGUID :one
SELECT * FROM blog_comment WHERE guid = ?;


-- name: DeleteBlogComment :exec
DELETE FROM blog_comment WHERE id = ?;
