package handler

import (
	"LSP_PNJ_NTG/entity"
	"LSP_PNJ_NTG/jenis_lsp"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type jenisLSPHandler struct {
	jenisLSPService jenis_lsp.Service
}

func NewJenisLSPHandler(jenisLSPService jenis_lsp.Service) *jenisLSPHandler {
	return &jenisLSPHandler{jenisLSPService}
}

func convertJToResponse(j entity.JenisLSP) jenis_lsp.JenisLSPResponse {
	return jenis_lsp.JenisLSPResponse{
		ID:   j.ID,
		Nama: j.Nama,
	}
}

func (h *jenisLSPHandler) GetJenisLSPs(c *gin.Context) {
	CorsPolicy(c)
	jenisLSPs, err := h.jenisLSPService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var JenisLSPsResponse []jenis_lsp.JenisLSPResponse

	for _, j := range jenisLSPs {
		JenisLSPResponse := convertJToResponse(j)
		JenisLSPsResponse = append(JenisLSPsResponse, JenisLSPResponse)
	}

	c.JSON(http.StatusOK, JenisLSPsResponse)
}

func (h *jenisLSPHandler) GetJenisLSP(c *gin.Context) {
	CorsPolicy(c)
	id := c.Param("id")
	idb, err := strconv.Atoi(id)

	j, err := h.jenisLSPService.FindById(idb)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	jenisLSPResponse := convertJToResponse(j)

	c.JSON(http.StatusOK, jenisLSPResponse)
}

func (h *jenisLSPHandler) DeleteJenisLSP(c *gin.Context) {
	CorsPolicy(c)
	id := c.Param("id")
	idb, err := strconv.Atoi(id)

	j, err := h.jenisLSPService.Delete(idb)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	lspResponse := convertJToResponse(j)

	deleteMessage := "Sucessfull Deleting Jenis LSP " + lspResponse.Nama

	c.JSON(http.StatusOK, gin.H{
		"data": deleteMessage,
	})
}

func (h *jenisLSPHandler) CreateJenisLSP(c *gin.Context) {
	CorsPolicy(c)
	var jenisLSPRequest jenis_lsp.JenisLSPRequest

	err := c.ShouldBindJSON(&jenisLSPRequest)
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

	jenisLSP, err := h.jenisLSPService.Create(jenisLSPRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, convertJToResponse(jenisLSP))

}

func (h *jenisLSPHandler) UpdateJenisLSP(c *gin.Context) {
	CorsPolicy(c)
	var jenisLSPUpdate jenis_lsp.JenisLSPRequest

	err := c.ShouldBindJSON(&jenisLSPUpdate)
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
	jenisLSP, err := h.jenisLSPService.Update(idb, jenisLSPUpdate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, convertJToResponse(jenisLSP))
}
