package payload

type SessionsPayload struct {
	SessionName  string `json:"session_name" validate:"required"`
	Capacity     int    `json:"capacity" validate:"required"`
	Dose         int    `json:"dose" validate:"required"`
	StartSession string `json:"start" validate:"required,max=5"`
	EndSession   string `json:"end" validate:"required,max=5"`
}

type SessionsUpdate struct {
	SessionName  string `json:"session_name"`
	Capacity     int    `json:"capacity"`
	Dose         int    `json:"dose"`
	StartSession string `json:"start"`
	EndSession   string `json:"end"`
}
type SessionsIsClose struct {
	IsClose bool `json:"is_close"`
}
