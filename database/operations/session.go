package operations

import (
	"errors"

	"gorm.io/gorm"

	"github.com/sentinel-official/dvpn-node/database/models"
)

// SessionInsertOne inserts a single Session record into the database.
func SessionInsertOne(db *gorm.DB, session *models.Session) error {
	if err := db.Create(session).Error; err != nil {
		return err
	}

	return nil
}

// SessionInsertMany inserts multiple Session records into the database.
func SessionInsertMany(db *gorm.DB, sessions []models.Session) error {
	if err := db.Create(&sessions).Error; err != nil {
		return err
	}

	return nil
}

// SessionFindOne retrieves a single session record from the database based on the provided query.
func SessionFindOne(db *gorm.DB, query map[string]interface{}) (*models.Session, error) {
	var session models.Session

	db = applyQuery(db, query)
	if err := db.First(&session).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &session, nil
}

// SessionFind retrieves multiple session records from the database based on the provided query.
func SessionFind(db *gorm.DB, query map[string]interface{}) ([]models.Session, error) {
	var sessions []models.Session

	db = applyQuery(db, query)
	if err := db.Find(&sessions).Error; err != nil {
		return nil, err
	}

	return sessions, nil
}

// SessionFindOneAndUpdate finds a single session record based on the provided query and updates it with the provided updates.
func SessionFindOneAndUpdate(db *gorm.DB, query, updates map[string]interface{}) (*models.Session, error) {
	session, err := SessionFindOne(db, query)
	if err != nil || session == nil {
		return nil, err
	}

	if err := db.Model(session).Updates(updates).Error; err != nil {
		return nil, err
	}

	return session, nil
}

// SessionUpdateMany updates multiple session records based on the provided query and updates them with the provided updates.
func SessionUpdateMany(db *gorm.DB, query map[string]interface{}, updates map[string]interface{}) error {
	db = applyQuery(db, query)
	if err := db.Model(&models.Session{}).Updates(updates).Error; err != nil {
		return err
	}

	return nil
}

// SessionFindOneAndDelete finds a single session record based on the provided query and deletes it.
func SessionFindOneAndDelete(db *gorm.DB, query map[string]interface{}) (*models.Session, error) {
	session, err := SessionFindOne(db, query)
	if err != nil || session == nil {
		return nil, err
	}

	if err := db.Model(session).Delete(nil).Error; err != nil {
		return nil, err
	}

	return session, nil
}

// SessionDeleteMany deletes multiple session records based on the provided query.
func SessionDeleteMany(db *gorm.DB, query map[string]interface{}) error {
	db = applyQuery(db, query)
	if err := db.Model(&models.Session{}).Delete(nil).Error; err != nil {
		return err
	}

	return nil
}
