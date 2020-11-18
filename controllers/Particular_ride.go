package controllers

import (
	"encoding/json"
	"fmt"
	"github/reccrides/Migration"
	"github/reccrides/Models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetAllCustomers godoc
// @Summary Get details of all customers
// @Description Get description of all customers
// @Tags customers
// @Accept  json
// @Produce  json
// @Param  saas_company_id path uint true "id of the company "
// @Success 200 {object} Models.Customer
// @Router /customer/{saas_company_id} [get]
func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := Migration.InitMigration()
	const query = ` 
	select 
	customer.id,
	customer.saas_company_id,
	front_end_user.id,
	front_end_user.last_name,
	front_end_user.first_name,
	front_end_user.email,
	front_end_user.phone_number,
	customer.send_pdf_email,
	customer.default_comment,
	customer.landline_number,
	customer.default_passenger_id,
	customer.front_end_user_id
	from customers customer
	LEFT OUTER JOIN front_end_users front_end_user on front_end_user.id = customer.front_end_user_id
	`
	var customers []Models.Customer
	err1 := db.Raw(query).Scan(&customers).Error
	if err1 != nil {
		fmt.Println(err1)
	}
	params := mux.Vars(r)
	result := make([]Models.Customer, 0)
	for _, item := range customers {
		saas_company_id, err := strconv.ParseUint(params["saas_company_id"], 10, 32)
		if err != nil {
			fmt.Println(err)
		}
		if uint(saas_company_id) == item.Saas_company_id {
			result = append(result, item)
		}
	}
	json.NewEncoder(w).Encode(result)
}

// GetAllCustomers godoc
// @Summary Get details of all customers
// @Description Get description of all customers
// @Tags customers
// @Accept  json
// @Produce  json
// @Param  saas_company_id path uint true "id of the company "
// @Param  id_customer path uint true "id of the customer "
// @Success 200 {object} Models.Customer
// @Router /customer/{saas_company_id}/{id_customer} [get]
func GetCustomersById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := Migration.InitMigration()
	const query = ` 
	select 
	customer.id,
	customer.saas_company_id,
	front_end_user.id,
	front_end_user.last_name,
	front_end_user.first_name,
	front_end_user.email,
	front_end_user.phone_number,
	customer.send_pdf_email,
	customer.default_comment,
	customer.landline_number,
	customer.default_passenger_id,
	customer.front_end_user_id,
	customer.referral_code
	from customers customer
	LEFT OUTER JOIN front_end_users front_end_user on front_end_user.id = customer.front_end_user_id
	`
	var customers []Models.Customer
	err1 := db.Raw(query).Scan(&customers).Error
	if err1 != nil {
		fmt.Println(err1)
	}
	params := mux.Vars(r)
	result := make([]Models.Customer, 0)
	for _, item := range customers {
		saas_company_id, err := strconv.ParseUint(params["saas_company_id"], 10, 32)
		id_customer, err := strconv.ParseUint(params["id_customer"], 10, 32)
		if err != nil {
			fmt.Println(err)
		}
		if uint(saas_company_id) == item.Saas_company_id && uint(id_customer) == item.Id {
			result = append(result, item)
		}
	}
	json.NewEncoder(w).Encode(result)
}

// GetAllPassengersEndpoint godoc
// @Summary Get details of all passengers
// @Description Get description of all passengers related to front_end_user
// @Tags passengers
// @Accept  json
// @Produce  json
// @Param  customer_id path uint true "id of the customers"
// @Success 200 {object} Models.Passengers
// @Router /passenger/{customer_id} [get]
func GetAllPassengersEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := Migration.InitMigration()
	query := `
	select
	customer.saas_company_id,
	passenger.customer_id,
	passenger.id,
	passenger.last_name,
	passenger.first_name,
	passenger.email,
	passenger.phone_number,
	customer.send_pdf_email,
	customer.default_comment,
	customer.landline_number,
	customer.default_passenger_id,
	customer.saas_company_id,
	customer.front_end_user_id
	from customers customer
	LEFT OUTER JOIN passengers passenger on passenger.customer_id = customer.id
	`

	var passengers []Models.Passengers
	err1 := db.Raw(query).Scan(&passengers).Error
	if err1 != nil {
		fmt.Println(err1)
	}
	params := mux.Vars(r)
	result := make([]Models.Passengers, 0)
	for _, item := range passengers {
		customer_id, err := strconv.ParseUint(params["customer_id"], 10, 32)
		if err != nil {
			fmt.Println(err)
		}
		if uint(customer_id) == item.Customer_id {
			result = append(result, item)
		}
	}
	json.NewEncoder(w).Encode(result)
}

// GetAllPassengersByIdEndpoint godoc
// @Summary Get details of all passengers
// @Description Get description of passengers related to front_end_user by id
// @Tags passengers
// @Accept  json
// @Produce  json
// @Param  customer_id path uint true "id of the customers"
// @Param  id path uint true "id of the passenger"
// @Success 200 {object} Models.Passengers
// @Router /passenger/{customer_id}/{id} [get]
func GetPassengersByIdEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := Migration.InitMigration()
	query := `
	select
	customer.saas_company_id,
	passenger.customer_id,
	passenger.id,
	passenger.last_name,
	passenger.first_name,
	passenger.email,
	passenger.phone_number,
	customer.send_pdf_email,
	customer.default_comment,
	customer.landline_number,
	customer.default_passenger_id,
	customer.saas_company_id,
	customer.front_end_user_id
	from customers customer
	LEFT OUTER JOIN passengers passenger on passenger.customer_id = customer.id
	`

	var passengers []Models.Passengers
	err1 := db.Raw(query).Scan(&passengers).Error
	if err1 != nil {
		fmt.Println(err1)
	}
	params := mux.Vars(r)
	result := make([]Models.Passengers, 0)
	for _, item := range passengers {
		customer_id, err := strconv.ParseUint(params["customer_id"], 10, 32)
		id, err := strconv.ParseUint(params["id"], 10, 32)
		if err != nil {
			fmt.Println(err)
		}
		if uint(customer_id) == item.Customer_id && uint(id) == item.Id {
			result = append(result, item)
		}
	}
	json.NewEncoder(w).Encode(result)
}

// GetAllCompaniesEndpoint godoc
// @Summary Get details of companies
// @Description Get description of companies related to front_end_user
// @Tags companies
// @Accept  json
// @Produce  json
// @Param  saas_company_id path uint true "id of the company"
// @Param  customer_id path uint true "id of the customer"
// @Success 200 {object} Models.Company
// @Router /companies/{saas_company_id}/{customer_id} [get]
func GetAllCompaniesEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := Migration.InitMigration()
	query := `
	select
	request.id,
	request.saas_company_id,
	request.customer_id,
	request.company_id,
	request.comment,
	request.passenger_id,
	company.name,
	company.company_address,
	company.default_comment,
	company.payment_type,
	company.company_siren,
	company.enable_send_invoice,
	company.payment_period_type,
	company.receive_email
	from requests request
	LEFT OUTER JOIN companies company on company.id = request.company_id
	LEFT OUTER JOIN customers customer on customer.id=request.customer_id
	where company.name IS NOT NULL
	`
	var requests []Models.Requests
	err1 := db.Raw(query).Scan(&requests).Error
	if err1 != nil {
		fmt.Println(err1)
	}
	params := mux.Vars(r)
	result := make([]Models.Requests, 0)
	for _, item := range requests {
		saas_company_id, err := strconv.ParseUint(params["saas_company_id"], 10, 32)
		customer_id, err := strconv.ParseUint(params["customer_id"], 10, 32)
		// id, err := strconv.ParseUint(params["id"], 10, 32)
		if err != nil {
			fmt.Println(err)
		}
		if uint(saas_company_id) == item.Saas_company_id && uint(customer_id) == item.Customer_id {
			result = append(result, item)
		}
	}
	json.NewEncoder(w).Encode(result)
}

// GetCompaniesByIdEndpoint godoc
// @Summary Get details of all companies
// @Description Get description of companies related to front_end_users by id
// @Tags companies
// @Accept  json
// @Produce  json
// @Param  saas_company_id path string true "id of the company"
// @Param  customer_id path uint true "id of the customer"
// @Param  company_id path uint true "id of the company"
// @Success 200 {object} Models.Company
// @Router /companies/{saas_company_id}/{customer_id}/{company_id} [get]
func GetCompaniesByIdEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := Migration.InitMigration()
	query := `
	select
	request.id,
	request.saas_company_id,
	request.customer_id,
	request.company_id,
	request.comment,
	request.passenger_id,
	company.name,
	company.company_address,
	company.default_comment,
	company.payment_type,
	company.company_siren,
	company.enable_send_invoice,
	company.payment_period_type,
	company.receive_email
	from requests request
	LEFT OUTER JOIN companies company on company.id = request.company_id
	LEFT OUTER JOIN customers customer on customer.id=request.customer_id
	where company.name IS NOT NULLs
	`
	var requests []Models.Requests
	err1 := db.Raw(query).Scan(&requests).Error
	if err1 != nil {
		fmt.Println(err1)
	}
	params := mux.Vars(r)
	result := make([]Models.Requests, 0)
	for _, item := range requests {
		saas_company_id, err := strconv.ParseUint(params["saas_company_id"], 10, 32)
		customer_id, err := strconv.ParseUint(params["customer_id"], 10, 32)
		company_id, err := strconv.ParseUint(params["company_id"], 10, 32)
		if err != nil {
			fmt.Println(err)
		}
		if uint(saas_company_id) == item.Saas_company_id && uint(customer_id) == item.Customer_id && uint(company_id) == item.Company_id {
			result = append(result, item)
		}
	}
	json.NewEncoder(w).Encode(result)
}

// GetRidesEndpoint godoc
// @Summary Get details of all rides
// @Description Get description of upcoming and current rides
// @Tags rides
// @Accept  json
// @Produce  json
// @Param  saas_company_id path uint true "id of the company"
// @Param  customer_id path uint true "customer_id of the companies"
// @Param  state path string true "id of the companies"
// @Success 200 {object} Models.Requests
// @Router /rides/{saas_company_id}/{customer_id}/{state} [get]
func GetRidesEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := Migration.InitMigration()
	query := `
	select
	request.state,
	request.id
	from requests request
	LEFT OUTER JOIN companies company on company.id = request.company_id
	LEFT OUTER JOIN customers customer on customer.id=request.customer_id
	where request.state='way_to'
	limit 3
	`
	var requests []Models.Requests
	var err error
	if err != nil {
		fmt.Printf("The request failed with error %s\n", err)
	}
	params := mux.Vars(r)
	// request_id, err := strconv.ParseUint(params["request_id"], 10, 32)
	saas_company_id, err := strconv.ParseUint(params["saas_company_id"], 10, 32)
	customer_id, err := strconv.ParseUint(params["customer_id"], 10, 32)
	err = db.Raw(query).Scan(&requests).Error
	json.NewEncoder(w).Encode(db.Find(&requests, &Models.Requests{State: params["state"],
		Saas_company_id: uint(saas_company_id), Customer_id: uint(customer_id)}))
}
