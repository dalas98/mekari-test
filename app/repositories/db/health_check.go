package db

import "github.com/dalas98/mekari-test/app/models"

func (r *DBRepository) HealthCheck() (*models.HealthCheck, error) {
	tx := r.db.Begin()
	hc := models.HealthCheck{}

	if err := tx.Raw("select 'OK' as message, now() as time").Scan(&hc).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &hc, nil
}
