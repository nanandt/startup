package transaction

import "bwastartup/user"

type GetCampaignTrnsactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}