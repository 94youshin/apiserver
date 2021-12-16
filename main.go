package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/youshintop/apiserver/config"
	"github.com/youshintop/apiserver/router"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {

	pflag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	gin.SetMode(viper.GetString("mode"))

	g := gin.New()

	middlerwares := []gin.HandlerFunc{}

	router.Load(g, middlerwares...)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address %s", viper.GetString("address"))
	log.Printf(http.ListenAndServe(viper.GetString("address"), g).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get("http://127.0.0.1:8080/sd/health")

		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		log.Print("Waitting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}

	return errors.New("cannot connect to the router.")
}
