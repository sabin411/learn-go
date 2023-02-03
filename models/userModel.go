package models

import "gorm.io/gorm"

type Body struct{
	Email string `gorm:"unique"`
	Password string 
}

type User struct {
	gorm.Model
 	Email string `gorm:"unique"`
	Password string 

}