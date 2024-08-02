package operations

import (
	"gorm.io/gorm"
)

// applyConditions applies conditions to a GORM query.
func applyConditions(query *gorm.DB, conditions map[string]interface{}) *gorm.DB {
	if len(conditions) > 0 {
		for key, value := range conditions {
			query = query.Where(key+" = ?", value)
		}
	} else {
		query = query.Where("1 = 1")
	}

	return query
}
