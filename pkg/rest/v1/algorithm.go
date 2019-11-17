package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"software-theory-final/internal/config"
	"software-theory-final/internal/utils"
	"software-theory-final/pkg/modules/model"
	service2 "software-theory-final/pkg/service"
)

func CreateAlgorithm(c *gin.Context) {
	var algo *model.Algorithm
	if err := c.BindJSON(&algo); err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.INVALID_FORM_PARAMETER))
		return
	}
	service := service2.NewAlgorithmService(config.GetMySQLClient())
	if err := service.CreateAlgorithm(algo); err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.CREATE_ALGORITHM_FAIL))
		return
	}
	c.JSON(http.StatusOK, utils.RespHelper(utils.SuccessResp()))
}

func GetAlgorithms(c *gin.Context) {
	service := service2.NewAlgorithmService(config.GetMySQLClient())
	aid := c.Query("id")
	if aid != "" {
		algo, err := service.GetAlgorithmById(aid)
		if err != nil {
			c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.ALGORITHM_NOT_EXISTS))
			return
		}
		c.JSON(http.StatusOK, utils.RespHelper(utils.SetData("algorithm", algo)))
		return
	}
	algos, err := service.GetAllAlgorithms()
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.ALGORITHM_NOT_EXISTS))
		return
	}
	c.JSON(http.StatusOK, utils.RespHelper(utils.SetData("algorithms", algos)))
}
