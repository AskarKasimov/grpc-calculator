package web

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type iDatabase interface {
	Login(login, password string) (int64, error)
	Register(login, password string) (int64, error)
}

type database struct {
	db *sql.DB
}

var db iDatabase

func DB() iDatabase { return db }

func init() {
	connStr := "user=admin password=admin dbname=yg sslmode=disable host=postgres"
	newConn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	db = &database{db: newConn}
}

func (d *database) Login(login, password string) (int64, error) {
	var id int64

	err := d.db.QueryRow("SELECT id FROM users WHERE login=$1 AND passwordHash=$2", login, password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (d *database) Register(login, password string) (int64, error) {
	var id int64

	err := d.db.QueryRow("INSERT INTO users (login, passwordHash) VALUES ($1, $2) RETURNING id", login, password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
