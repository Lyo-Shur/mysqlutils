package mysqlutils

import (
	"github.com/lyoshur/gorm/struct"
	"github.com/lyoshur/mysqlutils/mvc"
	"github.com/lyoshur/mysqlutils/mvc/info"
	"github.com/lyoshur/mysqlutils/session"
)

// 结构体
type DataBase = info.DataBase

// 获取数据库相关信息
//noinspection GoUnusedExportedFunction
func GetDataBaseInfo(templateEngine *_struct.TemplateEngine, dataBaseName string) info.DataBase {
	return info.GetDataBase(templateEngine, dataBaseName)
}

// 获取Mapper XML
//noinspection GoUnusedExportedFunction
func GetMapperXML(db info.DataBase, tableName string) string {
	return mvc.GetMapperXML(db, tableName)
}

type SessionFactory = session.Factory
type SessionFactoryBuilder = session.FactoryBuilder

// 获取会话工厂建造者
//noinspection GoUnusedExportedFunction
func GetSessionFactoryBuilder(dataSourceName string) *SessionFactoryBuilder {
	return session.GetSessionFactoryBuilder(dataSourceName)
}
