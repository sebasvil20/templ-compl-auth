package app

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func initializeDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "file::memory:")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE TABLE users (id TEXT PRIMARY KEY, email TEXT, password TEXT, admin INTEGER(1))")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("INSERT INTO users (id, email, password, admin) VALUES ('9f8wzRjPME', 'sebas@gmail.com', '$2a$10$XCTAOEjZ6XA0kk9Ju5QaeOS412IBfOuiB6LPmCEhTuwAsvj/Dk5tO', 1)")
	if err != nil {
		return nil, err
	}

	return db, nil
}
