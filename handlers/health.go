package handlers

import (
	"net/http"
)

// swagger:route GET /health health checkHealth
// Checks service availability
//
// responses:
//	200: healthResponse

// CheckHealth handles the health status of the service
func (h *Handler) CheckHealth(rw http.ResponseWriter, _ *http.Request) {
	h.l.Println("[DEBUG] Handle GET CheckHealth")

	if !h.db.IsAlive() {
		h.l.Println("Service unavailable")
		rw.WriteHeader(http.StatusOK)
		ToJSON(&HealthResponse{ //nolint
			Alive:   false,
			Message: "Service is not available",
		}, rw)

		return
	}

	h.l.Println("Service is available")
	rw.WriteHeader(http.StatusOK)
	ToJSON(&HealthResponse{ //nolint
		Alive:   true,
		Message: "Service is available",
	}, rw)
}
