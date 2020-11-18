package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ApiGetLocation(w http.ResponseWriter, r *http.Request) {
	url := "https://maps-manager.herokuapp.com/api/directions?provider=google&origin=36.6207016,2.938144&destination=27.3921384,-0.2326099&traffic=true&WithPath=true"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1OTQyMDk4NTMsInVzZXIiOiJ5dXNvIn0.sr_hXVm4aiGsmKZIh2T3Vs2iN7Pbz8isyAS4w9NDdHA"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Token", token)
	resp, err := http.DefaultClient.Do(req) // Call the API
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error to get result from api")
		}
		//Print the Result that is returned in the response of the API
		fmt.Println(string(bodyBytes))
	} else {
		fmt.Println("StatusCode is not ok")
	}
}
