package payload

type UpdateAccHistory struct {
	ID      string `json:"id_history"`
	NikUser string `json:"nik"`
	Status  string `json:"status"`
}

type HistoriesPayload struct {
	IdBooking  string `json:"id_booking"`
	NikUser    string `json:"nik"`
	IdSameBook string `json:"id_same_book"`
	Status     string `json:"status"`
}
