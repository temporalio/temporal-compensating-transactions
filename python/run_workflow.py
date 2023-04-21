import asyncio
import argparse

from run_worker import BreakfastWorkflow
from temporalio.client import Client


async def main():
    # Create client connected to server at the given address
    client = await Client.connect("localhost:7233")

    parser = argparse.ArgumentParser()
    parser.add_argument('--parallel', '--parallel-compensations', '-p', action='store_true')
    args = parser.parse_args()
    # Execute a workflow
    await client.execute_workflow(
        BreakfastWorkflow.run, args.parallel, id="breakfast-workflow", task_queue="breakfast-queue"
    )

    print("Started workflow")


if __name__ == "__main__":
    asyncio.run(main())
