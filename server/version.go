package server

import (
	"bytes"
	"runtime"
	"html/template"

)

var (
	Version = "1.0.0"
	BuildTime = "2020/01/01 14:15:16"
)

type VersionOptions struct {
	GitCommit string
	Version   string
	BuildTime string
	GoVersion string
	Os        string
	Arch      string
}

//版本输出模板
var versionTemplate = ` Version:      {{.Version}}
 Go version:   {{.GoVersion}}
 Built:        {{.BuildTime}}
 OS/Arch:      {{.Os}}/{{.Arch}}
 `

func GetVersion() string {
	var doc bytes.Buffer
	vo := VersionOptions{
		Version:   Version,
		BuildTime: BuildTime,
		GoVersion: runtime.Version(),
		Os:        runtime.GOOS,
		Arch:      runtime.GOARCH,
	}
	tmpl, _ := template.New("version").Parse(versionTemplate)
	tmpl.Execute(&doc, vo)
	return doc.String()
}


