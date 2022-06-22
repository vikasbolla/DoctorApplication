package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/vikasbolla/DoctorApplication/Doctor_application/domain"
)

var respdoctor domain.RespDoctor

var appointment domain.Appointment

func GetSuitableAppointment(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&appointment)
	Gettintvalues()
	var apptfstr string
	var appttstr string
	for i := 11; i < 13; i++ {
		apptfstr = apptfstr + string(appointment.From[i])
		appttstr = appttstr + string(appointment.To[i])
	}
	apptfrom, _ := strconv.Atoi(apptfstr)
	apptto, _ := strconv.Atoi(appttstr)
	for i := 0; i < len(Doctors); i++ {
		for j := 0; j < 2; j++ {
			if Slot[i].TimeInt[j].From >= apptfrom || Slot[i].TimeInt[j].To <= apptto {
				respdoctor.DoctorId = Slot[i].Id
				respdoctor.Fullname = Slot[i].Name
				break
			}
		}
		json.NewEncoder(w).Encode(respdoctor)
	}
}
