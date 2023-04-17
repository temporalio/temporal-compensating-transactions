from temporalio import activity

@activity.defn
async def get_bowl() -> None:
    print('Getting bowl')

@activity.defn
async def put_bowl_away() -> None:
   print('Putting bowl away')

@activity.defn
async def add_cereal() -> None:
   print('Adding cereal')

@activity.defn
async def put_cereal_back_in_box() -> None:
   print('Putting cereal back in box')

@activity.defn
async def add_milk() -> None:
    print('Adding milk')
    