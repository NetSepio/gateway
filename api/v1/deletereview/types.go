package deletereview

type DeleteReviewRequest struct {
	MetaDataUri string `json:"metaDataUri" binding:"required"`
}

type DeleteReviewPayload struct {
	TransactionHash    string `json:"transactionHash"`
	TransactionVersion int64  `json:"transactionVersion"`
}
