package models

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
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
	query := `
	    INSERT INTO snippets (title, content, created, expires)
	    VALUES ($1, $2, CURRENT_TIMESTAMP AT TIME ZONE 'UTC', CURRENT_TIMESTAMP AT TIME ZONE 'UTC' + $3 * INTERVAL '1 day')
	    RETURNING id`

	var id int
	err := m.DB.QueryRow(context.Background(), query, title, content, strconv.Itoa(expires)).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// Return a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	query := `SELECT id, title, content, created, expires FROM snippets WHERE id = $1`

	snippet := &Snippet{}
	err := m.DB.QueryRow(context.Background(), query, id).Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return snippet, nil
}

// Return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	query := `
		SELECT id, title, content, created, expires
		FROM snippets
		ORDER BY created DESC
		LIMIT 10`

	rows, err := m.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snippets []*Snippet
	for rows.Next() {
		s := &Snippet{}
		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return snippets, nil
}
