package main

import (
	"datastore-service/constants"
	"datastore-service/controller"
	"datastore-service/data_store"
	"datastore-service/pkg"
	"datastore-service/services"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Data store service")

	// set mode
	var mode constants.ServiceMode
	if len(os.Args) > 1 {
		mode = constants.ServiceMode(os.Args[1])
	} else {
		// Test mode is the default mode
		mode = constants.TestMode
	}

	if mode != constants.TestMode && mode != constants.DataStorePodMode {
		fmt.Println("Unknown mode selected")
		return
	}

	fmt.Printf("Mode -> %s\n", mode)

	// initialise app
	err := pkg.InitApp(mode)
	if err != nil {
		fmt.Printf("Error -> %s\n", err.Error())
		return
	}

	r := gin.Default()

	// set controllers
	controller.ExposeEndpoints(r, mode)

	// init database connection
	err = data_store.InitDataStore(constants.DATABASE_NAME, constants.TABLE_NAME)
	if err != nil {
		return
	}

	services.InitFileHandlerRepo(data_store.GetRepository())
	r.Run()
}
