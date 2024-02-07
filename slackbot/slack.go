package slackbot

import (
	"time"

	"github.com/slack-go/slack"
)

type SlackBot struct {
	client *slack.Client
}

func New(slackToken string) SlackBot {
	result := SlackBot{
		client: slack.New(slackToken, slack.OptionDebug(true)),
	}
	return result
}

func (s *SlackBot) Info(channelID string, message string) error {
	// Create the Slack attachment that we will send to the channel
	attachment := slack.Attachment{
		Pretext: "Pretext?",
		Text:    message,
		// Color Styles the Text, making it possible to have like Warnings etc.
		Color: "#36a64f",
		// Fields are Optional extra data!
		Fields: []slack.AttachmentField{
			{
				Title: "Date",
				Value: time.Now().String(),
			},
		},
	}

	_, _, err := s.client.PostMessage(
		channelID,
		// uncomment the item below to add a extra Header to the message
		//slack.MsgOptionText("New message from bot", false),
		slack.MsgOptionAttachments(attachment),
	)
	return err
}
