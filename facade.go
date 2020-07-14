package mysqlutils

import (
	"github.com/lyoshur/gorm/struct"
	"github.com/lyoshur/mysqlutils/info"
	"github.com/lyoshur/mysqlutils/mvc"
)

// 结构体
type DataBase = info.DataBase
type Holder = mvc.Holder

// 获取数据库相关信息
func GetDataBase(templateEngine *_struct.TemplateEngine, dataBaseName string) info.DataBase {
	return info.GetDataBase(templateEngine, dataBaseName)
}

// 获取Holder
func GetHolder(dataSourceName string) *mvc.Holder {
	return mvc.GetHolder(dataSourceName)
}

// 获取Mapper XML
func GetMapperXML(db info.DataBase, tableName string) string {
	return mvc.GetMapperXML(db, tableName)
}
