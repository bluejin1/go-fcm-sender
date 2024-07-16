package lib

import (
	"errors"
	"fcm-sender/helper"
	"fcm-sender/helper/zlog"
	"fcm-sender/pkg/rdb"
	"fcm-sender/pkg/rdb/rdb_master/models"
)

var (
	StoreInfoListKicc *map[string]models.StoreInfoShort = nil // 저장된 정보
	StoreInfoList     *map[int32]models.StoreInfoShort  = nil // 저장된 정보
	StoreInfoTime     int64                             = 0   // 갱신 시간
	StoreInfoKiccTime int64                             = 0   // 갱신 시간
)

func SetStoreListMemory() bool {

	storeListData := rdb.MasterModel.GetStoreInfoListAllShortWithKicc(5000)
	//zlog.Info().Msgf("storeListData : %s", storeListData)
	if storeListData != nil {
		//zlog.Info().Msgf("storeListData : %s", *storeListData)
		//helper.PrettyPrint(*storeListData)
		getStoreListData := *storeListData
		storeListKicc := map[string]models.StoreInfoShort{}
		StoreList := map[int32]models.StoreInfoShort{}
		if len(getStoreListData) > 0 {
			for _, data := range getStoreListData {
				key := data.StoreIdx
				code := data.KiccStoreCode
				StoreList[key] = data
				storeListKicc[code] = data
			}
		}
		if len(getStoreListData) > 0 {
			StoreInfoList = &StoreList
			StoreInfoTime = helper.GetNowTimestampInt64()
		}
		if len(storeListKicc) > 0 {
			StoreInfoListKicc = &storeListKicc
			StoreInfoKiccTime = helper.GetNowTimestampInt64()
		}
		zlog.Info().Msgf("StoreInfoTime : %s", StoreInfoTime)
		zlog.Info().Msgf("StoreInfoKiccTime : %s", StoreInfoKiccTime)

	} else {
		zlog.Error().Msgf("storeListData load fail %s", storeListData)
		return false
	}

	return true
}

func GetStoreInfoWithIdx(storeIdx int32) (res *models.StoreInfoShort, err error) {

	zlog.Info().Msgf("GetStoreInfoWithIdx : %s", storeIdx)

	if storeIdx < 1 {
		return nil, errors.New("storeIdx nil")
	}

	isLoaded := false
	checkTime := StoreInfoTime + 3600
	nowTime := helper.GetNowTimestampInt64()
	zlog.Info().Msgf("nowTime : %s", nowTime)
	zlog.Info().Msgf("checkTime : %s", checkTime)
	if nowTime > checkTime {
		isLoaded = SetStoreListMemory()
	}
	if StoreInfoList != nil {
		isLoaded = true
	}
	zlog.Info().Msgf("isLoaded : %s", isLoaded)

	if isLoaded {
		getStoreInfoList := *StoreInfoList
		storeInfoModel, ok := getStoreInfoList[storeIdx]
		zlog.Info().Msgf("storeInfoModel : %s", storeInfoModel)
		if ok {
			return &storeInfoModel, nil
		} else {
			storeInfo := rdb.MasterModel.GetStoreInfoShort(storeIdx)
			zlog.Info().Msgf("storeInfo : %c", storeInfo)
			helper.PrettyPrint(&storeInfo)
			if storeInfo != nil {
				return storeInfo, nil
			}
			return res, errors.New("storeInfo is nil")
		}

	} else {
		storeInfo := rdb.MasterModel.GetStoreInfoShort(storeIdx)
		zlog.Info().Msgf("storeInfo : %c", storeInfo)
		helper.PrettyPrint(&storeInfo)
		if storeInfo != nil {
			return storeInfo, nil
		}
		return res, errors.New("storeInfo is nil")
	}

	return res, nil

}

func GetStoreInfoWithCode(storeCode string) (res *models.StoreInfoShort, err error) {

	zlog.Info().Msgf("GetStoreInfoWithCode : %s", storeCode)

	if len(storeCode) < 1 {
		return nil, errors.New("storeCode nil")
	}

	isLoaded := false
	checkTime := StoreInfoTime + 3600
	nowTime := helper.GetNowTimestampInt64()
	zlog.Info().Msgf("nowTime : %s", nowTime)
	zlog.Info().Msgf("checkTime : %s", checkTime)
	if nowTime > checkTime {
		isLoaded = SetStoreListMemory()
	}
	if StoreInfoList != nil {
		isLoaded = true
	}
	zlog.Info().Msgf("isLoaded : %s", isLoaded)

	if isLoaded {
		getStoreInfoList := *StoreInfoListKicc
		storeInfoModel, ok := getStoreInfoList[storeCode]
		zlog.Info().Msgf("storeInfoModel : %s", storeInfoModel)
		if ok {
			return &storeInfoModel, nil
		} else {
			storeInfo := rdb.MasterModel.GetStoreInfoShortWithCode(storeCode)
			zlog.Info().Msgf("storeInfo : %c", storeInfo)
			helper.PrettyPrint(&storeInfo)
			if storeInfo != nil {
				return storeInfo, nil
			}
			return res, errors.New("storeInfo is nil")
		}

	} else {
		storeInfo := rdb.MasterModel.GetStoreInfoShortWithCode(storeCode)
		zlog.Info().Msgf("storeInfo : %c", storeInfo)
		helper.PrettyPrint(&storeInfo)
		if storeInfo != nil {
			return storeInfo, nil
		}
		return res, errors.New("storeInfo is nil")
	}

	return res, nil

}
