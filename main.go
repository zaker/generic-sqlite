package main

import (
	"database/sql"
	"fmt"

	"github.com/zaker/generic-sqlite/repo"
	"github.com/zaker/generic-sqlite/todo"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "my.db")
	if err != nil {
		panic(err)
	}

	r := repo.Repo[todo.Todo]{DB: db}
	err = r.CreateTable(todo.Todo{})
	if err != nil {
		panic(err)
	}

	t, err := r.Insert(todo.Todo{Title: "title 1", Body: "body", Status: 1234})
	if err != nil {
		panic(err)
	}
	fmt.Println("inserted todo", t, t.Id)

	todos, err := r.List(todo.Todo{})
	if err != nil {
		panic(err)
	}
	for _, t := range todos {
		fmt.Println(t)
	}
}
