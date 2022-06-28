package feedback

type PostFeedbackRequest struct {
	Feedback string `json:"feedback" binding:"required"`
}
