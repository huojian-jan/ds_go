package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type user struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	dsn := "root:123456@tcp(118.31.32.197:3306)/gin_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	db.AutoMigrate(&user{})
	db.Create(&user{Name: "zhangsan", Age: 18})

	log.Println("Connected to the database successfully!")
}
