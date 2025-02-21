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

	Response struct {
		Status	 int   					`json:"status"` 
		Data 		interface{} 		`json:"data,omitempty"`
		Err 		string					`json:"error,omitempty"`
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
			json.NewEncoder(w).Encode(&Response{
				Status: 400,
				Err: err.Error(),
			})
			return
		}

		if rq.FirstName == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{
				Status: 400,
				Err: "first name is required",
			})
			return
		}

		user, error := s.Create(rq.FirstName, rq.LastName, rq.Email, rq.Phone)

		if error != nil {
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(&Response{
				Status: 400,
				Err: error.Error(),
			})
			return
		}

			json.NewEncoder(w).Encode(&Response{
				Status: 200,
				Data: user,
			})
	}
}

func makeGetAllEnpoint(s Service) Controller {
	return func (w http.ResponseWriter, r *http.Request) {

		q := r.URL.Query()

		filters := Filterts{
			FirstName: q.Get("first_name"),
			LastName:  q.Get("last_name"),
		}

		users, err := s.GetAll(filters)
		if err != nil {
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(&Response{
				Status: 400,
				Err: err.Error(),
			})
			return
		}

		
			json.NewEncoder(w).Encode(&Response{
				Status: 200,
				Data: users,
			})
	}
}
func makeGetEnpoint(s Service) Controller {
	return func (w http.ResponseWriter, r *http.Request) {

		pah := mux.Vars(r)
		id := pah["id"]
		user, err := s.Get(id)

		if err != nil {
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(&Response{
				Status: 400,
				Err: err.Error(),
			})
			return
		}

			json.NewEncoder(w).Encode(&Response{
				Status: 200,
				Data: user,
			})
	}
}
func makeUpdateEnpoint(s Service) Controller {
	return func (w http.ResponseWriter, r *http.Request) {

		var rq UpdateReq

		err := json.NewDecoder(r.Body).Decode(&rq);

		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{
				Status: 400,
				Err: "Invalided request",
			})
			return
		}

		if rq.FirstName != nil && *rq.FirstName == "" {
    w.WriteHeader(400)
    json.NewEncoder(w).Encode(&Response{
				Status: 400,
				Err: "first name is required",
			})
    return
	}

	path := mux.Vars(r)
	id := path["id"]

	erro := s.Update(id, rq.FirstName, rq.LastName, rq.Email, rq.Phone)

		if erro != nil {
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(&Response{
				Status: 400,
				Err: erro.Error(),
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
				json.NewEncoder(w).Encode(&Response{
				Status: 400,
				Err:  err.Error(),
			})
			return
		}


			json.NewEncoder(w).Encode(map[string]bool{"Delete": true })
	}
}
