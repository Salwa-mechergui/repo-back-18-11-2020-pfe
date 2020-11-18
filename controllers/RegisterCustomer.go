package controllers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"github/reccrides/Migration"
// 	"github/reccrides/Models"
// 	"io/ioutil"
// 	"net/http"
// 	"time"
// )

// type Data struct {
// 	HomeAddress string
// }

// // CreateUser godoc
// // @Summary Create a new user
// // @Description Create a new user with the input payload
// // @Tags back_end_users
// // @Accept  json
// // @Produce  json
// // @Success 200 {object} Models.BackEndUser
// // @Router /register [post]
// func RegisterCustomer(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}
// 	// Unmarshal
// 	var data Data
// 	err = json.Unmarshal(body, &data)
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}
// 	db := Migration.InitMigration()
// 	var count int
// 	db.Table("back_end_users").Where("email = ?", data.Email).Count(&count)
// 	if count == 0 {
// 		// add information in addresses table
// 		address := Models.Addresse{AddressDescription: data.HomeAddress, AddressZipCode: "75075", AddressLat: data.HomeLatitude, AddressLong: data.HomeLongitude, CreatedAt: time.Now(), UpdatedAt: time.Now()}
// 		db.NewRecord(address)
// 		db.Create(&address)
// 		// add information in drivers table
// 		// DriverType == 1 ==> Manager
// 		driver := Models.Driver{DriverType: 1, DriverLicense: data.LicenceNumber, CreatedAt: time.Now(), UpdatedAt: time.Now(), Addresse: address, AddresseID: address.ID, DriverStatus: 0, Active: false, SignUpStep: 1}
// 		db.NewRecord(driver)
// 		db.Create(&driver)
// 		// create saasCompany relation with driver
// 		saasCompanyRelation := Models.SaasCompanyToDriverRelation{SaasCompanyID: data.SaasCompanyID, DriverID: driver.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()}
// 		db.NewRecord(saasCompanyRelation)
// 		db.Create(&saasCompanyRelation)
// 		fmt.Printf("\n company id : %d", saasCompanyRelation.ID)
// 		// create saasOffice relation with driver
// 		saasOfficeRelation := Models.SaasOfficeToDriverRelation{SaasOfficeID: data.Office, DriverID: driver.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()}
// 		db.NewRecord(saasOfficeRelation)
// 		db.Create(&saasOfficeRelation)
// 		fmt.Printf("\n office id : %d", saasOfficeRelation.ID)
// 		// add information in back_end_user table
// 		encryptedPassword, _ := HashPassword(data.Password)
// 		user := Models.BackEndUser{FirstName: data.FirstName, LastName: data.LastName, Email: data.Email, PhoneNumber: data.Number, EncryptedPassword: encryptedPassword, CreatedAt: time.Now(), UpdatedAt: time.Now(), Driver: driver, DriverID: driver.ID}
// 		db.NewRecord(user)
// 		db.Create(&user)
// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(user)
// 	} else {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		w.Write([]byte("address mail already used!"))
// 	}
// }
