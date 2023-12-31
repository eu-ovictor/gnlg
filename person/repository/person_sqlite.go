package repository

import (
	"database/sql"
	"fmt"

	"github.com/eu-ovictor/gnlg/person"
	_ "github.com/mattn/go-sqlite3"
)

type sqlitePersonRepository struct {
	DB *sql.DB
}

func NewSQLitePersonRepository(db *sql.DB) sqlitePersonRepository {
	return sqlitePersonRepository{
		DB: db,
	}
}

func (r sqlitePersonRepository) Add(p person.Person) error {
	query := `INSERT INTO person (name) VALUES (?)`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing add person query: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Name)
	if err != nil {
		return fmt.Errorf("error exec add person query: %w", err)
	}

	return nil
}

func (r sqlitePersonRepository) Edit(p person.Person) error {
	query := `UPDATE person SET name = ? WHERE rowid = ?`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing update person query: %w", err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(p.Name, p.ID); err != nil {
		return fmt.Errorf("error exec update person query: %w", err)
	}

	return nil
}

func (r sqlitePersonRepository) Fetch(ID int, name string) ([]person.Person, error) {
	var args []interface{}

	query := `SELECT rowid, name FROM person WHERE 1=1 `

	if ID != 0 {
		query += `AND rowid = ?`
		args = append(args, ID)
	}

	if name != "" {
		query += `AND name LIKE ?`
		args = append(args, name+"%")
	}

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error preparing fetch people query: %w", err)

	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, fmt.Errorf("error exec fetch people query: %w", err)
	}

	defer rows.Close()

	people := make([]person.Person, 0)

	for rows.Next() {
		person := person.Person{}

		err = rows.Scan(&person.ID, &person.Name)
		if err != nil {
			return nil, fmt.Errorf("error scanning fetch people query result: %w", err)
		}

		people = append(people, person)
	}

	return people, nil
}

func (r sqlitePersonRepository) Delete(ID int) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return fmt.Errorf("error begin transaction to delete person: %w", err)
	}

	deleteRelsQuery := `DELETE FROM relationship WHERE ancestor = ? OR descendant = ?`

	deleteRelsStmt, err := tx.Prepare(deleteRelsQuery)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error preparing delete person relationships query: %w", err)
	}
	defer deleteRelsStmt.Close()

	if _, err := deleteRelsStmt.Exec(ID, ID); err != nil {
		tx.Rollback()
		return fmt.Errorf("error exec delete person relationships query: %w", err)
	}

	deletePersonQuery := `DELETE FROM person WHERE rowid = ?`

	deletePersonStmt, err := tx.Prepare(deletePersonQuery)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error preparing delete person query: %w", err)
	}
	defer deleteRelsStmt.Close()

	if _, err := deletePersonStmt.Exec(ID); err != nil {
		tx.Rollback()
		return fmt.Errorf("error exec delete person query: %w", err)
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return fmt.Errorf("error commiting delete person queries: %w", err)
	}

	return nil
}
