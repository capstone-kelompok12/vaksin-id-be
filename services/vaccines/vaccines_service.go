package services

import (
	"errors"
	"vaksin-id-be/dto/payload"
	"vaksin-id-be/dto/response"
	m "vaksin-id-be/middleware"
	"vaksin-id-be/model"
	mysqlv "vaksin-id-be/repository/mysql/vaccines"

	"github.com/google/uuid"
)

type VaccinesService interface {
	CreateVaccine(authAdmin string, payloads payload.VaccinesPayload) (model.Vaccines, error)
	GetAllVaccines() ([]response.VaccinesResponse, error)
	GetVaccineByAdmin(idhealthfacilities string) ([]model.Vaccines, error)
	UpdateVaccine(id string, payloads payload.VaccinesUpdatePayload) (response.VaccinesResponse, error)
	DeleteVacccine(id string) error
}

type vaccinesService struct {
	VaccinesRepo mysqlv.VaccinesRepository
}

func NewVaccinesService(vaccinesRepo mysqlv.VaccinesRepository) *vaccinesService {
	return &vaccinesService{
		VaccinesRepo: vaccinesRepo,
	}
}

func (v *vaccinesService) CreateVaccine(authAdmin string, payloads payload.VaccinesPayload) (model.Vaccines, error) {
	var vaccineModel model.Vaccines

	id := uuid.NewString()

	idHealthFacilities, err := m.GetIdHealthFacilities(authAdmin)
	if err != nil {
		return vaccineModel, err
	}

	if err := v.VaccinesRepo.CheckNameExist(idHealthFacilities, payloads.Name); err == nil {
		return vaccineModel, errors.New("vaccine name already exist")
	}

	vaccineModel = model.Vaccines{
		ID:                 id,
		IdHealthFacilities: idHealthFacilities,
		Name:               payloads.Name,
		Stock:              payloads.Stock,
	}

	err = v.VaccinesRepo.CreateVaccine(vaccineModel)
	if err != nil {
		return vaccineModel, err
	}

	return vaccineModel, nil
}

func (v *vaccinesService) GetAllVaccines() ([]response.VaccinesResponse, error) {
	var vaccinesResponse []response.VaccinesResponse

	getVaccine, err := v.VaccinesRepo.GetAllVaccines()

	if err != nil {
		return vaccinesResponse, err
	}

	vaccinesResponse = make([]response.VaccinesResponse, len(getVaccine))

	for i, v := range getVaccine {
		vaccinesResponse[i] = response.VaccinesResponse{
			ID:    v.ID,
			Name:  v.Name,
			Stock: v.Stock,
		}
	}

	return vaccinesResponse, nil
}

func (v *vaccinesService) GetVaccineByAdmin(idhealthfacilities string) ([]model.Vaccines, error) {
	var vaccines []model.Vaccines
	idHealthFacilities, err := m.GetIdHealthFacilities(idhealthfacilities)
	if err != nil {
		return vaccines, err
	}

	vaccines, err = v.VaccinesRepo.GetVaccinesByIdAdmin(idHealthFacilities)
	if err != nil {
		return vaccines, err
	}

	return vaccines, nil
}

func (v *vaccinesService) UpdateVaccine(id string, payloads payload.VaccinesUpdatePayload) (response.VaccinesResponse, error) {
	var dataResp response.VaccinesResponse

	vaccineData := model.Vaccines{
		Name:  payloads.Name,
		Stock: payloads.Stock,
	}

	if err := v.VaccinesRepo.UpdateVaccine(vaccineData, id); err != nil {
		return dataResp, err
	}

	dataResp = response.VaccinesResponse{
		ID:    id,
		Name:  payloads.Name,
		Stock: payloads.Stock,
	}

	return dataResp, nil
}

func (v *vaccinesService) DeleteVacccine(id string) error {
	if err := v.VaccinesRepo.DeleteVacccine(id); err != nil {
		return err
	}

	return nil
}
