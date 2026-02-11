package dto

type UserPayload struct {
	FirstNameTH string `json:"first_name_th" validate:"omitempty,max=100,regexp=^[\u0E00-\u0E4C .-]+$"`
	LastNameTH  string `json:"last_name_th" validate:"omitempty,max=100,regexp=^[\u0E00-\u0E4C .-]*$"`
	Dob         string `json:"birth_date" validate:"required,datetime=2006-01-02"`
	Addresses   string `json:"addresses" validate:"omitempty,max=1000,regexp=^.+$"`
}

type DataInsertUser struct {
	ID           string `json:"id"`
	CreateTimeAt string `json:"create_time_at"`
}

type ResponseDataDashBoardUser struct {
	AllUsers []DataDashBoardUser `json:"all_users"`
	Total int `json:"total"`
}
type DataDashBoardUser struct {
	ID          string `json:"id"`
	FirstNameTH string `json:"first_name_th"`
	LastNameTH  string `json:"last_name_th" `
	FullNameTH  string `json:"full_name_th"`
	Dob         string `json:"birth_date"`
	Addresses   string `json:"addresses"`
	Age         int    `json:"age"`
}