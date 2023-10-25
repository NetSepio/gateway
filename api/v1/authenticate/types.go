package authenticate

type AuthenticateRequest struct {
	FlowId    string `json:"flowId" binding:"required"`
	Signature string `json:"signature" binding:"required"`
	PubKey    string `json:"pubKey" binding:"required"`
}

type AuthenticatePayload struct {
	Token string `json:"token"`
}
