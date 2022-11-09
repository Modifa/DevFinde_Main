package controller

import (
	"fmt"
	"net/http"

	models "github.com/Modifa/DevFinde_Main/models"
	services "github.com/Modifa/DevFinde_Main/services"
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
	lk.Description = u.Description
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
