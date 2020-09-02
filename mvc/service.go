package mvc

import (
	"github.com/lyoshur/gorm/table"
)

type Service interface {
	GetList(attr interface{}) (table.Table, error)
	GetCount(attr interface{}) (int64, error)
	Exist(attr interface{}) (bool, error)
	GetModel(attr interface{}) (table.Table, error)
	Update(attr interface{}) (int64, error)
	Insert(attr interface{}) (int64, error)
	Delete(attr interface{}) (int64, error)
}

type serviceImpl struct {
	dao DAO
}

// 获取service层
func GetService(dao DAO) Service {
	impl := serviceImpl{}
	impl.dao = dao
	return &impl
}

// 查询列表方法
func (serviceImpl *serviceImpl) GetList(attr interface{}) (table.Table, error) {
	return serviceImpl.dao.GetList(attr)
}

// 查询条数方法
func (serviceImpl *serviceImpl) GetCount(attr interface{}) (int64, error) {
	return serviceImpl.dao.GetCount(attr)
}

// 查询条数方法
func (serviceImpl *serviceImpl) Exist(attr interface{}) (bool, error) {
	return serviceImpl.dao.Exist(attr)
}

// 查询实体方法
func (serviceImpl *serviceImpl) GetModel(attr interface{}) (table.Table, error) {
	return serviceImpl.dao.GetModel(attr)
}

// 更新记录方法
func (serviceImpl *serviceImpl) Update(attr interface{}) (int64, error) {
	return serviceImpl.dao.Update(attr)
}

// 添加记录方法
func (serviceImpl *serviceImpl) Insert(attr interface{}) (int64, error) {
	return serviceImpl.dao.Insert(attr)
}

// 删除记录方法
func (serviceImpl *serviceImpl) Delete(attr interface{}) (int64, error) {
	return serviceImpl.dao.Delete(attr)
}
