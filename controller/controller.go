package controllers

import (
	"fmt"
	"net/http"
	_ "roar-dashboard-backend/models"
	"roar-dashboard-backend/service"

	"github.com/gin-gonic/gin"
)

type IController struct {
	service service.Service
}

// func (ic *IController) HomePage(c *gin.Context) {
// 	fmt.Println("Welcome to the HomePage!")
// 	fmt.Println("Endpoint Hit: homePage")
// 	c.JSON(http.StatusOK, "albums")

// }

func (ic *IController) GetFeatureRequestsPage(c *gin.Context) {
	authToken := c.Query("token")
	fmt.Println(authToken)
	err := ic.service.GetFeatureRequestsPage(authToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	}
}

// func (ic *IController) CreatePerson(c *gin.Context) {
// 	var person []models.Person
// 	err := c.ShouldBindJSON(&person)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, "")
// 	}
// 	result, err := ic.service.CreatePersonEntry(person)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, "")
// 	}
// 	resp := gin.H{
// 		"message": "Successfully inserted all records",
// 		"data":    result,
// 	}
// 	c.JSON(http.StatusOK, resp)
// }

// func (ic *IController) GetPerson(c *gin.Context) {

// }

func NewController() IController {
	return IController{
		service: service.NewService(),
	}
}
