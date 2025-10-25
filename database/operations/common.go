package operations

import (
	"gorm.io/gorm"
)

// applyQuery applies the given query to the GORM DB instance.
func applyQuery(db *gorm.DB, query map[string]interface{}) *gorm.DB {
	if len(query) > 0 {
		return db.Where(query)
	}

	return db.Where("1 = 1")
}
