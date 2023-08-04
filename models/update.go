package models

import "github.com/dsgomes/rest-api/db"

func Update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	sql := "UPDATE todos SET title=$2, description=$3, done=$4 WHERE id=$1"

	res, err := conn.Exec(sql, todo.ID, todo.Title, todo.Description, todo.Done)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
