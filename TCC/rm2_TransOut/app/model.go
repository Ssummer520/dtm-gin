package app

import (
	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmcli/logger"
	"github.com/gin-gonic/gin"
)

// TransOutUID 1
const TransOutUID = 1

// TransInUID 2
const TransInUID = 2

// Redis 1
const Redis = "redis"

// Mongo 1
const Mongo = "mongo"

// QsStartSvr quick start: start server
func QsStartSvr() {

}

type mainSwitchType struct {
	TransInResult         AutoEmptyString
	TransOutResult        AutoEmptyString
	TransInConfirmResult  AutoEmptyString
	TransOutConfirmResult AutoEmptyString
	TransInRevertResult   AutoEmptyString
	TransOutRevertResult  AutoEmptyString
	QueryPreparedResult   AutoEmptyString
	NextResult            AutoEmptyString
	JrpcResult            AutoEmptyString
	FailureReason         AutoEmptyString
}

// AutoEmptyString auto reset to empty when used once
type AutoEmptyString struct {
	value string
}

// SetOnce set a value once
func (s *AutoEmptyString) SetOnce(v string) {
	s.value = v
}

// Fetch fetch the stored value, then reset the value to empty
func (s *AutoEmptyString) Fetch() string {
	v := s.value
	s.value = ""
	if v != "" {
		logger.Debugf("fetch obtain not empty value: %s", v)
	}
	return v
}

// MainSwitch controls busi success or fail
var MainSwitch mainSwitchType

func InfoFromContext(c *gin.Context) *dtmcli.BranchBarrier {
	info := dtmcli.BranchBarrier{
		TransType: c.Query("trans_type"),
		Gid:       c.Query("gid"),
		BranchID:  c.Query("branch_id"),
		Op:        c.Query("op"),
	}
	return &info
}

type ReqHTTP struct {
	Amount int    `json:"amount"`
	UserID string `json:"userID"`
}
