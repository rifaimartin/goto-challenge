package main

import (
	"log"
	"pos-app/auth"
	"pos-app/cashier"
	"pos-app/handler"
	"pos-app/helper"

	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:rifaimartin123@tcp(127.0.0.1:3306)/pos-app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	cashierRepository := cashier.NewRepository(db)

	cashierService := cashier.NewService(cashierRepository)
	authService := auth.NewService()


	cashierHandler := handler.NewcashierHandler(cashierService, authService)

	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api/v1")

	api.GET("/cashiers", cashierHandler.GetAllCashiers)
	api.GET("/cashiers/:cashierId", cashierHandler.GetcashierByID)
	api.GET("/cashiers/:cashierId/passcode", cashierHandler.GetPassCode)

	api.POST("/cashiers/:cashierId/login", cashierHandler.Login)
	api.POST("/cashiers/:cashierId/logout", cashierHandler.Logout)
	api.POST("/cashiers", cashierHandler.CreateCashier)
	api.PUT("/cashiers/:cashierId", cashierHandler.UpdateCashier)	
	api.DELETE("/cashiers/:cashierId", cashierHandler.DeleteCashier)

	router.Run()
}


func authMiddleware(authService auth.Service, cashierService cashier.Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			fmt.Println("Bearer empity")
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, false, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// split by space example Bearer = Token
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, false, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, false, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		cashierID := int(claim["cashier_id"].(float64))

		cashier, err := cashierService.GetcashierByID(cashierID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, false, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// context currentcashier
		c.Set("currentcashier", cashier)
	}
}

