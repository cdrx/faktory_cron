# Unofficial Faktory Cron Scheduler

`faktory-cron` is a simple tool to send jobs to a [Faktory work server](https://github.com/contribsys/faktory) on a repeating schedule, with a cron like syntax.

## Installation

`faktory-cron` is a single binary, [download the latest version](https://github.com/cdrx/faktory_cron/releases) for Linux, macOS or Windows.

## Usage

A config file (`crontab.yaml`) provides a list of jobs and the schedule they should run on. For example:

```
faktory: tcp://localhost:7419
jobs:
  - job: test
    schedule: "@every 30s"
    args:
      - test
  - job: test
    schedule: 5 * * * *
    args:
      - 1
      - 2
    retries: 1
    queue: default
    priority: 5
  - job: test.cron
    type: cron
    schedule: "0 0 * * *"
    args:
      - curl
      - -X
      - GET
      - "https://geo.api.gouv.fr/communes?codePostal=69001&fields=nom,code,codesPostaux,codeDepartement,codeRegion,population&format=json&geometry=centre"
      - -H
      - "accept: application/json"
```

To run your jobs on these schedules, point `faktory-cron` at the file:

```
$ ./faktory-cron -config path/to/crontab.yaml
```

If you have a file called `crontab.yaml` in the current directory, then you can skip the `-config` option.

#### Job options

These are the configuration options for each job:

| Setting | Required | Description|
| -------- | -------- | -------- |
| job   | ✓ | The name of the job, as registered with Faktory  |
| type   | ✓ | Type of job: faktory or cron  |
| schedule  | ✓ | See below  |
| args  | ✓ | A list of arguments to send to the job |
| retries  | ✗ | Number of times Faktory will retry the job (if it fails) |
| queue  | ✗ | The queue in which this job should be placed. Will be placed in the default queue if not supplied. |
| priority  | ✗ | Can be between 1 and 9. Defaults to 5, jobs with a higher priority will skip the queue. |

#### Cron type

Cron type allows you to exec simple shell commands without using Faktory.

First job argument is the `command`, the next job arguments are the command arguments.

#### Cron syntax

See [this link](https://godoc.org/github.com/robfig/cron]) for a full reference, but in short:


Field name   | Mandatory? | Allowed values  | Allowed special characters
----------   | ---------- | --------------  | --------------------------
Seconds      | Yes        | 0-59            | * / , -
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , - ?
Month        | Yes        | 1-12 or JAN-DEC | * / , -
Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?

#### Interval syntax

Interal jobs may be defined by using `@every`, as in:

`@every 30m`

or

`@every 30s`

You may also use `@daily`, `@hourly`, `@weekly`, `@monthly` and so on.


#### Connecting to Faktory

If the `faktory` setting isn't in the config file, then `faktory-cron` will use the `FAKTORY_URL` environment variable.

Connections are made lazily, so if you restart Faktory then you don't need to restart `faktory-cron`.

### Docker

There is a Docker container at `cdrx/faktory-cron`, you can use it like this:

```
docker run -v path/to/crontab.yaml:/crontab.yaml -e FAKTORY_URL=tcp://:password@faktory:7419 cdrx/faktory-cron
```

You'll need to supply it a config file with the schedules and point it at faktory so it can queue the jobs.

## Limitations

If you run more than one copy of `faktory-cron` from the same config file, then tasks will be duplicated. Only run one copy of the service at a time.

Faktory doesn't provide any mechanism (yet) to check if a job has finished. If your job doesn't finish before the next one is scheduled, it will start two copies of the job.

## Contributing

Contributions welcome, particularly for things from this list:

- [ ] Tests
- [ ] Live reloading if the config file changes
- [ ] Supporting multiple configuration files in a folder, possibly with cron like folders (ala `/etc/cron/daily.d`)
- [ ] Documentation

## History

#### 0.5 - 2017-11-28

First release, works for me. Feedback would be appreciated.

## License

BSD, see the `LICENSE` file.
