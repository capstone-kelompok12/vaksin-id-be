package services

import (
	"vaksin-id-be/dto/payload"
	"vaksin-id-be/model"
	mysqla "vaksin-id-be/repository/mysql/addresses"
	mysqlh "vaksin-id-be/repository/mysql/health_facilities"

	"github.com/google/uuid"
)

type HealthFacilitiesService interface {
	CreateHealthFacilities(payloads payload.HealthFacilities) error
	//GetAllHealthFacilities() ([]payload.HealthFacilitiesRegister, error)       // wait admin and vaksin feature
	//GetHealthFacilities(name string) (payload.HealthFacilitiesRegister, error) // wait admin and vaksin feature
	UpdateHealthFacilities(payloads payload.HealthFacilities, id string) error
	DeleteHealthFacilities(id string) error
}

type healthFacilitiesService struct {
	HealthRepo  mysqlh.HealthFacilitiesRepository
	AddressRepo mysqla.AddressesRepository
}

func NewHealthFacilitiesService(healthRepo mysqlh.HealthFacilitiesRepository, addressRepo mysqla.AddressesRepository) *healthFacilitiesService {
	return &healthFacilitiesService{
		HealthRepo:  healthRepo,
		AddressRepo: addressRepo,
	}
}

func (h *healthFacilitiesService) CreateHealthFacilities(payloads payload.HealthFacilities) error {

	idHealthFacil := uuid.NewString()
	idAddr := uuid.NewString()

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
	return nil
}

func (h *healthFacilitiesService) UpdateHealthFacilities(payloads payload.HealthFacilities, id string) error {
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

	if err := h.HealthRepo.DeleteHealthFacilities(id); err != nil {
		return err
	}

	return nil
}
