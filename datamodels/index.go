package datamodels

import "github.com/OahcUil94/go-boilerplate/datasource"

func init() {
	datasource.PqDB.AutoMigrate(&Hero{})
	_ = datasource.PqDB.BulkAddCommentToColumn(&Hero{})
}
