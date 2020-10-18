package entity

type AuthResponse struct {
	Code         int    `json:"code"`
	Status       string `json:"status"`
	ErrorDetails string `json:"error_details"`
	ErrorType    string `json:"error_type"`
	Data         Data
}

type Data struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
