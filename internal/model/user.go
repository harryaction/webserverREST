package model

import "time"

type UserModel struct {
	Id        *string    `json:"id"`
	Name      *string    `json:"name"`
	Lastname  *string    `json:"lastname"`
	Age       int        `json:"age"`
	Birthdate *time.Time `json:"birthdate"`
}
