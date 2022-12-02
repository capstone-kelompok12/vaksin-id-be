package services

import (
	"vaksin-id-be/dto/payload"
	"vaksin-id-be/dto/response"
	"vaksin-id-be/model"
	mysqlv "vaksin-id-be/repository/mysql/vaccines"

	"github.com/google/uuid"
)

type AdminService interface {
	CreateVaccine(payloads payload.VaccinesPayload) error
	GetAllVaccines() ([]response.VaccinesResponse, error)
	UpdateVaccine(payloads payload.VaccinesPayload, id string) error
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

func (v *vaccinesService) CreateVaccine(payloads payload.VaccinesPayload) error {
	id := uuid.NewString()

	vaccineModel := model.Vaccines{
		ID:                 id,
		IdHealthFacilities: payloads.IdHealthFacilities,
		Name:               payloads.Name,
		Stock:              payloads.Stock,
	}

	err := v.VaccinesRepo.CreateVaccine(vaccineModel)
	if err != nil {
		return err
	}

	return nil
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

func (v *vaccinesService) UpdateVaccine(payloads payload.VaccinesPayload, id string) error {
	vaccineData := model.Vaccines{
		Name:  payloads.Name,
		Stock: payloads.Stock,
	}

	if err := v.VaccinesRepo.UpdateVaccine(vaccineData, id); err != nil {
		return err
	}
	return nil
}

func (v *vaccinesService) DeleteVacccine(id string) error {
	if err := v.VaccinesRepo.DeleteVacccine(id); err != nil {
		return err
	}

	return nil
}
