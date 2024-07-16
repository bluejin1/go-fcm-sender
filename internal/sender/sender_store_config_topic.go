package sender

import (
	"errors"
	"fcm-sender/cmd/fcm"
	"fcm-sender/helper"
	"fcm-sender/helper/env"
	"fcm-sender/helper/zlog"
	"fcm-sender/internal/types"
	"fmt"
	"strconv"
)

func SenderStoreConfigChangeAlarmToPartner(configData *types.StoreChangeConfig) error {

	zlog.Info().Msgf("SenderStoreConfigChangeAlarmToPartner : %s", configData)
	if configData == nil {
		zlog.Error().Msgf("configData nil : %s", configData)
		return errors.New("SenderStoreConfigChangeAlarmToPartner configData nil")
	}

	alarmData := *configData
	zlog.Info().Msgf("alarmData : %s", alarmData)

	storeNo := strconv.Itoa(int(alarmData.StoreIdx))
	topicName := "topic-store-" + storeNo
	if env.GetEnvEnvironment() == "development" {
		topicName = "test-topic-store-" + storeNo
	}

	title := alarmData.PushTitle
	pushData := &types.PartnerFcmPushData{
		FcmType:          "topic",
		Category:         alarmData.PushCategory,
		Title:            title,
		Body:             alarmData.PushBody,
		StoreNo:          storeNo,
		OrderNumber:      "",
		OrderStatus:      "",
		OrderReceiveType: "",
		ChangeKey:        alarmData.ConfigKey,
		ChangeVal:        alarmData.ConfigVal,
		AddBody:          "",
		IconType:         "",
		UserIdx:          "",
		Time:             alarmData.ModifyDate,
	}
	/*pushData := map[string]string{
		"type": "topic",
		//"Content":      alarmData.AlarmContents,
		"category": alarmData.PushCategory,
		"title":    title,
		"body":     alarmData.PushBody,
		//"location":     alarmData.AlarmLocation,
		"store_no":     storeNo,
		"order_number": "",
		"order_status": "",
		"change_key":   alarmData.ConfigKey,
		"change_val":   alarmData.ConfigVal,
		//"order_title":  "",
		//"order_name":   "",
		"add_body":  "",
		"icon_type": "",
		"userIdx":   "",
		//"score":        "10000",
		"time": alarmData.ModifyDate,
	}*/
	zlog.Info().Msgf("pushData : %s", pushData)
	helper.PrettyPrint(pushData)

	_, err := fcm.SendNotificationTopic(pushData, topicName)
	if err != nil {
		//log.Fatalln(err)
		fmt.Println("[SendNotificationTopic] client.Send error :", err, topicName)
	}

	return nil
}
