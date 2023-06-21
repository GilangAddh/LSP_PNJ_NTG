package lsp_sk

import (
	"LSP_PNJ_NTG/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.LSP_SK, error)
	FindById(ID int) (entity.LSP_SK, error)
	Create(sk entity.LSP_SK) (entity.LSP_SK, error)
	Delete(sk entity.LSP_SK) (entity.LSP_SK, error)
	Update(sk entity.LSP_SK) (entity.LSP_SK, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.LSP_SK, error) {
	var sk []entity.LSP_SK

	err := r.db.Find(&sk).Error

	return sk, err
}

func (r *repository) FindById(ID int) (entity.LSP_SK, error) {
	var sk entity.LSP_SK

	err := r.db.First(&sk, ID).Error

	return sk, err
}

func (r *repository) Create(sk entity.LSP_SK) (entity.LSP_SK, error) {
	err := r.db.Create(&sk).Error
	return sk, err
}

func (r *repository) Delete(sk entity.LSP_SK) (entity.LSP_SK, error) {
	err := r.db.Delete(&sk).Error

	return sk, err
}

func (r *repository) Update(sk entity.LSP_SK) (entity.LSP_SK, error) {

	err := r.db.Save(&sk).Error
	return sk, err
}
