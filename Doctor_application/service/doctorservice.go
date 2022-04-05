package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/vikasbolla/DoctorApplication/Doctor_application/domain"
)

var Doctors []domain.Doctor

//var doctor domain.Doctor

var Slot [100]domain.Doctordup

func DoctorDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Doctors)

}

func PostAvailability(w http.ResponseWriter, r *http.Request) {
	var doctor domain.Doctor
	w.Header().Add("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&doctor)
	Doctors = append(Doctors, doctor)
	fmt.Fprintf(w, "Successfully Posted")
}

func Gettintvalues() {
	for i := 0; i < len(Doctors); i++ {
		for j := 0; j < 2; j++ {
			var fnumber string = ""
			var tnumber string = ""
			for k := 11; k < 13; k++ {
				fnumber = fnumber + string(Doctors[i].Availability[j].From[k])
				tnumber = tnumber + string(Doctors[i].Availability[j].To[k])
			}
			Slot[i].Id = Doctors[i].DoctorId
			Slot[i].Name = Doctors[i].Fullname
			fnum, _ := strconv.Atoi(fnumber)
			tnum, _ := strconv.Atoi(tnumber)
			Slot[i].TimeInt[j].From = fnum
			Slot[i].TimeInt[j].To = tnum
		}
	}
	for i := 0; i < len(Doctors); i++ {
		fmt.Printf("%+v \n", Slot[i])
	}
}
