package models

import "gorm.io/gorm"

// user base model
type User struct {
	gorm.Model
	Name         string `json:"name" gorm:"not null; size:70"`
	Email        string `json:"email" gorm:"unique; not null; size:255"`
	Password     string `json:"password" gorm:"not null; size:255"`
	IsAdmin      bool   `json:"is_admin" gorm:"default:false"`
	Surname      string `json:"surname" gorm:"not null; size:70"`
	Phone_number string `json:"phone_number" gorm:"size:30"`
}

// AUTH
type RegisterUserReq struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Surname      string `json:"surname"`
	Phone_number string `json:"phone_number"`
} //register user req

type RegisterUserRes struct {
	gorm.Model
	Name         string
	Email        string
	Is_admin     bool
	Surname      string
	Phone_number string
} //register user response

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
} // login user request

type LoginUserRes struct {
	User
} // login user response

type VerifyTokenRes struct {
	ID       int    `json:"ID"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Is_admin bool   `json:"is_admin"`
	Surname  string `json:"surname"`
} // lo que responde el claim del toke
