package lsp

import (
	"LSP_PNJ_NTG/entity"
	"fmt"
)

type Service interface {
	FindAll() ([]entity.LSP, error)
	FindById(ID int) (entity.LSP, error)
	Create(lsp LSPRequest) (entity.LSP, error)
	Delete(ID int) (entity.LSP, error)
	Update(ID int, lsp LSPRequest) (entity.LSP, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]entity.LSP, error) {
	lsp, err := s.repository.FindAll()
	return lsp, err
}

func (s *service) FindById(ID int) (entity.LSP, error) {
	lsp, err := s.repository.FindById(ID)
	return lsp, err
}

func (s *service) Create(lspRequest LSPRequest) (entity.LSP, error) {
	kodePos, _ := lspRequest.KodePos.Int64()
	jenisLSPID, _ := lspRequest.JenisLSPID.Int64()
	lsp := entity.LSP{
		Kode:               lspRequest.Kode,
		Nama:               lspRequest.Nama,
		NamaKetua:          lspRequest.NamaKetua,
		NamaDewanPengarah:  lspRequest.NamaDewanPengarah,
		NoTelepon:          lspRequest.NoTelepon,
		NoWhatsapp:         lspRequest.NoWhatsapp,
		Alamat:             lspRequest.Alamat,
		Provinsi:           lspRequest.Provinsi,
		Kota:               lspRequest.Kota,
		Kecamatan:          lspRequest.Kecamatan,
		Desa:               lspRequest.Desa,
		KodePos:            int(kodePos),
		Website:            lspRequest.Website,
		NoLisensi:          lspRequest.NoLisensi,
		MasaBerlakuLisensi: lspRequest.MasaBerlakuLisensi,
		InstitusiInduk:     lspRequest.InstitusiInduk,
		JenisLSPID:         int(jenisLSPID),
	}
	newlsp, err := s.repository.Create(lsp)
	return newlsp, err
}

func (s *service) Delete(ID int) (entity.LSP, error) {
	lsp, err := s.repository.FindById(ID)
	delLSP, err := s.repository.Delete(lsp)

	return delLSP, err
}

func (s *service) Update(ID int, lspRequest LSPRequest) (entity.LSP, error) {
	lsp, err := s.repository.FindById(ID)

	kodePos, _ := lspRequest.KodePos.Int64()
	jenisLSPID, _ := lspRequest.JenisLSPID.Int64()

	lsp.Kode = lspRequest.Kode
	lsp.Nama = lspRequest.Nama
	lsp.NamaKetua = lspRequest.NamaKetua
	lsp.NamaDewanPengarah = lspRequest.NamaDewanPengarah
	lsp.NoTelepon = lspRequest.NoTelepon
	lsp.NoWhatsapp = lspRequest.NoWhatsapp
	lsp.Alamat = lspRequest.Alamat
	lsp.Provinsi = lspRequest.Provinsi
	lsp.Kota = lspRequest.Kota
	lsp.Kecamatan = lspRequest.Kecamatan
	lsp.Desa = lspRequest.Desa
	lsp.KodePos = int(kodePos)
	lsp.Website = lspRequest.Website
	lsp.NoLisensi = lspRequest.NoLisensi
	lsp.MasaBerlakuLisensi = lspRequest.MasaBerlakuLisensi
	lsp.InstitusiInduk = lspRequest.InstitusiInduk
	lsp.JenisLSPID = int(jenisLSPID)

	updateLSP, err := s.repository.Update(lsp)
	fmt.Println(updateLSP)
	return updateLSP, err
}
