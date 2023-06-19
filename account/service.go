package account

import "LSP_PNJ_NTG/entity"

type Service interface {
	Registration(account AccountRequest) (entity.Accounts, error)
	Authentification(email string) (entity.Accounts, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Registration(accountRequest AccountRequest) (entity.Accounts, error) {
	nik, _ := accountRequest.NIK.Int64()
	account := entity.Accounts{
		Nama:     accountRequest.Nama,
		NIK:      int(nik),
		Email:    accountRequest.Email,
		Password: accountRequest.Password,
		Role:     accountRequest.Role,
	}

	newAccount, err := s.repository.Registration(account)
	return newAccount, err
}

func (s *service) Authentification(email string) (entity.Accounts, error) {
	result, err := s.repository.Authentification(email)
	return result, err
}
