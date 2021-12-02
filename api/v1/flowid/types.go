package flowid

type GetFlowIdRequest struct {
	WalletAddress string `json:"walletAddress" binding:"required"`
}

type GetFlowIdResponse struct {
	Message string `json:"message"`
	FlowId  string `json:"flowId"`
}
