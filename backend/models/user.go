package models

import (
	_ "github.com/lib/pq"
)

type User struct {
	ID   int    `db:"id"`   // PostgreSQLの "id" (primary key)
	Age  int    `db:"age"`  // PostgreSQLの "age" (integer)
	Name string `db:"name"` // PostgreSQLの "name" (varchar(500))
	Role string `db:"role"` // PostgreSQLの "role" (char(15))
}

type UserData struct {
	User struct {
		Age  int    `json:"age"`
		Name string `json:"name"`
		Role string `json:"role"`
	} `json:"user"`
}

func (u *User) CreateUser() error {
	cmd := `INSERT INTO users(age, name, role) VALUES ($1, $2, $3)`
	_, err := Db.Exec(cmd,
		u.Age,
		u.Name,
		u.Role,
	)

	if err != nil {
		return err
	}
	return nil

}
