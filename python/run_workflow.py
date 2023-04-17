import asyncio

from run_worker import BreakfastWorkflow
from temporalio.client import Client


async def main():
    # Create client connected to server at the given address
    client = await Client.connect("localhost:7233")

    # Execute a workflow
    result = await client.execute_workflow(
        BreakfastWorkflow.run, id="breakfast-workflow", task_queue="breakfast-queue"
    )

    print(f"Result: {result}")


if __name__ == "__main__":
    asyncio.run(main())
