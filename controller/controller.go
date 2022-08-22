package controller

import (
	"net/http"

	models "github.com/Modifa/DevFinde_Main/models"
	services "github.com/Modifa/DevFinde_Main/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetDeveloperProfile(c *gin.Context) {
	var rb models.Returnblock
	var u models.DeveloperRequest
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	d, profile := services.GetDeveloperProfile(u.Username)
	if !d {
		rb.New(false, "User Does Not Exist", profile)
	} else if !profile.Activated {
		rb.New(false, "Account Not Activated", profile)
	} else if profile.Password != u.Password {
		rb.New(false, "Password Incorrect", profile)
	} else {
		rb.New(true, "", profile)
	}
}
