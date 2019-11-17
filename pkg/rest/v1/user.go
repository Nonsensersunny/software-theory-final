package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
	"os"
	"software-theory-final/internal/config"
	"software-theory-final/internal/log"
	"software-theory-final/internal/utils"
	"software-theory-final/pkg/modules/model"
	"software-theory-final/pkg/service"
	"time"
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

	dataDir := utils.GetDataSetsDir(dbUser.Id)
	predDir := utils.GetPredictionsDir(dbUser.Id)
	if _, err := os.Stat(dataDir); err == nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.DIR_ALREADY_EXISTS))
		return
	}
	if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
		log.Info(err, dataDir)
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.CREATE_DIR_FAIL))
		return
	}
	if _, err := os.Stat(predDir); err == nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.DIR_ALREADY_EXISTS))
		return
	}
	if err := os.MkdirAll(predDir, os.ModePerm); err != nil {
		log.Info(err, predDir)
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.CREATE_DIR_FAIL))
		return
	}
	c.JSON(http.StatusOK, utils.RespHelper(utils.SuccessResp()))
}

func Login(c *gin.Context) {
	var userRegist userRegistVO
	err := c.ShouldBindJSON(&userRegist)
	if err != nil || userRegist.Password == "" || userRegist.Mail == "" {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.INVALID_FORM_PARAMETER))
		return
	}

	userSrv := service.NewUserService(config.GetMySQLClient())
	dbUser, err := userSrv.GetUserByMailAndPwd(userRegist.Mail, userRegist.Password)
	if err != nil || dbUser.Id == "" {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.USER_LOGIN_FAIL))
		return
	}

	sessionId := uuid.NewV4().String()
	_ = config.GetRedisClient().Set(dbUser.Id+"-session_id", sessionId, time.Hour*12)
	c.SetCookie("uid", dbUser.Id, 3600, "/", "", false, false)
	c.SetCookie("session_id", sessionId, 3600, "/", "", false, false)
	c.JSON(http.StatusOK, utils.RespHelper(utils.SuccessResp()))
}

func GetLoginStatus(c *gin.Context) {
	uid, err := c.Cookie("uid")
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.INVALID_COOKIE))
		return
	}
	sessionId, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.INVALID_COOKIE))
		return
	}

	dbSessionId, err := config.GetRedisClient().Get(fmt.Sprintf("%s-session_id", uid))
	if dbSessionId != sessionId || err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.USER_NOT_LOGIN))
		return
	}
	c.JSON(http.StatusOK, utils.RespHelper(utils.SetData("login", 1)))
}
