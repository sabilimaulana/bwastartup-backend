package campaign

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
	FindByID(ID int) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Debug().Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByUserID(userID int) ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Debug().Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByID(ID int) (Campaign, error) {
	var campaign Campaign
	err := r.db.Debug().Where("id = ?", ID).Preload("CampaignImages").Preload("User").Find(&campaign).Error
	if err != nil {
		return campaign, err
	}

	if campaign.ID == 0 {
		return campaign, errors.New("no campaign found on that id")
	}

	return campaign, nil
}

func (r *repository) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Debug().Save(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) Update(campaign Campaign) (Campaign, error) {
	err := r.db.Debug().Save(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
