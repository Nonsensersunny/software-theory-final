package rest

import (
	"github.com/gin-gonic/gin"
	"software-theory-final/pkg/rest/v1"
)

func REST(engine *gin.Engine) {
	userRoute := engine.Group("/user")
	{
		userRoute.POST("/", v1.Register)
	}

	datasetRoute := engine.Group("/dataset")
	{
		datasetRoute.POST("/", v1.UploadDataset)
	}
}
