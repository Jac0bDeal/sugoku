# sugoku
![build-status](https://github.com/Jac0bDeal/sugoku/workflows/Build/badge.svg?branch=main)
![tests-status](https://github.com/Jac0bDeal/sugoku/workflows/Tests/badge.svg?branch=main)
[![go-report-card](https://goreportcard.com/badge/github.com/Jac0bDeal/sugoku)](https://goreportcard.com/report/github.com/Jac0bDeal/sugoku)

A sudoku generator and solver written in Go.

## Build
In order to build this, you need Go 1.19+ and Make installed.

From the project root, run
```shell script
make sugoku
```

This will build the binary at `bin/sugoku`.

## Running
Once built, the binary is run with
```shell script
./bin/sugoku
```

## Tests
If you want to run the tests (for some reason) use
```shell script
make test
```
