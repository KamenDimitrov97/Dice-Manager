package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDiceEndpoint(t *testing.T) {
	t.Run("Delete dice successfully", func(t *testing.T) {
		diceBag = []int{6, 8, 10}

		req, err := http.NewRequest("DELETE", "/delete", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(deleteDiceHandler)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		resp, _ := strconv.Atoi(string(rr.Body.String()))
		assert.GreaterOrEqual(t, resp, 1)
		assert.Equal(t, 2, len(diceBag))
	})

	t.Run("Delete dice not found", func(t *testing.T) {
		diceBag = []int{}

		req, err := http.NewRequest("DELETE", "/delete", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(deleteDiceHandler)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)

		assert.Contains(t, rr.Body.String(), "Dice not found")
	})
}

func TestAddDiceEndpoint(t *testing.T) {
	t.Run("Add dice successfully", func(t *testing.T) {
		diceBag = []int{}

		body := `{"dice": [6, 8, 10]}`
		req, err := http.NewRequest(http.MethodPost, "/dice/new", strings.NewReader(body))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(addDiceHandler)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		assert.Equal(t, []int{6, 8, 10}, diceBag)
	})

	t.Run("Add dice with wrong http method", func(t *testing.T) {
		diceBag = []int{}

		body := `{"dice": [6, 8, 10]}`
		req, err := http.NewRequest(http.MethodPut, "/dice/new", strings.NewReader(body))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(addDiceHandler)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)

		assert.Equal(t, []int{}, diceBag)
	})
}
