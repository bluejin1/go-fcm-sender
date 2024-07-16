package rdb

import (
	"errors"
	"fcm-sender/configs"
	"fcm-sender/helper/zlog"
	"fcm-sender/pkg/rdb/rdb_config"
	"fcm-sender/pkg/rdb/rdb_lib"
	"fcm-sender/pkg/rdb/rdb_master"
	"fcm-sender/pkg/rdb/rdb_master/dao"
)

var (
	//rdbMasterConfig *rdb_config.RdbServerMaster
	rdbConnMaster   *rdb_master.MasterDB = nil
	MasterRdbLoaded bool                 = false
)

var (
	MasterModel     *dao.GormMasterModel     = nil
	rdbMasterConfig *configs.RdbCommonConfig = nil
	rdbLogConfig    *configs.RdbCommonConfig = nil
)

func InitRdb() {
	if configs.IsUseRdbMasterDatabase() {
		rdbMasterConfig = configs.DefaultMasterRdbServerConfigFromEnv()
		//zlog.Printf("init store_rdb master env : %v", rdbMasterConfig)
		if rdbMasterConfig != nil {
			configData := *rdbMasterConfig
			rConfig, err := rdb_master.InitRdbMaster(&configData.Host, &configData.Port, &configData.User, &configData.Password)
			if err != nil {
				zlog.Error().Msgf("rdb_master.InitRdbMaster err %v", err)
			}
			if rConfig == nil {
				zlog.Error().Msgf("rdb_master.InitRdbMaster rConfig err %v", rConfig)
			}
		} else {
			zlog.Error().Msgf("init store_rdb master env nil")
		}

	}

}

func ConnectMaster() (mdb *rdb_master.MasterDB, err error) {
	if rdb_config.IsUseRdbMasterDatabase() {
		if rdbMasterConfig != nil {
			rdbMasterConfig = configs.DefaultMasterRdbServerConfigFromEnv()
		}
		if rdbMasterConfig == nil {
			zlog.Error().Msgf("rdbMasterConfig nil")
			return mdb, errors.New("rdbMasterConfig nil")
		}
		configData := *rdbMasterConfig

		var cErr error = nil
		rdbConnMaster, cErr = rdb_lib.ConnectMasterDatabases(&configData.Host, &configData.Port, &configData.User, &configData.Password)
		if cErr != nil {
			zlog.Error().Msgf("ConnectMasterDatabases err %v", cErr)
			return mdb, cErr
		}
		MasterRdbLoaded = true
		return rdbConnMaster, nil

	} else {
		MasterRdbLoaded = false
		return mdb, errors.New("master database not use")
	}
}

func LoadMasterGormModel() *dao.GormMasterModel {
	MasterModel = dao.NewMasterModel()
	return MasterModel
}

func MasterDbClose() {
	if rdbConnMaster != nil {
		rdbConnMaster.Close()
	}
}
