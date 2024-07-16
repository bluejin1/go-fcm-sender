package dao

import (
	"fcm-sender/pkg/rdb/rdb_master"
	"fmt"
	"gorm.io/gorm"
)

var (
	GMaster *GormMasterModel
	//GLogModel    *GormLogModel
)

type GormMasterModel struct {
	Gorm *gorm.DB
}

//type GormLogModel struct {
//	Gorm *gorm.DB
//}

/*func (m *GormMasterModel) SetCryptor(encryptor res_helper.EncryptFunc, decryptor res_helper.DecryptFunc) *GormMasterModel {
	res_helper.SetCryptorFunc(encryptor, decryptor)
	return m
}*/

func NewMasterModel() *GormMasterModel {

	if rdb_master.RdbConnMaster == nil {
		rdbMaster := new(rdb_master.MasterDB)
		dbConn, err := rdbMaster.ConnectMaster()
		if err != nil {
			fmt.Printf("RdbConnCollection.ConnectCollection err %v", err)
			return nil
		}
		GMaster = &GormMasterModel{
			Gorm: dbConn,
		}
		return GMaster
	}

	if GMaster == nil {
		//logModel = new(LogModel)
		GMaster = &GormMasterModel{
			Gorm: rdb_master.RdbConnMaster.DbConn,
		}
	}
	return GMaster
}

func NewLogModel() *GormMasterModel {

	if rdb_master.RdbConnMaster == nil {
		rdbMaster := new(rdb_master.MasterDB)
		dbConn, err := rdbMaster.ConnectMaster()
		if err != nil {
			fmt.Printf("RdbConnCollection.ConnectCollection err %v", err)
			return nil
		}
		GMaster = &GormMasterModel{
			Gorm: dbConn,
		}
		return GMaster
	}

	if GMaster == nil {
		//logModel = new(LogModel)
		GMaster = &GormMasterModel{
			Gorm: rdb_master.RdbConnMaster.DbConn,
		}
	}
	return GMaster
}
