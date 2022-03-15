package main

import (
	"testing"

	"github.com/Crunchy89/boilerplate_golang/utils/database"
	"github.com/gin-gonic/gin"
)

func TestMain(t *testing.T) {
	db := database.SetupDatabaseConnection()
	defer database.CloseDatabaseConnection(db)
	r := gin.Default()
	r.Run(":8080")
}
