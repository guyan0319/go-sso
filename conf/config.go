package conf

import (
	"sync"
)

type Config struct {
	Language string
	Token string
	Super string
	RedisPre string
	Host string
	OpenJwt bool
	Routes []string
}
var (
	Cfg  Config
	mutex   sync.Mutex
	declare sync.Once
)

func  Set(cfg Config) {
	mutex.Lock()
	Cfg.RedisPre=setDefault(cfg.RedisPre,"","go.sso.redis")
	Cfg.Language=setDefault(cfg.Language,"","cn")
	Cfg.Token=setDefault(cfg.Token,"","token")
	Cfg.Super=setDefault(cfg.Super,"","admin")//超级账户
	Cfg.Host=setDefault(cfg.Host,"","http://localhost:8282")//域名
	Cfg.Routes=cfg.Routes
	Cfg.OpenJwt=cfg.OpenJwt
	mutex.Unlock()
}
func setDefault( value,def ,defValue string) string {
	if value==def {
		return defValue
	}
	return value
}
