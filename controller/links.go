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

//AddDeveloperLink
func AddDeveloperLink(c *gin.Context) {

	db := services.DB{}
	var lk models.LinksRequest
	// var rb models.Returnblock
	var u models.LinksRequestDB
	var q models.IDRequest
	// var  models.DeveloperProfile

	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	// _, profile := services.GetDeveloperProfile(u.Username)
	q.Id = u.Id
	lk.LinkType = utils.StringToInt64(u.LinkType)
	lk.Link = u.Link
	lk.Id = u.Id

	_, err := db.SAVEONDB("dev_finder.fn_add_links", lk)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	//Update Developer Profile With New Link
	links, err := db.GetDeveloperLinks("dev_finder.fn_get_developer_links", q)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	//Use Go Routine

	go func() {
		for i := 0; i < len(links); i++ {
			services.SaveDeveloperLinks(links[i], u.Username)
		}
	}()
}

//
func UpdateDeveloperLink(c *gin.Context) {

	db := services.DB{}
	var rb models.Returnblock
	var t models.UpdateResumeLinkDB
	var resume models.UpdateResumeLinkPOST
	// var  models.DeveloperProfile

	if err := c.ShouldBindBodyWith(&t, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, rb.New(false, err.Error(), 0))
		return
	}
	// _, profile := services.GetDeveloperProfile(u.Username)

	resume.ID = t.ID
	resume.Link = t.Link
	resume.LinkId = t.LinkId
	errDel := db.SAVEONDBNPRETURN("dev_finder.fn_update_link", resume)
	if errDel != nil {
		fmt.Println("QueryRow failed: ", errDel.Error())
		errormessage := fmt.Sprintf("%v\n", errDel)
		c.JSON(http.StatusBadRequest, rb.New(false, errormessage, 0))
		return
	}
	//Update Developer Profile With New Link
	var q models.DBIDRequest
	q.ID = t.ID
	links, err := db.GetDeveloperLinks("dev_finder.fn_get_developer_links", q)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, rb.New(false, errormessage, 0))
		return
	}
	//Use Go Routine

	go func() {
		services.ClearDeveloperLinks(t.Username)
		for i := 0; i < len(links); i++ {
			services.SaveDeveloperLinks(links[i], t.Username)
		}
	}()

	//Return
	if errDel != nil {
		c.JSON(http.StatusOK, rb.New(false, "Could Not Update, Please try Again", 0))

		return
	} else {
		c.JSON(http.StatusOK, rb.New(true, "", 0))

		return
	}
}

//dev_finder.fn_delete_link

func DeleteDeveloperLink(c *gin.Context) {

	db := services.DB{}
	var rb models.Returnblock
	var t models.DeleteResumeLink
	var resume models.DeleteResumeLinkPOST

	if err := c.ShouldBindBodyWith(&t, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, rb.New(false, err.Error(), 0))
		return
	}

	resume.ID = t.ID
	resume.LinkId = t.LinkId
	_, errDel := db.SAVEONDB("dev_finder.fn_delete_link", resume)
	if errDel != nil {
		fmt.Println("QueryRow failed: ", errDel.Error())
		errormessage := fmt.Sprintf("%v\n", errDel)
		c.JSON(http.StatusBadRequest, rb.New(false, errormessage, 0))
		return
	}
	//Update Developer Profile With New Link
	var q models.DBIDRequest
	q.ID = t.ID
	links, err := db.GetDeveloperLinks("dev_finder.fn_get_developer_links", q)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, rb.New(false, errormessage, 0))
		c.JSON(http.StatusUnprocessableEntity, rb.New(false, err.Error(), 0))
		return
	}
	//Use Go Routine

	go func() {
		services.ClearDeveloperLinks(t.Username)
		for i := 0; i < len(links); i++ {
			services.SaveDeveloperLinks(links[i], t.Username)
		}
	}()

	if errDel != nil {
		c.JSON(http.StatusOK, rb.New(false, "Could Not Delete, Please try Again", 0))

		return
	} else {
		c.JSON(http.StatusOK, rb.New(true, "", 0))

		return
	}
}

func GetDeveloperLinksRedis(c *gin.Context) {

	var t models.ResumeRequestRedis
	var rb models.Returnblock
	if err := c.ShouldBindBodyWith(&t, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	exists, Links := services.GetDeveloperLinksRD(t.Username)
	if exists {
		c.JSON(http.StatusOK, rb.New(true, "", Links))
		return
	} else {
		c.JSON(http.StatusOK, rb.New(false, "No Fields Found", Links))
		return
	}
}
