package api

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := initSchema(db); err != nil {
		return nil, err
	}

	return db, nil
}

func initSchema(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		biography TEXT NOT NULL
	);
	`

	_, err := db.Exec(schema)
	return err
}

func CreateUserDB(db *sql.DB, user User) error {
	_, err := db.Exec(
		"INSERT INTO users (id, first_name, last_name, biography) VALUES ($1, $2, $3, $4)",
		user.ID, user.FirstName, user.LastName, user.Biography,
	)
	return err
}

func GetUserDB(db *sql.DB, id string) (User, error) {
	user := User{}
	err := db.QueryRow(
		"SELECT id, first_name, last_name, biography FROM users WHERE id = $1",
		id,
	).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Biography)

	if err == sql.ErrNoRows {
		return User{}, fmt.Errorf("user not found")
	}

	return user, err
}

func ListUsersDB(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, first_name, last_name, biography FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Biography); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, rows.Err()
}

func UpdateUserDB(db *sql.DB, user User) error {
	result, err := db.Exec(
		"UPDATE users SET first_name = $1, last_name = $2, biography = $3 WHERE id = $4",
		user.FirstName, user.LastName, user.Biography, user.ID,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

func DeleteUserDB(db *sql.DB, id string) (User, error) {
	user, err := GetUserDB(db, id)
	if err != nil {
		return User{}, err
	}

	_, err = db.Exec("DELETE FROM users WHERE id = $1", id)

	return user, err
}
