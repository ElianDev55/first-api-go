package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ElianDev55/first-api-go/internal/user"
	"github.com/ElianDev55/first-api-go/pkg/bootstrap"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {


	
	router := mux.NewRouter()
	_ = godotenv.Load()
	l := bootstrap.InitLogger()
	db,errDb := bootstrap.DBConnection()
	
	if errDb != nil {
		l.Fatal(errDb)
	}

	



	userRepo := user.NewRepo(l,db)
	userService := user.NewService(l,userRepo)
	userEnd := user.MakeEndPoints(userService)
	
	router.HandleFunc("/users", userEnd.GetAll).Methods("GET")
	router.HandleFunc("/users/{id}", userEnd.Get).Methods("GET")
	router.HandleFunc("/users", userEnd.Create).Methods("POST")
	router.HandleFunc("/users/{id}", userEnd.Update).Methods("PATCH")
	router.HandleFunc("/users/{id}", userEnd.Delete).Methods("DELETE")


	srv := &http.Server{
		Handler: http.TimeoutHandler(router, time.Second * 5, "Timeout"),
		Addr: "127.0.0.1:8000",
		ReadTimeout: 5  * time.Second ,
		WriteTimeout: 5  * time.Second  ,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}


}
