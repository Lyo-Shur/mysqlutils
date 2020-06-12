package mysqlutils

import (
	"github.com/Lyo-Shur/gorm"
	"github.com/Lyo-Shur/gorm/struct"
	"github.com/Lyo-Shur/mysqlutils/info"
	"github.com/Lyo-Shur/mysqlutils/mvc"
)

// 查询引擎
type MysqlUtils struct {
	gorm.Gorm
}

// 获取数据库相关信息
func (mysqlUtils MysqlUtils) GetDataBase(templateEngine *_struct.TemplateEngine, dataBaseName string) info.DataBase {
	return info.GetDataBase(templateEngine, dataBaseName)
}

// 获取Mapper XML
func (mysqlUtils MysqlUtils) GetMapperXML(db info.DataBase, tableName string) string {
	return mvc.GetMapperXML(db, tableName)
}
