package Models

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Passengers struct to get list of passengers
type Passengers struct {
	Saas_company_id     uint   `json:"company_id,omitempty"`
	Id                  uint   `json:"id,omitempty"`
	Customer_id         uint   `json:"customer_id,omitempty"`
	First_name          string `json:"first_name,omitempty"`
	Last_name           string `json:"last_name,omitempty"`
	Email               string `json:"email,omitempty"`
	Phone_number        string `json:"phone_number,omitempty"`
	Company_id          string `json:"company_id,omitempty"`
	language            string `json:"language,omitempty"`
	Deleted             bool   `json:"deleted,omitempty"`
	Locale              string `json:"locale,omitempty"`
	Send_invoice        bool   `json:"send_invoice,omitempty"`
	Show_price          string `json:"show_price,omitempty"`
	Type                string `json:"type,omitempty"`
	Default_favorite_id int    `json:"default_favorite_id,omitempty"`
	Default_comment     string `json:"default_comment,omitempty"`
	// Requests            []Requests `gorm:"many2many:requests_passenger;"`
}

//front_end_user struct to get list of front end users
type Front_end_user struct {
	Id              uint   `json:"id_front_user,omitempty"`
	Saas_company_id string `json:"front_saas_company_id,omitempty"`
	First_name      string `json:"first_name,omitempty"`
	Last_name       string `json:"last_name,omitempty"`
	Email           string `json:"email,omitempty"`
	Phone_number    string `json:"phone_number,omitempty"`
	Send_pdf_mail   string `json:"send_pdf_email,omitempty"`
	second_email    string `json:"second_email,omitempty"`
	Default_comment string `json:"default_comment,omitempty"`
	Landline_number string `json:"landline_number,omitempty"`
}

//Company struct to get list of companies
type Company struct {
	Id                  string `json:"id_front,omitempty"`
	Saas_company_id     string `json:"front_saas_company_id,omitempty"`
	Name                string `json:"name,omitempty"`
	Company_address     string `json:"company_address,omitempty"`
	Default_comment     string `json:"default_comment,omitempty"`
	Payement_type       string `json:"payement_type,omitempty"`
	Company_siren       string `json:"company_siren,omitempty"`
	enable_send_invoice string `json:"enable_send_invoice,omitempty"`
	payment_period_type string `json:"payment_period_type,omitempty"`
	receive_email       string `json:"receive_email,omitempty"`
}

//Company struct to get list of companies
type Customer struct {
	Id                   uint   `json:"id_customer,omitempty"`
	Saas_company_id      uint   `json:"saas_company_id,omitempty"`
	Front_end_user_id    uint   `json:"front_end_user_id,omitempty"`
	First_name           string `json:"first_name,omitempty"`
	Last_name            string `json:"last_name,omitempty"`
	Email                string `json:"email,omitempty"`
	Phone_number         string `json:"phone_number,omitempty"`
	Comment              string `json:"Comment,omitempty"`
	Referral_code        string `json:"referral_code,omitempty"`
	Default_passenger_id string `json:"default_passenger_id,omitempty"`
	Send_pdf_mail        string `json:"send_pdf_email,omitempty"`
	second_email         string `json:"second_email,omitempty"`
	Default_comment      string `json:"default_comment,omitempty"`
	Landline_number      string `json:"landline_number,omitempty"`
}

//request struct
type Requests struct {
	ID              uint   `json:"request_id,omitempty"`
	Saas_office_id  string `json:"saas_office_id,omitempty"`
	Saas_company_id uint   `json:"saas_company_id,omitempty"`
	Customer_id     uint   `json:"customer_id,omitempty"`
	Passenger_id    string `json:"passenger_id,omitempty"`
	Company_id      uint   `json:"company_id,omitempty"`
	Comment         string `json:"Comment,omitempty"`
	State           string `json:"state,omitempty"`
	First_name_fr   string `json:"first_name,omitempty"`
	Last_name_fr    string `json:"last_name,omitempty"`
	Email_fr        string `json:"email,omitempty"`
	Phone_number_fr string `json:"phone_number,omitempty"`
	First_name      string `json:"first_name,omitempty"`
	Last_name       string `json:"last_name,omitempty"`
	Email           string `json:"email,omitempty"`
	Phone_number    string `json:"phone_number,omitempty"`
	Id_ride         string `json:"id_ride,omitempty"`
	Id              uint   `json:"ride_id,omitempty"`
	Name            string `json:"Name,omitempty"`
	Referral_code   string `json:"Referral_code,omitempty"`
}

//rides struct
type Ride struct {
	ID              string `json:"id"`
	Ride_request_id string `json:"Ride_request_id"`
}

// VehicleType struct
type VehicleType struct {
	ID            uint   `gorm:"primary_key"`
	Designation   string `gorm:"type:varchar(255)"`
	IsAvailable   bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	KmPrice       int
	MinPrice      int
	PickupPrice   int
	TypeCode      int
	SaasCompanyID int
	saas_officeID int
	MaxPlace      int
	LogoVehicle   string `gorm:"type:varchar(255)"`
	name          string
	lat           int
	long          int
	types         string
	zip_code      string
	complement    string
}
type Addresse struct {
}
type Driver struct {
}
type SaasCompanyToDriverRelation struct {
}
type SaasOfficeToDriverRelation struct {
}
type BackEndUser struct {
}
