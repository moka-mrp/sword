package baseTpl



const  TplEnv = `# toml配置文件
# Wiki：https://github.com/toml-lang/toml
### 全局配置信息
Debug = true
Env = "local"
ID ="{{.Name}}"
###日志配置信息
[Log]
Handler = "file"
Dir = "./logs"
Level = "debug"
Name = "sword"
EnableFileName=true
EnableFuncName=true

###Redis配置相关信息(字典对象)
[Redis]
[Redis.Default]
Host = "127.0.0.1"
Port =6379
DB = 0
MaxIdle = 64
Wait = false
MaxActive =10
IdleTimeout = 180
ConnectTimeout = 3
ReadTimeout = 10
WriteTimeout = 10
Password="root"

[Redis.Center]
Host = "127.0.0.1"
Port = 6379
DB = 1
Password="root"
MaxIdle = 64
Wait = false
MaxActive = 10
IdleTimeout = 180
ConnectTimeout = 3
ReadTimeout = 10
WriteTimeout = 10

###Db配置相关信息
[Db]
Driver = "mysql"

[Db.Option]
MaxOpenConns = 128
MaxIdleConns = 32
IdleTimeout = 180 # second
Charset = "utf8mb4"
ConnMaxLifetime=180
ConnectTimeout = 3 # second

[Db.Master]
Host = "127.0.0.1"
Port = 3306
User = "root"
Password = "root"
Name = "test"

###Web服务配置相关信息
[Api]
Host = "0.0.0.0"
Port = 9999








`
