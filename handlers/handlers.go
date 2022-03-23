package handlers

import (
	"encoding/json"
	"flare/exercise/data"
	"io"
	"log"
)

type Handler struct {
	l  *log.Logger // logger
	db data.DB     // in-memory db
}

func NewHandler(l *log.Logger, f data.DB) *Handler {
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

// ToJSON writes JSON encoded value to writer
func ToJSON(i interface{}, w io.Writer) error {
	return json.NewEncoder(w).Encode(i)
}
