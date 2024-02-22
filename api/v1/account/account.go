package account

import (
	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/account")
	{
		g.POST("generate-auth-id", GenerateAuthId)
		g.Use(paseto.PASETO(true))
		g.POST("paseto-from-magic-link", PasetoFromMagicLink)
		g.POST("auth-google", authGoogle)
		g.Use(paseto.PASETO(false))
		g.DELETE("remove-mail", removeMail)
	}
}
