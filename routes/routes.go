package routes

import (
	"ALURA/jcqueiroz3/api-golang-gin/controllers"
	"ALURA/jcqueiroz3/api-golang-gin/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/students", controllers.ShowAllStudents)
	r.GET("/:name", controllers.Welcome)
	r.POST("/students", controllers.CreateNewStudent)
	r.GET("/students/:id", controllers.FindStudentByID)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.EditStudent)
	r.GET("/students/cpf/:cpf", controllers.FindStudentByCPF)
	r.GET("/index", controllers.ShowPageIndex)
	r.NoRoute(controllers.RouteNotFound)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":5000")
}
