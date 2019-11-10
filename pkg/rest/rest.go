package rest

import (
	"github.com/gin-gonic/gin"
	"software-theory-final/pkg/rest/v1"
)

func REST(engine *gin.Engine) {
	userRoute := engine.Group("/user")
	{
		userRoute.PUT("/", v1.Register)
		userRoute.POST("/", v1.Login)
		userRoute.GET("/", v1.GetLoginStatus)
	}

	datasetRoute := engine.Group("/dataset")
	{
		datasetRoute.PUT("/", v1.UploadDataset)
		datasetRoute.GET("/", v1.GetDataSets)
	}

	algorithmRoute := engine.Group("/algorithm")
	{
		algorithmRoute.PUT("/", v1.CreateAlgorithm)
		algorithmRoute.GET("/", v1.GetAlgorithms)
	}

	predictionRoute := engine.Group("/prediction")
	{
		predictionRoute.POST("/", v1.CreatePrediction)
		predictionRoute.GET("/", v1.GetPredictions)
	}
}
