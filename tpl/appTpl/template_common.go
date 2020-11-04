package appTpl

const TplCommon  = `package utils

import (
	"github.com/moka-mrp/sword-core/kernel/server"
	ser "github.com/moka-mrp/sword/server"
	"reflect"
)

//判断某个结构体指针是否是nil
//todo 注意通过  d ==nil是比对不出来的额,这个只能用来比对非指针类型
//@author sam@2020-08-26 11:42:21
func IsNil(i interface{}) bool {
	if i ==nil{
		return  true
	}
	//fmt.Printf("%#v\r\n",i)
	vi := reflect.ValueOf(i)//通过反射获取其对应的值
	if vi.Kind() == reflect.Ptr || vi.Kind() == reflect.Slice || vi.Kind()==reflect.Array  {
		return vi.IsNil()
	}
	return false
}


//返回版本号
//@author sam@2020-10-16 17:20:25
func GetVersion() string {
	return  "Sword(v"+ser.Version+") (sword-core v"+server.Version+")"
}`

