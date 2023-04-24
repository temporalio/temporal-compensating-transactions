from datetime import timedelta
from temporalio import workflow
from temporalio.common import RetryPolicy
import asyncio
import typing

with workflow.unsafe.imports_passed_through():
    from activities import (
        get_bowl,
        put_bowl_away,
        add_cereal,
        put_cereal_back_in_box,
        add_milk,
    )

# Trying only once to illustrate compensations easily when activities fail.
common_retry_policy = RetryPolicy(maximum_attempts=1)
time_delta = timedelta(seconds=5)


class Compensations:
    def __init__(self, parallel_compensations=False):
        self.parallel_compensations = parallel_compensations
        self.compensations = []

    def add(self, function: typing.Callable[..., typing.Awaitable[None]]):
        self.compensations.append(function)

    def __iadd__(self, function: typing.Callable[..., typing.Awaitable[None]]):
        self.add(function)
        return self

    async def compensate(self):
        if self.parallel_compensations:

            # Mini function separated out here purely for readability of the list 
            # comprehension below.
            def compensation_lambda(f):
                return workflow.execute_activity(
                    f,
                    start_to_close_timeout=time_delta,
                    retry_policy=common_retry_policy,
                )

            all_compensations = [compensation_lambda(c) for c in self.compensations]
            results = await asyncio.gather(*all_compensations)
            for result in results:
                if isinstance(result, Exception):
                    workflow.logger("failed to compensate: %s" % result)

        else:
            for f in reversed(self.compensations):
                try:
                    await workflow.execute_activity(
                        f,
                        start_to_close_timeout=time_delta,
                        retry_policy=common_retry_policy,
                    )
                except Exception as e:
                    workflow.logger("failed to compensate: %s" % result)


@workflow.defn
class BreakfastWorkflow:
    @workflow.run
    async def run(self, parallel_compensations) -> None:
        compensations = Compensations(parallel_compensations=parallel_compensations)
        try:
            await workflow.execute_activity(
                get_bowl,
                start_to_close_timeout=time_delta,
                retry_policy=common_retry_policy,
            )
            compensations += put_bowl_away
            await workflow.execute_activity(
                add_cereal,
                start_to_close_timeout=time_delta,
                retry_policy=common_retry_policy,
            )
            compensations += put_cereal_back_in_box
            await workflow.execute_activity(
                add_milk,
                start_to_close_timeout=time_delta,
                retry_policy=common_retry_policy,
            )
        except Exception:
            task = asyncio.create_task(compensations.compensate())
            # Ensure the compensations run in the face of cancelation.
            await asyncio.shield(task)
            raise
