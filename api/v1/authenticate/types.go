package authenticate

type AuthenticateRequest struct {
	WalletAddress string `json:"walletAddress"`
	FlowId        string `json:"flowId"`
	Signature     string `json:"signature"`
	PubKey        string `json:"pubKey"`
	SignatureSui  string `json:"signatureSui"`
	Message       string `json:"message"`
	AccessToken   string `json:"accessToken"`
	IdToken       string `json:"idToken"` // to be pass in bearer token [ AUTHORIZATION KEY ]
	ChainName     string `json:"chainName" binding:"required"`
	FlowId       string `json:"flowId" binding:"required"`
	Signature    string `json:"signature" binding:"omitempty,hexadecimal"`
	PubKey       string `json:"pubKey" binding:"omitempty"`
	SignatureSui string `json:"signatureSui" binding:"omitempty"`
	Message      string `json:"message" binding:"omitempty"`
	AccessToken  string `json:"accessToken" binding:"omitempty"`
	IdToken      string `json:"idToken" binding:"omitempty"` // to be pass in bearer token [ AUTHORIZATION KEY ]
	ChainName    string `json:"chainName" binding:"omitempty"`
}
type AuthenticateRequestNoSign struct {
	FlowId        string `json:"flowId" binding:"required"`
	WalletAddress string `json:"walletAddress" binding:"required"`
}

type AuthenticateRequestByWallet struct {
	ChainName     string `json:"chain" binding:"required"`
	WalletAddress string `json:"wallet_address" binding:"required"`
}

type AuthenticatePayload struct {
	Token  string `json:"token"`
	UserId string `json:"userId"`
}

type AuthenticateTokenPayload struct {
	UserId        string `json:"userId"`
	WalletAddress string `json:"walletAddress"`
}
