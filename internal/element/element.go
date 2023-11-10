package element

// 插入元素
type InsertEle struct {
	Column string
	Value  string
}

// 更新元素
type UpdateEle struct {
	Column string
	Value  string
}

// 搜索元素
type WhereEle struct {
	Column    string
	Condition string
	Value     interface{}
}

// 排序元素
type OrderEle struct {
	Column   string
	Sequence string
}

// 表名
type Table interface {
	TableName() string
}
