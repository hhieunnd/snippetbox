package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Define type snippet hold the data.
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Wrap sql.DB connection pool.
type SnippetModel struct {
	DB *pgxpool.Pool
}

// Insert new snippet into database.
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	query := `INSERT INTO snippets (title, content, created, expires) VALUES ($1, $2, CURRENT_TIMESTAMP AT TIME ZONE 'UTC', $3) RETURNING id`

	var id int
	err := m.DB.QueryRow(context.Background(), query, title, content, expires).Scan(&id)

	if err != nil {
		return 0, err
	}

	return 0, nil
}

// Return a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// Return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
