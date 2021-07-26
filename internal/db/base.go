package db

// 会话
type Session struct {
	Sql         string
	SelectSet   []string
	WhereSet    []whereEle
	OrderSet    []orderEle
	OffsetIndex int
	LimitIndex  int
	InsertSet   []insertEle
	UpdateSet   []updateEle
}

// 插入元素
type insertEle struct {
	column string
	value  string
}

// 更新元素
type updateEle struct {
	column string
	value  string
}

// 搜索元素
type whereEle struct {
	column    string
	condition string
	value     interface{}
}

// 排序元素
type orderEle struct {
	column   string
	sequence string
}

// 表名
type table interface {
	TableName() string
}
