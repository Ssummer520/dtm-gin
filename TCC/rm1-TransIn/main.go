package main

import (
	"context"
	"dtm-gin/TCC/rm1-TransIn/app"
	"fmt"
	"log"
	"net/http"
	"time"

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
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("错误:", err)
		return
	}
	time.Local = loc
	app := gin.Default()
	app.Use(gin.Recovery())
	app.Use(timeoutMiddleware(time.Second))
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

	s := &http.Server{
		Addr:           ":8882",
		Handler:        app,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("Starting server on %s\n", s.Addr)
	err = s.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed: %s", err)
	} else {

	}

}

func qsAddRoute(r *gin.Engine) {
	r.POST(qsBusiAPI+"/TccBTransInTry", app.TccBTransInTryHandler)
	r.POST(qsBusiAPI+"/TccBTransInConfirm", app.TccBTransInConfirmHandler)
	r.POST(qsBusiAPI+"/TccBTransInCancel", app.TccBTransInCancelHandler)

}
func timeoutMiddleware(duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), duration)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		done := make(chan struct{})
		go func() {
			c.Next()
			close(done)
		}()

		select {
		case <-ctx.Done():
			c.JSON(http.StatusGatewayTimeout, gin.H{"error": "request timed out"})
			c.Abort()
		case <-done:
		}
	}
}
