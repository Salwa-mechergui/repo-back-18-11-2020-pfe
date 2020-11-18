package Migration

import (
	"fmt"
	"github/reccrides/Helpers"
	"github/reccrides/Models"
	"log"

	"github.com/jinzhu/gorm"
)

//init migration to connect to db by parameters
func InitMigration() *gorm.DB {
	connection := Helpers.GetVarFromEnv()
	db, err := gorm.Open("postgres", connection)
	fmt.Printf(connection)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

//Migration to db
func Migration() {
	db := InitMigration()
	db.AutoMigrate(&Models.Front_end_user{})
	db.AutoMigrate(&Models.Passengers{})
	db.AutoMigrate(&Models.Company{})
	db.AutoMigrate(&Models.Requests{})
	db.AutoMigrate(&Models.Customer{})
	db.AutoMigrate(&Models.VehicleType{})
}
