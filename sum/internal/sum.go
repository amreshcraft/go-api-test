package internal

import (
	"encoding/json"
	"net/http"
	"sum/service"
)

type RequestFormat struct {
	A int `json:"a"`
	B int `json:"b"`
}

type ResponseFormat struct {
	Result int `json:"result"`
}



func HandlerSum(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	var data RequestFormat

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&data); err != nil {
		http.Error(w, "Invalid JSON format of data", http.StatusBadRequest)
		return
	}

	// validation
	if data.A == 0 || data.B == 0 {
		http.Error(w, "Fields 'a' and 'b' are required", http.StatusBadRequest)
		return
	}

	result := service.Sum(data.A, data.B)

	response := ResponseFormat{
		Result: result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}