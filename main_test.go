package main

import (
	"encoding/json"
	"flare/exercise/data"
	"flare/exercise/handlers"
	"flare/exercise/utils"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

func TestMain(t *testing.T) { //nolint
	l := log.New(os.Stdout, "username-api ", log.LstdFlags)
	f := data.NewDB()

	ph := handlers.NewHandler(l, f)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()

	// Set up endpoints for the GET router
	getRouter.HandleFunc("/username", ph.CheckUsername).Queries("q", "{query}")
	getRouter.HandleFunc("/health", ph.CheckHealth)

	t.Run("Test valid, taken usernames", func(t *testing.T) {
		var insertedNames []string
		for i := 0; i <= 100; i++ {
			n := utils.RandomInt(8, 11)
			name := utils.RandomStringValid(int(n))
			insertedNames = append(insertedNames, name)

			f.AddEntry(name)
		}

		var res handlers.UsernameResponse

		for _, name := range insertedNames {
			url := "http://localhost:9090/username?q=" + name
			req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
			if err != nil {
				t.Fatalf("Couldn't create request: %v\n", err)
			}

			w := httptest.NewRecorder()

			sm.ServeHTTP(w, req)

			if w.Code == http.StatusOK {
				t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, w.Code)
				if err := json.NewDecoder(w.Body).Decode(&res); err != nil {
					t.Fatalf("unable to deserialize response: %v", err)
				}
				if res.Available {
					t.Errorf("Want %v Got %v", false, res.Available)
				}
			} else {
				t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
			}
		}
	})

	t.Run("Test valid, available usernames", func(t *testing.T) {
		var newNames []string
		for i := 0; i <= 100; i++ {
			n := utils.RandomInt(8, 11)
			name := utils.RandomStringValid(int(n))
			newNames = append(newNames, name)
		}

		var userRes handlers.UsernameResponse

		for _, name := range newNames {
			url := "http://localhost:9090/username?q=" + name
			req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
			if err != nil {
				t.Fatalf("Couldn't create request: %v\n", err)
			}

			w := httptest.NewRecorder()

			sm.ServeHTTP(w, req)

			if w.Code == http.StatusOK {
				t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, w.Code)
				if err := json.NewDecoder(w.Body).Decode(&userRes); err != nil {
					t.Fatalf("unable to deserialize response: %v", err)
				}
				if !userRes.Available {
					t.Errorf("Want %v Got %v", true, userRes.Available)
				}
			} else {
				t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
			}
		}
	})

	t.Run("Test invalid usernames", func(t *testing.T) { //nolint
		var invalidNames []string
		for i := 0; i <= 100; i++ {
			n := utils.RandomInt(1, 20)
			name := utils.RandomStringInvalid(int(n))
			invalidNames = append(invalidNames, name)
		}

		var errRes handlers.GenericError
		for _, name := range invalidNames {
			url := "http://localhost:9090/username?q=" + name
			req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
			if err != nil {
				t.Fatalf("Couldn't create request: %v\n", err)
			}

			w := httptest.NewRecorder()

			sm.ServeHTTP(w, req)

			if w.Code == http.StatusBadRequest {
				t.Logf("Expected to get status %d is same as %d\n", http.StatusBadRequest, w.Code)
				if err := json.NewDecoder(w.Body).Decode(&errRes); err != nil {
					t.Fatalf("unable to deserialize response: %v", err)
				}
			} else {
				t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusBadRequest, w.Code)
			}
		}
	})

	t.Run("Test check health endpoint", func(t *testing.T) {
		url := "http://localhost:9090/health"
		req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
		if err != nil {
			t.Fatalf("Couldn't create request: %v\n", err)
		}

		w := httptest.NewRecorder()

		var healthRes handlers.HealthResponse

		sm.ServeHTTP(w, req)

		if w.Code == http.StatusOK {
			t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, w.Code)
			if err := json.NewDecoder(w.Body).Decode(&healthRes); err != nil {
				t.Fatalf("unable to deserialize response: %v", err)
			}
			if !healthRes.Alive {
				t.Errorf("Want service alive: %v Got service alive: %v", healthRes.Alive, false)
			}
		} else {
			t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
		}
	})
}
