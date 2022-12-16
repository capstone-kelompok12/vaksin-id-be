package services

import (
	"errors"
	"time"
	"vaksin-id-be/dto/payload"
	"vaksin-id-be/dto/response"
	"vaksin-id-be/model"
	mysqlb "vaksin-id-be/repository/mysql/bookings"
	mysqlh "vaksin-id-be/repository/mysql/histories"
	mysqls "vaksin-id-be/repository/mysql/sessions"
	mysqlu "vaksin-id-be/repository/mysql/users"

	"github.com/google/uuid"
)

type BookingService interface {
	CreateBooking(payloads []payload.BookingPayload) ([]response.BookingResponseCustom, error)
	GetAllBooking() ([]response.BookingResponse, error)
	GetBookingBysSession(id string) ([]response.BookingResponse, error)
	GetBookingDashboard() (response.Dashboard, error)
	GetBooking(id string) (response.BookingResponse, error)
	UpdateAccAttendend(payloads []payload.UpdateAccHistory) ([]response.HistoryResponse, error)
	UpdateBooking(payloads []payload.BookingUpdate) ([]response.BookingResponseCustom, error)
	DeleteBooking(id string) error
}

type bookingService struct {
	BookingRepo mysqlb.BookingRepository
	HistoryRepo mysqlh.HistoriesRepository
	SessionRepo mysqls.SessionsRepository
	UserRepo    mysqlu.UserRepository
}

func NewBookingService(bookingRepo mysqlb.BookingRepository, historyRepo mysqlh.HistoriesRepository, sessionRepo mysqls.SessionsRepository, userRepo mysqlu.UserRepository) *bookingService {
	return &bookingService{
		BookingRepo: bookingRepo,
		HistoryRepo: historyRepo,
		SessionRepo: sessionRepo,
		UserRepo:    userRepo,
	}
}

func (b *bookingService) CreateBooking(payloads []payload.BookingPayload) ([]response.BookingResponseCustom, error) {
	historyData := make([]model.VaccineHistories, len(payloads))
	bookingModel := make([]model.BookingSessions, len(payloads))
	resBooking := make([]response.BookingResponseCustom, len(payloads))

	idSameBooked := uuid.NewString()
	createdData := time.Now()
	updatedData := time.Now()
	idBooking := ""
	statusData := "OnProccess"

	for i, val := range payloads {
		defValue := 0
		idBooking = uuid.NewString()
		defId := idBooking

		_, err := b.HistoryRepo.GetHistoryByNIK(defId, val.NikUser)
		if err == nil {
			return resBooking, errors.New("already booked")
		}

		bookingModel[i] = model.BookingSessions{
			ID:        idBooking,
			IdSession: val.IdSession,
			Queue:     defValue,
			Status:    &statusData,
			CreatedAt: createdData,
			UpdatedAt: updatedData,
		}

		historyData[i] = model.VaccineHistories{
			ID:         uuid.NewString(),
			IdBooking:  defId,
			NikUser:    val.NikUser,
			IdSameBook: idSameBooked,
			Status:     &statusData,
			CreatedAt:  createdData,
			UpdatedAt:  updatedData,
		}
		err = b.BookingRepo.CreateBooking(bookingModel[i])
		if err != nil {
			return resBooking, err
		}
	}

	for _, v := range historyData {
		err := b.HistoryRepo.CreateHistory(v)
		if err != nil {
			return resBooking, err
		}
	}

	getSession, err := b.SessionRepo.GetSessionById(payloads[0].IdSession)
	if err != nil {
		return resBooking, err
	}

	getBookedByIdSession, err := b.BookingRepo.GetBookingBySession(payloads[0].IdSession)
	if err != nil {
		return resBooking, err
	}

	countBook := getSession.Capacity - len(getBookedByIdSession)

	for i, val := range bookingModel {
		getSessionId, err := b.SessionRepo.GetSessionById(val.IdSession)
		if err != nil {
			return resBooking, err
		}
		getHistorySame, err := b.HistoryRepo.GetHistoryByIdSameBook(idSameBooked)
		if err != nil {
			return resBooking, err
		}
		sessionData := response.BookingSessionCustom{
			ID:           getSessionId.ID,
			IdVaccine:    getSessionId.IdVaccine,
			SessionName:  getSessionId.SessionName,
			Capacity:     countBook,
			Dose:         getSessionId.Dose,
			Date:         getSessionId.Date,
			IsClose:      getSession.IsClose,
			StartSession: getSessionId.StartSession,
			EndSession:   getSessionId.EndSession,
			CreatedAt:    getSessionId.CreatedAt,
			UpdatedAt:    getSessionId.UpdatedAt,
			Vaccine:      getSession.Vaccine,
		}

		historyData := make([]response.BookingHistoryCustom, len(getHistorySame))
		if len(historyData) != 0 {
			historyData[i] = response.BookingHistoryCustom{
				ID:         getHistorySame[i].ID,
				IdBooking:  getHistorySame[i].IdBooking,
				NikUser:    getHistorySame[i].NikUser,
				IdSameBook: getHistorySame[i].IdSameBook,
				Status:     getHistorySame[i].Status,
				CreatedAt:  getHistorySame[i].CreatedAt,
				UpdatedAt:  getHistorySame[i].UpdatedAt,
				User:       getHistorySame[i].User,
			}
		}

		resBooking[i] = response.BookingResponseCustom{
			ID:        val.ID,
			IdSession: val.IdSession,
			Queue:     &val.Queue,
			Status:    val.Status,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
			Session:   sessionData,
			History:   historyData,
		}
	}

	return resBooking, nil
}

func (b *bookingService) UpdateBooking(payloads []payload.BookingUpdate) ([]response.BookingResponseCustom, error) {
	var data []response.BookingResponseCustom
	var bookingData []model.BookingSessions

	data = make([]response.BookingResponseCustom, len(payloads))
	bookingData = make([]model.BookingSessions, len(payloads))
	newRes := make([]response.BookingResponseCustom, len(payloads))
	newResReject := make([]response.BookingResponseCustom, len(payloads))

	var newQueue int
	var initQueu int = 0
	var getbackCap int

	for i, v := range payloads {

		getSession, err := b.SessionRepo.GetSessionById(v.IdSession)
		if err != nil {
			return data, err
		}

		if v.Status != "Accepted" && v.Status != "Rejected" {
			return data, errors.New("input status must Accepted or Rejected")
		}

		if v.Status == "Rejected" {

			bookingData := model.BookingSessions{
				ID:        v.ID,
				IdSession: v.IdSession,
				Status:    &v.Status,
			}

			if err := b.BookingRepo.UpdateBooking(bookingData); err != nil {
				return data, err
			}

			dataHistoryUpdate := model.VaccineHistories{
				Status: &v.Status,
			}

			idSameBook, err := b.HistoryRepo.GetHistoryByNIK(v.ID, v.NikUser)
			if err != nil {
				return data, err
			}

			newData, err := b.BookingRepo.GetBooking(v.ID)
			if err != nil {
				return data, err
			}

			getBookedByIdSession, err := b.BookingRepo.GetBookingBySessionDen(v.IdSession)
			if err != nil {
				return data, err
			}
			getbackCap = getSession.Capacity - len(getBookedByIdSession)

			getHistorySame, err := b.HistoryRepo.GetHistoryByIdSameBook(idSameBook.IdSameBook)
			if err != nil {
				return data, err
			}

			_, err = b.HistoryRepo.UpdateHistoryByNik(dataHistoryUpdate, v.NikUser, v.ID)
			if err != nil {
				return data, err
			}

			sessionData := response.BookingSessionCustom{
				ID:           getSession.ID,
				IdVaccine:    getSession.IdVaccine,
				SessionName:  getSession.SessionName,
				Capacity:     getbackCap,
				Dose:         getSession.Dose,
				Date:         getSession.Date,
				IsClose:      getSession.IsClose,
				StartSession: getSession.StartSession,
				EndSession:   getSession.EndSession,
				CreatedAt:    getSession.CreatedAt,
				UpdatedAt:    getSession.UpdatedAt,
				Vaccine:      getSession.Vaccine,
			}

			historyData := make([]response.BookingHistoryCustom, len(getHistorySame))
			if len(getHistorySame) != 0 {
				historyData[i] = response.BookingHistoryCustom{
					ID:         getHistorySame[i].ID,
					IdBooking:  getHistorySame[i].IdBooking,
					NikUser:    getHistorySame[i].NikUser,
					IdSameBook: getHistorySame[i].IdSameBook,
					Status:     getHistorySame[i].Status,
					CreatedAt:  getHistorySame[i].CreatedAt,
					UpdatedAt:  getHistorySame[i].UpdatedAt,
					User:       getHistorySame[i].User,
				}
			}

			newResReject[i] = response.BookingResponseCustom{
				ID:        newData.ID,
				IdSession: newData.IdSession,
				Queue:     &newData.Queue,
				Status:    newData.Status,
				CreatedAt: newData.CreatedAt,
				UpdatedAt: newData.UpdatedAt,
				Session:   sessionData,
				History:   historyData,
			}

			return newResReject, nil
		}

		checkStatus, err := b.BookingRepo.GetBooking(v.ID)
		if err != nil {
			return data, err
		}

		if *checkStatus.Status == "Accepted" {
			return data, errors.New("you already accept users")
		}

		lastQueue, err := b.BookingRepo.FindMaxQueue(v.IdSession)
		if err != nil {
			return data, err
		}
		newQueue = lastQueue.Queue

		if newQueue != 0 {
			initQueu = newQueue + 1
		} else {
			initQueu += 1
		}

		bookingData[i] = model.BookingSessions{
			ID:        v.ID,
			IdSession: v.IdSession,
			Queue:     initQueu,
			Status:    &v.Status,
		}
		initQueu += 1

		updateData, err := b.BookingRepo.UpdateBookingAcc(bookingData[i])
		if err != nil {
			return data, err
		}

		idSameBook, err := b.HistoryRepo.GetHistoryByNIK(v.ID, v.NikUser)
		if err != nil {
			return data, err
		}

		getHistorySame, err := b.HistoryRepo.GetHistoryByIdSameBook(idSameBook.IdSameBook)
		if err != nil {
			return data, err
		}

		getBookedByIdSession, err := b.BookingRepo.GetBookingBySessionDen(v.IdSession)
		if err != nil {
			return data, err
		}

		getbackCap = getSession.Capacity - len(getBookedByIdSession)

		sessionData := response.BookingSessionCustom{
			ID:           getSession.ID,
			IdVaccine:    getSession.IdVaccine,
			SessionName:  getSession.SessionName,
			Capacity:     getbackCap,
			Dose:         getSession.Dose,
			Date:         getSession.Date,
			IsClose:      getSession.IsClose,
			StartSession: getSession.StartSession,
			EndSession:   getSession.EndSession,
			CreatedAt:    getSession.CreatedAt,
			UpdatedAt:    getSession.UpdatedAt,
			Vaccine:      getSession.Vaccine,
		}
		historyData := make([]response.BookingHistoryCustom, len(getHistorySame))
		if len(getHistorySame) != 0 {
			historyData[i] = response.BookingHistoryCustom{
				ID:         getHistorySame[i].ID,
				IdBooking:  getHistorySame[i].IdBooking,
				NikUser:    getHistorySame[i].NikUser,
				IdSameBook: getHistorySame[i].IdSameBook,
				Status:     getHistorySame[i].Status,
				CreatedAt:  getHistorySame[i].CreatedAt,
				UpdatedAt:  getHistorySame[i].UpdatedAt,
				User:       getHistorySame[i].User,
			}
		}
		newRes[i] = response.BookingResponseCustom{
			ID:        updateData.ID,
			IdSession: updateData.IdSession,
			Queue:     &updateData.Queue,
			Status:    updateData.Status,
			CreatedAt: updateData.CreatedAt,
			UpdatedAt: updateData.UpdatedAt,
			Session:   sessionData,
			History:   historyData,
		}
	}
	return newRes, nil
}

func (b *bookingService) UpdateAccAttendend(payloads []payload.UpdateAccHistory) ([]response.HistoryResponse, error) {
	updateData := make([]response.HistoryResponse, len(payloads))

	var defaultVaccineCount int

	for i, val := range payloads {
		// update history berdasarkan id
		if val.Status != "Attended" && val.Status != "Absent" {
			return updateData, errors.New("input status must Attended or Absent")
		}

		updateModel := model.VaccineHistories{
			ID:     val.ID,
			Status: &val.Status,
		}

		_, err := b.HistoryRepo.UpdateHistory(updateModel, val.ID)
		if err != nil {
			return updateData, err
		}

		dataNik, err := b.HistoryRepo.CheckVaccineCount(val.NikUser)
		if err != nil {
			return updateData, err
		}

		defaultVaccineCount = len(dataNik)

		updateUserCount := model.Users{
			NIK:          val.NikUser,
			VaccineCount: defaultVaccineCount,
		}

		if err := b.UserRepo.UpdateUserProfile(updateUserCount); err != nil {
			return updateData, err
		}

		dataHistory, err := b.HistoryRepo.GetHistoryById(val.ID)
		if err != nil {
			return updateData, err
		}

		updateData[i] = response.HistoryResponse{
			ID:         dataHistory.ID,
			IdBooking:  dataHistory.IdBooking,
			NikUser:    dataHistory.NikUser,
			IdSameBook: dataHistory.IdBooking,
			Status:     dataHistory.Status,
			CreatedAt:  dataHistory.CreatedAt,
			UpdatedAt:  dataHistory.UpdatedAt,
			User:       *dataHistory.User,
		}
	}
	return updateData, nil
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
			Queue:     &v.Queue,
			Status:    v.Status,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Session:   *v.Session,
			History:   v.History,
		}
	}

	return bookingResponse, nil
}

func (b *bookingService) GetBookingBysSession(id string) ([]response.BookingResponse, error) {
	var responseBooking []response.BookingResponse

	getData, err := b.BookingRepo.GetBookingBySession(id)
	if err != nil {
		return responseBooking, err
	}

	responseBooking = make([]response.BookingResponse, len(getData))

	for i, val := range getData {
		responseBooking[i] = response.BookingResponse{
			ID:        val.ID,
			IdSession: val.IdSession,
			Queue:     &val.Queue,
			Status:    val.Status,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		}
	}

	return responseBooking, nil
}

func (b *bookingService) GetBooking(id string) (response.BookingResponse, error) {
	var responseBooking response.BookingResponse

	getData, err := b.BookingRepo.GetBooking(id)
	if err != nil {
		return responseBooking, err
	}

	responseBooking = response.BookingResponse{
		ID:        getData.ID,
		IdSession: getData.IdSession,
		Queue:     &getData.Queue,
		Status:    getData.Status,
		CreatedAt: getData.CreatedAt,
		UpdatedAt: getData.UpdatedAt,
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
