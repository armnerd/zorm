# Zorm

> 论如何攒一个 ORM

### Why

> 扯一个淡

* 学习一个技能的最佳实践就是教别人
* 了解一个事物原理的最佳实践就是自己造一个

### How

> 把大象装冰箱拢共需要几步

1. 打开连接
2. 执行操作
3. 关闭连接

### 实践

https://segmentfault.com/a/1190000021693989

1. 基于这篇文章的封装
2. 组装 sql
3. reflect 处理结果

### 反射

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

### Todo

* 事务，钩子，日志，join，tracing