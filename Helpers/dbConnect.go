package Helpers

import (
	"fmt"
	"os"
)

//get environnement parameters
func GetVarFromEnv() string {
	host := os.Getenv("DBHOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("DBPORT")
	if port == "" {
		port = "5432"
	}
	user := os.Getenv("DBUSER")
	if user == "" {
		user = "user"
	}
	password := os.Getenv("DBPASS")
	if password == "" {
		password = "123456"
	}
	dbname := os.Getenv("DBDATABASE")
	if dbname == "" {
		dbname = "bootcamp"
	}
	if os.Getenv("DATABASE_URL") != "" {
		ConnectionString := os.Getenv("DATABASE_URL")
		return ConnectionString
	}
	ConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	return ConnectionString
}

// //init migration to connect to db by parameters
// func InitMigration() *gorm.DB {
// 	connection := GetVarFromEnv()
// 	db, err := gorm.Open("postgres", connection)
// 	fmt.Printf(connection)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return db
// }

// //Migration to db
// func Migration() {
// 	db := InitMigration()
// 	db.AutoMigrate(&Models.Front_end_user{})
// 	db.AutoMigrate(&Models.Passenger{})
// 	db.AutoMigrate(&Models.Company{})
// 	db.AutoMigrate(&Models.Ride{})
// }
