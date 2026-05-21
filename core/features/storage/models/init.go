package models

import "gorm.io/gorm"

func AutoMigrateStorageModel(db *gorm.DB) error {
	for _, initFunc := range []func(*gorm.DB) error{
		InitChunk,
		InitFile,
		InitStorage,
		InitUpload,
	} {
		if err := initFunc(db); err != nil {
			return err
		}
	}
	return nil
}
