package mvc

import (
	"github.com/lyoshur/gorm/table"
)

type DAO interface {
	GetList(attr interface{}) (table.Table, error)
	GetCount(attr interface{}) (int64, error)
	Exist(attr interface{}) (bool, error)
	GetModel(attr interface{}) (table.Table, error)
	Update(attr interface{}) (int64, error)
	Insert(attr interface{}) (int64, error)
	Delete(attr interface{}) (int64, error)
}

type daoImpl struct {
	tableManager *table.Manager
}

// 获取dao层
func GetDAO(tableManager *table.Manager) DAO {
	daoImpl := daoImpl{}
	daoImpl.tableManager = tableManager
	return &daoImpl
}

// 查询列表方法
func (daoImpl *daoImpl) GetList(attr interface{}) (table.Table, error) {
	return daoImpl.tableManager.Query("GetList", attr)
}

// 查询条数方法
func (daoImpl *daoImpl) GetCount(attr interface{}) (int64, error) {
	return daoImpl.tableManager.Count("GetCount", attr)
}

// 查询记录存在
func (daoImpl *daoImpl) Exist(attr interface{}) (bool, error) {
	b, err := daoImpl.tableManager.Count("Exist", attr)
	return b > 0, err
}

// 查询实体方法
func (daoImpl *daoImpl) GetModel(attr interface{}) (table.Table, error) {
	return daoImpl.tableManager.Query("GetModel", attr)
}

// 更新记录方法
func (daoImpl *daoImpl) Update(attr interface{}) (int64, error) {
	_, num, err := daoImpl.tableManager.Exec("Update", attr)
	return num, err
}

// 添加记录方法
func (daoImpl *daoImpl) Insert(attr interface{}) (int64, error) {
	id, _, err := daoImpl.tableManager.Exec("Insert", attr)
	return id, err
}

// 删除记录方法
func (daoImpl *daoImpl) Delete(attr interface{}) (int64, error) {
	_, num, err := daoImpl.tableManager.Exec("Delete", attr)
	return num, err
}
