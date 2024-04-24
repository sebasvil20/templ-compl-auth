package app

import (
	"database/sql"

	"github.com/sebasvil20/templ-compl-auth/internal/config"
	"golang.org/x/crypto/bcrypt"
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

	saltedPass := "123" + config.PasswordSalt
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(saltedPass), bcrypt.DefaultCost)
	_, err = db.Exec("INSERT INTO users (id, email, password, admin) VALUES ('9f8wzRjPME', 'sebas@gmail.com', ?, 1)", hashedPassword)
	if err != nil {
		return nil, err
	}

	return db, nil
}
