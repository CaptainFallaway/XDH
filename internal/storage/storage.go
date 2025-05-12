package storage

import (
	"io"

	"github.com/CaptainFallaway/XDH/internal/api"
	"github.com/CaptainFallaway/XDH/internal/models"
)

type Storage interface {
	GetAllSurveys() ([]models.Survey, error)
	InsertSurvey(survey *models.Survey) error
	// UpdateSurvey updates the survey in the storage
	// But it also makes sure that the relation with the groupings are maintained.
	UpdateSurvey(surveyId string, survey *api.Survey) error
	// DeleteSurvey deletes the survey and the related groupings.
	DeleteSurvey(surveyId string) error

	GetGroupings(groupingIds []string) ([]models.Grouping, error)
	InsertGrouping(grouping *models.Grouping) error
	UpdateGrouping(groupingId string, grouping *api.Grouping) error

	io.Closer
}
