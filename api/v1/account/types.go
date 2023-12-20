package account

type CreateAccountRequest struct {
	IdToken string `json:"idToken" binding:"required"`
}

type CreateAccountResponse struct {
	Token  string `json:"token"`
	UserId string `json:"userId"`
}
