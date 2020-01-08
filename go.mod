module github.com/simanchou/go-gin-example

go 1.13

require (
	github.com/astaxie/beego v1.12.0
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ini/ini v1.51.1
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.11 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/sys v0.0.0-20200106114638-5f8ca72cd632 // indirect
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
