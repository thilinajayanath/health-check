package notify

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type snsPublishAPI interface {
	Publish(ctx context.Context, params *sns.PublishInput, optFns ...func(*sns.Options)) (*sns.PublishOutput, error)
}

func publishMessage(ctx context.Context, api snsPublishAPI, message, topicArn *string) (*sns.PublishOutput, error) {
	output, err := api.Publish(
		ctx,
		&sns.PublishInput{
			Message:  message,
			TopicArn: topicArn,
		},
	)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func Notify(msg string, topicArn *string) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("An error occured when loading AWS configurations: %v\n", err)
		return
	}

	client := sns.NewFromConfig(cfg)

	msgOutput, err := publishMessage(
		context.TODO(),
		client,
		&msg,
		topicArn,
	)
	if err != nil {
		log.Printf("An error occured when sending the SNS notification: %v\n", err)
		return
	}

	log.Printf("Alerted. Message ID %s\n", *msgOutput.MessageId)
}
