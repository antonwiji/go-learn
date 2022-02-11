package repository

import (
	"context"
	"person-go/entity"
)

type PersonRepository interface {
	Insert(ctx context.Context, person entity.Person) (entity.Person, error)
	FindById(ctx context.Context, id int32) (entity.Person, error)
	FindAll(ctx context.Context) ([]entity.Person, error)
}
