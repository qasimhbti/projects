package database

import (
	"encoding/json"
	"errors"
	"os"
)

// Client -
type Client struct {
	dbPath string
}

type databaseSchema struct {
	Users map[string]User
	Posts map[string]Post
}

// NewClient -
func NewClient(dbPath string) Client {
	return Client{
		dbPath: dbPath,
	}
}

// EnsureDB creates the database file if it doesn't exist
func (c Client) EnsureDB() error {
	_, err := os.ReadFile(c.dbPath)
	if errors.Is(err, os.ErrNotExist) {
		return c.createDB()
	}
	return err
}

func (c Client) createDB() error {
	dat, err := json.Marshal(databaseSchema{
		Users: make(map[string]User),
		Posts: make(map[string]Post),
	})
	if err != nil {
		return err
	}
	err = os.WriteFile(c.dbPath, dat, 0600)
	if err != nil {
		return err
	}
	return nil
}

func (c Client) updateDB(db databaseSchema) error {
	dat, err := json.Marshal(db)
	if err != nil {
		return err
	}
	err = os.WriteFile(c.dbPath, dat, 0600)
	if err != nil {
		return err
	}
	return nil
}

func (c Client) readDB() (databaseSchema, error) {
	dat, err := os.ReadFile(c.dbPath)
	if err != nil {
		return databaseSchema{}, err
	}
	db := databaseSchema{}
	err = json.Unmarshal(dat, &db)
	return db, err
}