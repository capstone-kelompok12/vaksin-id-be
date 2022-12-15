package payload

type UpdateAccHistory struct {
	ID      string `json:"id_history"`
	NikUser string `json:"nik"`
	Status  string `json:"status"`
}
