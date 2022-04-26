package helper

import (
	"gorm.io/gorm"
)

func CommitOrRollback(tx *gorm.DB) {
	err := recover()
	if err != nil {
		tx.Commit()
		panic(err)
	} else {
		tx.Commit()
	}
}
