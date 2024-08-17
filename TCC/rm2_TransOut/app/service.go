package app

import (
	"database/sql"
	"github.com/dtm-labs/dtmcli"
)

func TccBTransOutTryService(barrier *dtmcli.BranchBarrier, req *ReqHTTP) error {

	return barrier.CallWithDB(pdbGet(), func(tx *sql.Tx) error {
		return tccAdjustTrading(tx, req.UserID, -req.Amount)
	})
}

func TccBTransOutConfirmService(barrier *dtmcli.BranchBarrier, req *ReqHTTP) error {

	return barrier.CallWithDB(pdbGet(), func(tx *sql.Tx) error {
		return tccAdjustBalance(tx, req.UserID, -req.Amount)
	})
}

func TccBarrierTransOutCancelService(barrier *dtmcli.BranchBarrier, req *ReqHTTP) error {
	return barrier.CallWithDB(pdbGet(), func(tx *sql.Tx) error {
		return tccAdjustTrading(tx, TransOutUID, req.Amount)
	})
}
