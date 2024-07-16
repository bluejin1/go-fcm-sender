package fcm

import (
	"context"
	"fcm-sender/configs"
	"firebase.google.com/go/v4/messaging"
	"fmt"
)

// Firebase Cloud Messaging 단일 전송
func SendSingleNotification(title string, body string, fcmToken string) {

	// This registration token comes from the client FCM SDKs.
	registrationToken := fcmToken

	notification := &messaging.Notification{
		Title: title,
		Body:  body,
	}

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"Type":    title,
			"Content": body,
		},
		Token:        registrationToken,
		Notification: notification,
	}

	response, err := configs.FcmClient.Send(context.Background(), message)
	if err != nil {
		fmt.Println("[sendNotification] client.Send error :", err, registrationToken)
	}
	_ = response
}
