package cmd

import (
	"fcm-sender/pkg/rdb/rdb_master/lib"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"fcm-sender/cmd/mq"
	"fcm-sender/configs"
	"fcm-sender/helper/zlog"
	"fcm-sender/pkg/rdb"
	"fcm-sender/pkg/rdb/rdb_config"
	//"firebase.google.com/go/v4/messaging"
)

var (
	loggerConfig     configs.Logger
	mainServerConfig configs.MainServer
)

// Firebase APP 초기화
func init() {

	fmt.Printf(">>> init <<<")

	loggerConfig = configs.DefaultLoggerConfigFromEnv()
	zlog.Init(loggerConfig.Level, loggerConfig.LogStyle)

	mainServerConfig = configs.DefaultMainServerConfigFromEnv()
	// database set
	rdb.InitRdb()

	initLibrary()
}

func initLibrary() {
	zlog.Info().Msgf(">>> initLibrary <<<")

	err := configs.FcmInit()
	if err != nil {
		zlog.Error().Msgf("[main init] setConfig error :", err)
	}

	// store_rdb connect
	err = setConnectionServiceRdb()
	if err != nil {
		zlog.Error().Msgf("setConnectionServiceRdb err %v", err)
	}

}

func waitSignal(signals chan os.Signal) {
	s := <-signals
	zlog.Info().Msgf("Got System signal: %v", s)
	shutdown()
}

func shutdown() {
	zlog.Info().Msgf(">>> shutdown start <<<")

	if rdb.MasterRdbLoaded {
		rdb.MasterDbClose()
		zlog.Info().Msgf("MasterDbClose")
	}

	// mq consumer close
	/*if consumer.ConsumerRmq != nil {
		err := consumer.ConsumerRmqClose()
		if err != nil {
			zlog.Error().Msgf("ConsumerRmqClose error %v", err)
		}
	}*/

}

func setConnectionServiceRdb() error {
	zlog.Info().Msgf(">>> setConnectionServiceRdb <<")

	var rdbErr error = nil
	if rdb_config.IsUseRdbMasterDatabase() {
		_, mErr := rdb.ConnectMaster()
		if mErr != nil {
			zlog.Error().Msgf("store_rdb ConnectMaster err %v", mErr)
			rdbErr = mErr
		} else {
			// gorm model load - only master
			rdb.LoadMasterGormModel()
		}
		if rdb.MasterRdbLoaded {
			zlog.Info().Msgf("store_rdb.MasterRdbLoaded success", rdb.MasterRdbLoaded)

			isLoaded := lib.SetStoreListMemory()
			zlog.Info().Msgf("SetStoreListMemory isLoaded %c", isLoaded)

		} else {
			zlog.Info().Msgf("store_rdb.MasterRdbLoaded false", rdb.MasterRdbLoaded)
		}

	}

	return rdbErr
}

func OrderServerStart(commit string, buildTime string) {

	zlog.Info().Msgf("ConfigServerStart: %v (GitCommitShortHash: %v, BuildTime: %v)", mainServerConfig.Name, commit, buildTime)
	zlog.Info().Msgf("Logger Config: %v", loggerConfig)
	zlog.Info().Msgf("MainServer : %v", mainServerConfig)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)
	go waitSignal(sigs)

	// pos order receiver
	go func() {
		err := mq.CmdReceiverKiccOrder()
		if err != nil {
			if err != nil {
				zlog.Error().Msgf("CmdReceiverKiccOrder error %v", err)
			}
		}
	}()

	// store config
	go func() {
		err := mq.CmdReceiverStoreConfigChangeInfo()
		if err != nil {
			if err != nil {
				zlog.Error().Msgf("CmdReceiverStoreConfigChangeInfo error %v", err)
			}
		}
	}()

	// app order receiver
	err := mq.CmdReceiverAppOrder()
	if err != nil {
		if err != nil {
			zlog.Error().Msgf("CmdReceiverAppOrder error %v", err)
		}
	}

}
