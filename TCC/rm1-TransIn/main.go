package main

import (
	"dtm-gin/TCC/rm1-TransIn/app"

	"fmt"
	_ "github.com/dtm-labs/dtmcli/logger"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // 确保导入 MySQL 驱动
	"github.com/grafana/pyroscope-go"
	"runtime"
)

func main() {
	QsStartSvr()

}

// busi address
const qsBusiAPI = "/app/busi_start"
const qsBusiPort = 8882

// QsStartSvr quick start: start server
func QsStartSvr() {
	app := gin.Default()
	qsAddRoute(app)
	// These 2 lines are only required if you're using mutex or block profiling
	// Read the explanation below for how to set these rates:
	runtime.SetMutexProfileFraction(5)
	runtime.SetBlockProfileRate(5)

	pyroscope.Start(pyroscope.Config{
		ApplicationName: "TransIn.golang.app",

		// replace this with the address of pyroscope server
		ServerAddress: "http://localhost:4040",

		// you can disable logging by setting this to nil

		// you can provide static tags via a map:
		Tags: map[string]string{"hostname": "TransIn"},

		ProfileTypes: []pyroscope.ProfileType{
			// these profile types are enabled by default:
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,

			// these profile types are optional:
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})
	app.Run(fmt.Sprintf(":%d", qsBusiPort))

}

func qsAddRoute(r *gin.Engine) {
	r.POST(qsBusiAPI+"/TccBTransInTry", app.TccBTransInTryHandler)
	r.POST(qsBusiAPI+"/TccBTransInConfirm", app.TccBTransInConfirmHandler)
	r.POST(qsBusiAPI+"/TccBTransInCancel", app.TccBTransInCancelHandler)

}
