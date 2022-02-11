package repository

import (
	"context"
	"database/sql"
	"errors"
	"latihan-go/entity"
)

type repository struct {
	DB *sql.DB
}

func NewConnections(db *sql.DB) PersonRepository {
	return &repository{DB: db}
}

func (repo *repository) Insert(ctx context.Context, person entity.Person) (entity.Person, error) {
	script := "INSERT INTO person(nama) VALUES (?)"
	result, err := repo.DB.ExecContext(ctx, script, person.Nama)
	if err != nil {
		return person, nil
	}

	id, err := result.LastInsertId()
	if err != nil {
		return person, nil
	}
	person.Id = int32(id)

	return person, nil

}

func (repo *repository) FindbyId(ctx context.Context, id int32) (entity.Person, error) {

	script := "SELECT id, nama FROM person WHERE id = ? LIMIT 1"
	rows, err := repo.DB.QueryContext(ctx, script, id)
	person := entity.Person{}

	if err != nil {
		return person, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&person.Id, &person.Nama)
		return person, nil
	} else {
		return person, errors.New("Koneksi Gagal")
	}

}

func (repo *repository) FindAll(ctx context.Context) ([]entity.Person, error) {

	script := "SELECT id, nama FROM person"

	rows, err := repo.DB.QueryContext(ctx, script)
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
