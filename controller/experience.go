package controller

import (
	"fmt"
	"net/http"

	models "github.com/Modifa/DevFinde_Main/models"
	services "github.com/Modifa/DevFinde_Main/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func AddDeveloperExperience(c *gin.Context) {

	db := services.DB{}

	var u models.ExperienceRequest
	var q models.IDRequest

	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	q.Id = u.Id

	_, err := db.SAVEONDB("dev_finder.fn_add_experience", u)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	//
	Experiences, err := db.GetDeveloperExperience("dev_finder.fn_add_experience", q)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	//
	go func() {
		for i := 0; i < len(Experiences); i++ {
			services.SaveDeveloperExperience(Experiences[i], u.Username)
		}
	}()

}

//GetDeveloperExperienceRedis
func GetDeveloperExperienceRedis(c *gin.Context) {

	var t models.ResumeRequestRedis
	var rb models.Returnblock
	if err := c.ShouldBindBodyWith(&t, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	exists, exp := services.GetDeveloperExperienceRD(t.Username)
	if exists {
		c.JSON(http.StatusOK, rb.New(true, "", exp))
		return
	} else {
		c.JSON(http.StatusOK, rb.New(false, "No Fields Found", exp))
		return
	}
}
