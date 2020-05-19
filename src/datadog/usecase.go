package datadog

import (
	"log"
)

type Usecase interface {
	CreateUser(req AddUserRequest) (CreateDatadogResponse, error)
	DisableUser(usrHandle UserHandleRequest) (DisableDatadogResponse, error)
	GetAllUser() (GetUsersDatadogResponse, error)
	GetUser(usrHandle UserHandleRequest) (GetUserDatadogResponse, error)
	UpdateUser(req UpdateUserDatadogRequest) (CreateDatadogResponse, error)
}

type UsecaseImplementation struct {
	datadogDomain Domain
}

func NewUsecaseImplementation(datadogDomain Domain) *UsecaseImplementation {
	return &UsecaseImplementation{
		datadogDomain: datadogDomain,
	}
}

func (uc *UsecaseImplementation) CreateUser(req AddUserRequest) (CreateDatadogResponse, error) {

	resp, err := uc.datadogDomain.CreateUser(req)
	if err != nil {
		log.Println(err)
		return resp, err
	}
	return resp, err
}

func (uc *UsecaseImplementation) DisableUser(usrHandle UserHandleRequest) (DisableDatadogResponse, error) {

	resp, err := uc.datadogDomain.DisableUser(usrHandle.Handle)
	if err != nil {
		log.Println(err)
		return resp, err
	}
	return resp, err
}

func (uc *UsecaseImplementation) GetAllUser() (GetUsersDatadogResponse, error) {

	resp, err := uc.datadogDomain.GetAllUser()
	if err != nil {
		log.Println(err)
		return resp, err
	}
	return resp, err
}

func (uc *UsecaseImplementation) GetUser(usrHandle UserHandleRequest) (GetUserDatadogResponse, error) {

	resp, err := uc.datadogDomain.GetUser(usrHandle.Handle)
	if err != nil {
		log.Println(err)
		return resp, err
	}
	return resp, err
}

func (uc *UsecaseImplementation) UpdateUser(req UpdateUserDatadogRequest) (CreateDatadogResponse, error) {

	resp, err := uc.datadogDomain.UpdateUser(req.Handle, req.UpdateUser)
	if err != nil {
		log.Println(err)
		return resp, err
	}
	return resp, err
}
