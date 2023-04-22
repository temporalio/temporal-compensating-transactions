package app

import (
	"log"

	"go.temporal.io/sdk/workflow"
)

const BreakfastTaskQueue = "BREAKFAST_TASK_QUEUE"

type Compensations []any

func (s *Compensations) AddCompensation(activity any) {
	*s = append(*s, activity)
}

func (s Compensations) Compensate(ctx workflow.Context, in_parallel bool) {
	if !in_parallel {
		for i := len(s) - 1; i >= 0; i-- {
			errCompensation := workflow.ExecuteActivity(ctx, s[i]).Get(ctx, nil)
			if errCompensation != nil {
				log.Println("Executing compensation failed", errCompensation)
			}
		}
	} else {
		var pendingFutures []workflow.Future
		for i := len(s) - 1; i >= 0; i-- {
			f := workflow.ExecuteActivity(ctx, s[i])
			pendingFutures = append(pendingFutures, f)
		}
		for _, pending := range pendingFutures {
			if err := pending.Get(ctx, nil); err != nil {
				log.Println("Executing compensation failed", err)
			}
		}

	}
}
