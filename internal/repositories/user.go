package repositories

import (
	"fmt"
	"log"
	"webserverREST/datasource"
	"webserverREST/internal/model"
)

type User interface {
	Put(u model.UserModel) error
	Delete(ID string) error
	Get() ([]model.UserModel, error)
}

type user struct{}

func NewUser() User {
	return new(user)
}

func (user) Put(u model.UserModel) error {
	userData := `INSERT INTO public.api_users(uuid, 
                             name, 
                             lastname, 
                             birthdate) 
                             VALUES($1, $2, $3, $4::timestamptz) ON CONFLICT(uuid) DO UPDATE SET name = $2, lastname = $3, birthdate = $4`
	_, err := datasource.Db.Exec(userData, u.Id, u.Name, u.Lastname, u.Birthdate)
	if err != nil {
		log.Printf("Error adding data: %v", err)
		return fmt.Errorf("error adding data: %v", err)
	}
	return nil
}

func (user) Delete(id string) error {
	_, err := datasource.Db.Exec(`DELETE FROM public.api_users WHERE uuid = $1`, id)
	if err != nil {
		log.Printf("Error deleting data: %v", err)
		return fmt.Errorf("error deleting data: %v", err)
	}
	return nil
}

func (user) Get() ([]model.UserModel, error) {
	var uu []model.UserModel
	err := datasource.Db.Select(&uu, `SELECT uuid id, name, lastname, birthdate, -(DATE_PART('year', birthdate::date) - DATE_PART('year', now())) age FROM public.api_users`)
	if err != nil {
		log.Printf("Error getting data: %v", err)
		return uu, fmt.Errorf("error getting data: %v", err)
	}
	return uu, nil
}
