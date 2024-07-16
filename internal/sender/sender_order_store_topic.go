package sender

import (
	"fcm-sender/cmd/fcm"
	"fcm-sender/helper/env"
	"fcm-sender/helper/zlog"
	"fcm-sender/internal/types"
	"fcm-sender/pkg/rdb/rdb_master/models"

	"errors"
	"fmt"
	"strconv"
)

func SenderOrderAlarmToStore(getAlarmData *models.PartnerAlarm, orderInfo *models.AppOrder) error {

	zlog.Info().Msgf("SenderOrderAlarmToStore : %s", getAlarmData)
	if getAlarmData == nil {
		zlog.Error().Msgf("alarmData nil : %s", getAlarmData)
		return errors.New("SenderOrderAlarmToStore alarmData nil")
	}

	alarmData := *getAlarmData
	zlog.Info().Msgf("alarmData : %s", alarmData)

	storeNo := strconv.Itoa(int(alarmData.StoreIdx))
	topicName := "topic-store-" + storeNo
	if env.GetEnvEnvironment() == "development" {
		topicName = "test-topic-store-" + storeNo
	}

	title := alarmData.AlarmTitle
	pushData := &types.PartnerFcmPushData{
		FcmType:          "topic",
		Category:         alarmData.AlarmCategory,
		Title:            title,
		Body:             alarmData.AlarmContents,
		StoreNo:          storeNo,
		OrderNumber:      alarmData.OrderNo,
		OrderStatus:      strconv.Itoa(int(orderInfo.OrderStatus)),
		OrderReceiveType: orderInfo.ReceiveType,
		ChangeKey:        "",
		ChangeVal:        "",
		AddBody:          "",
		IconType:         "",
		UserIdx:          strconv.Itoa(int(alarmData.PartnerIdx)),
		Time:             alarmData.CreateDate,
	}
	/*pushData := map[string]string{
		"type": "topic",
		//"Content":      alarmData.AlarmContents,
		"category": alarmData.AlarmCategory,
		"title":    title,
		"body":     alarmData.AlarmContents,
		//"location":     alarmData.AlarmLocation,
		"store_no":     storeNo,
		"order_number": alarmData.OrderNo,
		"order_status": strconv.Itoa(int(orderInfo.OrderStatus)),
		//"order_title":  orderInfo.OrderTitle,
		//"order_name":   orderInfo.OrdererName,
		"change_key": "",
		"change_val": "",
		"add_body":   "",
		"icon_type":  "",
		"userIdx":    strconv.Itoa(int(alarmData.PartnerIdx)),
		//"score":        "10000",
		"time": alarmData.CreateDate,
	}*/
	zlog.Info().Msgf("pushData : %s", pushData)

	_, err := fcm.SendNotificationTopic(pushData, topicName)
	if err != nil {
		//log.Fatalln(err)
		fmt.Println("[SendNotificationTopic] client.Send error :", err, topicName)
	}

	return nil
}
