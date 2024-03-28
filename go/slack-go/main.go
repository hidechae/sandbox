package main

import (
	"fmt"
	"os"
	"time"

	"github.com/slack-go/slack"
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	api := slack.New(token, slack.OptionDebug(true))

	// メッセージの送信処理
	_, _, err := api.PostMessage(
		"#notification-test",
		slack.MsgOptionText("This is send test. Current time is "+time.Now().Format(time.DateTime), true),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}
