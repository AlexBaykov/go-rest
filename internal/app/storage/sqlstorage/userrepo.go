package sqlstorage

import (
	"github.com/AlexBaykov/go-rest-api/internal/app/model"
	"github.com/AlexBaykov/go-rest-api/internal/app/storage"
)

type UserRepo struct {
	storage storage.Storage
}

func (r *UserRepo) Create(*model.User) error {

}
