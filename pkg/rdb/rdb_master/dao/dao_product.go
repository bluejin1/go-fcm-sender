package dao

import (
	"fcm-sender/pkg/rdb/rdb_master/models"
)

func (m *GormMasterModel) GetProductInfo(productNo int32) *models.Product {
	if productNo < 1 {
		return nil
	}
	whereStr := "idx = ?"
	dmRows := &models.Product{}
	result := m.Gorm.Model(models.AppOrder{}).First(dmRows, whereStr, productNo)
	if result.RowsAffected > 0 {
		return dmRows
	}
	return nil
}

func (m *GormMasterModel) GetProductInfoListProductNos(productNos []int32) *[]models.Product {
	if len(productNos) < 1 {
		return nil
	}
	dmRows := make([]models.Product, 0)
	m.Gorm.Table(models.TableNameMasterProduct).Where("idx IN ? ", productNos).Limit(500).Scan(&dmRows)
	if len(dmRows) > 0 {
		return &dmRows
	}
	return nil
}
