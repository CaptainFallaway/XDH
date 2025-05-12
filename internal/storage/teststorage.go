package storage

import (
	"errors"

	"github.com/CaptainFallaway/XDH/internal/api"
	"github.com/CaptainFallaway/XDH/internal/models"
)

type testStorage struct {
	surveys   map[string]models.Survey
	groupings map[string]models.Grouping
}

func NewTestStorage() Storage {
	return &testStorage{
		surveys:   make(map[string]models.Survey),
		groupings: make(map[string]models.Grouping),
	}
}

func (s *testStorage) Close() error {
	return nil
}

func (s *testStorage) GetAllSurveys() ([]models.Survey, error) {
	surveys := make([]models.Survey, 0)
	for _, survey := range s.surveys {
		surveys = append(surveys, survey)
	}
	return surveys, nil
}

func (s *testStorage) InsertSurvey(survey *models.Survey) error {
	if _, exists := s.surveys[survey.Uid]; exists {
		return errors.New("survey already exists")
	}
	s.surveys[survey.Uid] = *survey
	return nil
}

func (s *testStorage) UpdateSurvey(surveyId string, survey *api.Survey) error {
	storedSurvey, exists := s.surveys[surveyId]
	if !exists {
		return errors.New("survey not found")
	}

	// Update the stored survey with the new values
	convertFields(survey, &storedSurvey)

	s.surveys[surveyId] = storedSurvey
	return nil
}

func (s *testStorage) DeleteSurvey(surveyId string) error {
	if _, exists := s.surveys[surveyId]; !exists {
		return errors.New("survey not found")
	}
	delete(s.surveys, surveyId)
	return nil
}

func (s *testStorage) GetGroupings(groupingIds []string) ([]models.Grouping, error) {
	groupings := make([]models.Grouping, 0)
	for _, id := range groupingIds {
		if grouping, exists := s.groupings[id]; exists {
			groupings = append(groupings, grouping)
		}
	}
	return groupings, nil
}

func (s *testStorage) InsertGrouping(grouping *models.Grouping) error {
	if _, exists := s.groupings[grouping.Uid]; exists {
		return errors.New("grouping already exists")
	}
	s.groupings[grouping.Uid] = *grouping
	return nil
}

func (s *testStorage) UpdateGrouping(groupingId string, grouping *api.Grouping) error {
	storedGrouping, exists := s.groupings[groupingId]
	if !exists {
		return errors.New("grouping not found")
	}

	// Update the stored grouping with the new values
	convertFields(grouping, &storedGrouping)

	s.groupings[groupingId] = storedGrouping
	return nil
}
