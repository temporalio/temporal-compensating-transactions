package main

import (
	"breakfast/app"
	"context"
	"flag"
	"log"

	"go.temporal.io/sdk/client"
)

func main() {

	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "breakfast-workflow",
		TaskQueue: app.BreakfastTaskQueue,
	}

	// Start the Workflow
	parallelCompensationsPtr := flag.Bool("parallel-compensations", false, "Execute compensations in parallel if possible.")
	flag.Parse()
	if !*parallelCompensationsPtr {
		_, err = c.ExecuteWorkflow(context.Background(), options, app.BreakfastWorkflow)
	} else {
		_, err = c.ExecuteWorkflow(context.Background(), options, app.BreakfastWorkflowParallel)
	}

	if err != nil {
		log.Fatalln("unable to start Workflow", err)
	}
}
