package rest

import (
	"github.com/gin-gonic/gin"
	v1 "software-theory-final/backend/pkg/rest/v1"
)

func REST(engine *gin.Engine) {
	userRoute := engine.Group("/user")
	{
		userRoute.POST("/join", v1.Register)
	}

	datasetRoute := engine.Group("/algo")
	{
		datasetRoute.POST("/file", v1.UploadDataset)
	}
}
