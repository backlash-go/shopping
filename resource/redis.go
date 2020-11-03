package resource

import (
	"errors"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"strings"
	"time"
)

//var pool *redis.Pool
var RDClient *redis.Client

func InitRedis() error {
	dsn := strings.Join([]string{viper.GetString("authRedis.host"), viper.GetString("authRedis.port")}, ":")
	RDClient = redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: viper.GetString("authRedis.password"),
		DB:       viper.GetInt("authRedis.db"),
	})
	pong, err := RDClient.Ping().Result()
	if err != nil {
		return err
	}
	logger.Info(pong)
	err = RDClient.Set("course_homework-running", "yeah", 0).Err()
	if err != nil {
		return errors.New("redis Set failed: " + dsn)
	}
	logger.Info("init redis")
	return nil
}

// GetRedis 获取redis链接实例
func GetRedis() *redis.Client {
	return RDClient
}

func SetHashValue(key string, value map[string]interface{}) error {
	err := RDClient.HMSet(key, value).Err()
	return err
}

func SetKeyTtl(key string, time time.Duration) error {
	err := RDClient.Expire(key, time).Err()
	return err
}

func RedisHmGet(key string,fields ...string) ([]interface{}, error) {
	s, err := RDClient.HMGet(key,fields...).Result()
	RDClient.Exists().Result()
	return s, err
}

func RedisGet(key string) (string, error) {
	result, err := RDClient.Get(key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return result, nil
}

func RedisExistKey(key string) bool {
	r, err := RedisGet(key)
	if r == "" || err != nil {
		return false
	} else {
		return true
	}
}
