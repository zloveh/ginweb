package conf

import (
	"ginweb/src/util"
)

var (
	Riskdb *DBInstance
)

type DBInstance struct {
	DBPool *util.DBPool
}

func NewDBInstance(c DBconfig) *DBInstance {
	DBIns := &DBInstance{}
	pool, err := util.NewDBPoolWithCharset(
		c.Usr,
		c.Pwd,
		c.Host,
		c.Port,
		c.DBname,
		c.MaxIdle,
		c.MaxOpen,
		c.PoolSize,
		"utf8mb4")

	if err != nil {
		panic("dbpool failed " + err.Error())
	}

	DBIns.DBPool = pool

	return DBIns
}

func InitDB(gconf *Config) {
	app_type := gconf.APPName

	if Riskdb == nil {
		util.Infof("%s init db: riskdb", app_type)
		Riskdb = NewDBInstance(gconf.RiskDBConfig)
	}
}
