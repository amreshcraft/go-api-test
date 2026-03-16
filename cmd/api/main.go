package main 

import (
	"log"
	"net/http"
	"github.com/amreshcraft/go-api-test/internal/handler"
)

func main()  {
	http.HandleFunc("/hello", handler.HelloHandler)

	log.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}