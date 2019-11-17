package v1

import "C"
import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"io"
	"net/http"
	"os"
	"software-theory-final/internal/config"
	"software-theory-final/internal/utils"
	"software-theory-final/pkg/modules/model"
	"software-theory-final/pkg/service"
)

func UploadDataset(c *gin.Context) {
	uid, err := c.Cookie("uid")
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.INVALID_COOKIE))
		return
	}

	desc := c.Request.PostFormValue("description")
	typ := c.Request.PostFormValue("type")
	name := c.Request.PostFormValue("name")

	file, _, err := c.Request.FormFile("upload")
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.UPLOAD_FILE_FAIL))
		return
	}

	filename := fmt.Sprintf("%s/%s", utils.GetDataSetsDir(uid), fmt.Sprintf("%s.csv", uuid.NewV1()))
	out, err := os.Create(filename)
	defer out.Close()
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.CREATE_FILE_FAIL))
		return
	}

	if _, err := io.Copy(out, file); err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.COPY_FILE_FAIL))
		return
	}

	service := service.NewDatasetService(config.GetMySQLClient())
	if err := service.CreateDataset(&model.Dataset{
		Uid:         uid,
		Path:        filename,
		Type:        typ,
		Description: desc,
		Name:        name,
	}); err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.CREATE_DATASET_FAIL))
		return
	}
	c.JSON(http.StatusOK, utils.RespHelper(utils.SetData("id", uid)))
}

func GetDatasetById(id string, c *gin.Context) {
	service := service.NewDatasetService(config.GetMySQLClient())
	dbDataset, err := service.GetDatasetById(id)
	if dbDataset.Id == "" || err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.DATASET_NOT_EXISTS))
		return
	}

	c.JSON(http.StatusOK, utils.RespHelper(utils.SetData("dataset", dbDataset)))
}

func GetDatasetsByUid(uid string, c *gin.Context) {

	service := service.NewDatasetService(config.GetMySQLClient())
	dataSets, err := service.GetDatasetsByUid(uid)
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.DATASET_NOT_EXISTS))
		return
	}
	c.JSON(http.StatusOK, utils.RespHelper(utils.SetData("datasets", dataSets)))
}

func GetDataSets(c *gin.Context) {
	id := c.Query("id")
	if id != "" {
		GetDatasetById(id, c)
		return
	}

	uid, err := c.Cookie("uid")
	if err != nil || uid == "" {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.INVALID_COOKIE))
		return
	}
	GetDatasetsByUid(uid, c)
}

func DelDatasetById(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusOK, utils.ErrorHelper(errors.New("id not valid"), utils.INVALID_PATH_PARAMETER))
		return
	}
	service := service.NewDatasetService(config.GetMySQLClient())
	err := service.DelDatasetById(id)
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.DELETE_DATASET_FAIL))
		return
	}
	c.JSON(http.StatusOK, utils.RespHelper(utils.SuccessResp()))
}
