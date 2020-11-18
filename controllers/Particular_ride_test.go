package controllers

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gorilla/mux"
// 	"github.com/stretchr/testify/assert"
// )

// //test_to_get_front_end_users
// func TestGetFrontUser(t *testing.T) {
// 	r, err := http.NewRequest("GET", "/customer/1", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	w := httptest.NewRecorder()
// 	handler := http.HandlerFunc(GetAllCustomers)
// 	handler.ServeHTTP(w, r)
// 	if status := w.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}
// }

// //test_to_get_front_end_users_by_id
// // func TestGetByIdFrontUser(t *testing.T) {
// // 	req, err := http.NewRequest("GET", "/customer/1/9738", nil)

// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}

// // 	rr := httptest.NewRecorder()

// // 	r := mux.NewRouter()

// // 	r.HandleFunc("/customer/{saas_company_id}/{id}", GetCustomersById).Methods("GET")

// // 	r.ServeHTTP(rr, req)

// // 	if status := rr.Code; status != http.StatusOK {
// // 		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
// // 	}

// // 	expected := `{
// // 		"id_front_user": 9738,
// // 		"front_saas_company_id": "1",
// // 		"first_name": "Rahma",
// // 		"last_name": "Waja",
// // 		"email": "rahma@yusofleet.com",
// // 		"phone_number": "+2165526536555167737"
// // 	}`

// // 	assert.JSONEq(t, expected, rr.Body.String(), "Response body differs")
// // }

// //test_to_get_passengers_related_to_front_end_users
// func TestGetPassenger(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/passenger/9898", nil)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()

// 	r := mux.NewRouter()

// 	r.HandleFunc("/passenger/{customer_id}", GetAllPassengersEndpoint).Methods("GET")

// 	r.ServeHTTP(rr, req)

// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
// 	}

// 	expected := `{ 
// 		"id": "12138",
// 		"customer_id": "9898",
// 		"first_name": "Rahma",
// 		"last_name": "Waja (Gmail)",
// 		"email": "rahmawaja@gmail.com",
// 		"phone_number": "+21655167737",
// 		"locale": "fr",
// 		"send_invoice": true,
// 		"show_price": "true",
// 		"type": "RidePassenger"
// 	}`

// 	assert.JSONEq(t, expected, rr.Body.String(), "Response body differs")
// }

// //test_to_get_passengers_related_to_front_end_users_by_id
// func TestGetByIdPassenger(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/passenger/9898/12138", nil)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()

// 	r := mux.NewRouter()

// 	r.HandleFunc("/passenger/{customer_id}/{id}", GetPassengersByIdEndpoint).Methods("GET")

// 	r.ServeHTTP(rr, req)

// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
// 	}

// 	expected := `{ 
// 		"id": "12138",
// 		"customer_id": "9898",
// 		"first_name": "Rahma",
// 		"last_name": "Waja (Gmail)",
// 		"email": "rahmawaja@gmail.com",
// 		"phone_number": "+21655167737",
// 		"locale": "fr",
// 		"send_invoice": true,
// 		"show_price": "true",
// 		"type": "RidePassenger"
// 	}`

// 	assert.JSONEq(t, expected, rr.Body.String(), "Response body differs")
// }

// //test_to_get_companies_related_to_front_end_users
// func TestGetCompanies(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/companies/1", nil)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()

// 	r := mux.NewRouter()

// 	r.HandleFunc("/companies/{saas_company_id}", GetAllCompaniesEndpoint).Methods("GET")

// 	r.ServeHTTP(rr, req)

// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
// 	}

// 	expected := `{ "id": "710",
// 	"front_saas_company_id": "1",
// 	"name": "achrefffco",
// 	"default_comment": "testzizou"}`

// 	assert.JSONEq(t, expected, rr.Body.String(), "Response body differs")
// }

// //test_to_get_companies_related_to_front_end_users_by_id
// func TestGetByIdCompanies(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/companies/1/710", nil)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()

// 	r := mux.NewRouter()

// 	r.HandleFunc("/companies/{saas_company_id}/{id}", GetCompaniesByIdEndpoint).Methods("GET")

// 	r.ServeHTTP(rr, req)

// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
// 	}

// 	expected := `{ "id": "710",
// 	"front_saas_company_id": "1",
// 	"name": "achrefffco",
// 	"default_comment": "testzizou"}`

// 	assert.JSONEq(t, expected, rr.Body.String(), "Response body differs")
// }

// //test_to_get_rides_related_to_front_end_users
// func TestGetRides(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/rides/1/9898", nil)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()

// 	r := mux.NewRouter()

// 	r.HandleFunc("/rides/{saas_company_id}/{customer_id}", GetRidesEndpoint).Methods("GET")

// 	r.ServeHTTP(rr, req)

// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
// 	}

// 	expected := `{
// 		"id": "60527",
// 		"front_saas_company_id": "1",
// 		"customer_id": "9898",
// 		"estimate_drop_off": "2017-05-04T21:15:00Z",
// 		"estimate_pick_up": "2017-05-04T21:00:00Z"
// 	}`

// 	assert.JSONEq(t, expected, rr.Body.String(), "Response body differs")
// }

// //test_to_get_rides_related_to_front_end_users_by_id
// func TestGetByIdRides(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/rides/1/9898/60527", nil)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()

// 	r := mux.NewRouter()

// 	r.HandleFunc("/rides/{saas_company_id}/{customer_id}/{id}", GetRidesByIdEndpoint).Methods("GET")

// 	r.ServeHTTP(rr, req)

// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
// 	}

// 	expected := `{
// 		"id": "60527",
// 		"front_saas_company_id": "1",
// 		"customer_id": "9898",
// 		"estimate_drop_off": "2017-05-04T21:15:00Z",
// 		"estimate_pick_up": "2017-05-04T21:00:00Z"
// 	}`

// 	assert.JSONEq(t, expected, rr.Body.String(), "Response body differs")
// }
