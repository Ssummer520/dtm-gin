package app

import (
	"database/sql"
	"github.com/dtm-labs/dtmcli"
)

func TccBTransInTryService(barrier *dtmcli.BranchBarrier, req *ReqHTTP) error {

	return barrier.CallWithDB(pdbGet(), func(tx *sql.Tx) error {
		return tccAdjustTrading(tx, req.UserID, req.Amount)
	})
}

func TccBTransInConfirmService(barrier *dtmcli.BranchBarrier, req *ReqHTTP) error {

	return barrier.CallWithDB(pdbGet(), func(tx *sql.Tx) error {
		return tccAdjustBalance(tx, req.UserID, req.Amount)
	})
}

func TccBarrierTransInCancelService(barrier *dtmcli.BranchBarrier, req *ReqHTTP) error {
	return barrier.CallWithDB(pdbGet(), func(tx *sql.Tx) error {
		return tccAdjustTrading(tx, req.UserID, -req.Amount)
	})
}
