package app

import (
	"time"

	"go.temporal.io/sdk/workflow"
	"go.uber.org/multierr"
)

func BreakfastWorkflow(ctx workflow.Context) (err error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	err = workflow.ExecuteActivity(ctx, GetBowl).Get(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			errCompensation := workflow.ExecuteActivity(ctx, PutBowlAway).Get(ctx, nil)
			err = multierr.Append(err, errCompensation)
		}
	}()

	err = workflow.ExecuteActivity(ctx, AddCereal).Get(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			errCompensation := workflow.ExecuteActivity(ctx, PutCerealBackInBox).Get(ctx, nil)
			err = multierr.Append(err, errCompensation)
		}
	}()

	err = workflow.ExecuteActivity(ctx, AddMilk).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
