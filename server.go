package main

import (
	"github.com/NoamBoni/gofoloapp/controllers"
	"github.com/NoamBoni/gofoloapp/routes"
)

func main() {
	defer controllers.Db.Close()
	
	router := routes.InitRouter()
	router.Run("localhost:8000")
}

