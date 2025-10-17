package controller

import (
	"net/http"

	"github.com/FallenCap/url-shortner/services"
	"github.com/gin-gonic/gin"
)

type ShortenRequst struct {
	URL string `json:"url" binding:"required"`
}

func ShortenUrl(c *gin.Context) {
	var req ShortenRequst

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Url is required"})
		return
	}

	shortened, err := services.CreateShortUrl(req.URL)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to shorten Url"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Short":    shortened.Short,
		"Original": shortened.Original,
	})
}

func RedirectToOriginal(c *gin.Context) {
	short := c.Param("short")

	url, err := services.GetOriginalUrl(short)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Url not found"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url.Original)
}
