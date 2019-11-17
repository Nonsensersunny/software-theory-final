package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/exec"
	"software-theory-final/internal/config"
	"software-theory-final/internal/utils"
	"software-theory-final/pkg/modules/model"
	"software-theory-final/pkg/service"
	"strconv"
	"strings"
)

type predictRequest struct {
	Test      string `json:"test"`
	Train     string `json:"train"`
	Algorithm string `json:"algorithm"`
}

func CreatePrediction(c *gin.Context) {
	uid, err := c.Cookie("uid")
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.INVALID_COOKIE))
		return
	}
	var predictR predictRequest
	if err := c.ShouldBindJSON(&predictR); err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.INVALID_FORM_PARAMETER))
		return
	}
	datasetService := service.NewDatasetService(config.GetMySQLClient())
	trainDataset, err1 := datasetService.GetDatasetById(predictR.Train)
	testDataset, err2 := datasetService.GetDatasetById(predictR.Test)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err1, utils.DATASET_NOT_EXISTS))
		return
	}
	algorithmService := service.NewAlgorithmService(config.GetMySQLClient())
	algo, err := algorithmService.GetAlgorithmById(predictR.Algorithm)
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err1, utils.ALGORITHM_NOT_EXISTS))
		return
	}
	cmd := exec.Command("python", algo.Path, trainDataset.Path, testDataset.Path, uid)
	buf, err := cmd.Output()
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.ALGORITHM_EXEC_FAIL))
		return
	}
	log.Println(buf)
	strBuf := strings.ReplaceAll(string(buf), "\n", "")
	output := strings.Split(strBuf, "@")
	if len(output) < 2 {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.ALGORITHM_EXEC_FAIL))
		return
	}
	accuracy, err := strconv.ParseFloat(output[1], 32)
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.ALGORITHM_EXEC_FAIL))
		return
	}
	predictionService := service.NewPredictionService(config.GetMySQLClient())
	pdc := &model.Prediction{
		Aid:      algo.Id,
		TrainId:  trainDataset.Id,
		TestId:   testDataset.Id,
		Accuracy: accuracy,
		Path:     output[0],
	}
	err = predictionService.CreatePrediction(pdc)
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.CREATE_PREDICTION_FAIL))
		return
	}
	c.JSON(http.StatusOK, utils.RespHelper(utils.SetData("result", pdc.Id)))
}

func GetPredictions(c *gin.Context) {
	predictionService := service.NewPredictionService(config.GetMySQLClient())
	id := c.Query("id")
	if id != "" {
		prediction, err := predictionService.GetPredictionById(id)
		if err != nil {
			c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.PREDICTION_NOT_EXISTS))
			return
		}
		c.JSON(http.StatusOK, utils.RespHelper(utils.SetData("prediction", prediction)))
		return
	}
	trainId := c.Query("train_id")
	if trainId != "" {
		predictions, err := predictionService.GetPredictionsByTrainId(trainId)
		if err != nil {
			c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.PREDICTION_NOT_EXISTS))
			return
		}
		c.JSON(http.StatusOK, utils.RespHelper(utils.SetData("prediction", predictions)))
		return
	}
	testId := c.Query("test_id")
	if testId != "" {
		predictions, err := predictionService.GetPredictionsByTestId(testId)
		if err != nil {
			c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.PREDICTION_NOT_EXISTS))
			return
		}
		c.JSON(http.StatusOK, utils.RespHelper(utils.SetData("prediction", predictions)))
		return
	}
	algorithmId := c.Query("algorithm_id")
	if algorithmId != "" {
		predictions, err := predictionService.GetPredictionsByAid(algorithmId)
		if err != nil {
			c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.PREDICTION_NOT_EXISTS))
			return
		}
		c.JSON(http.StatusOK, utils.RespHelper(utils.SetData("prediction", predictions)))
		return
	}
	c.JSON(http.StatusOK, utils.RespHelper(utils.SuccessResp()))
}
