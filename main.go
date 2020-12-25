package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func main() {
	r := gin.Default()
	//配置模板文件的静态文件路径
	r.Static("/static", "static")
	//配置模板文件路径
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//待办事项
	//添加
	//查看
	//修改
	//删除
	r.Run()
}
