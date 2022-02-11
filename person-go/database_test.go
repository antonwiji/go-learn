package main

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestGetData(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	_, err := db.ExecContext(ctx, "INSERT INTO person(id, nama) VALUES (2,'Regita');")

	if err != nil {
		panic(err)
	}

	fmt.Println("data Berhasil dikirim")
}

func TestShowData(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	resultsql := "SELECT * FROM person"
	rows, err := db.QueryContext(ctx, resultsql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var nama string

		err = rows.Scan(&id, &nama)
		if err != nil {
			panic(err)
		}
		fmt.Println("id :", id)
		fmt.Println("nama :", nama)
	}
}

func TestShowPerson(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	id := 4
	rows, err := db.QueryContext(ctx, "SELECT id,nama FROM person WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	db.Close()

	for rows.Next() {
		var id int
		var nama string
		err := rows.Scan(&id, &nama)
		if err != nil {
			fmt.Println("data tidak di temukan")
			panic(err)
		}
		fmt.Println("id", id)
		fmt.Println("nama", nama)
	}

}
