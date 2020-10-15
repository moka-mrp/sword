package server

import (
	"errors"
	"fmt"
	"os"
	"path"
)


//创建新项目
//@author sam@2020-10-14 10:07:44
func RunNew(args []string) error {
	//判断是否传递了项目名
	if len(args) == 0 {
		return errors.New("required project name")
	}
	P.Name = args[0]

	//获取go mod名，未设置则与项目名一致
	if P.ModuleName == "" {
		P.ModuleName = P.Name
	}

	//创建的项目的存储目录，未设置就是当前目录
	if P.Path != "" {
		P.Path = path.Join(P.Path, P.Name)
	} else {
		pwd, _ := os.Getwd()
		P.Path = path.Join(pwd, P.Name)
	}

	//创建项目
	if err := create(); err != nil {
		return err
	}
	//fmt.Printf("%#v\r\n",p)
	fmt.Printf("Project: %s\n", P.Name)
	fmt.Printf("Module Name: %s\n", P.ModuleName)
	fmt.Printf("Directory: %s\n\n", P.Path)
	fmt.Println("Congratulations.The application has been created.")
	return nil
}
