package main

import (
	"net/http"

	"github.com/uebung/data"
	"github.com/uebung/handlers"
)

func main() {
	mux := http.NewServeMux()

	storeTasks := data.NewStoreTasks()
	taskHandler := handlers.NewToDoHandler(storeTasks)

	mux.HandleFunc("/todo-list", taskHandler.GetToDos)

	http.ListenAndServe(":3000", mux)
}
