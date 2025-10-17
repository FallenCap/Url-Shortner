package services

import (
	"github.com/FallenCap/url-shortner/config"
	"github.com/FallenCap/url-shortner/models"
	"github.com/FallenCap/url-shortner/utils"
)

func CreateShortUrl(original string) (models.Url, error) {
	shortCode := utils.GenerateShortCode(6)

	url := models.Url{
		Original: original,
		Short:    shortCode,
	}

	result := config.DB.Create(&url)

	return url, result.Error
}

func GetOriginalUrl(shortCode string) (models.Url, error) {
	var url models.Url
	result := config.DB.Where("short = ?", shortCode).First(&url)

	return url, result.Error
}
