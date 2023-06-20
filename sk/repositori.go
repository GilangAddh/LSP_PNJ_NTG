package sk

import (
	"LSP_PNJ_NTG/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.SK, error)
	FindById(ID int) (entity.SK, error)
	Create(sk entity.SK) (entity.SK, error)
	Delete(sk entity.SK) (entity.SK, error)
	Update(sk entity.SK) (entity.SK, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.SK, error) {
	var sk []entity.SK

	err := r.db.Find(&sk).Error

	return sk, err
}

func (r *repository) FindById(ID int) (entity.SK, error) {
	var sk entity.SK

	err := r.db.First(&sk, ID).Error

	return sk, err
}

func (r *repository) Create(sk entity.SK) (entity.SK, error) {
	err := r.db.Create(&sk).Error
	return sk, err
}

func (r *repository) Delete(sk entity.SK) (entity.SK, error) {
	err := r.db.Delete(&sk).Error

	return sk, err
}

func (r *repository) Update(sk entity.SK) (entity.SK, error) {

	err := r.db.Save(&sk).Error
	return sk, err
}
