-- name: GetUserByID :one
SELECT user_id, full_name, email FROM users WHERE user_id = $1;


-- name: CreateUser :one
INSERT INTO users (full_name, email, hashed_password, role)
VALUES ($1, $2, $3, $4)
RETURNING user_id, full_name, email,role, created_at;
