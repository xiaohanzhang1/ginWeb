package main

import (
	//"net/http"

	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/xiaohanzhang1/ginWeb/router"
)

func init() {

}

func main() {
	InitRouter(gin.Default())
	fmt.Println("end")
}
