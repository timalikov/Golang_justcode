package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var ctx = context.Background()

type Data struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}

func getDataFromPostgreSQL() (*Data, error) {
	// Здесь должна быть реализация получения данных из PostgreSQL
	// Для примера возвращаем фиктивные данные
	return &Data{
		Field1: "example data",
		Field2: 123,
	}, nil
}

func setDataToRedis(rdb *redis.Client, key string, data *Data) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Устанавливаем данные в Redis с TTL в 5 минут
	if err := rdb.Set(ctx, key, dataBytes, 5*time.Minute).Err(); err != nil {
		return err
	}

	return nil
}

func getDataFromRedis(rdb *redis.Client, key string) (*Data, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil // Ключ не найден
	} else if err != nil {
		return nil, err
	}

	var data Data
	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func main() {
	// Инициализация клиента Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	defer rdb.Close()

	key := "my_data_key"

	cachedData, err := getDataFromRedis(rdb, key)
	if err != nil {
		log.Fatalf("Failed to get data from Redis: %v", err)
	}

	if cachedData == nil {
		data, err := getDataFromPostgreSQL()
		if err != nil {
			log.Fatalf("Failed to get data from PostgreSQL: %v", err)
		}

		err = setDataToRedis(rdb, key, data)
		if err != nil {
			log.Fatalf("Failed to set data to Redis: %v", err)
		}

		fmt.Println("Data from PostgreSQL:", data)
	} else {
		fmt.Println("Data from Redis:", cachedData)
	}
}
