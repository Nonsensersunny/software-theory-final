package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"software-theory-final/backend/internal/config"
	"software-theory-final/backend/internal/log"
	"software-theory-final/backend/internal/utils"
	"software-theory-final/backend/pkg/rest"
	"strconv"
)

type App struct {
	serverConf        *config.ServerConf
}

func Init() (app App) {
	log.Info("Initializing APP")
	var sc config.ServerConf
	serverConfig := sc.GetConfig()

	app = App{
		serverConf: serverConfig,
	}

	config.InitDB(*serverConfig)
	config.InitScheme()
	config.InitMailClient(app.serverConf.Mail)

	if _, err := os.Stat(utils.DATA_DIR); err != nil {
		log.Infof("no such dir named [%s] and now creating", utils.DATA_DIR)
	}
	if err := os.Mkdir(utils.DATA_DIR, os.ModePerm); err != nil {
		log.Errorf("creating data dir found error: %v", err)
	}
	return
}

func main() {
	app := Init()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = app.serverConf.Http.AllowOrigin
	fmt.Println("[INFO] CrosOrigin:", app.serverConf.Http.AllowOrigin)
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))

	rest.REST(r)

	r.Run(":" + strconv.Itoa(app.serverConf.Http.Port))
}
