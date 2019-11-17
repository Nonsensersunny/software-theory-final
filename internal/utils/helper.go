package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

func ErrorHelper(err error, statusCode int) gin.H {
	return gin.H{
		"error": err,
		"code":  statusCode,
		"data":  "",
	}
}

type M map[string]interface{}

func SetData(key string, val interface{}) M {
	return M{
		key: val,
	}
}

func SuccessResp() M {
	return SetData("status", "success")
}

func RespHelper(m M) gin.H {
	return gin.H{
		"code":  SUCCESS,
		"data":  m,
		"error": "",
	}
}


func GetDataSetsDir(id string) string {
	return fmt.Sprintf("data/%s/dataset", id)
}

func GetPredictionsDir(id string) string {
	return fmt.Sprintf("data/%s/prediction", id)
}

func GetFileContent(filepath string) (string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	btCnt, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(btCnt), nil
}