package repositories

import (
	"errors"

	"github.com/dsgomes/rest-api/internal/core/domain"
	"github.com/dsgomes/rest-api/internal/core/ports"
	"github.com/dsgomes/rest-api/internal/infra/db"
	_ "github.com/lib/pq"
)

type todoPostgresRepository struct {
	database db.Database
}

func NewTodoPostgresRepository(db db.Database) ports.TodoRepository {
	return &todoPostgresRepository{
		database: db,
	}
}

func (t *todoPostgresRepository) GetAll() (todos []domain.Todo, err error) {
	conn, err := t.database.OpenConnection()
	if err != nil {
		return
	}
	conn.Find(&todos)
	return
}

func (t *todoPostgresRepository) Get(id string) (todo domain.Todo, err error) {
	conn, err := t.database.OpenConnection()
	if err != nil {
		return
	}
	conn.First(&todo, id)
	return
}

func (t *todoPostgresRepository) Insert(todo *domain.Todo) (id string, err error) {
	conn, err := t.database.OpenConnection()
	if err != nil {
		return "", err
	}
	conn.Create(&todo)
	return
}

func (t *todoPostgresRepository) Update(id string, todo *domain.Todo) (int64, error) {
	conn, err := t.database.OpenConnection()
	if err != nil {
		return 0, err
	}

	conn.First(&todo, id)
	rows := conn.Save(&todo).RowsAffected
	if rows == 0 {
		return 0, errors.New("0 rows affected")
	}

	return rows, nil
}

func (t *todoPostgresRepository) Delete(id string) (int64, error) {
	var todo domain.Todo
	conn, err := t.database.OpenConnection()
	if err != nil {
		return 0, err
	}
	conn.First(&todo, id)
	rows := conn.Delete(&todo, id).RowsAffected
	if rows == 0 {
		return 0, errors.New("0 rows effected")
	}

	return rows, nil
}
