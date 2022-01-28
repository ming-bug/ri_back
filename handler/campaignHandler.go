package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ri/enum"
	"ri/model"
	. "ri/repository"
	"strconv"
)

// MyCampaigns Campaign info
// @Summary Get user basic campaign info
// @Description Get campaign info by email, current page and page size
// @Tags campaign
// @Accept application/json
// @Produce application/json
// @Param userUin query string true "user identify"
// @Param current query int true "list current page"
// @Param pageSize query int true "list page size"
// @Success 200 {object} model.ResultCont
// @Failure 404 {object} model.ResultCont
// @Router /api/campaign/myCampaigns [get]
func MyCampaigns(c *gin.Context) {
	result := model.NewResult(c)
	userUin := c.Query("userUin")
	current, _ := strconv.Atoi(c.Query("current"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	email, err := QueryUserByUin(userUin)
	if err != nil {
		result.Faild(http.StatusNotFound, err.Error(), enum.Error)
		return
	}
	campaigns, err := QueryCampaignByEmail(email, current, pageSize)
	if err != nil {
		result.Faild(http.StatusNotFound, err.Error(), enum.Error)
		return
	}
	validCampaigns := make([]model.Campaign, 0)
	invalidCampaigns := make([]model.Campaign, 0)
	for _, campaign := range campaigns {
		switch campaign.Status {
		case 5, 6, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18:
			validCampaigns = append(validCampaigns, campaign)
		default:
			invalidCampaigns = append(invalidCampaigns, campaign)
		}
	}
	result.Success(validCampaigns)
}
