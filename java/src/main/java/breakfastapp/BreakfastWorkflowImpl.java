package breakfastapp;

import io.temporal.activity.ActivityOptions;
import io.temporal.failure.ActivityFailure;
import io.temporal.workflow.Workflow;
import io.temporal.workflow.Saga;
import io.temporal.common.RetryOptions;

import java.time.Duration;
public class BreakfastWorkflowImpl implements BreakfastWorkflow {
    // RetryOptions specify how to automatically handle retries when Activities
    // fail.
    private final RetryOptions retryoptions = RetryOptions.newBuilder()
            .setInitialInterval(Duration.ofSeconds(1))
            .setMaximumAttempts(1)
            .build();
    private final ActivityOptions defaultActivityOptions = ActivityOptions.newBuilder()
            .setStartToCloseTimeout(Duration.ofSeconds(5))
            .setRetryOptions(retryoptions)
            .build();
    private final BreakfastActivity breakfastActivity = Workflow.newActivityStub(BreakfastActivity.class,
            defaultActivityOptions);

    // The transfer method is the entry point to the Workflow.
    // Activity method executions can be orchestrated here or from within other
    // Activity methods.
    @Override
    public void makeBreakfast(boolean parallelCompensations) {
        // You can set parallel compensations if appropriate with the Builder
        Saga saga = new Saga(new Saga.Options.Builder().setParallelCompensation(parallelCompensations).build());
        try {
            saga.addCompensation(breakfastActivity::putBowlAwayIfPresent);
            breakfastActivity.getBowl();
            saga.addCompensation(breakfastActivity::putCerealBackInBoxIfPresent);
            breakfastActivity.addCereal();
            breakfastActivity.addMilk();
        } catch (ActivityFailure e) {
            saga.compensate();
            throw e;
        }

    }
}
