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

func AddEducation(c *gin.Context) {

	db := services.DB{}

	var t models.EducationRequest
	var e models.Education

	if err := c.ShouldBindBodyWith(&t, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	e.ID = t.ID
	e.Intsitution = t.Intsitution
	e.Qualification_name = t.Qualification_name
	e.Qualification_type_ = utils.StringToInt64(t.Qualification_type_)
	e.Start_date = t.Start_date
	e.EndDate = t.EndDate

	_, err := db.SAVEONDB("dev_finder.fn_update_developer_resume", e)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	var q models.DBIDRequest
	q.ID = t.ID

	resp, err := db.GetEducation("dev_finder.fn_get_developer_education", q)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	//
	go func() {
		for i := 0; i < len(resp); i++ {
			services.SaveDeveloperEnducation(resp[i], t.UserName)
		}
	}()
}

//
func GetDeveloperEducation(c *gin.Context) {

	var t models.ResumeRequestRedis
	var rb models.Returnblock
	if err := c.ShouldBindBodyWith(&t, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	exists, education := services.GetDeveloperEducation(t.Username)
	if exists {
		c.JSON(http.StatusOK, rb.New(true, "", education))
		return
	} else {
		c.JSON(http.StatusOK, rb.New(false, "No Fields Found", education))
		return
	}
}
