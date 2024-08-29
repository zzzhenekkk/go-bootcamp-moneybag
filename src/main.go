package main

import (
	_ "day7/docs" // Замените на правильный путь
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"strconv"
)

// @title Swagger Example API
// @description API for managing coin calculations.
// @version 1.0
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /

func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/mincoins", GetMinCoins)
	r.GET("/mincoins2", GetMinCoins2)

	r.Run() // listen and serve on 0.0.0.0:8080
}

// GetMinCoins handles the HTTP request and calls minCoins to get the result.
func GetMinCoins(c *gin.Context) {
	valStr := c.Query("val")
	val, err := strconv.Atoi(valStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for 'val'"})
		return
	}

	coinsStr := c.QueryArray("coins")
	coins := make([]int, len(coinsStr))
	for i, coinStr := range coinsStr {
		coins[i], err = strconv.Atoi(coinStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value in 'coins'"})
			return
		}
	}

	result := minCoins(val, coins)
	c.JSON(http.StatusOK, result)
}

// GetMinCoins2 handles the HTTP request and calls minCoins2 to get the result.
func GetMinCoins2(c *gin.Context) {
	valStr := c.Query("val")
	val, err := strconv.Atoi(valStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for 'val'"})
		return
	}

	coinsStr := c.QueryArray("coins")
	coins := make([]int, len(coinsStr))
	for i, coinStr := range coinsStr {
		coins[i], err = strconv.Atoi(coinStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value in 'coins'"})
			return
		}
	}

	result := minCoins2(val, coins)
	c.JSON(http.StatusOK, result)
}
