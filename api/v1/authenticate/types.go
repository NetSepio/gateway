package authenticate

type AuthenticateRequest struct {
	FlowId       string `json:"flowId" binding:"required"`
	Signature    string `json:"signature" binding:"omitempty,hexadecimal"`
	PubKey       string `json:"pubKey" binding:"omitempty"`
	SignatureSui string `json:"signatureSui" binding:"omitempty"`
	Message      string `json:"message" binding:"omitempty"`
	AccessToken  string `json:"accessToken" binding:"omitempty"`
	IdToken      string `json:"idToken" binding:"omitempty"` // to be pass in bearer token [ AUTHORIZATION KEY ]
	ChainName    string `json:"chain_name" binding:"required"`
}
type AuthenticateRequestNoSign struct {
	FlowId        string `json:"flowId" binding:"required"`
	WalletAddress string `json:"walletAddress" binding:"required"`
}

type AuthenticatePayload struct {
	Token  string `json:"token"`
	UserId string `json:"userId"`
}

type AuthenticateTokenPayload struct {
	UserId        string `json:"userId"`
	WalletAddress string `json:"walletAddress"`
}
