package service

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"shopping/resource"
)

func ExistKey(key string) (int,error) {
	conn := resource.GetRedisPool().Get()
	defer conn.Close()
	return redis.Int(conn.Do("EXISTS",key))
	
}

func SetValue(key, value string, expires int)  error{
	conn := resource.GetRedisPool().Get()
	defer conn.Close()
	fmt.Println("aaa")
	_, err := conn.Do("SETEX",key,expires,value)
	return err
}

func GetValue(key string) (string, error) {
	conn := resource.GetRedisPool().Get()
	defer conn.Close()
	return redis.String(conn.Do("GET", key))
}

//func HmsetValue(tk string,cp string,em string)  {
//	var p1 struct {
//		Token  string `redis:"token"`
//		Cellphone string `redis:"cellphone"`
//		Email   string `redis:"email"`
//	}
//
//	conn := resource.GetRedisPool().Get()
//	defer conn.Close()
//	//_, err := conn.Do("HMSET",key,redis.Args{}.Add(key))
//	if _, err := conn.Do("HMSET", redis.Args{}.Add("id1").AddFlat(&p1)...); err != nil {
//		panic(err)
//	}
//
//	M := map[string]string{
//		"token":  tk,
//		"cellphone": cp,
//		"email":   em,
//	}
//
//	if _, err := conn.Do("HMSET", redis.Args{}.Add("id2").AddFlat(M)...); err != nil {
//		panic(err)
//	}
//}



func  HMSet(key ,cp , em string, expire  int) (err error) {
	M := map[string]string{
				//"token":  tk,
				"cellphone": cp,
				"email":   em,
			}
	conn := resource.GetRedisPool().Get()
	defer conn.Close()
	if _, err := conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(M)...); err != nil{
		panic(err)
	}

	if expire > 0 {
		if _, err = conn.Do("EXPIRE", key, int64(expire)); err!= nil{
			panic(err)
		}
	}
	if err != nil {
		return
	}
	return
}
