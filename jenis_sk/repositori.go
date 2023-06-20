package jenis_sk

import (
	"LSP_PNJ_NTG/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.JenisSK, error)
	FindById(ID int) (entity.JenisSK, error)
	Create(jenisSK entity.JenisSK) (entity.JenisSK, error)
	Delete(jenisSK entity.JenisSK) (entity.JenisSK, error)
	Update(jenisSK entity.JenisSK) (entity.JenisSK, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.JenisSK, error) {
	var jenisSK []entity.JenisSK

	err := r.db.Find(&jenisSK).Error

	return jenisSK, err
}

func (r *repository) FindById(ID int) (entity.JenisSK, error) {
	var jenisSK entity.JenisSK

	err := r.db.First(&jenisSK, ID).Error

	return jenisSK, err
}

func (r *repository) Create(jenisSK entity.JenisSK) (entity.JenisSK, error) {
	err := r.db.Create(&jenisSK).Error
	return jenisSK, err
}

func (r *repository) Delete(jenisSK entity.JenisSK) (entity.JenisSK, error) {
	err := r.db.Delete(&jenisSK).Error

	return jenisSK, err
}

func (r *repository) Update(jenisSK entity.JenisSK) (entity.JenisSK, error) {

	err := r.db.Save(&jenisSK).Error
	return jenisSK, err
}
