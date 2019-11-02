package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"software-theory-final/backend/internal/config"
	"software-theory-final/backend/internal/utils"
	"software-theory-final/backend/pkg/modules/model"
	"software-theory-final/backend/pkg/service"
)

func UploadDataset(c *gin.Context) {
	uid, err := c.Cookie("uid")
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.INVALID_COOKIE))
		return
	}

	desc := c.Request.PostFormValue("description")

	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.UPLOAD_FILE_FAIL))
		return
	}

	filename := fmt.Sprintf("%s/%s/%s", utils.DATA_DIR, uid, header.Filename)
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
		Description: desc,
	}); err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.CREATE_DATASET_FAIL))
		return
	}
	c.JSON(http.StatusOK, utils.RespHelper(utils.SuccessResp()))
}
