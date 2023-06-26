package handler

import (
	"LSP_PNJ_NTG/entity"
	"LSP_PNJ_NTG/skema_sertifikasi"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type skemaHandler struct {
	skemaService skema_sertifikasi.Service
}

func NewskemaHandler(skemaService skema_sertifikasi.Service) *skemaHandler {
	return &skemaHandler{skemaService}
}

func convertSSToResponse(ss entity.Skema_Sertifikasi) skema_sertifikasi.SkemaSertifikasiResponse {
	return skema_sertifikasi.SkemaSertifikasiResponse{
		ID:              ss.ID,
		Judul:           ss.Judul,
		JudulInggris:    ss.JudulInggris,
		Kode:            ss.Kode,
		KeteranganBukti: ss.KeteranganBukti,
		KedalamanBukti:  ss.KedalamanBukti,
		SKID:            ss.SKID,
	}
}

func (h *skemaHandler) GetSkemas(c *gin.Context) {
	CorsPolicy(c)
	skema_sertifikasis, err := h.skemaService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var skemasResponse []skema_sertifikasi.SkemaSertifikasiResponse

	for _, ss := range skema_sertifikasis {
		skemaResponse := convertSSToResponse(ss)
		skemasResponse = append(skemasResponse, skemaResponse)
	}

	c.JSON(http.StatusOK, skemasResponse)
}

func (h *skemaHandler) GetSkema(c *gin.Context) {
	CorsPolicy(c)
	id := c.Param("id")
	idb, err := strconv.Atoi(id)

	ss, err := h.skemaService.FindById(idb)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	skemaResponse := convertSSToResponse(ss)

	c.JSON(http.StatusOK, skemaResponse)
}

func (h *skemaHandler) DeleteSkema(c *gin.Context) {
	CorsPolicy(c)
	id := c.Param("id")
	idb, err := strconv.Atoi(id)

	ss, err := h.skemaService.Delete(idb)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	skemaResponse := convertSSToResponse(ss)

	deleteMessage := "Sucessfull Deleting Skema Sertifikasi " + skemaResponse.Judul

	c.JSON(http.StatusOK, gin.H{
		"data": deleteMessage,
	})
}

func (h *skemaHandler) CreateSkema(c *gin.Context) {
	CorsPolicy(c)
	var skemaRequest skema_sertifikasi.SkemaSertifikasiRequest

	err := c.ShouldBindJSON(&skemaRequest)
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

	skema, err := h.skemaService.Create(skemaRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, convertSSToResponse(skema))

}

func (h *skemaHandler) UpdateSkema(c *gin.Context) {
	CorsPolicy(c)
	var skemaUpdate skema_sertifikasi.SkemaSertifikasiRequest

	err := c.ShouldBindJSON(&skemaUpdate)
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
	skema, err := h.skemaService.Update(idb, skemaUpdate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, convertSSToResponse(skema))
}
