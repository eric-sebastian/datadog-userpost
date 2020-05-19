package datadog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Domain interface {
	CreateUser(newUser AddUserRequest) (CreateDatadogResponse, error)
	DisableUser(usrHandle string) (DisableDatadogResponse, error)
	GetAllUser() (GetUsersDatadogResponse, error)
	GetUser(usrHandle string) (GetUserDatadogResponse, error)
	UpdateUser(usrHandle string, updateUser AddUserRequest) (CreateDatadogResponse, error)
}

type DomainImplementation struct {
	APIKey         string `json:"api_key"`
	ApplicationKey string `json:"application_key"`
}

func NewDomainImplementation(APIKey string, ApplicationKey string) *DomainImplementation {
	return &DomainImplementation{
		APIKey:         APIKey,
		ApplicationKey: ApplicationKey,
	}
}

func (d *DomainImplementation) CreateUser(newUser AddUserRequest) (CreateDatadogResponse, error) {

	var resp CreateDatadogResponse

	url := fmt.Sprintf("https://api.datadoghq.com/api/v1/user")

	newUsr := &newUser
	out, err := json.Marshal(newUsr)

	if err != nil {
		log.Println(err)
		return resp, err
	}
	payload := strings.NewReader(string(out))

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		log.Println(err)
		return resp, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("DD-API-KEY", d.APIKey)
	req.Header.Add("DD-APPLICATION-KEY", d.ApplicationKey)

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return resp, err
	}
	defer res.Body.Close()

	resp.StatusCode = res.StatusCode
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return resp, err
	}
	err = json.Unmarshal([]byte(string(body)), &resp)
	if err != nil {
		log.Println(err)
		return resp, err
	}

	return resp, nil
}

func (d *DomainImplementation) DisableUser(usrHandle string) (DisableDatadogResponse, error) {

	var resp DisableDatadogResponse

	url := fmt.Sprintf("https://api.datadoghq.com/api/v1/user/" + usrHandle)

	payload := strings.NewReader("")

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, payload)
	if err != nil {
		log.Println(err)
		return resp, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("DD-API-KEY", d.APIKey)
	req.Header.Add("DD-APPLICATION-KEY", d.ApplicationKey)

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return resp, err
	}
	defer res.Body.Close()

	resp.StatusCode = res.StatusCode
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return resp, err
	}
	err = json.Unmarshal([]byte(string(body)), &resp)
	if err != nil {
		log.Println(err)
		return resp, err
	}

	return resp, nil
}

func (d *DomainImplementation) GetAllUser() (GetUsersDatadogResponse, error) {

	var resp GetUsersDatadogResponse

	url := fmt.Sprintf("https://api.datadoghq.com/api/v1/user")

	payload := strings.NewReader("")

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, payload)

	if err != nil {
		log.Println(err)
		return resp, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("DD-API-KEY", d.APIKey)
	req.Header.Add("DD-APPLICATION-KEY", d.ApplicationKey)

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return resp, err
	}
	defer res.Body.Close()

	resp.StatusCode = res.StatusCode
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return resp, err
	}
	err = json.Unmarshal([]byte(string(body)), &resp)
	if err != nil {
		log.Println(err)
		return resp, err
	}

	return resp, nil
}

func (d *DomainImplementation) GetUser(usrHandle string) (GetUserDatadogResponse, error) {

	var resp GetUserDatadogResponse

	url := fmt.Sprintf("https://api.datadoghq.com/api/v1/user/" + usrHandle)

	payload := strings.NewReader("")

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, payload)

	if err != nil {
		log.Println(err)
		return resp, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("DD-API-KEY", d.APIKey)
	req.Header.Add("DD-APPLICATION-KEY", d.ApplicationKey)

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return resp, err
	}
	defer res.Body.Close()

	resp.StatusCode = res.StatusCode
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return resp, err
	}
	err = json.Unmarshal([]byte(string(body)), &resp)
	if err != nil {
		log.Println(err)
		return resp, err
	}

	return resp, nil
}

func (d *DomainImplementation) UpdateUser(usrHandle string, updateUser AddUserRequest) (CreateDatadogResponse, error) {

	var resp CreateDatadogResponse

	url := fmt.Sprintf("https://api.datadoghq.com/api/v1/user/" + usrHandle)

	updateUsr := &updateUser
	out, err := json.Marshal(updateUsr)

	if err != nil {
		log.Println(err)
		return resp, err
	}
	payload := strings.NewReader(string(out))

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, payload)

	if err != nil {
		log.Println(err)
		return resp, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("DD-API-KEY", d.APIKey)
	req.Header.Add("DD-APPLICATION-KEY", d.ApplicationKey)

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return resp, err
	}
	defer res.Body.Close()

	resp.StatusCode = res.StatusCode
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return resp, err
	}
	err = json.Unmarshal([]byte(string(body)), &resp)
	if err != nil {
		log.Println(err)
		return resp, err
	}

	return resp, nil
}
