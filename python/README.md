# Compensating Transactions: A silly illustration of the this design pattern in Temporal through breakfast

## Instructions

Ensure you have Python 3.7 or later installed locally.


Create a virtual environment for your project and install the Temporal SDK:

```bash
python3 -m venv env
source env/bin/activate
python -m pip install temporalio pytest pytest-asyncio
```

Run the `temporal server start-dev` which starts the [Temporal server](https://docs.temporal.
io/docs/server/quick-install).
By default, you should see what things are running on your server at [http://localhost:8233/](http://localhost:8233/)

Run the worker and starter included in the project.

```bash
python run_worker.py
python run_workflow.py
```
