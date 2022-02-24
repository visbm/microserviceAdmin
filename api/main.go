package main

import (
	"database/sql"
	"fmt"
	"log"
	"microseviceAdmin/domain/model"
	"microseviceAdmin/webapp"
	"microseviceAdmin/webapp/server"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(filepath.Join("../", ".env"))
	if err != nil {
		log.Printf("Error loading .env file. %s", err.Error())
	}

	config := &webapp.Config{}
	config.NewConfig()
	initAccountDb(config)

	s := server.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}

// initAccountDb hashes accounts password from  DB, that creates by ini.sql
func initAccountDb(c *webapp.Config) {
	dataSourceName := c.PgDataSource()

	log.Printf("Connecting to database via %#v", dataSourceName)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 6; i++ {

		encryptedPassword, err := model.EncryptPassword(fmt.Sprintf("password"))
		if err != nil {
			log.Print(err)
		}

		result, err := db.Exec("UPDATE USERS set password = $1 WHERE id = $2",
			encryptedPassword, i+1)
		if err != nil {
			log.Print(err)
		}
		log.Println("password encrypted")
		log.Println(result.RowsAffected())
	}

}
