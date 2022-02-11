package repository

import (
	"context"
	"database/sql"
	"errors"
	"person-go/entity"
)

type personRepositoryImpl struct {
	DB *sql.DB
}

func NewPersonRepository(db *sql.DB) PersonRepository {
	return &personRepositoryImpl{DB: db}
}

func (repository *personRepositoryImpl) Insert(ctx context.Context, person entity.Person) (entity.Person, error) {
	script := "INSERT INTO person(nama) VALUES(?)"
	result, err := repository.DB.ExecContext(ctx, script, person.Nama)
	if err != nil {
		return person, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return person, err
	}
	person.Id = int32(id)
	return person, nil
}

func (repository *personRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Person, error) {
	script := "SELECT id, nama FROM person WHERE id = ? LIMIT 1"

	rows, err := repository.DB.QueryContext(ctx, script, id)
	person := entity.Person{}

	if err != nil {
		return person, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&person.Id, &person.Nama)
		return person, nil
	} else {
		return person, errors.New("id tidak di temukan")
	}
}

func (repository *personRepositoryImpl) FindAll(ctx context.Context) ([]entity.Person, error) {
	script := "SELECT id, nama FROM person"

	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var persons []entity.Person

	for rows.Next() {
		person := entity.Person{}
		rows.Scan(&person.Id, &person.Nama)
		persons = append(persons, person)
	}
	return persons, nil
}
