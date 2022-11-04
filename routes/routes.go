package routes

import (
	cont "github.com/Modifa/DevFinde_Main/controller"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	V1 := r.Group("/api/devfinder/")
	{
		V1.POST("GetDeveloperProfile", cont.GetDeveloperProfile)
		V1.POST("RegisterDeveloper", cont.RegisterDeveloper)
		V1.POST("AddDeveloperLink", cont.AddDeveloperLink)
		V1.POST("AddDeveloperExperience", cont.AddDeveloperExperience)
		V1.POST("UpdateDeveloperProfile", cont.UpdateDeveloperProfile)
		V1.GET("GetDBProfile", cont.GetDBProfile)
		V1.POST("UpdateImage", cont.UpdateImage)
		//Untested
		V1.POST("AddResume", cont.AddResume)
		V1.POST("AddResumeLink", cont.AddResumeLink)
		V1.POST("UpdateResume", cont.UpdateResume)
		V1.POST("AddEducation", cont.AddEducation)
	}
}
