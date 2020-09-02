package mvc

import (
	"github.com/lyoshur/gorm/generate"
	"github.com/lyoshur/mysqlutils/mvc/info"
)

const mapper = `
{{$ColumnNumber := (len .Table.Columns)}}
{{$key := ""}}{{range .Table.Indexs}}{{if eq .Name "PRIMARY"}}{{$key = .ColumnName}}{{end}}{{end}}
<xml>
    <sql>
        <key>GetList</key>
        <value>
            SELECT
                {{range $i, $v := .Table.Columns}}{{$.GraveAccent}}{{.Name}}{{$.GraveAccent}}{{if ne $v.Number $ColumnNumber}}, {{end}}{{end}}
            FROM
                {{$.GraveAccent}}{{.Table.Name}}{{$.GraveAccent}}
            WHERE
	{{range $i, $v := .Table.Columns}}
		{{if ne $v.Name $key}}
			{{ $t := (ClearType .Type) }}
			{{range $j, $c := $.IntTypes}}
				{{if eq $t $c}}
					{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
					{{$.LFlower}}{if ne .{{BigHump $v.Name}} 0}{{$.RFlower}} {{$.GraveAccent}}{{$v.Name}}{{$.GraveAccent}} = #{.{{BigHump $v.Name}}} AND {{$.LFlower}}{end}{{$.RFlower}}
					{{$.LFlower}}{end}{{$.RFlower}}
				{{end}}
			{{end}}
			{{range $j, $c := $.FloatTypes}}
				{{if eq $t $c}}
					{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
					{{$.LFlower}}{if ne .{{BigHump $v.Name}} 0}{{$.RFlower}} {{$.GraveAccent}}{{$v.Name}}{{$.GraveAccent}} = #{.{{BigHump $v.Name}}} AND {{$.LFlower}}{end}{{$.RFlower}}
					{{$.LFlower}}{end}{{$.RFlower}}
				{{end}}
			{{end}}
			{{range $j, $c := $.DateTypes}}
				{{if eq $t $c}}
					{{if eq $t "date"}}
						{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
						{{$.LFlower}}{if ne (.{{BigHump $v.Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} {{$v.Name}} = #{.{{BigHump $v.Name}}.Format "2006-01-02"} AND {{$.LFlower}}{end}{{$.RFlower}}
						{{$.LFlower}}{end}{{$.RFlower}}
					{{end}}
					{{if eq $t "time"}}
						{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
						{{$.LFlower}}{if ne (.{{BigHump $v.Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} {{$v.Name}} = #{.{{BigHump $v.Name}}.Format "15:04:05"} AND {{$.LFlower}}{end}{{$.RFlower}}
						{{$.LFlower}}{end}{{$.RFlower}}
					{{end}}
					{{if eq $t "year"}}
						{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
						{{$.LFlower}}{if ne (.{{BigHump $v.Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} {{$v.Name}} = #{.{{BigHump $v.Name}}.Format "2006"} AND {{$.LFlower}}{end}{{$.RFlower}}
						{{$.LFlower}}{end}{{$.RFlower}}
					{{end}}
					{{if eq $t "timestamp"}}
						{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
						{{$.LFlower}}{if ne (.{{BigHump $v.Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} {{$v.Name}} = #{.{{BigHump $v.Name}}.Format "20060102150405"} AND {{$.LFlower}}{end}{{$.RFlower}}
						{{$.LFlower}}{end}{{$.RFlower}}
					{{end}}
					{{if eq $t "datetime"}}
						{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
						{{$.LFlower}}{if ne (.{{BigHump $v.Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} {{$v.Name}} = #{.{{BigHump $v.Name}}.Format "2006-01-02 15:04:05"} AND {{$.LFlower}}{end}{{$.RFlower}}
						{{$.LFlower}}{end}{{$.RFlower}}
					{{end}}
				{{end}}
			{{end}}
			{{range $j, $c := $.StringTypes}}
				{{if eq $t $c}}
					{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}} {{$.GraveAccent}}{{$v.Name}}{{$.GraveAccent}} LIKE #{.{{BigHump $v.Name}}} AND {{$.LFlower}}{end}{{$.RFlower}}
				{{end}}
			{{end}}
		{{end}}
	{{end}}
            1 = 1
			{{$.LFlower}}{if .OrderBy}{{$.RFlower}}
				{{$.LFlower}}{if ne .OrderBy ""}{{$.RFlower}}
					ORDER BY {{$.LFlower}}{.OrderBy}{{$.RFlower}}
				{{$.LFlower}}{end}{{$.RFlower}}
			{{$.LFlower}}{end}{{$.RFlower}}
            {{$.LFlower}}{if .Start}{{$.RFlower}}
				{{$.LFlower}}{if ne .Start -1}{{$.RFlower}}
					LIMIT #{.Start},#{.Length}
				{{$.LFlower}}{end}{{$.RFlower}}
			{{$.LFlower}}{end}{{$.RFlower}}
        </value>
    </sql>
	<sql>
        <key>GetCount</key>
        <value>
            SELECT
                COUNT(1)
            FROM
                {{$.GraveAccent}}{{.Table.Name}}{{$.GraveAccent}}
            WHERE
	{{range $i, $v := .Table.Columns}}
		{{if ne $v.Name $key}}
			{{ $t := (ClearType .Type) }}
			{{range $j, $c := $.IntTypes}}
				{{if eq $t $c}}
					{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
					{{$.LFlower}}{if ne .{{BigHump $v.Name}} 0}{{$.RFlower}} {{$.GraveAccent}}{{$v.Name}}{{$.GraveAccent}} = #{.{{BigHump $v.Name}}} AND {{$.LFlower}}{end}{{$.RFlower}}
					{{$.LFlower}}{end}{{$.RFlower}}
				{{end}}
			{{end}}
			{{range $j, $c := $.FloatTypes}}
				{{if eq $t $c}}
					{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
					{{$.LFlower}}{if ne .{{BigHump $v.Name}} 0}{{$.RFlower}} {{$.GraveAccent}}{{$v.Name}}{{$.GraveAccent}} = #{.{{BigHump $v.Name}}} AND {{$.LFlower}}{end}{{$.RFlower}}
					{{$.LFlower}}{end}{{$.RFlower}}
				{{end}}
			{{end}}
			{{range $j, $c := $.DateTypes}}
				{{if eq $t $c}}
					{{if eq $t "date"}}
						{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
						{{$.LFlower}}{if ne (.{{BigHump $v.Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} {{$v.Name}} = #{.{{BigHump $v.Name}}.Format "2006-01-02"} AND {{$.LFlower}}{end}{{$.RFlower}}
						{{$.LFlower}}{end}{{$.RFlower}}
					{{end}}
					{{if eq $t "time"}}
						{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
						{{$.LFlower}}{if ne (.{{BigHump $v.Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} {{$v.Name}} = #{.{{BigHump $v.Name}}.Format "15:04:05"} AND {{$.LFlower}}{end}{{$.RFlower}}
						{{$.LFlower}}{end}{{$.RFlower}}
					{{end}}
					{{if eq $t "year"}}
						{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
						{{$.LFlower}}{if ne (.{{BigHump $v.Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} {{$v.Name}} = #{.{{BigHump $v.Name}}.Format "2006"} AND {{$.LFlower}}{end}{{$.RFlower}}
						{{$.LFlower}}{end}{{$.RFlower}}
					{{end}}
					{{if eq $t "timestamp"}}
						{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
						{{$.LFlower}}{if ne (.{{BigHump $v.Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} {{$v.Name}} = #{.{{BigHump $v.Name}}.Format "20060102150405"} AND {{$.LFlower}}{end}{{$.RFlower}}
						{{$.LFlower}}{end}{{$.RFlower}}
					{{end}}
					{{if eq $t "datetime"}}
						{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
						{{$.LFlower}}{if ne (.{{BigHump $v.Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} {{$v.Name}} = #{.{{BigHump $v.Name}}.Format "2006-01-02 15:04:05"} AND {{$.LFlower}}{end}{{$.RFlower}}
						{{$.LFlower}}{end}{{$.RFlower}}
					{{end}}
				{{end}}
			{{end}}
			{{range $j, $c := $.StringTypes}}
				{{if eq $t $c}}
					{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}} {{$.GraveAccent}}{{$v.Name}}{{$.GraveAccent}} LIKE #{.{{BigHump $v.Name}}} AND {{$.LFlower}}{end}{{$.RFlower}}
				{{end}}
			{{end}}
		{{end}}
	{{end}}
            1 = 1
        </value>
    </sql>
	<sql>
        <key>Exist</key>
        <value>
            SELECT
                COUNT(1)
            FROM
                {{$.GraveAccent}}{{.Table.Name}}{{$.GraveAccent}}
            WHERE
	{{range $i, $v := .Table.Columns}}
		{{if ne $v.Name $key}}
			{{ $t := (ClearType .Type) }}
			{{range $j, $c := $.IntTypes}}
				{{if eq $t $c}}
					{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
					{{$.LFlower}}{if ne .{{BigHump $v.Name}} 0}{{$.RFlower}} {{$.GraveAccent}}{{$v.Name}}{{$.GraveAccent}} = #{.{{BigHump $v.Name}}} AND {{$.LFlower}}{end}{{$.RFlower}}
					{{$.LFlower}}{end}{{$.RFlower}}
				{{end}}
			{{end}}
			{{range $j, $c := $.FloatTypes}}
				{{if eq $t $c}}
					{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
					{{$.LFlower}}{if ne .{{BigHump $v.Name}} 0}{{$.RFlower}} {{$.GraveAccent}}{{$v.Name}}{{$.GraveAccent}} = #{.{{BigHump $v.Name}}} AND {{$.LFlower}}{end}{{$.RFlower}}
					{{$.LFlower}}{end}{{$.RFlower}}
				{{end}}
			{{end}}
			{{range $j, $c := $.DateTypes}}
				{{if eq $t $c}}
					{{if eq $t "date"}}
						{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
						{{$.LFlower}}{if ne (.{{BigHump $v.Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} {{$v.Name}} = #{.{{BigHump $v.Name}}.Format "2006-01-02"} AND {{$.LFlower}}{end}{{$.RFlower}}
						{{$.LFlower}}{end}{{$.RFlower}}
					{{end}}
					{{if eq $t "time"}}
						{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
						{{$.LFlower}}{if ne (.{{BigHump $v.Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} {{$v.Name}} = #{.{{BigHump $v.Name}}.Format "15:04:05"} AND {{$.LFlower}}{end}{{$.RFlower}}
						{{$.LFlower}}{end}{{$.RFlower}}
					{{end}}
					{{if eq $t "year"}}
						{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
						{{$.LFlower}}{if ne (.{{BigHump $v.Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} {{$v.Name}} = #{.{{BigHump $v.Name}}.Format "2006"} AND {{$.LFlower}}{end}{{$.RFlower}}
						{{$.LFlower}}{end}{{$.RFlower}}
					{{end}}
					{{if eq $t "timestamp"}}
						{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
						{{$.LFlower}}{if ne (.{{BigHump $v.Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} {{$v.Name}} = #{.{{BigHump $v.Name}}.Format "20060102150405"} AND {{$.LFlower}}{end}{{$.RFlower}}
						{{$.LFlower}}{end}{{$.RFlower}}
					{{end}}
					{{if eq $t "datetime"}}
						{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
						{{$.LFlower}}{if ne (.{{BigHump $v.Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} {{$v.Name}} = #{.{{BigHump $v.Name}}.Format "2006-01-02 15:04:05"} AND {{$.LFlower}}{end}{{$.RFlower}}
						{{$.LFlower}}{end}{{$.RFlower}}
					{{end}}
				{{end}}
			{{end}}
			{{range $j, $c := $.StringTypes}}
				{{if eq $t $c}}
					{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}} {{$.GraveAccent}}{{$v.Name}}{{$.GraveAccent}} LIKE #{.{{BigHump $v.Name}}} AND {{$.LFlower}}{end}{{$.RFlower}}
				{{end}}
			{{end}}
		{{end}}
	{{end}}
            1 = 1
			LIMIT #{.Start},#{.Length}
        </value>
    </sql>
	<sql>
        <key>GetModel</key>
        <value>
            SELECT
                {{range $i, $v := .Table.Columns}}{{$.GraveAccent}}{{.Name}}{{$.GraveAccent}}{{if ne $v.Number $ColumnNumber}}, {{end}}{{end}}
            FROM
                {{$.GraveAccent}}{{.Table.Name}}{{$.GraveAccent}}
            WHERE
				{{$.GraveAccent}}{{$key}}{{$.GraveAccent}} = #{.}
        </value>
    </sql>
	<sql>
        <key>Update</key>
        <value>
            UPDATE {{$.GraveAccent}}{{.Table.Name}}{{$.GraveAccent}} SET
	{{range $i, $v := .Table.Columns}}
		{{if ne $v.Name $key}}
			{{ $t := (ClearType .Type) }}
			{{range $j, $c := $.IntTypes}}
				{{if eq $t $c}}
					{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
					{{$.LFlower}}{if ne .{{BigHump $v.Name}} 0}{{$.RFlower}} {{$.GraveAccent}}{{$v.Name}}{{$.GraveAccent}} = #{.{{BigHump $v.Name}}}, {{$.LFlower}}{end}{{$.RFlower}}
					{{$.LFlower}}{end}{{$.RFlower}}
				{{end}}
			{{end}}
			{{range $j, $c := $.FloatTypes}}
				{{if eq $t $c}}
					{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
					{{$.LFlower}}{if ne .{{BigHump $v.Name}} 0}{{$.RFlower}} {{$.GraveAccent}}{{$v.Name}}{{$.GraveAccent}} = #{.{{BigHump $v.Name}}}, {{$.LFlower}}{end}{{$.RFlower}}
					{{$.LFlower}}{end}{{$.RFlower}}
				{{end}}
			{{end}}
			{{range $j, $c := $.DateTypes}}
				{{if eq $t $c}}
					{{if eq $t "date"}}
						{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
						{{$.LFlower}}{if ne (.{{BigHump $v.Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} {{$v.Name}} = #{.{{BigHump $v.Name}}.Format "2006-01-02"}, {{$.LFlower}}{end}{{$.RFlower}}
						{{$.LFlower}}{end}{{$.RFlower}}
					{{end}}
					{{if eq $t "time"}}
						{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
						{{$.LFlower}}{if ne (.{{BigHump $v.Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} {{$v.Name}} = #{.{{BigHump $v.Name}}.Format "15:04:05"}, {{$.LFlower}}{end}{{$.RFlower}}
						{{$.LFlower}}{end}{{$.RFlower}}
					{{end}}
					{{if eq $t "year"}}
						{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
						{{$.LFlower}}{if ne (.{{BigHump $v.Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} {{$v.Name}} = #{.{{BigHump $v.Name}}.Format "2006"}, {{$.LFlower}}{end}{{$.RFlower}}
						{{$.LFlower}}{end}{{$.RFlower}}
					{{end}}
					{{if eq $t "timestamp"}}
						{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
						{{$.LFlower}}{if ne (.{{BigHump $v.Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} {{$v.Name}} = #{.{{BigHump $v.Name}}.Format "20060102150405"}, {{$.LFlower}}{end}{{$.RFlower}}
						{{$.LFlower}}{end}{{$.RFlower}}
					{{end}}
					{{if eq $t "datetime"}}
						{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
						{{$.LFlower}}{if ne (.{{BigHump $v.Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} {{$v.Name}} = #{.{{BigHump $v.Name}}.Format "2006-01-02 15:04:05"}, {{$.LFlower}}{end}{{$.RFlower}}
						{{$.LFlower}}{end}{{$.RFlower}}
					{{end}}
				{{end}}
			{{end}}
			{{range $j, $c := $.StringTypes}}
				{{if eq $t $c}}
					{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}} {{$.GraveAccent}}{{$v.Name}}{{$.GraveAccent}} = #{.{{BigHump $v.Name}}}, {{$.LFlower}}{end}{{$.RFlower}}
				{{end}}
			{{end}}
		{{end}}
	{{end}}
            WHERE {{$.GraveAccent}}{{$key}}{{$.GraveAccent}} = #{.{{BigHump $key}}}
        </value>
    </sql>
	<sql>
        <key>Insert</key>
        <value>
            INSERT INTO {{.Table.Name}}
			(
		{{range $i, $v := .Table.Columns}}
			{{if or (ne .Name $key) (eq $.Table.AutoIncrement -1)}}
				{{$.GraveAccent}}{{.Name}}{{$.GraveAccent}}{{if ne $v.Number $ColumnNumber}}, {{end}}
			{{end}}
		{{end}}
			)
            VALUES
			(
	{{range $i, $v := .Table.Columns}}
		{{if ne $v.Name $key}}
			{{ $t := (ClearType .Type) }}
			{{range $j, $c := $.IntTypes}}
				{{if eq $t $c}}
					{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
					{{$.LFlower}}{if ne .{{BigHump $v.Name}} 0}{{$.RFlower}} #{.{{BigHump $v.Name}}}, {{$.LFlower}}{end}{{$.RFlower}}
					{{$.LFlower}}{end}{{$.RFlower}}
				{{end}}
			{{end}}
			{{range $j, $c := $.FloatTypes}}
				{{if eq $t $c}}
					{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
					{{$.LFlower}}{if ne .{{BigHump $v.Name}} 0}{{$.RFlower}} #{.{{BigHump $v.Name}}}, {{$.LFlower}}{end}{{$.RFlower}}
					{{$.LFlower}}{end}{{$.RFlower}}
				{{end}}
			{{end}}
			{{range $j, $c := $.DateTypes}}
				{{if eq $t $c}}
					{{if eq $t "date"}}
						{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
						{{$.LFlower}}{if ne (.{{BigHump $v.Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} #{.{{BigHump $v.Name}}.Format "2006-01-02"}, {{$.LFlower}}{end}{{$.RFlower}}
						{{$.LFlower}}{end}{{$.RFlower}}
					{{end}}
					{{if eq $t "time"}}
						{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
						{{$.LFlower}}{if ne (.{{BigHump $v.Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} #{.{{BigHump $v.Name}}.Format "15:04:05"}, {{$.LFlower}}{end}{{$.RFlower}}
						{{$.LFlower}}{end}{{$.RFlower}}
					{{end}}
					{{if eq $t "year"}}
						{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
						{{$.LFlower}}{if ne (.{{BigHump $v.Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} #{.{{BigHump $v.Name}}.Format "2006"}, {{$.LFlower}}{end}{{$.RFlower}}
						{{$.LFlower}}{end}{{$.RFlower}}
					{{end}}
					{{if eq $t "timestamp"}}
						{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
						{{$.LFlower}}{if ne (.{{BigHump $v.Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} #{.{{BigHump $v.Name}}.Format "20060102150405"}, {{$.LFlower}}{end}{{$.RFlower}}
						{{$.LFlower}}{end}{{$.RFlower}}
					{{end}}
					{{if eq $t "datetime"}}
						{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}}
						{{$.LFlower}}{if ne (.{{BigHump $v.Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} #{.{{BigHump $v.Name}}.Format "2006-01-02 15:04:05"}, {{$.LFlower}}{end}{{$.RFlower}}
						{{$.LFlower}}{end}{{$.RFlower}}
					{{end}}
				{{end}}
			{{end}}
			{{range $j, $c := $.StringTypes}}
				{{if eq $t $c}}
					{{$.LFlower}}{if .{{BigHump $v.Name}}}{{$.RFlower}} #{.{{BigHump $v.Name}}}, {{$.LFlower}}{end}{{$.RFlower}}
				{{end}}
			{{end}}
		{{end}}
	{{end}}
			)
        </value>
    </sql>
    <sql>
        <key>Delete</key>
        <value>
            DELETE FROM {{$.GraveAccent}}{{.Table.Name}}{{$.GraveAccent}} WHERE {{$.GraveAccent}}{{$key}}{{$.GraveAccent}} = #{.}
        </value>
    </sql>
</xml>
`

func init() {
	// 初始化生成器
	builder = generate.GetBuilder(mapper)
}

// 生成器
var builder *generate.Builder

// 获取MapperXML
func GetMapperXML(db info.DataBase, tableName string) string {
	return builder.Execute(map[string]interface{}{
		"LFlower":      "{",
		"RFlower":      "}",
		"GraveAccent":  "`",
		"DataBaseName": db.Name,
		"Table":        db.GetTable(tableName),
		"IntTypes":     [...]string{"bit", "tinyint", "smallint", "mediumint", "int", "integer", "bigint"},
		"FloatTypes":   [...]string{"real", "double", "float", "decimal", "numeric"},
		"DateTypes":    [...]string{"date", "time", "year", "timestamp", "datetime"},
		"StringTypes":  [...]string{"char", "varchar", "text", "mediumtext", "longtext"},
	})
}
