package routes

import (
	"context"
	_ "database/sql"
	_ "fmt"
	"log"
	_ "net/http"
	_ "reflect"
	controllers "roar-dashboard-backend/controller"
	_ "roar-dashboard-backend/database"
	"strconv"
	_ "time"

	"github.com/gin-gonic/gin"
)

func HandleRequest(server *gin.Engine) {
	controller := controllers.NewController()
	server.GET("/roar-dashboard-backend/get-feature-requests/:id", controller.GetFeatureRequestsPage)
	server.GET("/roar-dashboard-backend/get-feature-requests", controller.GetAllFeatureRequestsPage)
	server.POST("/roar-dashboard-backend/create-feature-requests", controller.CreateFeatureRequest)
	err := server.Run(":" + strconv.Itoa(8080))
	if err != nil {
		log.Panic(context.Background(), "Error in starting the server %w", err)
	}
}
