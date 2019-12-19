module coder

go 1.12

require (
	github.com/Chronokeeper/anyxml v0.0.0-20160530174208-54457d8e98c6 // indirect
	github.com/CloudyKit/fastprinter v0.0.0-20170127035650-74b38d55f37a // indirect
	github.com/CloudyKit/jet v2.1.2+incompatible // indirect
	github.com/agrison/go-tablib v0.0.0-20160310143025-4930582c22ee // indirect
	github.com/agrison/mxj v0.0.0-20160310142625-1269f8afb3b4 // indirect
	github.com/bndr/gotabulate v1.1.2 // indirect
	github.com/clbanning/mxj v1.8.4 // indirect
	github.com/fatih/structs v1.1.0 // indirect

	github.com/gin-gonic/gin v1.4.0 // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/builder v0.0.0-20180826070321-377feedb49e3 // indirect
	github.com/go-xorm/core v0.0.0-20180322150003-0177c08cee88
	github.com/go-xorm/xorm v0.0.0-20180925133144-7a9249de3324
	github.com/jackc/pgx v3.5.0+incompatible // indirect
	github.com/jinzhu/gorm v1.9.10
	github.com/jmcvetta/randutil v0.0.0-20150817122601-2bb1b664bcff // indirect
	github.com/lunny/log v0.0.0-20160921050905-7887c61bf0de
	github.com/mitchellh/go-homedir v1.1.0
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/polaris1119/goutils v0.0.0-20190815094239-73c47df9b896
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	github.com/tealeg/xlsx v1.0.3 // indirect
	github.com/xormplus/builder v0.0.0-20190724032102-0ee351fedce9 // indirect
	github.com/xormplus/core v0.0.0-20190724072625-00f5a85ad6e0
	github.com/xormplus/xorm v0.0.0-20190827105519-9ce119234ab7
	gopkg.in/flosch/pongo2.v3 v3.0.0-20141028000813-5e81b817a0c4 // indirect
	gopkg.in/yaml.v2 v2.2.2
	xiaoshijie.com/go-xorm/core v0.0.0-00010101000000-000000000000 // indirect
	xiaoshijie.com/go-xorm/xorm v0.0.0-00010101000000-000000000000
	xiaoshijie.com/micro/common v0.0.0-00010101000000-000000000000 // indirect
)

replace xiaoshijie.com/go-xorm/xorm => ./customLib/go-xorm/xorm

replace xiaoshijie.com/go-xorm/core => ./customLib/go-xorm/core

replace xiaoshijie.com/micro/common => ./lanlan-library/micro-common
