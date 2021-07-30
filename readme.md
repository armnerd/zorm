# Zorm

> 如何攒一个 ORM

### Why ?

> 扯淡在先

* 看不懂别人的 ORM，所以自己写一个
* 哪里有痛点，哪里就有轮子
* 学习一个技能的最佳实践就是教别人
* 了解一个事物原理的最佳实践就是自己造一个

### How ?

> 把大象装冰箱拢共需要几步 ?

* 打开连接
* 执行操作
* 关闭连接

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

> 多个查询如何避免冲突 ?

### Todo

* 事务，钩子，日志，join，tracing