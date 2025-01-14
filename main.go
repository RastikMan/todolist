package main

import (
	"log"
	"net/http"
	"todo-list/handlers"
	"todo-list/storage"

	"github.com/gorilla/mux"
)

func main() {
	store := storage.NewMemoryStore()
	handler := &handlers.TaskHandler{Store: store}

	router := mux.NewRouter()
	router.HandleFunc("/tasks", handler.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", handler.AddTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", handler.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", handler.DeleteTask).Methods("DELETE")

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
