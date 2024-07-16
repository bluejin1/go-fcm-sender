package fcm

import (
	"context"
	"errors"
	"fcm-sender/configs"
	"fcm-sender/internal/types"
	"firebase.google.com/go/v4/messaging"
	"fmt"
	"github.com/fatih/structs"
)

func SendNotificationTopic(pushData *types.PartnerFcmPushData, topic string) (response string, err error) {

	fmt.Println("SendNotificationTopic:", topic)
	// This registration token comes from the client FCM SDKs.
	// The topic name can be optionally prefixed with "/topics/".
	//topic := "highScores"

	if pushData == nil {
		return response, errors.New("pushData is empty")
	}
	Data := make(map[string]string)
	for k, v := range structs.Map(pushData) {
		Data[k] = v.(string)
	}
	if len(Data) < 1 {
		return response, errors.New("data is empty")
	}

	notification := &messaging.Notification{
		Title: Data["title"],
		Body:  Data["body"],
	}

	// See documentation on defining a message payload.
	message := &messaging.Message{
		/*Data: map[string]string{
			"category": topic,
			"Type":     title,
			"Content":  body + topic,
			"score":    "850",
			"time":     "2:45",
		},*/
		Data:         Data,
		Topic:        topic,
		Notification: notification,
	}

	if configs.FcmClient == nil {
		fcmErr := configs.FcmInit()
		if fcmErr != nil {
			fmt.Printf("FcmInit err: %s\n", fcmErr)
			return response, fcmErr
		}
	}

	// Send a message to the devices subscribed to the provided topic.
	response, err = configs.FcmClient.Send(context.Background(), message)
	if err != nil {
		//log.Fatalln(err)
		fmt.Println("[sendNotification] client.Send error :", err, topic)
		return response, err
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)

	return response, err
}
