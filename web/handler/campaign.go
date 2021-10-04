package handler

import (
	"bwastartup/campaign"
	"bwastartup/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
	userService     user.Service
}

func NewCampaignHandler(campaignService campaign.Service, userService user.Service) *campaignHandler {
	return &campaignHandler{campaignService, userService}
}

func (h *campaignHandler) Index(c *gin.Context) {
	campaigns, err := h.campaignService.GetCampaigns(0)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "campaign_index.html", gin.H{"campaigns": campaigns})
}

func (h *campaignHandler) New(c *gin.Context) {
	users, err := h.userService.GetAllUser()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	input := campaign.FormCreateCampaignInput{
		Users: users,
	}

	c.HTML(http.StatusOK, "campaign_new.html", input)
}

func (h *campaignHandler) Create(c *gin.Context) {
	var input campaign.FormCreateCampaignInput
	err := c.ShouldBind(&input)
	if err != nil {
		users, e := h.userService.GetAllUser()
		if e != nil {
			c.HTML(http.StatusInternalServerError, "error.html", nil)
			return
		}

		input.Users = users
		input.Error = err
		c.HTML(http.StatusOK, "campaign_new.html", input)
		return
	}

	userCampaign, err := h.userService.GetUserByID(input.UserID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	futureCampaign := campaign.CreateCampaignInput{
		User:             userCampaign,
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		GoalAmount:       input.GoalAmount,
		Perks:            input.Perks,
	}

	_, err = h.campaignService.CreateCampaign(futureCampaign)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/campaigns")
}

func (h *campaignHandler) NewImage(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	c.HTML(http.StatusOK, "campaign_image.html", gin.H{"ID": id})
}

func (h *campaignHandler) CreateImage(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	file, err := c.FormFile("file")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	existingCampaign, err := h.campaignService.GetCampaignByID(campaign.GetCampaignDetailInput{ID: id})
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	path := fmt.Sprintf("images/%d-%s", existingCampaign.UserID, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	createCampaignImageInput := campaign.CreateCampaignImageInput{
		CampaignID: existingCampaign.ID,
		IsPrimary:  true,
		User:       existingCampaign.User,
	}

	_, err = h.campaignService.SaveCampaignImage(createCampaignImageInput, path)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/campaigns")
}

func (h *campaignHandler) Edit(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	existingCampaign, err := h.campaignService.GetCampaignByID(campaign.GetCampaignDetailInput{ID: id})
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	input := campaign.FormUpdateCampaignInput{
		ID:               existingCampaign.ID,
		Name:             existingCampaign.Name,
		ShortDescription: existingCampaign.ShortDescription,
		Description:      existingCampaign.Description,
		GoalAmount:       existingCampaign.GoalAmount,
		Perks:            existingCampaign.Perks,
		User:             existingCampaign.User,
	}

	c.HTML(http.StatusOK, "campaign_edit.html", input)
}
