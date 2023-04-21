package app

import (
	"log"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// A type with the compensations and their "disconnected contexts",
// which are distinct from the standard workflow.Contexts to protect
// against cancellation.
type FutureWithContext struct {
	context workflow.Context
	future  workflow.Future
}

func BreakfastWorkflowParallel(ctx workflow.Context) (err error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
		RetryPolicy:         &temporal.RetryPolicy{MaximumAttempts: 1},
	}

	ctx = workflow.WithActivityOptions(ctx, options)
	var pendingFutures []FutureWithContext

	defer func() error {
		// Run compensations last, if we encounter errors in normal execution.
		if err != nil {
			for _, pending := range pendingFutures {
				if err := pending.future.Get(pending.context, nil); err != nil {
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
			disconnectedCtx, _ := workflow.NewDisconnectedContext(ctx)
			f := workflow.ExecuteActivity(disconnectedCtx, PutBowlAway)
			pendingFutures = append(pendingFutures, FutureWithContext{disconnectedCtx, f})
		}
	}()

	err = workflow.ExecuteActivity(ctx, AddCereal).Get(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			disconnectedCtx, _ := workflow.NewDisconnectedContext(ctx)
			f := workflow.ExecuteActivity(disconnectedCtx, PutCerealBackInBox)
			pendingFutures = append(pendingFutures, FutureWithContext{disconnectedCtx, f})
		}
	}()

	err = workflow.ExecuteActivity(ctx, AddMilk).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
