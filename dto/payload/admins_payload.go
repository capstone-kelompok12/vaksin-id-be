package payload

type AdminsPayload struct {
	IdHealthFacilities string `json:"id_health_facilities"`
	Email              string `json:"email"`
	Password           string `json:"password"`
}

type LoginAdmin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
