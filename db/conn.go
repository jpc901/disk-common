package db

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jpc901/disk-common/conf"
	log "github.com/sirupsen/logrus"
)

type MySQLConn struct {
	conn *sql.DB
}

var (
	once sync.Once
	dbConn *MySQLConn
)

func GetDBInstance() *MySQLConn {
	once.Do(func() {
		dbConn = &MySQLConn{}
	})
	return dbConn
}

func (m *MySQLConn)GetConn() *sql.DB {
	return m.conn
}

func (m *MySQLConn) Init (config conf.MySQLConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DB,
	)
	var err error
	m.conn, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	m.conn.SetMaxIdleConns(config.MaxIdleConns)
	m.conn.SetMaxOpenConns(config.MaxOpenConns)

	err = m.conn.Ping()
	if err != nil {
		log.Fatal(err)
	}
}