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
	var we client.WorkflowRun
	if !*parallelCompensationsPtr {
		we, err = c.ExecuteWorkflow(context.Background(), options, app.BreakfastWorkflow)
	} else {
		we, err = c.ExecuteWorkflow(context.Background(), options, app.BreakfastWorkflowParallel)
	}

	if err != nil {
		log.Fatalln("unable to start Workflow", err)
	}

	// Check for errors
	err = we.Get(context.Background(), nil)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}

}
