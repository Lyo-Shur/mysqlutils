package session

//noinspection SpellCheckingInspection
import (
	"database/sql"
	"github.com/lyoshur/gorm"
	"github.com/lyoshur/gorm/core"
	gormsession "github.com/lyoshur/gorm/session"
	"github.com/lyoshur/gorm/struct"
	"github.com/lyoshur/gorm/table"
	"github.com/lyoshur/mysqlutils/mvc"
	"github.com/lyoshur/mysqlutils/mvc/info"
	"log"
	"strings"
)

// 会话工厂
type Factory struct {
	gormsession.Factory

	// 数据库信息
	DataBase info.DataBase
	// Service Map
	Services map[string]mvc.Service
}

// 语法糖+1
func (factory *Factory) GetService(serviceName string) mvc.Service {
	return factory.Services[serviceName]
}

// 会话工厂建造者
type FactoryBuilder struct {
	gormsession.FactoryBuilder

	// 数据库名
	databaseName string

	// 关闭Service
	isServiceClosed bool
}

// 获取会话工厂建造者
func GetSessionFactoryBuilder(driverName string, dataSourceName string) *FactoryBuilder {
	fb := FactoryBuilder{}
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	fb.DB = db
	fb.Logger = &core.PrintLogger{}
	fb.databaseName = strings.Split(strings.Split(dataSourceName, "/")[1], "?")[0]
	fb.isServiceClosed = false
	return &fb
}

// 设置日志
func (fb *FactoryBuilder) SetLogger(logger core.Logger) *FactoryBuilder {
	fb.Logger = logger
	return fb
}

// 关闭Service
func (fb *FactoryBuilder) CloseServices() *FactoryBuilder {
	fb.isServiceClosed = true
	return fb
}

// 建造
func (fb *FactoryBuilder) Build() *Factory {
	// 创建引擎
	tableEngine := table.Engine{}
	tableEngine.Init(fb.DB)

	structEngine := _struct.Engine{}
	structEngine.Init(fb.DB)

	// 组装session
	session := gormsession.Session{
		TableTemplateEngine:  table.GetTemplateEngine(&tableEngine),
		StructTemplateEngine: _struct.GetTemplateEngine(&structEngine),
	}

	// 组装会话工厂
	factory := &Factory{}
	factory.DB = fb.DB
	factory.Logger = fb.Logger
	factory.Session = session

	// 临时关闭日志
	structEngine.SetLogger(&core.NoLogger{})
	tableEngine.SetLogger(&core.NoLogger{})

	// 如果关闭了service 不加载下面
	if fb.isServiceClosed {
		return factory
	}

	factory.Services = make(map[string]mvc.Service)

	// 获取默认数据库信息
	factory.DataBase = info.GetDataBase(session.StructTemplateEngine, fb.databaseName)

	// 遍历填充
	length := len(factory.DataBase.Tables)
	for i := 0; i < length; i++ {
		tableName := factory.DataBase.Tables[i].Name
		// 逐层组装Service
		mapper := mvc.GetMapperXML(factory.DataBase, tableName)
		manager := gorm.GetTableManagerByString(session.TableTemplateEngine, mapper)
		dao := mvc.GetDAO(manager)
		factory.Services[tableName] = mvc.GetService(dao)
	}

	// 重新开启日志
	tableEngine.SetLogger(fb.Logger)
	structEngine.SetLogger(fb.Logger)

	return factory
}
