package transaction

type GetCampaignTrnsactionsInput struct {
	ID int `uri:"id" binding:"required"`
}