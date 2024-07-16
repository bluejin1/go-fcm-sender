package sender

import (
	"errors"
	"fcm-sender/cmd/fcm"
	"fcm-sender/helper"
	"fcm-sender/helper/zlog"
	"fcm-sender/internal/types"
	"fcm-sender/pkg/rdb"
	"fcm-sender/pkg/rdb/rdb_master/models"
	"strconv"
)

// 레시피 푸시
func SenderOrderRecipeAlarmToPartner(getAlarmData *models.PartnerAlarm, orderInfo *models.AppOrder) error {

	zlog.Info().Msgf("SenderOrderAlarmToStore : %s", getAlarmData)
	if getAlarmData == nil {
		zlog.Error().Msgf("alarmData nil : %s", getAlarmData)
		return errors.New("SenderOrderAlarmToStore alarmData nil")
	}

	alarmData := *getAlarmData
	zlog.Info().Msgf("alarmData : %s", alarmData)

	storeIdx := getAlarmData.StoreIdx
	storeMemberTokenList := rdb.MasterModel.GetLogErrorListWithEndpoint(&storeIdx, 100)
	zlog.Info().Msgf("storeMemberTokenList : %s", storeMemberTokenList)
	helper.PrettyPrint(storeMemberTokenList)
	if storeMemberTokenList == nil {
		zlog.Error().Msgf("storeMemberTokenList nil : %s", storeMemberTokenList)
		return errors.New("GetLogErrorListWithEndpoint storeMemberTokenList nil")
	}

	allowStoreAlarmRecipe := false
	if alarmData.StoreIdx > 0 {
		storeInfo := rdb.MasterModel.GetStoreInfo(alarmData.StoreIdx)
		zlog.Info().Msgf("storeInfo : %s", storeInfo)
		helper.PrettyPrint(storeInfo)
		if storeInfo != nil {
			if storeInfo.StoreAlarmRecipe == 1 {
				allowStoreAlarmRecipe = true
			}
		}
	}
	if allowStoreAlarmRecipe != true {
		zlog.Error().Msgf("allowStoreAlarmRecipe false : %s", allowStoreAlarmRecipe)
		return errors.New("store allowStoreAlarmRecipe false")
	}

	storeNo := strconv.Itoa(int(alarmData.StoreIdx))

	title := alarmData.AlarmTitle
	body := alarmData.AlarmContents
	pushData := &types.PartnerFcmPushData{
		FcmType:          "topic",
		Category:         alarmData.AlarmCategory,
		Title:            title,
		Body:             body,
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
		"type": "push",
		//"Content":      alarmData.AlarmContents,
		"category": alarmData.AlarmCategory,
		"title":    title,
		"body":     body,
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

	fcmTokens := []string{}
	tokenList := *storeMemberTokenList
	for i := range tokenList {
		memberToken := tokenList[i].AppUuid
		if len(memberToken) < 1 {
			continue
		}
		if tokenList[i].AllowRecipePush != 1 {
			continue
		}
		fcmTokens = append(fcmTokens, memberToken)
	}
	zlog.Info().Msgf("fcmTokens : %s", fcmTokens)
	helper.PrettyPrint(fcmTokens)

	if len(fcmTokens) > 0 {
		failedTokens, err := fcm.SendNotifications(title, body, fcmTokens, pushData)
		if len(failedTokens) > 0 {
			zlog.Error().Msgf("failedTokens list : %s", failedTokens)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

// 레시피 푸시 - kicc order
func SenderKiccOrderRecipeAlarmToPartner(getAlarmData *models.PartnerAlarm) error {

	zlog.Info().Msgf("SenderKiccOrderRecipeAlarmToPartner : %s", getAlarmData)
	if getAlarmData == nil {
		zlog.Error().Msgf("alarmData nil : %s", getAlarmData)
		return errors.New("SenderOrderAlarmToStore alarmData nil")
	}

	alarmData := *getAlarmData
	zlog.Info().Msgf("alarmData : %s", alarmData)

	storeIdx := getAlarmData.StoreIdx
	storeMemberTokenList := rdb.MasterModel.GetLogErrorListWithEndpoint(&storeIdx, 100)
	zlog.Info().Msgf("storeMemberTokenList : %s", storeMemberTokenList)
	helper.PrettyPrint(storeMemberTokenList)
	if storeMemberTokenList == nil {
		zlog.Error().Msgf("storeMemberTokenList nil : %s", storeMemberTokenList)
		return errors.New("GetLogErrorListWithEndpoint storeMemberTokenList nil")
	}

	allowStoreAlarmRecipe := false
	if alarmData.StoreIdx > 0 {
		storeInfo := rdb.MasterModel.GetStoreInfo(alarmData.StoreIdx)
		zlog.Info().Msgf("storeInfo : %s", storeInfo)
		helper.PrettyPrint(storeInfo)
		if storeInfo != nil {
			if storeInfo.StoreAlarmRecipe == 1 {
				allowStoreAlarmRecipe = true
			}
		}
	}
	if allowStoreAlarmRecipe != true {
		zlog.Error().Msgf("allowStoreAlarmRecipe false : %s", allowStoreAlarmRecipe)
		return errors.New("store allowStoreAlarmRecipe false")
	}

	storeNo := strconv.Itoa(int(alarmData.StoreIdx))

	title := alarmData.AlarmTitle
	body := alarmData.AlarmContents
	pushData := &types.PartnerFcmPushData{
		FcmType:          "push",
		Category:         alarmData.AlarmCategory,
		Title:            title,
		Body:             body,
		StoreNo:          storeNo,
		OrderNumber:      alarmData.OrderNo,
		OrderStatus:      "6",
		OrderReceiveType: "",
		ChangeKey:        "",
		ChangeVal:        "",
		AddBody:          "",
		IconType:         "",
		UserIdx:          strconv.Itoa(int(alarmData.PartnerIdx)),
		Time:             alarmData.CreateDate,
	}
	zlog.Info().Msgf("pushData : %s", pushData)

	fcmTokens := []string{}
	tokenList := *storeMemberTokenList
	for i := range tokenList {
		memberToken := tokenList[i].AppUuid
		if len(memberToken) < 1 {
			continue
		}
		if tokenList[i].AllowRecipePush != 1 {
			continue
		}
		fcmTokens = append(fcmTokens, memberToken)
	}
	zlog.Info().Msgf("fcmTokens : %s", fcmTokens)
	helper.PrettyPrint(fcmTokens)

	if len(fcmTokens) > 0 {
		failedTokens, err := fcm.SendNotifications(title, body, fcmTokens, pushData)
		if len(failedTokens) > 0 {
			zlog.Error().Msgf("failedTokens list : %s", failedTokens)
		}
		if err != nil {
			return err
		}
	}

	return nil
}
