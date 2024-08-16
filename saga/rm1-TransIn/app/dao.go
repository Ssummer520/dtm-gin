package app

import (
	"database/sql"
	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmcli/dtmimp"
	"github.com/dtm-labs/dtmcli/logger"
	"strings"
)

func SagaAdjustBalance(db dtmcli.DB, uid int, amount int, result string) error {
	if strings.Contains(result, dtmcli.ResultFailure) {
		return dtmcli.ErrFailure
	}
	_, err := dtmimp.DBExec(BusiConf.Driver, db, "update test.user_account set balance = balance + ?  ,  update_time = NOW()  where user_id = ?", amount, uid)
	return err
}

var StoreHost = "localhost"

// BusiConf 1
var BusiConf = dtmcli.DBConf{
	Driver:   "mysql",
	Host:     StoreHost,
	Port:     3306,
	User:     "root",
	Password: "sa123456",
}

func pdbGet() *sql.DB {
	db, err := dtmimp.PooledDB(BusiConf)
	logger.FatalIfError(err)
	return db
}
