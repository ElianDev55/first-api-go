package user

import (
	"encoding/json"
	"net/http"
)

type (

	Controller func (w http.ResponseWriter, r *http.Request)

	EndPoints struct {
		Create Controller
		GetAll Controller
		Get Controller
		Update Controller
		Delete Controller
	}

	CreateReq struct {

		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
		Email string `json:"emial"`
		Phone string `json:"phone"`
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

		error := s.Create(rq.FirstName, rq.LastName, rq.Email, rq.Phone)

		if error != nil {
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(ErrorResponse{
				Error: "first name is required",
			})
			return
		}

			json.NewEncoder(w).Encode(rq)
	}
}

func makeGetAllEnpoint(s Service) Controller {
	return func (w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
func makeGetEnpoint(s Service) Controller {
	return func (w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
func makeUpdateEnpoint(s Service) Controller {
	return func (w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func makeDeleteEnpoint(s Service) Controller {
	return func (w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
