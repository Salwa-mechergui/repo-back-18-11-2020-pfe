package controllers

import (
	"encoding/json"
	"github/reccrides/Migration"
	"github/reccrides/Models"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllCustomers godoc
// @Summary Get
// @Description
// @Tags VehicleType
// @Accept  json
// @Produce  json
// @Param  saas_officeID path uint true "id of the saas_office"
// @Success 200 {object} Models.VehicleType
// @Router /vehicule/{saas_officeID} [get]
func GetVehicleTypes(w http.ResponseWriter, r *http.Request) {
	db := Migration.InitMigration()
	params := mux.Vars(r)
	saasParam := params["saas_officeID"]
	var types []Models.VehicleType
	db.Where("saas_office_id = ?", saasParam).Where("is_available = ?", true).Where("deleted = ?", false).Find(&types)
	if len(types) != 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(types)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("data not found"))
	}
}
