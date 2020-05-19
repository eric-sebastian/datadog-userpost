package datadog

type AddUserRequest struct {
	AccessRole string `json:"access_role"`
	Disabled   bool   `json:"disabled"`
	Email      string `json:"email"`
	Handle     string `json:"handle"`
	Name       string `json:"name"`
}

type User struct {
	Disabled   bool   `json:"disabled"`
	Handle     string `json:"handle"`
	Name       string `json:"name"`
	Title      string `json:"title"`
	IsAdmin    bool   `json:"is_admin"`
	Role       string `json:"role"`
	AccessRole string `json:"access_role"`
	Verified   bool   `json:"verified"`
	Email      string `json:"email"`
	Icon       string `json:"icon"`
}

type ErrorResponse struct {
	Status     string   `json:"status"`
	Code       int      `json:"code"`
	Errors     []string `json:"errors"`
	Statuspage string   `json:"statuspage"`
	Twitter    string   `json:"twitter"`
	Email      string   `json:"email"`
}

type CreateDatadogResponse struct {
	User       User     `json:"user"`
	StatusCode int      `json:"statuscode"`
	Errors     []string `json:"errors"`
}

type UserHandleRequest struct {
	Handle string `json:"handle"`
}

type DisableDatadogResponse struct {
	Message    string   `json:"message"`
	StatusCode int      `json:"statuscode"`
	Errors     []string `json:"errors"`
}

type GetUsersDatadogResponse struct {
	Users      []User   `json:"users"`
	StatusCode int      `json:"statuscode"`
	Errors     []string `json:"errors"`
}

type GetUserDatadogResponse struct {
	User       User     `json:"user"`
	StatusCode int      `json:"statuscode"`
	Errors     []string `json:"errors"`
}

type UpdateUserDatadogRequest struct {
	Handle     string         `json:"handle"`
	UpdateUser AddUserRequest `json:"updateuser"`
}
