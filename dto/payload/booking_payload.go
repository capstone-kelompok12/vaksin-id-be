package payload

type BookingPayload struct {
	NikUser   string `json:"nik"`
	IdSession string `json:"id_session"`
}

type BookingUpdate struct {
	ID     string `json:"booking_id"`
	Status string `json:"status"`
}