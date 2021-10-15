package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Todo struct {
	gorm.Model
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func main() {
	db, err := gorm.Open("mysql", "root:root1234@(127.0.0.1:13306)/bubble?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 自动迁移
	db.AutoMigrate(&Todo{})
	fmt.Println("Start the learning journey of go web ...")
}
