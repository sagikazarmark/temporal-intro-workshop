# Temporal Intro Workshop

This repository contains example code for my [Temporal Intro Workshop](https://sagikazarmark.hu/slides/workshops/temporal-intro).

Record: https://youtu.be/UwdGmdTO3Ts


## Prerequisites

1. Git, Make, etc.
2. Make sure you have the latest [Go](https://golang.org/) and [Docker](https://www.docker.com/get-started) installed


## Usage

1. Checkout this repository
2. Run `make up`
3. Wait for Temporal to start
4. Check if Temporal is running with `make ps`
5. Start a new shell with `make shell`

Alternatively, you can use the following alias for the `tctl` commands instead of opening a new shell:

```shell
alias tctl='docker compose exec temporal-admin-tools tctl'
```


## Running a workflow using the CLI

You can run a workflow using the CLI with the following command:

```bash
tctl workflow run --taskqueue workshop --execution_timeout 60 --workflow_type WORKFLOW_TYPE -i 'arg1 arg2...'
```

As a best practice, workflows generally have a single input struct (to remain compatible with other languages).
By default, Temporal uses JSON encoding, so such workflow execution looks like this:

```bash
tctl workflow run --taskqueue workshop --execution_timeout 60 --workflow_type example01 -i '{"A": 1, "B": 2}'
```

You can shorten the command a lot by using shorthands for commands and options:

```bash
tctl wf run --tq workshop --et 60 --wt example01 -i '{"A": 1, "B": 2}'
```

Last, but not least, if you want to start a workflow without waiting for its result,
you can do so by using the `start` command instead of `run`:

```bash
tctl wf start --tq workshop --et 60 --wt example01 -i '{"A": 1, "B": 2}'
```


## Quering workflow state using the CLI

Workflows can register query handlers to expose state about themselves. You can query that state using the following command:

```bash
tctl workflow query --workflow_id 72daa600-3cac-49b0-9e86-277a47c80a87 --query_type current_number
```

Or using a shorter version:

```bash
tctl wf query --wid 72daa600-3cac-49b0-9e86-277a47c80a87 --qt current_number
```

There is a special query type called `__stack_trace` that gives you the current stack trace of the workflow.
Useful if a workflow is stuck for a long time and you want to check where it stopped.


## Signaling a workflow using the CLI

Workflows can register query handlers to expose state about themselves. You can query that state using the following command:

```bash
tctl workflow signal --workflow_id 72daa600-3cac-49b0-9e86-277a47c80a87 --name set_number --input '2'
```

Or using a shorter version:

```bash
tctl wf signal --wid 72daa600-3cac-49b0-9e86-277a47c80a87 -n set_number -i '2'
```

There is a special query type called `__stack_trace` that gives you the current stack trace of the workflow.
Useful if a workflow is stuck for a long time and you want to check where it stopped.


## Cleaning up

Once you are finished with the workshop, you can clean up all resources (containers) by running the following command:

```bash
make down
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
