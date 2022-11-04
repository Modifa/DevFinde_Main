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
	// developerID := profile.Id
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

//Add Developer Links\
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

//Update User Developer profile
//UpdateProfile

func UpdateDeveloperProfile(c *gin.Context) {

	db := services.DB{}

	// var rb models.Returnblock
	var u models.UpdateProfile
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	number, _ := utils.FormatMobileNumber(u.Mobile)
	u.Mobile = number
	_, err := db.SAVEONDB("dev_finder.fn_update_developer_details", u)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	var t models.GetDeveloperProfile
	t.EmailAddress = u.Email

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

func GetDBProfile(c *gin.Context) {

	db := services.DB{}

	var t models.GetDeveloperProfile

	if err := c.ShouldBindBodyWith(&t, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	resp, err := db.GetDeveloperProfile("dev_finder.fn_get_developer_profile", t)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	fmt.Print(resp)

}

//Update Developer Profile Image
func UpdateImage(c *gin.Context) {

	db := services.DB{}

	var t models.UpdateImageDB
	var image models.UpdateImage

	if err := c.ShouldBindBodyWith(&t, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	image.ID = t.ID
	image.ImageURL = t.ImageURL
	_, err := db.SAVEONDB("dev_finder.fn_update_developer_image", image)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	var q models.GetDeveloperProfile
	q.EmailAddress = t.Email

	resp, err := db.GetDeveloperProfile("dev_finder.fn_get_developer_profile", q)
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

//Add Resume Link
func AddResumeLink(c *gin.Context) {

	db := services.DB{}

	var t models.AddResumeLinkDB
	var resume models.AddResumeLink

	if err := c.ShouldBindBodyWith(&t, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	resume.ID = t.ID
	resume.ResumeURL = t.ResumeURL
	_, err := db.SAVEONDB("dev_finder.fn_add_resume_link", resume)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	var q models.GetDeveloperProfile
	q.EmailAddress = t.Email

	resp, err := db.GetDeveloperProfile("dev_finder.fn_get_developer_profile", q)
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

func UpdateResume(c *gin.Context) {

	db := services.DB{}

	var t models.AddResumeLinkDB
	var resume models.AddResumeLink

	if err := c.ShouldBindBodyWith(&t, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	resume.ID = t.ID
	resume.ResumeURL = t.ResumeURL
	_, err := db.SAVEONDB("dev_finder.fn_update_developer_resume", resume)
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

//

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
	e.Qualification_type_ = t.Qualification_type_
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
//dev_finder.fn_get_developer_resume
