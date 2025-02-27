package enrollment

import (
	"log"

	"github.com/ElianDev55/first-api-go/internal/domain"
)


type service struct {
	log *log.Logger
	repo Repository
}


type Service interface {
	Create(userID, courseID string) (*domain.Enrollment, error)
	GetAll()([]domain.Enrollment , error)
}


func NewService(log *log.Logger, repo Repository) Service{
	return &service{
		repo: repo,
		log: log,
	}
}


func (s service) Create(userID, courseID string) (*domain.Enrollment, error) {

	enroll := &domain.Enrollment{
		UserID: userID ,
		CourseID: courseID,
		Status: "P",
	}

	err := s.repo.Create(enroll)

	if err != nil {
		return nil,err
	}

	return enroll, nil

}

func (s service) GetAll() ([]domain.Enrollment, error) {

	enrollments, err := s.repo.GetAll()

	if err != nil {
		return nil,err
	}

	return enrollments,nil

}
