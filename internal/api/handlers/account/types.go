package account

const (
	AUTH_GOOGLE_APP = "google"
	AUTH_APPLE_APP = "apple"
)

type CreateAccountRequest struct {
	IdToken string `json:"idToken" binding:"required"`
}
type AppAccountRequest struct {
	Email string `json:"email" binding:"required"`
}
type AuthAppAccountRequest struct {
	AppleID  string `json:"appleId"`
	Email    string `json:"email"`
	AuthType string `json:"authType" binding:"required"`
}

type AppAccountRegisterApple struct {
	Email   string `json:"email" binding:"required"`
	AppleId string `json:"appleId" binding:"required"`
}

type CreateAccountResponse struct {
	Token  string `json:"token"`
	UserId string `json:"userId"`
}

type PasetoFromMagicLinkRequest struct {
	Code    string `json:"code" binding:"required"`
	EmailId string `json:"emailId" binding:"required"`
}

type PasetoFromMagicLinkResponse struct {
	Token string `json:"token"`
}

type GenerateAuthIdRequest struct {
	Email string `json:"email" binding:"required"`
}
