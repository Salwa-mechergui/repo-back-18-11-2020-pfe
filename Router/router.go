package Router

import (
	"github/reccrides/controllers"

	"github.com/gorilla/mux"
)

//router setup
func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	//Route handles & endpoints

	// route to the authentication
	router.HandleFunc("/login", controllers.CreateEndpoint).Methods("POST")
	// route to get the front_end_users
	router.HandleFunc("/customer/{saas_company_id}", controllers.GetAllCustomers).Methods("GET")
	// route to get the front_end_user by Id
	router.HandleFunc("/customer/{saas_company_id}/{id_customer}", controllers.GetCustomersById).Methods("GET")
	// route to get the passengers
	router.HandleFunc("/passenger/{customer_id}", controllers.GetAllPassengersEndpoint).Methods("GET")
	// route to get the passengers by Id
	router.HandleFunc("/passenger/{customer_id}/{id}", controllers.GetPassengersByIdEndpoint).Methods("GET")
	// route to get the companies
	router.HandleFunc("/company/{saas_company_id}/{customer_id}", controllers.GetAllCompaniesEndpoint).Methods("GET")
	// route to get the companies by Id
	router.HandleFunc("/company/{saas_company_id}/{customer_id}/{company_id}", controllers.GetCompaniesByIdEndpoint).Methods("GET")
	// route to get the rides
	router.HandleFunc("/ride/{saas_company_id}/{customer_id}/{state}", controllers.GetRidesEndpoint).Methods("GET")
	// route to get the vehicule
	router.HandleFunc("/direction", controllers.GetVehicleTypes).Methods("GET")
	router.HandleFunc("/vehicule/{saas_officeID}", controllers.GetVehicleTypes).Methods("GET")
	// router.HandleFunc("/customer", controllers.RegisterCustomer).Methods("POST")
	// router.HandleFunc("/passenger", controllers.RegisterPassenger).Methods("POST")
	router.HandleFunc("/location", controllers.ApiGetLocation).Methods("GET")
	// Create
	// router.HandleFunc("/ridecreation", RideCreation).Methods("POST")
	return router
}
