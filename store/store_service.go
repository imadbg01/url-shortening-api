package store

import (
	"context"
	// "fmt"
	// "github.com/go-redis/redis"
	// "time"
)

type StorageService struct {
	redisClient string
}

var (
	storageService = &StorageService{}
	ctx            = context.Background()
)
