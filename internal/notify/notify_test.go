package notify

import (
	"context"
	"strconv"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type mockPublishAPI func(ctx context.Context, params *sns.PublishInput, optFns ...func(*sns.Options)) (*sns.PublishOutput, error)

func (m mockPublishAPI) Publish(ctx context.Context, params *sns.PublishInput, optFns ...func(*sns.Options)) (*sns.PublishOutput, error) {
	return m(ctx, params, optFns...)
}

func TestPublishMessage(t *testing.T) {
	cases := []struct {
		client        func(t *testing.T) snsPublishAPI
		message       string
		topicArn      string
		expectedMsgID string
	}{
		{
			client: func(t *testing.T) snsPublishAPI {
				return mockPublishAPI(
					func(ctx context.Context, params *sns.PublishInput, optFns ...func(*sns.Options)) (*sns.PublishOutput, error) {
						t.Helper()
						if params.Message == nil {
							t.Fatal("message cannot be nil")
						}
						if e, a := "testTopic", *params.TopicArn; e != a {
							t.Errorf("expect %v, got %v", e, a)
						}

						messageId := "msg123"
						return &sns.PublishOutput{
							MessageId: &messageId,
						}, nil
					},
				)
			},
			message:  "test",
			topicArn: "testTopic",
		},
	}

	for i, tt := range cases {
		t.Run(
			strconv.Itoa(i), func(t *testing.T) {
				ctx := context.TODO()
				output, err := publishMessage(ctx, tt.client(t), &tt.message, &tt.topicArn)
				if err != nil {
					t.Fatalf("expect no error, got %v", err)
				}
				if *output.MessageId != "msg123" {
					t.Fatalf("expect message ID msg123, got %v", *output.MessageId)
				}
			},
		)
	}
}
