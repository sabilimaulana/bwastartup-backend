package transaction

import "gorm.io/gorm"

type Repository interface {
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository {
	return repository{db}
}

func (r *repository) GetByCampaignID(CampaignID int) ([]Transaction, error) {
	var transactions []Transaction

	err := r.db.Debug().Preload("User").Where("campaign_id = ?", CampaignID).Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
