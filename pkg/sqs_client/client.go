package sqs_client

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQSClient struct {
	sqs *sqs.SQS
}

func NewSQSClient(region string) (*SQSClient, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, err
	}

	return &SQSClient{
		sqs: sqs.New(sess),
	}, nil
}

func (c *SQSClient) SendMessage(queueURL string, messageBody string) error {
	input := &sqs.SendMessageInput{
		MessageBody: aws.String(messageBody),
		QueueUrl:    aws.String(queueURL),
	}

	_, err := c.sqs.SendMessage(input)
	return err
}
