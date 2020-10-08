package storage

import "github.com/AlexBaykov/go-rest/internal/app/model"

//UserRepo is an interface representing some way of storing users
type UserRepo interface {
	Create(*model.User) error
	Find(id int) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}
