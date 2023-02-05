package mariadb

const (
	findUserByIDQuery = `
		SELECT
			id,
			email,
			hashed_password
		FROM
			users
		WHERE
			id = ?
	`
)
