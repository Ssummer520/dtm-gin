package main

import (
	"fmt"
	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmcli/dtmimp"
	"github.com/dtm-labs/dtmcli/logger"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	app := gin.Default()
	app.POST(qsBusiAPI+"/TransOut", func(c *gin.Context) {
		info := infoFromContext(c)
		var req ReqHTTP
		c.ShouldBindJSON(&req)
		log.Printf("TransOut:%v,gid:%v", req.Amount, info.Gid)
		c.JSON(http.StatusOK, dtmimp.OrString(MainSwitch.QueryPreparedResult.Fetch(), dtmcli.ResultSuccess))
	})
	app.POST(qsBusiAPI+"/TransOutCompensate", func(c *gin.Context) {
		info := infoFromContext(c)
		var req ReqHTTP
		c.ShouldBindJSON(&req)
		log.Printf("TransOutCompensate:%vgid:%v", req.Amount, info.Gid)
		c.JSON(http.StatusOK, dtmimp.OrString(MainSwitch.QueryPreparedResult.Fetch(), dtmcli.ResultSuccess))
	})
	log.Printf("quick start examples listening at %d", qsBusiPort)

	app.Run(fmt.Sprintf(":%d", qsBusiPort))
}

// busi address
const qsBusiAPI = "/api/busi_start"
const qsBusiPort = 8880

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

type ReqHTTP struct {
	Amount int `json:"amount"`
}

func infoFromContext(c *gin.Context) *dtmcli.BranchBarrier {
	info := dtmcli.BranchBarrier{
		TransType: c.Query("trans_type"),
		Gid:       c.Query("gid"),
		BranchID:  c.Query("branch_id"),
		Op:        c.Query("op"),
	}
	return &info
}
