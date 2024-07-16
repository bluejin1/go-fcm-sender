package dao

import (
	"errors"
	"fcm-sender/pkg/rdb/rdb_master/models"
)

func (m *GormMasterModel) GetMasterAppOrder(orderNo *string) *models.AppOrder {
	if orderNo == nil || *orderNo == "" {
		return nil
	}
	whereStr := "orderNo = ?"
	dmRows := &models.AppOrder{}
	result := m.Gorm.Model(models.AppOrder{}).First(dmRows, whereStr, *orderNo)
	if result.RowsAffected > 0 {
		return dmRows
	}
	return nil
}

func (m *GormMasterModel) UpdateAppOrderApp(orderNo *string, updateData *models.AppOrder) error {
	if orderNo == nil || *orderNo == "" {
		return errors.New("empty orderNo")
	}
	result := m.Gorm.Model(models.AppOrder{}).Where(
		"orderNo = ?", *orderNo).Updates(updateData)
	if result.RowsAffected > 0 {
		return nil
	}
	return result.Error
}

func (m *GormMasterModel) UpdateAppOrderAppWithMap(orderNo *string, updateData *map[string]interface{}) (error, int64) {
	if orderNo == nil || *orderNo == "" {
		return errors.New("empty orderNo"), 0
	}
	result := m.Gorm.Model(models.AppOrder{}).Where(
		"orderNo = ?", *orderNo).Updates(*updateData)
	if result.RowsAffected > 0 {
		return nil, result.RowsAffected
	}
	if result.RowsAffected == 0 {
		return nil, result.RowsAffected
	}
	return result.Error, result.RowsAffected
}

func (m *GormMasterModel) InsertAppOrderAppWithMap(orderNo *string, insertData *map[string]interface{}) error {
	data := *insertData
	data["orderNo"] = orderNo
	result := m.Gorm.Model(models.AppOrder{}).Create(data)
	if result.RowsAffected > 0 {
		return nil
	}
	return result.Error
}

func (m *GormMasterModel) DeleteAppOrderApp(orderNo *string) error {
	if orderNo == nil || *orderNo == "" {
		return errors.New("empty orderNo")
	}
	result := m.Gorm.Model(&models.AppOrder{}).Delete(&models.AppOrder{}, "orderNo = ?", *orderNo)
	if result.RowsAffected > 0 {
		return nil
	}
	return result.Error
}

func (m *GormMasterModel) CreateAppOrderApp(createData *models.AppOrder) (res string, err error) {
	if createData.OrderNo == "" {
		return res, errors.New("empty orderNo")
	}
	result := m.Gorm.Create(createData)
	if result.RowsAffected > 0 && result.Error == nil {
		return createData.OrderNo, nil
	}
	return res, result.Error
}

func (m *GormMasterModel) CreateAppOrderAppWithMap(createData map[string]interface{}) (res *map[string]interface{}, err error) {
	if createData["orderNo"] == "" {
		return res, errors.New("empty orderNo")
	}
	result := m.Gorm.Model(models.AppOrder{}).Create(createData)
	if result.RowsAffected > 0 && result.Error == nil {
		return &createData, nil
	}
	return res, result.Error
}
