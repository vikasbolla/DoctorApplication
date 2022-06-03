package domain

//Doctor Details && Time Slots
type Doctor struct {
	DoctorId     string  `json:"doctorId"`
	Fullname     string  `json:"fullname"`
	Availability []Avail `json:"availability"`
}

type Avail struct {
	From string `json:"from"`
	To   string `json:"to"`
}

//Suitable Appointment Time Slots with Id
type Appointment struct {
	AppointmentId string `json:"appoinmentid"`
	From          string `json:"from"`
	To            string `json:"to"`
}

//Doctor duplicate data
//Duplicate of Doctor Details && int Time Slots for decision making
type Doctordup struct {
	Id      string `json:"doctorid"`
	Name    string `json:"fullname"`
	TimeInt [2]TimeInt
}

type TimeInt struct {
	From int `json:"from"`
	To   int `json:"to"`
}

//Final Response with required details
type RespDoctor struct {
	DoctorId string `json:"doctorId"`
	Fullname string `json:"fullname"`
}
