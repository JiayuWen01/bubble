package main

import (
	"bubble/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/ini.v1"
	"net/http"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func main() {
	fmt.Println("Start the learning journey of go web ...")
	var cfg = new(config.AppConfig)
	err := ini.MapTo(cfg, "./config/config.ini")
	if err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	dsn := config.GetDSN(&cfg.MySQLConfig)
	fmt.Println(dsn)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 自动迁移
	db.AutoMigrate(&Todo{})

	r := gin.Default()

	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	vGroup := r.Group("v1")
	{
		vGroup.POST("/todo", func(c *gin.Context) {
			var todo Todo
			c.BindJSON(&todo)
			db.Debug().Create(&todo)
			c.JSON(http.StatusOK, todo)
		})

		vGroup.GET("/todo", func(c *gin.Context) {
			var todoList []*Todo
			db.Debug().Find(&todoList)
			c.JSON(http.StatusOK, todoList)
		})
		vGroup.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}

			todo := db.Debug().Where("id=?", id).First(&Todo{})
			c.BindJSON(&todo)
			db.Debug().Save(todo)
			c.JSON(http.StatusOK, todo)
		})
		vGroup.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}

			fmt.Println(id)
			if err := db.Where("id=?", id).Delete(&Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
			} else {
				c.JSON(http.StatusOK, gin.H{id: "deleted"})
			}
		})
	}

	if err := r.Run(":9000"); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
