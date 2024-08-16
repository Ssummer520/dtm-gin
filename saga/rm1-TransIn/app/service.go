package app

import (
	"database/sql"
	"fmt"
	"github.com/dtm-labs/dtmcli"
)

func SagaAdjustBalanceService(barrier *dtmcli.BranchBarrier, req *ReqHTTP) error {
	if barrier == nil {
		return fmt.Errorf(dtmcli.ResultFailure)
	}

	return barrier.CallWithDB(pdbGet(), func(tx *sql.Tx) error {
		return SagaAdjustBalance(tx, 1000, req.Amount, req.TransInResult)
	})
}

func SagaAdjustBalanceCompensateService(barrier *dtmcli.BranchBarrier, req *ReqHTTP) error {
	if barrier == nil {
		return fmt.Errorf(dtmcli.ResultFailure)
	}
	return barrier.CallWithDB(pdbGet(), func(tx *sql.Tx) error {
		return SagaAdjustBalance(tx, 1000, -req.Amount, req.TransInResult)

	})
}
