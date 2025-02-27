package enrollment

import (
	"encoding/json"
	"net/http"

	"github.com/ElianDev55/first-api-go/pkg/meta"
)

type (

	Controller func (w http.ResponseWriter, r *http.Request)

	EndPoints struct {
		Create 		Controller
		GetAll 		Controller
	}

	CreateReq struct {
		UserId 	string `json:"user_id"`
		CourseId 		string `json:"course_id"`
	}

	Response struct {
		Status	 int   					`json:"status"` 
		Data 		interface{} 		`json:"data,omitempty"`
		Err 		string					`json:"error,omitempty"`
		Meta 		*meta.Meta			`json:"meta,omitempty"`
	}

)


func MakeEndPoints(s Service) EndPoints  {
	return EndPoints{
		Create: makeCreateEnpoint(s),
		GetAll: makeGetAllEnpoint(s),
	}
}


func makeCreateEnpoint(s Service) Controller {
	return func (w http.ResponseWriter, r *http.Request) {

		var rq CreateReq

		err := json.NewDecoder(r.Body).Decode(&rq)

		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{
				Status: 400,
				Err: err.Error(),
			})
			return
		}

		enrollment, errEnroll := s.Create(rq.UserId,rq.CourseId)

		if errEnroll != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{
				Status: 400,
				Err: err.Error(),
			})
			return
		}

		json.NewEncoder(w).Encode(&Response{
				Status: 200,
				Data: enrollment,
			})

	}
}


func makeGetAllEnpoint(s Service) Controller {
	return func (w http.ResponseWriter, r *http.Request) {

	
		users, err := s.GetAll()
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
