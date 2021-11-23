package mysqlplus

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type mySqlPlusInterface interface {
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

type mySqlPlusStruct struct {
	isInit bool     // 是否初始化
	link   string   // 数据库连接字符串
	db     *sqlx.DB // 数据库连接对象
}

var MySqlPlus mySqlPlusInterface = (*mySqlPlusStruct)(nil)

// Create 创建数据库连接对象
func (mySqlPlus *mySqlPlusStruct) Create(link string) (*mySqlPlusStruct, error) {
	if nil == mySqlPlus || !mySqlPlus.isInit || mySqlPlus.link != link {
		db, err := sqlx.Open("mysql", link)
		if err != nil {
			return nil, err
		}
		mySqlPlus = &mySqlPlusStruct{
			db:     db,
			link:   link,
			isInit: true,
		}
		return mySqlPlus, nil
	}
	return mySqlPlus, nil
}

// SetMaxOpenConns 设置最大连接数
func (mySqlPlus *mySqlPlusStruct) SetMaxOpenConns(max int) {
	mySqlPlus.db.SetMaxOpenConns(max)
}

// SetMaxIdleConns 设置最大空闲数
func (mySqlPlus *mySqlPlusStruct) SetMaxIdleConns(max int) {
	mySqlPlus.db.SetMaxIdleConns(max)
}

// SetConnMaxLifetime 设置最大重连时间
func (mySqlPlus *mySqlPlusStruct) SetConnMaxLifetime(max time.Duration) {
	mySqlPlus.db.SetConnMaxLifetime(max * time.Nanosecond)
}

// SetConnMaxIdleTime 设置最大空闲时间
func (mySqlPlus *mySqlPlusStruct) SetConnMaxIdleTime(max time.Duration) {
	mySqlPlus.db.SetConnMaxIdleTime(max * time.Nanosecond)
}

// GetDB 获取数据库对象
func (mySqlPlus *mySqlPlusStruct) GetDB() *sqlx.DB {
	return mySqlPlus.db
}

// Select 执行Select查询方法
// result interface 接收结果接口
// model 数据模型
// table 表名
// where where条件语句
// whereParams 条件参数
func (mySqlPlus *mySqlPlusStruct) Select(result, model interface{}, table string, where string, whereParams ...interface{}) error {
	sql, params := Query.Select(table, where, model, whereParams...)
	return mySqlPlus.db.Select(result, sql, params...)
}

// Update 执行Update修改方法
// table 表名
// set 设置值的结构体
// where where条件语句
// whereParams 条件参数
func (mySqlPlus *mySqlPlusStruct) Update(table string, set interface{}, where string, whereParams ...interface{}) (sql.Result, error) {
	sql, params := Query.Update(table, set, where, whereParams...)
	return mySqlPlus.db.Exec(sql, params...)
}

// Delete 执行Delete删除方法
// table 表名
// where where条件语句
// whereParams 条件参数
func (mySqlPlus *mySqlPlusStruct) Delete(table string, where string, whereParams ...interface{}) (sql.Result, error) {
	sql, params := Query.Delete(table, where, whereParams...)
	return mySqlPlus.db.Exec(sql, params...)
}

// Insert 执行Insert增加方法
// table 表名
// val 设置值的结构体
func (mySqlPlus *mySqlPlusStruct) Insert(table string, val interface{}) (sql.Result, error) {
	sql, params := Query.Insert(table, val)
	return mySqlPlus.db.Exec(sql, params...)
}
