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

	_, err = stmt.Exec(p.Name)
	if err != nil {
		return fmt.Errorf("error exec add person query: %w", err)
	}

	return nil
}

func (r sqlitePersonRepository) Edit(p person.Person) (int64, error) {
	query := `UPDATE person SET name = ? WHERE rowid = ?`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("error preparing update person query: %w", err)
	}

	res, err := stmt.Exec(p.Name, p.ID)
	if err != nil {
		return 0, fmt.Errorf("error exec update person query: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("error getting affected rows count: %w", err)
	}

	return rowsAffected, nil
}

func (r sqlitePersonRepository) Fetch() ([]person.Person, error) {
	query := `SELECT rowid, name FROM person`

	rows, err := r.DB.Query(query)
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
