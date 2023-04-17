package breakfastapp;

import io.temporal.workflow.WorkflowInterface;
import io.temporal.workflow.WorkflowMethod;

@WorkflowInterface
public interface BreakfastWorkflow {

    // The Workflow method is called by the initiator either via code or CLI.
    @WorkflowMethod
    void makeBreakfast();
}