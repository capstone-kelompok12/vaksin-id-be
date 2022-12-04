package payload

import "time"

type SessionsPayload struct {
	IdHealthFacilities string `json:"id_health_facilities"`
	SessionName        string `json:"session_name"`
	Capacity           int    `json:"capacity"`
	StartSession       time.Time
	EndSession         time.Time
}

type SessionsUpdate struct {
	ID            string `json:"session_id"`
	SessionName   string `json:"session_name"`
	Capacity      int    `json:"capacity"`
	SessionStatus bool   `json:"session_status"`
	StartSession  time.Time
	EndSession    time.Time
}
