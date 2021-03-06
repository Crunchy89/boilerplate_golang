package main

import (
	"github.com/Crunchy89/boilerplate_golang/utils/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = database.SetupDatabaseConnection()
)

func main() {
	defer database.CloseDatabaseConnection(db)
	r := gin.Default()
	r.Run(":8080")
}
