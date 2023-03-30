-- Code generated by sqlc-crud. DO NOT EDIT.

-- name: InsertUser :exec
INSERT INTO user (
	 guid, name
) VALUES (
	 ?, ?
);
-- name: GetUser :one
SELECT * FROM user WHERE id = ?;


-- name: GetUserByGUID :one
SELECT * FROM user WHERE guid = ?;


-- name: DeleteUser :exec
DELETE FROM user WHERE id = ?;