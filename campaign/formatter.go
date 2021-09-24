package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{
		ID:               campaign.ID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		ImageURL:         "",
		Slug:             campaign.Slug,
	}

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormater := []CampaignFormatter{}

	for _, c := range campaigns {
		campaignsFormater = append(campaignsFormater, FormatCampaign(c))
	}

	return campaignsFormater
}

type CampaignDetailFormatter struct {
	ID               int                      `json:"id"`
	UserID           int                      `json:"user_id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"short_description"`
	Description      string                   `json:"description"`
	ImageURL         string                   `json:"image_url"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"current_amount"`
	Slug             string                   `json:"slug"`
	Perks            []string                 `json:"perks"`
	User             CampaignUserFormatter    `json:"user"`
	Images           []CampaignImageFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	images := []CampaignImageFormatter{}

	for _, image := range campaign.CampaignImages {
		campaignImageFormatter := CampaignImageFormatter{
			ImageURL:  image.FileName,
			IsPrimary: false,
		}

		if image.IsPrimary == 1 {
			campaignImageFormatter.IsPrimary = true
		}

		images = append(images, campaignImageFormatter)
	}

	campaignUserFormatter := CampaignUserFormatter{
		Name:     campaign.User.Name,
		ImageURL: campaign.User.AvatarFileName,
	}

	campaignFormatter := CampaignDetailFormatter{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
		ImageURL:         "",
		User:             campaignUserFormatter,
		Images:           images,
	}

	if len(campaign.CampaignImages) > 0 {
		for index, image := range campaign.CampaignImages {
			if image.IsPrimary == 1 {
				campaignFormatter.ImageURL = campaign.CampaignImages[index].FileName
			}
		}

	}

	perks := strings.Split(campaign.Perks, ", ")
	campaignFormatter.Perks = perks

	return campaignFormatter
}
