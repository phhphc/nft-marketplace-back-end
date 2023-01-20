package clients

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgreClient struct {
	Database *sql.DB
}

func NewPostgreClient(uri string) (*PostgreClient, error) {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgreClient{
		Database: db,
	}, nil
}

func (c *PostgreClient) Disconnect() error {
	return c.Database.Close()
}
