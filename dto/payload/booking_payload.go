package payload

type BookingPayload struct {
	NikUser   string `json:"nik"`
	IdSession string `json:"id_session"`
}

type BookingUpdate struct {
	ID        string `json:"booking_id"`
	NikUser   string `json:"nik"`
	IdSession string `json:"id_session"`
	Status    string `json:"status"`
}

type BookingCancel struct {
	ID        string `json:"booking_id"`
	IdSession string `json:"id_session"`
}
