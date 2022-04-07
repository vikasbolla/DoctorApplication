package testing

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/vikasbolla/DoctorApplication/Doctor_application/service"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/getall", service.DoctorDetails).Methods("GET")
	return router
}

func TestDoctorDetails(t *testing.T) {
	request, _ := http.NewRequest("GET", "/getall", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "ok response is expected")
}
