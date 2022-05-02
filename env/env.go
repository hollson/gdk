package env

import (
	"os"
	"strings"
	"sync"
)

// 应用版本环境变量
// 参考：https://en.wikipedia.org/wiki/Deployment_environment
const (
	Development = "development" // 开发版
	Alpha       = "alpha"       // 内测版
	Beta        = "beta"        // 公测版
	Staging     = "staging"     // 验收/审核版
	Production  = "production"  // 生产版
)

var (
	once       sync.Once
	currentEnv string
)

// Environment 应用环境(默认Key: SERVER_ENV)
func Environment(key ...string) string {
	once.Do(func() {
		k := "SERVER_ENV"
		if len(key) == 0 {
			k = key[0]
		}
		env := strings.ToLower(os.Getenv(k))
		switch env {
		case "alpha":
			currentEnv = Alpha
		case "beta":
			currentEnv = Beta
		case "stag", "stage", "staging":
			currentEnv = Staging
		case "prod", "produce", "production":
			currentEnv = Production
		default:
			currentEnv = Development
		}
	})
	return currentEnv
}
