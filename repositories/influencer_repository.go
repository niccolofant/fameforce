package repositories

import (
	"database/sql"

	"github.com/ocintnaf/fameforce/entities"
)

const (
	findAllSqlStatement = "SELECT * FROM influencers"
)

type influencerRepository struct {
	db *sql.DB
}

type InfluencerRepository interface {
	FindAll() ([]entities.InfluencerEntity, error)
}

func NewInfluencerRepository(db *sql.DB) *influencerRepository {
	return &influencerRepository{db: db}
}

func (r *influencerRepository) FindAll() ([]entities.InfluencerEntity, error) {
	var influencers []entities.InfluencerEntity

	rows, err := r.db.Query(findAllSqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var influencer entities.InfluencerEntity

		err := rows.Scan(&influencer.ID, &influencer.Name)
		if err != nil {
			return nil, err
		}

		influencers = append(influencers, influencer)
	}

	return influencers, nil
}
