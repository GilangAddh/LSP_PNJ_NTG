package jenis_lsp

import (
	"LSP_PNJ_NTG/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.JenisLSP, error)
	FindById(ID int) (entity.JenisLSP, error)
	Create(jenisLSP entity.JenisLSP) (entity.JenisLSP, error)
	Delete(jenisLSP entity.JenisLSP) (entity.JenisLSP, error)
	Update(jenisLSP entity.JenisLSP) (entity.JenisLSP, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.JenisLSP, error) {
	var jenisLSP []entity.JenisLSP

	err := r.db.Find(&jenisLSP).Error

	return jenisLSP, err
}

func (r *repository) FindById(ID int) (entity.JenisLSP, error) {
	var jenisLSP entity.JenisLSP

	err := r.db.First(&jenisLSP, ID).Error

	return jenisLSP, err
}

func (r *repository) Create(jenisLSP entity.JenisLSP) (entity.JenisLSP, error) {
	err := r.db.Create(&jenisLSP).Error
	return jenisLSP, err
}

func (r *repository) Delete(jenisLSP entity.JenisLSP) (entity.JenisLSP, error) {
	err := r.db.Delete(&jenisLSP).Error

	return jenisLSP, err
}

func (r *repository) Update(jenisLSP entity.JenisLSP) (entity.JenisLSP, error) {

	err := r.db.Save(&jenisLSP).Error
	return jenisLSP, err
}
