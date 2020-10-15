package tpl

const TplGoMod = `module {{.ModuleName}}

go 1.13

require (
	github.com/coreos/etcd v3.3.13+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/golang/protobuf v1.3.3
	github.com/mitchellh/mapstructure v1.1.2
	github.com/moka-mrp/sword-core v0.1.1
	github.com/robfig/cron v1.2.0
	github.com/rogpeppe/fastuuid v1.2.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.7.1
	golang.org/x/crypto v0.0.0-20190605123033-f99c8df09eb5
	google.golang.org/grpc v1.21.1
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0
)
`