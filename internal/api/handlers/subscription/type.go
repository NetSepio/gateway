package subscription

import "netsepio-gateway-v1.1/models"

type BuyErebrusNFTResponse struct {
	ClientSecret string `json:"clientSecret"`
}

type SubscriptionResponse struct {
	Subscription *models.Subscription `json:"subscription,omitempty"`
	Status       string               `json:"status"`
}

type CreateSubscriptionPayload struct {
	PlanId string `json:"plan_id" binding:"required"`
}