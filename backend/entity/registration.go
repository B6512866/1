package entity

import (
	"gorm.io/gorm"
	"github.com/asaskevich/govalidator"
)

type Registration struct {
	gorm.Model
	Name string `valid:"required~Name Is Required"`
	Email string `valid:"email~Required Type Email"`
	StudentID string `valid:"matches(^B[0-9]{7}$)~Student must be B + 7 digits"`
	Age int `valid:"range(18|60)~Age must be 18-60"`
}

func (r *Registration) Validate() (bool,error) {
	return govalidator.ValidateStruct(r)
}