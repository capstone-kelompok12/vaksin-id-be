package services

import (
	"fmt"
	"vaksin-id-be/dto/payload"
	"vaksin-id-be/dto/response"
	"vaksin-id-be/model"
	mysqlv "vaksin-id-be/repository/mysql/vaccines"

	"github.com/google/uuid"
)

type VaccinesService interface {
	CreateVaccine(idhealth string, payloads payload.VaccinesPayload) (model.Vaccines, error)
	GetAllVaccines() ([]response.VaccinesResponse, error)
	GetVaccineByAdmin(idhealthfacilities string) ([]model.Vaccines, error)
	GetVaccineDashboard() ([]response.DashboardVaccine, error)
	GetVaccinesCount() ([]response.VaccinesStockResponse, error)
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

func (v *vaccinesService) CreateVaccine(idhealth string, payloads payload.VaccinesPayload) (model.Vaccines, error) {
	var vaccineModel model.Vaccines

	id := uuid.NewString()

	vaccineData, err := v.VaccinesRepo.CheckNameDosisExist(idhealth, payloads.Name, payloads.Dose)
	if err == nil {
		fmt.Println(vaccineData.ID)
		addStock := payloads.Stock + vaccineData.Stock
		payloadUpdate := payload.VaccinesUpdatePayload{
			Stock: addStock,
		}

		dataUpdate, err := v.UpdateVaccine(vaccineData.ID, payloadUpdate)
		if err != nil {
			return vaccineModel, err
		}

		vaccineModel = model.Vaccines{
			ID:                 dataUpdate.ID,
			IdHealthFacilities: idhealth,
			Name:               payloads.Name,
			Stock:              addStock,
			Dose:               payloads.Dose,
			CreatedAt:          vaccineData.CreatedAt,
			UpdatedAt:          vaccineData.UpdatedAt,
		}

		return vaccineModel, nil
	}

	vaccineModel = model.Vaccines{
		ID:                 id,
		IdHealthFacilities: idhealth,
		Name:               payloads.Name,
		Dose:               payloads.Dose,
		Stock:              payloads.Stock,
	}

	err = v.VaccinesRepo.CreateVaccine(vaccineModel)
	if err != nil {
		return vaccineModel, err
	}

	return vaccineModel, nil
}

func (v *vaccinesService) GetVaccinesCount() ([]response.VaccinesStockResponse, error) {
	var vaccinesResponse []response.VaccinesStockResponse

	getName, err := v.VaccinesRepo.GetVaccineByName()
	if err != nil {
		return vaccinesResponse, err
	}

	vaccinesResponse = make([]response.VaccinesStockResponse, len(getName))

	for i, val := range getName {
		vaccinesResponse[i] = response.VaccinesStockResponse{
			Name:  val.Name,
			Stock: val.Stock,
		}
	}

	return vaccinesResponse, nil
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
			Dose:  v.Dose,
			Stock: v.Stock,
		}
	}

	return vaccinesResponse, nil
}

func (v *vaccinesService) GetVaccineByAdmin(idhealthfacilities string) ([]model.Vaccines, error) {
	var vaccines []model.Vaccines

	vaccines, err := v.VaccinesRepo.GetVaccinesByIdAdmin(idhealthfacilities)
	if err != nil {
		return vaccines, err
	}

	return vaccines, nil
}

func (v *vaccinesService) UpdateVaccine(id string, payloads payload.VaccinesUpdatePayload) (response.VaccinesResponse, error) {
	var dataResp response.VaccinesResponse

	vaccineData := model.Vaccines{
		Name:  payloads.Name,
		Dose:  payloads.Dose,
		Stock: payloads.Stock,
	}

	if err := v.VaccinesRepo.UpdateVaccine(vaccineData, id); err != nil {
		return dataResp, err
	}

	dataResp = response.VaccinesResponse{
		ID:    id,
		Name:  payloads.Name,
		Dose:  payloads.Dose,
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

func (v *vaccinesService) GetVaccineDashboard() ([]response.DashboardVaccine, error) {
	var vaccinesResponse []response.DashboardVaccine

	getVaccine, err := v.VaccinesRepo.GetAllVaccines()

	if err != nil {
		return vaccinesResponse, err
	}

	vaccinesResponse = make([]response.DashboardVaccine, len(getVaccine))

	for i, v := range getVaccine {
		vaccinesResponse[i] = response.DashboardVaccine{
			Name: v.Name,
			Dose: v.Dose,
		}
	}

	return vaccinesResponse, nil
}
