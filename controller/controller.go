package controller

import (
	"fmt"
	"net/http"

	models "github.com/Modifa/DevFinde_Main/models"
	services "github.com/Modifa/DevFinde_Main/services"
	utils "github.com/Modifa/DevFinde_Main/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func RegisterDeveloper(c *gin.Context) {

	db := services.DB{}

	// var rb models.Returnblock
	var u models.DeveloperRegister
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	number, _ := utils.FormatMobileNumber(u.Mobile_number)
	u.Mobile_number = number
	_, err := db.RegisterDeveloper("dev_finder.fn_register_developer", u)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	var t models.GetDeveloperProfile
	t.EmailAddress = u.Email_address

	resp, err := db.GetDeveloperProfile("dev_finder.fn_get_developer_profile", t)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	go func() {
		err1 := services.SaveDeveloperprofile(resp[0])
		if err != nil {
			fmt.Println(err1)
		}
	}()

}

func GetDeveloperProfile(c *gin.Context) {
	var rb models.Returnblock
	var u models.DeveloperRequest
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	d, profile := services.GetDeveloperProfile(u.Username)
	active := utils.StringToBool(profile.Active)
	if !d {
		c.JSON(http.StatusOK, rb.New(false, "Acoount Does Not Exist", profile))
		return
	} else if !active {
		c.JSON(http.StatusOK, rb.New(false, "Account Not Activated", profile))
		return
	} else if profile.Password != u.Password {
		c.JSON(http.StatusOK, rb.New(false, "Password Incorrect", profile))
		return
	} else {
		c.JSON(http.StatusOK, rb.New(true, "", profile))
		return
	}
}
