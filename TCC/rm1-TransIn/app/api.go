package app

import (
	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmcli/dtmimp"
	"github.com/dtm-labs/dtmcli/logger"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// TccBTransInTryHandler TRY
func TccBTransInTryHandler(c *gin.Context) {
	info := InfoFromContext(c)
	inputDto := reqFrom(c)
	barrier := MustBarrierFromGin(c)

	log.Printf("TccBTransInTryHandler:%vgid:%v", inputDto.Amount, info.Gid)
	err := TccBTransInTryService(barrier, inputDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtmimp.OrString(MainSwitch.QueryPreparedResult.Fetch(), dtmcli.ResultFailure))
		return
	}
	c.JSON(http.StatusOK, dtmimp.OrString(MainSwitch.QueryPreparedResult.Fetch(), dtmcli.ResultSuccess))
}

// TccBTransInConfirmHandler Confirm
func TccBTransInConfirmHandler(c *gin.Context) {
	info := InfoFromContext(c)

	inputDto := reqFrom(c)
	barrier := MustBarrierFromGin(c)
	log.Printf("TccBTransInConfirmHandler:%v,gid:%v", inputDto.Amount, info.Gid)

	err := TccBTransInConfirmService(barrier, inputDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtmimp.OrString(MainSwitch.QueryPreparedResult.Fetch(), dtmcli.ResultFailure))
		return
	}
	c.JSON(http.StatusOK, dtmimp.OrString(MainSwitch.QueryPreparedResult.Fetch(), dtmcli.ResultSuccess))

}

// TccBTransInCancelHandler Cancel
func TccBTransInCancelHandler(c *gin.Context) {
	info := InfoFromContext(c)

	inputDto := reqFrom(c)
	barrier := MustBarrierFromGin(c)
	log.Printf("TccBTransInCancelHandler:%v,gid:%v", inputDto.Amount, info.Gid)

	err := TccBarrierTransInCancelService(barrier, inputDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtmimp.OrString(MainSwitch.QueryPreparedResult.Fetch(), dtmcli.ResultFailure))
		return
	}
	c.JSON(http.StatusOK, dtmimp.OrString(MainSwitch.QueryPreparedResult.Fetch(), dtmcli.ResultSuccess))

}
func reqFrom(c *gin.Context) *ReqHTTP {
	v, ok := c.Get("trans_req")
	if !ok {
		req := ReqHTTP{}
		err := c.BindJSON(&req)
		logger.FatalIfError(err)
		c.Set("trans_req", &req)
		v = &req
	}
	return v.(*ReqHTTP)
}

// MustBarrierFromGin 1
func MustBarrierFromGin(c *gin.Context) *dtmcli.BranchBarrier {
	ti, err := dtmcli.BarrierFromQuery(c.Request.URL.Query())
	logger.FatalIfError(err)
	ti.BarrierTableName = " dtm_barrier.barrier"
	return ti
}
