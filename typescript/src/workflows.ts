import { proxyActivities } from '@temporalio/workflow'
import * as activities from './activities'


const { getBowl, putBowlAway, addCereal, putCerealBackInBox, addMilk } =
  proxyActivities<typeof activities>({
    startToCloseTimeout: '1 minute',
    retry: { maximumAttempts: 1 }
  })

type Compensation = () => Promise<void>

export async function breakfastWorkflow(compensateInParallel = false): Promise<void> {
  const compensations: Compensation[] = []
  try {
    await getBowl()
    compensations.unshift(putBowlAway)
    await addCereal()
    compensations.unshift(putCerealBackInBox)

    await addMilk()
  } catch (err) {
    await compensate(compensations, compensateInParallel)
    throw err
  }
}


async function compensate(compensations: Compensation[], compensateInParallel = false) {
  if (compensateInParallel) {
    const outcomes = await Promise.allSettled(compensations.map(comp => comp()))
    for (const outcome of outcomes) {
      if (outcome.status === 'rejected') {
        console.error(`failed to compensate: ${outcome.reason.message}`)
      }
    }
    return
  }


  for (const comp of compensations) {
    try {
      await comp()
    } catch (err) {
      console.error(`failed to compensate: ${err}`)
    }
  }
}
