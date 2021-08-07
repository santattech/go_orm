package utility

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "gorm_posts"
)

func GetConnection() *gorm.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
		panic("failed to connect to the database")
		log.Fatal(err)
	} else {
		log.Println("DB Connection established...")
	}

	return db
}
