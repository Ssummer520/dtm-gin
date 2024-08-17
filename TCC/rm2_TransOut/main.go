package main

import (
	"dtm-gin/TCC/rm2_TransOut/app"
	"fmt"
	_ "github.com/dtm-labs/dtmcli/logger"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // 确保导入 MySQL 驱动
	"github.com/grafana/pyroscope-go"
	_ "github.com/grafana/pyroscope-go"
	"runtime"
	"time"
)

// busi address
const qsBusiAPI = "/app/busi_start"
const qsBusiPort = 8888

func main() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("错误:", err)
		return
	}
	time.Local = loc
	r := gin.Default()
	r.POST(qsBusiAPI+"/TccBTransOutTry", app.TccBTransOutTryHandler)
	r.POST(qsBusiAPI+"/TccBTransOutConfirm", app.TccBTransOutConfirmHandler)
	r.POST(qsBusiAPI+"/TccBTransOutCancel", app.TccBTransOutCancelHandler)

	// These 2 lines are only required if you're using mutex or block profiling
	// Read the explanation below for how to set these rates:
	runtime.SetMutexProfileFraction(5)
	runtime.SetBlockProfileRate(5)
	fmt.Println()
	pyroscope.Start(pyroscope.Config{
		ApplicationName: "TransOut.golang.app",

		// replace this with the address of pyroscope server
		ServerAddress: "http://localhost:4040",

		// you can disable logging by setting this to nil

		// you can provide static tags via a map:
		Tags: map[string]string{"hostname": "TransOut"},

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

	// your code goes here

	r.Run(fmt.Sprintf(":%d", qsBusiPort))
}
