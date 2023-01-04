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

//Add Resume Link
func AddResumeLink(c *gin.Context) {

	db := services.DB{}
	var rb models.Returnblock
	var t models.AddResumeLinkDB
	var resume models.AddResumeLinkDBPOST

	if err := c.ShouldBindBodyWith(&t, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	resume.ID = t.ID
	resume.Link = t.Link
	resume.LinkType = utils.StringToInt64(t.LinkType)
	_, errAdd := db.SAVEONDB("dev_finder.fn_add_links", resume)
	if errAdd != nil {
		fmt.Println("QueryRow failed: ", errAdd.Error())
		errormessage := fmt.Sprintf("%v\n", errAdd)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, errAdd.Error())
		return
	}

	var q models.IDRequest
	q.Id = t.ID

	resp, err := db.GetDeveloperLinks("dev_finder.fn_get_developer_links", q)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	go func() {
		for i := 0; i < len(resp); i++ {
			err1 := services.SaveDeveloperLinks(resp[i], t.Username)
			if err != nil {
				fmt.Println(err1)
			}
		}
	}()
	if errAdd != nil {
		c.JSON(http.StatusOK, rb.New(false, "Resume Not Uploaded", 0))

		return
	} else {
		c.JSON(http.StatusOK, rb.New(true, "", 1))

		return

	}
}

//
func AddResume(c *gin.Context) {

	db := services.DB{}

	var t models.AddResumeDB
	var resume models.AddResume

	if err := c.ShouldBindBodyWith(&t, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	resume.ID = t.ID
	resume.ResumeURL = t.ResumeURL
	_, err := db.SAVEONDB("dev_finder.fn_add_resume", resume)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	var q models.DBIDRequest
	q.ID = t.ID

	resp, err := db.GetResume("dev_finder.fn_get_developer_resume", q)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	go func() {
		err1 := services.SaveDeveloperResume(resp[0])
		if err1 != nil {
			fmt.Println(err1)
		}
	}()
}

func Update(c *gin.Context) {

	// db := services.DB{}

	var t models.UpdateResumeLinkDB
	var resume models.UpdateResumeLinkPOST

	if err := c.ShouldBindBodyWith(&t, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	resume.ID = t.ID
	// resume.ResumeURL = t.ResumeURL
	// _, err := db.SAVEONDB("dev_finder.fn_update_developer_resume", resume)
	// if err != nil {
	// 	fmt.Println("QueryRow failed: ", err.Error())
	// 	errormessage := fmt.Sprintf("%v\n", err)
	// 	c.JSON(http.StatusBadRequest, errormessage)
	// 	c.JSON(http.StatusUnprocessableEntity, err.Error())
	// 	return
	// }

	// var q models.DBIDRequest
	// q.ID = t.ID

	// resp, err := db.GetResume("dev_finder.fn_get_developer_resume", q)
	// if err != nil {
	// 	fmt.Println("QueryRow failed: ", err.Error())
	// 	errormessage := fmt.Sprintf("%v\n", err)
	// 	c.JSON(http.StatusBadRequest, errormessage)
	// 	c.JSON(http.StatusUnprocessableEntity, err.Error())
	// 	return
	// }

	// go func() {
	// 	err1 := services.SaveDeveloperResume(resp[0])
	// 	if err1 != nil {
	// 		fmt.Println(err1)
	// 	}
	// }()
}

//GetDeveloperResume

func GetDeveloperResume(c *gin.Context) {

	var t models.ResumeRequestRedis
	var rb models.Returnblock
	if err := c.ShouldBindBodyWith(&t, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	exists, resume := services.GetDeveloperResume(t.Username)
	if exists {
		c.JSON(http.StatusOK, rb.New(true, "", resume))
		return
	} else {
		c.JSON(http.StatusOK, rb.New(false, "Resume Not Uploaded", resume))
		return
	}
}

//dev_finder.fn_add_resume_desc
//ResumedescReq

func AddResumeDesc(c *gin.Context) {

	db := services.DB{}

	var t models.DBIDRequest
	var resume models.ResumedescReq

	if err := c.ShouldBindBodyWith(&resume, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	See, err := db.SAVEONDB("dev_finder.fn_add_resume_desc", resume)
	fmt.Println("error", See)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	t.ID = resume.ID
	resp, err := db.GetDeveloperResumeDesc("dev_finder.fn_get_developer_resume_desc", t)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	//SaveDeveloperResumeDesc
	var W models.ResumedescRedis
	W.Dateadded = resp[0].Dateadded
	W.Description = resp[0].Description

	go func() {
		err1 := services.SaveDeveloperResumeDesc(W, resp[0].Username)
		if err != nil {
			fmt.Println(err1)
		}
	}()
}

//ResumedescRedisUP
func UpdateResumeDesc(c *gin.Context) {

	db := services.DB{}
	var rb models.Returnblock
	var t models.ResumedescRedisUP
	var resume models.ResumeDesc
	var q models.DBIDRequest

	if err := c.ShouldBindBodyWith(&t, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	resume.Description = t.Description
	resume.Developer_ID = t.Developer_ID

	_, err := db.SAVEONDB("dev_finder.fn_update_resume_desc", resume)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, rb.New(false, errormessage, 0))

		return
	}
	q.ID = resume.Developer_ID
	resp, err := db.GetDeveloperResumeDesc("dev_finder.fn_get_developer_resume_desc", q)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, rb.New(false, errormessage, 0))

		return
	}
	//SaveDeveloperResumeDesc
	var W models.ResumedescRedis
	W.Dateadded = resp[0].Dateadded
	W.Description = resp[0].Description
	W.ResID = resp[0].ResID

	go func() {
		err1 := services.SaveDeveloperResumeDesc(W, resp[0].Username)
		if err != nil {
			fmt.Println(err1)
		}
	}()

	c.JSON(http.StatusOK, rb.New(true, "", W))

}

//GetDeveloperResumeDesc
func GetDeveloperResumedesc(c *gin.Context) {

	var t models.ResumeRequestRedis
	var rb models.Returnblock
	if err := c.ShouldBindBodyWith(&t, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	exists, resume := services.GetDeveloperResumeDesc(t.Username)
	if exists {
		c.JSON(http.StatusOK, rb.New(true, "", resume))
		return
	} else {
		c.JSON(http.StatusOK, rb.New(false, "Resume Not Uploaded", resume))
		return
	}
}
