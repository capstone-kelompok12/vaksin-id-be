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
	UpdateCancelBooking(payloads payload.BookingCancel, nik string) (response.BookingResponseCustom, error)
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
	statusData := "OnProcess"

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
			NikUser:   val.NikUser,
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

		idSameBook, err := b.HistoryRepo.GetHistoryByNIK(val.ID, val.NikUser)
		if err != nil {
			return resBooking, err
		}

		sessionData := response.BookingSessionCustom{
			ID:           getSessionId.ID,
			IdVaccine:    getSessionId.IdVaccine,
			SessionName:  getSessionId.SessionName,
			CapacityLeft: countBook,
			Capacity:     getSession.Capacity,
			Dose:         getSessionId.Dose,
			Date:         getSessionId.Date,
			IsClose:      getSession.IsClose,
			StartSession: getSessionId.StartSession,
			EndSession:   getSessionId.EndSession,
			CreatedAt:    getSessionId.CreatedAt,
			UpdatedAt:    getSessionId.UpdatedAt,
			Vaccine:      getSession.Vaccine,
		}

		historyData := response.BookingHistoryCustom{
			ID:         idSameBook.ID,
			IdBooking:  idSameBook.IdBooking,
			NikUser:    idSameBook.NikUser,
			IdSameBook: idSameBook.IdSameBook,
			Status:     idSameBook.Status,
			CreatedAt:  idSameBook.CreatedAt,
			UpdatedAt:  idSameBook.UpdatedAt,
			User:       idSameBook.User,
		}

		resBooking[i] = response.BookingResponseCustom{
			ID:        val.ID,
			IdSession: val.IdSession,
			NikUser:   val.NikUser,
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

	if payloads[0].Status == "Rejected" {
		for i, v := range payloads {

			getSession, err := b.SessionRepo.GetSessionById(v.IdSession)
			if err != nil {
				return data, err
			}

			if v.Status != "Accepted" && v.Status != "Rejected" {
				return data, errors.New("input status must Accepted or Rejected")
			}
			bookingData[i] = model.BookingSessions{
				ID:        v.ID,
				IdSession: v.IdSession,
				NikUser:   v.NikUser,
				Status:    &v.Status,
			}

			updateData, err := b.BookingRepo.UpdateBookingAcc(bookingData[i])
			if err != nil {
				return data, err
			}

			statusHistory := "Absent"

			dataHistoryUpdate := model.VaccineHistories{
				Status: &statusHistory,
			}

			idSameBook, err := b.HistoryRepo.GetHistoryByNIK(v.ID, v.NikUser)
			if err != nil {
				return data, err
			}

			getBookedByIdSession, err := b.BookingRepo.GetBookingBySessionDen(v.IdSession)
			if err != nil {
				return data, err
			}
			getbackCap = getSession.Capacity - len(getBookedByIdSession)

			_, err = b.HistoryRepo.UpdateHistoryByNik(dataHistoryUpdate, v.NikUser, v.ID)
			if err != nil {
				return data, err
			}

			sessionData := response.BookingSessionCustom{
				ID:           getSession.ID,
				IdVaccine:    getSession.IdVaccine,
				SessionName:  getSession.SessionName,
				CapacityLeft: getbackCap,
				Capacity:     getSession.Capacity,
				Dose:         getSession.Dose,
				Date:         getSession.Date,
				IsClose:      getSession.IsClose,
				StartSession: getSession.StartSession,
				EndSession:   getSession.EndSession,
				CreatedAt:    getSession.CreatedAt,
				UpdatedAt:    getSession.UpdatedAt,
				Vaccine:      getSession.Vaccine,
			}

			historyData := response.BookingHistoryCustom{
				ID:         idSameBook.ID,
				IdBooking:  idSameBook.IdBooking,
				NikUser:    idSameBook.NikUser,
				IdSameBook: idSameBook.IdSameBook,
				Status:     idSameBook.Status,
				CreatedAt:  idSameBook.CreatedAt,
				UpdatedAt:  idSameBook.UpdatedAt,
				User:       idSameBook.User,
			}

			newResReject[i] = response.BookingResponseCustom{
				ID:        updateData.ID,
				IdSession: updateData.IdSession,
				NikUser:   updateData.NikUser,
				Queue:     &updateData.Queue,
				Status:    updateData.Status,
				CreatedAt: updateData.CreatedAt,
				UpdatedAt: updateData.UpdatedAt,
				Session:   sessionData,
				History:   historyData,
			}
		}

		return newResReject, nil
	}

	for i, v := range payloads {

		getSession, err := b.SessionRepo.GetSessionById(v.IdSession)
		if err != nil {
			return data, err
		}

		if v.Status != "Accepted" && v.Status != "Rejected" {
			return data, errors.New("input status must Accepted or Rejected")
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
			NikUser:   v.NikUser,
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

		getBookedByIdSession, err := b.BookingRepo.GetBookingBySessionDen(v.IdSession)
		if err != nil {
			return data, err
		}

		getbackCap = getSession.Capacity - len(getBookedByIdSession)

		sessionData := response.BookingSessionCustom{
			ID:           getSession.ID,
			IdVaccine:    getSession.IdVaccine,
			SessionName:  getSession.SessionName,
			CapacityLeft: getbackCap,
			Capacity:     getSession.Capacity,
			Dose:         getSession.Dose,
			Date:         getSession.Date,
			IsClose:      getSession.IsClose,
			StartSession: getSession.StartSession,
			EndSession:   getSession.EndSession,
			CreatedAt:    getSession.CreatedAt,
			UpdatedAt:    getSession.UpdatedAt,
			Vaccine:      getSession.Vaccine,
		}

		historyData := response.BookingHistoryCustom{
			ID:         idSameBook.ID,
			IdBooking:  idSameBook.IdBooking,
			NikUser:    idSameBook.NikUser,
			IdSameBook: idSameBook.IdSameBook,
			Status:     idSameBook.Status,
			CreatedAt:  idSameBook.CreatedAt,
			UpdatedAt:  idSameBook.UpdatedAt,
			User:       idSameBook.User,
		}

		newRes[i] = response.BookingResponseCustom{
			ID:        updateData.ID,
			IdSession: updateData.IdSession,
			NikUser:   updateData.NikUser,
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

		if err := b.UserRepo.UpdateAccUserProfile(updateUserCount); err != nil {
			return updateData, err
		}

		dataHistory, err := b.HistoryRepo.GetHistoryByIdNoUserHistory(val.ID)
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

func (b *bookingService) UpdateCancelBooking(payloads payload.BookingCancel, nik string) (response.BookingResponseCustom, error) {
	var responseData response.BookingResponseCustom

	statusBooking := "Rejected"
	statusHistory := "Absent"

	updateModelBooking := model.BookingSessions{
		ID:        payloads.ID,
		IdSession: payloads.IdSession,
		NikUser:   nik,
		Status:    &statusBooking,
	}

	err := b.BookingRepo.UpdateBooking(updateModelBooking)
	if err != nil {
		return responseData, err
	}

	updateModelHistory := model.VaccineHistories{
		Status: &statusHistory,
	}

	_, err = b.HistoryRepo.UpdateHistoryByNik(updateModelHistory, nik, payloads.ID)
	if err != nil {
		return responseData, err
	}

	dataBooking, err := b.BookingRepo.GetBooking(payloads.ID)
	if err != nil {
		return responseData, err
	}

	dataSession, err := b.SessionRepo.GetSessionById(payloads.IdSession)
	if err != nil {
		return responseData, err
	}

	countHistory, err := b.HistoryRepo.GetHistoryByNIK(payloads.ID, nik)
	if err != nil {
		return responseData, err
	}

	getSession, err := b.SessionRepo.GetSessionById(payloads.IdSession)
	if err != nil {
		return responseData, err
	}

	getBookedByIdSession, err := b.BookingRepo.GetBookingBySessionDen(payloads.IdSession)
	if err != nil {
		return responseData, err
	}
	getbackCap := getSession.Capacity - len(getBookedByIdSession)

	sessionData := response.BookingSessionCustom{
		ID:           dataSession.ID,
		IdVaccine:    dataSession.IdVaccine,
		SessionName:  dataSession.SessionName,
		Capacity:     dataSession.Capacity,
		CapacityLeft: getbackCap,
		Dose:         dataSession.Dose,
		Date:         dataSession.Date,
		IsClose:      dataSession.IsClose,
		StartSession: dataSession.StartSession,
		EndSession:   dataSession.EndSession,
		CreatedAt:    dataSession.CreatedAt,
		UpdatedAt:    dataSession.UpdatedAt,
		Vaccine:      dataSession.Vaccine,
	}

	historyData := response.BookingHistoryCustom{
		ID:         countHistory.ID,
		IdBooking:  countHistory.IdBooking,
		NikUser:    countHistory.NikUser,
		IdSameBook: countHistory.IdSameBook,
		Status:     countHistory.Status,
		CreatedAt:  countHistory.CreatedAt,
		UpdatedAt:  countHistory.UpdatedAt,
		User:       countHistory.User,
	}

	responseData = response.BookingResponseCustom{
		ID:        dataBooking.ID,
		IdSession: dataBooking.IdSession,
		Queue:     &dataBooking.Queue,
		Status:    dataBooking.Status,
		CreatedAt: dataBooking.CreatedAt,
		UpdatedAt: dataBooking.UpdatedAt,
		Session:   sessionData,
		History:   historyData,
	}
	return responseData, nil
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
