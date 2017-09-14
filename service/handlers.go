package service

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
    "github.com/validatepolicy/dbclient"
    "fmt"
)

var DBClient dbclient.IBoltClient

func ValidatePolicy(w http.ResponseWriter, r *http.Request) {

	// Read the 'policyNumber' path parameter from the mux map
	var policyNumber = mux.Vars(r)["policyNumber"]

        // Read the policy struct BoltDB
	policy, err := DBClient.QueryPolicy(policyNumber)

        // If err, return a 404
	if err != nil {
                fmt.Println("Some error occured serving " + policyNumber + ": " + err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

        // If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(policy)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}