

export type Compensation = () => Promise<void>

export async function compensate(compensations: Compensation[], compensateInParallel = false) {
  if (compensateInParallel) {
    await Promise.allSettled(
      compensations.map((comp) =>
        comp().catch((err) => console.error(`failed to compensate: $error`))
      )
    )
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
