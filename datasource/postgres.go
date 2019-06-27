package datasource

import (
	"fmt"
	"github.com/OahcUil94/go-boilerplate/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type pqdb struct {
	*gorm.DB
}

func (p *pqdb) initDB() {
	var getURI = func () string {
		return fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			config.Postgres.Host, config.Postgres.Port,
			config.Postgres.User, config.Postgres.Password,
			config.Postgres.DB, config.Postgres.SSLMode)
	}

	db, err := gorm.Open("postgres", getURI())
	if err != nil {
		msg := fmt.Sprint("init postgres db error, errmsg: ", err.Error())
		panic(msg)
	}

	if err := db.DB().Ping(); err != nil {
		msg := fmt.Sprintf(
			"ping postgres db error, host=%s, port=%s, errmsg: %s",
			config.Postgres.Host,
			config.Postgres.Port, err.Error())

		panic(msg)
	}

	db.LogMode(true)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(30)
	db.DB().SetConnMaxLifetime(time.Hour)

	p.DB = db
}

func (p *pqdb) AddCommentToColumn(v interface{}) (err error) {
	scope := p.NewScope(v)
	tableName := scope.TableName()
	fields := scope.GetStructFields()

	for _, field := range fields {
		if field == nil || field.TagSettings["PQ_COMMENT"] == "" {
			continue
		}

		if err = p.Exec(
			fmt.Sprintf(`comment on column %s."%s" is '%s'`, tableName,
				field.DBName, field.TagSettings["PQ_COMMENT"]),
		).Error; err != nil {
			break
		}
	}

	return
}

func (p *pqdb) BulkAddCommentToColumn(arr ...interface{}) (err error) {
	for _, v := range arr {
		if err = p.AddCommentToColumn(v); err != nil {
			break
		}
	}

	return
}

var PqDB = &pqdb{}
