package app

import (
	"log"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func BreakfastWorkflow(ctx workflow.Context, parallel_compensations bool) (err error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
		RetryPolicy:         &temporal.RetryPolicy{MaximumAttempts: 1},
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	var compensations Compensations

	err = workflow.ExecuteActivity(ctx, GetBowl).Get(ctx, nil)
	compensations.AddCompensation(PutBowlAway)
	if err != nil {
		return err
	}

	err = workflow.ExecuteActivity(ctx, AddCereal).Get(ctx, nil)
	compensations.AddCompensation(PutCerealBackInBox)
	if err != nil {
		return err
	}

	err = workflow.ExecuteActivity(ctx, AddMilk).Get(ctx, nil)

	log.Println("just before")
	log.Println(err)

	defer func() (err error) {
		log.Println("in the compensation func")
		log.Println(err)
		if err != nil {
			log.Println("3333333")

			// activity failed, and workflow context is canceled
			disconnectedCtx, _ := workflow.NewDisconnectedContext(ctx)
			compensations.Compensate(disconnectedCtx, parallel_compensations)
		}
		return err
	}()

	return err
}
