package orm

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 設定連線字串
func setConnection(driver, host, port, database, username, password string) string {
	switch driver {
	case "mysql":
		return username + ":" + password + "@tcp(" + host + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Asia%2FTaipei"
	}
	return ""
}

// NewOrmConnection 建立連線
func NewOrmConnection(connName string, conf *Host) (err error) {
	connectName := setConnection(
		"mysql",
		conf.Host,
		conf.Port,
		conf.DB,
		conf.Username,
		conf.Password,
	)
	db, err := gorm.Open("mysql", connectName)
	if err == nil {
		db.LogMode(conf.LogMode)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(conf.MaxConn)
	db.DB().SetConnMaxLifetime(time.Minute * 5)

	//Add to pool map
	DbPool.AddPool(connName, db)

	return
}
