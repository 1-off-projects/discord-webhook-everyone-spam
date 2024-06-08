package main

import (
	"fmt"
	"time"

	"github.com/DisgoOrg/disgohook"
)

func main() {
	webhookID := "hook id"
	webhookToken := "hook token"
	webhook, err := disgohook.NewWebhookClientByToken(nil, nil, webhookID+"/"+webhookToken)
	if err != nil {
		fmt.Println("Error creating webhook client:", err)
		return
	}
	_, err = webhook.SendContent("<@&1246560273746231356>")
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
	time.Sleep(1000000000)
	fmt.Println("pong")
}
