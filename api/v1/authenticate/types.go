package authenticate

type AuthenticateRequest struct {
	FlowId       string  `json:"flowId" binding:"required"`
	Signature    string  `json:"signature" binding:"omitempty,hexadecimal,startswith=0x"`
	PubKey       string  `json:"pubKey" binding:"omitempty,hexadecimal,startswith=0x"`
	SignatureSui string  `json:"signatureSui" binding:"omitempty"`
	PubkeySui    []uint8 `json:"pubkeySui" binding:"omitempty"`
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
