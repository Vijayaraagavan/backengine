package scripts

import (
	// "bufio"
	"fmt"
	// "log"
	// "os"

	pg "github.com/go-pg/pg/v10"

	"backengine/models"
)

func initUsers(pgdb *pg.DB) {
	// users := []models.User{
	// 	{
	// 		// Id: 1,
	// 		// Person_id: 1,
	// 		Password_hash: "4f3s4df3sdf3s1s",
	// 	},
	// 	{
	// 		// Id: 2,
	// 		// Person_id: 2,
	// 		Password_hash: "4d42sf2ss2Sfs",
	// 	},
	// }

	// persons := []models.Person{
	// 	{
	// 		// Id: 1,
	// 		Name: "vijayaragavan",
	// 		PhoneNo: 9134343131,
	// 		UserId: 1,
	// 	},
	// 	{
	// 		// Id: 2,
	// 		Name: "Hari",
	// 		PhoneNo: 894641343,
	// 		UserId: 2,
	// 	},
	// }
	// qs := []string{
		// "INSERT INTO users VALUES (1, '4313131sdsdsds')",
    // "INSERT INTO persons VALUES (1, 'vijayaragv', 913434311), (2, 'hari', 876443434)",
	// }
	// for _, q := range qs {
	// 	_, err := pgdb.Exec(q)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// obj := models.Person{Name: "chopper", PhoneNo: 943543434, UserId: 1}
	// obj := []models.User{}
	// _ = pgdb.Model(&obj).Where("id = ?", 2).Select()
	// fmt.Println(obj)

	ob := models.Person{}
	_ = pgdb.Model(&ob).Relation("User").Where("person.id = ?", 1).Select()
	fmt.Printf("Person %+v", ob)
	// _, err := pgdb.Model(&obj).Insert()
	// if err != nil {
	// 			panic(err)
	// 		}
	// fmt.Println(obj)
	// a := new(models.Person) 
	
	// pgdb.Model(a).Where("id = ?", 2).Column("id", "password_hash").Select()
	// pgdb.Query(a, "SELECT * from users where id = ?", 1)
	// b := a
	// fmt.Println(*a)
	// fmt.Println(*b)

	// for _, user := range users {
	// 	fmt.Println(user)
	// 	if _, err := pgdb.Model(&user).Insert(); err != nil {
	// 		fmt.Println("error in initiating user table", err)
	// 	}
	// }
	// for _, person := range persons {
	// 	if _, err := pgdb.Model(&person).Insert(); err != nil {
	// 		fmt.Println("error in initiating person table", err)
	// 	}
	// }
}

func initMasters(pgdb *pg.DB) {
	initUsers(pgdb)
	
}