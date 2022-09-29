package db

import (
	pg "github.com/go-pg/pg/v10"
	// orm "github.com/go-pg/pg/v10/orm"
	"sync"
	config "github.com/Vijayaraagavan/backengine/config"
	"time"
	// "context"
	"log"
	"os"
	"fmt"
	// "github.com/Vijayaraagavan/backengine/models"
)

var (
	pgdb *pg.DB
	once sync.Once
)

func dial(cfg config.Config) *pg.DB{
	var options = &pg.Options{
		Network:      "tcp",
		Addr:         cfg.Database.URI,
		User:         cfg.Database.Username,
		Password:     cfg.Database.Password,
		Database:     cfg.Database.Name,
		DialTimeout:  60 * time.Second,
		IdleTimeout:  60 * time.Second,
		PoolSize:     50,
		PoolTimeout:  60 * time.Second,
		ReadTimeout:  90 * time.Second,
		WriteTimeout: 90 * time.Second,
	}

	conn := pg.Connect(options)

	// if err := conn.Ping(ctx); err != nil {
	// 	fmt.Println("error in db connection")
	// }

	
	return conn
}

func Connect() *pg.DB {
	cfg := config.Get()
	pgdb = dial(cfg)
	if pgdb == nil {
		fmt.Println("no db connection")
	}
	log.Println("db connected", pgdb)
	return pgdb
}

func Get() *pg.DB {
	if pgdb == nil {
		log.Println("no db connection, exit db")
		os.Exit(1)
	}
	return pgdb
}

func Close() {
	if pgdb == nil {
		return
	}

	if err := pgdb.Close(); err != nil {
		fmt.Printf("Unable to close PostgreSQL DB while shutdown, err: %v", err)
		return
	}

	log.Println("Closing PostgreSQL DB")
}

// func Createtables(db *pg.DB) {
// 	fmt.Println("here")
// 	fmt.Println((*(models.User))(nil))
// 	newModel := ((*(models.User))(nil))
// 	// fmt.Println(*(models.User))
// 	err := db.Model(newModel).CreateTable(&orm.CreateTableOptions{
// 		IfNotExists: true,
// 	})
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("created tables")
// 	newData := &models.User{
// 		Id: 3,
// 		Person_id: 5,
// 		Password_hash: "541fds3d13s",
// 	}
// 	_, err = db.Model(newData).Insert()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }