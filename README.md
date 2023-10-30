# health-check

A simple monitoring application based on periodic pings to a HTTP endpoint. Sends a SNS notification once a given number of pings are missed.

## Endpoints

The health-check application runs on port `8080` and there are two endpoints.

- `/` - This endpoint listens for ping messages, which are `HTTP Post` requests.
- `/reset` - This endpoint resets the alert.

For example:

```bash
curl -X POST localhost:8080/ # ping the endpoint
curl -X POST localhost:8080/reset # resets the alert
```

## Building and running the application

```bash
# clone the git repo
git clone git@github.com:thilinajayanath/health-check.git

# build the application
cd health-check
go build -o ./app cmd/health-check/main.go

# run the application
./app -topic-arn "<topic-arn>"
```

## Input Arguments

| Argument  | Description                                                                             |
| --------- | --------------------------------------------------------------------------------------- |
| topic-arn | AWS SNS topic where the alert is sent to. **(Required)**                                |
| interval  | Time period in minutes between pings. Defaults to 15 minutes. (optional)                |
| threshold | Number of pings to miss before alerting. Defaults to 1. (optional)                      |
| realert   | Time to send the alert again once an alert is sent. Defaults to 120 minutes. (optional) |

## Requirements

- AWS SNS topic
- Configured AWS credentials profile with access to the SNS topic
