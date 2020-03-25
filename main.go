package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"shopping/handler"
	"shopping/handler/middle"
	"shopping/resource"
	"syscall"
	"time"
)

func main() {
	configFile := flag.String("conf", "config/config.yaml", "path of config file")
	flag.Parse()
	viper.SetConfigFile(*configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("viper read config is failed, err is %v configFile is %s ", err, configFile)
	}
	log.Println("logger init ")
    //init mysql
	dbConf := viper.GetStringMapString("database")
	resource.InitDB(dbConf["user"], dbConf["password"], dbConf["host"], dbConf["port"], dbConf["name"])

	//init redis
	authRedisConf := viper.GetStringMapString("authRedis")
	resource.InitRedis(fmt.Sprintf("%s:%s", authRedisConf["host"], authRedisConf["port"]), authRedisConf["password"], authRedisConf["db"])

	//logger
	resource.InitLogger()

	e := echo.New()
	//e.Use(handle.ResTime)

	e.Use(middle.LoginValidate)

//uuid
//	u1 := (uuid.NewV4())
//	fmt.Printf("UUIDv4: %s\n", u1)
//	fmt.Printf("%T,%s\n", u1,u1)
//
//	u2 := u1.String()
//	fmt.Printf("%T,%s\n", u2,u2)
	for _, h := range handler.GetRouters() {
		e.Add(h.Method, h.Path, h.Hf)
	}
	e.Logger.Fatal(e.Start(":1234"))

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	s := <-signalCh
	fmt.Println(s)
	switch s {
	case syscall.SIGTERM, syscall.SIGINT:
		c, _ := context.WithTimeout(context.Background(), time.Second*5)
		_ = e.Shutdown(c)
		log.Println("receive signal is ", s)
		os.Exit(0)
	}
}