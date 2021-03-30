package handler

import (
	"bwastartup/campaign"
	"github.com/gin-gonic/gin"
	"net/http"
)

type campaignHandler struct{
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) Index(c *gin.Context)  {
	campaigns, err := h.campaignService.GetCampaigns(0)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "campaign_index.html", gin.H{"campaigns":campaigns})
}