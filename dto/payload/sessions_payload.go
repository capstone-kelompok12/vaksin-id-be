package payload

type SessionsPayload struct {
	IdVaccine    string `json:"id_vaccine" validate:"required"`
	SessionName  string `json:"session_name" validate:"required"`
	Capacity     int    `json:"capacity" validate:"required"`
	Date         string `json:"date" gorm:"not null" validate:"required"`
	StartSession string `json:"start" validate:"required,max=5"`
	EndSession   string `json:"end" validate:"required,max=5"`
}

type SessionsUpdate struct {
	IdVaccine    string `json:"id_vaccine"`
	SessionName  string `json:"session_name"`
	Capacity     int    `json:"capacity"`
	Date         string `json:"date"`
	StartSession string `json:"start"`
	EndSession   string `json:"end"`
}
type SessionsIsClose struct {
	IsClose bool `json:"is_close"`
}
