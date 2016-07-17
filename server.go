package main

import(
	"gopkg.in/appleboy/gin-jwt.v2"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	if port == "" {
		port = "8080"
	}

	// the jwt middleware
	authMiddleware := &jwt.GinJWTMiddleware{
	  Realm:      "test zone",
	  Key:        []byte("secret key"),
	  Timeout:    time.Hour,
	  MaxRefresh: time.Hour * 24,
	  Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
	    if (userId == "admin" && password == "admin") || (userId == "test" && password == "test") {
	      return userId, true
	    }

	    return userId, false
	  },
	  Authorizator: func(userId string, c *gin.Context) bool {
	    if userId == "admin" {
	      return true
	    }

	    return false
	  },
	  Unauthorized: func(c *gin.Context, code int, message string) {
	    c.JSON(code, gin.H{
	      "code":    code,
	      "message": message,
	    })
	  },
	}

	router.POST("/login", authMiddleware.LoginHandler)

	v1 := router.Group("/v1")
	v1.Use(authMiddleware.MiddlewareFunc())
	{
		v1.GET("/refresh_token", authMiddleware.RefreshHandler)
	}



	endless.ListenAndServe(":" + port, router)
}