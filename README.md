## Simple [machinery](https://github.com/RichardKnop/machinery) usage

App starts machinery server with:

1. Task "PingTask" - simply prints "Pong"
2. A periodic task - "periodic-notifier" which initiates "PingTask" once per minute

**Expected result**

PingTask is started once per minute

**Actual result**

PingTask executed only once


### How to start

Start redis (used as broker and result backend):
```
docker-compose up redis
```

Start the app
```
go run app/main.go
```