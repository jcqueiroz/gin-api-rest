package main

import (
	"ALURA/jcqueiroz3/api-golang-gin/database"
	"ALURA/jcqueiroz3/api-golang-gin/routes"
)

func main() {
	database.ConnectWithDatabase()

	routes.HandleRequests()

}
