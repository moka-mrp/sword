package appTpl

const TplCommonTest=`package utils

import (
	"fmt"
	"testing"
)


type Student struct {
	Name  string
}
func fnil()(*Student){
	return  nil
}


func TestIsNil(t *testing.T){
	//1.未初始化的切片
	var slice01 []string
	fmt.Println(IsNil(slice01)) //true
	//2.空的切片
	slice02:=make([]string,0)
	fmt.Println(IsNil(slice02)) //false
	//3.结构体
	s:=Student{}
	fmt.Println(IsNil(s)) //false
	s2:=new(Student)
	fmt.Println(IsNil(s2)) //false
	s3:=&Student{}
	fmt.Println(IsNil(s3)) //false
	//4.空结构体指针
	s4:=fnil()
	fmt.Println(IsNil(s4)) //true
	//5.nil
	fmt.Println(IsNil(nil)) //true

}`