package sk

import (
	"LSP_PNJ_NTG/entity"
	"fmt"
)

type Service interface {
	FindAll() ([]entity.SK, error)
	FindById(ID int) (entity.SK, error)
	Create(sk SKRequest) (entity.SK, error)
	Delete(ID int) (entity.SK, error)
	Update(ID int, sk SKRequest) (entity.SK, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]entity.SK, error) {
	sk, err := s.repository.FindAll()
	return sk, err
}

func (s *service) FindById(ID int) (entity.SK, error) {
	sk, err := s.repository.FindById(ID)
	return sk, err
}

func (s *service) Create(skRequest SKRequest) (entity.SK, error) {
	jenisSKID, _ := skRequest.JenisSKID.Int64()
	sk := entity.SK{
		JudulStandar:         skRequest.JudulStandar,
		NoStandar:            skRequest.NoStandar,
		LegalitasPerundangan: skRequest.LegalitasPerundangan,
		Sektor:               skRequest.Sektor,
		SubSektor:            skRequest.SubSektor,
		Penerbit:             skRequest.Penerbit,
		JenisSKID:            int(jenisSKID),
	}
	newsk, err := s.repository.Create(sk)
	return newsk, err
}

func (s *service) Delete(ID int) (entity.SK, error) {
	sk, err := s.repository.FindById(ID)
	delsk, err := s.repository.Delete(sk)

	return delsk, err
}

func (s *service) Update(ID int, skRequest SKRequest) (entity.SK, error) {
	sk, err := s.repository.FindById(ID)

	jenisSKID, _ := skRequest.JenisSKID.Int64()

	sk.JudulStandar = skRequest.JudulStandar
	sk.NoStandar = skRequest.NoStandar
	sk.LegalitasPerundangan = skRequest.LegalitasPerundangan
	sk.Sektor = skRequest.Sektor
	sk.SubSektor = skRequest.SubSektor
	sk.Penerbit = skRequest.Penerbit
	sk.JenisSKID = int(jenisSKID)

	updatesk, err := s.repository.Update(sk)
	fmt.Println(updatesk)
	return updatesk, err
}
