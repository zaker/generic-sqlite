package todo

import "fmt"

type Todo struct {
	Id     int64
	Title  string
	Body   string
	Status int
}

func (t Todo) SetID(id int64) any {
	t.Id = id
	return t
}

func (t Todo) InsertArgs() []any {
	return []any{t.Title, t.Body, t.Status}
}

func (t Todo) InsertQuery() string {
	return "INSERT INTO todo(title, body, status) values(?,?,?)"
}

func (t Todo) ListQuery() string {
	return "SELECT title,body,status FROM todo"
}

func (t Todo) TableSchema() string {
	return `
	CREATE TABLE IF NOT EXISTS todo (
	  title TEXT,
	  body TEXT,
	  status INT
	)`
}

func (t Todo) ValueRefs() []any {
	return []any{&t.Title, &t.Body, &t.Status}
}

func (t Todo) Scan(dbScan func(...any) error) (any, error) {
	var newTodo Todo
	err := dbScan(&newTodo.Title, &newTodo.Body, &newTodo.Status)
	return newTodo, err
}

func (t Todo) String() string {
	return fmt.Sprintf("Title:%s, Body:%s, Status:%d", t.Title, t.Body, t.Status)
}
