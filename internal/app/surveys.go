package app

import (
	"github.com/CaptainFallaway/XDH/internal/api"
	"github.com/CaptainFallaway/XDH/internal/grouping"
	"github.com/CaptainFallaway/XDH/internal/models"
	"github.com/CaptainFallaway/XDH/internal/parsers"
	"github.com/google/uuid"
)

// CreateSurvey generates a new UID for the survey and inserts it into the storage
func (app *App) CreateSurvey(surveyDto *api.Survey, dataSourcePath string) *models.Survey {
	// Parse the data source
	scans, err := parsers.Parse(dataSourcePath)
	if err != nil {
		app.HandleError(err)
		return nil
	}

	// Create groupings from the parsed scans
	groupings, err := grouping.MakeBoatGroupings(scans)
	if err != nil {
		app.HandleError(err)
		return nil
	}

	// Generate a new UID for the survey
	uid, err := uuid.NewV7()
	if err != nil {
		app.HandleError(err)
		return nil
	}

	// Create the relationship between the survey and groupings
	groupingIds := make([]string, len(groupings))

	for i, grouping := range groupings {
		groupingIds[i] = grouping.Uid
	}

	survey := &models.Survey{
		Uid:              uid.String(),
		Surveyor:         surveyDto.Surveyor,
		Date:             surveyDto.Date,
		Location:         surveyDto.Location,
		InstrumentSerial: surveyDto.InstrumentSerial,
		GroupingIds:      groupingIds,
	}

	// Insert the groupings in the storage
	err = app.insertGroupings(groupings)
	if err != nil {
		app.HandleError(err)
		return nil
	}

	// Insert the survey in the storage
	err = app.Storage.InsertSurvey(survey)
	if err != nil {
		app.HandleError(err)
	}

	return survey
}

func (app *App) UpdateSurvey(surveyId string, survey *api.Survey) {
	err := app.Storage.UpdateSurvey(surveyId, survey)
	if err != nil {
		app.HandleError(err)
		return
	}
}

// GetSurveys returns a list of all surveys from the storage
func (app *App) GetSurveys() []models.Survey {
	surveys, err := app.Storage.GetAllSurveys()
	if err != nil {
		app.HandleError(err)
		return nil
	}

	return surveys
}

// DeleteSurvey deletes the survey and the related groupings
func (app *App) DeleteSurvey(surveyId string) {
	err := app.Storage.DeleteSurvey(surveyId)
	if err != nil {
		app.HandleError(err)
		return
	}
}
