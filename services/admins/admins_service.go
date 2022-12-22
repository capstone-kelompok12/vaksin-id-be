package services

import (
	"errors"
	"vaksin-id-be/dto/payload"
	"vaksin-id-be/dto/response"
	m "vaksin-id-be/middleware"
	"vaksin-id-be/model"
	mysqla "vaksin-id-be/repository/mysql/admins"
	"vaksin-id-be/util"
)

type AdminService interface {
	LoginAdmin(payloads payload.Login) (response.Login, error)
	GetAllAdmins() ([]response.AdminResponse, error)
	GetAdmins(id string) (response.AdminResponse, error)
	UpdateAdmins(payloads payload.AdminsPayload, id string) (response.AdminProfileResponse, error)
	DeleteAdmins(id string) error
}

type adminService struct {
	AdminRepo mysqla.AdminsRepository
}

func NewAdminService(adminRepo mysqla.AdminsRepository) *adminService {
	return &adminService{
		AdminRepo: adminRepo,
	}
}

func (a *adminService) LoginAdmin(payloads payload.Login) (response.Login, error) {
	var loginResponse response.Login

	adminModel := model.Admins{
		Email:    payloads.Email,
		Password: payloads.Password,
	}

	adminData, err := a.AdminRepo.LoginAdmins(adminModel)
	if err != nil {
		return loginResponse, err
	}

	isValid := util.CheckPasswordHash(payloads.Password, adminData.Password)
	if !isValid {
		return loginResponse, errors.New("wrong password")
	}

	token, errToken := m.CreateTokenAdmin(adminData.ID, adminData.IdHealthFacilities, adminData.Email)

	if errToken != nil {
		return loginResponse, err
	}

	loginResponse = response.Login{
		Token: token,
	}

	return loginResponse, nil
}

func (a *adminService) GetAllAdmins() ([]response.AdminResponse, error) {
	var adminResponse []response.AdminResponse

	getData, err := a.AdminRepo.GetAllAdmins()
	if err != nil {
		return adminResponse, err
	}

	adminResponse = make([]response.AdminResponse, len(getData))

	for i, data := range getData {
		adminResponse[i] = response.AdminResponse{
			ID:                 data.ID,
			IdHealthFacilities: data.IdHealthFacilities,
			Email:              data.Email,
			CreatedAt:          data.CreatedAt,
			UpdatedAt:          data.UpdatedAt,
		}
	}

	return adminResponse, nil
}

func (a *adminService) GetAdmins(id string) (response.AdminResponse, error) {
	var responseAdmin response.AdminResponse

	getData, err := a.AdminRepo.GetAdmins(id)

	if err != nil {
		return responseAdmin, err
	}

	responseAdmin = response.AdminResponse{
		ID:                 getData.ID,
		IdHealthFacilities: getData.IdHealthFacilities,
		Email:              getData.Email,
		CreatedAt:          getData.CreatedAt,
		UpdatedAt:          getData.UpdatedAt,
	}

	return responseAdmin, nil
}

func (a *adminService) UpdateAdmins(payloads payload.AdminsPayload, id string) (response.AdminProfileResponse, error) {
	var dataResp response.AdminProfileResponse
	hashPass, err := util.HashPassword(payloads.Password)
	if err != nil {
		return dataResp, err
	}

	adminData := model.Admins{
		IdHealthFacilities: payloads.IdHealthFacilities,
		Email:              payloads.Email,
		Password:           hashPass,
	}

	dataResp = response.AdminProfileResponse{
		ID:    id,
		Email: payloads.Email,
	}

	if err := a.AdminRepo.UpdateAdmins(adminData, id); err != nil {
		return dataResp, err
	}
	return dataResp, nil
}

func (a *adminService) DeleteAdmins(id string) error {

	if err := a.AdminRepo.DeleteAdmins(id); err != nil {
		return err
	}

	return nil
}
