package breakfastapp;

import io.temporal.api.common.v1.WorkflowExecution;
import io.temporal.client.WorkflowClient;
import io.temporal.client.WorkflowOptions;
import io.temporal.serviceclient.WorkflowServiceStubs;

public class InitiateBreakfastMaking {

    public static void main(String[] args) throws Exception {

        WorkflowServiceStubs service = WorkflowServiceStubs.newLocalServiceStubs();
        WorkflowOptions options = WorkflowOptions.newBuilder()
                .setTaskQueue(Shared.BREAKFAST_TASK_QUEUE)
                .setWorkflowId("breakfast-workflow")
                .build();
        WorkflowClient client = WorkflowClient.newInstance(service);
        BreakfastWorkflow workflow = client.newWorkflowStub(BreakfastWorkflow.class, options);
        WorkflowExecution we = WorkflowClient.start(workflow::makeBreakfast);
        System.out.printf("\nWorkflowID: %s RunID: %s", we.getWorkflowId(), we.getRunId());
        System.exit(0);
    }
}
