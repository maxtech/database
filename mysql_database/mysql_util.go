package mysql_database

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
    "log"
)

type mySQLPool struct {
    linkString string
    loggerName string
    logSqlFlag bool
}

var MySQLPool *mySQLPool

func init() {
    MySQLPool = new(mySQLPool)
}

func (mp *mySQLPool) Init(address, username, password, dbname string, logSql bool) {
    mp.linkString = fmt.Sprintf(
        "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&timeout=2s",
        username,
        password,
        address,
        dbname)
    mp.loggerName = dbname
    mp.logSqlFlag = logSql
}

func (mp *mySQLPool) GetEngine() *xorm.Engine {
    engine, err := xorm.NewEngine("mysql", mp.linkString)
    if err != nil {
        log.Fatal("Database connect error: ", err)
    }
    engine.ShowExecTime(mp.logSqlFlag)
    engine.ShowSQL(mp.logSqlFlag)
    engine.SetLogger(
        newLogger(fmt.Sprintf("[%s]", mp.loggerName), xorm.DEFAULT_LOG_FLAG))
    return engine
}

func (mp *mySQLPool) CheckMySQLEngine(engine *xorm.Engine) bool {
    err := engine.Ping()
    if err != nil {
        return false
    }
    return true
}
