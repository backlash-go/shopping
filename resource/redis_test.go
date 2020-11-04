package resource

import (
	"fmt"
	"github.com/go-redis/redis"
	"shopping/handler"
	"testing"
)

func TestGetRedis(t *testing.T) {
	RDClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "Hexian%!g5tg",
		DB:       11,
	})

	result, _ := RedisHmGet("wq", "name")
	if len(result) == 0 {
		fmt.Println("result is nil")
	}
	logger.Info("aaaa")
	fmt.Printf("%T,%v\n", result, result)
	fmt.Println(result[0])

	handler.MdSalt("xixianbin123")

}
