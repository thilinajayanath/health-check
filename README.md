# health-check

A simple monitoring app based on periodic pings to a HTTP endpoint. Sends a SNS notification once a given number of pings are missed.

## Endpoints

The health-check application runs on port `8080` and there are two endpoints.

- `/` - This endpoint listens for ping messages, which are `HTTP Post` requests.
- `/reset` - This endpoint resets the alert.

For example:

```bash
curl -X POST localhost:8080/ # ping the endpoint
curl -X POST localhost:8080/reset # resets the alert
```

## Input Arguments

interval - Time period in minutes between pings. Defaults to 15 minutes.  
threshold - Number of pings to miss before alerting. Defaults to 1.  
realert - Time to send the alert again once an alert is sent. Defaults to 120 minutes.  
topic-arn - AWS SNS topic where the alert is sent to. **(Required)**

## Requirements

- AWS SNS topic
- Configured AWS credentials profile with access to the SNS topic
