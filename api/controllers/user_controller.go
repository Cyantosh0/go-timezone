package controllers

import (
	"net/http"
	"time"

	"github.com/Cyantosh0/gorm-crud/api/repositories"
	"github.com/Cyantosh0/gorm-crud/constants"
	"github.com/Cyantosh0/gorm-crud/lib"
	"github.com/Cyantosh0/gorm-crud/models"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	repository repositories.UserRepository
}

func NewUserController(repository repositories.UserRepository) UserController {
	return UserController{
		repository: repository,
	}
}

func (uc UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := uc.repository.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": user})
}

func (u UserController) UpdateUser(c *gin.Context) {
	paramID := c.Param("id")

	var user models.User
	err := u.repository.First(&user, "id = ?", lib.ParseUUID(paramID)).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := u.repository.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": user})
}

func (u UserController) RetrieveUser(c *gin.Context) {
	paramID := c.Param("id")

	var user models.User
	err := u.repository.First(&user, "id = ?", lib.ParseUUID(paramID)).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": user})
}

func (u UserController) GellAllUsers(c *gin.Context) {
	today := c.Query("today")
	date := c.Query("date")
	fromDate := c.Query("fromDate")
	toDate := c.Query("toDate")

	timeZone := c.Query("timeZone")

	var users []models.User

	chain := u.repository.Model(&models.User{})

	if today == "true" {
		chain = chain.Where("convert_tz(created_at,'UTC',?) >= ? AND convert_tz(created_at,'UTC',?) <= ?", timeZone, time.Now().Format(constants.ISOFormat), timeZone, time.Now().AddDate(0, 0, 1).Format(constants.ISOFormat))
	} else {
		if date != "" {
			d, err := time.Parse("2006-01-02T15:04:05.000Z", date)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
				return
			}

			chain = chain.Where("convert_tz(created_at,'UTC',?) >= ? AND convert_tz(created_at,'UTC',?) <= ?", timeZone, date, timeZone, d.AddDate(0, 0, 1))
		} else {
			if fromDate != "" {
				chain = chain.Where("created_at >= ?", fromDate)
			}
			if toDate != "" {
				chain = chain.Where("created_at <= ?", toDate)
			}
		}
	}

	err := chain.Find(&users).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": users})
}

func (u UserController) DeleteUser(c *gin.Context) {
	paramID := c.Param("id")

	err := u.repository.Where("id = ?", lib.ParseUUID(paramID)).Delete(&models.User{}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "user deleted"})
}

func (u UserController) UpdateBasicInformation(c *gin.Context) {
	paramID := c.Param("id")

	var user models.User
	err := u.repository.First(&user, "id = ?", lib.ParseUUID(paramID)).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = u.repository.UpdateBasicInformation(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": user})
}
