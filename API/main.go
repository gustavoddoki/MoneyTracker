package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gustavoddoki/MoneyTracker/API/configs"
	"github.com/gustavoddoki/MoneyTracker/API/db"
	"github.com/gustavoddoki/MoneyTracker/API/handlers"
)

func main() {

	db.CreateDatabase()

	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.AddTransactionHandler)
	r.Put("/{id}", handlers.UpdateTransactionHandler)
	r.Delete("/{id}", handlers.DeleteTransactionHandler)
	r.Get("/", handlers.GetTransactionHandler)
	r.Get("/{id}", handlers.GetTransactionByIDHandler)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
