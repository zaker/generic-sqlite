package repo

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type Item interface {
	SetID(int64) any
	TableSchema() string
	InsertQuery() string
	InsertArgs() []any
	ListQuery() string
	Scan(func(...any) error) (any, error)
}

type Repo[T Item] struct {
	DB *sql.DB
}

func (r Repo[T]) CreateTable(item T) error {
	_, err := r.DB.Exec(T.TableSchema(item))

	return err
}

func (r Repo[T]) Insert(item T) (*T, error) {
	res, err := r.DB.Exec(item.InsertQuery(), item.InsertArgs()...)
	if err != nil {

		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	item = item.SetID(id).(T)

	return &item, nil
}

func (r Repo[T]) List(item T) ([]T, error) {

	rows, err := r.DB.Query(item.ListQuery())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []T
	for rows.Next() {
		rowItem, err := item.Scan(rows.Scan)
		if err != nil {
			return nil, err
		}
		all = append(all, rowItem.(T))

	}
	return all, nil
}
