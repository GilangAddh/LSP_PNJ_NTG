package skema_sertifikasi

import (
	"LSP_PNJ_NTG/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.Skema_Sertifikasi, error)
	FindById(ID int) (entity.Skema_Sertifikasi, error)
	Create(skema entity.Skema_Sertifikasi) (entity.Skema_Sertifikasi, error)
	Delete(skema entity.Skema_Sertifikasi) (entity.Skema_Sertifikasi, error)
	Update(skema entity.Skema_Sertifikasi) (entity.Skema_Sertifikasi, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.Skema_Sertifikasi, error) {
	var skema []entity.Skema_Sertifikasi

	err := r.db.Find(&skema).Error

	return skema, err
}

func (r *repository) FindById(ID int) (entity.Skema_Sertifikasi, error) {
	var skema entity.Skema_Sertifikasi

	err := r.db.First(&skema, ID).Error

	return skema, err
}

func (r *repository) Create(skema entity.Skema_Sertifikasi) (entity.Skema_Sertifikasi, error) {
	err := r.db.Create(&skema).Error
	return skema, err
}

func (r *repository) Delete(skema entity.Skema_Sertifikasi) (entity.Skema_Sertifikasi, error) {
	err := r.db.Delete(&skema).Error

	return skema, err
}

func (r *repository) Update(skema entity.Skema_Sertifikasi) (entity.Skema_Sertifikasi, error) {

	err := r.db.Save(&skema).Error
	return skema, err
}
