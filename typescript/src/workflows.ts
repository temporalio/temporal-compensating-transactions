import { proxyActivities } from '@temporalio/workflow'
import * as activities from './activities'
import {Compensation, compensate} from './compensations'



const { getBowl, putBowlAway: putBowlAwayIfPresent, addCereal, putCerealBackInBox: putCerealBackInBoxIfPresent, addMilk } =
  proxyActivities<typeof activities>({
    startToCloseTimeout: '1 minute',
    retry: { maximumAttempts: 1 }
  })

export async function breakfastWorkflow(compensateInParallel = false): Promise<void> {
  const compensations: Compensation[] = []
  try {
    compensations.unshift(putBowlAwayIfPresent)
    await getBowl()
    compensations.unshift(putCerealBackInBoxIfPresent)
    await addCereal()
    await addMilk()
  } catch (err) {
    await compensate(compensations, compensateInParallel)
    throw err
  }
}