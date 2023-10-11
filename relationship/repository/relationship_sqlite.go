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

func (r sqliteRelationshipRepository) Add(members relationship.Members) error {
	query := `INSERT INTO relationship (ancestor, descendant) VALUES (?, ?)`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing add relationship query: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(members.Parent, members.Child)
	if err != nil {
		return fmt.Errorf("error exec add relationship query: %w", err)
	}

	return nil
}

func (r sqliteRelationshipRepository) FetchByID(ID int64) ([]relationship.NamedMembers, error) {
	query := `
        WITH RECURSIVE genealogy AS (
             SELECT r.* FROM relationship r WHERE descendant = ?
             UNION ALL
             SELECT r.* FROM relationship r
             JOIN genealogy g ON r.descendant = g.ancestor
        )
        SELECT anc.name AS ancestor, desc.name AS descendant
        FROM genealogy g
        JOIN person anc ON anc.rowid = g.ancestor
        JOIN person desc ON desc.rowid = g.descendant;
    `

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error preparing fetch relationships by id query: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(ID)
	if err != nil {
		return nil, fmt.Errorf("error exec fetch relationships by id query: %w", err)
	}
	defer rows.Close()

	members := make([]relationship.NamedMembers, 0)

	for rows.Next() {
		member := relationship.NamedMembers{}

		err = rows.Scan(&member.Parent, &member.Child)
		if err != nil {
			return nil, fmt.Errorf("error scanning fetch relationships by id query result: %w", err)
		}

		members = append(members, member)

	}

	return members, nil
}
