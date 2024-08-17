package main

import (
	"fmt"
	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmcli/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/lithammer/shortuuid/v3"
	"log"
	"time"
)

func main() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("错误:", err)
		return
	}
	time.Local = loc
	app := gin.Default()
	app.Use(gin.Recovery())
	app.GET("/test", func(c *gin.Context) {
		QsFireRequest()
		log.Printf("TransOut")
		c.JSON(200, "sss")
	})

	app.Run(":1111")

}

const qsBusiAPI = "/app/busi_start"
const qsBusiPortIN = 8882
const qsBusiPortOUT = 8888
const dtmServer = "http://localhost:36789/api/dtmsvr"

var qsBusiIN = fmt.Sprintf("http://localhost:%d%s", qsBusiPortIN, qsBusiAPI)
var qsBusiOUT = fmt.Sprintf("http://localhost:%d%s", qsBusiPortOUT, qsBusiAPI)
var qsBusiIN1 = fmt.Sprintf("http://host.docker.internal:%d%s", qsBusiPortIN, qsBusiAPI)
var qsBusiOUT1 = fmt.Sprintf("http://host.docker.internal:%d%s", qsBusiPortOUT, qsBusiAPI)

type ReqHTTP struct {
	Amount int    `json:"amount"`
	UserID string `json:"userID"`
}

func QsFireRequest() string {
	defer func() {
		if err := recover(); err != nil {

		}
	}()
	reqIn := &ReqHTTP{Amount: 30, UserID: "1000"}  // load of micro-service
	reqOut := &ReqHTTP{Amount: 30, UserID: "1001"} // load of micro-service
	// DtmServer is the url of dtm
	logger.Debugf("tcc simple transaction begin")
	gid := shortuuid.New()
	err := dtmcli.TccGlobalTransaction(dtmServer, gid, func(tcc *dtmcli.Tcc) (*resty.Response, error) {
		resp, err := tcc.CallBranch(reqOut, qsBusiOUT+"/TccBTransOutTry", qsBusiOUT1+"/TccBTransOutConfirm", qsBusiOUT1+"/TccBTransOutCancel")
		if err != nil {
			return resp, err
		}
		return tcc.CallBranch(reqIn, qsBusiIN+"/TccBTransInTry", qsBusiIN1+"/TccBTransInConfirm", qsBusiIN1+"/TccBTransInCancel")
	})

	logger.Errorf(err.Error())
	return gid

}
