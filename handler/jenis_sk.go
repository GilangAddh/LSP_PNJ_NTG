package handler

import (
	"LSP_PNJ_NTG/entity"
	"LSP_PNJ_NTG/jenis_sk"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type jenisSKHandler struct {
	jenisSKService jenis_sk.Service
}

func NewJenisSKHandler(jenisSKService jenis_sk.Service) *jenisSKHandler {
	return &jenisSKHandler{jenisSKService}
}

func convertJSToResponse(j entity.JenisSK) jenis_sk.JenisSKResponse {
	return jenis_sk.JenisSKResponse{
		ID:   j.ID,
		Nama: j.Nama,
	}
}

func (h *jenisSKHandler) GetJenisSKs(c *gin.Context) {
	CorsPolicy(c)
	jenisSKs, err := h.jenisSKService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var jenisSKsResponse []jenis_sk.JenisSKResponse

	for _, j := range jenisSKs {
		jenisSKResponse := convertJSToResponse(j)
		jenisSKsResponse = append(jenisSKsResponse, jenisSKResponse)
	}

	c.JSON(http.StatusOK, jenisSKsResponse)
}

func (h *jenisSKHandler) GetJenisSK(c *gin.Context) {
	CorsPolicy(c)
	id := c.Param("id")
	idb, err := strconv.Atoi(id)

	j, err := h.jenisSKService.FindById(idb)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	jenisSKResponse := convertJSToResponse(j)

	c.JSON(http.StatusOK, jenisSKResponse)
}

func (h *jenisSKHandler) DeleteJenisSK(c *gin.Context) {
	CorsPolicy(c)
	id := c.Param("id")
	idb, err := strconv.Atoi(id)

	j, err := h.jenisSKService.Delete(idb)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	lspResponse := convertJSToResponse(j)

	deleteMessage := "Sucessfull Deleting Jenis LSP " + lspResponse.Nama

	c.JSON(http.StatusOK, gin.H{
		"data": deleteMessage,
	})
}

func (h *jenisSKHandler) CreateJenisSK(c *gin.Context) {
	CorsPolicy(c)
	var jenisSKRequest jenis_sk.JenisSKRequest

	err := c.ShouldBindJSON(&jenisSKRequest)
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

	jenisSK, err := h.jenisSKService.Create(jenisSKRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, convertJSToResponse(jenisSK))

}

func (h *jenisSKHandler) UpdatejenisSK(c *gin.Context) {
	CorsPolicy(c)
	var jenisSKUpdate jenis_sk.JenisSKRequest

	err := c.ShouldBindJSON(&jenisSKUpdate)
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
	jenisSK, err := h.jenisSKService.Update(idb, jenisSKUpdate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, convertJSToResponse(jenisSK))
}
