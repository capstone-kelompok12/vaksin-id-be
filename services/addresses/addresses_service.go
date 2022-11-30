package services

import (
	"vaksin-id-be/dto/payload"
	m "vaksin-id-be/middleware"
	"vaksin-id-be/model"
	mysql "vaksin-id-be/repository/mysql/addresses"
)

type AddressService interface {
	GetAddressUser(nik string) (model.Addresses, error)
	UpdateUserAddress(payloads payload.UpdateAddress, nik string) error
}

type addressService struct {
	AddressRepo mysql.AddressesRepository
}

func NewAddressesService(addressRepo mysql.AddressesRepository) *addressService {
	return &addressService{
		AddressRepo: addressRepo,
	}
}

func (a *addressService) GetAddressUser(nik string) (model.Addresses, error) {
	var address model.Addresses

	getUserNik, err := m.GetUserNik(nik)
	if err != nil {
		return address, err
	}

	dataAddress, err := a.AddressRepo.GetAddressUser(getUserNik)
	if err != nil {
		return address, err
	}

	return dataAddress, nil
}

func (a *addressService) UpdateUserAddress(payloads payload.UpdateAddress, nik string) error {

	getUserNik, err := m.GetUserNik(nik)
	if err != nil {
		return err
	}

	newAddress := model.Addresses{
		CurrentAddress: payloads.CurrentAddress,
		District:       payloads.District,
		City:           payloads.City,
		Province:       payloads.Province,
		Longitude:      payloads.Longitude,
		Latitude:       payloads.Latitude,
	}

	if err := a.AddressRepo.UpdateAddressUser(newAddress, getUserNik); err != nil {
		return err
	}
	return nil
}
