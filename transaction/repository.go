package transaction

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}
// Repository ...
type Repository interface {
	GetByCampaignID (campaignID int) ([]Transaction, error)
	GetByUserID (userID int)([]Transaction, error)
}

// NewRepository ...
// exported func NewRepository
func NewRepository(db *gorm.DB) *repository{
	return &repository{db}
}

func (r *repository) GetByCampaignID (campaignID int) ([]Transaction, error){
	var transactions []Transaction

	err := r.db.Preload("User").Where("campaign_id = ?", campaignID).Order("id DESC").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) GetByUserID (userID int)([]Transaction, error){
	var transactions []Transaction

	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Where("user_id = ?", userID).Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}