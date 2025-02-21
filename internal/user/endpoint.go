package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type (

	Controller func (w http.ResponseWriter, r *http.Request)

	EndPoints struct {
		Create 		Controller
		GetAll 		Controller
		Get 			Controller
		Update 		Controller
		Delete 		Controller
	}

	CreateReq struct {
		FirstName 	string `json:"first_name"`
		LastName 		string `json:"last_name"`
		Email 			string `json:"emial"`
		Phone 			string `json:"phone"`
	}

	UpdateReq struct {
		FirstName 	*string `json:"first_name"`
		LastName 		*string `json:"last_name"`
		Email 			*string `json:"emial"`
		Phone 			*string `json:"phone"`
	}
	ErrorResponse struct {
		Error string `json:"error"`
	}

)

func MakeEndPoints(s Service) EndPoints  {
	return EndPoints{
		Create: makeCreateEnpoint(s),
		GetAll: makeGetAllEnpoint(s),
		Get: makeGetEnpoint(s),
		Update: makeUpdateEnpoint(s),
		Delete: makeDeleteEnpoint(s),
	}
}


func makeCreateEnpoint(s Service) Controller {
	return func (w http.ResponseWriter, r *http.Request) {

		var rq CreateReq

		err := json.NewDecoder(r.Body).Decode(&rq);

		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorResponse{
				Error: "Invalid request format",
			})
			return
		}

		if rq.FirstName == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorResponse{
				Error: "first name is required",
			})
			return
		}

		user, error := s.Create(rq.FirstName, rq.LastName, rq.Email, rq.Phone)

		if error != nil {
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(ErrorResponse{
				Error: error.Error(),
			})
			return
		}

			json.NewEncoder(w).Encode(user)
	}
}

func makeGetAllEnpoint(s Service) Controller {
	return func (w http.ResponseWriter, r *http.Request) {

		users, err := s.GetAll()

		if err != nil {
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(ErrorResponse{
				Error: err.Error(),
			})
			return
		}

			json.NewEncoder(w).Encode(users)
	}
}
func makeGetEnpoint(s Service) Controller {
	return func (w http.ResponseWriter, r *http.Request) {

		pah := mux.Vars(r)
		id := pah["id"]
		user, err := s.Get(id)

		if err != nil {
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(ErrorResponse{
				Error: err.Error(),
			})
			return
		}

			json.NewEncoder(w).Encode(user)
	}
}
func makeUpdateEnpoint(s Service) Controller {
	return func (w http.ResponseWriter, r *http.Request) {

		var rq UpdateReq

		err := json.NewDecoder(r.Body).Decode(&rq);

		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorResponse{
				Error: "Invalid request format",
			})
			return
		}

		if rq.FirstName != nil && *rq.FirstName == "" {
    w.WriteHeader(400)
    json.NewEncoder(w).Encode(ErrorResponse{"first name is required"})
    return
	}

	path := mux.Vars(r)
	id := path["id"]

	erro := s.Update(id, rq.FirstName, rq.LastName, rq.Email, rq.Phone)

		if erro != nil {
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(ErrorResponse{
				Error: err.Error(),
			})
			return
		}

			json.NewEncoder(w).Encode(map[string]bool{"Uodate": true})
	}
}

func makeDeleteEnpoint(s Service) Controller {
	return func (w http.ResponseWriter, r *http.Request) {

		path := mux.Vars(r)
		id := path["id"]

		err := s.Delete(id)

		if err != nil {
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(ErrorResponse{
				Error: err.Error(),
			})
			return
		}


			json.NewEncoder(w).Encode(map[string]bool{"Delete": true })
	}
}
