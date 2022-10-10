package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeController struct {
}

func (h HomeController) Index(c *gin.Context) {
	session := sessions.Default(c)

	user := session.Get("user")
	fmt.Println(user)

	c.HTML(http.StatusOK, "index", gin.H{
		"title": "consumer reviews",
		"user":  user,
	})
}
