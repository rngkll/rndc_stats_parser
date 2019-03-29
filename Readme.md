# rndc parser in go

## Build

## Development
To view the help
`go run main.go -h`

To view the list of options
`go run main.go -sFile ../test/named_stats.txt -list`

To view the list of options
`go run main.go -sFile ../test/named_stats.txt -option <selected option>`

## Linux and bind

Most of the bind servers are running in a linux environment, if you don't do the build in the same OS, your executable won't run in a linux environment, to solve this, I do a build with docker.

## Cronjob

## Get linux binary
```
docker build . -t rndc_parser
docker run -it rndc_parser bash
```
In other terminal, run:
```
docker cp rndc_parser:/go/src/app/rndc_parser.linux .
```