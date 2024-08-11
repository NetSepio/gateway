package leaderboard

// ApplyRoutes applies router to gin Router
// func ApplyRoutes(r *gin.RouterGroup) {
// 	g := r.Group("/reviewerdetails")
// 	{
// 		g.GET("", getProfile)
// 	}
// }

// func update(c *gin.Context) {
// 	db := dbconfig.GetDb()
// 	var request GetReviewerDetailsQuery
// 	err := c.BindQuery(&request)

// 	payload := GetReviewerDetailsPayload{
// 		Name:              user.Name,
// 		WalletAddress:     user.WalletAddress,
// 		ProfilePictureUrl: user.ProfilePictureUrl,
// 		Discord:           user.Discord,
// 		Twitter:           user.Twitter,
// 	}
// 	httpo.NewSuccessResponseP(200, "Profile fetched successfully", payload).SendD(c)
// }
