package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"track_doctor/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectDb() {
	p := config.Config("PG_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println("Error parsing config str to int")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Dhaka", config.Config("PG_HOST"), config.Config("PG_USER"), config.Config("PG_PASSWORD"), config.Config("PG_DBNAME"), port)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal("Failed to connect database", err)
		os.Exit(2)
	}
	log.Println("Connected")
	fmt.Println("DB connected")
	DB = DbInstance{
		Db: db,
	}
}
