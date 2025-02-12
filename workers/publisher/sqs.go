package publisher

import (
	"context"
	"fmt"
	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/sqs"
)

type Option func(*sqs.Config)

func MustGetStandardQueuePublisher(ctx context.Context, qn string, opts ...Option) *sqs.Publisher {
	return sqs.NewPublisher(MustGetStandardQueue(ctx, qn, opts...))
}

func GetStandardQueuePublisher(ctx context.Context, qn string, opts ...Option) (pub *sqs.Publisher, err error) {
	queue, err := GetStandardQueue(ctx, qn, opts...)
	if err != nil {
		return
	}

	pub = sqs.NewPublisher(queue)

	return
}

func MustGetFifoQueuePublisher(ctx context.Context, qn string, opts ...Option) *sqs.Publisher {
	return sqs.NewPublisher(MustGetFifoQueue(ctx, qn, opts...))
}

func GetFifoQueuePublisher(ctx context.Context, qn string, opts ...Option) (pub *sqs.Publisher, err error) {
	queue, err := GetFifoQueue(ctx, qn, opts...)
	if err != nil {
		return
	}

	pub = sqs.NewPublisher(queue)

	return
}

func GetStandardQueue(ctx context.Context, qn string, opts ...Option) (queue *sqs.Queue, err error) {
	awsConfig := getAwsConfig(ctx, opts...)
	queue, err = sqs.NewStandardQueue(ctx, qn, awsConfig)
	if err != nil {
		return
	}

	return
}

func MustGetStandardQueue(ctx context.Context, qn string, opts ...Option) *sqs.Queue {
	awsConfig := getAwsConfig(ctx, opts...)
	queue, err := sqs.NewStandardQueue(ctx, qn, awsConfig)
	if err != nil {
		panic(fmt.Sprintf("error while getting queue, err: %s", err.Error()))
	}

	return queue
}

func GetFifoQueue(ctx context.Context, qn string, opts ...Option) (queue *sqs.Queue, err error) {
	awsConfig := getAwsConfig(ctx, opts...)
	queue, err = sqs.NewFifoQueue(ctx, qn, awsConfig)
	if err != nil {
		return
	}

	return
}

func MustGetFifoQueue(ctx context.Context, qn string, opts ...Option) *sqs.Queue {
	awsConfig := getAwsConfig(ctx, opts...)
	queue, err := sqs.NewFifoQueue(ctx, qn, awsConfig)
	if err != nil {
		panic(fmt.Sprintf("error while getting queue, err: %s", err.Error()))
	}

	return queue
}

func getAwsConfig(ctx context.Context, opts ...Option) *sqs.Config {
	c := &sqs.Config{
		Account:  config.GetString(ctx, "aws.sqs.account"),
		Region:   config.GetString(ctx, "aws.sqs.region"),
		Endpoint: config.GetString(ctx, "aws.endpoint"),
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithAccount(account string) func(*sqs.Config) {
	return func(c *sqs.Config) {
		c.Account = account
	}
}

func WithRegion(region string) func(*sqs.Config) {
	return func(c *sqs.Config) {
		c.Region = region
	}
}

func WithEndpoint(endpoint string) func(*sqs.Config) {
	return func(c *sqs.Config) {
		c.Endpoint = endpoint
	}
}
