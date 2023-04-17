package app

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func BreakfastWorkflowParallel(ctx workflow.Context) (err error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, options)
	pendingFutures := make([]workflow.Future, 0)

	defer func() error {
		// Run compensations last, if we encounter errors in normal execution.
		if err != nil {
			for _, future := range pendingFutures {
				if err := future.Get(ctx, nil); err != nil {
					return err
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
		f := workflow.ExecuteActivity(ctx, PutBowlAway)
		AddCompensationParallel(ctx, f, pendingFutures, &err)
	}()

	err = workflow.ExecuteActivity(ctx, AddCereal).Get(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		f := workflow.ExecuteActivity(ctx, PutCerealBackInBox)
		AddCompensationParallel(ctx, f, pendingFutures, &err)
	}()

	err = workflow.ExecuteActivity(ctx, AddMilk).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}

func AddCompensationParallel(ctx workflow.Context, f workflow.Future, pendingFutures []workflow.Future, err *error) {
	if err != nil {
		pendingFutures = append(pendingFutures, f)
	}
}
