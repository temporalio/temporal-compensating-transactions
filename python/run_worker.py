import asyncio

from temporalio import activity, workflow
from temporalio.client import Client
from temporalio.worker import Worker

from activities import (
    get_bowl,
    put_bowl_away_if_present,
    add_cereal,
    put_cereal_back_in_box_if_present,
    add_milk,
)
from workflows import BreakfastWorkflow


async def main():
    client = await Client.connect("localhost:7233", namespace="default")
    # Run the worker
    worker = Worker(
        client,
        task_queue="breakfast-queue",
        workflows=[BreakfastWorkflow],
        activities=[
            get_bowl,
            put_bowl_away_if_present,
            add_cereal,
            put_cereal_back_in_box_if_present,
            add_milk,
        ],
    )
    await worker.run()


if __name__ == "__main__":
    asyncio.run(main())
