package service

import (
	"io/ioutil"
	"software-theory-final/pkg/modules/model"
	"software-theory-final/pkg/modules/mysql"
	"os"
)

type PredictionService struct {
	client *mysql.Client
}

func NewPredictionService(client *mysql.Client) *PredictionService {
	return &PredictionService{
		client: client,
	}
}

func (p *PredictionService) CreatePrediction(prediction *model.Prediction) error {
	return p.client.DB.Table("predictions").Create(&prediction).Error
}

func (p *PredictionService) GetPredictionById(id string) (*model.Prediction, error) {
	prediction := &model.Prediction{}
	err := p.client.DB.Table("predictions").Where("id = ?", id).Scan(prediction).Error
	f, err := os.Open(prediction.Path)
	if err != nil {
		return nil, err
	}
	btCnt, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	prediction.Result = string(btCnt)
	return prediction, err
}

func (p *PredictionService) GetPredictionsByTrainId(tid string) ([]*model.Prediction, error) {
	var predictions []*model.Prediction
	err := p.client.DB.Table("predictions").Where("train_id = ?", tid).Scan(&predictions).Error
	return predictions, err
}

func (p *PredictionService) GetPredictionsByTestId(tid string) ([]*model.Prediction, error) {
	var predictions []*model.Prediction
	err := p.client.DB.Table("predictions").Where("test_id = ?", tid).Scan(&predictions).Error
	return predictions, err
}

func (p *PredictionService) GetPredictionsByAid(aid string) ([]*model.Prediction, error) {
	var predictions []*model.Prediction
	err := p.client.DB.Table("predictions").Where("aid = ?", aid).Scan(&predictions).Error
	return predictions, err
}