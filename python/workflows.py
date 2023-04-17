from datetime import timedelta
from temporalio import workflow
from temporalio.exceptions import ActivityError, CancelledError
import asyncio
import typing

with workflow.unsafe.imports_passed_through():
    from activities import get_bowl, put_bowl_away, add_cereal, put_cereal_back_in_box, add_milk


class Compensations:
   def __init__(self, parallel_compensations=False):
       self.parallel_compensations = parallel_compensations
       self.compensations = []


   def add(self, function: typing.Awaitable):
       self.compensations.append(function)


   def __iadd__(self, function: typing.Awaitable):
       self.add(function)
       return self
  
   async def compensate(self):
       if self.parallel_compensations:
           try:
               await asyncio.gather(self.compensations, return_exceptions=True)
           except Exception as e:
               print('failed to compensate: %s' % e)
       else:
            for f in reversed(self.compensations):
               try:
                    await f()
               except Exception as e:
                    print('failed to compensate: %s' % e)

@workflow.defn
class BreakfastWorkflow:
    @workflow.run
    async def run(self) -> None:
       compensations = Compensations()
       try:
           await workflow.execute_activity(get_bowl, start_to_close_timeout=timedelta(seconds=5))
           compensations += put_bowl_away
           await workflow.execute_activity(add_cereal, start_to_close_timeout=timedelta(seconds=5))
           compensations += put_cereal_back_in_box
           await workflow.execute_activity(add_milk, start_to_close_timeout=timedelta(seconds=5))
       except (ActivityError, CancelledError):
           task = asyncio.create_task(compensations.compensate())
           # Ensure the compensations run in the face of cancelation.
           await asyncio.shield(task)
           raise
