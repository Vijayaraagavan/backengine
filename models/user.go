package models

import (
	fiber "github.com/gofiber/fiber/v2"
	pg "github.com/go-pg/pg/v10"
	// "fmt"
)
type User struct {
	Id int64 `json:"id"`
	// Person_id *Person `json:"person_id" pg:"rel:belongs-to"`
	UserName string `json:"username" pg:",unique"`
	Password_hash string `json:"password_hash"`
	Person_id int64	`pg:,unique`
	Person *Person `pg:"rel:has-one"`
}

func (u *User) Insert(c *fiber.Ctx, pgdb *pg.DB) error {
	_, err := pgdb.Model(u).Insert()
	return err
}