package main

import (
	"context"
	"fmt"
	"latihan-go/repository"
	"testing"
)

func TestFindID(t *testing.T) {
	commit := repository.NewConnections(GetConnection())
	ctx := context.Background()

	result, err := commit.FindbyId(ctx, 2)

	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
