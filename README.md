# mysql_plus

#### 介绍

mysql_plus 自动生成mysql单表操作语句的封装式框架

核心技术以结构体反射为主

此框架内核采用sqlx,故数据库结构体tag采用db

封装Util结构体方法,基础处理UUID

core.go 文件全面封装sql自动生成处理逻辑, 以反射方式生成insert、update、select、delete语句

root.go 基础封装sqlx并内置db,以调用core.go中核心自动生成sql语句方法进行基础sql语句调用

#### 测试(TEST)

本框架测试用例基本全面封装于default_test.go方法之中

数据库采用mysql, 测试用库名为: mw, 测试用表为: tb_dict_info

sql语句为:
```
create table tb_dict_info
(
    id             varchar(32)            not null comment '字典id'
        primary key,
    dict_name      varchar(64) default '' not null comment '字典名称',
    dict_key       varchar(32) default '' null comment '字典键',
    dict_value     varchar(64) default '' null comment '字典值',
    dict_before_id varchar(32)            null comment '上级字典id',
    create_time    bigint      default 0  null comment '创建时间',
    update_time    bigint      default 0  null comment '修改时间',
    create_user    varchar(32) default '' null comment '创建者',
    status         int         default 1  null comment '字典状态 1 正常 2 停用 3 删除'
)
    comment '字典信息表';
```

#### 使用(USING)
```
mysqlplus.Query{
    Insert(table string, val interface{}) (string, []interface{})                                           // Insert 自动生成Insert语句方法
	Update(table string, set interface{}, where string, whereParams ...interface{}) (string, []interface{}) // Update 自动生成Update语句方法
	Delete(table, where string, whereParams ...interface{}) (string, []interface{})                         // Delete 自动生成Delete语句方法
	Select(table, where string, model interface{}, whereParams ...interface{}) (string, []interface{})      // Select 自动生成Select语句方法
}
```
```
mysqlplus.MySqlPlus{
    Create(link string) (*mySqlPlusStruct, error)                                                       // 创建数据库连接对象
	SetMaxOpenConns(max int)                                                                            // 设置最大连接数
	SetMaxIdleConns(max int)                                                                            // 设置最大空闲数
	SetConnMaxLifetime(max time.Duration)                                                               // 设置最大重连时间
	SetConnMaxIdleTime(max time.Duration)                                                               // 设置最大空闲时间
	GetDB() *sqlx.DB                                                                                    // 获取数据库对象
	Select(result, model interface{}, table string, where string, whereParams ...interface{}) error     // 自动生成查询语句并进行查询操作
	Update(table string, set interface{}, where string, whereParams ...interface{}) (sql.Result, error) // 自动生成修改语句并执行修改操作
	Delete(table string, where string, whereParams ...interface{}) (sql.Result, error)                  // 自动生成删除语句并执行删除操作
	Insert(table string, val interface{}) (sql.Result, error)                                           // 自动生成增加语句并执行增加操作
}
```