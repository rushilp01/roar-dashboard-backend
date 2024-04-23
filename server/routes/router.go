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

// type Person struct {
// 	Id   int    `uri:"id" json:"id"`
// 	Name string `uri:"name" json:"name"`
// }
// type PId struct {
// 	Id int `uri:"id" json:"id"`
// }

func HandleRequest(server *gin.Engine) {
	controller := controllers.NewController()
	server.GET("/roar-dashboard-backend/get-feature-requests", controller.GetFeatureRequestsPage)
	// server.POST("/person", controller.CreatePerson)
	// server.GET("/person/:name", GetPerson)
	err := server.Run(":" + strconv.Itoa(8080))
	if err != nil {
		log.Panic(context.Background(), "Error in starting the server %w", err)
	}
}

// func GetPerson(c *gin.Context) {
// 	var pId Person
// 	_ = c.ShouldBindUri(&pId)
// 	record, _ := getPerson(pId)
// 	resp := gin.H{
// 		"message": "Successfully fetched all records with name = " + pId.Name,
// 		"data":    record,
// 	}
// 	c.JSON(http.StatusOK, resp)
// }

// func getPerson(pid Person) ([]Person, error) {
// 	var person []Person
// 	DB := database.Db
// 	query := "SELECT id, name FROM test where name = (?)"
// 	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancelfunc()
// 	stmt, err := DB.PrepareContext(ctx, query)
// 	if err != nil {
// 		log.Printf("Error %s when preparing SQL statement", err)
// 		return nil, err
// 	}
// 	defer stmt.Close()
// 	results, err := stmt.QueryContext(ctx, pid.Name)
// 	if err != nil {
// 		log.Printf("Error %s when inserting row into products table", err)
// 		return nil, err
// 	}
// 	for results.Next() {
// 		if err := results.Scan(&pid.Id, &pid.Name); err != nil {
// 			return nil, err
// 		}
// 		fmt.Println(reflect.TypeOf(pid), pid)
// 		person = append(person, pid)
// 	}
// 	return person, nil
// }
