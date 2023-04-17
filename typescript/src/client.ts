import { Connection, Client } from '@temporalio/client';
import { breakfastWorkflow } from './workflows';
import { nanoid } from 'nanoid';

async function run() {
  // Connect to the default Server location (localhost:7233)
  const connection = await Connection.connect();

  const client = new Client({
    connection,
  });

  const handle = await client.workflow.start(breakfastWorkflow, {
    args: [],
    taskQueue: 'make-breakfast',
    // in practice, use a meaningful business ID, like customerId or transactionId
    workflowId: 'workflow-' + nanoid(),
  });
  console.log(`Started workflow ${handle.workflowId}`);
}

run().catch((err) => {
  console.error(err);
  process.exit(1);
});
