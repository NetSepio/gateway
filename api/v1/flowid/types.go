package flowid

type GetFlowIdRequest struct {
	WalletAddress string `json:"walletAddress" binding:"required"`
}

type GetFlowIdPayload struct {
	Eula   string `json:"eula"`
	FlowId string `json:"flowId"`
}
