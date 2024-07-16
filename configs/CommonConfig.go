package configs

import "fcm-sender/helper/env"

type AddressInfo struct {
	Host string // localhost,
	Port string
}

var (
	EnvServerType = env.GetEnv("SERVER_LIVE_TYPE", SERVER_LIVE_TYPE)
)
