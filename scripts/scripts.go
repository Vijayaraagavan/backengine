package scripts

import (
	"fmt"

	pg "github.com/go-pg/pg/v10"
	orm "github.com/go-pg/pg/v10/orm"

	models "backengine/models"
	db "backengine/db"
)

func getModels(pgdb *pg.DB) []interface{} {
	return []interface{}{
		&models.Person{},
		&models.User{},
	}
}

func createTables(pgdb *pg.DB) {
	for _, schema := range getModels(pgdb) {
		if err := pgdb.Model(schema).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
			FKConstraints: true,
			
		}); err != nil {
			fmt.Printf("cannot create table => %v because %v\n", schema, err)
		}

	}
	// qs := []string{
	// 	"CREATE TABLE users (id int, password_hash text)",
	// 	"CREATE TABLE persons (id int, name text, phone_no int, user_id int)",
	// }
	// for _, q := range qs {
	// 	_, err := pgdb.Exec(q)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
}

func createIndexes(pgdb *pg.DB) {
	for _, q := range []string{
		fmt.Sprintf("CREATE INDEX userId_idx on %s(person_id)", db.User),
	} {
		if _, err := pgdb.Exec(q); err != nil {
			fmt.Printf("cannot create index: %v\n", err)
		}
	}
}

func Run(pgdb *pg.DB) {
	createTables(pgdb)
	fmt.Println("done")
	// createIndexes(pgdb)

	initMasters(pgdb)
}