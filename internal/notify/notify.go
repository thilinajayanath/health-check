package notify

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func Notify(msg string, topicArn *string) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("An error occured when loading AWS configurations: %v\n", err)
	}

	client := sns.NewFromConfig(cfg)

	output, err := client.Publish(context.TODO(), &sns.PublishInput{
		Message:  &msg,
		TopicArn: topicArn,
	})
	if err != nil {
		log.Printf("An error occured when sending the SNS notification: %v\n", err)
	}

	log.Printf("Alerted. Message ID %s\n", *output.MessageId)
}
