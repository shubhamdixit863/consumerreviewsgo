package controllers

import (
	"consumerreviewsgo/forms"
	"consumerreviewsgo/models"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type AuthController struct {
}

func (h AuthController) Signup(c *gin.Context) {
	dat, err := os.ReadFile("./views/register.html")
	if err != nil {
		fmt.Println(err)
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", dat)
	//c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(" <div class=\"modal\"> <!-- Login modal --> <div id=\"login-modal\"> <div class=\"login-from register\"> <h3>Create an account</h3> <form  method=\"POST\" ><div class=\"form-group\"> <input type=\"text\" name=\"username\" id=\"username_register \" placeholder=\"User Name\" required=\"\"> </div> <div class=\"form-group\"> <input type=\"email\" id=\"email \" name=\"email\" placeholder=\"Email\" required=\"\"> </div> <div class=\"form-group\"> <input type=\"password\" id=\"password \" name=\"password\" placeholder=\"Password\" required=\"\"> </div> <div class=\"form-group full-width\"> <div class=\"checkbox-wrap\"> <input type='checkbox' name='remember-me' value='remember-me' id=\"agree-policy\" /> <label for=\"agree-policy\">I agree to the <a href=\"#\">Privacy Policy</a></label> </div> <div class=\"checkbox-wrap\"> <input type='checkbox' name='remember-me' value='remember-me' id=\"agree-terms\" /> <label for=\"agree-terms\">I agree to the <a href=\"#\">Terms and Conditions</a></label> </div> </div> <div class=\"form-group\"> <button type=\"button\"  id=\"registerButton\"  class=\"theme-btn btn-style-one\">Register</button> </form></div> <div class=\"form-group\"> <span class=\"text\">Or connect with</span> </div> <div class=\"form-group\"> <a href=\"#\" class=\"social-btn facebook-btn\"><span class=\"fab fa-facebook-f\"></span> Facebook</a> <a href=\"#\" class=\"social-btn google-btn\"><span class=\"fab fa-google\"></span> Google</a> </div> </div> </div> <!-- End Login Module --> </div> "))
}

func (h AuthController) SignupPost(c *gin.Context) {
	var requestBody forms.UserSignup

	session := sessions.Default(c)
	user := session.Get("user")
	if user != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Please logout first"})
		return
	}

	if err := c.BindJSON(&requestBody); err != nil {

	}
	signup, err := (models.User{}).Signup(requestBody)
	if err != nil {
		c.Data(http.StatusInternalServerError, "application/json", []byte(err.Error()))

	}

	session.Set("user", signup)
	if err := session.Save(); err != nil {
		c.Data(http.StatusOK, "application/json", []byte(err.Error()))
		return
	}
	c.Data(http.StatusOK, "application/json", []byte(`{"message":"Success"}`))
}

func (h AuthController) Login(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<div class=\"model\"> <!-- Login modal --> <div id=\"login-modal\"> <div class=\"login-from\"> <h3>Login</h3> <div class=\"form-group\"> <input type=\"text\" name=\"username\" placeholder=\"User Name\" required=\"\"> </div> <div class=\"form-group\"> <input type=\"password\" name=\"password\" placeholder=\"Password\" required=\"\"> </div> <div class=\"form-group\"> <div class=\"checkbox-wrap\"> <input type='checkbox' name='remember-me' value='remember-me' id=\"remember-me\" /> <label for=\"remember-me\">Remember Me</label> </div> <a href=\"#\" class=\"forgot-pass\">Forgot Password?</a> </div> <div class=\"form-group\"> <button type=\"submit\" class=\"theme-btn btn-style-one\">Login</button> </div> <div class=\"form-group\"> <span class=\"text\">Or connect with</span> </div> <div class=\"form-group\"> <a href=\"#\" class=\"social-btn facebook-btn\"><span class=\"fab fa-facebook-f\"></span> Facebook</a> <a href=\"#\" class=\"social-btn google-btn\"><span class=\"fab fa-google\"></span> Google</a> </div> <div class=\"form-group bottom-text\"> <div class=\"text\">Don’t have an account?</div> <a href=\"#\" class=\"signup-link call-modal\">Register</a> </div> </div> </div> <!-- End Login Module --> </div>"))

}
