package handler

import (
	"LSP_PNJ_NTG/account"
	"LSP_PNJ_NTG/entity"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type accountHandler struct {
	accountService account.Service
}

func NewAccountHandler(accountService account.Service) *accountHandler {
	return &accountHandler{accountService}
}

func (h *accountHandler) Registration(c *gin.Context) {
	CorsPolicy(c)
	var accountRequest account.AccountRequest

	err := c.ShouldBindJSON(&accountRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		}
	}

	account, err := h.accountService.Registration(accountRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "email already registered",
		})
		return
	}

	c.JSON(http.StatusOK, convertAToResponse(account))
}

func (h *accountHandler) Authentification(c *gin.Context) {
	CorsPolicy(c)
	var signinRequest account.SignInRequest

	err := c.ShouldBindJSON(&signinRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		}
	}

	account, err := h.accountService.Authentification(signinRequest.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "email not registered yet"})
		return
	}

	if account.Password != signinRequest.Password {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "incorrect password"})
		return
	}

	c.JSON(http.StatusOK, convertAToResponse(account))
}

func convertAToResponse(a entity.Accounts) account.AccountResponse {
	return account.AccountResponse{
		ID:       a.ID,
		Nama:     a.Nama,
		NIK:      a.NIK,
		Email:    a.Email,
		Password: a.Password,
		Role:     a.Role,
	}
}
