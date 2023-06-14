package jenis_lsp

import (
	"LSP_PNJ_NTG/entity"
	"fmt"
)

type Service interface {
	FindAll() ([]entity.JenisLSP, error)
	FindById(ID int) (entity.JenisLSP, error)
	Create(jenisLSP JenisLSPRequest) (entity.JenisLSP, error)
	Delete(ID int) (entity.JenisLSP, error)
	Update(ID int, jenisLSP JenisLSPRequest) (entity.JenisLSP, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]entity.JenisLSP, error) {
	jenisLSP, err := s.repository.FindAll()
	return jenisLSP, err
}

func (s *service) FindById(ID int) (entity.JenisLSP, error) {
	jenisLSP, err := s.repository.FindById(ID)
	return jenisLSP, err
}

func (s *service) Create(jenisLSPRequest JenisLSPRequest) (entity.JenisLSP, error) {
	jenisLSP := entity.JenisLSP{
		Nama: jenisLSPRequest.Nama,
	}
	newJenisLSP, err := s.repository.Create(jenisLSP)
	return newJenisLSP, err
}

func (s *service) Delete(ID int) (entity.JenisLSP, error) {
	jenisLSP, err := s.repository.FindById(ID)
	delJenisLSP, err := s.repository.Delete(jenisLSP)

	return delJenisLSP, err
}

func (s *service) Update(ID int, jenisLSPRequest JenisLSPRequest) (entity.JenisLSP, error) {
	jenisLSP, err := s.repository.FindById(ID)

	jenisLSP.Nama = jenisLSPRequest.Nama

	updateJenisLSP, err := s.repository.Update(jenisLSP)
	fmt.Println(updateJenisLSP)
	return updateJenisLSP, err
}
