package controller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigofrumento/petChallenge/models"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	router.GET("/tutores", func(c *gin.Context) {
		// chama o método GetAllUsers
		users, err := models.GetAllUsers(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		//retornar em JSON usuários
		c.JSON(http.StatusOK, users)
	})
}
