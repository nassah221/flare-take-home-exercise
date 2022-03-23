package handlers

import (
	"fmt"
	"net/http"
	"regexp"
)

var invalidUsernameMsg = `Please make sure your username follows these rules
1. Does not start with a number or special character
2. Does not contain a special character
3. Has at least 8 and at most 12 alphanumeric characters`

// swagger:route GET /username username checkUsername
// Checks the supplied username for availability
//
// responses:
// 	200: usernameResponse
// 	400: errorResponse
//
// parameters:
// 	+ name: query
//	  in: query
//	  required: true
//	  type: string

// CheckUsername handles checking user availability
func (h *Handler) CheckUsername(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("[DEBUG] Handle GET CheckUsername")
	query := r.URL.Query().Get("q")

	if query == "" {
		h.l.Println("Empty query parameter")
		http.Error(rw, "Empty query parameter", http.StatusBadRequest)
		return
	}

	// Deliberately using a simple regex which checks against the following conditions
	// 1. Username must start with an alphabet, upper or lowercase
	// 2. Username must have at least 8, and at most 12 characters
	// Other than the above rules, the username can have any combination of alphanumeric characters
	reg := regexp.MustCompile("^[A-Za-z][A-Za-z0-9_]{7,11}$")
	name := reg.FindString(query)
	if name == "" {
		h.l.Printf("Username [%s] does not match scheme", query)

		rw.WriteHeader(http.StatusBadRequest)
		ToJSON(&GenericError{ //nolint
			Message: fmt.Sprintf("Invalid user name scheme. %s", invalidUsernameMsg),
		}, rw)

		return
	}

	h.l.Printf("[DEBUG] Got username: %s", query)

	if exists := h.db.CheckUsername(query); !exists {
		h.l.Printf("Username [%s] is available for use\n", query)

		rw.WriteHeader(http.StatusOK)
		ToJSON(&UsernameResponse{ //nolint
			Available: true,
			Message:   fmt.Sprintf("Username [%s] is available for use\n", query),
		}, rw)

		return
	}

	h.l.Printf("Username [%s] already exists\n", query)

	rw.WriteHeader(http.StatusOK)
	ToJSON(&UsernameResponse{ //nolint
		Available: false,
		Message:   fmt.Sprintf("Username [%s] already exists\n", query),
	}, rw)
}
