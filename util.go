package mysqlplus

import (
	"reflect"
	"strings"

	"github.com/go-basic/uuid"
)

type utilInterface interface {
	GetColumns(val interface{}, tag string) ([]string, []interface{}) // 获取字段数组及设值数组
	GetAllTags(val interface{}, tag string) []string                  // 获取当前结构体对应标签数据
	IsBlank(value reflect.Value) bool                                 // 判断值是否为空
	UUID() string                                                     // 获取UUID
	UUIDReplaceAll(old, new string) string                            // UUIDReplaceAll 获取UUID并进行替换字符

}

type utilStruct struct{}

var Util utilInterface = (*utilStruct)(nil)

// GetColumns 基础内核方法 获取参数数据
// 仅返回数据非空的字段数据及数据结果
func (u *utilStruct) GetColumns(val interface{}, tag string) ([]string, []interface{}) {

	columns := make([]string, 0)
	params := make([]interface{}, 0)

	typeOf := reflect.TypeOf(val)
	valueOf := reflect.ValueOf(val)

	for i := 0; i < typeOf.NumField(); i++ {
		column := typeOf.Field(i).Tag.Get(tag)
		value := valueOf.Field(i).Interface()
		if !u.IsBlank(reflect.ValueOf(value)) {
			columns = append(columns, column)
			params = append(params, value)
		}
	}

	return columns, params
}

// GetAllTags 基础内核方法 获取参数标签数据
// 获取所有字段数据及数据结果
func (u *utilStruct) GetAllTags(val interface{}, tag string) []string {

	columns := make([]string, 0)

	typeOf := reflect.TypeOf(val)

	for i := 0; i < typeOf.NumField(); i++ {
		column := typeOf.Field(i).Tag.Get("db")
		columns = append(columns, column)
	}

	return columns
}

// IsBlank 判断值是否为空
func (u *utilStruct) IsBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

// UUID 获取UUID
func (u *utilStruct) UUID() string {
	return uuid.New()
}

// UUIDReplaceAll 获取UUID并进行替换字符
func (u *utilStruct) UUIDReplaceAll(old, new string) string {
	return strings.ReplaceAll(u.UUID(), old, new)
}
