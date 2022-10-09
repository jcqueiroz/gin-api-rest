package controllers

import (
	"ALURA/jcqueiroz3/api-golang-gin/database"
	"ALURA/jcqueiroz3/api-golang-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

func ShowAllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(200, students)

}

func Welcome(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API say:": "Hello" + name + ", all right?",
	})

}

// CreateNewStudent godoc
// @Summary      Create new Student values
// @Description  Create New Student
// @Tags         Add-swagger-by-jcqueiroz
// @Accept       json
// @Produce      json
// @Param        students body models.Student true "Models of student"
// @Success      200  {object}  models.Student
// @Failure      400  {object}  httputil.HTTPError
// @Router       /students [post]
func CreateNewStudent(c *gin.Context) {
	var Student models.Student
	if err := c.ShouldBindJSON(&Student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.CheckDataStudent(&Student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&Student)
	c.JSON(http.StatusOK, Student)
}

// FindStudentByID godoc
// @Summary      Find Student By ID values
// @Description  Find Student Created
// @Tags         Add-swagger-by-jcqueiroz
// @Accept       json
// @Produce      json
// @Param        id   path     int    true     "Student ID"
// @Success      200  {object}  models.Student
// @Failure      400  {object}  httputil.HTTPError
// @Router       /students/{id} [get]
func FindStudentByID(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Student not found!"})
		return
	}

	c.JSON(http.StatusOK, student)
}

// DeleteStudent godoc
// @Summary      Delete Student values
// @Description  Delete Student Created
// @Tags         Add-swagger-by-jcqueiroz
// @Accept       json
// @Produce      json
// @Param        id   path     int    true     "Student ID"
// @Success      200  {object}  models.Student
// @Failure      400  {object}  httputil.HTTPError
// @Router       /students/{id} [delete]
func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.Unscoped().Delete(&student, id)
	c.JSON(http.StatusOK, gin.H{"data": "Student Selected Deleted with success"})
}

// EditStudent godoc
// @Summary      Edit an Student
// @Description  Edit by json Student
// @Tags         Add-swagger-by-jcqueiroz
// @Accept       json
// @Produce      json
// @Param        id   path      int    true     "Student ID"
// @Param        id   body      models.Student    true     "Edit Student"
// @Success      200  {object}  models.Student
// @Failure      400  {object}  httputil.HTTPError
// @Router       /students/:id [patch]
func EditStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.CheckDataStudent(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Save(&student)
	c.JSON(http.StatusOK, student)
}

// FindStudentByCPF godoc
// @Summary      Find Student By CPF values
// @Description  Find Student Created
// @Tags         Add-swagger-by-jcqueiroz
// @Accept       json
// @Produce      json
// @Param        cpf    query     string  false  "student search by cpf"  Format(email)
// @Success      200  {object}  models.Student
// @Failure      400  {object}  httputil.HTTPError
// @Router       /students [get]
func FindStudentByCPF(c *gin.Context) {
	var student models.Student
	cpf := c.Param("cpf")
	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Student not found!"})
		return
	}
	c.JSON(http.StatusOK, student)

}

func ShowPageIndex(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"students": students,
	})

}

func RouteNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
