package controllers

import (
	"net/http"
	"roar-dashboard-backend/models"
	_ "roar-dashboard-backend/models"
	"roar-dashboard-backend/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IController struct {
	service service.Service
}

func (ic *IController) GetFeatureRequestsPage(c *gin.Context) {
	feature_id := c.Param("id")
	feature, err := strconv.Atoi(feature_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	}
	feature_details, err := ic.service.GetFeatureRequestsPage(feature)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	}
	c.JSON(http.StatusOK, feature_details)
}

func (ic *IController) GetAllFeatureRequestsPage(c *gin.Context) {
	all_feature_details, err := ic.service.GetAllFeatureRequestsPage()
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	}
	c.JSON(http.StatusOK, all_feature_details)
}

func (ic *IController) CreateFeatureRequest(c *gin.Context) {
	var createFeature models.FeatureRequest
	err := c.ShouldBindJSON(&createFeature)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	}
	err = ic.service.CreateFeature(createFeature)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
	}
	c.JSON(http.StatusOK, map[string]interface{}{"message": "success", "status_code": 200})
}

func NewController() IController {
	return IController{
		service: service.NewService(),
	}
}
