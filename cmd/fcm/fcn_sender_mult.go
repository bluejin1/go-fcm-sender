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

// Firebase Cloud Messaging 다중 전송
func SendNotifications(title string, body string, fcmTokens []string, pushData *types.PartnerFcmPushData) (failedTokens []string, err error) {

	fmt.Println("SendNotifications:")

	if len(fcmTokens) == 0 {
		return failedTokens, errors.New("fcmTokens is empty")
	}
	if pushData == nil {
		return failedTokens, errors.New("pushData is empty")
	}
	Data := make(map[string]string)
	for k, v := range structs.Map(pushData) {
		Data[k] = v.(string)
	}
	if len(Data) < 1 {
		return failedTokens, errors.New("data is empty")
	}

	// fcm tokens
	registrationTokens := fcmTokens

	notification := &messaging.Notification{
		Title: title,
		Body:  body,
	}

	message := &messaging.MulticastMessage{
		Data:         Data,
		Tokens:       registrationTokens,
		Notification: notification,
	}

	if configs.FcmClient == nil {
		fcmErr := configs.FcmInit()
		if fcmErr != nil {
			fmt.Printf("FcmInit err: %s\n", fcmErr)
			return failedTokens, fcmErr
		}
	}

	br, err := configs.FcmClient.SendMulticast(context.Background(), message)
	if err != nil {
		fmt.Printf("configs.FcmClient.SendMulticast err: %s\n", err)
		return failedTokens, err
	}

	fmt.Println("[sendNotifications] ", br.SuccessCount, " messages were sent successfully, ", br.FailureCount, " messages fail count")

	if br.FailureCount > 0 {
		//var failedTokens []string
		for idx, resp := range br.Responses {
			if !resp.Success {
				// The order of responses corresponds to the order of the registration tokens.
				failedTokens = append(failedTokens, registrationTokens[idx])
			}
		}

		fmt.Printf("[sendNotifications] List of tokens that caused failures: %v\n", failedTokens)
	}
	return failedTokens, nil
}
