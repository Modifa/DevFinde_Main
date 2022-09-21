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
	}
}
