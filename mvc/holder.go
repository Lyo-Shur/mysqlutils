package mvc

import (
	"github.com/Lyo-Shur/gorm"
	"github.com/Lyo-Shur/gorm/core"
	"github.com/Lyo-Shur/gorm/struct"
	"github.com/Lyo-Shur/gorm/table"
	"github.com/Lyo-Shur/gorm/table/mvc"
	"github.com/Lyo-Shur/mysqlutils/info"
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
	utils := gorm.Gorm{}
	structEngine := utils.GetStructEngine()
	structEngine.InitDB("mysql", dataSourceName)
	structEngine.SetLogger(&core.NoLogger{})
	holder.StructTemplateEngine = utils.GetStructTemplateEngine(structEngine)

	// 获取Table模板引擎
	tableEngine := utils.GetTableEngine()
	tableEngine.InitDB("mysql", dataSourceName)
	tableEngine.SetLogger(&core.NoLogger{})
	holder.TableTemplateEngine = utils.GetTableTemplateEngine(tableEngine)

	// 获取默认数据库信息
	holder.DataBase = info.GetDataBase(holder.StructTemplateEngine, databaseName)

	// 遍历填充
	length := len(holder.DataBase.Tables)
	for i := 0; i < length; i++ {
		tableName := holder.DataBase.Tables[i].Name
		// 逐层组装Service
		mapper := GetMapperXML(holder.DataBase, tableName)
		manager := utils.GetTableManagerByString(holder.TableTemplateEngine, mapper)
		dao := utils.GetDAO(manager)
		holder.Services[tableName] = mvc.GetService(dao)
	}
	return &holder
}