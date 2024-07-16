package rdb

//type GormLogModel struct {
//	Gorm *gorm.DB
//}

/*func (m *GormMasterModel) SetCryptor(encryptor res_helper.EncryptFunc, decryptor res_helper.DecryptFunc) *GormMasterModel {
	res_helper.SetCryptorFunc(encryptor, decryptor)
	return m
}*/
/*
func NewMasterModel() *rdb_dao.GormMasterModel {

	if RdbConnMaster == nil {
		rdbMaster := new(MasterDB)
		dbConn, err := rdbMaster.ConnectMaster()
		if err != nil {
			fmt.Printf("RdbConnCollection.ConnectCollection err %v", err)
			return nil
		}
		rdb_dao.GMaster = &rdb_dao.GormMasterModel{
			Gorm: dbConn,
		}
		return rdb_dao.GMaster
	}

	if rdb_dao.GMaster == nil {
		//logModel = new(LogModel)
		rdb_dao.GMaster = &rdb_dao.GormMasterModel{
			Gorm: RdbConnMaster.DbConn,
		}
	}
	return rdb_dao.GMaster
}

func NewLogModel() *rdb_dao.GormMasterModel {

	if RdbConnMaster == nil {
		rdbMaster := new(MasterDB)
		dbConn, err := rdbMaster.ConnectMaster()
		if err != nil {
			fmt.Printf("RdbConnCollection.ConnectCollection err %v", err)
			return nil
		}
		rdb_dao.GMaster = &rdb_dao.GormMasterModel{
			Gorm: dbConn,
		}
		return rdb_dao.GMaster
	}

	if rdb_dao.GMaster == nil {
		//logModel = new(LogModel)
		rdb_dao.GMaster = &rdb_dao.GormMasterModel{
			Gorm: RdbConnMaster.DbConn,
		}
	}
	return rdb_dao.GMaster
}
*/
