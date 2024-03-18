package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/alwarmra/golang-webserver/handlers"
	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := handlers.NewProducts(l)

	r := mux.NewRouter()
	getRouter := r.Methods(http.MethodGet).Subrouter()
	postRouter := r.Methods(http.MethodPost).Subrouter()
	putRouter := r.Methods(http.MethodPut).Subrouter()
	deleteRouter := r.Methods(http.MethodDelete).Subrouter()

	postRouter.Use(ph.MiddleWareValidateProduct)
	putRouter.Use(ph.MiddleWareValidateProduct)

	getRouter.HandleFunc("/", ph.GetProducts)
	postRouter.HandleFunc("/", ph.CreateProduct)
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	deleteRouter.HandleFunc("/{id:[0-9]+}", ph.DeleteProduct)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 180 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal()
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	l.Println("Got signal", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
