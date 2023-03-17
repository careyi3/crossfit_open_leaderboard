package importer

import (
	"errors"

	"github.com/careyi3/crossfit_open_leaderboard/models"
	"gorm.io/gorm"
)

func Import(result *models.Response, db *gorm.DB) {
	for i := 0; i < len(result.LeaderboardRows); i++ {
		entrant := result.LeaderboardRows[i].Entrant
		scores := result.LeaderboardRows[i].Scores
		athlete := &models.Athlete{}
		res := db.Where("competitor_id = ?", entrant.CompetitorId).First(&athlete)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			db.Create(
				&models.Athlete{
					CompetitorId: int(entrant.CompetitorId),
					Name:         entrant.CompetitorName,
					Age:          int(entrant.Age),
					Gender:       entrant.Gender,
				},
			)
			db.Where("competitor_id = ?", entrant.CompetitorId).First(&athlete)
		}

		for j := 0; j < len(scores); j++ {
			score := scores[j]
			result := &models.Result{}
			res := db.Where("athlete_id = ? AND ordinal = ?", athlete.Id, score.Ordinal).First(&result)
			if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
				db.Create(
					&models.Result{
						AthleteId:   athlete.Id,
						Ordinal:     int(score.Ordinal),
						Score:       int(score.Score),
						Description: score.ScoreDisplay,
					},
				)
			}
		}
	}
}
