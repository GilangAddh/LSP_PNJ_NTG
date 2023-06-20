package jenis_sk

import (
	"LSP_PNJ_NTG/entity"
	"fmt"
)

type Service interface {
	FindAll() ([]entity.JenisSK, error)
	FindById(ID int) (entity.JenisSK, error)
	Create(jenisSK JenisSKRequest) (entity.JenisSK, error)
	Delete(ID int) (entity.JenisSK, error)
	Update(ID int, jenisSK JenisSKRequest) (entity.JenisSK, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]entity.JenisSK, error) {
	jenisSK, err := s.repository.FindAll()
	return jenisSK, err
}

func (s *service) FindById(ID int) (entity.JenisSK, error) {
	jenisSK, err := s.repository.FindById(ID)
	return jenisSK, err
}

func (s *service) Create(jenisSKRequest JenisSKRequest) (entity.JenisSK, error) {
	jenisSK := entity.JenisSK{
		Nama: jenisSKRequest.Nama,
	}
	newjenisSK, err := s.repository.Create(jenisSK)
	return newjenisSK, err
}

func (s *service) Delete(ID int) (entity.JenisSK, error) {
	jenisSK, err := s.repository.FindById(ID)
	deljenisSK, err := s.repository.Delete(jenisSK)

	return deljenisSK, err
}

func (s *service) Update(ID int, jenisSKRequest JenisSKRequest) (entity.JenisSK, error) {
	jenisSK, err := s.repository.FindById(ID)

	jenisSK.Nama = jenisSKRequest.Nama

	updatejenisSK, err := s.repository.Update(jenisSK)
	fmt.Println(updatejenisSK)
	return updatejenisSK, err
}
