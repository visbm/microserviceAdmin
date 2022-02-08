package session

import (
	"database/sql"
	"microseviceAdmin/webapp"
	"os"

	"github.com/antonlindstrom/pgstore"
)

type SessionStore struct {
	DB      *sql.DB
	PGStore *pgstore.PGStore
}
var sstore SessionStore

func OpenSessionStore(c *webapp.Config) error {

	dataSourceName := c.PgDataSource()
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	sstore.DB = db

	PGStore, err := pgstore.NewPGStoreFromPool(db, []byte(os.Getenv("SESSION_KEY")))
	if err != nil {
		return err
	}
	sstore.PGStore = PGStore

	return nil

}
