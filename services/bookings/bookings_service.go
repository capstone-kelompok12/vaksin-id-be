package services

import (
	"vaksin-id-be/dto/payload"
	"vaksin-id-be/dto/response"
	"vaksin-id-be/model"
	mysqlb "vaksin-id-be/repository/mysql/bookings"

	"github.com/google/uuid"
)

type BookingService interface {
	CreateBooking(payloads payload.BookingPayload) (payload.BookingPayload, error)
	UpdateBooking(payloads payload.BookingUpdate, id string) (response.BookingResponse, error)
	GetAllBooking() ([]response.BookingResponse, error)
	GetBookingDashboard() (response.Dashboard, error)
	GetBooking(id string) (response.BookingResponse, error)
	DeleteBooking(id string) error
}

type bookingService struct {
	BookingRepo mysqlb.BookingRepository
}

func NewBookingService(bookingRepo mysqlb.BookingRepository) *bookingService {
	return &bookingService{
		BookingRepo: bookingRepo,
	}
}

func (b *bookingService) CreateBooking(payloads payload.BookingPayload) (payload.BookingPayload, error) {
	var dataResp payload.BookingPayload
	id := uuid.NewString()

	bookingModel := model.BookingSessions{
		ID:        id,
		NikUser:   payloads.NikUser,
		IdSession: payloads.IdSession,
	}

	err := b.BookingRepo.CreateBooking(bookingModel)
	if err != nil {
		return dataResp, err
	}
	dataResp = payloads

	return dataResp, nil
}

func (b *bookingService) UpdateBooking(payloads payload.BookingUpdate, id string) (response.BookingResponse, error) {
	var data response.BookingResponse

	bookingData := model.BookingSessions{
		ID:     payloads.ID,
		Status: payloads.Status,
	}

	if err := b.BookingRepo.UpdateBooking(bookingData, id); err != nil {
		return data, err
	}

	data, err := b.GetBooking(id)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (b *bookingService) GetAllBooking() ([]response.BookingResponse, error) {
	var bookingResponse []response.BookingResponse

	getBooking, err := b.BookingRepo.GetAllBooking()

	if err != nil {
		return bookingResponse, err
	}

	bookingResponse = make([]response.BookingResponse, len(getBooking))

	for i, v := range getBooking {
		bookingResponse[i] = response.BookingResponse{
			ID:        v.ID,
			NikUser:   v.NikUser,
			IdSession: v.IdSession,
			Queue:     v.Queue,
			Status:    v.Status,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			User:      *v.User,
			Session:   *v.Session,
		}
	}

	return bookingResponse, nil
}

func (b *bookingService) GetBooking(id string) (response.BookingResponse, error) {
	var responseBooking response.BookingResponse

	getData, err := b.BookingRepo.GetBooking(id)
	if err != nil {
		return responseBooking, err
	}

	responseBooking = response.BookingResponse{
		ID:        getData.ID,
		NikUser:   getData.NikUser,
		IdSession: getData.IdSession,
		Queue:     getData.Queue,
		Status:    getData.Status,
		CreatedAt: getData.CreatedAt,
		UpdatedAt: getData.UpdatedAt,
		User:      *getData.User,
	}

	return responseBooking, nil
}

func (b *bookingService) DeleteBooking(id string) error {
	if err := b.BookingRepo.DeleteBooking(id); err != nil {
		return err
	}

	return nil
}

func (b *bookingService) GetBookingDashboard() (response.Dashboard, error) {
	var bookingResponse response.Dashboard

	getBooking, err := b.BookingRepo.GetAllBooking()

	if err != nil {
		return bookingResponse, err
	}

	count := len(getBooking)

	bookingResponse = response.Dashboard{Booking: count}

	return bookingResponse, nil
}
