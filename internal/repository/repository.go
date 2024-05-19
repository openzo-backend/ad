package repository

import (
	"github.com/google/uuid"
	"github.com/tanush-128/openzo_backend/ad/internal/models"

	"gorm.io/gorm"
)

type AdRepository interface {
	CreateAd(Ad models.Ad) (models.Ad, error)
	GetAdByID(id string) (models.Ad, error)
	
	GetAdsByPincode(pincode string) ([]models.Ad, error)

	// Add more methods for other Ad operations (GetAdByEmail, UpdateAd, etc.)

}

type adRepository struct {
	db *gorm.DB
}

func NewAdRepository(db *gorm.DB) AdRepository {

	return &adRepository{db: db}
}

func (r *adRepository) CreateAd(Ad models.Ad) (models.Ad, error) {
	Ad.ID = uuid.New().String()

	tx := r.db.Create(&Ad)

	if tx.Error != nil {
		return models.Ad{}, tx.Error
	}

	return Ad, nil
}

func (r *adRepository) GetAdByID(id string) (models.Ad, error) {
	var Ad models.Ad
	tx := r.db.Where("id = ?", id).First(&Ad)
	if tx.Error != nil {
		return models.Ad{}, tx.Error
	}

	return Ad, nil
}



func (r *adRepository) GetAdsByPincode(pincode string) ([]models.Ad, error) {
	var Ads []models.Ad
	tx := r.db.Where("pincode = ?", pincode).Find(&Ads)
	if tx.Error != nil {
		return []models.Ad{}, tx.Error
	}

	return Ads, nil
}

