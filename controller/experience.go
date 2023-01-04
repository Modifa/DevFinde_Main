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
	var rb models.Returnblock
	var u models.ExperienceRequest
	var q models.IDRequest
	var exp models.ExperienceRequestDB

	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	//
	exp.Id = u.Id
	exp.Company = u.Company
	exp.Description = u.Description
	exp.Title = u.Title
	exp.Start_date = u.Start_date
	exp.End_date = u.End_date

	err := db.SAVEONDBNPRETURN("dev_finder.fn_add_experience", exp)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, rb.New(false, errormessage, 0))
		return
	}
	//
	q.Id = u.Id
	Experiences, err := db.GetDeveloperExperience("dev_finder.fn_add_experience", q)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, rb.New(false, errormessage, 0))
		return
	}
	//
	go func() {
		services.ClearDeveloperExperience(u.Username)
		for i := 0; i < len(Experiences); i++ {
			services.SaveDeveloperExperience(Experiences[i], u.Username)
		}
	}()

	c.JSON(http.StatusOK, rb.New(true, "", 0))

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
