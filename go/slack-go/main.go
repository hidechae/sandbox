package main

import (
	"context"
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

var c *slack.Client

func main() {
	token := os.Getenv("SLACK_TOKEN")
	c = slack.New(token, slack.OptionDebug(true))
	ctx := context.Background()

	//if err := SendSimpleMessage(ctx, "This is send test. Current time is "+time.Now().Format(time.DateTime)); err != nil {
	//	fmt.Printf("%s\n", err)
	//}

	if err := SendBlockMessage(ctx, "Hello World", "This is body text."); err != nil {
		fmt.Printf("%s\n", err)
	}
}

func SendSimpleMessage(ctx context.Context, text string) error {
	_, _, err := c.PostMessageContext(ctx, "#notification-test", slack.MsgOptionText(text, true))
	return err
}

func SendBlockMessage(ctx context.Context, title, body string) error {
	_, _, err := c.PostMessageContext(ctx, "#notification-test", slack.MsgOptionBlocks(
		// 区切り線
		slack.NewDividerBlock(),

		// テキストのみのセクション
		&slack.SectionBlock{
			Type: slack.MBTSection,
			Text: &slack.TextBlockObject{
				Type: "mrkdwn",
				Text: "これはタイトルです",
			},
		},

		// テキストとフィールドのセクション
		slack.NewSectionBlock(
			&slack.TextBlockObject{
				Type: "mrkdwn",
				Text: fmt.Sprintf("*%s*", title),
			},
			[]*slack.TextBlockObject{
				{
					Type: "plain_text",
					Text: body,
				},
			},
			nil,
		),
	))
	return err
}
