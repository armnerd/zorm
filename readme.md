# Zorm

> 论如何攒一个 ORM

### Why ?

> 扯淡在先

* 学习一个技能的最佳实践就是教别人
* 了解一个事物原理的最佳实践就是自己造一个

### How ?

> 把大象装冰箱拢共需要几步 ?

* 打开连接
* 执行操作
* 关闭连接

> 其实呢 ?

https://segmentfault.com/a/1190000021693989

* 就是基于这篇文章的封装
* 组装 sql
* reflect 处理结果

> 我想链式查询 ?

* 一个充满了各种 sql 元素的 struct

```go
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
```

* return self

```go
func (s *Session) Select(field []string) *Session {
	s.SelectSet = append(s.SelectSet, field...)
	return s
}
```

> 我怎么知道是哪个表 ?

* 一个有获取表名方法的 interface

```go
type table interface {
	TableName() string
}
```

* 创建一个表 struct, 并实现这个接口
* tag 可以通过反射获取字段名以及 Json 序列化时用

```go
type Demo struct {
	Id   int    `json:"id"` 
	Name string `json:"name"`
}

func (d Demo) TableName() string {
	return "demo"
}
```

* 获取表名

```go
func (s *Session) getSqlForBalabala (table table) {
	tableName := table.TableName()
}
```

> 查询出来怎么向上返回 ?

* 是时候使用反射了

```go
// 获取表 struct 的 type
tableType := reflect.ValueOf(table).Type()

// 搞一个新实例出来
newElement := reflect.New(tableType).Elem()

// 获取成员数量
newElement.NumField()

// 获取每个成员的内存地址
newElement.Field(index).Addr().Interface()

// 获取每个成员的数据类型
newElement.Field(index).Kind().String()

// 获取每个成员的值
newElement.Field(index)

// 获取表 struct 的 tag
tableInfo := reflect.TypeOf(table)
tableInfo.Field(index).Tag.Get("json")
```

> 多个操作如何避免冲突 ?

* 起一个新的 session 实例

```go
session := db.Session{}
```

* 每个 session 实例都会在操作后把用到的 sql 元素初始化

```go
func (s *Session) cleanUpForSelect() {
	s.Sql = ""
	s.SelectSet = []string{}
	s.WhereSet = []whereEle{}
	s.OrderSet = []orderEle{}
	s.OffsetIndex = 0
	s.LimitIndex = 0
}
```

### Todo

* 事务，钩子，日志，join，tracing