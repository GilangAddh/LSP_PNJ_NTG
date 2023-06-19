package account

import (
	"LSP_PNJ_NTG/entity"
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Registration(a entity.Accounts) (entity.Accounts, error)
	Authentification(email string) (entity.Accounts, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Registration(a entity.Accounts) (entity.Accounts, error) {
	var existingAccount entity.Accounts

	result := r.db.Where("email = ?", a.Email).First(&existingAccount)
	err := errors.New("error")

	if result.RowsAffected == 0 {
		err = r.db.Create(&a).Error
	}

	return a, err
}

func (r *repository) Authentification(email string) (entity.Accounts, error) {
	var account entity.Accounts

	err := r.db.Where("email = ?", email).First(&account).Error

	return account, err
}
