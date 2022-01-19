package delegateartifactcreation

type DelegateArtifactCreationRequest struct {
	CreatorAddress string `json:"creatorAddress" binding:"required"`
	MetaDataHash   string `json:"metaDataHash" binding:"required"`
}

type DelegateArtifactCreationPayload struct {
	TransactionHash string `json:"transactionHash"`
}
