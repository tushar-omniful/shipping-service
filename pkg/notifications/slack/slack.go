package slack

import (
	"context"
	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/log"
	"github.com/slack-go/slack"
)

var slackClient *slack.Client

func Get() *slack.Client {
	return slackClient
}

func Set(ctx context.Context) {
	slackClient = slack.New(config.GetString(ctx, "notification.slack.token"))
}

func SendNotification(ctx context.Context, message string) {
	_, _, err := Get().PostMessage(
		config.GetString(ctx, "notification.slack.channelID"),
		slack.MsgOptionText(message, false),
	)

	if err != nil {
		log.Errorf("unable to send slack notification for bulk schedule layout configuration. err :: %v", err)
	}

}
