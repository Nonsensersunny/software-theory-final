package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"software-theory-final/internal/config"
	"software-theory-final/internal/log"
	"software-theory-final/internal/utils"
	"software-theory-final/pkg/modules/model"
	"software-theory-final/pkg/service"
)

type userRegistVO struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var userRegist userRegistVO
	err := c.ShouldBindJSON(&userRegist)
	if err != nil || userRegist.Password == "" || userRegist.Mail == "" {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.INVALID_FORM_PARAMETER))
		return
	}

	userSrv := service.NewUserService(config.GetMySQLClient())
	dbUser, err := userSrv.CreateUser(&model.User{
		Mail:     userRegist.Mail,
		Password: userRegist.Password,
	})
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.USER_REGISTER_FAIL))
		return
	}

	dataDir := fmt.Sprintf("%s/%s", utils.DATA_DIR, dbUser.Id)
	if _, err := os.Stat(dataDir); err == nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.DIR_ALREADY_EXISTS))
		return
	}
	if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
		log.Info(err, dataDir)
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.CREATE_DIR_FAIL))
		return
	}
	c.JSON(http.StatusOK, utils.RespHelper(utils.SuccessResp()))
}
