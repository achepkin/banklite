package main

import (
	"github.com/achepkin/banklite/internal/app"
	_ "github.com/achepkin/banklite/internal/app"
	"github.com/achepkin/banklite/internal/app/provider"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	cfg := app.Read()
	p := provider.NewProvider(&cfg)

	r := mux.NewRouter()

	//r.HandleFunc("/accounts", handlers.CreateAccount).Methods("POST")
	r.HandleFunc("/accounts/{id}", p.AccountHandler().GetAccount).Methods("GET")
	// Create Account
	r.HandleFunc("/accounts", p.AccountHandler().CreateAccount).Methods("POST")
	r.HandleFunc("/accounts", p.AccountHandler().ListAccounts).Methods("GET")
	r.HandleFunc("/accounts/{id}/transactions", p.TransactionHandler().CreateTransaction).Methods("POST")
	r.HandleFunc("/accounts/{id}/transactions", p.TransactionHandler().GetTransactions).Methods("GET")
	r.HandleFunc("/transfer", p.TransactionHandler().Transfer).Methods("POST")

	log.Printf("Starting HTTP server on http://localhost%s", cfg.HTTPAddr)

	//srv := &http.Server{
	//	Handler: r,
	//	Addr:    cfg.HTTPAddr,
	//}

	//// listen to OS signals and gracefully shutdown HTTP server
	//stopped := make(chan struct{})
	//go func() {
	//	sigint := make(chan os.Signal, 1)
	//	signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	//	<-sigint
	//	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	//	_ = ctx
	//	defer cancel()
	//	if err := srv.Shutdown(ctx); err != nil {
	//		log.Printf("HTTP Server Shutdown Error: %v", err)
	//	}
	//	close(stopped)
	//}()

	//log.Fatal(srv.ListenAndServe())
	err := http.ListenAndServe(cfg.HTTPAddr, r)
	if err != nil {
		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
	}

	//<-stopped
	//
	//log.Printf("Have a nice day!")
}
