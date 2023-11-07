package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// Токен для авторизации
const AuthToken = "tokenXcxzcasdKLDSAdxc"

// Middleware для проверки токена
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != AuthToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}

// Структура данных, которую использовал для примера
type Entity struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

// Валидатор для проверки входящих данных
var validate = validator.New()

// Пример хранилища данных
var entities = []Entity{
	{ID: "1", Name: "Entity1"},
	{ID: "2", Name: "Entity2"},
}

func main() {
	router := gin.Default()

	// Группировка маршрутов
	api := router.Group("/api", AuthMiddleware())
	{
		api.GET("/entities", getEntities)
		api.POST("/entities", createEntity)
		// Добавьте здесь дополнительные CRUD операции
	}

	router.Run(":3000") // Запускаем сервер на порту 3000
}

func getEntities(c *gin.Context) {
	c.JSON(http.StatusOK, entities)
}

func createEntity(c *gin.Context) {
	var newEntity Entity
	if err := c.ShouldBindWith(&newEntity, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	entities = append(entities, newEntity)
	c.JSON(http.StatusCreated, newEntity)
}
