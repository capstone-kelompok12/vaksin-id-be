package services

import (
	"errors"
	"vaksin-id-be/dto/payload"
	"vaksin-id-be/dto/response"
	m "vaksin-id-be/middleware"
	"vaksin-id-be/model"
	mysqlad "vaksin-id-be/repository/mysql/admins"
	"vaksin-id-be/util"

	"github.com/google/uuid"
)

type AdminService interface {
	RegisterAdmin(payloads payload.AdminsPayload) error
	LoginAdmin(payloads payload.LoginAdmin) (response.Login, error)
	GetAllAdmin() ([]response.AdminsResponse, error)
	GetAdmin(id string) (response.AdminsResponse, error)
	UpdateAdmin(payloads payload.AdminsPayload, id string) error
	DeleteAdmin(id string) error
}

type adminService struct {
	AdminRepo mysqlad.AdminRepository
}

func NewAdminService(adminServ mysqlad.AdminRepository) *adminService {
	return &adminService{
		AdminRepo: adminServ,
	}
}

func (a *adminService) RegisterAdmin(payloads payload.AdminsPayload) error {
	id := uuid.NewString()

	passAdmin, err := util.HashPassword(payloads.Password)
	if err != nil {
		return err
	}

	dataAdmin := model.Admins{
		ID:                 id,
		IdHealthFacilities: payloads.IdHealthFacilities,
		Email:              payloads.Email,
		Password:           passAdmin,
	}

	if err := a.AdminRepo.RegisterAdmin(dataAdmin); err != nil {
		return err
	}

	return nil
}

func (a *adminService) LoginAdmin(payloads payload.LoginAdmin) (response.Login, error) {
	var loginResponse response.Login

	adminModel := model.Admins{
		Email:    payloads.Email,
		Password: payloads.Password,
	}

	userData, err := a.AdminRepo.LoginAdmin(adminModel)
	if err != nil {
		return loginResponse, err
	}

	isValid := util.CheckPasswordHash(payloads.Password, userData.Password)
	if !isValid {
		return loginResponse, errors.New("wrong password")
	}

	token, errToken := m.CreateTokenAdmin(userData.ID, userData.Email)

	if errToken != nil {
		return loginResponse, err
	}

	loginResponse = response.Login{
		Token: token,
	}

	return loginResponse, nil
}

func (a *adminService) GetAllAdmin() ([]response.AdminsResponse, error) {
	allAdmins, err := a.AdminRepo.GetAllAdmin()
	responses := make([]response.AdminsResponse, len(allAdmins))

	if err != nil {
		return responses, err
	}

	for i, val := range allAdmins {
		responses[i] = response.AdminsResponse{
			ID:                 val.ID,
			IdHealthFacilities: val.IdHealthFacilities,
			Email:              val.Email,
			CreatedAt:          val.CreatedAt,
			UpdatedAt:          val.UpdatedAt,
		}
	}

	return responses, nil
}

func (a *adminService) GetAdmin(id string) (response.AdminsResponse, error) {
	var responses response.AdminsResponse

	dataAdmin, err := a.AdminRepo.GetAdmin(id)
	if err != nil {
		return responses, err
	}

	responses = response.AdminsResponse{
		ID:                 dataAdmin.ID,
		IdHealthFacilities: dataAdmin.IdHealthFacilities,
		Email:              dataAdmin.Email,
		CreatedAt:          dataAdmin.CreatedAt,
		UpdatedAt:          dataAdmin.UpdatedAt,
	}

	return responses, nil
}

func (a *adminService) UpdateAdmin(payloads payload.AdminsPayload, id string) error {

	passAdmin, err := util.HashPassword(payloads.Password)
	if err != nil {
		return err
	}

	dataAdmin := model.Admins{
		IdHealthFacilities: payloads.IdHealthFacilities,
		Email:              payloads.Email,
		Password:           passAdmin,
	}

	if err := a.AdminRepo.UpdateAdmin(dataAdmin, id); err != nil {
		return err
	}

	return nil
}

func (a *adminService) DeleteAdmin(id string) error {

	if err := a.AdminRepo.DeleteAdmin(id); err != nil {
		return err
	}

	return nil
}
