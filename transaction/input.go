package transaction

import "bwastartup/user"

// GetCampaignTrnsactionsInput is...
// exported func GetCampaignTrnsactionsInput
type GetCampaignTrnsactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}