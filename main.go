package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ElianDev55/first-api-go/internal/course"
	"github.com/ElianDev55/first-api-go/internal/enrollment"
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
	router.HandleFunc("/users", userEnd.Create).Methods("POST")
	router.HandleFunc("/users/{id}", userEnd.Get).Methods("GET")
	router.HandleFunc("/users/{id}", userEnd.Update).Methods("PATCH")
	router.HandleFunc("/users/{id}", userEnd.Delete).Methods("DELETE")

	courseRepo := course.NewRepo(l,db)
	courseService := course.NewService(l,courseRepo)
	courseEnd := course.MakeEndPoints(courseService)

	router.HandleFunc("/courses", courseEnd.Create).Methods("POST")
	router.HandleFunc("/courses", courseEnd.GetAll).Methods("GET")
	router.HandleFunc("/courses/{id}", courseEnd.Get).Methods("GET")
	router.HandleFunc("/courses/{id}", courseEnd.Update).Methods("PATCH")
	router.HandleFunc("/courses/{id}", courseEnd.Delete).Methods("DELETE")


	enrollmentRepo := enrollment.NewRepo(l,db)
	enrollmenService := enrollment.NewService(l,enrollmentRepo)
	enrollmenEnd := enrollment.MakeEndPoints(enrollmenService)

	router.HandleFunc("/enrollments", enrollmenEnd.Create).Methods("POST")
	router.HandleFunc("/enrollments", enrollmenEnd.GetAll).Methods("GET")


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
