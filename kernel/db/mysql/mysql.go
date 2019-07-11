package mysql

import (
	"webserver/kernel/common"

	"github.com/WeiWeiWesley/simple_orm"
	_ "github.com/go-sql-driver/mysql"
)

//NewPool Add a new pool
func NewPool(hostName string, conf common.MySQL) error {
	return orm.NewOrmConnection(hostName, common.MySQLDriver, &orm.Host{
		DB:       conf.DB,
		Host:     conf.Host,
		Port:     conf.Port,
		Username: conf.Username,
		Password: conf.Password,
		MaxConn:  conf.MaxConn,
		LogMode:  conf.LogMode,
	})
}

//GetPool Fetch a conn pool by name
func GetPool(poolName string) (*orm.DB, error) {
	return orm.GetDbConn(poolName)
}
