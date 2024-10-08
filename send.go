package crmhazar_pkg_push

import (
	"context"
	"firebase.google.com/go/messaging"
	"log"
)

func (c *Client) SendPush(ctx context.Context, title, body, image, page, action, token string) error {
	message := &messaging.Message{
		Data: map[string]string{
			"score":  "850",
			"time":   "2:45",
			"title":  title,
			"body":   body,
			"image":  image,
			"page":   page,
			"action": action,
		},
		Token: token,
	}

	response, err := c.FcmClient.Send(ctx, message)
	if err != nil {
		log.Println("FcmClient.Send ", err.Error())
		return err
	}
	log.Println("Response success => ", response)
	return nil
}

func (c *Client) SendMultiPush(ctx context.Context, title, body, image, page, action string, tokens []string) error {
	message := &messaging.MulticastMessage{
		Data: map[string]string{
			"score":  "850",
			"time":   "2:45",
			"title":  title,
			"body":   body,
			"image":  image,
			"page":   page,
			"action": action,
		},
		Tokens: tokens,
	}

	response, err := c.FcmClient.SendMulticast(ctx, message)
	if err != nil {
		log.Println("FcmClient.SendMulticast ", err.Error())
	}
	if response != nil {
		log.Println("Response success count : ", response.SuccessCount)
		log.Println("Response failure count : ", response.FailureCount)
	}

	return nil
}
