package main

import (
    "github.com/gin-gonic/gin"

    "gin_jwt/apis"
    "gin_jwt/middleware"
)

func main() {
    r := gin.Default()
    r.POST("/login", apis.Login)
    r.POST("/register", apis.Register)

    taR := r.Group("/data")
    taR.Use(middleware.JWTMiddleWare)
    {
        taR.GET("/test", apis.Test)
    }
    r.Run(":8080")
}
