package models

import (
	"consumerreviewsgo/db"
	"consumerreviewsgo/forms"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	UserName string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func (h User) Signup(userPayload forms.UserSignup) (*User, error) {
	id := uuid.NewV4()
	user := User{
		Id:       id,
		UserName: userPayload.UserName,
		Email:    userPayload.Email,
		Password: userPayload.Password,
	}
	model := ElasticSearchOperations{
		Client: db.GetDb(),
		CTX:    context.Background(),
		Index:  "user",
	}

	err := model.Insert(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}
