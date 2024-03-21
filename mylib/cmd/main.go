package main

import (
	"mylib/database"
	"mylib/routers"
)

func main() {
	db := database.SQLInit()

	gorm := database.GormInit(db)

	defer db.Close()

	start := routers.StartServer(db, gorm)
	start.Run(":8080")
}
