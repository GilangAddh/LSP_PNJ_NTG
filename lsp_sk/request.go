package lsp_sk

import (
	"encoding/json"
)

type LSPSKRequest struct {
	LSPID json.Number `json:"LSPID" binding"required number"`
	SKID  json.Number `json:"SKID" binding"required number"`
}
