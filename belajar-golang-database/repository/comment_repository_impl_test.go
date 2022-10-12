package repository

import (
	belajargolangdatabase "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())

	ctx := context.Background()

	comment := entity.Comment{
		Email:   "test@gmail.com",
		Comment: "ini adalah comment",
	}
	result, err := commentRepository.Insert(ctx, comment)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("result", result)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())

	ctx := context.Background()

	result, err := commentRepository.FindById(ctx, 20)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("result", result)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())

	ctx := context.Background()

	result, err := commentRepository.FindAll(ctx)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	for _, comment := range result {
		fmt.Println(comment)
	}
}
