package admin

import (
	"fmt"
	"gormdemo01/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}

type AdminController struct {
	BaseController
}

func (con AdminController) Admin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/index.html", gin.H{
		"title": "index page",
	})
}

func (con AdminController) News(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "default/news.html", gin.H{
		"title": "news page",
	})
}

func (con AdminController) AdminNews(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/news.html", gin.H{
		"title": "news page",
	})
}

type UserController struct {
	BaseController
}

func (userController UserController) Users(c *gin.Context) {
	//query db
	userList := []models.User{}
	models.DB.Find(&userList)

	c.JSON(http.StatusOK, userList)
}

func (userController UserController) DeleteUser(c *gin.Context) {
	var user models.User
	uid := c.Param("uid")
	models.DB.Delete(&user, uid)
	c.JSON(http.StatusOK, gin.H{
		"message": "user was deleted with id " + uid,
	})

}

func (userController UserController) AddUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		userController.failed(c)
	} else {
		fmt.Println("retrive data:", user)
		result := models.DB.Create(&user)
		if result.Error != nil {
			userController.failed(c)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message":  "user added!",
				"userInfo": user,
			})
		}
	}
}

func (con UserController) EditUser(c *gin.Context) {
	c.String(http.StatusOK, "user Edited")
}
