package service

import (
	"roar-dashboard-backend/models"
	"roar-dashboard-backend/repository"
)

type Service struct {
	repo repository.Repo
}

func (ser *Service) GetFeatureRequestsPage(feature_id int) (models.FeatureRequest, error) {
	feature_details, err := ser.repo.GetFeatureDetails(feature_id)
	if err != nil {
		return models.FeatureRequest{}, err
	}
	return feature_details, nil
}

func (ser *Service) GetAllFeatureRequestsPage() ([]models.FeatureRequest, error) {
	all_feature_details, err := ser.repo.GetAllFeatureDetails()
	if err != nil {
		return nil, err
	}
	return all_feature_details, nil
}

func (ser *Service) CreateFeature(createFeature models.FeatureRequest) error {
	err := ser.repo.CreateFeature(createFeature)
	if err != nil {
		return err
	}
	return nil
}

func NewService() Service {
	return Service{
		repo: repository.NewRepo(),
	}
}
