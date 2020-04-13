package repositories

import (
	"github.com/jmoiron/sqlx"
	"log"
	"webserverREST/internal/model"
	"webserverREST/internal/tools"
)

type User interface {
	Put(u *model.UserModel) error
	Delete(ID string) error
	Get() ([]model.UserModel, error)
}

type user struct {
	DB *sqlx.DB
}

func NewUser(db *sqlx.DB) User {
	return user{db}
}

func (a user) Put(u *model.UserModel) error {
	userData := `INSERT INTO public.api_users(uuid, 
                             name, 
                             lastname, 
                             birthdate) 
                             VALUES($1, $2, $3, $4::timestamptz) ON CONFLICT(uuid) DO UPDATE SET name = $2, lastname = $3, birthdate = $4`
	_, err := a.DB.Exec(userData, u.Id, u.Name, u.Lastname, u.Birthdate)
	if err != nil {
		log.Printf("Error adding data: %v", err)
		return err
	}
	return nil
}

func (a user) Delete(id string) error {
	_, err := a.DB.Exec(`DELETE FROM public.api_users WHERE uuid = $1`, id)
	if err != nil {
		log.Printf("Error deleting data: %v", err)
		return err
	}
	return nil
}

func (a user) Get() ([]model.UserModel, error) {
	var uu []model.UserModel
	err := a.DB.Select(&uu, `SELECT uuid id, name, lastname, birthdate, 0 age FROM public.api_users`)
	if err != nil {
		log.Printf("Error getting data: %v", err)
		return uu, err
	}
	for i, b := range uu {
		uu[i].Age = tools.Age(*b.Birthdate)
	}
	return uu, nil
}
