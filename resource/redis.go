package resource

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

var pool *redis.Pool

func GetRedisPool() *redis.Pool  {
	return pool
}

func InitRedis(add,pwd,choiceDB string)  {
	log.Println("initRedis choice db is ", choiceDB)
	pool = &redis.Pool{
		TestOnBorrow: nil,
		MaxIdle:      10000,
		MaxActive:    10000,
		IdleTimeout:  5 * time.Second,
		Wait:         false,
		Dial: func() (conn redis.Conn, err error) {
			c, err := redis.Dial("tcp",add)
			if err != nil{
				return nil,err
			}
			if pwd == "" {
				return c,nil
			}
			if _, err := c.Do("AUTH",pwd); err != nil{
				_ = c.Close()
				return nil, err
			}
			if _, err := c.Do("SELECT",choiceDB); err != nil{
				_ = c.Close()
				return nil,err
			}

			return c,nil

		},
	}


}



