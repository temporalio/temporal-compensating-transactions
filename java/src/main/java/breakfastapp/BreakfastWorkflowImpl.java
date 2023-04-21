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
            .setMaximumInterval(Duration.ofSeconds(100))
            .setBackoffCoefficient(2)
            .setMaximumAttempts(500)
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
    public void makeBreakfast() {
        // You can set parallel compensations if appropriate with the Builder
        Saga saga = new Saga(new Saga.Options.Builder().build());
        try {
            breakfastActivity.getBowl();
            saga.addCompensation(breakfastActivity::putBowlAway);
            breakfastActivity.addCereal();
            saga.addCompensation(breakfastActivity::putCerealBackInBox);
            breakfastActivity.addMilk();
        } catch (ActivityFailure e) {
            saga.compensate();
            throw e;
        }

    }
}
