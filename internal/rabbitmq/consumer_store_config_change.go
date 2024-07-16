package rabbitmq

import (
	"encoding/json"
	"errors"
	"fcm-sender/helper"
	"fcm-sender/helper/zlog"
	"fcm-sender/internal/sender"
	"fcm-sender/internal/types"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func NewStoreConfigChangeInfoReceiver(mqMessage *amqp.Delivery) error {

	log.Printf("NewStoreConfigChangeInfoReceiver %C", mqMessage)

	log.Printf("Received a message: %s", mqMessage.Body)

	configInfo := types.StoreChangeConfig{}
	err := json.Unmarshal([]byte(mqMessage.Body), &configInfo)
	if err != nil {
		zlog.Error().Msgf("Error %s", err)
		return err
	}
	zlog.Debug().Msgf("configInfo: %+v", configInfo)
	helper.PrettyPrint(configInfo)

	if len(configInfo.ConfigKey) < 1 || len(configInfo.ConfigVal) < 1 {
		zlog.Error().Msgf("ConfigKey empty %s", configInfo)
		return errors.New("ConfigKey empty")
	}

	// FCM 전송
	senderErr := sender.SenderStoreConfigChangeAlarmToPartner(&configInfo)
	if senderErr != nil {
		zlog.Error().Msgf("SenderStoreConfigChangeAlarmToPartner Error %s", senderErr)
		return senderErr
	}

	return nil
}
