package repository

import (
	"database/sql"
	"fmt"

	"github.com/eu-ovictor/gnlg/relationship"
)

type sqliteRelationshipRepository struct {
	DB *sql.DB
}

func NewSQLiteRelationshipRepository(db *sql.DB) sqliteRelationshipRepository {
	return sqliteRelationshipRepository{
		DB: db,
	}
}

func (r sqliteRelationshipRepository) Add(rel relationship.Relationship) error {
	query := `INSERT INTO relationship (ancestor, descendant) VALUES (?, ?)`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing add relationship query: %w", err)
	}

	_, err = stmt.Exec(rel.Parent, rel.Child)
	if err != nil {
		return fmt.Errorf("error exec add relationship query: %w", err)
	}

	return nil
}
