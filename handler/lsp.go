package handler

import (
	"LSP_PNJ_NTG/entity"
	"LSP_PNJ_NTG/lsp"

	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type lspHandler struct {
	lspService lsp.Service
}

func NewLSPHandler(lspService lsp.Service) *lspHandler {
	return &lspHandler{lspService}
}

func convertLToResponse(l entity.LSP) lsp.LSPResponse {
	return lsp.LSPResponse{
		ID:                 l.ID,
		Kode:               l.Kode,
		Nama:               l.Nama,
		NamaKetua:          l.NamaKetua,
		NamaDewanPengarah:  l.NamaDewanPengarah,
		NoTelepon:          l.NoTelepon,
		NoWhatsapp:         l.NoWhatsapp,
		Alamat:             l.Alamat,
		Provinsi:           l.Provinsi,
		Kota:               l.Kota,
		Kecamatan:          l.Kecamatan,
		Desa:               l.Desa,
		KodePos:            l.KodePos,
		Website:            l.Website,
		NoLisensi:          l.NoLisensi,
		MasaBerlakuLisensi: l.MasaBerlakuLisensi,
		InstitusiInduk:     l.InstitusiInduk,
		JenisLSPID:         l.JenisLSPID,
	}
}

func (h *lspHandler) GetLSPs(c *gin.Context) {
	CorsPolicy(c)
	lsps, err := h.lspService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var LSPsResponse []lsp.LSPResponse

	for _, l := range lsps {
		LSPResponse := convertLToResponse(l)
		LSPsResponse = append(LSPsResponse, LSPResponse)
	}

	c.JSON(http.StatusOK, LSPsResponse)
}

func (h *lspHandler) GetLSP(c *gin.Context) {
	CorsPolicy(c)
	id := c.Param("id")
	idb, err := strconv.Atoi(id)

	l, err := h.lspService.FindById(idb)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	lspResponse := convertLToResponse(l)

	c.JSON(http.StatusOK, lspResponse)
}

func (h *lspHandler) CreateLSP(c *gin.Context) {
	CorsPolicy(c)
	var lspRequest lsp.LSPRequest

	err := c.ShouldBindJSON(&lspRequest)
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

	lsp, err := h.lspService.Create(lspRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, convertLToResponse(lsp))

}

func (h *lspHandler) DeleteLSP(c *gin.Context) {
	CorsPolicy(c)
	id := c.Param("id")
	idb, err := strconv.Atoi(id)

	l, err := h.lspService.Delete(idb)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	lspResponse := convertLToResponse(l)

	deleteMessage := "Sucessfull Deleting LSP " + lspResponse.Nama

	c.JSON(http.StatusOK, gin.H{
		"data": deleteMessage,
	})
}

func (h *lspHandler) UpdateLSP(c *gin.Context) {
	CorsPolicy(c)
	var lspUpdate lsp.LSPRequest

	err := c.ShouldBindJSON(&lspUpdate)
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
	lsp, err := h.lspService.Update(idb, lspUpdate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, convertLToResponse(lsp))
}
