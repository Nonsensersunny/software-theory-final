package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"software-theory-final/internal/config"
	"software-theory-final/internal/log"
	"software-theory-final/internal/utils"
	"software-theory-final/pkg/modules/model"
	"software-theory-final/pkg/rest"
	"software-theory-final/pkg/service"
	"strconv"
)

type App struct {
	serverConf        *config.ServerConf
}

func (app *App) init_algos() {
	algorithmService := service.NewAlgorithmService(config.GetMySQLClient())
	algorithmService.CreateAlgorithm(&model.Algorithm{
		Name: "Bayes",
		Path: "algorithm/Bayes.py",
		Description: "贝叶斯",
	})
	algorithmService.CreateAlgorithm(&model.Algorithm{
		Name: "kNN",
		Path: "algorithm/kNN.py",
		Description: "k临近算法",
	})
	algorithmService.CreateAlgorithm(&model.Algorithm{
		Name: "Logistic Regression",
		Path: "algorithm/LogisticRegression.py",
		Description: "Logistic回归",
	})
	algorithmService.CreateAlgorithm(&model.Algorithm{
		Name: "Random Forest",
		Path: "algorithm/RandomForest.py",
		Description: "随机森林",
	})
	algorithmService.CreateAlgorithm(&model.Algorithm{
		Name: "SGDC",
		Path: "algorithm/SGDC.py",
		Description: "随机梯度下降",
	})
	algorithmService.CreateAlgorithm(&model.Algorithm{
		Name: "SVM",
		Path: "algorithm/SVM.py",
		Description: "支持向量机",
	})
	algorithmService.CreateAlgorithm(&model.Algorithm{
		Name: "NN",
		Path: "algorithm/NN.py",
		Description: "全连接神经网络",
	})
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
	app.init_algos()
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
