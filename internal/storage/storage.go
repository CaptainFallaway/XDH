package storage

import "github.com/CaptainFallaway/XDH/internal/models"

type Storage interface {
	GetAllSurveys() ([]*models.Survey, error)
	InsertSurvey(survey *models.Survey) error
	UpdateSurvey(survey *models.Survey) error
	DeleteSurvey(uid string) error

	InsertGroupings(surveryId string, groupings []*models.Grouping) error
	GetGroupings(surveyId string) ([]*models.Grouping, error)

	InsertScans(groupingId string, scans []*models.Scan) error
	UpdateScan(scan *models.Scan) error
}
