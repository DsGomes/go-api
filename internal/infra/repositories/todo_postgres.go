package repositories

import (
	"time"

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
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM todos")
	if err != nil {
		return
	}

	for rows.Next() {
		var todo domain.Todo

		err = rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.Done,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return
}

func (t *todoPostgresRepository) Get(id string) (todo domain.Todo, err error) {
	conn, err := t.database.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow("SELECT * FROM todos WHERE id=$1", id)

	err = row.Scan(
		&todo.ID,
		&todo.Title,
		&todo.Description,
		&todo.Done,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)

	return
}

func (t *todoPostgresRepository) Insert(todo *domain.Todo) (id int64, err error) {
	conn, err := t.database.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}

func (t *todoPostgresRepository) Update(id string, todo *domain.Todo) (int64, error) {
	conn, err := t.database.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	sql := "UPDATE todos SET title=$2, description=$3, done=$4, updated_at=$5 WHERE id=$1"

	res, err := conn.Exec(sql, id, todo.Title, todo.Description, todo.Done, time.Now())
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (t *todoPostgresRepository) Delete(id string) (int64, error) {
	conn, err := t.database.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec("DELETE FROM todos WHERE id=$1", id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
