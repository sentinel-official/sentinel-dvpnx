package operations

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/sentinel-official/dvpn-node/database/models"
)

// SessionInsertOne inserts a single Session record into the database.
func SessionInsertOne(db *gorm.DB, session *models.Session) error {
	if result := db.Create(session); result.Error != nil {
		return result.Error
	}

	return nil
}

// SessionInsertMany inserts multiple Session records into the database.
func SessionInsertMany(db *gorm.DB, sessions []models.Session) error {
	if result := db.Create(&sessions); result.Error != nil {
		return result.Error
	}

	return nil
}

// SessionFindOne finds a single Session record matching the provided conditions.
func SessionFindOne(db *gorm.DB, conditions map[string]interface{}) (*models.Session, error) {
	var session models.Session

	query := applyConditions(db.Model(&models.Session{}), conditions)
	if result := query.First(&session); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &session, nil
}

// SessionFind finds multiple Session records matching the provided conditions.
func SessionFind(db *gorm.DB, conditions map[string]interface{}) ([]models.Session, error) {
	var sessions []models.Session

	query := applyConditions(db.Model(&models.Session{}), conditions)
	if result := query.Find(&sessions); result.Error != nil {
		return nil, result.Error
	}

	return sessions, nil
}

// SessionFindOneAndUpdate finds a single Session record, updates it with new values, and returns the updated record.
func SessionFindOneAndUpdate(db *gorm.DB, conditions map[string]interface{}, updates map[string]interface{}) (*models.Session, error) {
	var session models.Session

	query := applyConditions(db.Model(&models.Session{}), conditions)
	if result := query.Clauses(clause.Returning{}).Updates(updates).First(&session); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &session, nil
}

// SessionUpdateMany updates multiple Session records with new values based on the provided conditions.
func SessionUpdateMany(db *gorm.DB, conditions map[string]interface{}, updates map[string]interface{}) error {
	query := applyConditions(db.Model(&models.Session{}), conditions)
	if result := query.Updates(updates); result.Error != nil {
		return result.Error
	}

	return nil
}

// SessionFindOneAndDelete finds a single Session record, deletes it, and returns the deleted record.
func SessionFindOneAndDelete(db *gorm.DB, conditions map[string]interface{}) (*models.Session, error) {
	var session models.Session

	query := applyConditions(db.Model(&models.Session{}), conditions)
	if result := query.Clauses(clause.Returning{}).Delete(&session); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &session, nil
}

// SessionDeleteMany deletes multiple Session records based on the provided conditions.
func SessionDeleteMany(db *gorm.DB, conditions map[string]interface{}) error {
	query := applyConditions(db.Model(&models.Session{}), conditions)
	if result := query.Delete(&models.Session{}); result.Error != nil {
		return result.Error
	}

	return nil
}
