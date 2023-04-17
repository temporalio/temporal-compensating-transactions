package breakfastapp;

import static org.mockito.ArgumentMatchers.eq;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.verify;

import io.temporal.client.WorkflowClient;
import io.temporal.client.WorkflowOptions;
import io.temporal.testing.TestWorkflowEnvironment;
import io.temporal.worker.Worker;

import org.junit.After;
import org.junit.Before;
import org.junit.Test;

public class RecipeCreatorGenerationWorkflowTest {

    private TestWorkflowEnvironment testEnv;
    private Worker worker;
    private WorkflowClient workflowClient;

    @Before
    public void setUp() {
        testEnv = TestWorkflowEnvironment.newInstance();
        worker = testEnv.newWorker(Shared.BREAKFAST_TASK_QUEUE);
        worker.registerWorkflowImplementationTypes(BreakfastWorkflowImpl.class);
        workflowClient = testEnv.getWorkflowClient();
    }

    @After
    public void tearDown() {
        testEnv.close();
    }

    @Test
    public void testTransfer() {
        BreakfastActivity breakfast = mock(BreakfastActivityImpl.class);
        worker.registerActivitiesImplementations(breakfast);

        testEnv.start();
        WorkflowOptions options = WorkflowOptions.newBuilder()
                .setTaskQueue(Shared.BREAKFAST_TASK_QUEUE)
                .build();
        BreakfastWorkflow workflow = workflowClient.newWorkflowStub(BreakfastWorkflow.class, options);
        workflow.makeBreakfast();
        verify(breakfast).getBowl();
        verify(breakfast).addCereal();
        verify(breakfast).addMilk();
    }
}
