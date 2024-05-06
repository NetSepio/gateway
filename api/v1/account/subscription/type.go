package subscription

import "github.com/NetSepio/gateway/models"

type BuyErebrusNFTResponse struct {
	ClientSecret string `json:"clientSecret"`
}

type SubscriptionResponse struct {
	Subscription *models.Subscription `json:"subscription,omitempty"`
	Status       string               `json:"status"`
}
