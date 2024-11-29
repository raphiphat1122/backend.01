package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Employee routes
	r.GET("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET EMPLOYEE METHOD",
		})
	})

	r.POST("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST EMPLOYEE METHOD",
		})
	})

	r.PUT("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT EMPLOYEE METHOD",
		})
	})

	r.DELETE("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE EMPLOYEE METHOD",
		})
	})

	// Customer routes
	r.GET("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET CUSTOMER METHOD",
		})
	})

	r.POST("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST CUSTOMER METHOD",
		})
	})

	r.PUT("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT CUSTOMER METHOD",
		})
	})

	r.DELETE("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE CUSTOMER METHOD",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
