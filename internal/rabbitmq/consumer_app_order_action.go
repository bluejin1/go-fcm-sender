package rabbitmq

import (
	"encoding/json"
	"errors"
	"fcm-sender/helper"
	"fcm-sender/helper/zlog"
	"fcm-sender/internal/sender"
	"fcm-sender/pkg/rdb"
	"fcm-sender/pkg/rdb/rdb_master/lib"
	"fcm-sender/pkg/rdb/rdb_master/models"
	"fmt"
	"log"
	"strconv"
	"strings"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewAppOrderReceiver(mqMessage *amqp.Delivery) error {

	log.Printf("NewAppOrderReceiver %C", mqMessage)

	log.Printf("Received a message: %s", mqMessage.Body)

	//helper.PrettyPrint(mqMessage.Body)

	getOrderInfo := &models.AppOrder{}
	err := json.Unmarshal([]byte(mqMessage.Body), getOrderInfo)
	if err != nil {
		zlog.Error().Msgf("Error %s", err)
		return err
	}
	zlog.Debug().Msgf("getOrderInfo: %+v", *getOrderInfo)
	helper.PrettyPrint(getOrderInfo)

	/*orderInfo := getOrderInfo
	helper.PrettyPrint(orderInfo)
	zlog.Debug().Msgf("orderInfo.MemberNo : %s", orderInfo.MemberNo)*/

	orderNo := getOrderInfo.OrderNo
	if len(orderNo) == 0 {
		zlog.Error().Msgf("Store Idx Error : %s", orderNo)
		return errors.New("orderNo Error")
	}
	zlog.Debug().Msgf("orderNo : %s", orderNo)

	// 주문정보를 DB에서 가져온다.
	orderInfo := rdb.MasterModel.GetMasterAppOrder(&orderNo)
	if orderInfo == nil {
		zlog.Error().Msgf("orderInfo nil : %s", orderInfo)
		return errors.New("orderInfo nil Error")
	}

	helper.PrettyPrint(*orderInfo)
	zlog.Debug().Msgf("orderInfo.MemberNo : %s", orderInfo.MemberNo)

	productDataList := make([]models.AppOrderProduct, 0)
	productRaw := orderInfo.ProductData
	if len(orderInfo.ProductData) != 0 {
		fmt.Printf("productRaw %s\n", productRaw)
		err = json.Unmarshal([]byte(productRaw), &productDataList)
	}
	helper.PrettyPrint(productDataList)

	// store info
	storeIdx := orderInfo.StoreIdx
	if storeIdx < 1 {
		zlog.Error().Msgf("Store Idx Error : %d", storeIdx)
		return errors.New("Store Idx Error")
	}

	//storeInfo := rdb.MasterModel.GetStoreInfoFull(storeIdx)
	storeInfo, err := lib.GetStoreInfoWithIdx(storeIdx)
	if err != nil {
		zlog.Error().Msgf("storeInfo empty Error : %s", storeInfo)
		return err
	}
	zlog.Info().Msgf("storeInfo : %c", storeInfo)
	helper.PrettyPrint(&storeInfo)

	isOrderAlarm := false
	if storeInfo.StoreAlarmOrder == 1 {
		isOrderAlarm = true
	}
	isRecipeAlarm := false
	if storeInfo.StoreAlarmRecipe == 1 {
		isRecipeAlarm = true
	}

	zlog.Info().Msgf("isOrderAlarm : %c", isOrderAlarm)
	zlog.Info().Msgf("isRecipeAlarm : %c", isRecipeAlarm)

	if isOrderAlarm == false && isRecipeAlarm == false {
		return nil
	}
	zlog.Debug().Msgf("isRecipeAlarm : %c", isRecipeAlarm)

	productIdx := []int32{}
	// 레시피 정보
	if isRecipeAlarm == true {
		if len(productDataList) > 0 {

			for no, data := range productDataList {
				zlog.Debug().Msgf("product : %s", data.No)
				zlog.Debug().Msgf("no : %s", no)
				//no, _ = strconv.ParseInt(product.No, 10, 64)
				tmp, _ := strconv.ParseInt(data.No, 10, 64)
				zlog.Debug().Msgf("tmp : %s", tmp)
				//productIdx[no] = int32(tmp)
				productIdx = append(productIdx, int32(tmp))
			}
		}
	}
	helper.PrettyPrint(productIdx)
	zlog.Debug().Msgf("productIdx : %s", productIdx)

	// 주문에 있는 상품정보를 가져옴
	var productList []models.Product
	productListRaw := rdb.MasterModel.GetProductInfoListProductNos(productIdx)
	if productListRaw != nil {
		productList = *productListRaw
	} else {
		isRecipeAlarm = false
	}

	helper.PrettyPrint(productList)
	zlog.Debug().Msgf("productList : %s", productList)

	recipeDataList := make([]models.AlarmRecipeData, 0)
	recipeProductNames := []string{}
	recipeIdxList := []string{}
	if len(productList) > 0 && isRecipeAlarm == true {
		for _, product := range productList {
			if product.RecipeIdx > 0 {

				recipeData := models.AlarmRecipeData{
					Id:   strconv.Itoa(int(product.RecipeIdx)),
					Name: product.ProductName + "(ICED)",
					Type: "ICED",
					Icon: "",
				}

				if helper.InArray(recipeData.Id, recipeIdxList) == -1 {
					recipeDataList = append(recipeDataList, recipeData)
					recipeIdxList = append(recipeIdxList, recipeData.Id)
				}

			}
			if product.HotRecipeIdx > 0 {
				recipeData := models.AlarmRecipeData{
					Id:   strconv.Itoa(int(product.HotRecipeIdx)),
					Name: product.ProductName + "(HOT)",
					Type: "HOT",
					Icon: "",
				}
				if helper.InArray(recipeData.Id, recipeIdxList) == -1 {
					recipeDataList = append(recipeDataList, recipeData)
					recipeIdxList = append(recipeIdxList, recipeData.Id)
				}
			}
			recipeProductNames = append(recipeProductNames, product.ProductName)
		}
	}
	alarmRecipeData, _ := json.Marshal(recipeDataList)

	// 발송지: APP,KICC, POS, ADMIN
	appType := "APP"
	zlog.Info().Msgf("orderInfo.ProductData : %s", orderInfo.ProductData)

	orderProductDataList := make([]models.OrderProduct, 0)
	jErr := json.Unmarshal([]byte(orderInfo.ProductData), &orderProductDataList)
	if jErr != nil {
		zlog.Error().Msgf("orderProductDataList json Error : %s", jErr)
	}
	zlog.Info().Msgf("orderProductDataList : %s", orderProductDataList)
	helper.PrettyPrint(orderProductDataList)

	orderProductData, _ := json.Marshal(orderProductDataList)
	zlog.Info().Msgf("orderProductData : %s", orderProductData)
	helper.PrettyPrint(orderProductData)

	orderData := models.AlarmOrderData{
		AppType:  appType,
		Id:       orderNo,
		Name:     orderInfo.OrderTitle,
		Amount:   strconv.Itoa(int(orderInfo.OrderAmount)),
		Products: orderProductData,
	}
	helper.PrettyPrint(orderData)
	zlog.Info().Msgf("orderData : %s", orderData)

	alarmOrderData, _ := json.Marshal(orderData)

	alarmData := "[]"
	// 푸시발송여부 체크 - 0~90-미발송, 100 발송, 101-발송생략, 99-발송에러,102-GOFCM
	var pushCheck int8 = 102
	// 알림상태 - 1: 정상, 0:미노출, 9:삭제
	var alarmStatus int8 = 1
	// 알림카테고리 - NOTICE, QA, RECIPE
	alarmCategory := "ORDER"

	zlog.Info().Msgf("isOrderAlarm : %c", isOrderAlarm)
	zlog.Info().Msgf("isRecipeAlarm : %c", isRecipeAlarm)

	if isOrderAlarm == false && isRecipeAlarm == false {
		return nil
	}
	// test
	//isOrderAlarm = false
	//isRecipeAlarm = true

	alarmTitle := "주문이 접수되었습니다."
	alarmContents := "APP 오더로부터 주문이 접수되었습니다.(주문번호:" + orderNo + ")"
	if isOrderAlarm == false && isRecipeAlarm == true {
		if len(recipeIdxList) < 1 {
			zlog.Error().Msgf("recipeIdxList len zero %s", recipeIdxList)
			return errors.New("recipeIdxList len zero")
		}
		alarmTitle = "[레시피 도착] " + strings.Join(recipeProductNames, ", ")
		alarmContents = "신규 주문에 대한 레시피를 확인해 보세요.(주문번호:" + orderNo + ")"
		alarmCategory = "RECIPE"
	}

	partnerAlarmData := models.PartnerAlarm{
		PushCheck:     pushCheck,
		AlarmStatus:   alarmStatus,
		IsRead:        0,
		AlarmTitle:    alarmTitle,
		AlarmLocation: appType,
		AlarmCategory: alarmCategory,
		AlarmContents: alarmContents,
		AlarmData:     alarmData,
		PartnerIdx:    0,
		StoreIdx:      storeIdx,
		OrderNo:       orderNo,
		OrderData:     string(alarmOrderData),
		RecipeIdxs:    strings.Join(recipeIdxList, "|"),
		RecipeData:    string(alarmRecipeData),
		CreateUser:    "FCM:" + strconv.Itoa(int(orderInfo.MemberNo)),
		CreateDate:    helper.GetNowDbDatetime(),
		CreateTime:    helper.GetNowTimestampInt32(),
	}

	zlog.Debug().Msgf("partnerAlarmData : %s", partnerAlarmData)
	helper.PrettyPrint(partnerAlarmData)

	alarmIdx, resErr := rdb.MasterModel.CreatePartnerAlarmApp(&partnerAlarmData)
	if resErr != nil {
		zlog.Error().Msgf("Error %s", resErr)
		return resErr
	}
	zlog.Debug().Msgf("alarmIdx : %s", alarmIdx)

	// FCM 전송
	if isOrderAlarm {
		senderErr := sender.SenderOrderAlarmToStore(&partnerAlarmData, orderInfo)
		if senderErr != nil {
			zlog.Error().Msgf("SenderOrderAlarmToStore Error %s", senderErr)
			return senderErr
		}
	} else if isRecipeAlarm {
		if len(recipeIdxList) < 1 {
			zlog.Error().Msgf("recipeIdxList len zero %s", recipeIdxList)
			return errors.New("recipeIdxList len zero")
		}
		senderErr := sender.SenderOrderRecipeAlarmToPartner(&partnerAlarmData, orderInfo)
		if senderErr != nil {
			zlog.Error().Msgf("SenderOrderRecipeAlarmToPartner Error %s", senderErr)
			return senderErr
		}
	}

	return nil
}
