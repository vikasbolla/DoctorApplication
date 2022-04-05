package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/getall", DoctorDetails).Methods("GET")
	myRouter.HandleFunc("/postavailability", PostAvailability).Methods("POST")
	myRouter.HandleFunc("/getsuitableappointment", GetSuitableAppointment).Methods("POST")
	log.Fatal((http.ListenAndServe(":9009", myRouter)))
}
