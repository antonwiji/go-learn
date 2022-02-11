package main

import (
	"context"
	"fmt"
	"person-go/entity"
	repo "person-go/repository"
	"testing"
)

func TestPersonsql(t *testing.T) {
	commitRepository := repo.NewPersonRepository(GetConnection())
	ctx := context.Background()

	person := entity.Person{
		Nama: "Aku Ganteng nian",
	}

	result, err := commitRepository.Insert(ctx, person)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}

func TestPersonsqlFindId(t *testing.T) {
	commitRepository := repo.NewPersonRepository(GetConnection())
	ctx := context.Background()

	person, err := commitRepository.FindById(ctx, 3)
	if err != nil {
		panic(err)
	}

	fmt.Println(person)

}

func TestPersonsqlFindAll(t *testing.T) {
	commitRepository := repo.NewPersonRepository(GetConnection())
	ctx := context.Background()

	person, err := commitRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, person := range person {
		fmt.Println(person)
	}
}
