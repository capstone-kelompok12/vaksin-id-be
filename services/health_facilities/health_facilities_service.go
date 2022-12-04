package services

import (
	"vaksin-id-be/dto/payload"
	"vaksin-id-be/model"
	mysqla "vaksin-id-be/repository/mysql/addresses"
	mysqladm "vaksin-id-be/repository/mysql/admins"
	mysqlh "vaksin-id-be/repository/mysql/health_facilities"
	"vaksin-id-be/util"

	"github.com/google/uuid"
)

type HealthFacilitiesService interface {
	CreateHealthFacilities(payloads payload.HealthFacilities) error
	GetAllHealthFacilities() ([]model.HealthFacilities, error)
	GetHealthFacilities(name string) (model.HealthFacilities, error)
	UpdateHealthFacilities(payloads payload.UpdateHealthFacilities, id string) error
	DeleteHealthFacilities(id string) error
}

type healthFacilitiesService struct {
	HealthRepo  mysqlh.HealthFacilitiesRepository
	AddressRepo mysqla.AddressesRepository
	AdminRepo   mysqladm.AdminsRepository
}

func NewHealthFacilitiesService(healthRepo mysqlh.HealthFacilitiesRepository, addressRepo mysqla.AddressesRepository, adminRepo mysqladm.AdminsRepository) *healthFacilitiesService {
	return &healthFacilitiesService{
		HealthRepo:  healthRepo,
		AddressRepo: addressRepo,
		AdminRepo:   adminRepo,
	}
}

func (h *healthFacilitiesService) CreateHealthFacilities(payloads payload.HealthFacilities) error {

	idHealthFacil := uuid.NewString()
	idAddr := uuid.NewString()
	idAdmin := uuid.NewString()

	healthFacil := model.HealthFacilities{
		ID:       idHealthFacil,
		Email:    payloads.Email,
		PhoneNum: payloads.PhoneNum,
		Name:     payloads.Name,
		Image:    payloads.Image,
	}

	if err := h.HealthRepo.CreateHealthFacilities(healthFacil); err != nil {
		return err
	}

	address := model.Addresses{
		ID:                 idAddr,
		IdHealthFacilities: &idHealthFacil,
		NikUser:            nil,
		CurrentAddress:     payloads.CurrentAddress,
		District:           payloads.District,
		City:               payloads.City,
		Province:           payloads.Province,
		Longitude:          payloads.Longitude,
		Latitude:           payloads.Latitude,
	}

	if err := h.AddressRepo.CreateAddress(address); err != nil {
		return err
	}

	hashPass, err := util.HashPassword(payloads.PasswordAdmin)
	if err != nil {
		return err
	}

	adminModel := model.Admins{
		ID:                 idAdmin,
		IdHealthFacilities: idHealthFacil,
		Email:              payloads.EmailAdmin,
		Password:           hashPass,
	}

	err = h.AdminRepo.RegisterAdmins(adminModel)
	if err != nil {
		return err
	}
	return nil
}

func (h *healthFacilitiesService) GetAllHealthFacilities() ([]model.HealthFacilities, error) {
	allData, err := h.HealthRepo.GetAllHealthFacilities()
	if err != nil {
		return allData, err
	}

	return allData, err
}

func (h *healthFacilitiesService) GetHealthFacilities(name string) (model.HealthFacilities, error) {
	data, err := h.HealthRepo.GetHealthFacilities(name)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (h *healthFacilitiesService) UpdateHealthFacilities(payloads payload.UpdateHealthFacilities, id string) error {
	healthFacil := model.HealthFacilities{
		Email:    payloads.Email,
		PhoneNum: payloads.PhoneNum,
		Name:     payloads.Name,
		Image:    payloads.Image,
	}

	if err := h.HealthRepo.UpdateHealthFacilities(healthFacil, id); err != nil {
		return err
	}

	address := model.Addresses{
		CurrentAddress: payloads.CurrentAddress,
		District:       payloads.District,
		City:           payloads.City,
		Province:       payloads.Province,
		Longitude:      payloads.Longitude,
		Latitude:       payloads.Latitude,
	}

	if err := h.AddressRepo.UpdateAddressHealthFacilities(address, id); err != nil {
		return err
	}
	return nil
}

func (h *healthFacilitiesService) DeleteHealthFacilities(id string) error {
	if err := h.AddressRepo.DeleteAddressHealthFacilities(id); err != nil {
		return err
	}

	if err := h.AdminRepo.DeleteAdminsByHealth(id); err != nil {
		return err
	}

	if err := h.HealthRepo.DeleteHealthFacilities(id); err != nil {
		return err
	}

	return nil
}
