package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pdkkid/go_api_backend/cassandra"
	"github.com/pdkkid/go_api_backend/product"
)

type heartbeatRes struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func main() {
	CassandraSession := cassandra.Session
	defer CassandraSession.Close()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", heartbeat)
	router.HandleFunc("/product", product.Post)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Heartbeat sent")
	json.NewEncoder(w).Encode(heartbeatRes{Status: "OK", Code: 200})
}
