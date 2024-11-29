# gojs

## Intro
This is a poc demo to show how to share data between two go processes launched via exec.Command, streaming data from the first process into a buffer and then streaming the buffer into the stdin of the second process.

## Demo
To run the demo, first build the typescript packages:
- `cd ts1 && npm i && npm run build`
- `cd ts2 && npm i && npm run build`

Then install the go package:
- `go install`

Then run the exec2 command:
- `gojs exec2`

You will see data piped from the stdout of the `ts1` app, into the stdin of the `ts2` app, via the ob outputBytes buffer in the go command.