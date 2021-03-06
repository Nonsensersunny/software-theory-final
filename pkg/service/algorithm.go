package service

import (
	"software-theory-final/pkg/modules/model"
	"software-theory-final/pkg/modules/mysql"
)

type AlgorithmService struct {
	client *mysql.Client
}

func NewAlgorithmService(client *mysql.Client) *AlgorithmService {
	return &AlgorithmService{
		client: client,
	}
}

func (a *AlgorithmService) CreateAlgorithm(algorithm *model.Algorithm) error {
	return a.client.DB.Table("algorithms").Create(&algorithm).Error
}

func (a *AlgorithmService) GetAlgorithmById(id string) (*model.Algorithm, error) {
	algorithm := &model.Algorithm{}
	err := a.client.DB.Table("algorithms").Where("id = ?", id).Scan(algorithm).Error
	return algorithm, err
}

func (a *AlgorithmService) GetAllAlgorithms() ([]*model.Algorithm, error) {
	var algos []*model.Algorithm
	err := a.client.DB.Table("algorithms").Find(&algos).Error
	return algos, err
}
