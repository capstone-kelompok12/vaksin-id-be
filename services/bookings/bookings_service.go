package services

import (
	"errors"
	"fmt"
	"time"
	"vaksin-id-be/dto/payload"
	"vaksin-id-be/dto/response"
	"vaksin-id-be/model"
	mysqlb "vaksin-id-be/repository/mysql/bookings"
	mysqlh "vaksin-id-be/repository/mysql/histories"
	mysqls "vaksin-id-be/repository/mysql/sessions"

	"github.com/google/uuid"
)

type BookingService interface {
	CreateBooking(payloads []payload.BookingPayload) ([]model.VaccineHistories, error)
	UpdateBooking(payloads payload.BookingUpdate, id, nik string) ([]response.BookingResponse, error)
	GetAllBooking() ([]response.BookingResponse, error)
	GetBookingDashboard() (response.Dashboard, error)
	GetBooking(id string) ([]response.BookingResponse, error)
	DeleteBooking(id string) error
}

type bookingService struct {
	BookingRepo mysqlb.BookingRepository
	HistoryRepo mysqlh.HistoriesRepository
	SessionRepo mysqls.SessionsRepository
}

func NewBookingService(bookingRepo mysqlb.BookingRepository, historyRepo mysqlh.HistoriesRepository, sessionRepo mysqls.SessionsRepository) *bookingService {
	return &bookingService{
		BookingRepo: bookingRepo,
		HistoryRepo: historyRepo,
		SessionRepo: sessionRepo,
	}
}

func (b *bookingService) CreateBooking(payloads []payload.BookingPayload) ([]model.VaccineHistories, error) {
	var historyData []model.VaccineHistories = make([]model.VaccineHistories, len(payloads))
	var bookingModel []model.BookingSessions = make([]model.BookingSessions, len(payloads))
	idSameBooked := uuid.NewString()
	createdData := time.Now()
	updatedData := time.Now()

	for i, val := range payloads {
		defValue := 0

		_, err := b.HistoryRepo.GetHistoryByNIK(payloads[i].NikUser)
		if err == nil {
			return historyData, errors.New("already booked")
		}

		bookingModel[i] = model.BookingSessions{
			ID:        uuid.NewString(),
			IdSession: val.IdSession,
			Queue:     &defValue,
			CreatedAt: createdData,
			UpdatedAt: updatedData,
		}

		historyData[i] = model.VaccineHistories{
			ID:         uuid.NewString(),
			IdBooking:  bookingModel[i].ID,
			NikUser:    val.NikUser,
			IdSameBook: idSameBooked,
			CreatedAt:  createdData,
			UpdatedAt:  updatedData,
		}
		err = b.BookingRepo.CreateBooking(bookingModel[i])
		if err != nil {
			return historyData, err
		}
	}

	for _, v := range historyData {
		err := b.HistoryRepo.CreateHistory(v)
		if err != nil {
			return historyData, err
		}
	}

	getSession, err := b.SessionRepo.GetSessionById(payloads[0].IdSession)
	if err != nil {
		return historyData, err
	}

	countBook := getSession.Capacity - len(payloads)

	updateSession := model.Sessions{
		Capacity: countBook,
	}

	err = b.SessionRepo.UpdateSession(updateSession, payloads[0].IdSession)
	if err != nil {
		return historyData, err
	}

	getData, err := b.HistoryRepo.GetHistoryByIdSameBook(idSameBooked)
	if err != nil {
		return historyData, err
	}

	return getData, nil
}

func (b *bookingService) UpdateBooking(payloads payload.BookingUpdate, id, nik string) ([]response.BookingResponse, error) {
	var data []response.BookingResponse
	var bookingData []model.BookingSessions

	if payloads.Status != "Accepted" && payloads.Status != "Denied" {
		return data, errors.New("input status must Accepted or Denied")
	}

	dataById, err := b.GetBooking(id)
	if err != nil {
		return data, err
	}
	getSession, err := b.SessionRepo.GetSessionById(dataById[0].IdSession)
	if err != nil {
		return data, err
	}
	getbackCap := getSession.Capacity + len(dataById)

	data = make([]response.BookingResponse, len(dataById))
	bookingData = make([]model.BookingSessions, len(dataById))

	if payloads.Status == "Denied" {

		updateSession := model.Sessions{
			Capacity: getbackCap,
		}

		err = b.SessionRepo.UpdateSession(updateSession, dataById[0].IdSession)
		if err != nil {
			return data, err
		}

		bookingData := model.BookingSessions{
			ID:     payloads.ID,
			Status: &payloads.Status,
		}

		if err := b.BookingRepo.UpdateBooking(bookingData, id, nik); err != nil {
			return data, err
		}
		return data, nil
	}

	// countData, err := b.HistoryRepo.GetHistoryByIdBooking(id)
	// if err != nil {
	// 	return data, nil
	// }

	for i, val := range data {
		var tempQueu int
		// manyCount := len(countData)
		def := 0
		if val.Queue == &def {
			tempQueu = 1
		}
		// else {
		// 	tempQueu = manyCount + 1
		// }
		fmt.Println(tempQueu)
		bookingData[i] = model.BookingSessions{
			ID:     val.ID,
			Status: val.Status,
			Queue:  val.Queue,
		}

		if err := b.BookingRepo.UpdateBookingAcc(bookingData, id); err != nil {
			return data, err
		}
		// if data[0].Queue == nil {
		// 	bookingData = model.BookingSessions{
		// 		ID:     payloads.ID,
		// 		Status: &payloads.Status,
		// 		Queue:  &tempQueue,
		// 	}
		// } else if !data[0] {
		// 	tempQueue += 1
		// 	bookingData = model.BookingSessions{
		// 		ID:     payloads.ID,
		// 		Status: &payloads.Status,
		// 		Queue:  &tempQueue,
		// 	}
		// }
	}
	// newQueue := dataById[0]

	// bookingData := model.BookingSessions{
	// 	ID:     payloads.ID,
	// 	Status: &payloads.Status,
	// }

	// data, err = b.GetBooking(id)
	// if err != nil {
	// 	return data, err
	// }

	return dataById, nil
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
			IdSession: v.IdSession,
			Queue:     v.Queue,
			Status:    v.Status,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Session:   *v.Session,
			History:   v.History,
		}
	}

	return bookingResponse, nil
}

func (b *bookingService) GetBooking(id string) ([]response.BookingResponse, error) {
	var responseBooking []response.BookingResponse

	getData, err := b.BookingRepo.GetBooking(id)
	if err != nil {
		return responseBooking, err
	}

	responseBooking = make([]response.BookingResponse, len(getData))

	for i, val := range getData {
		responseBooking[i] = response.BookingResponse{
			ID:        val.ID,
			IdSession: val.IdSession,
			Queue:     val.Queue,
			Status:    val.Status,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		}
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
