package controllers

import (
	"fmt"
	// "github.com/valyala/fasthttp"
	fiber "github.com/gofiber/fiber/v2"
	// pg "github.com/go-pg/pg/v10"
	uuid "github.com/pborman/uuid"
	"errors"
	// "sync"

	"backengine/models"
	"backengine/util"
	// "backengine/errs"
	"backengine/db"
)

type register struct {
	Person models.Person	`json:"person"`
	User models.User		`json:"user"`
}

func NewFriend(c *fiber.Ctx) error {
	fmt.Printf("%T", string(c.Body()))
	

	c.Cookie(&fiber.Cookie{
        Name:  "test",
        Value: "SomeThing",
		Path: "/register",
    })
	pgdb := db.Get()
	
	member := register{}
	if err := c.BodyParser(&member); err != nil {
		return err
	}
	fmt.Printf("%+v", member)
	modelPerson := member.Person
	modelPerson.Uuid = uuid.New()

	//insert person
	if err := modelPerson.Insert(c, pgdb); err != nil {
		return c.JSON(FailureResp(c, err, "insert failed in person"))
	}

	person_id := modelPerson.Id

	// get person id and insert user
	member.User.Person_id = person_id
	pass := member.User.Password_hash
	pass_hash := util.SHAEncoding(pass)
	member.User.Password_hash = pass_hash

	if err := member.User.Insert(c, pgdb); err != nil {
		return c.JSON(FailureResp(c, err, "insert failed in user"))
	}
	
	// return c.SendString("Welcome On-board " + c.Params("name"))
	return c.JSON(models.Resp{Message: "you have successfully resigstered", Data: modelPerson})
}

func GetUser(c *fiber.Ctx) error {
	pgdb := db.Get()
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.JSON(FailureResp(c, errors.New("json parsing failed"), "no field in db"))
	}
	person := models.User{}
	// err := pgdb.Model(&person).Where("user_name = ?", user.UserName).Column("user_name", "id").Select();
	err := pgdb.Model(&person).Relation("Person").Where("user_name = ?", user.UserName).Column("user.user_name").Select();
	if err != nil {
		return c.JSON(FailureResp(c, err, "no user found"))
	}
	// pgdb.Model(&person).Where("user_name = ?", user.UserName).ColumnExpr("count()").Select(&x);
	// pgdb.Model(&user).WherePK().Select();
	fmt.Printf("%+v\n", person)
	return c.JSON(person)
}