package storage

import (
	"fmt"

	"github.com/CaptainFallaway/XDH/internal/api"
	"github.com/CaptainFallaway/XDH/internal/models"
	"github.com/dgraph-io/badger/v4"
)

// TODO: Clean up the code by using the DRY principle
// and remove the duplicate code in the InsertSurvey and InsertGrouping methods
// and the UpdateSurvey and UpdateGrouping methods

type badgerStorage struct {
	surveys   *badger.DB
	groupings *badger.DB
}

func NewBadgerStorage(surveysPath, groupingsPath string) (Storage, error) {
	surveys, err := badger.Open(badger.DefaultOptions(surveysPath))
	if err != nil {
		return nil, err
	}

	groupings, err := badger.Open(badger.DefaultOptions(groupingsPath))
	if err != nil {
		return nil, err
	}

	return &badgerStorage{
		surveys:   surveys,
		groupings: groupings,
	}, nil
}

func (s *badgerStorage) Close() error {
	if err := s.surveys.Close(); err != nil {
		return err
	}
	if err := s.groupings.Close(); err != nil {
		return err
	}
	return nil
}

func (s *badgerStorage) GetAllSurveys() ([]models.Survey, error) {
	surveys := make([]models.Survey, 0)

	err := s.surveys.View(func(txn *badger.Txn) error {
		iter := txn.NewIterator(badger.DefaultIteratorOptions)
		defer iter.Close()

		var item *badger.Item

		for iter.Rewind(); iter.Valid(); iter.Next() {
			item = iter.Item()

			err := item.Value(func(val []byte) error {
				survey, err := decodeObj[models.Survey](val)
				if err != nil {
					return err
				}

				surveys = append(surveys, *survey)
				fmt.Println("A Survey:", survey)

				return nil
			})

			return err
		}

		return nil
	})

	return surveys, err
}

func (s *badgerStorage) InsertSurvey(survey *models.Survey) error {
	return s.surveys.Update(func(txn *badger.Txn) error {
		data, err := encodeObj(survey)
		if err != nil {
			return err
		}

		err = txn.Set([]byte(survey.Uid), data)
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *badgerStorage) UpdateSurvey(surveyId string, survey *api.Survey) error {
	return s.surveys.Update(func(txn *badger.Txn) error {
		// Retrieve the existing survey to make sure some values are not changed
		storedSurveyItem, err := txn.Get([]byte(surveyId))
		if err != nil {
			return err
		}

		var storedSurvey *models.Survey

		err = storedSurveyItem.Value(func(val []byte) error {
			var err error
			storedSurvey, err = decodeObj[models.Survey](val)
			return err
		})

		if err != nil {
			return err
		}

		// Update the survey with the new values
		// but keep the existing values for the fields that are not in the new survey
		convertFields(survey, storedSurvey)

		// Encode the survey to bytes
		// and store it in the database
		data, err := encodeObj(storedSurvey)
		if err != nil {
			return err
		}

		err = txn.Set([]byte(surveyId), data)
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *badgerStorage) deleteSurvey(uid string) error {
	return s.surveys.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(uid))
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *badgerStorage) deleteGroupings(groupingIds []string) error {
	return s.groupings.Update(func(txn *badger.Txn) error {
		for _, uid := range groupingIds {
			err := txn.Delete([]byte(uid))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// DeleteSurvey deletes the survey and the related groupings
//
// It first retrieves the groupings associated with the survey and then deletes them.
// Finally, it deletes the survey itself.
func (s *badgerStorage) DeleteSurvey(surveyId string) error {
	var groupingIds []string

	// First, retrieve the groupings associated with the survey
	err := s.surveys.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(surveyId))
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			survey, err := decodeObj[models.Survey](val)
			if err != nil {
				return err
			}

			groupingIds = survey.GroupingIds
			return nil
		})

		if err != nil {
			return err
		}

		return nil
	})

	// Delete the groupings associated with the survey
	err = s.deleteGroupings(groupingIds)
	if err != nil {
		return err
	}

	// Finally, delete the survey itself
	return s.deleteSurvey(surveyId)
}

func (s *badgerStorage) GetGroupings(groupingIds []string) ([]models.Grouping, error) {
	var groupings []models.Grouping

	err := s.groupings.View(func(txn *badger.Txn) error {
		for _, uid := range groupingIds {
			item, err := txn.Get([]byte(uid))
			if err != nil {
				return err
			}

			err = item.Value(func(val []byte) error {
				grouping, err := decodeObj[models.Grouping](val)
				if err != nil {
					return err
				}

				groupings = append(groupings, *grouping)
				return nil
			})

			if err != nil {
				return err
			}
		}

		return nil
	})

	return groupings, err
}

func (s *badgerStorage) InsertGrouping(grouping *models.Grouping) error {
	return s.groupings.Update(func(txn *badger.Txn) error {
		data, err := encodeObj(grouping)
		if err != nil {
			return err
		}

		err = txn.Set([]byte(grouping.Uid), data)
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *badgerStorage) UpdateGrouping(groupingId string, grouping *api.Grouping) error {
	return s.groupings.Update(func(txn *badger.Txn) error {
		// Retrieve the existing grouping to make sure some values are not changed
		storedSurveyItem, err := txn.Get([]byte(groupingId))
		if err != nil {
			return err
		}

		var storedGrouping *models.Grouping

		err = storedSurveyItem.Value(func(val []byte) error {
			var err error
			storedGrouping, err = decodeObj[models.Grouping](val)
			return err
		})

		if err != nil {
			return err
		}

		// Update the survey with the new values
		// but keep the existing values for the fields that are not in the new survey
		convertFields(grouping, storedGrouping)

		// Encode the survey to bytes
		// and store it in the database
		data, err := encodeObj(storedGrouping)
		if err != nil {
			return err
		}

		err = txn.Set([]byte(groupingId), data)
		if err != nil {
			return err
		}

		return nil
	})
}
