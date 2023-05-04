from temporalio import activity

# To introduce a failure, add an Exception in one of these methods.


@activity.defn
async def get_bowl() -> None:
    print("Getting bowl")


@activity.defn
async def put_bowl_away_if_present() -> None:
    print("Putting bowl away if bowl is out")


@activity.defn
async def add_cereal() -> None:
    print("Adding cereal")


@activity.defn
async def put_cereal_back_in_box_if_present() -> None:
    print("Putting cereal back in box if there is cereal")


@activity.defn
async def add_milk() -> None:
    print("Adding milk")
