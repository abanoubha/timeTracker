# Time Tracker

A simple tiny program that traks the total time you work on a laptop/desktop computer.

## Develop

```sh
go mod tidy

go run .
```

## Build the program

```sh
go build -o timeTracker .
```

## How it works ?

Log the time the program started (startTime) and the time it is terminated (endTime) and the duration it run.

Just put the program (timeTracker) in the list of startup programs in your operating system (Ubuntu in my case).

## How to use the program ?

1. clone the repository by `git clone https://github.com/abanoubha/timeTracker.git`
2. go into the cloned repo. directory/folder by `cd timeTracker`
3. build the program from source by `go mod tidy && go build -o timeTracker .`
4. add the path to timeTracker program file in the startup (according to your operating system)

## Roadmap: versioned tasks

- v25.07.07
  - track time by saving the initial start time and the end time into a log file
- next
  - add CLI args `timeTracker -v|--view` to see the log file in terminal/stdout
