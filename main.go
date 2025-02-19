package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ElianDev55/first-api-go/internal/user"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {


	
	router := mux.NewRouter()
	_ = godotenv.Load()
	

	host := os.Getenv("DB_HOST")
	userDb:= os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	nameDd := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, userDb, password, nameDd, port  )
	db, errDd := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db = db.Debug()
	_ = db.AutoMigrate(&user.User{})

	if errDd != nil {
		fmt.Println(errDd)
	}


	userService := user.NewService()
	userEnd := user.MakeEndPoints(userService)
	
	router.HandleFunc("/users", userEnd.Create).Methods("GET")
	router.HandleFunc("/users", userEnd.Create).Methods("POST")
	router.HandleFunc("/users", userEnd.Update).Methods("PATCH")
	router.HandleFunc("/users", userEnd.Delete).Methods("DELETE")


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
