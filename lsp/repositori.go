package lsp

import (
	"LSP_PNJ_NTG/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.LSP, error)
	FindById(ID int) (entity.LSP, error)
	Create(lsp entity.LSP) (entity.LSP, error)
	Delete(lsp entity.LSP) (entity.LSP, error)
	Update(lsp entity.LSP) (entity.LSP, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.LSP, error) {
	var lsp []entity.LSP

	err := r.db.Find(&lsp).Error

	return lsp, err
}

func (r *repository) FindById(ID int) (entity.LSP, error) {
	var lsp entity.LSP

	err := r.db.First(&lsp, ID).Error

	return lsp, err
}

func (r *repository) Create(lsp entity.LSP) (entity.LSP, error) {
	err := r.db.Create(&lsp).Error
	return lsp, err
}

func (r *repository) Delete(lsp entity.LSP) (entity.LSP, error) {
	err := r.db.Delete(&lsp).Error

	return lsp, err
}

func (r *repository) Update(lsp entity.LSP) (entity.LSP, error) {

	err := r.db.Save(&lsp).Error
	return lsp, err
}
