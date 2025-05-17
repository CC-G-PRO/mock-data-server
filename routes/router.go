package routes

import (
	"mockdata-server/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	login := r.Group("/login")
	{
		login.POST("/auth/google", handler.GoogleLogin)
	}

	parser := r.Group("/report")
	{
		parser.POST("/report/validate", handler.IsValidPdf)
		parser.POST("/report", handler.ParsePdf)
	}
	course := r.Group("course")
	{
		course.GET("/courses", handler.SearchSubject)
	}
	carts := r.Group("/carts")
	{
		carts.POST("", handler.AddToCart)
		carts.PATCH("", handler.UpdateCart)
		carts.DELETE("/:lectureId", handler.DeleteFromCart)
		carts.GET("", handler.GetCart)
	}

	timetables := r.Group("/timetables")
	{
		timetables.GET("", handler.GenerateTimetables)
		timetables.POST("/filter", handler.FilterTimetables)
		timetables.GET("/:id", handler.GetTimetableDetail)
	}
}
