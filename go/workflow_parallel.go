package app

import (
	"log"
	"time"

	"go.temporal.io/sdk/workflow"
)

func BreakfastWorkflowParallel(ctx workflow.Context) (err error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, options)
	var pendingFutures []workflow.Future

	defer func() error {
		// Run compensations last, if we encounter errors in normal execution.
		if err != nil {
			for _, future := range pendingFutures {
				if err := future.Get(ctx, nil); err != nil {
					log.Println("Executing compensation failed", err)
				}
			}
		}
		return nil
	}()

	err = workflow.ExecuteActivity(ctx, GetBowl).Get(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			f := workflow.ExecuteActivity(ctx, PutBowlAway)
			pendingFutures = append(pendingFutures, f)
		}
	}()

	err = workflow.ExecuteActivity(ctx, AddCereal).Get(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			f := workflow.ExecuteActivity(ctx, PutCerealBackInBox)
			pendingFutures = append(pendingFutures, f)
		}
	}()

	err = workflow.ExecuteActivity(ctx, AddMilk).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
