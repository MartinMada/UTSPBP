package main

import (
	"UTS_PBP/controller"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/rooms", controller.GetAllRooms).Methods("GET")
	router.HandleFunc("/detail/rooms", controller.GetDetailRooms).Methods("GET")
	router.HandleFunc("/insert/room", controller.InsertRoom).Methods("POST")
	router.HandleFunc("/leave/room", controller.LeaveRoom).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Connected to Port 8080")
	log.Println("Connected to Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
