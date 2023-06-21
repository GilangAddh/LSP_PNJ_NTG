package lsp_sk

import (
	"LSP_PNJ_NTG/entity"
	"fmt"
)

type Service interface {
	FindAll() ([]entity.LSP_SK, error)
	FindById(ID int) (entity.LSP_SK, error)
	Create(sk LSPSKRequest) (entity.LSP_SK, error)
	Delete(ID int) (entity.LSP_SK, error)
	Update(ID int, sk LSPSKRequest) (entity.LSP_SK, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]entity.LSP_SK, error) {
	lsp_sk, err := s.repository.FindAll()
	return lsp_sk, err
}

func (s *service) FindById(ID int) (entity.LSP_SK, error) {
	lsp_sk, err := s.repository.FindById(ID)
	return lsp_sk, err
}

func (s *service) Create(LSPSKRequest LSPSKRequest) (entity.LSP_SK, error) {
	LSPID, _ := LSPSKRequest.LSPID.Int64()
	SKID, _ := LSPSKRequest.SKID.Int64()
	lsp_sk := entity.LSP_SK{
		LSPID: int(LSPID),
		SKID:  int(SKID),
	}
	lsp_sk, err := s.repository.Create(lsp_sk)
	return lsp_sk, err
}

func (s *service) Delete(ID int) (entity.LSP_SK, error) {
	lsp_sk, err := s.repository.FindById(ID)
	del_lsp_sk, err := s.repository.Delete(lsp_sk)

	return del_lsp_sk, err
}

func (s *service) Update(ID int, LSPSKRequest LSPSKRequest) (entity.LSP_SK, error) {
	sk, err := s.repository.FindById(ID)

	SKID, _ := LSPSKRequest.SKID.Int64()
	LSPID, _ := LSPSKRequest.LSPID.Int64()

	sk.SKID = int(SKID)
	sk.LSPID = int(LSPID)

	updatesk, err := s.repository.Update(sk)
	fmt.Println(updatesk)
	return updatesk, err
}
