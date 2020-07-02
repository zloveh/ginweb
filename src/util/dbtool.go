package util

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var MySQLPoolTimeOutError error

func init() {
	rand.Seed(time.Now().UnixNano())
	MySQLPoolTimeOutError = errors.New("mysql pool get conn timeout")
}

type DBPool struct {
	// pools   []*sql.DB
	// channel chan bool
	pools chan *sql.DB
}

func NewDBPoolWithConfig(conf string, dbName string) (*DBPool, error) {
	config, err := NewConfig(conf)
	if err != nil {
		return nil, errors.New("读取数据库配置文件错误")
	}
	host := config.MustValue(dbName, "host")
	port := config.MustInt(dbName, "port", 3306)
	username := config.MustValue(dbName, "username")
	password := config.MustValue(dbName, "password")
	database := config.MustValue(dbName, "database")
	maxIdle := config.MustInt(dbName, "maxIdle", 2)
	maxOpen := config.MustInt(dbName, "maxOpen", 2)
	poolSize := config.MustInt(dbName, "poolSize", 5)
	dbPool, err := NewDBPool(username, password, host, port, database, maxIdle, maxOpen, poolSize)
	return dbPool, err
}

/*
   使用默认的charset utf8
*/
func NewDBPool(username string, password string, host string, port int, database string, maxIdle int, maxOpen int, poolSize int) (*DBPool, error) {
	return NewDBPoolWithCharset(username, password, host, port, database, maxIdle, maxOpen, poolSize, "utf8")
}

/*
 */
func NewDBPoolWithCharset(username string, password string, host string, port int, database string, maxIdle int, maxOpen int, poolSize int, charset string) (*DBPool, error) {
	var (
		url            = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", username, password, host, port, database, charset)
		dbPool *DBPool = &DBPool{
			pools: make(chan *sql.DB, poolSize),
		}
		conn *sql.DB
		err  error
	)
	if poolSize <= 0 {
		return nil, errors.New("没有传递pool size")
	}

	loopSize := poolSize
	for {
		if loopSize <= 0 {
			break
		}
		loopSize--
		if conn, err = sql.Open("mysql", url); err != nil {
			fmt.Println("db 连接失败 %v", err)
			continue
		}
		if err = conn.Ping(); err != nil {
			fmt.Println("db 连接失败 %v", err)
			continue
		}
		conn.SetMaxIdleConns(maxIdle)
		conn.SetMaxOpenConns(maxOpen)
		dbPool.pools <- conn
	}
	if len(dbPool.pools) == 0 {
		return nil, fmt.Errorf("连接mysql池为空")
	}
	// dbPool.channel = make(chan bool, poolSize)
	return dbPool, nil
}

/*
获取db连接
*/
func (this *DBPool) GetConn() *sql.DB {
	deadline := time.After(1 * time.Second)
	select {
	case <-deadline:
		return nil
	case c := <-this.pools:
		return c
	}
	return nil
}

/*
释放链接
*/
func (this *DBPool) Release(c *sql.DB) {
	this.pools <- c
}
