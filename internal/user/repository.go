package user

import (
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type Repository interface {
	Create(user *User) error
	GetAll(filters Filterts)([]User , error)
	Get(id string) (*User, error)
	Delete(id string) error
	Update(id string, firstName *string, lastName *string, email *string, phone *string) error 
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

func (repo *repo) Create(user *User) error {
	repo.log.Println("User from repo")
	user.ID =  uuid.New().String()

	if err := repo.db.Create(user).Error; err != nil {
		repo.log.Println(err)
		return err
	}

	repo.log.Println("User has been create with id ", user.ID)

	return nil
}

func (repo *repo) GetAll(filters Filterts) ([]User , error)  {
	repo.log.Println("GetAll User from repo")

	var u []User

	tx := repo.db.Model(&u)
	tx = applyFilters(tx, filters)
	result := tx.Order("created_at desc").Find(&u)

	if result.Error != nil {
		return nil, result.Error
	}

	return u, nil
}

func (repo *repo) Get(id string) (*User , error)  {
	repo.log.Println("Get User by id from repo")

	user := User{ID: id}

	result := repo.db.First(&user)

	if result.Error != nil {
		return nil, result.Error
	}


	return &user, nil
}


func (repo *repo) Delete(id string) error {

	repo.log.Println("Get User by id from repo")
	user := User{ID:id}
	result := repo.db.Delete(&user)

		if result.Error != nil {
		return result.Error
	}

	return nil

}

func (r *repo) Update(id string, firstName *string, lastName *string, email *string, phone *string) error {
    
    values := make(map[string]interface{})

    if firstName != nil {
        values["first_name"] = *firstName
    }

    if lastName != nil {
        values["last_name"] = *lastName
    }

    if email != nil {
        values["email"] = *email
    }

    if phone != nil {
        values["phone"] = *phone
    }

    result := r.db.Model(&User{}).Where("id = ?", id).Updates(values)
    if result.Error != nil {
        return result.Error
    }

    return nil
}

func applyFilters(tx *gorm.DB, filters Filterts) *gorm.DB {

	if filters.FirstName != "" {
		filters.FirstName = fmt.Sprintf("%%%s%%", strings.ToLower(filters.FirstName))
		tx = tx.Where("lower(first_name) like ?", filters.FirstName)
	}

	if filters.LastName != "" {
		filters.LastName = fmt.Sprintf("%%%s%%", strings.ToLower(filters.LastName))
		tx = tx.Where("lower(last_name) like ?", filters.LastName)
	}

	return tx

}

