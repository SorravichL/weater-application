package tests

import (
	"go-weather-api/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"io"
)

func TestWeather_NoCityProvided(t *testing.T) {
	r := router.SetupRouter()

	req, _ := http.NewRequest("GET", "/weather", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "city is required")
}

func TestWeather_ValidCity(t *testing.T) {
	r := router.SetupRouter()

	req, _ := http.NewRequest("GET", "/weather?city=Bangkok", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	body, err := io.ReadAll(w.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(body), "Bangkok") // assuming Visual Crossing returns "Bangkok" in JSON
}
