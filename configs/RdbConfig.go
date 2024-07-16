package configs

import (
	"fcm-sender/helper/env"
)

type RdbCommonConfig struct {
	RdbType            string
	Timeout            string
	MaxIdleConnections string
	MaxOpenConnections string
	Charset            string
	Timezone           string
	User               string
	Password           string
	Database           string
	ConnectionStr      string
	DbType             string
	Host               string // localhost,
	Port               string
}

func IsUseRdbMasterDatabase() bool {
	useMasterDatabase := env.GetEnv("RDB_USE_MASTER_DB", RDB_USE_MASTER_DB)
	if useMasterDatabase == "true" {
		return true
	}
	return false
}

var rdbConfigMaster *RdbCommonConfig = nil

func DefaultMasterRdbServerConfigFromEnv() *RdbCommonConfig {

	if rdbConfigMaster != nil {
		return rdbConfigMaster
	}
	configMaster := RdbCommonConfig{
		RdbType:            env.GetEnv("RDB_TYPE", RDB_TYPE),
		User:               env.GetEnv("RDB_USER", RDB_USER),
		Password:           env.GetEnv("RDB_PW", RDB_PW),
		Database:           env.GetEnv("RDB_DBNAME", RDB_DBNAME),
		DbType:             CodeRdbConfigDatabaseTypeMaster,
		Timeout:            env.GetEnv("RDB_TIMEOUT", RDB_TIMEOUT),
		MaxIdleConnections: env.GetEnv("RDB_MAX_IDLE_CNT", RDB_MAX_IDLE_CNT),
		MaxOpenConnections: env.GetEnv("RDB_MAX_OPEN_CONN", RDB_MAX_OPEN_CONN),
		Charset:            env.GetEnv("RDB_CHARSET", RDB_CHARSET),
		Timezone:           env.GetEnv("RDB_TIMEZONE", RDB_TIMEZONE),
		Host:               env.GetEnv("RDB_HOST", RDB_HOST),
		Port:               env.GetEnv("RDB_PORT", RDB_PORT),
	}

	if env.GetEnvEnvironment() == "development" {
		configMaster = RdbCommonConfig{
			RdbType:            env.GetEnv("RDB_TYPE", RDB_TYPE),
			User:               env.GetEnv("RDB_USER", RDB_USER_TEST),
			Password:           env.GetEnv("RDB_PW", RDB_PW_TEST),
			Database:           env.GetEnv("RDB_DBNAME", RDB_DBNAME_TEST),
			DbType:             CodeRdbConfigDatabaseTypeMaster,
			Timeout:            env.GetEnv("RDB_TIMEOUT", RDB_TIMEOUT),
			MaxIdleConnections: env.GetEnv("RDB_MAX_IDLE_CNT", RDB_MAX_IDLE_CNT),
			MaxOpenConnections: env.GetEnv("RDB_MAX_OPEN_CONN", RDB_MAX_OPEN_CONN),
			Charset:            env.GetEnv("RDB_CHARSET", RDB_CHARSET),
			Timezone:           env.GetEnv("RDB_TIMEZONE", RDB_TIMEZONE),
			Host:               env.GetEnv("RDB_HOST", RDB_HOST_TEST),
			Port:               env.GetEnv("RDB_PORT", RDB_PORT_TEST),
		}
	}
	rdbConfigMaster = &configMaster
	return rdbConfigMaster
}
