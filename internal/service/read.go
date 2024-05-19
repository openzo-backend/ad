package service

import (
	"github.com/gin-gonic/gin"
	"github.com/tanush-128/openzo_backend/ad/internal/models"
)

func (s *adService) GetAdByID(ctx *gin.Context, id string) (models.Ad, error) {
	ad, err := s.adRepository.GetAdByID(id)
	if err != nil {
		return models.Ad{}, err
	}

	return ad, nil
}

func (s *adService) GetAdsByPincode(ctx *gin.Context, pincode string) ([]models.Ad, error) {
	ads, err := s.adRepository.GetAdsByPincode(pincode)
	if err != nil {
		return []models.Ad{}, err
	}

	return ads, nil
}

