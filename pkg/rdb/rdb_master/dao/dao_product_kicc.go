package dao

import "fcm-sender/pkg/rdb/rdb_master/models"

func (m *GormMasterModel) GetProductListWithKiccCodes(codes []string, limit int) *[]models.ProductInfoFull {
	if len(codes) < 1 {
		return nil
	}
	dmRows := make([]models.ProductInfoFull, 0)
	whereStr := "kicc.kicc_product_code IN ?"
	selectQuery := "kicc.*, product.*"
	joinQuery := "INNER JOIN " + models.TableNameMasterProduct + " AS product ON product.idx = kicc.product_idx"
	result := m.Gorm.Table(models.TableNameKiccProduct+" AS kicc").Select(selectQuery).Joins(joinQuery).Where(whereStr, codes).Limit(limit).Scan(&dmRows)
	if result.RowsAffected > 0 {
		return &dmRows
	}
	return nil

}
