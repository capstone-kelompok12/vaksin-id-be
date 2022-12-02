package payload

type VaccinesPayload struct {
	IdHealthFacilities string `json:"id_health_facilities"`
	Name               string `json:"name"`
	Stock              int    `json:"stock"`
}
