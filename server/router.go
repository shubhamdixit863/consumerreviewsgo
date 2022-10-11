package server

import (
	"consumerreviewsgo/controllers"
	"consumerreviewsgo/middlewares"
	"consumerreviewsgo/models"
	"encoding/gob"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.HTMLRender = ginview.Default()
	router.Static("/static", "./static")
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session", store))
	gob.Register(models.User{})

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	v1 := router.Group("/")
	{
		homeGroup := v1.Group("/")
		{
			home := new(controllers.HomeController)
			homeGroup.GET("/", home.Index)
		}

		authGroup := v1.Group("auth")
		{
			auth := new(controllers.AuthController)
			authGroup.GET("/login", auth.Login)
			authGroup.GET("/signup", auth.Signup)
			authGroup.POST("/signup", auth.SignupPost)
		}

		listingGroup := v1.Group("listing")
		{
			listing := new(controllers.ListingController)
			listingGroup.GET("/post", listing.GetListingPage)

		}
	}
	router.Use(middlewares.AuthMiddleware())
	return router

}
