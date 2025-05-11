package storage

import (
	"errors"

	"github.com/CaptainFallaway/XDH/internal/models"
)

type testStorage struct {
	Surveys   []Survey
	Groupings []Grouping
	Scans     []Scan
}

func NewTestStorage() Storage {
	return &testStorage{
		Surveys:   make([]Survey, 0),
		Groupings: make([]Grouping, 0),
		Scans:     make([]Scan, 0),
	}
}

func (ts *testStorage) GetAllSurveys() ([]*models.Survey, error) {
	surveys := make([]*models.Survey, len(ts.Surveys))

	for i, survey := range ts.Surveys {
		temp := new(models.Survey)
		convertFields(survey, temp)
		surveys[i] = temp
	}

	return surveys, nil
}

func (ts *testStorage) InsertSurvey(survey *models.Survey) error {
	temp := new(Survey)
	convertFields(survey, temp)
	ts.Surveys = append(ts.Surveys, *temp)
	return nil
}

func (ts *testStorage) UpdateSurvey(survey *models.Survey) error {
	temp := new(Survey)
	convertFields(survey, temp)

	for i, s := range ts.Surveys {
		if s.Uid == survey.Uid {
			ts.Surveys[i] = *temp
			return nil
		}
	}

	return errors.New("Survey not found")
}

func valInSlice(val string, slice []string) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func (ts *testStorage) DeleteSurvey(uid string) error {
	for i, s := range ts.Surveys {
		if s.Uid == uid {
			ts.Surveys = append(ts.Surveys[:i], ts.Surveys[i+1:]...)
		}
	}

	var associatedGroupingUids []string

	for i, g := range ts.Groupings {
		if g.SurveyID == uid {
			ts.Groupings = append(ts.Groupings[:i], ts.Groupings[i+1:]...)
			associatedGroupingUids = append(associatedGroupingUids, g.Uid)
		}
	}

	for i, s := range ts.Scans {
		if valInSlice(s.GroupingID, associatedGroupingUids) {
			ts.Scans = append(ts.Scans[:i], ts.Scans[i+1:]...)
		}
	}

	return nil
}

func (ts *testStorage) InsertGroupings(surveryId string, groupings []*models.Grouping) error {
	for _, grouping := range groupings {
		temp := new(Grouping)
		convertFields(grouping, temp)

		temp.SurveyID = surveryId

		ts.Groupings = append(ts.Groupings, *temp)
	}

	return nil
}

func (ts *testStorage) GetGroupings(surveyId string) ([]*models.Grouping, error) {
	groupings := make([]*models.Grouping, 0)

	for _, grouping := range ts.Groupings {
		if grouping.SurveyID == surveyId {
			temp := new(models.Grouping)
			convertFields(grouping, temp)
			groupings = append(groupings, temp)
		}
	}

	return groupings, nil
}

func (ts *testStorage) InsertScans(groupingId string, scans []*models.Scan) error {
	for _, scan := range scans {
		temp := new(Scan)
		convertFields(scan, temp)

		temp.GroupingID = groupingId

		ts.Scans = append(ts.Scans, *temp)
	}

	return nil
}

func (ts *testStorage) UpdateScan(scan *models.Scan) error {
	for i, s := range ts.Scans {
		if s.Uid == scan.Uid {
			temp := new(Scan)
			convertFields(scan, temp)
			ts.Scans[i] = *temp
			return nil
		}
	}

	return errors.New("Scan not found")
}
