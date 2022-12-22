package services

import (
	"vaksin-id-be/dto/payload"
	"vaksin-id-be/model"
	mysql "vaksin-id-be/repository/mysql/addresses"
)

type AddressService interface {
	GetAddressUser(nik string) (model.Addresses, error)
	UpdateUserAddress(payloads payload.UpdateAddress, nik string) (payload.UpdateAddress, error)
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

	dataAddress, err := a.AddressRepo.GetAddressUser(nik)
	if err != nil {
		return address, err
	}

	return dataAddress, nil
}

func (a *addressService) UpdateUserAddress(payloads payload.UpdateAddress, nik string) (payload.UpdateAddress, error) {
	var addressResp payload.UpdateAddress

	newAddress := model.Addresses{
		CurrentAddress: payloads.CurrentAddress,
		District:       payloads.District,
		City:           payloads.City,
		Province:       payloads.Province,
		Longitude:      payloads.Longitude,
		Latitude:       payloads.Latitude,
	}

	if err := a.AddressRepo.UpdateAddressUser(newAddress, nik); err != nil {
		return addressResp, err
	}

	addressResp = payloads
	return addressResp, nil
}
