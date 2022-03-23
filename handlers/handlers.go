package handlers

import (
	"encoding/json"
	"flare/exercise/data"
	"io"
	"log"
)

type Handler struct {
	l  *log.Logger
	db data.Filter
}

func NewHandler(l *log.Logger, f data.Filter) *Handler {
	return &Handler{l, f}
}

type UsernameResponse struct {
	Available bool   `json:"available"`
	Message   string `json:"message"`
}

type HealthResponse struct {
	Alive   bool   `json:"alive"`
	Message string `json:"message"`
}

type GenericError struct {
	Message string `json:"message"`
}

func ToJSON(i interface{}, w io.Writer) error {
	return json.NewEncoder(w).Encode(i)
}
