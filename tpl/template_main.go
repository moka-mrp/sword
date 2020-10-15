package tpl

const TplMain = `package main


import "{{.ModuleName}}/cmd"

//亮剑入口
//@author sam@2020-07-27 09:13:28
func main() {
	cmd.Execute()
}
`