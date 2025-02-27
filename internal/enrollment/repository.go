package enrollment

import (
	"log"

	"github.com/ElianDev55/first-api-go/internal/domain"
	"gorm.io/gorm"
)

type Repository interface {
	Create(enrollment *domain.Enrollment) error
	GetAll() ([]domain.Enrollment, error)
}


type repo struct {
	log *log.Logger
	db *gorm.DB
}

func NewRepo(log *log.Logger, db *gorm.DB) Repository {
	return &repo{
		log: log,
		db: db,
	}
}


func (repo *repo) Create(enrollment *domain.Enrollment) error {
	repo.log.Println("enrollment from repo")

	if err := repo.db.Create(enrollment).Error; err != nil {
		repo.log.Println(err)
		return err
	}

	repo.log.Println("domain.User has been create with id ", enrollment.ID)

	return nil
}


func (repo *repo) GetAll() ([]domain.Enrollment, error) {

	var e []domain.Enrollment

	result := repo.db.Model(&e).Order("created_at desc").Find(&e)

	if result.Error != nil {
		return nil, result.Error
	}

	return e,nil

}
