package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/rafaelSoaresAlmeida/ecom-api/cmd/api"
	"github.com/rafaelSoaresAlmeida/ecom-api/config"
	"github.com/rafaelSoaresAlmeida/ecom-api/db"
)

func main() {
	
	db, err := db.NewSqlStorage(mysql.Config {
		User:  config.Envs.DbUser,
		Passwd: config.Envs.DbPassword,
		Addr: config.Envs.DbAddress,
		DBName: config.Envs.DbName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)
	
	server := api.NewApiServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB up and running...")
}