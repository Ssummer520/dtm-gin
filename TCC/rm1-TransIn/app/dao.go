package app

import (
	"database/sql"
	"fmt"
	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmcli/dtmimp"
	"github.com/dtm-labs/dtmcli/logger"
)

var StoreHost = "localhost"

// BusiConf 1
var BusiConf = dtmcli.DBConf{
	Driver:   "mysql",
	Host:     StoreHost,
	Port:     3306,
	User:     "sa",
	Password: "sa123456",
}

func pdbGet() *sql.DB {
	db, err := dtmimp.PooledDB(BusiConf)
	logger.FatalIfError(err)
	return db
}

func tccAdjustTrading(db dtmcli.DB, uid int, amount int) error {
	affected, err := dtmimp.DBExec(BusiConf.Driver, db, `update dtm_busi.user_account
		set trading_balance=trading_balance+?
		where user_id=? and trading_balance + ? + balance >= 0`, amount, uid, amount)
	if err == nil && affected == 0 {
		return fmt.Errorf("update error, maybe balance not enough")
	}
	return err
}
func tccAdjustBalance(db dtmcli.DB, uid int, amount int) error {
	affected, err := dtmimp.DBExec(BusiConf.Driver, db, `update dtm_busi.user_account
		set trading_balance=trading_balance-?,
		balance=balance+? where user_id=?`, amount, amount, uid)
	if err == nil && affected == 0 {
		return fmt.Errorf("update user_account 0 rows")
	}
	return err
}
