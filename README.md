# Go-server

This is my first ever web-server written in Go. I am trying to learn how to structure an Go-project
and how to use interfaces in an effective way. I also try to implement tests and other toolings to
manage the project so I may actually use this as a guide when building actual production grade projects
for myself.

## Setup

Easiest way to setup your dev-tools is by using [asdf](https://asdf-vm.com/guide/getting-started.html). Once setup and [go-plugin](https://github.com/asdf-community/asdf-golang) installed, you are good to go. Go to
project root-folder and run command:

```
asdf install
```

## Running tests

To run tests, simply run:

```
go test ./...
```

in the root of the project. You may add `-v` flag to get more verbose output from the tests if you wish.

## Formatting the code

To format the code after making changes, run `gofmt -w .` in the project root.

## Building & running

To build the app, simply run `go build .` in the project root. After that you can run the app by simply running
`./web1`. Currently the app responds from port `8080`.
