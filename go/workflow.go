package app

import (
	"log"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func BreakfastWorkflow(ctx workflow.Context) (err error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
		RetryPolicy:         &temporal.RetryPolicy{MaximumAttempts: 1},
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	err = workflow.ExecuteActivity(ctx, GetBowl).Get(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			// activity failed, and workflow context is canceled
			disconnectedCtx, _ := workflow.NewDisconnectedContext(ctx)
			errCompensation := workflow.ExecuteActivity(disconnectedCtx, PutBowlAway).Get(ctx, nil)
			if errCompensation != nil {
				log.Println("Executing bowl compensation failed", errCompensation)
			}
		}
	}()

	err = workflow.ExecuteActivity(ctx, AddCereal).Get(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			disconnectedCtx, _ := workflow.NewDisconnectedContext(ctx)
			errCompensation := workflow.ExecuteActivity(disconnectedCtx, PutCerealBackInBox).Get(ctx, nil)
			if errCompensation != nil {
				log.Println("Executing cereal compensation failed", errCompensation)
			}
		}
	}()

	err = workflow.ExecuteActivity(ctx, AddMilk).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
