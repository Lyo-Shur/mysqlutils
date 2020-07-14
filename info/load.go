package info

import (
	"github.com/lyoshur/gorm/core"
	"github.com/lyoshur/gorm/struct"
	"log"
)

// 表信息XML
const tableXML = `
<xml>
    <mapper column="TABLE_NAME" parameter="Name"/>
    <mapper column="ENGINE" parameter="Engine"/>
    <mapper column="TABLE_COLLATION" parameter="Collation"/>
    <mapper column="TABLE_COMMENT" parameter="Comment"/>
	<mapper column="AUTO_INCREMENT" parameter="AutoIncrement"/>
    <sql>
        <key>GetList</key>
        <value>
            SELECT
                TABLE_NAME, ENGINE, TABLE_COLLATION, TABLE_COMMENT, IFNULL(AUTO_INCREMENT, -1)
            FROM
                information_schema.TABLES
            WHERE
                TABLE_SCHEMA = #{.DataBaseName} AND TABLE_TYPE = 'BASE TABLE'
        </value>
    </sql>
</xml>
`

// 列信息XML
const columnXML = `
<xml>
    <mapper column="ORDINAL_POSITION" parameter="Number"/>
    <mapper column="COLUMN_NAME" parameter="Name"/>
    <mapper column="COLUMN_TYPE" parameter="Type"/>
    <mapper column="IS_NULLABLE" parameter="NullAble"/>
    <mapper column="COLUMN_DEFAULT" parameter="Defaule"/>
    <mapper column="COLUMN_COMMENT" parameter="Comment"/>
    <sql>
        <key>GetList</key>
        <value>
            SELECT
                ORDINAL_POSITION,
                COLUMN_NAME,
                COLUMN_TYPE,
                IS_NULLABLE,
                IFNULL(COLUMN_DEFAULT, ''),
                COLUMN_COMMENT
            FROM
                information_schema.COLUMNS
            WHERE
                TABLE_SCHEMA = #{.DataBaseName} AND TABLE_NAME = #{.TableName}
        </value>
    </sql>
</xml>
`

// 索引XML
const indexXML = `
<xml>
    <mapper column="INDEX_NAME" parameter="Name"/>
    <mapper column="COLUMN_NAME" parameter="ColumnName"/>
    <mapper column="NON_UNIQUE" parameter="Unique"/>
    <mapper column="INDEX_TYPE" parameter="Type"/>
    <sql>
        <key>GetList</key>
        <value>
            SELECT
                INDEX_NAME,
                COLUMN_NAME,
                NON_UNIQUE,
                INDEX_TYPE
            FROM
                information_schema.STATISTICS
            WHERE
                TABLE_SCHEMA = #{.DataBaseName} AND TABLE_NAME = #{.TableName}
        </value>
    </sql>
</xml>
`

// 外键XML
const keyXML = `
<xml>
    <mapper column="COLUMN_NAME" parameter="ColumnName"/>
    <mapper column="REFERENCED_TABLE_NAME" parameter="RelyTable"/>
    <mapper column="REFERENCED_COLUMN_NAME" parameter="RelyColumnName"/>
    <sql>
        <key>GetList</key>
        <value>
            SELECT
                COLUMN_NAME, REFERENCED_TABLE_NAME, REFERENCED_COLUMN_NAME
            FROM
                information_schema.KEY_COLUMN_USAGE
            WHERE
                CONSTRAINT_NAME != 'PRIMARY' AND
                TABLE_SCHEMA = REFERENCED_TABLE_SCHEMA AND
                TABLE_SCHEMA = #{.DataBaseName} AND TABLE_NAME = #{.TableName}
        </value>
    </sql>
</xml>
`

// 获取数据库相关信息
func GetDataBase(structTemplateEngine *_struct.TemplateEngine, dataBaseName string) DataBase {
	db := DataBase{}
	db.Name = dataBaseName
	// 查询数据
	db.SetTables(getTables(structTemplateEngine, db.Name))
	for i := 0; i < len(db.Tables); i++ {
		db.Tables[i].SetColumns(getColumns(structTemplateEngine, db.Name, db.Tables[i].Name))
		db.Tables[i].Indexs = getIndexs(structTemplateEngine, db.Name, db.Tables[i].Name)
		db.Tables[i].Keys = getKeys(structTemplateEngine, db.Name, db.Tables[i].Name)
	}
	return db
}

// 获取表相关信息
func getTables(templateEngine *_struct.TemplateEngine, dataBaseName string) []table {
	tableManager := _struct.GetManager(templateEngine, core.GetXmlConfig(tableXML), table{})
	vs, err := tableManager.Query("GetList", map[string]interface{}{
		"DataBaseName": dataBaseName,
	})
	if err != nil {
		log.Println(err)
	}
	// 转换列表
	l := len(vs)
	list := make([]table, l)
	for i := 0; i < l; i++ {
		list[i] = *vs[i].Interface().(*table)
	}
	return list
}

// 获取列相关信息
func getColumns(templateEngine *_struct.TemplateEngine, dataBaseName string, tableName string) []column {
	columnManager := _struct.GetManager(templateEngine, core.GetXmlConfig(columnXML), column{})
	vs, err := columnManager.Query("GetList", map[string]interface{}{
		"DataBaseName": dataBaseName,
		"TableName":    tableName,
	})
	if err != nil {
		log.Println(err)
	}
	// 转换列表
	l := len(vs)
	list := make([]column, l)
	for i := 0; i < l; i++ {
		list[i] = *vs[i].Interface().(*column)
	}
	return list
}

// 获取索引相关信息
func getIndexs(templateEngine *_struct.TemplateEngine, dataBaseName string, tableName string) []index {
	indexManager := _struct.GetManager(templateEngine, core.GetXmlConfig(indexXML), index{})
	vs, err := indexManager.Query("GetList", map[string]interface{}{
		"DataBaseName": dataBaseName,
		"TableName":    tableName,
	})
	if err != nil {
		log.Println(err)
	}
	// 转换列表
	l := len(vs)
	list := make([]index, l)
	for i := 0; i < l; i++ {
		list[i] = *vs[i].Interface().(*index)
	}
	return list
}

// 获取外键相关信息
func getKeys(templateEngine *_struct.TemplateEngine, dataBaseName string, tableName string) []key {
	keyManager := _struct.GetManager(templateEngine, core.GetXmlConfig(keyXML), key{})
	vs, err := keyManager.Query("GetList", map[string]interface{}{
		"DataBaseName": dataBaseName,
		"TableName":    tableName,
	})
	if err != nil {
		log.Println(err)
	}
	// 转换列表
	l := len(vs)
	list := make([]key, l)
	for i := 0; i < l; i++ {
		list[i] = *vs[i].Interface().(*key)
	}
	return list
}
