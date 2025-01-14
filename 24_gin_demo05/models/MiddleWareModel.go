package models

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GlobalMiddleWare1(c *gin.Context) {
	fmt.Println("hello, i am global middleware!")
}
