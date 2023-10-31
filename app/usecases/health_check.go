package usecases

import (
	"github.com/dalas98/mekari-test/app/models"
)

func (a *APIGOUsecase) HealthCheck() (*models.HealthCheck, error) {

	hc, err := a.DBRepository.HealthCheck()
	if err != nil {
		return nil, err
	}

	return hc, nil
}
