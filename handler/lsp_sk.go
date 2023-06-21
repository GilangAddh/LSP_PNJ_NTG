package handler

import (
	"LSP_PNJ_NTG/entity"
	"LSP_PNJ_NTG/lsp_sk"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LSPSKHandler struct {
	LSPSKService lsp_sk.Service
}

func NewLSPSKHandler(LSPSKService lsp_sk.Service) *LSPSKHandler {
	return &LSPSKHandler{LSPSKService}
}

func convertLSToResponse(j entity.LSP_SK) lsp_sk.LSPSKResponse {
	return lsp_sk.LSPSKResponse{
		ID:    j.ID,
		LSPID: j.LSPID,
		SKID:  j.SKID,
	}
}

func (h *LSPSKHandler) GetLSPSKs(c *gin.Context) {
	CorsPolicy(c)
	LSPSKs, err := h.LSPSKService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var LSPSKsResponse []lsp_sk.LSPSKResponse

	for _, j := range LSPSKs {
		LSPSKResponse := convertLSToResponse(j)
		LSPSKsResponse = append(LSPSKsResponse, LSPSKResponse)
	}

	c.JSON(http.StatusOK, LSPSKsResponse)
}

func (h *LSPSKHandler) GetLSPSK(c *gin.Context) {
	CorsPolicy(c)
	id := c.Param("id")
	idb, err := strconv.Atoi(id)

	j, err := h.LSPSKService.FindById(idb)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	LSPSKResponse := convertLSToResponse(j)

	c.JSON(http.StatusOK, LSPSKResponse)
}

func (h *LSPSKHandler) DeleteLSPSK(c *gin.Context) {
	CorsPolicy(c)
	id := c.Param("id")
	idb, err := strconv.Atoi(id)

	j, err := h.LSPSKService.Delete(idb)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	lspResponse := convertLSToResponse(j)

	deleteMessage := "Sucessfull Deleting LSP SK ID" + strconv.Itoa(lspResponse.ID)

	c.JSON(http.StatusOK, gin.H{
		"data": deleteMessage,
	})
}

func (h *LSPSKHandler) CreateLSPSK(c *gin.Context) {
	CorsPolicy(c)
	var LSPSKRequest lsp_sk.LSPSKRequest

	err := c.ShouldBindJSON(&LSPSKRequest)
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

	LSPSK, err := h.LSPSKService.Create(LSPSKRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, convertLSToResponse(LSPSK))

}

func (h *LSPSKHandler) UpdateLSPSK(c *gin.Context) {
	CorsPolicy(c)
	var LSPSKUpdate lsp_sk.LSPSKRequest

	err := c.ShouldBindJSON(&LSPSKUpdate)
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
	LSPSK, err := h.LSPSKService.Update(idb, LSPSKUpdate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, convertLSToResponse(LSPSK))
}
