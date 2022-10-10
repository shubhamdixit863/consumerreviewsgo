package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListingController struct {
}

func (h ListingController) GetListingPage(c *gin.Context) {
	session := sessions.Default(c)

	user := session.Get("user")
	if user != nil {
		c.HTML(http.StatusOK, "add-listing", gin.H{
			"title": "add listing",
			"user":  user,
		})

	}

	c.HTML(http.StatusOK, "index", gin.H{
		"title": "index",
		"user":  user,
	})

}
