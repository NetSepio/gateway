package discordauth

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/util/pkg/httphelper"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/discord-auth")
	{
		g.GET("", discordAuth)
	}
}

func discordAuth(c *gin.Context) {
	// TODO: validate state
	endpoint := oauth2.Endpoint{
		AuthURL:   "https://discord.com/api/oauth2/authorize",
		TokenURL:  "https://discord.com/api/oauth2/token",
		AuthStyle: oauth2.AuthStyleInParams,
	}
	conf := &oauth2.Config{
		RedirectURL:  envconfig.EnvVars.DISCORD_REDIRECT_URL,
		ClientID:     envconfig.EnvVars.DISCORD_CLIENT_ID,
		ClientSecret: envconfig.EnvVars.DISCORD_CLIENT_SECRET,
		Scopes:       []string{"identify"},
		Endpoint:     endpoint,
	}
	token, err := conf.Exchange(context.Background(), c.Request.FormValue("code"))
	if err != nil {
		logwrapper.Errorf("failed to exchange token: %s", err)
		httphelper.ErrResponse(c, 500, "")
		return
	}
	res, err := conf.Client(context.Background(), token).Get("https://discord.com/api/users/@me")
	if err != nil {
		logwrapper.Errorf("failed to query client: %s", err)
		httphelper.ErrResponse(c, 500, "")
		return
	}
	r, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logwrapper.Errorf("failed to get data from body of exchange token response: %s", err)
		httphelper.ErrResponse(c, 500, "")
		return
	}
	var discordRes DicordExchangeRes
	json.Unmarshal(r, &discordRes)
	httphelper.SuccessResponse(c, "Discord OAuth2 handled", nil)
}
