package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanush-128/openzo_backend/ad/internal/models"
	"github.com/tanush-128/openzo_backend/ad/internal/service"
	"github.com/tanush-128/openzo_backend/ad/internal/utils"
)

type Handler struct {
	adService service.AdService
}

func NewHandler(adService *service.AdService) *Handler {
	return &Handler{adService: *adService}
}

func (h *Handler) CreateAd(ctx *gin.Context) {
	var ad models.Ad

	ad.Name = ctx.PostForm("name")
	ad.URL = ctx.PostForm("url")
	ad.StoreID = ctx.PostForm("store_id")
	ad.Pincode = ctx.PostForm("pincode")
	ad.ExpiresAt = utils.StringToTime(ctx.PostForm("expires_at"))

	createdAd, err := h.adService.CreateAd(ctx, ad)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdAd)

}

func (h *Handler) GetAdByID(ctx *gin.Context) {
	id := ctx.Param("id")

	ad, err := h.adService.GetAdByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, ad)
}

func (h *Handler) GetAdsByPincode(ctx *gin.Context) {
	pincode := ctx.Param("pincode")

	ads, err := h.adService.GetAdsByPincode(ctx, pincode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, ads)
}
