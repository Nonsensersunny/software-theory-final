package service

import (
	"software-theory-final/pkg/modules/model"
	"software-theory-final/pkg/modules/mysql"
)

type DatasetService struct {
	client *mysql.Client
}


func NewDatasetService(client *mysql.Client) *DatasetService {
	return &DatasetService{
		client: client,
	}
}

func (d *DatasetService) CreateDataset(dataset *model.Dataset) error {
	return d.client.DB.Table("datasets").Create(&dataset).Error
}

func (d *DatasetService) GetDatasetById(id string) (*model.Dataset, error) {
	dataset := &model.Dataset{}
	err := d.client.DB.Table("datasets").Where("id = ?", id).Scan(dataset).Error
	return dataset, err
}

func (d *DatasetService) GetDatasetsByUid(uid string) ([]*model.Dataset, error) {
	var datasets []*model.Dataset
	err := d.client.DB.Table("datasets").Where("uid = ?", uid).Scan(&datasets).Error
	return datasets, err
}