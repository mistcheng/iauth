package bootstrap

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
	"github.com/mistcheng/ilib/ijwt"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
)

func AppConfig()  {
	mysqlUser := beego.AppConfig.String("mysql.user")
	mysqlPass := beego.AppConfig.String("mysql.pass")
	mysqlHost := beego.AppConfig.String("mysql.host")
	mysqlDatabase := beego.AppConfig.String("mysql.database")
	ds := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", mysqlUser, mysqlPass, mysqlHost, mysqlDatabase)

	orm.RegisterDataBase("default", "mysql", ds)
	orm.DefaultTimeLoc = time.UTC
	orm.Debug = beego.AppConfig.DefaultBool("orm.debug", false)

	// config jwt
	ijwt.SetJWTMode(ijwt.JWTMode(beego.AppConfig.DefaultInt("jwt.mode", 0)))
	if jwtSecret := beego.AppConfig.String("jwt.secret"); len(jwtSecret) > 0 {
		ijwt.SetJWTSecret(jwtSecret)
	}

	if jwtPrivateKeyPath := beego.AppConfig.String("jwt.private_key_pem_path"); len(jwtPrivateKeyPath) > 0 {
		ijwt.SetJWTPrivateKey(jwtPrivateKeyPath)
	}

	if jwtPublicKeyPath := beego.AppConfig.String("jwt.public_key_pem_path"); len(jwtPublicKeyPath) > 0 {
		ijwt.SetJWTPublicKey(jwtPublicKeyPath)
	}

	// logs config
	logs.SetLogger(beego.AppConfig.String("log.driver"))
	logs.SetLevel(beego.AppConfig.DefaultInt("log.level", logs.LevelInfo))
}