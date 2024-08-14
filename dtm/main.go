package main

import (
	"fmt"
	"github.com/dtm-labs/dtmcli"
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v3"
	"log"
)

func main() {
	app := gin.Default()

	app.GET("/test", func(c *gin.Context) {
		QsFireRequest()
		log.Printf("TransOut")
		c.JSON(200, "sss")
	})
	app.Run(":1111")

}

const qsBusiAPI = "/api/busi_start"
const qsBusiPortIN = 8881
const qsBusiPortOUT = 8880
const dtmServer = "http://localhost:36789/api/dtmsvr"

var qsBusiIN = fmt.Sprintf("http://host.docker.internal:%d%s", qsBusiPortIN, qsBusiAPI)
var qsBusiOUT = fmt.Sprintf("http://host.docker.internal:%d%s", qsBusiPortOUT, qsBusiAPI)

func QsFireRequest() string {
	req := &gin.H{"amount": 30} // load of micro-service
	// DtmServer is the url of dtm
	saga := dtmcli.NewSaga(dtmServer, shortuuid.New()).
		// add a TransOut sub-transaction，forward operation with url: qsBusi+"/TransOut", reverse compensation operation with url: qsBusi+"/TransOutCompensate"
		Add(qsBusiOUT+"/TransOut", qsBusiOUT+"/TransOutCompensate", req).
		// add a TransIn sub-transaction, forward operation with url: qsBusi+"/TransIn", reverse compensation operation with url: qsBusi+"/TransInCompensate"
		Add(qsBusiIN+"/TransIn", qsBusiIN+"/TransInCompensate", req)
	// submit the created saga transaction，dtm ensures all sub-transactions either complete or get revoked
	saga.RetryInterval = 3
	saga.RequestTimeout = 10
	err := saga.Submit()

	if err != nil {
		panic(err)
	}
	return saga.Gid
}
