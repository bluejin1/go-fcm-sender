package rabbitmq

import (
	"encoding/json"
	"fcm-sender/helper"
	"fcm-sender/helper/zlog"
	"fcm-sender/internal/sender"
	"fcm-sender/internal/types"
	"fcm-sender/pkg/rdb/rdb_master/models"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func NewKiccOrderReceiver(mqMessage *amqp.Delivery) error {

	log.Printf("NewAppOrderReceiver %C", mqMessage)

	log.Printf("Received a message: %s", mqMessage.Body)

	//helper.PrettyPrint(mqMessage.Body)

	orderInfo := &types.KiccOrderInfo{}
	err := json.Unmarshal([]byte(mqMessage.Body), orderInfo)

	if err != nil {
		zlog.Error().Msgf("Error %s", err)
		return err
	}
	zlog.Debug().Msgf("orderInfo: %+v", *orderInfo)
	helper.PrettyPrint(orderInfo)

	partnerAlarmData := models.PartnerAlarm{}

	// FCM 전송
	senderErr := sender.SenderKiccOrderRecipeAlarmToPartner(&partnerAlarmData)
	if senderErr != nil {
		zlog.Error().Msgf("SenderOrderRecipeAlarmToPartner Error %s", senderErr)
		return senderErr
	}

	return nil
}
