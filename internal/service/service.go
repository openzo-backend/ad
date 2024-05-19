package service

import (
	"github.com/gin-gonic/gin"
	"github.com/tanush-128/openzo_backend/ad/internal/models"
	"github.com/tanush-128/openzo_backend/ad/internal/pb"
	"github.com/tanush-128/openzo_backend/ad/internal/repository"
	"github.com/tanush-128/openzo_backend/ad/internal/utils"
)

type AdService interface {

	//CRUD
	CreateAd(ctx *gin.Context, req models.Ad) (models.Ad, error)
	GetAdByID(ctx *gin.Context, id string) (models.Ad, error)
	GetAdsByPincode(ctx *gin.Context, pincode string) ([]models.Ad, error)

}

type adService struct {
	adRepository repository.AdRepository
	imageClient     pb.ImageServiceClient
}

func NewAdService(adRepository repository.AdRepository,
	imageClient pb.ImageServiceClient,
) AdService {
	return &adService{adRepository: adRepository, imageClient: imageClient}
}

func (s *adService) CreateAd(ctx *gin.Context, req models.Ad) (models.Ad, error) {

	err := ctx.Request.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		return models.Ad{}, err
	}

	file, err := ctx.FormFile("image")
	if err == nil {

		imageBytes, err := utils.FileHeaderToBytes(file)
		if err != nil {
			return models.Ad{}, err
		}

		Image, err := s.imageClient.UploadImage(ctx, &pb.ImageMessage{
			ImageData: imageBytes,
		})
		if err != nil {
			return models.Ad{}, err
		}

		req.Image = Image.Url
	}
	createdAd, err := s.adRepository.CreateAd(req)
	if err != nil {
		return models.Ad{}, err // Propagate error
	}

	return createdAd, nil
}

