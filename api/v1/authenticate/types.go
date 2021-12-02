package authenticate

import jwt "github.com/golang-jwt/jwt/v4"

type customClaims struct {
	WalletAddress string `json:"walletAddress"`
	jwt.RegisteredClaims
}
type AuthenticateRequest struct {
	FlowId    string `json:"flowId" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}
