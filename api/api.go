package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"time"
)

func InitialAPI(listener net.Listener) {
	logrus.Info("initial api start")
	router := SetupRouter()
	err := http.Serve(listener, router)
	if err != nil {
		panic(err)
	}
}

const appName = "appName"
const secretKey = "aiadvanc"
const authorization = "Authorization"

type User struct {
	Id           string `json:"id"`
	MobileNumber string `json:"mobileNumber"`
	NkpUserId    string `json:"nkpUserId"`
	Channel      string `json:"channel"`
}

type Response struct {
	Code  string      `json:"code"`
	Data  interface{} `json:"data"`
	MSG   string      `json:"message"`
	Extra string      `json:"extra"`
}

func HandlerInterceptorAdapter() gin.HandlerFunc {
	//Accept-Language : en-US
	//Accept-Language : id
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("X-ADV-TOKEN")
		if err != nil {
			ctx.Header("Content-Type", "application/json")
			ctx.JSON(http.StatusUnauthorized, &Response{Code: "UNAUTHORIZED", MSG: "Unauthorized request calling to get CurrentUserId"})
			ctx.Abort()
			return
		}
		var user User
		err = ParseToken(token, secretKey, &user)
		if err != nil {
			ctx.Header("Content-Type", "application/json")
			ctx.JSON(http.StatusUnauthorized, &Response{Code: "UNAUTHORIZED", MSG: "Unauthorized request calling to get CurrentUserId"})
			ctx.Abort()
			return
		}
		fmt.Println(token)
		t := time.Now()
		// Set example variable
		ctx.Set("example", "12345")
		// before request
		ctx.Next()
		// after request
		latency := time.Since(t)
		logrus.Info(latency)
		// access the status we are sending
		status := ctx.Writer.Status()
		logrus.Info(status)
	}
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(HandlerInterceptorAdapter())
	router.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "UP", "message": "OK"}) })
	router.GET("/test", func(c *gin.Context) {
		logrus.Info("test api")
		c.JSON(200, gin.H{"status": "UP", "message": "OK"})
	})
	return router
}
