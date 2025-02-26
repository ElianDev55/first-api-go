package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)




type Course struct {

	ID 							string	  		`json:"id" gorm:"type:char(36);not null;primary_key;unique"`
	Name 						string 				`json:"name" gorm:"type:varchar(50);not null"`
	StartDate 			time.Time  		`json:"start_date"`
	EndDate  				time.Time 		`json:"end_date"`
	CreatedAt 			*time.Time 		`json:"-"` 
	UpdatedAt  			*time.Time 		`json:"-"` 
	Deleted 				*time.Time 		`json:"-"` 

}



func (u *Course) BeforeCreate(tx *gorm.DB) (err error) {
    if u.ID == "" {
        u.ID = uuid.New().String()
    }
    return
}
