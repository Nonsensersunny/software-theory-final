package rest

import (
	"github.com/gin-gonic/gin"
	v1 "software-theory-final/backend/pkg/rest/v1"
)

func REST(engine *gin.Engine) {
	userRoute := engine.Group("/user")
	{
		userRoute.POST("/", v1.Register)
	}

	datasetRoute := engine.Group("/dataset")
	{
		datasetRoute.POST("/file", v1.UploadDataset)
	}
}
