package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Post struct {
    Title   string `json:"title"`
    Content string `json:"content"`
}

var posts []Post

func main() {
    r := gin.Default()

    // Загрузка всех постов блога
    r.GET("/posts", func(c *gin.Context) {
        c.JSON(http.StatusOK, posts)
    })

    // Добавление нового поста в блог
    r.POST("/posts", func(c *gin.Context) {
        var newPost Post
        if err := c.BindJSON(&newPost); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        posts = append(posts, newPost)
        c.JSON(http.StatusCreated, newPost)
    })

    // Обслуживание статических файлов
    r.Static("/static", "./static") // Замените "./static" на путь к вашим статическим файлам

    // Обработка запросов для index.html
    r.GET("/", func(c *gin.Context) {
        c.File("./index.html") // Замените "./index.html" на путь к вашему файлу index.html
    })

    // Запуск сервера на порту 8080
    r.Run(":8080")
}
