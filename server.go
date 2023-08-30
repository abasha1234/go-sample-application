package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	// TODO: import as directed in the install flow
	"github.com/newrelic/go-agent/v3/newrelic"

	// TODO: import for gin instrumentation
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
)

func main() {
	go client()

	//TODO: paste your creds for the app
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("<APP NAME>"),
		newrelic.ConfigLicense("xxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	if err != nil {
		log.Panic(err)
	}

	transactions := make(chan string)

	go func() {
		for {
			// Generate a new transaction.
			transaction := fmt.Sprintf("Transaction generated at %s", time.Now().Format(time.RFC3339))
			transactions <- transaction
			time.Sleep(time.Second)
		}
	}()

	// Create a Gin router.
	r := gin.Default()
	r.Use(nrgin.Middleware(app))

	r.GET("/transaction", func(c *gin.Context) {
		transaction := <-transactions
		c.JSON(200, gin.H{
			"transaction": transaction,
		})
	})

	r.Run()
}
