package app

import (
	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmcli/dtmimp"
	"github.com/dtm-labs/dtmcli/logger"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// TransInCompensateHandler 回滚操作
func TransInCompensateHandler(c *gin.Context) {
	info := InfoFromContext(c)
	inputDto := reqFrom(c)
	barrier := MustBarrierFromGin(c)

	log.Printf("TransInCompensate:%vgid:%v", inputDto.Amount, info.Gid)
	err := SagaAdjustBalanceService(barrier, inputDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtmimp.OrString(MainSwitch.QueryPreparedResult.Fetch(), dtmcli.ResultFailure))
		return
	}
	c.JSON(http.StatusOK, dtmimp.OrString(MainSwitch.QueryPreparedResult.Fetch(), dtmcli.ResultSuccess))
}

// TransInHandler 正常事务
func TransInHandler(c *gin.Context) {
	info := InfoFromContext(c)

	inputDto := reqFrom(c)
	barrier := MustBarrierFromGin(c)
	log.Printf("TransIn:%v,gid:%v", inputDto.Amount, info.Gid)

	err := SagaAdjustBalanceCompensateService(barrier, inputDto)
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
	ti.BarrierTableName = " dtm_barrier.barrier"
	logger.FatalIfError(err)
	return ti
}
