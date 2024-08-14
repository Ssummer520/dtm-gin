package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	app := gin.Default()
	app.POST(qsBusiAPI+"/TransOut", func(c *gin.Context) {
		log.Printf("TransOut")

		c.JSON(200, "")
	})
	app.POST(qsBusiAPI+"/TransOutCompensate", func(c *gin.Context) {
		log.Printf("TransOutCompensate")
		c.JSON(200, "")
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
