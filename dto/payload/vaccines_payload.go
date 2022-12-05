package payload

type VaccinesPayload struct {
	Name  string `json:"name" validate:"required"`
	Stock int    `json:"stock"`
}

type VaccinesUpdatePayload struct {
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}
