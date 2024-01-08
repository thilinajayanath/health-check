package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/thilinajayanath/health-check/internal/handler"
)

func parseInput() (int, int, int, *string) {
	interval := flag.Int("interval", 15, "Time duration between pings in minutes")
	realert := flag.Int("realert", 120, "Time duration between alerts in minutes")
	threshold := flag.Int("threshold", 1, "Number of ping misses before alerting")
	topicArn := flag.String("topic-arn", "", "Topic of the SNS ARN to send the message")

	flag.Parse()

	return *interval, *realert, *threshold, topicArn
}

func main() {
	interval, realert, threshold, topicArn := parseInput()

	if *topicArn == "" {
		fmt.Println("ARN of the SNS topic is missing")
		os.Exit(0)
	}

	log.Printf("Setting the time period between pings to %d minutes and the threshold to %d\n", interval, threshold)

	handler.HandleRequests(interval, realert, threshold, topicArn)
}
