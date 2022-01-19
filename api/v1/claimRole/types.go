package claimrole

type ClaimRoleRequest struct {
	Signature string `json:"signature" binding:"required"`
	FlowId    string `json:"flowId" binding:"required"`
}

type ClaimRolePayload struct {
	TransactionHash string `json:"transactionHash"`
}
