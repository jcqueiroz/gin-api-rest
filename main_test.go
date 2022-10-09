package main

import (
	"ALURA/jcqueiroz3/api-golang-gin/controllers"
	"ALURA/jcqueiroz3/api-golang-gin/database"
	"ALURA/jcqueiroz3/api-golang-gin/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateStudentMock() {
	student := models.Student{Name: "Name of student test", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&student)
	ID = int(student.ID)

}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestVerifyStatusCodeInWelcomeWithParameter(t *testing.T) {
	r := SetupTestRoutes()
	r.GET("/:name", controllers.Welcome)
	req, _ := http.NewRequest("GET", "/jean", nil)
	answer := httptest.NewRecorder()
	r.ServeHTTP(answer, req)
	assert.Equal(t, http.StatusOK, answer.Code, "should be the same")
	mockOfTheAnswer := `{"API say:": "Hellojean, all right?"}`
	answerBody, _ := ioutil.ReadAll(answer.Body)
	assert.Equal(t, mockOfTheAnswer, string(answerBody))
	fmt.Println(answerBody)
	fmt.Println(mockOfTheAnswer)
}

func TestListAllStudentsHandler(t *testing.T) {
	database.ConnectWithDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupTestRoutes()
	r.GET("/students", controllers.ShowAllStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	answer := httptest.NewRecorder()
	r.ServeHTTP(answer, req)
	assert.Equal(t, http.StatusOK, answer.Code)
	fmt.Println(answer.Body)
	fmt.Println(answer.Body)
}

func TestSearchStudentByCPFHandler(t *testing.T) {
	database.ConnectWithDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupTestRoutes()
	r.GET("/students/cpf/:cpf", controllers.FindStudentByCPF)
	req, _ := http.NewRequest("GET", "/students/cpf/12345678901", nil)
	answer := httptest.NewRecorder()
	r.ServeHTTP(answer, req)
	assert.Equal(t, http.StatusOK, answer.Code)
}

func TestFindStudentByIdHandler(t *testing.T) {
	database.ConnectWithDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupTestRoutes()
	r.GET("/students/:id", controllers.FindStudentByID)
	pathSearch := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathSearch, nil)
	answer := httptest.NewRecorder()
	r.ServeHTTP(answer, req)
	var studentMock models.Student
	json.Unmarshal(answer.Body.Bytes(), &studentMock)
	assert.Equal(t, "Name of student test", studentMock.Name, "The names should be equals")
	assert.Equal(t, "12345678901", studentMock.CPF)
	assert.Equal(t, "123456789", studentMock.RG)
	assert.Equal(t, http.StatusOK, answer.Code)
}

func TestDeleteStudentHandler(t *testing.T) {
	database.ConnectWithDatabase()
	CreateStudentMock()
	r := SetupTestRoutes()
	r.DELETE("/students/:id", controllers.DeleteStudent)
	pathSearch := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathSearch, nil)
	answer := httptest.NewRecorder()
	r.ServeHTTP(answer, req)
	assert.Equal(t, http.StatusOK, answer.Code)
}

func TestEditOneStudentHandler(t *testing.T) {
	database.ConnectWithDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupTestRoutes()
	r.PATCH("/students/:id", controllers.EditStudent)
	student := models.Student{Name: "Name of student test", CPF: "47123456789", RG: "123456700"}
	valueJson, _ := json.Marshal(student)
	pathForEdit := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathForEdit, bytes.NewBuffer(valueJson))
	answer := httptest.NewRecorder()
	r.ServeHTTP(answer, req)
	var studentMockUpdated models.Student
	json.Unmarshal(answer.Body.Bytes(), &studentMockUpdated)
	assert.Equal(t, "47123456789", studentMockUpdated.CPF)
	assert.Equal(t, "123456700", studentMockUpdated.RG)
	assert.Equal(t, "Name of student test", studentMockUpdated.Name)
}
