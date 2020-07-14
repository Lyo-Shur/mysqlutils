package mvc

import (
	"github.com/lyoshur/gorm"
	"github.com/lyoshur/gorm/core"
	"github.com/lyoshur/gorm/struct"
	"github.com/lyoshur/gorm/table"
	"github.com/lyoshur/gorm/table/mvc"
	"github.com/lyoshur/mysqlutils/info"
	"strings"
)

type Holder struct {
	// 结构体 模板引擎
	StructTemplateEngine *_struct.TemplateEngine
	// table 模板引擎
	TableTemplateEngine *table.TemplateEngine
	// 数据库信息
	DataBase info.DataBase
	// Service Map
	Services map[string]mvc.Service
}

func GetHolder(dataSourceName string) *Holder {
	holder := Holder{}
	holder.Services = make(map[string]mvc.Service)
	// 解析数据库名
	databaseName := strings.Split(strings.Split(dataSourceName, "/")[1], "?")[0]

	// 获取结构体模板引擎
	structEngine := gorm.StructEngine{}
	structEngine.InitDB("mysql", dataSourceName)
	holder.StructTemplateEngine = gorm.GetStructTemplateEngine(&structEngine)

	// 获取Table模板引擎
	tableEngine := gorm.TableEngine{}
	tableEngine.InitDB("mysql", dataSourceName)
	holder.TableTemplateEngine = gorm.GetTableTemplateEngine(&tableEngine)

	// 临时关闭日志
	structEngine.SetLogger(&core.NoLogger{})
	tableEngine.SetLogger(&core.NoLogger{})

	// 获取默认数据库信息
	holder.DataBase = info.GetDataBase(holder.StructTemplateEngine, databaseName)

	// 遍历填充
	length := len(holder.DataBase.Tables)
	for i := 0; i < length; i++ {
		tableName := holder.DataBase.Tables[i].Name
		// 逐层组装Service
		mapper := GetMapperXML(holder.DataBase, tableName)
		manager := gorm.GetTableManagerByString(holder.TableTemplateEngine, mapper)
		dao := gorm.GetDAO(manager)
		holder.Services[tableName] = mvc.GetService(dao)
	}

	// 重新开启日志
	structEngine.SetLogger(&core.PrintLogger{})
	tableEngine.SetLogger(&core.PrintLogger{})

	return &holder
}
