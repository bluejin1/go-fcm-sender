package rdb_config

import (
	"fcm-sender/configs"
	"sync"

	"fcm-sender/helper/env"
)

var (
	RdbMasterDatabaseNameLive string = configs.RdbMasterDatabaseNameLive
	RdbMasterDatabaseNameTest string = configs.RdbMasterDatabaseNameTest
)

var (
	EnvServerType *string = nil
)

var (
	RdbConfigMaster                 *RdbServerMaster
	rdbConfigOnce                   sync.Once
	RdbHost                         *string
	RdbPort                         *string
	RdbUser                         *string
	RdbPassword                     *string
	CodeRdbConfigDatabaseTypeMaster string = "master"
)

func GetRdbMasterDatabaseName() string {
	if env.GetEnvEnvironment() == "development" {
		return RdbMasterDatabaseNameTest
	}
	return RdbMasterDatabaseNameLive
}

func IsUseRdbMasterDatabase() bool {
	useDatabase := env.GetEnv("RDB_USE_MASTER_DB", "true")
	if useDatabase == "true" {
		return true
	}
	return false
}

func GetRdbClusterConnectionStr() string {
	clusterUrl := env.GetEnv("RDB_CLUSTER_URL", "")
	return clusterUrl
}

func GetRdbConnectionStr() string {
	var dbConnection string = ""
	if RdbConfigMaster == nil {
		return dbConnection
	}
	if RdbConfigMaster.User == "" || RdbConfigMaster.Password == "" || RdbConfigMaster.Address.Host == "" || RdbConfigMaster.Database == "" {
		return dbConnection
	}
	dbConnection = "" + RdbConfigMaster.User + ":" + RdbConfigMaster.Password + "@tcp(" + RdbConfigMaster.Address.Host + ":" + RdbConfigMaster.Address.Port + ")/" + RdbConfigMaster.Database + "?charset=" + RdbConfigMaster.MaxSetting.Charset + "&parseTime=True&loc=" + RdbConfigMaster.MaxSetting.Timezone
	return dbConnection
}

func InitSetConfigMaster(masterHost *string, masterPort *string, maserUser *string, masterPass *string) *RdbServerMaster {
	RdbHost = masterHost
	RdbPort = masterPort
	RdbUser = maserUser
	RdbPassword = masterPass
	return SetMasterRdbServerConfigFromEnv()
}

func SetMasterRdbServerConfigFromEnv() *RdbServerMaster {
	rdbConfigOnce.Do(func() {

		RdbConfigMaster = &RdbServerMaster{
			RdbType:       env.GetEnv("RDB_TYPE", configs.RdbDatabasesType),
			User:          env.GetEnv("RDB_USER", *RdbUser),
			Password:      env.GetEnv("RDB_PW", *RdbPassword),
			Database:      env.GetEnv("RDB_DBNAME", GetRdbMasterDatabaseName()),
			DbType:        CodeRdbConfigDatabaseTypeMaster,
			ConnectionStr: GetRdbClusterConnectionStr(),
			Address: AddressInfo{
				Host: env.GetEnv("RDB_HOST", *RdbHost),
				Port: env.GetEnv("RDB_PORT", *RdbPort),
			},
			MaxSetting: RdbCommonConfig{
				Timeout:            env.GetEnv("RDB_TIMEOUT", configs.RdbTimeout),
				MaxIdleConnections: env.GetEnv("RDB_MAX_IDLE_CNT", configs.RdbMaxIdleCnt),
				MaxOpenConnections: env.GetEnv("RDB_MAX_OPEN_CONN", configs.RdbMaxOpenConn),
				Charset:            env.GetEnv("RDB_CHARSET", configs.RdbCharset),
				Timezone:           env.GetEnv("RDB_TIMEZONE", configs.RdbTimezone),
			},
			DebugLevel: env.GetEnv("DEBUG_LEVEL", configs.RdbDebugLevel),
		}
	})
	return RdbConfigMaster
}
