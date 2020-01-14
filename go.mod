module github.com/simanchou/go-gin-example

go 1.13

require (
	github.com/EDDYCJY/go-gin-example v0.0.0-20191007083155-a98c25f2172a
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.0
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ini/ini v1.51.1
	github.com/go-openapi/spec v0.19.5 // indirect
	github.com/go-openapi/swag v0.19.6 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/gomodule/redigo v2.0.1-0.20180401191855-9352ab68be13+incompatible
	github.com/jinzhu/gorm v1.9.12
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mailru/easyjson v0.7.0 // indirect
	github.com/mattn/go-isatty v0.0.11 // indirect
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.5
	github.com/unknwon/com v1.0.1
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553 // indirect
	golang.org/x/sys v0.0.0-20200107162124-548cf772de50 // indirect
	golang.org/x/tools v0.0.0-20200113040837-eac381796e91 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/yaml.v2 v2.2.7 // indirect
)

replace (
	github.com/simachou/go-gin-example/conf => /home/siman/go-application/go-gin-example/pkg/conf
	github.com/simachou/go-gin-example/middleware => /home/siman/go-application/go-gin-example/middleware
	github.com/simachou/go-gin-example/models => /home/siman/go-application/go-gin-example/models
	github.com/simachou/go-gin-example/pkg/e => /home/siman/go-application/go-gin-example/pkg/e
	github.com/simachou/go-gin-example/pkg/setting => /home/siman/go-application/go-gin-example/pkg/setting
	github.com/simachou/go-gin-example/pkg/util => /home/siman/go-application/go-gin-example/pkg/util
	github.com/simachou/go-gin-example/routers => /home/siman/go-application/go-gin-example/routers
//	github.com/simanchou/go-gin-example/routers/api/v1 => /home/siman/go-application/go-gin-example/routers/api/v1
)
