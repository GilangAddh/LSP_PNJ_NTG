package handler

import (
	"LSP_PNJ_NTG/entity"
	"LSP_PNJ_NTG/sk"

	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type skHandler struct {
	skService sk.Service
}

func NewSKHandler(skService sk.Service) *skHandler {
	return &skHandler{skService}
}

func convertSToResponse(s entity.SK) sk.SKResponse {
	return sk.SKResponse{
		ID:                   s.ID,
		JudulStandar:         s.JudulStandar,
		NoStandar:            s.NoStandar,
		LegalitasPerundangan: s.LegalitasPerundangan,
		Sektor:               s.Sektor,
		SubSektor:            s.SubSektor,
		Penerbit:             s.Penerbit,
		JenisSKID:            s.JenisSKID,
	}
}

func (h *skHandler) GetSKs(c *gin.Context) {
	CorsPolicy(c)
	sks, err := h.skService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var sksResponse []sk.SKResponse

	for _, l := range sks {
		skResponse := convertSToResponse(l)
		sksResponse = append(sksResponse, skResponse)
	}

	c.JSON(http.StatusOK, sksResponse)
}

func (h *skHandler) GetSK(c *gin.Context) {
	CorsPolicy(c)
	id := c.Param("id")
	idb, err := strconv.Atoi(id)

	l, err := h.skService.FindById(idb)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	skResponse := convertSToResponse(l)

	c.JSON(http.StatusOK, skResponse)
}

func (h *skHandler) CreateSK(c *gin.Context) {
	CorsPolicy(c)
	var skRequest sk.SKRequest

	err := c.ShouldBindJSON(&skRequest)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	sk, err := h.skService.Create(skRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, convertSToResponse(sk))

}

func (h *skHandler) DeleteSK(c *gin.Context) {
	CorsPolicy(c)
	id := c.Param("id")
	idb, err := strconv.Atoi(id)

	l, err := h.skService.Delete(idb)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	skResponse := convertSToResponse(l)

	deleteMessage := "Sucessfull Deleting sk " + skResponse.JudulStandar

	c.JSON(http.StatusOK, gin.H{
		"data": deleteMessage,
	})
}

func (h *skHandler) UpdateSK(c *gin.Context) {
	CorsPolicy(c)
	var skUpdate sk.SKRequest

	err := c.ShouldBindJSON(&skUpdate)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	id := c.Param("id")
	idb, err := strconv.Atoi(id)
	sk, err := h.skService.Update(idb, skUpdate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, convertSToResponse(sk))
}
