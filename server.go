package main

import (
	"github.com/NoamBoni/gofoloapp/models"
	"github.com/NoamBoni/gofoloapp/routes"
)

func main() {
	defer models.Db.Close()
	router := routes.InitRouter()
	router.Run("localhost:8000")
}
