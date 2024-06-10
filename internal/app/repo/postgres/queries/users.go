package queries

const (
	QueryCreateUser = `
		INSERT INTO users (email, username, password) VALUES ($1, $2, $3) RETURNING id
		`
	QueryGetUserByEmail = `
		SELECT id, username, password FROM users WHERE email = $1
	`

	QueryGetUserByID = `
		SELECT id, email, username, password FROM users WHERE id = $1
	`
)
