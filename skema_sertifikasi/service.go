package skema_sertifikasi

import (
	"LSP_PNJ_NTG/entity"
	"fmt"
)

type Service interface {
	FindAll() ([]entity.Skema_Sertifikasi, error)
	FindById(ID int) (entity.Skema_Sertifikasi, error)
	Create(skema SkemaSertifikasiRequest) (entity.Skema_Sertifikasi, error)
	Delete(ID int) (entity.Skema_Sertifikasi, error)
	Update(ID int, skema SkemaSertifikasiRequest) (entity.Skema_Sertifikasi, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]entity.Skema_Sertifikasi, error) {
	skema, err := s.repository.FindAll()
	return skema, err
}

func (s *service) FindById(ID int) (entity.Skema_Sertifikasi, error) {
	skema, err := s.repository.FindById(ID)
	return skema, err
}

func (s *service) Create(skemaRequest SkemaSertifikasiRequest) (entity.Skema_Sertifikasi, error) {
	SKID, _ := skemaRequest.SKID.Int64()
	skema := entity.Skema_Sertifikasi{
		Judul:           skemaRequest.Judul,
		JudulInggris:    skemaRequest.JudulInggris,
		Kode:            skemaRequest.Kode,
		KeteranganBukti: skemaRequest.KeteranganBukti,
		KedalamanBukti:  skemaRequest.KedalamanBukti,
		SKID:            int(SKID),
	}
	newskema, err := s.repository.Create(skema)
	return newskema, err
}

func (s *service) Delete(ID int) (entity.Skema_Sertifikasi, error) {
	sk, err := s.repository.FindById(ID)
	delsk, err := s.repository.Delete(sk)

	return delsk, err
}

func (s *service) Update(ID int, skemaRequest SkemaSertifikasiRequest) (entity.Skema_Sertifikasi, error) {
	skema, err := s.repository.FindById(ID)

	SKID, _ := skemaRequest.SKID.Int64()

	skema.Judul = skemaRequest.Judul
	skema.JudulInggris = skemaRequest.JudulInggris
	skema.Kode = skemaRequest.Kode
	skema.KeteranganBukti = skemaRequest.KeteranganBukti
	skema.KedalamanBukti = skemaRequest.KedalamanBukti
	skema.SKID = int(SKID)

	updateskema, err := s.repository.Update(skema)
	fmt.Println(updateskema)
	return updateskema, err
}
