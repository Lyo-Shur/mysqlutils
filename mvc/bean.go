/*
此文件主要是一个查询使用的bean
负责查询的实体，可选择组合Query结构体，来拥有排序和分页查询功能
*/
package mvc

// 不使用分页
var closePaging int64 = -1

// 查询结构体
type Query struct {
	// 排序
	OrderBy string
	// 记录开始位置
	Start int64
	// 查询长度
	Length int64
}

// 查询结构体初始化方法
func (query *Query) Init() {
	query.OrderBy = ""
	query.Start = closePaging
	query.Length = 10
}
