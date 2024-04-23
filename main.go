package main

import (
	_ "context"
	"fmt"
	_ "log"

	//"reflect"
	"roar-dashboard-backend/database"
	Server "roar-dashboard-backend/server"
	router "roar-dashboard-backend/server/routes"
	_ "strconv"
)

func main() {
	DB, _ := database.Connect()
	fmt.Println("////////////", DB)
	server := Server.Start()
	router.HandleRequest(server)
}
