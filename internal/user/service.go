package user

import "log"


type (
	service struct {
	log *log.Logger
	repo Repository
}

	Service interface {
	Create(firstName, lastName, email, phone string) (*User, error)
	GetAll(filters Filterts)([]User, error)
	Get(id string)(*User, error)
	Update(id string, firstName *string, lastName *string, email *string, phone *string) error
	Delete(id string) error
}
Filterts struct {
	FirstName string
	LastName string
}

) 


func NewService(log *log.Logger, repo Repository) Service{
	return &service{
		repo: repo,
		log: log,
	}
}

func (s service) Create(firstName, lastName, email, phone string) (*User, error)  {
	s.log.Println("Create user service")
	user := User{
		FirstName: firstName,
		LastName: lastName,
		Email: email,
		Phone: phone,
	}
	

	if err := s.repo.Create(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s service) GetAll(filters Filterts) ([]User, error)  {
	s.log.Println("GetAll user service")
	
	users,err := s.repo.GetAll(filters)

	if err != nil {
		return nil, err
	}

	return users, nil
}


func (s service) Get(id string) (*User, error)  {
	s.log.Println("Get user service")
	
	user,err := s.repo.Get(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s service) Update(id string, firstName *string, lastName *string, email *string, phone *string) error  {
	s.log.Println("Update user service")
	
	err := s.repo.Update(id, firstName, lastName, email, phone)

	if err != nil {
		return err
	}

	return nil
}


func (s service) Delete(id string) error  {
	s.log.Println("Delete user service")
	
	err := s.repo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
