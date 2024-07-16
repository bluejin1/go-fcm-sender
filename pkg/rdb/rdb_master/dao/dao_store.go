package dao

import "fcm-sender/pkg/rdb/rdb_master/models"

func (m *GormMasterModel) GetStoreInfoFull(storeIdx int32) *models.StoreInfoFull {
	if storeIdx < 1 {
		return nil
	}
	dmRows := &models.StoreInfoFull{}
	selectQuery := "kicc.*, store.*"
	joinQuery := "INNER JOIN " + models.TableNameMasterStoreKicc + " AS kicc ON store.store_idx = kicc.store_idx"
	result := m.Gorm.Table(models.TableNameMasterStore+" AS store").Select(selectQuery).Joins(joinQuery).Where("store.store_idx = ? ", storeIdx).Limit(1).Scan(&dmRows)
	if result.RowsAffected > 0 {
		return dmRows
	}
	return nil
}

func (m *GormMasterModel) GetStoreInfoFullWithKiccCode(kiccStoreCode string) *models.StoreInfoFull {
	if len(kiccStoreCode) < 1 {
		return nil
	}
	dmRows := &models.StoreInfoFull{}
	selectQuery := "kicc.*, store.*"
	joinQuery := "INNER JOIN " + models.TableNameMasterStore + " AS store ON store.store_idx = kicc.store_idx"
	result := m.Gorm.Table(models.TableNameMasterStoreKicc+" AS kicc").Select(selectQuery).Joins(joinQuery).Where("kicc.kicc_store_code = ? ", kiccStoreCode).Limit(1).Scan(&dmRows)
	if result.RowsAffected > 0 {
		return dmRows
	}
	return nil
}

func (m *GormMasterModel) GetStoreInfo(storeIdx int32) *models.Store {
	if storeIdx < 1 {
		return nil
	}
	whereStr := "store_idx = ?"
	dmRows := &models.Store{}
	result := m.Gorm.Model(models.Store{}).First(dmRows, whereStr, storeIdx)
	if result.RowsAffected > 0 {
		return dmRows
	}
	return nil
}

func (m *GormMasterModel) GetStoreInfoListAllShortWithKicc(limit int) *[]models.StoreInfoShort {

	dmRows := make([]models.StoreInfoShort, 0)
	whereStr := "store.store_status = ?"
	selectQuery := "kicc.*, store.*"
	joinQuery := "LEFT JOIN " + models.TableNameMasterStoreKicc + " AS kicc ON store.store_idx = kicc.store_idx"
	result := m.Gorm.Table(models.TableNameMasterStore+" AS store").Select(selectQuery).Joins(joinQuery).Where(whereStr, 1).Limit(limit).Scan(&dmRows)
	if result.RowsAffected > 0 {
		return &dmRows
	}
	return nil
}

func (m *GormMasterModel) GetStoreInfoShort(storeIdx int32) *models.StoreInfoShort {
	if storeIdx < 1 {
		return nil
	}
	dmRows := &models.StoreInfoShort{}
	selectQuery := "kicc.*, store.*"
	joinQuery := "INNER JOIN " + models.TableNameMasterStoreKicc + " AS kicc ON store.store_idx = kicc.store_idx"
	result := m.Gorm.Table(models.TableNameMasterStore+" AS store").Select(selectQuery).Joins(joinQuery).Where("store.store_idx = ? ", storeIdx).Limit(1).Scan(&dmRows)
	if result.RowsAffected > 0 {
		return dmRows
	}
	return nil
}

func (m *GormMasterModel) GetStoreInfoShortWithCode(storeCode string) *models.StoreInfoShort {
	if len(storeCode) < 1 {
		return nil
	}
	dmRows := &models.StoreInfoShort{}
	selectQuery := "kicc.*, store.*"
	joinQuery := "LEFT JOIN " + models.TableNameMasterStore + " AS store ON store.store_idx = kicc.store_idx"
	result := m.Gorm.Table(models.TableNameMasterStoreKicc+" AS kicc").Select(selectQuery).Joins(joinQuery).Where("kicc.kicc_store_code = ? ", storeCode).Limit(1).Scan(&dmRows)
	if result.RowsAffected > 0 {
		return dmRows
	}
	return nil
}
