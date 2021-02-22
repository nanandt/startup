package transaction

import "bwastartup/user"

// GetCampaignTrnsactionsInput is...
// exported func GetCampaignTrnsactionsInput
type GetCampaignTrnsactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionInput struct {
	Amount int 	`json:"amount" binding:"required"`
	CampaignID int `json:"campaign_id" binding:"required"`
	User user.User
}