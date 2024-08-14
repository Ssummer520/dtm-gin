package main

import (
	"fmt"
	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmcli/dtmimp"
	"github.com/dtm-labs/dtmcli/logger"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	QsStartSvr()

}

// busi address
const qsBusiAPI = "/api/busi_start"
const qsBusiPort = 8881

// QsStartSvr quick start: start server
func QsStartSvr() {
	app := gin.Default()
	qsAddRoute(app)
	log.Printf("quick start examples listening at %d", qsBusiPort)

	app.Run(fmt.Sprintf(":%d", qsBusiPort))

}

func qsAddRoute(app *gin.Engine) {
	app.POST(qsBusiAPI+"/TransIn", func(c *gin.Context) {
		log.Printf("TransIn")
		// c.JSON(200, "")

		time.Sleep(time.Second * 5)
		c.JSON(http.StatusConflict, dtmimp.OrString(MainSwitch.QueryPreparedResult.Fetch(), dtmcli.ResultFailure)) // Status 409 for Failure. Won't be retried
	})
	app.POST(qsBusiAPI+"/TransInCompensate", func(c *gin.Context) {
		log.Printf("TransInCompensate")
		c.JSON(200, "")
	})

}
func string2DtmError(str string) error {
	return map[string]error{
		dtmcli.ResultFailure: dtmcli.ErrFailure,
		dtmcli.ResultOngoing: dtmcli.ErrOngoing,
		dtmcli.ResultSuccess: nil,
		"":                   nil,
	}[str]
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
