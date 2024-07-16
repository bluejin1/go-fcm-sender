package rdb_lib

import (
	"errors"
	"fcm-sender/pkg/rdb/rdb_config"
	"fcm-sender/pkg/rdb/rdb_master"
)

func ConnectMasterDatabases(masterHost *string, masterPort *string, maserUser *string, masterPass *string) (mdb *rdb_master.MasterDB, err error) {
	if rdb_config.IsUseRdbMasterDatabase() != true {
		return mdb, errors.New("master databases env off")
	}
	if rdb_master.IsRdbMasterInit && rdb_master.RdbConnMaster != nil {
		if rdb_master.RdbConnMaster.IsConnect {
			if rdb_master.RdbConnMaster.DbConn == nil {
				_, errCon := rdb_master.RdbConnMaster.DbConn.DB()
				if errCon == nil {
					return rdb_master.RdbConnMaster, err
				}
			} else {
				return rdb_master.RdbConnMaster, err
			}
		}
	}
	_, errConf := rdb_master.InitRdbMaster(masterHost, masterPort, maserUser, masterPass)
	if errConf != nil {
		return mdb, err
	}
	rdbMaster := new(rdb_master.MasterDB)
	if _, connErr := rdbMaster.ConnectMaster(); connErr != nil {
		return mdb, connErr
	}
	rdb_master.RdbConnMaster = rdbMaster
	return rdb_master.RdbConnMaster, nil
}

func HasConnectionMasterDatabase() bool {
	if rdb_config.IsUseRdbMasterDatabase() != true {
		return false
	}
	if rdb_master.IsRdbMasterInit != true || rdb_master.RdbConnMaster.IsConnect != true {
		return false
	}
	_, errCon := rdb_master.RdbConnMaster.DbConn.DB()
	if errCon != nil {
		return false
	}
	return true
}
