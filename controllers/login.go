package controllers

import(
	"fmt"
	fiber "github.com/gofiber/fiber/v2"
	// pg "github.com/go-pg/pg/v10"
	"errors"
	"backengine/models"
	"backengine/util"
	"backengine/db"
	"backengine/auth/core"
	"log"
)


func Login(c *fiber.Ctx) error {
	fmt.Println("in login verification")
	login := models.Login{}
	if err := c.BodyParser(&login); err != nil {
		return c.JSON(FailureResp(c, err, "error parsing json"))
	}
	fmt.Printf("%+v", login)
	pgdb := db.Get()
	user := models.User{}
	fmt.Println(util.SHAEncoding(login.Password))
	err := pgdb.Model(&user).Relation("Person").Where("user_name = ? and password_hash = ?", login.UserName, util.SHAEncoding(login.Password)).Select()
	// err := pgdb.Model(&user).Where("user_name = ?", login.UserName).Where("password_hash = ?", util.SHAEncoding(login.Password)).Select()
	fmt.Printf("%+v", user)
	if err != nil {
		return c.JSON(FailureResp(c, errors.New(fmt.Sprintf("%v", err)), "no user found"))
	}
	PrepareToken(*user.Person, c)
	return c.JSON(user.Person)
}

func PrepareToken(modelPerson models.Person, c *fiber.Ctx) error {
	//token generation
	backend := Authentication.InitJWTAuthenticationBackend()
	token, err := backend.GenerateToken(modelPerson)
	if err != nil {
		return c.JSON(FailureResp(c, errors.New(fmt.Sprintf("%v", err)), "error in generating token"))
	}
	c.Set("Authorization", "Bearer " + token)
	log.Println(token, err)
	return nil
}

// aug 15 142
// aug 15 207
// somedate within 26 63
// aug 26 165
// sep 10 10
// saturday 30
// sep 17 100
// sep 18 180
// sep 25 120 + 10 
// total 142 + 207 + 63 + 165 + 10 + 30 + 100 + 180 + 130 = 947

