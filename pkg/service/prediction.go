package service

import (
	"software-theory-final/pkg/modules/model"
	"software-theory-final/pkg/modules/mysql"
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
	return prediction, err
}

func (p *PredictionService) GetPredictionsByDid(did string) ([]*model.Prediction, error) {
	var predictions []*model.Prediction
	err := p.client.DB.Table("predictions").Where("did = ?", did).Scan(&predictions).Error
	return predictions, err
}

func (p *PredictionService) GetPredictionsByAid(aid string) ([]*model.Prediction, error) {
	var predictions []*model.Prediction
	err := p.client.DB.Table("predictions").Where("aid = ?", aid).Scan(&predictions).Error
	return predictions, err
}