package models

import (
	fiber "github.com/gofiber/fiber/v2"
	pg "github.com/go-pg/pg/v10"
	// "fmt"
)

type Person struct {
	Id int64
	Uuid string
	Name string	`json:"name"`
	PhoneNo int	`json:"phoneNo" pg:",unique"`
	// UserId int
	// User *User `pg:"rel:has-one"`
}

func (p *Person) Insert(c *fiber.Ctx, pgdb *pg.DB) error {
	_, err := pgdb.Model(p).Insert()
	return err
}