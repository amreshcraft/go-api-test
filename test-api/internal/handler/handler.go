package handler

import(
	"encoding/json"
	"net/http"
	"log"
)


type ResponseMessage struct {
	Message string `json:"message"`
}


func HelloHandler(w http.ResponseWriter, r *http.Request){
	responseData := ResponseMessage{
		Message: "Hello World",
	}
  log.Println("Get Request")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData)
}


